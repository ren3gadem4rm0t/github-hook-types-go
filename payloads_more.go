package github

// CommitCommentPayload represents the webhook payload sent for commit_comment events.
type CommitCommentPayload struct {
	WebhookPayload
	Comment struct {
		ID        int64     `json:"id"`
		NodeID    string    `json:"node_id"`
		URL       string    `json:"url"`
		HTMLURL   string    `json:"html_url"`
		User      User      `json:"user"`
		Position  *int      `json:"position"`
		Line      *int      `json:"line"`
		Path      *string   `json:"path"`
		CommitID  string    `json:"commit_id"`
		CreatedAt Timestamp `json:"created_at"`
		UpdatedAt Timestamp `json:"updated_at"`
		Body      string    `json:"body"`
	} `json:"comment"`
}

// ContentReferencePayload represents the webhook payload sent for content_reference events.
type ContentReferencePayload struct {
	WebhookPayload
	ContentReference struct {
		ID        int64  `json:"id"`
		NodeID    string `json:"node_id"`
		Reference string `json:"reference"`
	} `json:"content_reference"`
}

// CreatePayload represents the webhook payload sent for create events.
type CreatePayload struct {
	WebhookPayload
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
}

// DeletePayload represents the webhook payload sent for delete events.
type DeletePayload struct {
	WebhookPayload
	Ref        string `json:"ref"`
	RefType    string `json:"ref_type"`
	PusherType string `json:"pusher_type"`
}

// DeployKeyPayload represents the webhook payload sent for deploy_key events.
type DeployKeyPayload struct {
	WebhookPayload
	Key struct {
		ID        int64     `json:"id"`
		Key       string    `json:"key"`
		URL       string    `json:"url"`
		Title     string    `json:"title"`
		CreatedAt Timestamp `json:"created_at"`
		Verified  bool      `json:"verified"`
		ReadOnly  bool      `json:"read_only"`
	} `json:"key"`
}

// DeploymentPayload represents the webhook payload sent for deployment events.
type DeploymentPayload struct {
	WebhookPayload
	Deployment struct {
		URL                   string    `json:"url"`
		ID                    int64     `json:"id"`
		NodeID                string    `json:"node_id"`
		SHA                   string    `json:"sha"`
		Ref                   string    `json:"ref"`
		Task                  string    `json:"task"`
		Payload               string    `json:"payload"`
		Environment           string    `json:"environment"`
		Description           string    `json:"description"`
		Creator               User      `json:"creator"`
		CreatedAt             Timestamp `json:"created_at"`
		UpdatedAt             Timestamp `json:"updated_at"`
		StatusesURL           string    `json:"statuses_url"`
		RepositoryURL         string    `json:"repository_url"`
		TransientEnvironment  bool      `json:"transient_environment"`
		ProductionEnvironment bool      `json:"production_environment"`
	} `json:"deployment"`
}

// DeploymentStatusPayload represents the webhook payload sent for deployment_status events.
type DeploymentStatusPayload struct {
	WebhookPayload
	DeploymentStatus struct {
		URL            string    `json:"url"`
		ID             int64     `json:"id"`
		NodeID         string    `json:"node_id"`
		State          string    `json:"state"`
		Creator        User      `json:"creator"`
		Description    string    `json:"description"`
		Environment    string    `json:"environment"`
		TargetURL      string    `json:"target_url"`
		CreatedAt      Timestamp `json:"created_at"`
		UpdatedAt      Timestamp `json:"updated_at"`
		DeploymentURL  string    `json:"deployment_url"`
		RepositoryURL  string    `json:"repository_url"`
		LogURL         string    `json:"log_url"`
		EnvironmentURL string    `json:"environment_url"`
	} `json:"deployment_status"`
	Deployment struct {
		URL                   string    `json:"url"`
		ID                    int64     `json:"id"`
		NodeID                string    `json:"node_id"`
		SHA                   string    `json:"sha"`
		Ref                   string    `json:"ref"`
		Task                  string    `json:"task"`
		Payload               string    `json:"payload"`
		Environment           string    `json:"environment"`
		Description           string    `json:"description"`
		Creator               User      `json:"creator"`
		CreatedAt             Timestamp `json:"created_at"`
		UpdatedAt             Timestamp `json:"updated_at"`
		StatusesURL           string    `json:"statuses_url"`
		RepositoryURL         string    `json:"repository_url"`
		TransientEnvironment  bool      `json:"transient_environment"`
		ProductionEnvironment bool      `json:"production_environment"`
	} `json:"deployment"`
}

// DiscussionPayload represents the webhook payload sent for discussion events.
type DiscussionPayload struct {
	WebhookPayload
	Discussion struct {
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		Number            int       `json:"number"`
		Title             string    `json:"title"`
		User              User      `json:"user"`
		State             string    `json:"state"`
		Locked            bool      `json:"locked"`
		Comments          int       `json:"comments"`
		CreatedAt         Timestamp `json:"created_at"`
		UpdatedAt         Timestamp `json:"updated_at"`
		AuthorAssociation string    `json:"author_association"`
		ActiveLockReason  *string   `json:"active_lock_reason"`
		Body              string    `json:"body"`
		TimelineURL       string    `json:"timeline_url"`
		RepositoryURL     string    `json:"repository_url"`
		Category          struct {
			ID           int64     `json:"id"`
			NodeID       string    `json:"node_id"`
			RepositoryID int64     `json:"repository_id"`
			Emoji        string    `json:"emoji"`
			Name         string    `json:"name"`
			Description  string    `json:"description"`
			CreatedAt    Timestamp `json:"created_at"`
			UpdatedAt    Timestamp `json:"updated_at"`
			IsAnswerable bool      `json:"is_answerable"`
		} `json:"category"`
	} `json:"discussion"`
}

// DiscussionCommentPayload represents the webhook payload sent for discussion_comment events.
type DiscussionCommentPayload struct {
	WebhookPayload
	Discussion struct {
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		Number            int       `json:"number"`
		Title             string    `json:"title"`
		User              User      `json:"user"`
		State             string    `json:"state"`
		Locked            bool      `json:"locked"`
		Comments          int       `json:"comments"`
		CreatedAt         Timestamp `json:"created_at"`
		UpdatedAt         Timestamp `json:"updated_at"`
		AuthorAssociation string    `json:"author_association"`
		ActiveLockReason  *string   `json:"active_lock_reason"`
		Body              string    `json:"body"`
		TimelineURL       string    `json:"timeline_url"`
		RepositoryURL     string    `json:"repository_url"`
		Category          struct {
			ID           int64     `json:"id"`
			NodeID       string    `json:"node_id"`
			RepositoryID int64     `json:"repository_id"`
			Emoji        string    `json:"emoji"`
			Name         string    `json:"name"`
			Description  string    `json:"description"`
			CreatedAt    Timestamp `json:"created_at"`
			UpdatedAt    Timestamp `json:"updated_at"`
			IsAnswerable bool      `json:"is_answerable"`
		} `json:"category"`
	} `json:"discussion"`
	Comment struct {
		ID                int64     `json:"id"`
		NodeID            string    `json:"node_id"`
		DiscussionID      int64     `json:"discussion_id"`
		User              User      `json:"user"`
		CreatedAt         Timestamp `json:"created_at"`
		UpdatedAt         Timestamp `json:"updated_at"`
		AuthorAssociation string    `json:"author_association"`
		Body              string    `json:"body"`
		HTMLURL           string    `json:"html_url"`
		ParentID          *int64    `json:"parent_id"`
		ChildCommentCount int       `json:"child_comment_count"`
		RepositoryURL     string    `json:"repository_url"`
	} `json:"comment"`
}

// ForkPayload represents the webhook payload sent for fork events.
type ForkPayload struct {
	WebhookPayload
	Forkee Repository `json:"forkee"`
}

// GitHubAppAuthorizationPayload represents the webhook payload sent for github_app_authorization events.
type GitHubAppAuthorizationPayload struct {
	WebhookPayload
}

// GollumPayload represents the webhook payload sent for gollum events.
type GollumPayload struct {
	WebhookPayload
	Pages []struct {
		PageName string `json:"page_name"`
		Title    string `json:"title"`
		Summary  string `json:"summary,omitempty"`
		Action   string `json:"action"`
		SHA      string `json:"sha"`
		HTMLURL  string `json:"html_url"`
	} `json:"pages"`
}

// InstallationPayload represents the webhook payload sent for installation events.
type InstallationPayload struct {
	WebhookPayload
	Installation struct {
		ID                  int64     `json:"id"`
		NodeID              string    `json:"node_id"`
		AppID               int64     `json:"app_id"`
		AppSlug             string    `json:"app_slug"`
		TargetID            int64     `json:"target_id"`
		TargetType          string    `json:"target_type"`
		RepositorySelection string    `json:"repository_selection"`
		Account             User      `json:"account"`
		AccessTokensURL     string    `json:"access_tokens_url"`
		RepositoriesURL     string    `json:"repositories_url"`
		HTMLURL             string    `json:"html_url"`
		CreatedAt           Timestamp `json:"created_at"`
		UpdatedAt           Timestamp `json:"updated_at"`
		Events              []string  `json:"events"`
		Permissions         struct {
			Issues       string `json:"issues"`
			Contents     string `json:"contents"`
			Metadata     string `json:"metadata"`
			SingleFile   string `json:"single_file"`
			PullRequests string `json:"pull_requests"`
		} `json:"permissions"`
	} `json:"installation"`
}

// InstallationRepositoriesPayload represents the webhook payload sent for installation_repositories events.
type InstallationRepositoriesPayload struct {
	WebhookPayload
	Installation struct {
		ID                  int64     `json:"id"`
		NodeID              string    `json:"node_id"`
		AppID               int64     `json:"app_id"`
		AppSlug             string    `json:"app_slug"`
		TargetID            int64     `json:"target_id"`
		TargetType          string    `json:"target_type"`
		RepositorySelection string    `json:"repository_selection"`
		Account             User      `json:"account"`
		AccessTokensURL     string    `json:"access_tokens_url"`
		RepositoriesURL     string    `json:"repositories_url"`
		HTMLURL             string    `json:"html_url"`
		CreatedAt           Timestamp `json:"created_at"`
		UpdatedAt           Timestamp `json:"updated_at"`
		Events              []string  `json:"events"`
		Permissions         struct {
			Issues       string `json:"issues"`
			Contents     string `json:"contents"`
			Metadata     string `json:"metadata"`
			SingleFile   string `json:"single_file"`
			PullRequests string `json:"pull_requests"`
		} `json:"permissions"`
	} `json:"installation"`
	RepositoriesAdded   []Repository `json:"repositories_added"`
	RepositoriesRemoved []Repository `json:"repositories_removed"`
	RepositorySelection string       `json:"repository_selection"`
}

// LabelPayload represents the webhook payload sent for label events.
type LabelPayload struct {
	WebhookPayload
	Label   Label `json:"label"`
	Changes struct {
		Name        *ChangedFrom `json:"name,omitempty"`
		Color       *ChangedFrom `json:"color,omitempty"`
		Description *ChangedFrom `json:"description,omitempty"`
	} `json:"changes,omitempty"`
}

// MarketplacePurchasePayload represents the webhook payload sent for marketplace_purchase events.
type MarketplacePurchasePayload struct {
	WebhookPayload
	MarketplacePurchase struct {
		Account         User       `json:"account"`
		BillingCycle    string     `json:"billing_cycle"`
		NextBillingDate *Timestamp `json:"next_billing_date"`
		UnitCount       int        `json:"unit_count"`
		Plan            struct {
			ID          int64    `json:"id"`
			Name        string   `json:"name"`
			Description string   `json:"description"`
			Monthly     bool     `json:"monthly_price_in_cents"`
			Yearly      bool     `json:"yearly_price_in_cents"`
			PriceModel  string   `json:"price_model"`
			UnitName    string   `json:"unit_name"`
			Bullets     []string `json:"bullets"`
		} `json:"plan"`
		OnFreeTrial     bool       `json:"on_free_trial"`
		FreeTrialEndsOn *Timestamp `json:"free_trial_ends_on"`
		PendingChange   *struct {
			Plan struct {
				ID          int64    `json:"id"`
				Name        string   `json:"name"`
				Description string   `json:"description"`
				Monthly     bool     `json:"monthly_price_in_cents"`
				Yearly      bool     `json:"yearly_price_in_cents"`
				PriceModel  string   `json:"price_model"`
				UnitName    string   `json:"unit_name"`
				Bullets     []string `json:"bullets"`
			} `json:"plan"`
			EffectiveDate *Timestamp `json:"effective_date"`
		} `json:"pending_change"`
	} `json:"marketplace_purchase"`
}

// MemberPayload represents the webhook payload sent for member events.
type MemberPayload struct {
	WebhookPayload
	Member  User `json:"member"`
	Changes struct {
		Permission *ChangedFrom `json:"permission,omitempty"`
		Role       *ChangedFrom `json:"role,omitempty"`
	} `json:"changes,omitempty"`
}

// MembershipPayload represents the webhook payload sent for membership events.
type MembershipPayload struct {
	WebhookPayload
	Scope  string `json:"scope"`
	Member User   `json:"member"`
	Team   Team   `json:"team"`
}

// MetaPayload represents the webhook payload sent for meta events.
type MetaPayload struct {
	WebhookPayload
	HookID int64 `json:"hook_id"`
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
