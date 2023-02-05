package ghintegration

import (
	"encoding/json"
	"fmt"
	"log"
	"server/common"
)

func fetchStarredRepos() Repo {
	var repoList []RepoResponse
	var repoData Repo = Repo{}

	rawRepoList, _ := common.BearerAuthAPICall("https://api.github.com/users/shashank-priyadarshi/starred?per_page=100&page=1&sort=pushed", authToken)
	err := json.Unmarshal(rawRepoList, &repoList)
	if err != nil {
		log.Println("Unable to unmarshal raw repo response: ", err)
	}

	for _, repo := range repoList {
		repoData = Repo{
			Count: repoData.Count + 1,
			List:  append(repoData.List, fmt.Sprintf("%v,%v", repo.Name, repo.URL)),
		}
	}

	return repoData
}
