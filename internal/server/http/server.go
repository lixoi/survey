package internalhttp

import (
	"context"
	"net/http"

	"github.com/lixoi/survey/internal/app"
)

type Server struct { // TODO
	srv *http.Server
	//storage app.Storage
}

/*
type Logger interface { // TODO
}


type Application interface { // TODO
}
*/

func NewServer(logger app.Logger, a app.App) *Server {
	mux := http.NewServeMux()

	h := NewService(a)

	mux.HandleFunc("/hello", loggingMiddleware(h.GetHello))

	//mux.HandleFunc("/user/add/<id>", loggingMiddleware(h.AddUserId))
	//mux.HandleFunc("/user/del/<id>", loggingMiddleware(h.DelUserId))
	//mux.HandleFunc("/question/<id>/<index>", loggingMiddleware(h.GetStats))
	// websocket handler
	//mux.HandleFunc("/result/<id>", loggingMiddleware(h.StatStream))

	return &Server{srv: &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}}
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	s.srv.ListenAndServe()
	s.Stop(ctx)
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	<-ctx.Done()
	return nil
}

// TODO
