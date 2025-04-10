# GitHub Webhook Handler with Gin and Zerolog

This example demonstrates how to create a GitHub webhook handler using the Gin web framework and Zerolog for logging. It shows how to:

1. Create middleware for logging incoming requests
2. Validate GitHub webhook signatures
3. Parse and process different webhook event types
4. Structure a Gin-based webhook server

## Features

- **Complete Event Coverage**: Handles all GitHub webhook event types with properly typed structs
- **Structured Logging**: Uses Zerolog for contextual, structured logging of all events
- **Signature Validation**: Validates webhook signatures for security
- **RESTful API Design**: Follows best practices for HTTP API design
- **Context Preservation**: Maintains request context between middleware

## Setup

First, make sure you have Go installed (version 1.19 or higher). Then, initialize the Go modules:

```bash
go mod tidy
```

## Running the Server

Start the server with:

```bash
go run main.go
```

By default, the server listens on port 3000. You can change this by setting the `PORT` environment variable.

## Environment Variables

- `PORT`: The port on which the server listens (default: 3000)
- `WEBHOOK_SECRET`: The secret used to validate GitHub webhook signatures

Example:

```bash
WEBHOOK_SECRET=your_webhook_secret PORT=8080 go run main.go
```

## Endpoints

- `GET /health`: Health check endpoint
- `POST /api/webhooks/github`: GitHub webhook endpoint

## Supported Event Types

This example handles all GitHub webhook event types:

- Issue events
- Pull request events
- Push events
- Repository events
- Deployment events
- Security events
- Release events
- Workflow events
- Discussion events
- Wiki events
- Installation events
- Team events
- Organization events
- Project events
- ...and many more!

## Testing with GitHub

1. Create a webhook in your GitHub repository with:
   - Payload URL: `https://your-server-url/api/webhooks/github`
   - Content type: `application/json`
   - Secret: The same secret you set in the `WEBHOOK_SECRET` environment variable
   - Events: Choose the events you want to receive (you can select "Send me everything")

2. GitHub will send a ping event when you create the webhook. The server will log this event and respond with a success message.

## Custom Logging

The example uses Zerolog for structured logging with customized field names for each event type. The logs include:

- Common fields: event type, repository, action, delivery ID
- Event-specific fields: PR number, issue title, commit SHAs, etc.
- Timestamp and request details
- Error context for failures

You can modify the log format in the `setupLogger` function to output JSON instead of the console format if you're sending logs to a log aggregator.
