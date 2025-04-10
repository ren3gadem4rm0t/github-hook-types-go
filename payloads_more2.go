package github

import (
	"encoding/json"
	"time"
)

// MilestonePayload represents the webhook payload sent for milestone events.
type MilestonePayload struct {
	WebhookPayload
	Milestone Milestone `json:"milestone"`
	Changes   struct {
		Title       *ChangedFrom `json:"title,omitempty"`
		Description *ChangedFrom `json:"description,omitempty"`
		DueOn       *ChangedFrom `json:"due_on,omitempty"`
	} `json:"changes,omitempty"`
}

// OrganizationPayload represents the webhook payload sent for organization events.
type OrganizationPayload struct {
	WebhookPayload
	Invitation struct {
		ID     int64  `json:"id"`
		NodeID string `json:"node_id"`
		Login  string `json:"login"`
		Email  string `json:"email"`
		Role   string `json:"role"`
	} `json:"invitation,omitempty"`
	Membership struct {
		URL    string `json:"url"`
		State  string `json:"state"`
		Role   string `json:"role"`
		OrgURL string `json:"organization_url"`
		User   User   `json:"user"`
	} `json:"membership,omitempty"`
	Changes struct {
		DefaultRepositoryPermission  *ChangedFrom `json:"default_repository_permission,omitempty"`
		MembersCanCreateRepositories *ChangedFrom `json:"members_can_create_repositories,omitempty"`
	} `json:"changes,omitempty"`
}

// OrgBlockPayload represents the webhook payload sent for org_block events.
type OrgBlockPayload struct {
	WebhookPayload
	BlockedUser User `json:"blocked_user"`
}

// PackagePayload represents the webhook payload sent for package events.
type PackagePayload struct {
	WebhookPayload
	Package struct {
		ID             int64     `json:"id"`
		Name           string    `json:"name"`
		PackageType    string    `json:"package_type"`
		HTMLURL        string    `json:"html_url"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Owner          User      `json:"owner"`
		PackageVersion struct {
			ID           int64     `json:"id"`
			Version      string    `json:"version"`
			Summary      string    `json:"summary"`
			Body         string    `json:"body"`
			BodyHTML     string    `json:"body_html"`
			HTMLURL      string    `json:"html_url"`
			CreatedAt    time.Time `json:"created_at"`
			UpdatedAt    time.Time `json:"updated_at"`
			PackageFiles []struct {
				DownloadURL string    `json:"download_url"`
				ID          int64     `json:"id"`
				Name        string    `json:"name"`
				SHA256      string    `json:"sha256"`
				SHA1        string    `json:"sha1"`
				MD5         string    `json:"md5"`
				Size        int       `json:"size"`
				CreatedAt   time.Time `json:"created_at"`
				UpdatedAt   time.Time `json:"updated_at"`
			} `json:"package_files"`
		} `json:"package_version"`
	} `json:"package"`
}

// PageBuildPayload represents the webhook payload sent for page_build events.
type PageBuildPayload struct {
	WebhookPayload
	ID    int64 `json:"id"`
	Build struct {
		URL    string `json:"url"`
		Status string `json:"status"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
		Pusher    User      `json:"pusher"`
		Commit    string    `json:"commit"`
		Duration  int       `json:"duration"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"build"`
}

// ProjectPayload represents the webhook payload sent for project events.
type ProjectPayload struct {
	WebhookPayload
	Project struct {
		ID         int64     `json:"id"`
		NodeID     string    `json:"node_id"`
		Name       string    `json:"name"`
		Body       string    `json:"body"`
		Number     int       `json:"number"`
		State      string    `json:"state"`
		HTMLURL    string    `json:"html_url"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		Creator    User      `json:"creator"`
		ColumnsURL string    `json:"columns_url"`
	} `json:"project"`
	Changes struct {
		Name *ChangedFrom `json:"name,omitempty"`
		Body *ChangedFrom `json:"body,omitempty"`
	} `json:"changes,omitempty"`
}

// ProjectCardPayload represents the webhook payload sent for project_card events.
type ProjectCardPayload struct {
	WebhookPayload
	ProjectCard struct {
		ID         int64     `json:"id"`
		NodeID     string    `json:"node_id"`
		Note       string    `json:"note"`
		Creator    User      `json:"creator"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		ProjectURL string    `json:"project_url"`
		ColumnURL  string    `json:"column_url"`
		ContentURL string    `json:"content_url,omitempty"`
		ProjectID  int64     `json:"project_id"`
		ColumnID   int64     `json:"column_id"`
		Archived   bool      `json:"archived"`
	} `json:"project_card"`
	Changes struct {
		Note *ChangedFrom `json:"note,omitempty"`
	} `json:"changes,omitempty"`
}

// ProjectColumnPayload represents the webhook payload sent for project_column events.
type ProjectColumnPayload struct {
	WebhookPayload
	ProjectColumn struct {
		ID         int64     `json:"id"`
		NodeID     string    `json:"node_id"`
		Name       string    `json:"name"`
		ProjectURL string    `json:"project_url"`
		CardsURL   string    `json:"cards_url"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"project_column"`
	Changes struct {
		Name *ChangedFrom `json:"name,omitempty"`
	} `json:"changes,omitempty"`
}

// PublicPayload represents the webhook payload sent for public events.
type PublicPayload struct {
	WebhookPayload
}

// PullRequestReviewPayload represents the webhook payload sent for pull_request_review events.
type PullRequestReviewPayload struct {
	WebhookPayload
	Review struct {
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		User              User      `json:"user"`
		Body              string    `json:"body"`
		CommitID          string    `json:"commit_id"`
		HTMLURL           string    `json:"html_url"`
		PullRequestURL    string    `json:"pull_request_url"`
		State             string    `json:"state"`
		AuthorAssociation string    `json:"author_association"`
		SubmittedAt       time.Time `json:"submitted_at"`
	} `json:"review"`
	PullRequest PullRequest `json:"pull_request"`
	Changes     struct {
		Body *ChangedFrom `json:"body,omitempty"`
	} `json:"changes,omitempty"`
}

// PullRequestReviewCommentPayload represents the webhook payload sent for pull_request_review_comment events.
type PullRequestReviewCommentPayload struct {
	WebhookPayload
	Comment struct {
		ID                  int64     `json:"id"`
		NodeID              string    `json:"node_id"`
		Path                string    `json:"path"`
		Position            *int      `json:"position"`
		OriginalPosition    int       `json:"original_position"`
		CommitID            string    `json:"commit_id"`
		OriginalCommitID    string    `json:"original_commit_id"`
		User                User      `json:"user"`
		Body                string    `json:"body"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
		HTMLURL             string    `json:"html_url"`
		PullRequestURL      string    `json:"pull_request_url"`
		AuthorAssociation   string    `json:"author_association"`
		DiffHunk            string    `json:"diff_hunk"`
		PullRequestReviewID int64     `json:"pull_request_review_id"`
		InReplyToID         *int64    `json:"in_reply_to_id"`
	} `json:"comment"`
	PullRequest PullRequest `json:"pull_request"`
	Changes     struct {
		Body *ChangedFrom `json:"body,omitempty"`
	} `json:"changes,omitempty"`
}

// ReleasePayload represents the webhook payload sent for release events.
type ReleasePayload struct {
	WebhookPayload
	Release struct {
		ID              int64     `json:"id"`
		NodeID          string    `json:"node_id"`
		TagName         string    `json:"tag_name"`
		TargetCommitish string    `json:"target_commitish"`
		Name            string    `json:"name"`
		Draft           bool      `json:"draft"`
		Author          User      `json:"author"`
		Prerelease      bool      `json:"prerelease"`
		CreatedAt       time.Time `json:"created_at"`
		PublishedAt     time.Time `json:"published_at"`
		AssetsURL       string    `json:"assets_url"`
		TarballURL      string    `json:"tarball_url"`
		ZipballURL      string    `json:"zipball_url"`
		HTMLURL         string    `json:"html_url"`
		Body            string    `json:"body"`
		Assets          []struct {
			URL                string    `json:"url"`
			BrowserDownloadURL string    `json:"browser_download_url"`
			ID                 int64     `json:"id"`
			NodeID             string    `json:"node_id"`
			Name               string    `json:"name"`
			Label              string    `json:"label"`
			State              string    `json:"state"`
			ContentType        string    `json:"content_type"`
			Size               int       `json:"size"`
			DownloadCount      int       `json:"download_count"`
			CreatedAt          time.Time `json:"created_at"`
			UpdatedAt          time.Time `json:"updated_at"`
			Uploader           User      `json:"uploader"`
		} `json:"assets"`
	} `json:"release"`
}

// RepositoryPayload represents the webhook payload sent for repository events.
type RepositoryPayload struct {
	WebhookPayload
	Changes struct {
		DefaultBranch *ChangedFrom `json:"default_branch,omitempty"`
		Description   *ChangedFrom `json:"description,omitempty"`
		Homepage      *ChangedFrom `json:"homepage,omitempty"`
	} `json:"changes,omitempty"`
}

// RepositoryDispatchPayload represents the webhook payload sent for repository_dispatch events.
type RepositoryDispatchPayload struct {
	WebhookPayload
	Branch        string          `json:"branch"`
	ClientPayload json.RawMessage `json:"client_payload"`
}

// RepositoryImportPayload represents the webhook payload sent for repository_import events.
type RepositoryImportPayload struct {
	WebhookPayload
	Status string `json:"status"`
}

// RepositoryVulnerabilityAlertPayload represents the webhook payload sent for repository_vulnerability_alert events.
type RepositoryVulnerabilityAlertPayload struct {
	WebhookPayload
	Alert struct {
		ID                  int64      `json:"id"`
		AffectedRange       string     `json:"affected_range"`
		AffectedPackageName string     `json:"affected_package_name"`
		ExternalReference   string     `json:"external_reference"`
		ExternalIdentifier  string     `json:"external_identifier"`
		FixedIn             string     `json:"fixed_in"`
		DismissedAt         *time.Time `json:"dismissed_at"`
		DismissedBy         *User      `json:"dismissed_by"`
		DismissedReason     string     `json:"dismissed_reason"`
		GhsaID              string     `json:"ghsa_id"`
		CVEID               string     `json:"cve_id"`
	} `json:"alert"`
}
