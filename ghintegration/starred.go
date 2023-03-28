package ghintegration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/common"
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
			log.Println("Unable to unmarshal raw repo response: ", err)
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
