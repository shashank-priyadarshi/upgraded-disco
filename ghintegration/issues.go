package ghintegration

import (
	"encoding/json"
	"log"
	"server/common"
)

func getIssueData() []string {
	var issueList []IssueRequest
	var issues []string

	// api call to fetch list of all issues
	rawIssueResposne, _ := common.BearerAuthAPICall("https://api.github.com/user/issues?per_page=100&page=1", authToken)
	err := json.Unmarshal(rawIssueResposne, &issueList)
	if err != nil {
		log.Println("Unable to unmarshal raw issue list response: ", err)
	}
	// according to date, increment issue count
	// recording issue count & issue names for corresponding duration
	for _, issue := range issueList {
		issues = append(issues, issue.Title)
	}
	return issues
}
