package main

import (
	"context"
	"log"
	"net/http"

	adventarv1 "github.com/adventar/adventar/backend/pkg/gen/adventar/v1"
	"github.com/adventar/adventar/backend/pkg/gen/adventar/v1/adventarv1connect"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type AdventarServer struct {
	adventarv1connect.UnimplementedAdventarHandler // TODO: remove
}

func (s *AdventarServer) GetCalendar(
	ctx context.Context,
	req *connect.Request[adventarv1.GetCalendarRequest],
) (*connect.Response[adventarv1.GetCalendarResponse], error) {
	id := req.Msg.CalendarId
	res := connect.NewResponse(&adventarv1.GetCalendarResponse{
		Calendar: &adventarv1.Calendar{
			Id: id,
		},
	})
	return res, nil
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(adventarv1connect.NewAdventarHandler(&AdventarServer{}))
	err := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
	log.Fatalf("listen failed: %v", err)
}
