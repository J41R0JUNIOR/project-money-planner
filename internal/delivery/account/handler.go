package account

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateAccount(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	panic("unimplemented")
}

func (h *Handler) GetAccounts(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	panic("unimplemented")
}