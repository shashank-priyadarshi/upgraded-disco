package ghintegration

import (
	"encoding/json"
	"fmt"
	"server/config"
	"time"

	mongoconnection "server/db/mongo"
)

func main() {
	fmt.Println("Plugin execution started!")
	pluginTriggeredRecently <- true
	rawGitHubData, rawGraphData, err := triggerIntegration() // github integration core logic to fetch data
	if err != nil {
		fmt.Println(err)
	}

	githubData, graphData := GitHubData{}, GraphData{}
	err = json.Unmarshal(rawGitHubData, &githubData)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(rawGraphData, &graphData)
	if err != nil {
		fmt.Println(err)
	}

	// getting collection names from environment
	gitHubDataCollection, issueCollection, graphCollection := config.FetchConfig().Collections.GITHUBDATA, config.FetchConfig().Collections.TODOS, config.FetchConfig().GRAPHDATA

	// writing githubdata data to mongodb
	err = mongoconnection.WriteDataToCollection(gitHubDataCollection, githubData)
	if err != nil {
		fmt.Println(err)
	}

	// writing graph data to mongodb
	err = mongoconnection.WriteDataToCollection(graphCollection, graphData)
	if err != nil {
		fmt.Println(err)
	}

	// writing issue data to mongodb
	err = mongoconnection.WriteDataToCollection(issueCollection, struct {
		Issues []string `json:"issues"`
	}{Issues: getIssueData()})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data written to MongoDB successfully!")
}

// Fetching & saving raw data for past 30 days
func triggerIntegration() (gitHubData, graphData []byte, err error) {
	fmt.Println("Triggered Integration...")
	// fetching week wise pr, commit, loc data
	// saving date for past 30 days in array
	graphData, _ = json.Marshal(GraphData{
		WeekData: fetchRepoWiseData(),
	})

	// saving week wise pr, commit, loc data
	gitHubData, _ = json.Marshal(GitHubData{
		StarredRepos: fetchStarredRepos(),                                          // fetching list of starred repos
		Time:         time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)), // saving time for latest trigger
	})
	fmt.Println("Plugin execution completed!")
	return
}

// saving date for past 30 days in arraya
func addDate() (scmActivity []SCMActivity) {
	scmActivity = []SCMActivity{}
	num := 0
	for num < 31 {
		scmActivity = append(scmActivity, SCMActivity{
			Date: time.Now().AddDate(0, 0, -num).String(),
		})
		num++
	}
	return
}
