//Package issue
/*
@Title: bind.go
@Description
@Author: kkw 2023/6/8 16:29
*/
package issue

import "time"

type WebHookReq struct {
	Action string `json:"action"`
	Issue  struct {
		HtmlUrl               string        `json:"html_url"`
		Labels                []LabelReq    `json:"labels"`
		State                 string        `json:"state"`
		Locked                bool          `json:"locked"`
		Assignee              interface{}   `json:"assignee"`
		Assignees             []interface{} `json:"assignees"`
		Milestone             interface{}   `json:"milestone"`
		Comments              int           `json:"comments"`
		CreatedAt             time.Time     `json:"created_at"`
		UpdatedAt             time.Time     `json:"updated_at"`
		ClosedAt              time.Time     `json:"closed_at"`
		AuthorAssociation     string        `json:"author_association"`
		ActiveLockReason      interface{}   `json:"active_lock_reason"`
		Body                  string        `json:"body"`
		TimelineURL           string        `json:"timeline_url"`
		PerformedViaGithubApp interface{}   `json:"performed_via_github_app"`
		StateReason           string        `json:"state_reason"`
	} `json:"issue"`
}

type LabelReq struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}
