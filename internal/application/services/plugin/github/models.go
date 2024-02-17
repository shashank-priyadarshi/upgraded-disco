package main

type starredRepositoriesResponse struct {
	Data struct {
		Viewer struct {
			StarredRepositories struct {
				TotalCount int      `json:"totalCount"`
				PageInfo   PageInfo `json:"pageInfo"`
				Nodes      []Node   `json:"nodes"`
			} `json:"starredRepositories"`
		} `json:"viewer"`
	} `json:"data"`
}

type followedUsersResponse struct {
	Data struct {
		Viewer struct {
			Following struct {
				TotalCount int      `json:"totalCount"`
				PageInfo   PageInfo `json:"pageInfo"`
				Nodes      []Node   `json:"nodes"`
			} `json:"following"`
		} `json:"viewer"`
	} `json:"data"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

type Node struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Languages   struct {
		Nodes []Node `json:"nodes"`
	} `json:"languages"`
}
