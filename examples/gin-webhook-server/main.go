package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ren3gadem4rm0t/github-hook-types-go"
	"github.com/ren3gadem4rm0t/github-hook-types-go/webhook"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func setupLogger() {
	// Configure error stack marshaling
	zerolog.TimeFieldFormat = time.RFC3339

	// Set the global log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Create a console writer
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Create the logger
	log = zerolog.New(output).With().
		Timestamp().
		Caller().
		Logger()
}

// LoggerMiddleware is a Gin middleware that logs requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log request details
		log.Info().
			Str("method", method).
			Str("path", path).
			Int("status", c.Writer.Status()).
			Dur("latency", latency).
			Str("client-ip", c.ClientIP()).
			Str("user-agent", c.Request.UserAgent()).
			Msg("Request processed")
	}
}

// GithubWebhookMiddleware creates a middleware that processes GitHub webhook events
func GithubWebhookMiddleware(secret string) gin.HandlerFunc {
	handler := webhook.NewHandler(secret)

	return func(c *gin.Context) {
		// Read request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Error().Err(err).Msg("Failed to read request body")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}

		// Close the body and reset it for downstream handlers
		if err := c.Request.Body.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close request body")
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Validate signature
		if err := handler.ValidateSignature(c.Request, body); err != nil {
			log.Error().Err(err).Msg("Invalid webhook signature")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
			return
		}

		// Extract event info
		eventType := github.WebhookEventType(c.GetHeader(webhook.EventTypeHeader))
		deliveryID := c.GetHeader(webhook.DeliveryIDHeader)

		// Store event info in the context for handlers
		c.Set("event_type", string(eventType))
		c.Set("delivery_id", deliveryID)

		// Log the webhook event
		log.Info().
			Str("event_type", string(eventType)).
			Str("delivery_id", deliveryID).
			Msg("Received GitHub webhook event")

		c.Next()
	}
}

// HandleGithubWebhook handles GitHub webhook events
func HandleGithubWebhook(c *gin.Context) {
	// Read request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Get event type and delivery ID from the context
	eventType := c.GetString("event_type")
	deliveryID := c.GetString("delivery_id")

	// Parse the payload based on the event type
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

	// Unmarshal the payload
	if err := json.Unmarshal(body, parsedPayload); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal payload"})
		return
	}

	// Process the webhook event
	switch webhook := parsedPayload.(type) {
	case *github.PingPayload:
		log.Info().Str("zen", webhook.Zen).Int64("hook_id", webhook.HookID).Msg("Ping received")

	case *github.PushPayload:
		log.Info().
			Str("repo", webhook.Repository.FullName).
			Str("pusher", webhook.PusherPerson.Name).
			Str("ref", webhook.Ref).
			Str("before", webhook.Before).
			Str("after", webhook.After).
			Int("commits", len(webhook.Commits)).
			Msg("Push event received")

	case *github.CheckRunPayload:
		conclusion := ""
		if webhook.CheckRun.Conclusion != nil {
			conclusion = *webhook.CheckRun.Conclusion
		}
		log.Info().
			Str("action", webhook.Action).
			Str("name", webhook.CheckRun.Name).
			Str("status", webhook.CheckRun.Status).
			Str("conclusion", conclusion).
			Str("repo", webhook.Repository.FullName).
			Msg("Check run event received")

	case *github.CheckSuitePayload:
		conclusion := ""
		if webhook.CheckSuite.Conclusion != nil {
			conclusion = *webhook.CheckSuite.Conclusion
		}
		log.Info().
			Str("action", webhook.Action).
			Str("status", webhook.CheckSuite.Status).
			Str("conclusion", conclusion).
			Str("head_branch", webhook.CheckSuite.HeadBranch).
			Str("repo", webhook.Repository.FullName).
			Msg("Check suite event received")

	case *github.IssuesPayload:
		log.Info().
			Str("action", webhook.Action).
			Int("issue_number", webhook.Issue.Number).
			Str("title", webhook.Issue.Title).
			Str("user", webhook.Issue.User.Login).
			Str("state", webhook.Issue.State).
			Msg("Issue event received")

	case *github.IssueCommentPayload:
		log.Info().
			Str("action", webhook.Action).
			Int("issue_number", webhook.Issue.Number).
			Str("comment_user", webhook.Comment.User.Login).
			Str("comment_body", webhook.Comment.Body).
			Msg("Issue comment event received")

	case *github.PullRequestPayload:
		log.Info().
			Str("action", webhook.Action).
			Int("pr_number", webhook.PullRequest.Number).
			Str("title", webhook.PullRequest.Title).
			Str("user", webhook.PullRequest.User.Login).
			Str("base_ref", webhook.PullRequest.Base.Ref).
			Str("head_ref", webhook.PullRequest.Head.Ref).
			Bool("merged", webhook.PullRequest.Merged).
			Msg("Pull request event received")

	case *github.PullRequestReviewPayload:
		log.Info().
			Str("action", webhook.Action).
			Int("pr_number", webhook.PullRequest.Number).
			Str("reviewer", webhook.Review.User.Login).
			Str("state", webhook.Review.State).
			Msg("Pull request review event received")

	case *github.PullRequestReviewCommentPayload:
		log.Info().
			Str("action", webhook.Action).
			Int("pr_number", webhook.PullRequest.Number).
			Str("commenter", webhook.Comment.User.Login).
			Str("path", webhook.Comment.Path).
			Msg("Pull request review comment event received")

	case *github.CommitCommentPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("commit_id", webhook.Comment.CommitID).
			Str("commenter", webhook.Comment.User.Login).
			Str("repo", webhook.Repository.FullName).
			Msg("Commit comment event received")

	case *github.CreatePayload:
		log.Info().
			Str("ref", webhook.Ref).
			Str("ref_type", webhook.RefType).
			Str("repo", webhook.Repository.FullName).
			Msg("Create event received")

	case *github.DeletePayload:
		log.Info().
			Str("ref", webhook.Ref).
			Str("ref_type", webhook.RefType).
			Str("repo", webhook.Repository.FullName).
			Msg("Delete event received")

	case *github.DeployKeyPayload:
		log.Info().
			Int64("key_id", webhook.Key.ID).
			Str("title", webhook.Key.Title).
			Msg("Deploy key event received")

	case *github.DeploymentPayload:
		log.Info().
			Str("environment", webhook.Deployment.Environment).
			Str("sha", webhook.Deployment.SHA).
			Str("ref", webhook.Deployment.Ref).
			Str("task", webhook.Deployment.Task).
			Str("description", webhook.Deployment.Description).
			Str("repo", webhook.Repository.FullName).
			Msg("Deployment event received")

	case *github.DeploymentStatusPayload:
		log.Info().
			Str("state", webhook.DeploymentStatus.State).
			Str("environment", webhook.DeploymentStatus.Environment).
			Str("description", webhook.DeploymentStatus.Description).
			Int64("deployment_id", webhook.DeploymentStatus.ID).
			Str("repo", webhook.Repository.FullName).
			Msg("Deployment status event received")

	case *github.ForkPayload:
		log.Info().
			Str("forked_repo", webhook.Repository.FullName).
			Str("fork", webhook.Forkee.FullName).
			Str("forked_by", webhook.Sender.Login).
			Msg("Fork event received")

	case *github.GollumPayload:
		pages := len(webhook.Pages)
		action := ""
		if pages > 0 {
			action = webhook.Pages[0].Action
		}
		log.Info().
			Int("pages_count", pages).
			Str("action", action).
			Str("repo", webhook.Repository.FullName).
			Msg("Wiki page event received")

	case *github.InstallationPayload:
		log.Info().
			Str("action", webhook.Action).
			Int64("installation_id", webhook.Installation.ID).
			Str("account", webhook.Installation.Account.Login).
			Msg("Installation event received")

	case *github.InstallationRepositoriesPayload:
		log.Info().
			Str("action", webhook.Action).
			Int64("installation_id", webhook.Installation.ID).
			Str("account", webhook.Installation.Account.Login).
			Int("repos_added", len(webhook.RepositoriesAdded)).
			Int("repos_removed", len(webhook.RepositoriesRemoved)).
			Msg("Installation repositories event received")

	case *github.LabelPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("label", webhook.Label.Name).
			Str("color", webhook.Label.Color).
			Str("repo", webhook.Repository.FullName).
			Msg("Label event received")

	case *github.MemberPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("member", webhook.Member.Login).
			Str("repo", webhook.Repository.FullName).
			Msg("Member event received")

	case *github.MembershipPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("scope", webhook.Scope).
			Str("member", webhook.Member.Login).
			Str("team", webhook.Team.Name).
			Str("org", webhook.Organization.Login).
			Msg("Membership event received")

	case *github.MilestonePayload:
		log.Info().
			Str("action", webhook.Action).
			Int("number", webhook.Milestone.Number).
			Str("title", webhook.Milestone.Title).
			Str("state", webhook.Milestone.State).
			Str("repo", webhook.Repository.FullName).
			Msg("Milestone event received")

	case *github.OrganizationPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("org", webhook.Organization.Login).
			Str("sender", webhook.Sender.Login).
			Msg("Organization event received")

	case *github.PackagePayload:
		log.Info().
			Str("action", webhook.Action).
			Str("package_name", webhook.Package.Name).
			Str("package_type", webhook.Package.PackageType).
			Str("org", webhook.Organization.Login).
			Msg("Package event received")

	case *github.PageBuildPayload:
		log.Info().
			Str("status", webhook.Build.Status).
			Str("repo", webhook.Repository.FullName).
			Msg("Page build event received")

	case *github.ProjectPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("project_name", webhook.Project.Name).
			Str("state", webhook.Project.State).
			Str("repo", webhook.Repository.FullName).
			Msg("Project event received")

	case *github.ProjectCardPayload:
		log.Info().
			Str("action", webhook.Action).
			Int64("project_id", webhook.ProjectCard.ProjectID).
			Int64("column_id", webhook.ProjectCard.ColumnID).
			Str("repo", webhook.Repository.FullName).
			Msg("Project card event received")

	case *github.ProjectColumnPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("column_name", webhook.ProjectColumn.Name).
			Str("repo", webhook.Repository.FullName).
			Msg("Project column event received")

	case *github.PublicPayload:
		log.Info().
			Str("repo", webhook.Repository.FullName).
			Str("sender", webhook.Sender.Login).
			Msg("Repository made public")

	case *github.ReleasePayload:
		log.Info().
			Str("action", webhook.Action).
			Str("tag", webhook.Release.TagName).
			Str("name", webhook.Release.Name).
			Bool("draft", webhook.Release.Draft).
			Bool("prerelease", webhook.Release.Prerelease).
			Str("repo", webhook.Repository.FullName).
			Msg("Release event received")

	case *github.RepositoryPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("repo", webhook.Repository.FullName).
			Str("sender", webhook.Sender.Login).
			Msg("Repository event received")

	case *github.RepositoryVulnerabilityAlertPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("repo", webhook.Repository.FullName).
			Str("affected_package", webhook.Alert.AffectedPackageName).
			Str("affected_range", webhook.Alert.AffectedRange).
			Str("fixed_in", webhook.Alert.FixedIn).
			Msg("Repository vulnerability alert received")

	case *github.SecurityAdvisoryPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("ghsa_id", webhook.SecurityAdvisory.GHSAID).
			Str("summary", webhook.SecurityAdvisory.Summary).
			Str("severity", webhook.SecurityAdvisory.Severity).
			Msg("Security advisory event received")

	case *github.StarPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("repo", webhook.Repository.FullName).
			Str("user", webhook.Sender.Login).
			Msg("Star event received")

	case *github.StatusPayload:
		log.Info().
			Str("state", webhook.State).
			Str("context", webhook.Context).
			Str("description", webhook.Description).
			Str("sha", webhook.SHA).
			Str("repo", webhook.Repository.FullName).
			Msg("Status event received")

	case *github.TeamPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("team", webhook.Team.Name).
			Str("org", webhook.Organization.Login).
			Msg("Team event received")

	case *github.TeamAddPayload:
		log.Info().
			Str("team", webhook.Team.Name).
			Str("repo", webhook.Repository.FullName).
			Str("org", webhook.Organization.Login).
			Msg("Team add event received")

	case *github.WatchPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("repo", webhook.Repository.FullName).
			Str("user", webhook.Sender.Login).
			Msg("Watch event received")

	case *github.WorkflowDispatchPayload:
		log.Info().
			Str("workflow", webhook.Workflow).
			Str("ref", webhook.Ref).
			Str("repo", webhook.Repository.FullName).
			Msg("Workflow dispatch event received")

	case *github.WorkflowJobPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("workflow_name", webhook.WorkflowJob.Name).
			Str("status", webhook.WorkflowJob.Status).
			Str("conclusion", webhook.WorkflowJob.Conclusion).
			Int64("run_id", webhook.WorkflowJob.RunID).
			Str("repo", webhook.Repository.FullName).
			Msg("Workflow job event received")

	case *github.WorkflowRunPayload:
		log.Info().
			Str("action", webhook.Action).
			Str("workflow_name", webhook.WorkflowRun.Name).
			Str("event", webhook.WorkflowRun.Event).
			Str("status", webhook.WorkflowRun.Status).
			Str("conclusion", webhook.WorkflowRun.Conclusion).
			Str("repo", webhook.Repository.FullName).
			Msg("Workflow run event received")

	case *github.RegistryPackagePayload:
		log.Info().
			Str("action", webhook.Action).
			Str("package_name", webhook.RegistryPackage.Name).
			Str("package_type", webhook.RegistryPackage.PackageType).
			Str("version", webhook.RegistryPackage.PackageVersion.Version).
			Str("owner", webhook.RegistryPackage.Owner.Login).
			Msg("Registry package event received")

	case *github.RepositoryDispatchPayload:
		clientPayload := "null"
		if webhook.ClientPayload != nil {
			clientPayload = string(webhook.ClientPayload)
		}
		log.Info().
			Str("action", webhook.Action).
			Str("branch", webhook.Branch).
			Str("client_payload", clientPayload).
			Str("repo", webhook.Repository.FullName).
			Msg("Repository dispatch event received")

	default:
		log.Warn().
			Str("event_type", eventType).
			Msg("Unhandled event type. Consider adding a specific handler for this event type in the main.go switch statement.")
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "event": eventType, "delivery_id": deliveryID})
}

func main() {
	// Setup logger
	setupLogger()

	// Get webhook secret from environment
	secret := os.Getenv("WEBHOOK_SECRET")
	if secret == "" {
		log.Warn().Msg("WEBHOOK_SECRET not set, webhook signature validation will be skipped")
	}

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	router := gin.New()

	// Use middleware
	router.Use(LoggerMiddleware())
	router.Use(gin.Recovery())

	// Define routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// GitHub webhook endpoint
	webhookGroup := router.Group("/api/webhooks/github")
	webhookGroup.Use(GithubWebhookMiddleware(secret))
	webhookGroup.POST("", HandleGithubWebhook)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Info().Str("port", port).Msg("Starting server")
	if err := router.Run(":" + port); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
