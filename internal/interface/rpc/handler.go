package rpc

import (
	"context"
	"errors"

	"github.com/peppydays/go-grpc-member/internal/service"
	"github.com/peppydays/go-grpc-member/pkg/contract"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BaseHandler struct {
	MemberService *service.MemberService
}

func NewBaseHandler(memberService *service.MemberService) *BaseHandler {
	return &BaseHandler{
		MemberService: memberService,
	}
}

func (h *BaseHandler) Error(_ context.Context, err error) error {
	switch err {
	case nil:
		return nil
	default:
		return status.Error(codes.Internal, errors.New("internal error").Error())
	}
}

type HealthHandler struct {
	contract.UnimplementedHealthServer
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(ctx context.Context, req *contract.HealthCheckRequest) (*contract.HealthCheckResponse, error) {
	return &contract.HealthCheckResponse{
		Status: contract.HealthCheckResponse_SERVING,
	}, nil
}
