package main

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	github "github.com/ren3gadem4rm0t/github-hook-types-go"
)

func init() {
	// Configure zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Use a pretty console writer when in development
	if os.Getenv("AWS_EXECUTION_ENV") == "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

// Response represents the Lambda response format
type Response events.APIGatewayProxyResponse

// Handler function processes GitHub webhook events through AWS Lambda
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	// Get webhook secret from environment
	webhookSecret := os.Getenv("GITHUB_WEBHOOK_SECRET")

	// Extract event headers
	eventType := request.Headers["X-GitHub-Event"]
	if eventType == "" {
		// Check for lowercase headers as API Gateway might normalize them
		eventType = request.Headers["x-github-event"]
	}

	deliveryID := request.Headers["X-GitHub-Delivery"]
	if deliveryID == "" {
		deliveryID = request.Headers["x-github-delivery"]
	}

	// Log incoming webhook
	log.Info().
		Str("event_type", eventType).
		Str("delivery_id", deliveryID).
		Msg("Received GitHub webhook event")

	// Verify webhook signature if secret is provided
	if webhookSecret != "" {
		// Get signature from headers
		signature := request.Headers["X-Hub-Signature-256"]
		if signature == "" {
			signature = request.Headers["x-hub-signature-256"]
		}

		if signature == "" {
			log.Error().Msg("Missing signature in webhook request")
			return createErrorResponse(400, "Missing signature"), nil
		}

		// Manually validate the signature since we don't have an http.Request
		if !strings.HasPrefix(signature, "sha256=") {
			log.Error().Msg("Invalid signature format")
			return createErrorResponse(401, "Invalid signature format"), nil
		}

		sig, err := hex.DecodeString(strings.TrimPrefix(signature, "sha256="))
		if err != nil {
			log.Error().Err(err).Msg("Error decoding signature")
			return createErrorResponse(401, "Invalid signature"), nil
		}

		mac := hmac.New(sha256.New, []byte(webhookSecret))
		mac.Write([]byte(request.Body))
		expectedMAC := mac.Sum(nil)

		if !hmac.Equal(sig, expectedMAC) {
			log.Error().Msg("Signature validation failed")
			return createErrorResponse(401, "Invalid signature"), nil
		}
	}

	// Parse webhook based on event type
	var parsedPayload interface{}

	switch github.WebhookEventType(eventType) {
	case github.CheckRunEvent:
		parsedPayload = &github.CheckRunPayload{}
	case github.CheckSuiteEvent:
		parsedPayload = &github.CheckSuitePayload{}
	case github.CommitCommentEvent:
		parsedPayload = &github.CommitCommentPayload{}
	case github.ContentReferenceEvent:
		parsedPayload = &github.ContentReferencePayload{}
	case github.CreateEvent:
		parsedPayload = &github.CreatePayload{}
	case github.DeleteEvent:
		parsedPayload = &github.DeletePayload{}
	case github.DeployKeyEvent:
		parsedPayload = &github.DeployKeyPayload{}
	case github.DeploymentEvent:
		parsedPayload = &github.DeploymentPayload{}
	case github.DeploymentStatusEvent:
		parsedPayload = &github.DeploymentStatusPayload{}
	case github.DiscussionEvent:
		parsedPayload = &github.DiscussionPayload{}
	case github.DiscussionCommentEvent:
		parsedPayload = &github.DiscussionCommentPayload{}
	case github.ForkEvent:
		parsedPayload = &github.ForkPayload{}
	case github.GitHubAppAuthorizationEvent:
		parsedPayload = &github.AppAuthorizationPayload{}
	case github.GollumEvent:
		parsedPayload = &github.GollumPayload{}
	case github.InstallationEvent:
		parsedPayload = &github.InstallationPayload{}
	case github.InstallationRepositoriesEvent:
		parsedPayload = &github.InstallationRepositoriesPayload{}
	case github.IssueCommentEvent:
		parsedPayload = &github.IssueCommentPayload{}
	case github.IssuesEvent:
		parsedPayload = &github.IssuesPayload{}
	case github.LabelEvent:
		parsedPayload = &github.LabelPayload{}
	case github.MarketplacePurchaseEvent:
		parsedPayload = &github.MarketplacePurchasePayload{}
	case github.MemberEvent:
		parsedPayload = &github.MemberPayload{}
	case github.MembershipEvent:
		parsedPayload = &github.MembershipPayload{}
	case github.MetaEvent:
		parsedPayload = &github.MetaPayload{}
	case github.MilestoneEvent:
		parsedPayload = &github.MilestonePayload{}
	case github.OrganizationEvent:
		parsedPayload = &github.OrganizationPayload{}
	case github.OrgBlockEvent:
		parsedPayload = &github.OrgBlockPayload{}
	case github.PackageEvent:
		parsedPayload = &github.PackagePayload{}
	case github.PageBuildEvent:
		parsedPayload = &github.PageBuildPayload{}
	case github.PingEvent:
		parsedPayload = &github.PingPayload{}
	case github.ProjectEvent:
		parsedPayload = &github.ProjectPayload{}
	case github.ProjectCardEvent:
		parsedPayload = &github.ProjectCardPayload{}
	case github.ProjectColumnEvent:
		parsedPayload = &github.ProjectColumnPayload{}
	case github.PublicEvent:
		parsedPayload = &github.PublicPayload{}
	case github.PullRequestEvent:
		parsedPayload = &github.PullRequestPayload{}
	case github.PullRequestReviewEvent:
		parsedPayload = &github.PullRequestReviewPayload{}
	case github.PullRequestReviewCommentEvent:
		parsedPayload = &github.PullRequestReviewCommentPayload{}
	case github.PushEvent:
		parsedPayload = &github.PushPayload{}
	case github.ReleaseEvent:
		parsedPayload = &github.ReleasePayload{}
	case github.RegistryPackageEvent:
		parsedPayload = &github.RegistryPackagePayload{}
	case github.RepositoryDispatchEvent:
		parsedPayload = &github.RepositoryDispatchPayload{}
	case github.RepositoryEvent:
		parsedPayload = &github.RepositoryPayload{}
	case github.RepositoryImportEvent:
		parsedPayload = &github.RepositoryImportPayload{}
	case github.RepositoryVulnerabilityAlertEvent:
		parsedPayload = &github.RepositoryVulnerabilityAlertPayload{}
	case github.SecurityAdvisoryEvent:
		parsedPayload = &github.SecurityAdvisoryPayload{}
	case github.SponsorshipEvent:
		parsedPayload = &github.SponsorshipPayload{}
	case github.StarEvent:
		parsedPayload = &github.StarPayload{}
	case github.StatusEvent:
		parsedPayload = &github.StatusPayload{}
	case github.TeamEvent:
		parsedPayload = &github.TeamPayload{}
	case github.TeamAddEvent:
		parsedPayload = &github.TeamAddPayload{}
	case github.WatchEvent:
		parsedPayload = &github.WatchPayload{}
	case github.WorkflowDispatchEvent:
		parsedPayload = &github.WorkflowDispatchPayload{}
	case github.WorkflowJobEvent:
		parsedPayload = &github.WorkflowJobPayload{}
	case github.WorkflowRunEvent:
		parsedPayload = &github.WorkflowRunPayload{}
	default:
		parsedPayload = &github.WebhookPayload{}
	}

	// Unmarshal the webhook payload
	if err := json.Unmarshal([]byte(request.Body), &parsedPayload); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal payload")
		return createErrorResponse(400, "Failed to unmarshal payload"), nil
	}

	// Process the webhook event
	if err := processWebhook(eventType, deliveryID, parsedPayload); err != nil {
		log.Error().Err(err).Msg("Failed to process webhook")
		return createErrorResponse(500, "Failed to process webhook"), nil
	}

	// Return a successful response
	return Response{
		StatusCode: 202,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"status": "success", "event": "` + eventType + `", "delivery_id": "` + deliveryID + `"}`,
	}, nil
}

// processWebhook handles different webhook event types
func processWebhook(eventType, deliveryID string, payload interface{}) error {
	switch webhook := payload.(type) {
	case *github.PingPayload:
		log.Info().
			Str("zen", webhook.Zen).
			Int64("hook_id", webhook.HookID).
			Msg("Ping received")

	case *github.PushPayload:
		log.Info().
			Str("repo", webhook.Repository.FullName).
			Str("ref", webhook.Ref).
			Int("commits", len(webhook.Commits)).
			Msg("Push event received")

		// Handle push event (example: trigger a build process)
		// triggerBuild(webhook.Repository.FullName, webhook.Ref)

	case *github.IssuesPayload:
		log.Info().
			Str("action", webhook.Action).
			Int("issue_number", webhook.Issue.Number).
			Str("title", webhook.Issue.Title).
			Msg("Issue event received")

		// Example: Post to Slack when issues are created
		// if webhook.Action == "opened" {
		//     notifySlack("New issue: " + webhook.Issue.Title)
		// }

	case *github.PullRequestPayload:
		log.Info().
			Str("action", webhook.Action).
			Int("pr_number", webhook.PullRequest.Number).
			Str("title", webhook.PullRequest.Title).
			Msg("Pull request event received")

		// Example: Start code review process
		// if webhook.Action == "opened" || webhook.Action == "synchronize" {
		//     startCodeReview(webhook.Repository.FullName, webhook.PullRequest.Number)
		// }

	default:
		// For other event types, just log that we received them
		log.Info().
			Str("event_type", eventType).
			Str("delivery_id", deliveryID).
			Msg("Unhandled event type received")
	}

	return nil
}

// createErrorResponse builds an error response with specified status code and message
func createErrorResponse(statusCode int, message string) Response {
	body, _ := json.Marshal(map[string]string{
		"error": message,
	})

	return Response{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}

func main() {
	lambda.Start(Handler)
}
