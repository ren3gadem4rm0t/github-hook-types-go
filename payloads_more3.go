package github

import "encoding/json"

// SecurityAdvisoryPayload represents the webhook payload sent for security_advisory events.
type SecurityAdvisoryPayload struct {
	WebhookPayload
	SecurityAdvisory struct {
		GHSAID      string   `json:"ghsa_id"`
		CVEID       []string `json:"cve_id"`
		Summary     string   `json:"summary"`
		Description string   `json:"description"`
		Severity    string   `json:"severity"`
		Identifiers []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"identifiers"`
		References      json.RawMessage `json:"references"`
		PublishedAt     Timestamp       `json:"published_at"`
		UpdatedAt       Timestamp       `json:"updated_at"`
		Vulnerabilities []struct {
			Package struct {
				Ecosystem string `json:"ecosystem"`
				Name      string `json:"name"`
			} `json:"package"`
			Severity               string `json:"severity"`
			VulnerableVersionRange string `json:"vulnerable_version_range"`
			FirstPatchedVersion    struct {
				Identifier string `json:"identifier"`
			} `json:"first_patched_version"`
		} `json:"vulnerabilities"`
		CVSS struct {
			VectorString string  `json:"vector_string"`
			Score        float64 `json:"score"`
		} `json:"cvss"`
		CWES []struct {
			CWEID string `json:"cwe_id"`
			Name  string `json:"name"`
		} `json:"cwes"`
	} `json:"security_advisory"`
}

// SponsorshipPayload represents the webhook payload sent for sponsorship events.
type SponsorshipPayload struct {
	WebhookPayload
	Sponsorship struct {
		NodeID       string    `json:"node_id"`
		CreatedAt    Timestamp `json:"created_at"`
		Sponsor      User      `json:"sponsor"`
		Sponsorable  User      `json:"sponsorable"`
		PrivacyLevel string    `json:"privacy_level"`
		Tier         struct {
			NodeID                string    `json:"node_id"`
			CreatedAt             Timestamp `json:"created_at"`
			Description           string    `json:"description"`
			MonthlyPriceInCents   int       `json:"monthly_price_in_cents"`
			MonthlyPriceInDollars int       `json:"monthly_price_in_dollars"`
			Name                  string    `json:"name"`
			IsOneTime             bool      `json:"is_one_time"`
			IsCustomAmount        bool      `json:"is_custom_amount"`
		} `json:"tier"`
	} `json:"sponsorship"`
}

// StarPayload represents the webhook payload sent for star events.
type StarPayload struct {
	WebhookPayload
	Starred_at *Timestamp `json:"starred_at"`
}

// StatusPayload represents the webhook payload sent for status events.
type StatusPayload struct {
	WebhookPayload
	ID          int64     `json:"id"`
	SHA         string    `json:"sha"`
	Name        string    `json:"name"`
	TargetURL   string    `json:"target_url"`
	Context     string    `json:"context"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	CommitURL   string    `json:"commit_url"`
	CreatedAt   Timestamp `json:"created_at"`
	UpdatedAt   Timestamp `json:"updated_at"`
	Branches    []struct {
		Name   string `json:"name"`
		Commit struct {
			SHA string `json:"sha"`
			URL string `json:"url"`
		} `json:"commit"`
		Protected bool `json:"protected"`
	} `json:"branches"`
}

// TeamPayload represents the webhook payload sent for team events.
type TeamPayload struct {
	WebhookPayload
	Team    Team `json:"team"`
	Changes struct {
		Description *ChangedFrom `json:"description,omitempty"`
		Name        *ChangedFrom `json:"name,omitempty"`
		Privacy     *ChangedFrom `json:"privacy,omitempty"`
		Repository  *struct {
			Permissions struct {
				From struct {
					Admin bool `json:"admin"`
					Pull  bool `json:"pull"`
					Push  bool `json:"push"`
				} `json:"from"`
			} `json:"permissions"`
		} `json:"repository,omitempty"`
	} `json:"changes,omitempty"`
}

// TeamAddPayload represents the webhook payload sent for team_add events.
type TeamAddPayload struct {
	WebhookPayload
	Team Team `json:"team"`
}

// WatchPayload represents the webhook payload sent for watch events.
type WatchPayload struct {
	WebhookPayload
}

// WorkflowDispatchPayload represents the webhook payload sent for workflow_dispatch events.
type WorkflowDispatchPayload struct {
	WebhookPayload
	Inputs   map[string]string `json:"inputs"`
	Ref      string            `json:"ref"`
	Workflow string            `json:"workflow"`
}

// WorkflowJobPayload represents the webhook payload sent for workflow_job events.
type WorkflowJobPayload struct {
	WebhookPayload
	WorkflowJob struct {
		ID          int64     `json:"id"`
		RunID       int64     `json:"run_id"`
		RunURL      string    `json:"run_url"`
		NodeID      string    `json:"node_id"`
		HeadSHA     string    `json:"head_sha"`
		URL         string    `json:"url"`
		Status      string    `json:"status"`
		Conclusion  string    `json:"conclusion"`
		StartedAt   Timestamp `json:"started_at"`
		CompletedAt Timestamp `json:"completed_at"`
		Name        string    `json:"name"`
		Steps       []struct {
			Name        string    `json:"name"`
			Status      string    `json:"status"`
			Conclusion  string    `json:"conclusion"`
			Number      int       `json:"number"`
			StartedAt   Timestamp `json:"started_at"`
			CompletedAt Timestamp `json:"completed_at"`
		} `json:"steps"`
		CheckRunURL     string   `json:"check_run_url"`
		Labels          []string `json:"labels"`
		RunnerID        int64    `json:"runner_id"`
		RunnerName      string   `json:"runner_name"`
		RunnerGroupID   int64    `json:"runner_group_id"`
		RunnerGroupName string   `json:"runner_group_name"`
	} `json:"workflow_job"`
}

// WorkflowRunPayload represents the webhook payload sent for workflow_run events.
type WorkflowRunPayload struct {
	WebhookPayload
	WorkflowRun struct {
		ID                 int64         `json:"id"`
		NodeID             string        `json:"node_id"`
		Name               string        `json:"name"`
		HeadBranch         string        `json:"head_branch"`
		HeadSHA            string        `json:"head_sha"`
		Path               string        `json:"path"`
		RunNumber          int           `json:"run_number"`
		Event              string        `json:"event"`
		Status             string        `json:"status"`
		Conclusion         string        `json:"conclusion"`
		WorkflowID         int64         `json:"workflow_id"`
		CheckSuiteID       int64         `json:"check_suite_id"`
		CheckSuiteNodeID   string        `json:"check_suite_node_id"`
		URL                string        `json:"url"`
		HTMLURL            string        `json:"html_url"`
		PullRequests       []PullRequest `json:"pull_requests"`
		CreatedAt          Timestamp     `json:"created_at"`
		UpdatedAt          Timestamp     `json:"updated_at"`
		RunStartedAt       Timestamp     `json:"run_started_at"`
		JobsURL            string        `json:"jobs_url"`
		LogsURL            string        `json:"logs_url"`
		CheckSuiteURL      string        `json:"check_suite_url"`
		ArtifactsURL       string        `json:"artifacts_url"`
		CancelURL          string        `json:"cancel_url"`
		RerunURL           string        `json:"rerun_url"`
		PreviousAttemptURL string        `json:"previous_attempt_url"`
		WorkflowURL        string        `json:"workflow_url"`
		HeadCommit         Commit        `json:"head_commit"`
		Repository         Repository    `json:"repository"`
		HeadRepository     Repository    `json:"head_repository"`
	} `json:"workflow_run"`
}

// RegistryPackagePayload represents the webhook payload sent for registry_package events.
type RegistryPackagePayload struct {
	WebhookPayload
	RegistryPackage struct {
		ID             int64     `json:"id"`
		Name           string    `json:"name"`
		PackageType    string    `json:"package_type"`
		HTMLURL        string    `json:"html_url"`
		CreatedAt      Timestamp `json:"created_at"`
		UpdatedAt      Timestamp `json:"updated_at"`
		Owner          User      `json:"owner"`
		PackageVersion struct {
			ID              int64             `json:"id"`
			Version         string            `json:"version"`
			Summary         string            `json:"summary"`
			Description     string            `json:"description"`
			CreatedAt       Timestamp         `json:"created_at"`
			UpdatedAt       Timestamp         `json:"updated_at"`
			HTMLURL         string            `json:"html_url"`
			PackageFiles    []json.RawMessage `json:"package_files,omitempty"`
			ReleaseMetadata json.RawMessage   `json:"metadata,omitempty"`
		} `json:"package_version"`
	} `json:"registry_package"`
}
