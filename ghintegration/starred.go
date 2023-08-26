package ghintegration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/common"

	logger "github.com/rs/zerolog/log"
)

// fetching starred repo list
func fetchStarredRepos() (repoData Repo) {
	var rawData []byte
	page, statusCode, params, repoList := 0, 200, queryParams, []RepoResponse{}
	params = append([]string{"sort pushed"}, params...)

	for statusCode == http.StatusOK {
		// fetching issues for user
		page++
		tempIssueList := []RepoResponse{}
		params[2] = fmt.Sprintf("page %v", page)
		rawData, statusCode = common.BearerAuthAPICall("https://api.github.com/users/shashank-priyadarshi/starred", authToken, params...)
		err := json.Unmarshal(rawData, &tempIssueList)
		if err != nil {
			logger.Info().Err(err).Msg("Error unmarshalling data starred repo data for user from github api: ")
		} else {
			repoList = append(repoList, tempIssueList...)
		}
	}

	repoData = Repo{}

	for _, repo := range repoList {
		repoData = Repo{
			Count: repoData.Count + 1,
			List:  append(repoData.List, fmt.Sprintf("%v,%v", repo.Name, repo.URL)),
		}
	}

	return
}
