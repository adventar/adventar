package service

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/adventar/adventar/backend/grpc-server/grpc/adventar/v1"
	"github.com/adventar/adventar/backend/grpc-server/model"
	"github.com/adventar/adventar/backend/grpc-server/util"
	"github.com/golang/protobuf/ptypes/empty"
)

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
	var calendar model.Calendar
	var user model.User
	selectSQL := `
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
		where c.id = ?
	`

	row := s.db.QueryRow(selectSQL, in.GetCalendarId())
	err := row.Scan(
		&calendar.ID,
		&calendar.Title,
		&calendar.Description,
		&calendar.Year,
		&user.ID,
		&user.Name,
		&user.IconURL,
	)

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

	pbUser := &pb.User{Id: user.ID, Name: user.Name, IconUrl: user.IconURL}
	pbCalendar := &pb.Calendar{
		Id:          calendar.ID,
		Title:       calendar.Title,
		Description: calendar.Description,
		Year:        calendar.Year,
		Owner:       pbUser,
		EntryCount:  int32(len(entries)),
	}

	return &pb.GetCalendarResponse{Calendar: pbCalendar, Entries: entries}, nil
}

// CreateCalendar creates a calendar.
func (s *Service) CreateCalendar(ctx context.Context, in *pb.CreateCalendarRequest) (*pb.Calendar, error) {
	now, err := util.CurrentDate()
	if err != nil {
		return nil, err
	}

	if now.Month < 11 {
		return nil, status.Errorf(codes.FailedPrecondition, "Calendars can not create now.")
	}

	currentUser, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
	}

	stmt, err := s.db.Prepare("insert into calendars(user_id, title, description, year) values(?, ?, ?, ?)")
	if err != nil {
		return nil, xerrors.Errorf("Failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(currentUser.ID, in.GetTitle(), in.GetDescription(), now.Year)
	if err != nil {
		return nil, xerrors.Errorf("Failed query to insert into calendar: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, xerrors.Errorf("Failed to get last id: %w", err)
	}

	var calendar model.Calendar
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
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
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

	var calendar model.Calendar
	err = s.db.QueryRow("select id, user_id, title, description, year from calendars where id = ? and user_id = ?", in.GetCalendarId(), currentUser.ID).Scan(&calendar.ID, &calendar.UserID, &calendar.Title, &calendar.Description, &calendar.Year)
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
		return nil, status.Errorf(codes.PermissionDenied, "Authentication failed")
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
