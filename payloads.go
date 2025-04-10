package github

// PingPayload represents the webhook payload sent for a ping event.
type PingPayload struct {
	WebhookPayload
	Zen    string `json:"zen"`
	HookID int64  `json:"hook_id"`
	Hook   struct {
		Type   string   `json:"type"`
		ID     int64    `json:"id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSsl string `json:"insecure_ssl"`
			URL         string `json:"url"`
		} `json:"config"`
	} `json:"hook"`
}

// PushPayload represents the webhook payload sent for a push event.
type PushPayload struct {
	WebhookPayload
	Ref          string       `json:"ref"`
	Before       string       `json:"before"`
	After        string       `json:"after"`
	Created      bool         `json:"created"`
	Deleted      bool         `json:"deleted"`
	Forced       bool         `json:"forced"`
	BaseRef      *string      `json:"base_ref"`
	Compare      string       `json:"compare"`
	Commits      []Commit     `json:"commits"`
	HeadCommit   *Commit      `json:"head_commit"`
	PusherPerson PusherPerson `json:"pusher"`
}

// PusherPerson represents a user who pushed a commit.
type PusherPerson struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Commit represents a git commit in a webhook payload.
type Commit struct {
	ID        string       `json:"id"`
	TreeID    string       `json:"tree_id"`
	Distinct  bool         `json:"distinct"`
	Message   string       `json:"message"`
	Timestamp Timestamp    `json:"timestamp"`
	URL       string       `json:"url"`
	Author    CommitAuthor `json:"author"`
	Committer CommitAuthor `json:"committer"`
	Added     []string     `json:"added"`
	Removed   []string     `json:"removed"`
	Modified  []string     `json:"modified"`
}

// CommitAuthor represents the author or committer of a commit.
type CommitAuthor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username,omitempty"`
}

// IssuesPayload represents the webhook payload sent for issues events.
type IssuesPayload struct {
	WebhookPayload
	Issue   Issue `json:"issue"`
	Changes struct {
		Title    *ChangedFrom `json:"title,omitempty"`
		Body     *ChangedFrom `json:"body,omitempty"`
		Labels   *ChangedFrom `json:"labels,omitempty"`
		Assignee *ChangedFrom `json:"assignee,omitempty"`
	} `json:"changes,omitempty"`
}

// Issue represents an issue in a GitHub repository.
type Issue struct {
	ID               int64      `json:"id"`
	NodeID           string     `json:"node_id"`
	URL              string     `json:"url"`
	RepositoryURL    string     `json:"repository_url"`
	LabelsURL        string     `json:"labels_url"`
	CommentsURL      string     `json:"comments_url"`
	EventsURL        string     `json:"events_url"`
	HTMLURL          string     `json:"html_url"`
	Number           int        `json:"number"`
	State            string     `json:"state"`
	Title            string     `json:"title"`
	Body             string     `json:"body"`
	User             User       `json:"user"`
	Labels           []Label    `json:"labels"`
	Assignee         *User      `json:"assignee"`
	Assignees        []User     `json:"assignees"`
	Milestone        *Milestone `json:"milestone"`
	Locked           bool       `json:"locked"`
	ActiveLockReason string     `json:"active_lock_reason,omitempty"`
	Comments         int        `json:"comments"`
	PullRequest      *struct {
		URL      string `json:"url"`
		HTMLURL  string `json:"html_url"`
		DiffURL  string `json:"diff_url"`
		PatchURL string `json:"patch_url"`
	} `json:"pull_request,omitempty"`
	ClosedAt          *Timestamp `json:"closed_at"`
	CreatedAt         Timestamp  `json:"created_at"`
	UpdatedAt         Timestamp  `json:"updated_at"`
	AuthorAssociation string     `json:"author_association"`
}

// Label represents a label on an issue or pull request.
type Label struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

// Milestone represents a GitHub milestone.
type Milestone struct {
	ID           int64      `json:"id"`
	NodeID       string     `json:"node_id"`
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	Number       int        `json:"number"`
	State        string     `json:"state"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      User       `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	CreatedAt    Timestamp  `json:"created_at"`
	UpdatedAt    Timestamp  `json:"updated_at"`
	ClosedAt     *Timestamp `json:"closed_at"`
	DueOn        *Timestamp `json:"due_on"`
}

// ChangedFrom represents a changed value in a webhook payload.
type ChangedFrom struct {
	From interface{} `json:"from"`
}

// IssueCommentPayload represents the webhook payload sent for issue_comment events.
type IssueCommentPayload struct {
	WebhookPayload
	Issue   Issue   `json:"issue"`
	Comment Comment `json:"comment"`
	Changes struct {
		Body *ChangedFrom `json:"body,omitempty"`
	} `json:"changes,omitempty"`
}

// Comment represents a comment on an issue or pull request.
type Comment struct {
	ID                    int64     `json:"id"`
	NodeID                string    `json:"node_id"`
	URL                   string    `json:"url"`
	HTMLURL               string    `json:"html_url"`
	Body                  string    `json:"body"`
	User                  User      `json:"user"`
	CreatedAt             Timestamp `json:"created_at"`
	UpdatedAt             Timestamp `json:"updated_at"`
	IssueURL              string    `json:"issue_url,omitempty"`
	AuthorAssociation     string    `json:"author_association"`
	PerformedViaGitHubApp *struct {
		ID     int64  `json:"id"`
		NodeID string `json:"node_id"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
	} `json:"performed_via_github_app,omitempty"`
}

// PullRequestPayload represents the webhook payload sent for pull_request events.
type PullRequestPayload struct {
	WebhookPayload
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Changes     struct {
		Title *ChangedFrom `json:"title,omitempty"`
		Body  *ChangedFrom `json:"body,omitempty"`
		Base  *struct {
			Ref struct {
				From string `json:"from"`
			} `json:"ref"`
			SHA struct {
				From string `json:"from"`
			} `json:"sha"`
		} `json:"base,omitempty"`
	} `json:"changes,omitempty"`
}

// PullRequest represents a GitHub pull request.
type PullRequest struct {
	ID                  int64             `json:"id"`
	NodeID              string            `json:"node_id"`
	URL                 string            `json:"url"`
	HTMLURL             string            `json:"html_url"`
	DiffURL             string            `json:"diff_url"`
	PatchURL            string            `json:"patch_url"`
	IssueURL            string            `json:"issue_url"`
	Number              int               `json:"number"`
	State               string            `json:"state"`
	Locked              bool              `json:"locked"`
	Title               string            `json:"title"`
	Body                string            `json:"body"`
	CreatedAt           Timestamp         `json:"created_at"`
	UpdatedAt           Timestamp         `json:"updated_at"`
	ClosedAt            *Timestamp        `json:"closed_at"`
	MergedAt            *Timestamp        `json:"merged_at"`
	MergeCommitSHA      *string           `json:"merge_commit_sha"`
	Assignee            *User             `json:"assignee"`
	Assignees           []User            `json:"assignees"`
	RequestedReviewers  []User            `json:"requested_reviewers"`
	RequestedTeams      []Team            `json:"requested_teams"`
	Labels              []Label           `json:"labels"`
	Milestone           *Milestone        `json:"milestone"`
	Draft               bool              `json:"draft"`
	User                User              `json:"user"`
	Base                PullRequestBranch `json:"base"`
	Head                PullRequestBranch `json:"head"`
	AuthorAssociation   string            `json:"author_association"`
	Merged              bool              `json:"merged"`
	Mergeable           *bool             `json:"mergeable"`
	Rebaseable          *bool             `json:"rebaseable"`
	MergeableState      string            `json:"mergeable_state"`
	MergedBy            *User             `json:"merged_by"`
	Comments            int               `json:"comments"`
	ReviewComments      int               `json:"review_comments"`
	MaintainerCanModify bool              `json:"maintainer_can_modify"`
	Commits             int               `json:"commits"`
	Additions           int               `json:"additions"`
	Deletions           int               `json:"deletions"`
	ChangedFiles        int               `json:"changed_files"`
}

// PullRequestBranch represents a branch in a pull request.
type PullRequestBranch struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	SHA   string     `json:"sha"`
	User  User       `json:"user"`
	Repo  Repository `json:"repo"`
}

// Team represents a GitHub team.
type Team struct {
	ID              int64  `json:"id"`
	NodeID          string `json:"node_id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	Privacy         string `json:"privacy"`
	URL             string `json:"url"`
	HTMLURL         string `json:"html_url"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
	Permission      string `json:"permission"`
}

// CheckRunPayload represents the webhook payload sent for check_run events.
type CheckRunPayload struct {
	WebhookPayload
	CheckRun CheckRun `json:"check_run"`
}

// CheckRun represents a GitHub check run.
type CheckRun struct {
	ID          int64      `json:"id"`
	NodeID      string     `json:"node_id"`
	HeadSHA     string     `json:"head_sha"`
	ExternalID  string     `json:"external_id"`
	URL         string     `json:"url"`
	HTMLURL     string     `json:"html_url"`
	DetailsURL  string     `json:"details_url"`
	Status      string     `json:"status"`
	Conclusion  *string    `json:"conclusion"`
	StartedAt   Timestamp  `json:"started_at"`
	CompletedAt *Timestamp `json:"completed_at"`
	Output      struct {
		Title            string `json:"title"`
		Summary          string `json:"summary"`
		Text             string `json:"text"`
		AnnotationsCount int    `json:"annotations_count"`
		AnnotationsURL   string `json:"annotations_url"`
	} `json:"output"`
	Name       string `json:"name"`
	CheckSuite struct {
		ID           int64         `json:"id"`
		NodeID       string        `json:"node_id"`
		HeadBranch   string        `json:"head_branch"`
		HeadSHA      string        `json:"head_sha"`
		Status       string        `json:"status"`
		Conclusion   *string       `json:"conclusion"`
		URL          string        `json:"url"`
		Before       string        `json:"before"`
		After        string        `json:"after"`
		PullRequests []PullRequest `json:"pull_requests"`
		App          struct {
			ID          int64     `json:"id"`
			NodeID      string    `json:"node_id"`
			Owner       User      `json:"owner"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			ExternalURL string    `json:"external_url"`
			HTMLURL     string    `json:"html_url"`
			CreatedAt   Timestamp `json:"created_at"`
			UpdatedAt   Timestamp `json:"updated_at"`
		} `json:"app"`
		CreatedAt Timestamp `json:"created_at"`
		UpdatedAt Timestamp `json:"updated_at"`
	} `json:"check_suite"`
	App struct {
		ID          int64     `json:"id"`
		NodeID      string    `json:"node_id"`
		Owner       User      `json:"owner"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ExternalURL string    `json:"external_url"`
		HTMLURL     string    `json:"html_url"`
		CreatedAt   Timestamp `json:"created_at"`
		UpdatedAt   Timestamp `json:"updated_at"`
	} `json:"app"`
	PullRequests []PullRequest `json:"pull_requests"`
}

// CheckSuitePayload represents the webhook payload sent for check_suite events.
type CheckSuitePayload struct {
	WebhookPayload
	CheckSuite struct {
		ID           int64         `json:"id"`
		NodeID       string        `json:"node_id"`
		HeadBranch   string        `json:"head_branch"`
		HeadSHA      string        `json:"head_sha"`
		Status       string        `json:"status"`
		Conclusion   *string       `json:"conclusion"`
		URL          string        `json:"url"`
		Before       string        `json:"before"`
		After        string        `json:"after"`
		PullRequests []PullRequest `json:"pull_requests"`
		App          struct {
			ID          int64     `json:"id"`
			NodeID      string    `json:"node_id"`
			Owner       User      `json:"owner"`
			Name        string    `json:"name"`
			Description string    `json:"description"`
			ExternalURL string    `json:"external_url"`
			HTMLURL     string    `json:"html_url"`
			CreatedAt   Timestamp `json:"created_at"`
			UpdatedAt   Timestamp `json:"updated_at"`
		} `json:"app"`
		CreatedAt Timestamp `json:"created_at"`
		UpdatedAt Timestamp `json:"updated_at"`
	} `json:"check_suite"`
}
