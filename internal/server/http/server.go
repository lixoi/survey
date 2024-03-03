package internalhttp

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/loads"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	flags "github.com/jessevdk/go-flags"
	"github.com/lixoi/survey/internal/app"
	"github.com/lixoi/survey/internal/server/grpc/api"
	"github.com/lixoi/survey/restapi"
	"github.com/lixoi/survey/restapi/operations"
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

	fs := http.FileServer(http.Dir("./../../swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
	mux := http.NewServeMux()
	mux.Handle("/", s.rmux)                                        // Handle gRPC-gateway requests
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs)) // Handle Swagger UI requests
	if err := http.ListenAndServe(s.addr, mux); err != nil {
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

func (s *Server) StartSwagger(ctx context.Context) error {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	sapi := operations.NewSwaggerAPI(swaggerSpec)
	server := restapi.NewServer(sapi)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "api/api.proto"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range sapi.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = api.RegisterICHSurveyHandlerFromEndpoint(ctx, s.rmux, ":50051", opts)
	if err != nil {
		return nil
	}
	server.SetHandler(s.rmux)

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

// TODO
