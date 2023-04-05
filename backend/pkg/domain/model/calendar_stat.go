package model

import adventarv1 "github.com/adventar/adventar/backend/pkg/gen/proto/adventar/v1"

type CalendarStat struct {
	Year           int32
	CalendarsCount int32
	EntriesCount   int32
}

func (x *CalendarStat) ToProto() *adventarv1.CalendarStat {
	return &adventarv1.CalendarStat{
		Year:           x.Year,
		CalendarsCount: x.CalendarsCount,
		EntriesCount:   x.EntriesCount,
	}
}
