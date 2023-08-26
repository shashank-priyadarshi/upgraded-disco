package ghintegration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/common"
	"strings"
	"time"

	logger "github.com/rs/zerolog/log"
)

func fetchRepoWiseData() (scmActivity []SCMActivity) {
	var rawData []byte
	repoListData, statusCode, page := []RepoResponse{}, http.StatusOK, 0

	scmActivity = addDate()

	commonqueryParams := []string{"per_page 100", "page 1"}
	repoQueryParams := commonqueryParams
	repoQueryParams = append([]string{"sort pushed"}, repoQueryParams...)

	for statusCode == http.StatusOK {
		// fetching repos for user
		tempRepoList := []RepoResponse{}
		page++
		repoQueryParams[2] = fmt.Sprintf("page %v", page)
		rawData, statusCode = common.BearerAuthAPICall("https://api.github.com/user/repos", authToken, repoQueryParams...)
		err := json.Unmarshal(rawData, &tempRepoList)
		if err != nil {
			logger.Info().Err(err).Msg("Error unmarshalling data repo data from github api: ")
		} else {
			repoListData = append(repoListData, tempRepoList...)
		}
	}

	err := createCleanDir()
	if err != nil {
		logger.Info().Err(err).Msg("Unable to create clean directory: ")
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
		err = cloneRepo(repo.URL, fmt.Sprintf("./../%v/%v", "cloned-repos", repo.Name))
		if err != nil {
			logger.Info().Err(err).Msg("Unable to clone repo: ")
		}
		commitList, pullRequestList := fetchPaginatedRepoWiseData(repo.CommitsURL, repo.PRURL, commonqueryParams)
		scmActivity = appendRepoWiseData(scmActivity, commitList, pullRequestList)
	}
	return
}

func fetchPaginatedRepoWiseData(commitsURL, prURL string, commonQueryParams []string) ([]CommitResponse, []PullRequestResponse) {
	page, prFlag, commitFlag, commitList, prList := 0, false, false, []CommitResponse{}, []PullRequestResponse{}
	for !prFlag || !commitFlag {
		// fetching repos for user
		page++
		tempCommitList, tempPRList := []CommitResponse{}, []PullRequestResponse{}
		commonQueryParams[1] = fmt.Sprintf("page %v", page)

		rawData, statusCode := common.BearerAuthAPICall(strings.Split(commitsURL, "{/sha}")[0], authToken, commonQueryParams...)
		if statusCode != http.StatusOK {
			commitFlag = true
		} else {
			if err := json.Unmarshal(rawData, &tempCommitList); err != nil {
				logger.Info().Err(err).Msg("Unable to unmarshal raw commit response: ")
			} else {
				commitList = append(commitList, tempCommitList...)
			}
		}

		rawData, statusCode = common.BearerAuthAPICall(strings.Split(prURL, "{/number}")[0], authToken, commonQueryParams...)
		if statusCode != http.StatusOK {
			prFlag = true
		} else {
			if err := json.Unmarshal(rawData, &tempPRList); err != nil {
				logger.Info().Err(err).Msg("Unable to unmarshal raw pr response: ")
			} else {
				prList = append(prList, tempPRList...)
			}
		}
	}
	return commitList, prList
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
