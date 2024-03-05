package internalgrpc

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"sort"
	"time"

	"github.com/lixoi/survey/internal/app"
	"github.com/lixoi/survey/internal/config"
	log "github.com/lixoi/survey/internal/logger"
	"github.com/lixoi/survey/internal/server/grpc/api"
	"github.com/lixoi/survey/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	api.UnimplementedICHSurveyServer
	strg app.Storage
	logg log.Logger
}

func New(strg app.Storage, logg log.Logger) *GRPCServer {
	return &GRPCServer{
		strg: strg,
		logg: logg,
	}
}

func (s *GRPCServer) AddCandidate(ctx context.Context, req *api.UserInfoRequest) (*api.StatusResponse, error) {
	existTo := time.Time{}
	if req.ExistTo > 0 {
		existTo = time.Now().AddDate(0, 0, int(req.ExistTo))
	}
	user := storage.User{
		ID:            req.UserId,
		BaseQ:         req.BaseQuestion.String(),
		FirstProfileQ: req.FirstGuestion.String(),
		SecProfileQ:   req.SecondGuestion.String(),
		ExistTo:       existTo,
	}
	message := ""
	err := s.strg.AddUser(ctx, user)
	if err != nil {
		message = err.Error()
	}

	return &api.StatusResponse{
		Message: message,
	}, err
}

func (s *GRPCServer) DeleteCandidate(ctx context.Context, req *api.UserIdRequest) (*api.StatusResponse, error) {
	message := ""
	err := s.strg.DeleteUser(ctx, req.UserId)
	if err != nil {
		message = err.Error()
	}

	return &api.StatusResponse{
		Message: message,
	}, err
}

func (s *GRPCServer) StartSurvey(ctx context.Context, req *api.UserIdRequest) (*api.QuestionResponse, error) {
	res := &api.QuestionResponse{}
	question, err := s.strg.StartSurveyFor(ctx, req.UserId)
	if err != nil {
		res.Message = err.Error()
		res.Number = 0
		res.Question = ""
	} else {
		res.Message = ""
		res.Question = question.Question
		res.Number = question.QuestionNumber
		res.UserId = question.UserID
	}

	return res, err
}

func (s *GRPCServer) SetAnswer(ctx context.Context, req *api.AnswerRequest) (*api.QuestionResponse, error) {
	res := &api.QuestionResponse{}
	nextQuestion, err := s.strg.SetAnswerFor(ctx, req.UserId, req.Number, req.Answer)
	if err != nil {
		res.Message = err.Error()
		res.Number = 0
		res.Question = ""
	} else {
		res.Message = ""
		res.Question = nextQuestion.Question
		res.Number = nextQuestion.QuestionNumber
		res.UserId = nextQuestion.UserID
	}

	return res, err
}

func (s *GRPCServer) SetFinishCandidate(ctx context.Context, req *api.UserIdRequest) (*api.StatusResponse, error) {
	message := ""
	err := s.strg.FinishSurveyFor(ctx, req.UserId)
	if err != nil {
		message = err.Error()
	}

	return &api.StatusResponse{
		Message: message,
	}, err
}

func (s *GRPCServer) GetSurveyForCandidate(ctx context.Context, req *api.UserIdRequest) (*api.SurveyResponse, error) {
	res := &api.SurveyResponse{}
	qList, err := s.strg.GetSurveyFor(ctx, req.UserId)
	if err != nil {
		res.Mesage = err.Error()
		return res, err
	}
	if len(qList) == 0 {
		res.Mesage = "Not questions for user"
		return res, errors.New(res.Mesage)
	}
	res.Qs = s.marshalingList(ctx, qList)
	res.Mesage = ""

	return res, nil
}

func (s *GRPCServer) Start(ctx context.Context, config config.Config) error {
	l, err := net.Listen("tcp", ":"+config.Server.GrpcPort)
	if err != nil {
		s.logg.Error(err.Error())
		return err
	}

	defer s.Stop(ctx)

	tlsCredentials, err := s.loadTLSCredentials(config)
	if err != nil {
		s.logg.Error("cannot load TLS credentials: " + err.Error())
		return err
	}

	srv := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.ChainUnaryInterceptor(
			UnaryServerRequestValidatorInterceptor(ValidateReq),
		),
	)
	api.RegisterICHSurveyServer(srv, s)
	reflection.Register(srv)

	return srv.Serve(l)
}

func (s *GRPCServer) Stop(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

func (s *GRPCServer) marshalingList(ctx context.Context, qList []storage.Survey) (res []*api.Survey) {
	usr, err := s.strg.GetInfoFor(ctx, qList[0].UserID)
	if err != nil {
		return
	}
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].QuestionNumber < qList[j].QuestionNumber
	})
	beforeTime := &usr.SurveyStart
	for _, v := range qList {
		latency := "not time"
		if v.AnsweredAt != nil && beforeTime != nil {
			latency = (*v.AnsweredAt).Sub(*beforeTime).String()
		}
		s := &api.Survey{
			UserId:   v.UserID,
			Title:    v.Title,
			Question: v.Question,
			Answer:   v.Answer,
			Number:   v.QuestionNumber,
			Latency:  latency,
		}
		res = append(res, s)
		beforeTime = v.AnsweredAt
	}

	return res
}

func (s *GRPCServer) loadTLSCredentials(conf config.Config) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(conf.Certs.SrvCert, conf.Certs.SrvKey)
	if err != nil {
		return nil, err
	}
	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
