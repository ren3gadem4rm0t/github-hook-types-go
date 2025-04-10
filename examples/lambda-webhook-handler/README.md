# GitHub Webhook AWS Lambda Handler

This example demonstrates how to use the `github-hook-types-go` library to handle GitHub webhook events using AWS Lambda and API Gateway.

## Features

- Processes incoming GitHub webhook events via AWS Lambda
- Verifies webhook signatures for security
- Handles all GitHub webhook event types
- Structured logging with zerolog
- Example event handlers for common webhooks (push, issues, pull requests)

## Prerequisites

- Go 1.19 or later
- AWS account
- AWS CLI configured locally

## Setup

1. Clone this repository
2. Build the Lambda function:

```bash
cd examples/lambda-webhook-handler
go mod tidy
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
zip lambda-function.zip bootstrap
```

## Deployment

You can deploy this Lambda function using the AWS CLI, AWS SAM, Terraform, or the AWS Console.

### AWS CLI Example

```bash
aws lambda create-function \
  --function-name github-webhook-handler \
  --runtime provided.al2 \
  --handler bootstrap \
  --role arn:aws:iam::ACCOUNT_ID:role/lambda-execution-role \
  --zip-file fileb://lambda-function.zip \
  --environment Variables={GITHUB_WEBHOOK_SECRET=your_webhook_secret}
```

### API Gateway Setup

1. Create a new REST API in API Gateway
2. Create a resource (e.g., `/webhook`)
3. Create a POST method and integrate it with your Lambda function
4. Deploy the API to a stage (e.g., `prod`)
5. Use the resulting URL as your GitHub webhook URL

## Setting up GitHub Webhooks

1. Go to your GitHub repository settings
2. Click on "Webhooks" and then "Add webhook"
3. Set the Payload URL to your API Gateway endpoint
4. Set Content type to `application/json`
5. Set the Secret to the same value as the `GITHUB_WEBHOOK_SECRET` environment variable
6. Choose which events you want to trigger the webhook
7. Click "Add webhook"

## Extending the Example

This example provides a basic structure for handling GitHub webhooks. You can extend it by:

1. Implementing more event handlers in the `processWebhook` function
2. Adding integration with other AWS services (SNS, SQS, DynamoDB, etc.)
3. Setting up CloudWatch alarms for monitoring

## Security Considerations

- Always verify webhook signatures to prevent spoofing
- Store your webhook secret in AWS Secrets Manager for production use
- Configure appropriate IAM roles with least privilege
- Enable AWS X-Ray for tracing and debugging

## License

This example is licensed under the same terms as the main library.
