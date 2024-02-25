package internalgrpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lixoi/survey/internal/server/grpc/api"
)

type Validator func(req interface{}) error

func UnaryServerRequestValidatorInterceptor(validator Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validator(req); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "%s is rejected by validate interceptor. Error: %v", info.FullMethod, err)
		}
		return handler(ctx, req)
	}
}

func ValidateReq(req interface{}) error {
	switch r := req.(type) {
	case *api.UserInfoRequest:
		if r.UserId == 0 {
			// generate user_id
			return nil
		}
	case *api.UserIdRequest:
		if r.UserId == 0 {
			return errors.New("interceptor validator: not current user id")
		}
	case *api.AnswerRequest:
		if r.UserId == 0 || r.Number == 0 {
			return errors.New("interceptor validator: not current user id or index of question")
		}
	}
	return nil
}
