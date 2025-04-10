# GitHub Webhook Types for Go

A comprehensive Go library providing strongly-typed definitions for all GitHub webhook events and payloads. This library helps developers build applications that receive and process GitHub webhook events with proper type safety and full IntelliSense support.

## Features

- Complete coverage of all GitHub webhook event types
- Strongly-typed structs for each webhook payload
- JSON tags for direct unmarshaling of webhook payloads
- Helper functions for webhook signature validation
- Simple HTTP handler for webhook processing
- Detailed documentation for all types and fields

## Installation

```bash
go get github.com/ren3gadem4rm0t/github-hook-types-go
```

## Usage

### Basic Example

Here's a simple example of how to use this library to handle GitHub webhooks:

```go
package main

import (
 "log"
 "net/http"
 "os"

 "github.com/ren3gadem4rm0t/github-hook-types-go"
 "github.com/ren3gadem4rm0t/github-hook-types-go/webhook"
)

func main() {
 // Get webhook secret from environment
 secret := os.Getenv("WEBHOOK_SECRET")
 
 // Create a new webhook handler
 handler := webhook.NewHandler(secret)

 // Set up a webhook handler endpoint
 http.HandleFunc("/webhook", handler.HandleWebhook(handleGitHubEvent))

 // Start the server
 port := os.Getenv("PORT")
 if port == "" {
  port = "3000"
 }
 
 log.Printf("Server starting on port %s...\n", port)
 log.Fatal(http.ListenAndServe(":"+port, nil))
}

// handleGitHubEvent processes GitHub webhook events
func handleGitHubEvent(event *github.WebhookEvent) error {
 log.Printf("Received %s event (Delivery ID: %s)\n", event.Type, event.DeliveryID)

 // Type switch to handle different event types
 switch event.Type {
 case github.PingEvent:
  payload := event.Payload.(*github.PingPayload)
  log.Printf("Ping received! Zen: %s", payload.Zen)
  
 case github.PushEvent:
  payload := event.Payload.(*github.PushPayload)
  log.Printf("Push to %s by %s", 
   payload.Repository.FullName, 
   payload.PusherPerson.Name)
  
 case github.IssuesEvent:
  payload := event.Payload.(*github.IssuesPayload)
  log.Printf("Issue #%d %s by %s: %s",
   payload.Issue.Number,
   payload.Action,
   payload.Issue.User.Login,
   payload.Issue.Title)
   
 default:
  log.Printf("Received unhandled event type: %s", event.Type)
 }

 return nil
}
```

### Manual Webhook Processing

If you need more control over the webhook processing flow:

```go
func webhookHandler(w http.ResponseWriter, r *http.Request) {
 // Create a webhook handler with your secret
 handler := webhook.NewHandler("your-webhook-secret")
 
 // Process the webhook
 event, err := handler.ProcessWebhook(r)
 if err != nil {
  http.Error(w, err.Error(), http.StatusBadRequest)
  return
 }
 
 // Handle the event based on its type
 switch event.Type {
 case github.PushEvent:
  payload := event.Payload.(*github.PushPayload)
  // Process push event...
 case github.IssuesEvent:
  payload := event.Payload.(*github.IssuesPayload)
  // Process issues event...
 // Handle other event types...
 }
 
 w.WriteHeader(http.StatusOK)
}
```

### Using with Gin Framework

For applications using the Gin web framework, check out the [Gin webhook example](examples/gin-webhook-server/) which demonstrates:

- Creating logging middleware with Zerolog
- Setting up webhook signature validation middleware
- Handling GitHub webhook events in a Gin router
- Properly parsing and processing webhook payloads

## Supported Event Types

This library supports all GitHub webhook event types, including:

- `check_run`
- `check_suite`
- `commit_comment`
- `content_reference`
- `create`
- `delete`
- `deploy_key`
- `deployment`
- `deployment_status`
- `discussion`
- `discussion_comment`
- `fork`
- `github_app_authorization`
- `gollum`
- `installation`
- `installation_repositories`
- `issue_comment`
- `issues`
- `label`
- `marketplace_purchase`
- `member`
- `membership`
- `meta`
- `milestone`
- `organization`
- `org_block`
- `package`
- `page_build`
- `ping`
- `project`
- `project_card`
- `project_column`
- `public`
- `pull_request`
- `pull_request_review`
- `pull_request_review_comment`
- `push`
- `release`
- `repository_dispatch`
- `repository`
- `repository_import`
- `repository_vulnerability_alert`
- `security_advisory`
- `sponsorship`
- `star`
- `status`
- `team`
- `team_add`
- `watch`
- `workflow_dispatch`
- `workflow_job`
- `workflow_run`

## License

MIT
