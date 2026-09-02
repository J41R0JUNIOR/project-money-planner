package account

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	Handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{
		Handler: handler,
	}
}

func (r *Router) Handle(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	switch {
	case event.RequestContext.HTTP.Method == http.MethodPost &&
		event.RequestContext.HTTP.Path == "/account":
		return r.Handler.CreateAccount(ctx, event)

	case event.RequestContext.HTTP.Method == http.MethodGet &&
		event.RequestContext.HTTP.Path == "/account":
		return r.Handler.GetAccounts(ctx, event)
	default:

		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body:       `{"message":"route not found"}`,
		}, nil
	}
}
