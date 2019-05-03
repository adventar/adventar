package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/adventar/adventar/grpc-server/adventar/v1"
)

type verifier interface {
	VerifyIDToken(string) *AuthResult
}

// Service holds data used by grpc functions.
type Service struct {
	db       *sql.DB
	verifier verifier
}

// NewService creates a new Service.
func NewService(db *sql.DB, verifier verifier) *Service {
	return &Service{db: db, verifier: verifier}
}

func (s *Service) serve(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterAdventarServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
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
		return nil, err
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
			return nil, err
		}
		calendar.Owner = &user
		calendars = append(calendars, &calendar)
	}

	if len(calendars) != 0 {
		err := s.bindEntryCount(calendars)
		if err != nil {
			return nil, err
		}
	}

	return &pb.ListCalendarsResponse{Calendars: calendars}, nil
}

// GetCalendar returns a calendar.
func (s *Service) GetCalendar(ctx context.Context, in *pb.GetCalendarRequest) (*pb.GetCalendarResponse, error) {
	var calendar calendar
	row := s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetCalendarId())
	err := row.Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}
	pbCalendar := &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}
	return &pb.GetCalendarResponse{Calendar: pbCalendar}, nil
}

// CreateCalendar creates a calendar.
func (s *Service) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(currentUser.ID, in.GetTitle(), in.GetDescription(), time.Now().Year())
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var calendar calendar
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", lastID).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

// UpdateCalendar updates the calendar.
func (s *Service) UpdateCalendar(ctx context.Context, in *pb.UpdateCalendarRequest) (*pb.Calendar, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	stmt, err := s.db.Prepare("update calendars set title = ?, description = ? where id = ? and user_id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(in.GetTitle(), in.GetDescription(), in.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, err
	}

	var calendar calendar
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ?", in.GetCalendarId()).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
	if err != nil {
		return nil, err
	}

	return &pb.Calendar{Id: calendar.ID, Title: calendar.Title, Description: calendar.Description, Year: calendar.Year}, nil
}

// DeleteCalendar deletes the calendar.
func (s *Service) DeleteCalendar(ctx context.Context, in *pb.DeleteCalendarRequest) (*empty.Empty, error) {
	currentUser, err := s.getCurrentUser(ctx)
	stmt, err := s.db.Prepare("delete from calendars where id = ? and user_id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(in.GetCalendarId(), currentUser.ID)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// CreateEntry creates a entry.
func (s *Service) CreateEntry(ctx context.Context, in *pb.CreateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	var year int
	row := s.db.QueryRow("select year from calendars where id = ?", in.GetCalendarId())
	err = row.Scan(&year)
	if err != nil {
		return nil, err
	}

	day := in.GetDay()
	if day < 1 || day > 25 {
		return nil, fmt.Errorf("Invalid day: %d", day)
	}

	stmt, err := s.db.Prepare("insert into entries(user_id, calendar_id, date) values(?, ?, ?)")
	if err != nil {
		return nil, err
	}
	dateStr := fmt.Sprintf("%d-12-%d", year, in.GetDay())
	res, err := stmt.Exec(currentUser.ID, in.GetCalendarId(), dateStr)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var entryID int64
	err = s.db.QueryRow("select id from calendars where id = ?", lastID).Scan(&entryID)
	if err != nil {
		return nil, err
	}

	return &pb.Entry{Id: entryID}, nil
}

// UpdateEntry updates the entry.
func (s *Service) UpdateEntry(ctx context.Context, in *pb.UpdateEntryRequest) (*pb.Entry, error) {
	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	stmt, err := s.db.Prepare("update entries set comment = ?, url = ? where id = ? and user_id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(in.GetComment(), in.GetUrl(), in.GetEntryId(), currentUser.ID)
	if err != nil {
		return nil, err
	}

	return &pb.Entry{Id: in.GetEntryId()}, nil
}

// DeleteEntry deletes the entry.
func (s *Service) DeleteEntry(ctx context.Context, in *pb.DeleteEntryRequest) (*empty.Empty, error) {
	stmt, err := s.db.Prepare("delete from entries where id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(in.GetEntryId())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

// SignIn validates the id token.
func (s *Service) SignIn(ctx context.Context, in *pb.SignInRequest) (*empty.Empty, error) {
	authResult := s.verifier.VerifyIDToken(in.GetJwt())
	var userID int
	err := s.db.QueryRow("select id from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&userID)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		stmt, err := s.db.Prepare("insert into users (name, auth_uid, auth_provider, icon_url) values (?, ?, ?, ?)")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(authResult.Name, authResult.AuthUID, authResult.AuthProvider, authResult.IconURL)
		if err != nil {
			return nil, err
		}
	} else {
		stmt, err := s.db.Prepare("update users set icon_url = ? where id = ?")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(authResult.IconURL, userID)
		if err != nil {
			return nil, err
		}
	}

	return &empty.Empty{}, nil
}

func (s *Service) getCurrentUser(ctx context.Context) (*user, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("not found metadata")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, fmt.Errorf("not found authorization in metadata")
	}

	authResult := s.verifier.VerifyIDToken(values[0])

	var user user
	err := s.db.QueryRow("select id, name, icon_url from users where auth_provider = ? and auth_uid = ?", authResult.AuthProvider, authResult.AuthUID).Scan(&user.ID, &user.Name, &user.IconURL)
	if err != nil {
		return nil, err
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
		return err
	}

	entryCounts := map[int64]int32{}
	for rows.Next() {
		var cid int64
		var count int32
		if err := rows.Scan(&cid, &count); err != nil {
			return err
		}
		entryCounts[cid] = count
	}

	for _, c := range calendars {
		c.EntryCount = entryCounts[c.Id]
	}

	return nil
}
