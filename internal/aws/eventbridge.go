package aws

import "github.com/aws/aws-sdk-go-v2/service/eventbridge"

type EventBridgeService struct {
	client *eventbridge.Client
}

func (f *ServiceFactory) NewEventBridgeService() *EventBridgeService {
	return &EventBridgeService{client: f.NewEventBridgeClient()}
}
