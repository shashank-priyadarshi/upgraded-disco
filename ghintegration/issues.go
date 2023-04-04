package ghintegration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/common"

	logger "github.com/rs/zerolog/log"
)

func getIssueData() []string {
	var rawData []byte
	page, statusCode, params, issues, issueList := 0, 200, queryParams, []string{}, []IssueRequest{}

	for statusCode == http.StatusOK {
		// fetching issues for user
		page++
		tempIssueList := []IssueRequest{}
		params[1] = fmt.Sprintf("page %v", page)
		rawData, statusCode = common.BearerAuthAPICall("https://api.github.com/user/issues", authToken, params...)
		err := json.Unmarshal(rawData, &tempIssueList)
		if err != nil {
			logger.Info().Err(err).Msg("Error unmarshalling data: ")
		} else {
			issueList = append(issueList, tempIssueList...)
		}
	}

	// according to date, increment issue count
	// recording issue count & issue names for corresponding duration
	for _, issue := range issueList {
		issues = append(issues, fmt.Sprintf("%v,%v", issue.Title, issue.URL))
	}
	return issues
}
