package internalhttp

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lixoi/survey/internal/app"
	"github.com/lixoi/survey/internal/server/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct { // TODO
	//srv  *http.Server
	rmux *runtime.ServeMux
	addr string
	logg app.Logger
	//storage app.Storage
}

/*
type Logger interface { // TODO
}


type Application interface { // TODO
}
*/

func NewServer(logger app.Logger, a app.App) *Server {
	rmux := runtime.NewServeMux()
	/*
	   mux := http.NewServeMux()
	   mux.Handle("/", rmux)
	   mux.Handle("/set_answer", rmux)

	   h := NewHandler(a)
	   mux.HandleFunc("/hello", loggingMiddleware(h.GetHello, logger))
	   mux.HandleFunc("/user/add/<id>", loggingMiddleware(h.AddUserId))
	   mux.HandleFunc("/user/del/<id>", loggingMiddleware(h.DelUserId))
	   mux.HandleFunc("/question/<id>/<index>", loggingMiddleware(h.GetStats))
	   // websocket handler
	   //mux.HandleFunc("/result/<id>", loggingMiddleware(h.StatStream))

	   	return &Server{srv: &http.Server{
	   		Addr:    ":8080",
	   		Handler: rmux,
	   	},

	   		logg: logger,
	   	}
	*/
	return &Server{
		rmux: rmux,
		addr: ":8080",
		logg: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	// TODO
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterICHSurveyHandlerFromEndpoint(ctx, s.rmux, ":50051", opts)
	if err := http.ListenAndServe(s.addr, s.rmux); err != nil {
		//if err := s.srv.ListenAndServe(); err != nil {
		s.logg.Error("Not start server: " + err.Error())
		return err
	}
	s.Stop(ctx)
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	// TODO
	<-ctx.Done()
	return nil
}

// TODO
