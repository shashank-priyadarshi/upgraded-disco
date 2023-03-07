package ghintegration

import (
	"encoding/json"
	"fmt"
	"log"
	"server/common"
	"strings"
	"time"
)

func fetchRepoWiseData() (scmActivity []SCMActivity) {
	var repoListData []RepoResponse
	var commitList []CommitResponse
	var pullRequestList []PullRequestResponse

	scmActivity = addDate()

	// fetching repos for user
	rawRepoList, _ := common.BearerAuthAPICall("https://api.github.com/user/repos?per_page=100&page=1&sort=pushed", authToken)
	err := json.Unmarshal(rawRepoList, &repoListData)
	if err != nil {
		log.Println("Unable to unmarshal raw repo response: ", err)
	}

	err = createCleanDir()
	if err != nil {
		log.Println("Unable to create clean dir: ", err)
	}

	for _, repo := range repoListData {
		if repo.Forked { // if repo is forked, do not execute logic
			continue
		}

		prDate, err := time.Parse("2006-01-02", repo.PushedAt[:10])
		if err != nil {
			log.Printf("Unable to parse date: %v, %v", repo.PushedAt, err)
		}
		if 31-(int(prDate.Unix()/(24*60*60))-days) > 31 {
			continue
		}

		// cloning repo
		err = cloneRepo(repo.URL, fmt.Sprintf("./%v/", "cloned-repos"))
		if err != nil {
			log.Println("Unable to clone repo: ", err)
		}

		// api call to fetch commits in repo.Name for this iteration
		rawCommitResponse, _ := common.BearerAuthAPICall(fmt.Sprintf("%v?per_page=100&page=1", strings.Split(repo.CommitsURL, "{/sha}")[0]), authToken)
		err = json.Unmarshal(rawCommitResponse, &commitList)
		if err != nil {
			log.Println("Unable to unmarshal raw commit response: ", err)
		}

		// api call to fetch pull requests in repo.Name for this iteration
		rawPullRequestResponse, _ := common.BearerAuthAPICall(fmt.Sprintf("%v?per_page=100&page=1", strings.Split(repo.PRURL, "{/number}")[0]), authToken)
		err = json.Unmarshal(rawPullRequestResponse, &pullRequestList)
		if err != nil {
			log.Println("Unable to unmarshal raw pull request response: ", err)
		}

		scmActivity = appendRepoWiseData(scmActivity, commitList, pullRequestList)
	}
	return
}

func appendRepoWiseData(scmActivity []SCMActivity, commitList []CommitResponse, pullRequestList []PullRequestResponse) (temp []SCMActivity) {
	// according to date, increment commit data
	// recording commit count for corresponding duration
	for _, commit := range commitList {
		commitDate, _ := time.Parse("2006-01-02", commit.Commit.Committer.Date[:10])
		timeElapsed := 31 - (int(commitDate.Unix()/(24*60*60)) - days)
		if timeElapsed < 31 && timeElapsed >= 0 {
			scmActivity[timeElapsed].Commit++
		}
	}

	// according to date, increment pr data
	// recording pr count for corresponding duration
	for _, pr := range pullRequestList {
		prDate, _ := time.Parse("2006-01-02", pr.CreatedAt[:10])
		timeElapsed := 31 - (int(prDate.Unix()/(24*60*60)) - days) // days since event = days since Jan 1, '70 - (days since Jan 1, '70 to 31 days ago)
		if timeElapsed < 31 && timeElapsed >= 0 {
			scmActivity[timeElapsed].PR++
		}
	}
	temp = scmActivity
	return
}
