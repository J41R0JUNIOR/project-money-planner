# Project Money Planner - Serverless AWS Edition

A Go-based serverless money planning application deployed on AWS Lambda with DynamoDB.

## Architecture

```
API Gateway
    ↓
Lambda Function (Go)
    ↓
DynamoDB (Users, Transactions, Projections)
```

## Prerequisites

- AWS CLI v2
- AWS SAM CLI
- Go 1.26.2 or later
- An AWS account with appropriate permissions

## Project Structure

```
.
├── main.go                          # Lambda handler
├── go.mod                           # Go modules
├── go.sum                           # Module dependencies
├── template.yaml                    # SAM/CloudFormation template
├── build.sh                         # Build script
├── src/
│   ├── domain/                      # Domain models
│   │   ├── User.go
│   │   ├── Transaction.go
│   │   ├── Projection.go
│   │   └── TypeTransaction.go
│   ├── repositories/                # Repository interfaces
│   │   ├── UserRepository.go
│   │   ├── TransactionRepository.go
│   │   ├── ProjectionRepository.go
│   │   └── dynamodb/                # DynamoDB implementations
│   │       ├── user_repository.go
│   │       ├── transaction_repository.go
│   │       └── projection_repository.go
│   └── handlers/                    # Lambda HTTP handlers
│       ├── user_handler.go
│       ├── transaction_handler.go
│       └── projection_handler.go
```

## API Endpoints

### Users
- `POST /users` - Create a new user
- `GET /users/{id}` - Get user by ID
- `DELETE /users/{id}` - Delete a user

### Transactions
- `POST /transactions` - Create a new transaction
- `GET /transactions` - List all transactions
- `GET /transactions/{id}` - Get transaction by ID
- `DELETE /transactions/{id}` - Delete a transaction

### Projections
- `POST /projections` - Create a new projection
- `GET /projections` - List all projections
- `GET /projections/{date}` - Get projection by date
- `DELETE /projections/{date}` - Delete a projection

## Local Testing

### 1. Build the project
```bash
chmod +x build.sh
./build.sh
```

### 2. Test with SAM local

```bash
# Start SAM local
sam local start-api

# Test endpoints (in another terminal)
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"ID": 1, "Username": "john", "Email": "john@example.com", "Password": "secret"}'
```

## Deployment to AWS

### 1. Build for Lambda
```bash
./build.sh
```

### 2. Deploy with SAM
```bash
sam deploy \
  --template-file template.yaml \
  --stack-name project-money-planner \
  --region us-east-1 \
  --parameter-overrides Environment=prod \
  --capabilities CAPABILITY_IAM
```

### 3. Get API endpoint
```bash
aws cloudformation describe-stacks \
  --stack-name project-money-planner \
  --query 'Stacks[0].Outputs[?OutputKey==`ApiEndpoint`].OutputValue' \
  --output text
```

## Environment Variables

The Lambda function expects these environment variables (set automatically by CloudFormation):

- `USERS_TABLE` - DynamoDB Users table name
- `TRANSACTIONS_TABLE` - DynamoDB Transactions table name
- `PROJECTIONS_TABLE` - DynamoDB Projections table name

## AWS Services Used

- **AWS Lambda** - Serverless compute
- **API Gateway** - HTTP API
- **DynamoDB** - NoSQL database
- **CloudFormation** - Infrastructure as Code

## Cost Optimization

The template uses:
- **DynamoDB on-demand billing** - Pay only for what you use
- **Lambda** - Automatically scales, pay per execution
- **API Gateway** - RESTful API with automatic caching

## Example Requests

### Create User
```bash
curl -X POST https://api-endpoint/users \
  -H "Content-Type: application/json" \
  -d '{
    "ID": 1,
    "Username": "john_doe",
    "Email": "john@example.com",
    "Password": "secure_password"
  }'
```

### Create Transaction
```bash
curl -X POST https://api-endpoint/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "ID": 1,
    "Type": "EXPENSE",
    "Amount": 50.00,
    "Date": "2024-01-15",
    "Description": "Grocery shopping",
    "IsRecurring": false
  }'
```

### Create Projection
```bash
curl -X POST https://api-endpoint/projections \
  -H "Content-Type: application/json" \
  -d '{
    "Date": "2024-02-01T00:00:00Z",
    "Balance": 5000.00
  }'
```

## Next Steps

1. Customize domain models as needed
2. Add authentication/authorization (AWS Cognito)
3. Add logging (CloudWatch)
4. Add monitoring (CloudWatch Alarms)
5. Set up CI/CD pipeline (GitHub Actions, CodePipeline)

## Troubleshooting

### Lambda cold starts
- Consider using reserved concurrency
- Pre-warm Lambda function with scheduled events

### DynamoDB costs
- Monitor table throughput
- Enable point-in-time recovery for prod
- Set up auto-scaling if needed

### API Gateway issues
- Check CloudWatch logs in Lambda
- Verify IAM permissions for DynamoDB access

## License

MIT
