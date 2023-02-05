package ghintegration

import "time"

type RepoResponse struct {
	Name  string `json:"name"`
	URL   string `json:"html_url"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	Forked     bool   `json:"fork"`
	CommitsURL string `json:"commits_url"`
	PRURL      string `json:"pulls_url"`
}

type CommitResponse struct {
	Commit struct {
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

type PullRequestResponse struct {
	State string `json:"state"`
	User  struct {
		Login string `json:"login"`
	} `json:"user"`
	CreatedAt string `json:"created_at"`
}

type IssueRequest struct {
	URL       string `json:"url"`
	Title     string `json:"title"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	ClosedAt  string `json:"closed_at"`
	Assignee  struct {
		Login string `json:"login"`
	} `json:"assignee"`
}

type GitHubData struct {
	Time         time.Time     `json:"execution_time"`
	StarredRepos Repo          `json:"starredrepos"`
	Issues       []string      `json:"issues"`
	WeekData     []SCMActivity `json:"weekdata"`
}

type Repo struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}

type SCMActivity struct {
	PR     int    `json:"pr"`
	LOC    int    `json:"loc"`
	Date   string `json:"date"`
	Commit int    `json:"commit"`
}

type Issue struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
