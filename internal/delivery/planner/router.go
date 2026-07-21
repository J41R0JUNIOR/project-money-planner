package delivery

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Router struct {
	handler *PlannerHandler
}

func NewRouter(
	handler *PlannerHandler,
) *Router {

	return &Router{
		handler: handler,
	}
}

func (r *Router) Handle(
	ctx context.Context,
	event events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {

	switch {
	case event.RequestContext.HTTP.Method == http.MethodPost &&
		event.RawPath == "/planner/event":

		return r.handler.CreateEvent(ctx, event)

	case event.RequestContext.HTTP.Method == http.MethodDelete &&
		event.RawPath == "/planner/event":

		return r.handler.DeleteEvent(ctx, event)

	case event.RequestContext.HTTP.Method == http.MethodPut &&
		event.RawPath == "/planner/event":

		return r.handler.UpdateEvent(ctx, event)
	default:

		return events.APIGatewayV2HTTPResponse{
			StatusCode: http.StatusNotFound,
			Body:       `{"message":"route not found"}`,
		}, nil
	}
}
