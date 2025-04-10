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
	http.HandleFunc("/api/webhook/github", handler.HandleWebhook(handleGitHubEvent))

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

	// Handle different event types
	switch event.Type {
	case github.PingEvent:
		payload := event.Payload.(*github.PingPayload)
		log.Printf("Ping received! Zen: %s", payload.Zen)

	case github.PushEvent:
		payload := event.Payload.(*github.PushPayload)
		log.Printf("Push to %s by %s", payload.Repository.FullName, payload.PusherPerson.Name)

		// Log each commit
		for _, commit := range payload.Commits {
			log.Printf("  Commit: %s", commit.Message)
		}

	case github.IssuesEvent:
		payload := event.Payload.(*github.IssuesPayload)
		log.Printf("Issue #%d %s by %s: %s",
			payload.Issue.Number,
			payload.Action,
			payload.Issue.User.Login,
			payload.Issue.Title)

	case github.IssueCommentEvent:
		payload := event.Payload.(*github.IssueCommentPayload)
		log.Printf("Comment on issue #%d %s by %s: %s",
			payload.Issue.Number,
			payload.Action,
			payload.Comment.User.Login,
			payload.Comment.Body)

	case github.PullRequestEvent:
		payload := event.Payload.(*github.PullRequestPayload)
		log.Printf("Pull request #%d %s by %s: %s",
			payload.PullRequest.Number,
			payload.Action,
			payload.PullRequest.User.Login,
			payload.PullRequest.Title)

	case github.WatchEvent:
		payload := event.Payload.(*github.WatchPayload)
		log.Printf("Repository %s starred by %s",
			payload.Repository.FullName,
			payload.Sender.Login)

	case github.ReleaseEvent:
		payload := event.Payload.(*github.ReleasePayload)
		log.Printf("Release %s %s by %s",
			payload.Release.TagName,
			payload.Action,
			payload.Release.Author.Login)

	default:
		log.Printf("Received unhandled event type: %s", event.Type)
	}

	return nil
}
