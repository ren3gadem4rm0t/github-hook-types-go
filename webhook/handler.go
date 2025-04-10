// Package webhook provides utilities for handling GitHub webhook deliveries.
package webhook

import (
	"crypto/hmac"
	"crypto/sha1" // #nosec G505 - keeping for backward compatibility with GitHub API
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ren3gadem4rm0t/github-hook-types-go"
)

const (
	// SignatureHeader is the GitHub header containing the HMAC-SHA1 hexdigest.
	// DEPRECATED: Use SignatureHeader256 instead as SHA-1 is considered cryptographically weak.
	SignatureHeader = "X-Hub-Signature"

	// SignatureHeader256 is the GitHub header containing the HMAC-SHA256 hexdigest.
	SignatureHeader256 = "X-Hub-Signature-256"

	// EventTypeHeader is the GitHub header containing the event type.
	EventTypeHeader = "X-GitHub-Event"

	// DeliveryIDHeader is the GitHub header containing the unique webhook delivery ID.
	DeliveryIDHeader = "X-GitHub-Delivery"
)

// Handler processes webhook requests from GitHub.
type Handler struct {
	secret []byte
}

// NewHandler creates a new webhook handler with the given secret.
func NewHandler(secret string) *Handler {
	return &Handler{
		secret: []byte(secret),
	}
}

// ValidateSignature validates the signature in the request against the webhook secret.
// It returns an error if the signature is invalid or missing.
func (h *Handler) ValidateSignature(r *http.Request, payload []byte) error {
	if len(h.secret) == 0 {
		// No secret configured, so signature validation is skipped
		return nil
	}

	signature256 := r.Header.Get(SignatureHeader256)
	if signature256 != "" {
		return h.validateSignatureSHA256(signature256, payload)
	}

	signature := r.Header.Get(SignatureHeader)
	if signature != "" {
		log.Println("WARNING: Using deprecated SHA-1 signature validation. Configure your webhook to use SHA-256.")
		return h.validateSignatureSHA1(signature, payload)
	}

	return errors.New("missing signature headers")
}

// validateSignatureSHA1 validates an HMAC-SHA1 signature.
// DEPRECATED: GitHub is transitioning away from SHA-1. Use SHA-256 when possible.
// #nosec G401 - keeping for backward compatibility with GitHub API
func (h *Handler) validateSignatureSHA1(signature string, payload []byte) error {
	if !strings.HasPrefix(signature, "sha1=") {
		return errors.New("invalid signature format")
	}

	sig, err := hex.DecodeString(strings.TrimPrefix(signature, "sha1="))
	if err != nil {
		return fmt.Errorf("error decoding signature: %v", err)
	}

	mac := hmac.New(sha1.New, h.secret)
	_, _ = mac.Write(payload)
	expectedMAC := mac.Sum(nil)

	if !hmac.Equal(sig, expectedMAC) {
		return errors.New("signature validation failed")
	}

	return nil
}

// validateSignatureSHA256 validates an HMAC-SHA256 signature.
func (h *Handler) validateSignatureSHA256(signature string, payload []byte) error {
	if !strings.HasPrefix(signature, "sha256=") {
		return errors.New("invalid signature format")
	}

	sig, err := hex.DecodeString(strings.TrimPrefix(signature, "sha256="))
	if err != nil {
		return fmt.Errorf("error decoding signature: %v", err)
	}

	mac := hmac.New(sha256.New, h.secret)
	_, _ = mac.Write(payload)
	expectedMAC := mac.Sum(nil)

	if !hmac.Equal(sig, expectedMAC) {
		return errors.New("signature validation failed")
	}

	return nil
}

// ProcessWebhook processes a webhook request and returns the corresponding webhook event.
// It validates the signature if a secret is configured and returns an error if validation fails.
func (h *Handler) ProcessWebhook(r *http.Request) (*github.WebhookEvent, error) {
	// Read the request body
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading request body: %v", err)
	}

	// Use a named return to handle errors from deferred operations
	var webhookEvent *github.WebhookEvent
	var retErr error

	// Defer closing the request body
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil && retErr == nil {
			retErr = fmt.Errorf("error closing request body: %v", closeErr)
		}
	}()

	// Validate the signature
	if err := h.ValidateSignature(r, payload); err != nil {
		return nil, fmt.Errorf("invalid signature: %v", err)
	}

	// Extract event info from headers
	eventType := github.WebhookEventType(r.Header.Get(EventTypeHeader))
	if eventType == "" {
		return nil, errors.New("missing event type header")
	}

	deliveryID := r.Header.Get(DeliveryIDHeader)
	if deliveryID == "" {
		return nil, errors.New("missing delivery ID header")
	}

	// Parse the payload based on the event type
	var parsedPayload any
	switch eventType {
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

	// Unmarshal the payload
	if err := json.Unmarshal(payload, parsedPayload); err != nil {
		return nil, fmt.Errorf("error unmarshaling payload: %v", err)
	}

	webhookEvent = &github.WebhookEvent{
		Type:       eventType,
		DeliveryID: deliveryID,
		Payload:    parsedPayload,
	}
	return webhookEvent, retErr
}

// HandleWebhook is a convenience method that provides an http.HandlerFunc for processing webhooks.
// It calls the provided callback function with the parsed webhook event.
func (h *Handler) HandleWebhook(callback func(*github.WebhookEvent) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := h.ProcessWebhook(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := callback(event); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
