package internalhttp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lixoi/survey/internal/app"
	"github.com/lixoi/survey/internal/config"
	"github.com/lixoi/survey/internal/server/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	rmux *runtime.ServeMux
	logg app.Logger
}

/*
type Logger interface {
}
type Application interface {
}
*/

func loadTLSCredentials(conf config.Config) (tls.Certificate, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(conf.Certs.SrvCert, conf.Certs.SrvKey)
	if err != nil {
		return serverCert, err
	}

	return serverCert, nil
}

func New(logger app.Logger) *Server {
	rmux := runtime.NewServeMux()
	return &Server{
		rmux: rmux,
		logg: logger,
	}
}

func (s *Server) Start(ctx context.Context, conf config.Config) error {
	cert, err := loadTLSCredentials(conf)
	if err != nil {
		s.logg.Error("not load certs: " + err.Error())
		return fmt.Errorf("not load certs, check config file")
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		s.logg.Error("not parse certificates: " + err.Error())
		return fmt.Errorf("not parse certificates")
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(cert.Leaf)
	cp := credentials.NewClientTLSFromCert(certPool, "")

	opts := []grpc.DialOption{grpc.WithTransportCredentials(cp)}
	err = api.RegisterICHSurveyHandlerFromEndpoint(ctx, s.rmux, ":"+conf.Server.GrpcPort, opts)

	mux := http.NewServeMux()
	mux.Handle("/", s.rmux) // Handle gRPC-gateway requests

	if conf.Server.Swagger {
		fs := http.FileServer(http.Dir("./../../swagger"))
		http.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
		mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs)) // Handle Swagger UI requests
	}

	gwServer := http.Server{
		Addr: conf.Server.HostName + ":" + conf.Server.HttpPort,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
		Handler: mux,
	}

	if err := gwServer.ListenAndServeTLS("", ""); err != nil {
		s.logg.Error("Not start server: " + err.Error())
		return err
	}

	s.Stop(ctx)
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
