package auth

import (
	"context"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
)


type Router struct {
	authHandler *Handler
}


func NewRouter(
	authHandler *Handler,
) *Router {

	return &Router{
		authHandler: authHandler,
	}
}


func (r *Router) Handle(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	switch {

	case event.RequestContext.HTTP.Method == http.MethodPost &&
		event.RawPath == "/signup":

		return r.authHandler.SignUp(ctx, event)

	case event.RequestContext.HTTP.Method == http.MethodPost &&
		event.RawPath == "/signin":

		return r.authHandler.SignIn(ctx, event)

	default:

		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body: `{"message":"route not found"}`,
		}, nil
	}
}