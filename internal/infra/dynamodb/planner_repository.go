package dynamodb

import (
	"context"
	"money-manager/internal/domain/planner"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PlannerRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewPlannerRepository(client *dynamodb.Client, tableName string) *PlannerRepository {
	return &PlannerRepository{
		client:    client,
		tableName: tableName,
	}
}

func (p *PlannerRepository) ReadEvent(userId string, ctx context.Context) ([]domain.PlannedEvent, error) {
	events, err := p.client.Query(ctx, &dynamodb.QueryInput{
		TableName:              aws.String(p.tableName),
		KeyConditionExpression: aws.String("PK = :pk AND begins_with(SK, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "USER#" + userId},
			":sk": &types.AttributeValueMemberS{Value: "EVENT#"},
		},
	})

	if err != nil {
		return nil, err
	}

	var plannedEvents []domain.PlannedEvent

	for _, item := range events.Items {
		var event EventItem
		err = attributevalue.UnmarshalMap(item, &event)
		if err != nil {
			return nil, err
		}

		startDateParsed, err := time.Parse(time.RFC3339, event.StartDate)
		if err != nil {
			return nil, err
		}

		plannedEvent := domain.PlannedEvent{
			Id:         event.Id,
			UserId:     event.UserId,
			AccountId:  event.AccountId,
			CategoryId: event.CategoryId,
			Name:	  event.Name,
			StartDate:  startDateParsed,
			Description: event.Description,
			Status:      domain.PlannedEventStatus(event.PlannedEventStatus),
		}
		plannedEvents = append(plannedEvents, plannedEvent)
	}

	return plannedEvents, nil
}

func (p *PlannerRepository) DeleteEvent(userId string, eventID string, ctx context.Context) error {
	_, err := p.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(p.tableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "USER#" + userId},
			"SK": &types.AttributeValueMemberS{Value: "EVENT#" + eventID},
		},
	})
	return err
}

func (p *PlannerRepository) SaveEvent(event domain.PlannedEvent, ctx context.Context) error {
	newEvent := EventItem{
		PK:                 "USER#" + event.UserId,
		SK:                 "EVENT#" + event.Id,
		Id:                 event.Id,
		UserId:             event.UserId,
		AccountId:          event.AccountId,
		CategoryId:         event.CategoryId,
		Name:               event.Name,
		StartDate:          event.StartDate.Format(time.RFC3339),
		Description:        event.Description,
		PlannedEventStatus: string(event.Status),
		Amount: MoneyItem{
			Amount:   event.Amount.Amount,
			Currency: string(event.Amount.Currency),
		},
	}

	av, err := attributevalue.MarshalMap(newEvent)

	_, err = p.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(p.tableName),
		Item:      av,
	})

	return err
}

// DeleteCategory implements [repository.PlannerRepository].
func (p *PlannerRepository) DeleteCategory(categoryID string, ctx context.Context) error {
	panic("unimplemented")
}

// SaveCategory implements [repository.PlannerRepository].
func (p *PlannerRepository) SaveCategory(category domain.Category, ctx context.Context) error {
	panic("unimplemented")
}



// DeleteRecurringEvent implements [repository.PlannerRepository].
func (p *PlannerRepository) DeleteRecurringEvent(eventID string, ctx context.Context) error {
	panic("unimplemented")
}

// SavePlannedTransfer implements [repository.PlannerRepository].
func (p *PlannerRepository) SavePlannedTransfer(transfer domain.PlannedTransfer, ctx context.Context) error {
	panic("unimplemented")
}

// DeletePlannedTransfer implements [repository.PlannerRepository].
func (p *PlannerRepository) DeletePlannedTransfer(transferID string, ctx context.Context) error {
	panic("unimplemented")
}