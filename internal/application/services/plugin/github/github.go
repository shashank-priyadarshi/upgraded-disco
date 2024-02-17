package main

import (
	"errors"
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"net/http"
)

const (
	STARRED_REPOSITORIES           = `query { viewer { starredRepositories(first: 100, orderBy: {field: STARRED_AT, direction: DESC}){ totalCount pageInfo{ hasNextPage endCursor } nodes{ name url description languages(first: 10){nodes{name}}}}}}`
	STARRED_REPOSITORIES_PAGINATOR = `query { viewer { starredRepositories(first: 100, after: "%s", orderBy: {field: STARRED_AT, direction: DESC}){ totalCount pageInfo{ hasNextPage endCursor } nodes{ name url description languages(first: 10){nodes{name}}}}}}`
	FOLLOWED_USERS                 = `query { viewer { following(first: 100) { totalCount pageInfo { hasNextPage endCursor } nodes { name url}}}}`
	FOLLOWED_USERS_PAGINATOR       = `query { viewer { following(first: 100, after: "%s") { totalCount pageInfo { hasNextPage endCursor } nodes { name url}}}}`
)

type GitHub struct {
	endpoint, token string
}

func (g GitHub) Run(i ...interface{}) (constants.HTTPStatusCode, error) {

	if config, ok := i[0].(models.PluginsConfig); !ok {
		return constants.HTTPStatusPreconditionFailed, errors.New("this plugin requires GitHub API endpoint and User's GitHub token to run")
	} else {
		g = GitHub{
			endpoint: config.Endpoint,
			token:    config.Token,
		}
	}

	return g.run()
}

func (g GitHub) run() (constants.HTTPStatusCode, error) {
	if err := g.getStarredRepos(); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error while fetching starred repositories: %v", err)
	}

	if err := g.getFollowedUsers(); err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error while fetching starred repositories: %v", err)
	}

	return http.StatusOK, nil
}

func (g GitHub) getStarredRepos() (err error) {
	var starredRepositoriesCount int
	var starredRepositories []Node
	var data interface{}

	if data, err = getData(g.endpoint, STARRED_REPOSITORIES, g.token, starredRepositoriesResponse{}); err != nil {
		return err
	}

	response := data.(starredRepositoriesResponse)
	starredRepositoriesCount = response.Data.Viewer.StarredRepositories.TotalCount
	starredRepositories = response.Data.Viewer.StarredRepositories.Nodes

	for response.Data.Viewer.StarredRepositories.PageInfo.HasNextPage {
		if data, err = getData(g.endpoint, fmt.Sprintf(STARRED_REPOSITORIES_PAGINATOR, response.Data.Viewer.StarredRepositories.PageInfo.EndCursor), g.token, starredRepositoriesResponse{}); err != nil {
			return err
		}
		response = data.(starredRepositoriesResponse)
		starredRepositories = append(starredRepositories, response.Data.Viewer.StarredRepositories.Nodes...)
	}

	fmt.Println(starredRepositoriesCount)

	return
}

func (g GitHub) getFollowedUsers() (err error) {
	var followedUsersCount int
	var followedUsers []Node
	var data interface{}

	if data, err = getData(g.endpoint, FOLLOWED_USERS, g.token, followedUsersResponse{}); err != nil {
		return err
	}

	response := data.(followedUsersResponse)
	followedUsers = response.Data.Viewer.Following.Nodes
	followedUsersCount = response.Data.Viewer.Following.TotalCount

	for response.Data.Viewer.Following.PageInfo.HasNextPage {
		if data, err = getData(g.endpoint, fmt.Sprintf(FOLLOWED_USERS_PAGINATOR, response.Data.Viewer.Following.PageInfo.EndCursor), g.token, followedUsersResponse{}); err != nil {
			return err
		}

		response = data.(followedUsersResponse)
		followedUsers = append(followedUsers, response.Data.Viewer.Following.Nodes...)
	}
	fmt.Println(followedUsersCount)

	return
}

func getData(endpoint, query, token string, data interface{}) (interface{}, error) {

	//responseBytes, err := http2.HttpClient("POST", endpoint, query, "Bearer", token)
	//if err != nil {
	//	return err, fmt.Errorf("HTTP POST request to endpoint %s to fetch starred repositories failed with status err: %v", endpoint, err)
	//}
	//
	//if err = json.Unmarshal(responseBytes, &data); err != nil {
	//	return nil, fmt.Errorf("error unmarshaling starred repository payload: %v", err)
	//}

	return data, nil
}
