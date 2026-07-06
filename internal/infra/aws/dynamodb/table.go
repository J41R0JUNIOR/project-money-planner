package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"errors"
	"log"
)

func CreateTable(client *dynamodb.Client, tableName string) error {

	if client == nil {
		return fmt.Errorf("dynamodb client is nil")
	}

	// Check if the table already exists
	exists, err := TableBasics{DynamoDbClient: client, TableName: tableName}.tableExists(context.Background())
	if err != nil {
		return fmt.Errorf("failed to check if table exists: %v", err)
	}
	if exists {
		fmt.Printf("Table %v already exists.\n", tableName)
		return nil
	}

	// Define the table schema and create the table
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("PK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SK"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("PK"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("SK"),
				KeyType:       types.KeyTypeRange,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
		TableName:   aws.String(tableName),
	}

	_, err = client.CreateTable(context.Background(), input)
	if err != nil {
		return err
	}

	fmt.Println("Created the table", tableName)
	return nil
}


// TableBasics encapsulates the Amazon DynamoDB service actions used in the examples.
// It contains a DynamoDB service client that is used to act on the specified table.
type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}



// TableExists determines whether a DynamoDB table exists.
func (basics TableBasics) tableExists(ctx context.Context) (bool, error) {
	exists := true
	_, err := basics.DynamoDbClient.DescribeTable(
		ctx, &dynamodb.DescribeTableInput{TableName: aws.String(basics.TableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", basics.TableName)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", basics.TableName, err)
		}
		exists = false
	}
	return exists, err
}