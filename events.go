package github

import (
	"crypto/hmac"
	"crypto/sha1" // #nosec G505 - keeping for backward compatibility with GitHub API
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// WebhookEventType represents a GitHub webhook event type.
type WebhookEventType string

// GitHub webhook event types.
const (
	CheckRunEvent                     WebhookEventType = "check_run"
	CheckSuiteEvent                   WebhookEventType = "check_suite"
	CommitCommentEvent                WebhookEventType = "commit_comment"
	ContentReferenceEvent             WebhookEventType = "content_reference"
	CreateEvent                       WebhookEventType = "create"
	DeleteEvent                       WebhookEventType = "delete"
	DeployKeyEvent                    WebhookEventType = "deploy_key"
	DeploymentEvent                   WebhookEventType = "deployment"
	DeploymentStatusEvent             WebhookEventType = "deployment_status"
	DiscussionEvent                   WebhookEventType = "discussion"
	DiscussionCommentEvent            WebhookEventType = "discussion_comment"
	ForkEvent                         WebhookEventType = "fork"
	GitHubAppAuthorizationEvent       WebhookEventType = "github_app_authorization"
	GollumEvent                       WebhookEventType = "gollum"
	InstallationEvent                 WebhookEventType = "installation"
	InstallationRepositoriesEvent     WebhookEventType = "installation_repositories"
	IssueCommentEvent                 WebhookEventType = "issue_comment"
	IssuesEvent                       WebhookEventType = "issues"
	LabelEvent                        WebhookEventType = "label"
	MarketplacePurchaseEvent          WebhookEventType = "marketplace_purchase"
	MemberEvent                       WebhookEventType = "member"
	MembershipEvent                   WebhookEventType = "membership"
	MetaEvent                         WebhookEventType = "meta"
	MilestoneEvent                    WebhookEventType = "milestone"
	OrganizationEvent                 WebhookEventType = "organization"
	OrgBlockEvent                     WebhookEventType = "org_block"
	PackageEvent                      WebhookEventType = "package"
	PageBuildEvent                    WebhookEventType = "page_build"
	PingEvent                         WebhookEventType = "ping"
	ProjectEvent                      WebhookEventType = "project"
	ProjectCardEvent                  WebhookEventType = "project_card"
	ProjectColumnEvent                WebhookEventType = "project_column"
	PublicEvent                       WebhookEventType = "public"
	PullRequestEvent                  WebhookEventType = "pull_request"
	PullRequestReviewEvent            WebhookEventType = "pull_request_review"
	PullRequestReviewCommentEvent     WebhookEventType = "pull_request_review_comment"
	PushEvent                         WebhookEventType = "push"
	ReleaseEvent                      WebhookEventType = "release"
	RegistryPackageEvent              WebhookEventType = "registry_package"
	RepositoryDispatchEvent           WebhookEventType = "repository_dispatch"
	RepositoryEvent                   WebhookEventType = "repository"
	RepositoryImportEvent             WebhookEventType = "repository_import"
	RepositoryVulnerabilityAlertEvent WebhookEventType = "repository_vulnerability_alert"
	SecurityAdvisoryEvent             WebhookEventType = "security_advisory"
	SponsorshipEvent                  WebhookEventType = "sponsorship"
	StarEvent                         WebhookEventType = "star"
	StatusEvent                       WebhookEventType = "status"
	TeamEvent                         WebhookEventType = "team"
	TeamAddEvent                      WebhookEventType = "team_add"
	WatchEvent                        WebhookEventType = "watch"
	WorkflowDispatchEvent             WebhookEventType = "workflow_dispatch"
	WorkflowJobEvent                  WebhookEventType = "workflow_job"
	WorkflowRunEvent                  WebhookEventType = "workflow_run"
)

// WebhookEventHeader is the HTTP header key used to determine the webhook event type.
const WebhookEventHeader = "X-GitHub-Event"

// WebhookDeliveryHeader is the HTTP header key used for the unique webhook delivery ID.
const WebhookDeliveryHeader = "X-GitHub-Delivery"

// WebhookSignatureHeader is the HTTP header key used for webhook signature validation.
const WebhookSignatureHeader = "X-Hub-Signature"

// WebhookSignatureHeader256 is the HTTP header key used for SHA-256 webhook signature validation.
const WebhookSignatureHeader256 = "X-Hub-Signature-256"

// GetEventType extracts the webhook event type from the HTTP request headers.
func GetEventType(r *http.Request) WebhookEventType {
	return WebhookEventType(r.Header.Get(WebhookEventHeader))
}

// GetDeliveryID extracts the webhook delivery ID from the HTTP request headers.
func GetDeliveryID(r *http.Request) string {
	return r.Header.Get(WebhookDeliveryHeader)
}

// WebhookEvent contains metadata about the webhook event, including its type,
// delivery ID, and payload.
type WebhookEvent struct {
	Type       WebhookEventType
	DeliveryID string
	Payload    interface{}
}

// ParseWebhook parses a webhook from an HTTP request, identifying the event type
// and appropriate payload structure. It returns an error if the event type is
// unknown or if the payload cannot be parsed.
func ParseWebhook(r *http.Request) (*WebhookEvent, error) {
	eventType := GetEventType(r)
	deliveryID := GetDeliveryID(r)

	// Verify we have a known event type
	if eventType == "" {
		return nil, fmt.Errorf("missing event type in headers")
	}

	// Read and parse the request body
	var payload interface{}
	var err error

	switch eventType {
	case CheckRunEvent:
		payload = new(CheckRunPayload)
	case CheckSuiteEvent:
		payload = new(CheckSuitePayload)
	case CommitCommentEvent:
		payload = new(CommitCommentPayload)
	case ContentReferenceEvent:
		payload = new(ContentReferencePayload)
	case CreateEvent:
		payload = new(CreatePayload)
	case DeleteEvent:
		payload = new(DeletePayload)
	case DeployKeyEvent:
		payload = new(DeployKeyPayload)
	case DeploymentEvent:
		payload = new(DeploymentPayload)
	case DeploymentStatusEvent:
		payload = new(DeploymentStatusPayload)
	case DiscussionEvent:
		payload = new(DiscussionPayload)
	case DiscussionCommentEvent:
		payload = new(DiscussionCommentPayload)
	case ForkEvent:
		payload = new(ForkPayload)
	case GitHubAppAuthorizationEvent:
		payload = new(AppAuthorizationPayload)
	case GollumEvent:
		payload = new(GollumPayload)
	case InstallationEvent:
		payload = new(InstallationPayload)
	case InstallationRepositoriesEvent:
		payload = new(InstallationRepositoriesPayload)
	case IssueCommentEvent:
		payload = new(IssueCommentPayload)
	case IssuesEvent:
		payload = new(IssuesPayload)
	case LabelEvent:
		payload = new(LabelPayload)
	case MarketplacePurchaseEvent:
		payload = new(MarketplacePurchasePayload)
	case MemberEvent:
		payload = new(MemberPayload)
	case MembershipEvent:
		payload = new(MembershipPayload)
	case MetaEvent:
		payload = new(MetaPayload)
	case MilestoneEvent:
		payload = new(MilestonePayload)
	case OrganizationEvent:
		payload = new(OrganizationPayload)
	case OrgBlockEvent:
		payload = new(OrgBlockPayload)
	case PackageEvent:
		payload = new(PackagePayload)
	case PageBuildEvent:
		payload = new(PageBuildPayload)
	case PingEvent:
		payload = new(PingPayload)
	case ProjectEvent:
		payload = new(ProjectPayload)
	case ProjectCardEvent:
		payload = new(ProjectCardPayload)
	case ProjectColumnEvent:
		payload = new(ProjectColumnPayload)
	case PublicEvent:
		payload = new(PublicPayload)
	case PullRequestEvent:
		payload = new(PullRequestPayload)
	case PullRequestReviewEvent:
		payload = new(PullRequestReviewPayload)
	case PullRequestReviewCommentEvent:
		payload = new(PullRequestReviewCommentPayload)
	case PushEvent:
		payload = new(PushPayload)
	case ReleaseEvent:
		payload = new(ReleasePayload)
	case RegistryPackageEvent:
		payload = new(RegistryPackagePayload)
	case RepositoryDispatchEvent:
		payload = new(RepositoryDispatchPayload)
	case RepositoryEvent:
		payload = new(RepositoryPayload)
	case RepositoryImportEvent:
		payload = new(RepositoryImportPayload)
	case RepositoryVulnerabilityAlertEvent:
		payload = new(RepositoryVulnerabilityAlertPayload)
	case SecurityAdvisoryEvent:
		payload = new(SecurityAdvisoryPayload)
	case SponsorshipEvent:
		payload = new(SponsorshipPayload)
	case StarEvent:
		payload = new(StarPayload)
	case StatusEvent:
		payload = new(StatusPayload)
	case TeamEvent:
		payload = new(TeamPayload)
	case TeamAddEvent:
		payload = new(TeamAddPayload)
	case WatchEvent:
		payload = new(WatchPayload)
	case WorkflowDispatchEvent:
		payload = new(WorkflowDispatchPayload)
	case WorkflowJobEvent:
		payload = new(WorkflowJobPayload)
	case WorkflowRunEvent:
		payload = new(WorkflowRunPayload)
	default:
		// If we don't recognize the event type, use a generic payload structure
		payload = new(WebhookPayload)
	}

	// Decode the JSON payload
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, fmt.Errorf("failed to parse webhook payload: %v", err)
	}

	return &WebhookEvent{
		Type:       eventType,
		DeliveryID: deliveryID,
		Payload:    payload,
	}, nil
}

// ValidateSignature validates the webhook signature against the payload and secret.
// It supports both SHA-1 and SHA-256 signatures.
func ValidateSignature(r *http.Request, payload []byte, secret string) error {
	if secret == "" {
		// No secret configured, so signature validation is skipped
		return nil
	}

	secretBytes := []byte(secret)

	// Try SHA-256 first (preferred)
	signature256 := r.Header.Get(WebhookSignatureHeader256)
	if signature256 != "" {
		// Validate the SHA-256 signature
		if !strings.HasPrefix(signature256, "sha256=") {
			return errors.New("invalid SHA-256 signature format")
		}

		sig, err := hex.DecodeString(strings.TrimPrefix(signature256, "sha256="))
		if err != nil {
			return fmt.Errorf("error decoding SHA-256 signature: %v", err)
		}

		mac := hmac.New(sha256.New, secretBytes)
		_, _ = mac.Write(payload)
		expectedMAC := mac.Sum(nil)

		if !hmac.Equal(sig, expectedMAC) {
			return errors.New("SHA-256 signature validation failed")
		}

		return nil
	}

	// Fall back to SHA-1 if SHA-256 is not present
	signature := r.Header.Get(WebhookSignatureHeader)
	if signature != "" {
		// Log a warning about SHA-1 usage
		// We use fmt.Fprintf to stderr since this is a library function and we don't want to
		// enforce a specific logging package on users
		fmt.Fprintf(os.Stderr, "WARNING: Using deprecated SHA-1 signature validation. Configure your webhook to use SHA-256.\n")

		// Validate the SHA-1 signature
		if !strings.HasPrefix(signature, "sha1=") {
			return errors.New("invalid SHA-1 signature format")
		}

		sig, err := hex.DecodeString(strings.TrimPrefix(signature, "sha1="))
		if err != nil {
			return fmt.Errorf("error decoding SHA-1 signature: %v", err)
		}

		// #nosec G401 - keeping SHA-1 for backward compatibility with GitHub API
		mac := hmac.New(sha1.New, secretBytes)
		_, _ = mac.Write(payload)
		expectedMAC := mac.Sum(nil)

		if !hmac.Equal(sig, expectedMAC) {
			return errors.New("SHA-1 signature validation failed")
		}

		return nil
	}

	return errors.New("missing signature headers")
}
