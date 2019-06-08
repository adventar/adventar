package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
)

type verifier interface {
	VerifyIDToken(string) (*AuthResult, error)
}

type metaFetcher interface {
	Fetch(string) (*SiteMeta, error)
}

// Service holds data used by grpc functions.
type Service struct {
	db          *sql.DB
	verifier    verifier
	metaFetcher metaFetcher
}

// NewService creates a new Service.
func NewService(db *sql.DB, verifier verifier, metaFetcher metaFetcher) *Service {
	return &Service{db: db, verifier: verifier, metaFetcher: metaFetcher}
}

func (s *Service) serve(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{})
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpc_logrus.ReplaceGrpcLogger(logger)

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logger, opts...),
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
				resp, err := handler(ctx, req)
				s, _ := status.FromError(err)
				if s.Code() == codes.Unknown {
					fmt.Printf("%+v\n", err)
				}
				return resp, err
			},
		),
	)
	pb.RegisterAdventarServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// ListCalendars lists calendars.
func (s *Service) ListCalendars(ctx context.Context, in *pb.ListCalendarsRequest) (*pb.ListCalendarsResponse, error) {
	conditionQueries := []string{"c.year = ?"}
	limitQuery := ""
	conditionValues := []interface{}{in.GetYear()}
	if in.GetUserId() != 0 {
		conditionQueries = append(conditionQueries, "c.user_id = ?")
		conditionValues = append(conditionValues, in.GetUserId())
	}
	if in.GetQuery() != "" {
		conditionQueries = append(conditionQueries, "(c.title like ? or c.description like ?)")
		conditionValues = append(conditionValues, "%"+in.GetQuery()+"%", "%"+in.GetQuery()+"%")
	}
	if in.GetPageSize() != 0 {
		limitQuery = "limit ?"
		conditionValues = append(conditionValues, in.GetPageSize())
	}
	sql := fmt.Sprintf(`
		select
			c.id,
			c.title,
			c.description,
			c.year,
			u.id,
			u.name,
			u.icon_url
		from calendars as c
		inner join users as u on u.id = c.user_id
		where %s
		order by c.id desc
		%s
	`, strings.Join(conditionQueries, " and "), limitQuery)

	rows, err := s.db.Query(sql, conditionValues...)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendars: %w", err)
	}
	defer rows.Close()
	var calendars []*pb.Calendar
	for rows.Next() {
		var calendar pb.Calendar
		var user pb.User
		err := rows.Scan(
			&calendar.Id,
			&calendar.Title,
			&calendar.Description,
			&calendar.Year,
			&user.Id,
			&user.Name,
			&user.IconUrl,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed to scan row: %w", err)
		}
		calendar.Owner = &user
		calendars = append(calendars, &calendar)
	}

	if len(calendars) != 0 {
		err := s.bindEntryCount(calendars)
		if err != nil {
			return nil, xerrors.Errorf("Failed to bind entry count: %w", err)
		}
	}

	return &pb.ListCalendarsResponse{Calendars: calendars}, nil
}

// GetCalendar returns a calendar.
func (s *Service) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.GetCalendarResponse, error) {
	var calendar calendar
	row := s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetCalendarId())
	err := row.Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}

	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	entries, err := s.findEntries(calendar.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed to find entries: %w", err)
	}

	pbCalendar := &pb.Calendar{
		Id:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
		EntryCount:  int32(len(entries)),
	}
	return &pb.GetCalendarResponse{Calendar: pbCalendar, Entries: entries}, nil
}

// CreateCalendar creates a calendar.
func (s *Service) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}
	if currentUser == nil {
		return nil, status.Errorf(codes.PermissionDenied, "Unauthorized user")
	}

	stmt, err := s.db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(currentUser.ID, in.GetTitle(), in.GetDescription(), time.Now().Year())
	if err != nil {
		return nil, xerrors.Errorf("Failed query to insert into calendar: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf("Failed to get last id: %w", err)
	}

	var calendar calendar
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", lastID).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

// UpdateCalendar updates the calendar.
func (s *Service) UpdateCalendar(ctx context.Context, in *pb.UpdateCalendarRequest) (*pb.Calendar, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}
	stmt, err := s.db.Prepare("update calendars set title = ?, description = ? where id = ? and user_id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(in.GetTitle(), in.GetDescription(), in.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update calendar: %w", err)
	}

	var calendar calendar
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetCalendarId()).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "Calendar not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

// DeleteCalendar deletes the calendar.
func (s *Service) DeleteCalendar(ctx context.Context, in *pb.DeleteCalendarRequest) (*empty.Empty, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	stmt, err := s.db.Prepare("delete from calendars where id = ? and user_id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(in.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to delete calendar: %w", err)
	}
	return &empty.Empty{}, nil
}

// ListEntries lists entries.
func (s *Service) ListEntries(ctx context.Context, in *pb.ListEntriesRequest) (*pb.ListEntriesResponse, error) {
	conditionQueries := []string{"e.user_id = ?"}
	conditionValues := []interface{}{in.GetUserId()}

	if in.GetYear() != 0 {
		conditionQueries = append(conditionQueries, "c.year = ?")
		conditionValues = append(conditionValues, in.GetYear())
	}

	sql := fmt.Sprintf(`
		select
			e.id,
			e.day,
			e.title,
			e.comment,
			e.url,
			e.image_url,
			c.id,
			c.title,
			c.description,
			u.id,
			u.name,
			u.icon_url
		from entries as e
		inner join users as u on u.id = e.user_id
		inner join calendars as c on c.id = e.calendar_id
		where %s
		order by e.day
	`, strings.Join(conditionQueries, " and "))

	rows, err := s.db.Query(sql, conditionValues...)
	defer rows.Close()
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entries: %w", err)
	}

	entries := []*pb.Entry{}
	for rows.Next() {
		var e pb.Entry
		var c pb.Calendar
		var u pb.User
		err := rows.Scan(
			&e.Id,
			&e.Day,
			&e.Title,
			&e.Comment,
			&e.Url,
			&e.ImageUrl,
			&c.Id,
			&c.Title,
			&c.Description,
			&u.Id,
			&u.Name,
			&u.IconUrl,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed to scan row: %w", err)
		}
		e.Calendar = &c
		e.Owner = &u
		entries = append(entries, &e)
	}

	return &pb.ListEntriesResponse{Entries: entries}, nil
}

// CreateEntry creates a entry.
func (s *Service) CreateEntry(ctx context.Context, in *pb.CreateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	var year int
	err = s.db.QueryRow("select year from calendars where id = ?", in.GetCalendarId()).Scan(&year)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "Calendar not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	day := in.GetDay()
	if day < 1 || day > 25 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid day: %d", day)
	}

	// TODO: Specify default value by schema definition.
	stmt, err := s.db.Prepare("insert into entries(user_id, calendar_id, day, comment, url, title, image_url) values(?, ?, ?, '', '', '', '')")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(currentUser.ID, in.GetCalendarId(), day)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to insert into entry: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf("Failed to get last id: %w", err)
	}

	var entryID int64
	err = s.db.QueryRow("select id from entries where id = ?", lastID).Scan(&entryID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return &pb.Entry{Id: entryID}, nil
}

// UpdateEntry updates the entry.
func (s *Service) UpdateEntry(ctx context.Context, in *pb.UpdateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	stmt, err := s.db.Prepare("update entries set comment = ?, url = ? where id = ? and user_id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(in.GetComment(), in.GetUrl(), in.GetEntryId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update entry: %w", err)
	}

	if in.GetUrl() != "" {
		m, err := s.metaFetcher.Fetch(in.GetUrl())
		// TODO: Ignore error
		if err != nil {
			return nil, xerrors.Errorf("Failed to fetch url: %w", err)
		}
		stmt, err = s.db.Prepare("update entries set title = ?, image_url = ? where id = ? and user_id = ?")
		if err != nil {
			return nil, xerrors.Errorf("Failed to prepare query: %w", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(m.Title, m.ImageURL, in.GetEntryId(), currentUser.ID)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update entry: %w", err)
		}
	}

	var comment string
	var url string
	var title string
	var imageURL string
	err = s.db.QueryRow("select comment, url, title, image_url from entries where id = ?", in.GetEntryId()).Scan(&comment, &url, &title, &imageURL)
	if err == sql.ErrNoRows {
		return nil, status.Errorf(codes.NotFound, "Entry not found")
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entry: %w", err)
	}

	return &pb.Entry{Id: in.GetEntryId(), Comment: comment, Url: url, Title: title, ImageUrl: imageURL}, nil
}

// DeleteEntry deletes the entry.
func (s *Service) DeleteEntry(ctx context.Context, in *pb.DeleteEntryRequest) (*empty.Empty, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	// TODO: Calendar owner can cancel entry
	stmt, err := s.db.Prepare("delete from entries where id = ? and user_id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare entry: %w", err)
	}

	_, err = stmt.Exec(in.GetEntryId(), currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to delete entry: %w", err)
	}

	return &empty.Empty{}, nil
}

// SignIn validates the id token.
func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.User, error) {
	authResult, err := s.verifier.VerifyIDToken(in.GetJwt())
	if err != nil {
		return nil, xerrors.Errorf("Failed to verify token: %w", err)
	}

	var userID int64
	var userName string
	err = s.db.QueryRow("select id, name from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&userID, &userName)

	if err != nil && err != sql.ErrNoRows {
		return nil, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	u := &pb.User{Name: authResult.Name, IconUrl: authResult.IconURL}

	if err == sql.ErrNoRows {
		stmt, err := s.db.Prepare("insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)")
		if err != nil {
			return nil, xerrors.Errorf("Failed to prepare query: %w", err)
		}
		defer stmt.Close()
		res, err := stmt.Exec(authResult.Name, authResult.AuthUID, authResult.AuthProvider, authResult.IconURL)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to insert into user: %w", err)
		}
		u.Id, err = res.LastInsertId()
		if err != nil {
			return nil, xerrors.Errorf("Failed to get last id: %w", err)
		}
	} else {
		stmt, err := s.db.Prepare("update users set icon_url = ? where id = ?")
		if err != nil {
			return nil, xerrors.Errorf("Failed to prepare query: %w", err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(authResult.IconURL, userID)
		if err != nil {
			return nil, xerrors.Errorf("Failed query to update user: %w", err)
		}
		u.Id = userID
		u.Name = userName
	}

	return u, nil
}

// UpdateUser updates user info.
func (s *Service) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}
	name := in.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is blank")
	}

	stmt, err := s.db.Prepare("update users set name = ? where id = ?")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}

	_, err = stmt.Exec(name, currentUser.ID)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to update user: %w", err)
	}

	return &pb.User{Id: currentUser.ID, Name: name, IconUrl: currentUser.IconURL}, nil
}

func (s *Service) getCurrentUser(ctx context.Context) (*user, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, xerrors.Errorf("Metadata not found")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, xerrors.Errorf("Authorization metadata not found")
	}

	authResult, err := s.verifier.VerifyIDToken(values[0])
	if err != nil {
		return nil, xerrors.Errorf("Failed to verify token: %w", err)
	}

	var user user
	err = s.db.QueryRow("select id, name, icon_url from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&user.ID, &user.Name, &user.IconURL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch user: %w", err)
	}

	return &user, nil
}

func (s *Service) bindEntryCount(calendars []*pb.Calendar) error {
	ids := []interface{}{}
	interpolations := []string{}

	for _, c := range calendars {
		ids = append(ids, c.Id)
		interpolations = append(interpolations, "?")
	}

	sql := fmt.Sprintf("select calendar_id, count(*) from entries where calendar_id in (%s) group by calendar_id", strings.Join(interpolations, ","))
	rows, err := s.db.Query(sql, ids...)
	if err != nil {
		return xerrors.Errorf("Failed query to fetch calendar: %w", err)
	}

	entryCounts := map[int64]int32{}
	for rows.Next() {
		var cid int64
		var count int32
		if err := rows.Scan(&cid, &count); err != nil {
			return xerrors.Errorf("Failed to scan row: %w", err)
		}
		entryCounts[cid] = count
	}

	for _, c := range calendars {
		c.EntryCount = entryCounts[c.Id]
	}

	return nil
}

func (s *Service) findEntries(cid int64) ([]*pb.Entry, error) {
	rows, err := s.db.Query(`
		select
			e.id,
			e.day,
			e.title,
			e.comment,
			e.url,
			e.image_url,
			u.id,
			u.name,
			u.icon_url
		from entries as e
		inner join users as u on u.id = e.user_id
		where e.calendar_id = ?
		order by e.day
	`, cid)

	if err != nil {
		return nil, xerrors.Errorf("Failed query to fetch entries: %w", err)
	}

	entries := []*pb.Entry{}
	for rows.Next() {
		var e pb.Entry
		var u pb.User
		err := rows.Scan(
			&e.Id,
			&e.Day,
			&e.Title,
			&e.Comment,
			&e.Url,
			&e.ImageUrl,
			&u.Id,
			&u.Name,
			&u.IconUrl,
		)
		if err != nil {
			return nil, xerrors.Errorf("Failed to scan row: %w", err)
		}
		e.Owner = &u
		entries = append(entries, &e)
	}

	return entries, nil
}
