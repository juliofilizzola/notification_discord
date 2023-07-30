package domain

type User struct {
	Login        string `json:"login,omitempty"`
	Id           int    `json:"id,omitempty"`
	NodeId       string `json:"node_id,omitempty"`
	AvatarUrl    string `json:"avatar_url,omitempty"`
	Url          string `json:"url,omitempty"`
	HtmlUrl      string `json:"html_url,omitempty"`
	FollowersUrl string `json:"followers_url,omitempty"`
	GravatarId   string `json:"gravatar_id,omitempty"`
	Type         string `json:"type,omitempty"`
}

type Repo struct {
	Id       int    `json:"id,omitempty"`
	NodeId   string `json:"node_id,omitempty"`
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Private  bool   `json:"private,omitempty"`
}

type Repository struct {
	Repo
	Owner User `json:"owner"`
}
type Label struct {
	Id          int    `json:"id,omitempty"`
	NodeId      string `json:"node_id,omitempty"`
	Url         string `json:"url,omitempty"`
	Name        string `json:"name,omitempty"`
	Color       string `json:"color,omitempty"`
	Default     bool   `json:"default,omitempty"`
	Description string `json:"description,omitempty"`
}

type PullRequest struct {
	Url                 string  `json:"url,omitempty"`
	Id                  int     `json:"id,omitempty"`
	User                User    `json:"user"`
	NodeId              string  `json:"node_id,omitempty"`
	HtmlUrl             string  `json:"html_url,omitempty"`
	DiffUrl             string  `json:"diff_url,omitempty"`
	PatchUrl            string  `json:"patch_url,omitempty"`
	IssueUrl            string  `json:"issue_url,omitempty"`
	Number              int     `json:"number,omitempty"`
	State               string  `json:"state,omitempty"`
	Locked              bool    `json:"locked,omitempty"`
	Title               string  `json:"title,omitempty"`
	Body                string  `json:"body,omitempty"`
	CreatedAt           string  `json:"created_at,omitempty"`
	UpdatedAt           string  `json:"updated_at,omitempty"`
	ClosedAt            string  `json:"closed_at,omitempty"`
	MergedAt            string  `json:"merged_at,omitempty"`
	MergeCommitSha      string  `json:"merge_commit_sha,omitempty"`
	Assignee            User    `json:"assignee"`
	Assignees           []User  `json:"assignees,omitempty"`
	Repo                Repo    `json:"repo"`
	RequestedReviewers  []User  `json:"requested_reviewers"`
	Labels              []Label `json:"labels"`
	Comments            int     `json:"comments,omitempty"`
	ReviewComments      int     `json:"review_comments,omitempty"`
	MaintainerCanModify bool    `json:"maintainer_can_modify,omitempty"`
	Commits             int     `json:"commits,omitempty"`
	Additions           int     `json:"additions,omitempty"`
	Deletions           int     `json:"deletions,omitempty"`
	ChangedFiles        int     `json:"changed_files,omitempty"`
	CommitsUrl          string  `json:"commits_url,omitempty"`
	ReviewCommentsUrl   string  `json:"review_comments_url,omitempty"`
	ReviewCommentUrl    string  `json:"review_comment_url,omitempty"`
	CommentsUrl         string  `json:"comments_url,omitempty"`
	StatusesUrl         string  `json:"statuses_url,omitempty"`
}
type Organization struct {
	Id               int    `json:"id,omitempty"`
	Login            string `json:"login,omitempty"`
	ModeId           string `json:"mode_id,omitempty"`
	Url              string `json:"url,omitempty"`
	ReposUrl         string `json:"repos_url,omitempty"`
	EventsUrl        string `json:"events_url,omitempty"`
	HooksUrl         string `json:"hooks_url,omitempty"`
	IssuesUrl        string `json:"issues_url,omitempty"`
	MembersUrl       string `json:"members_url,omitempty"`
	PublicMembersUrl string `json:"public_members_url,omitempty"`
	AvatarUrl        string `json:"avatar_url,omitempty"`
	Description      string `json:"description,omitempty"`
}

type Github struct {
	Number       int          `json:"number,omitempty"`
	Action       string       `json:"action,omitempty"`
	PullRequest  PullRequest  `json:"pull_request"`
	Repository   Repository   `json:"repository"`
	Sender       User         `json:"sender"`
	Organization Organization `json:"organization"`
}
