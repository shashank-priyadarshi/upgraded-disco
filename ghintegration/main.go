package ghintegration

import (
	"encoding/json"
	"fmt"
	"server/config"
	"server/mongoconnection"
	"time"
)

func main() {
	fmt.Println("Plugin execution started!")
	pluginTriggeredRecently <- true
	rawData, err := triggerIntegration() // github integration core logic to fetch data
	if err != nil {
		panic(err)
	}

	var githubData GitHubData
	err = json.Unmarshal(rawData, &githubData)
	if err != nil {
		panic(err)
	}

	// getting collection names from environment
	gitHubDataCollection, issueCollection := config.FetchConfig().Collections.GITHUBDATA, config.FetchConfig().Collections.TODOS

	// writing githubdata data to mongodb
	err = mongoconnection.WriteDataToCollection(gitHubDataCollection, githubData)
	if err != nil {
		panic(err)
	}

	// writing issue data to mongodb
	err = mongoconnection.WriteDataToCollection(issueCollection, struct {
		Issues []string `json:"issues"`
	}{Issues: getIssueData()})
	if err != nil {
		panic(err)
	}
	fmt.Println("Data written to MongoDB successfully!")
}

// Fetching & saving raw data for past 30 days
func triggerIntegration() ([]byte, error) {
	fmt.Println("Triggered Integration...")

	_, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println(time.LoadLocation("Asia/Calcutta"))
	}

	scmActivity := fetchRepoWiseData() // fetching week wise pr, commit, loc data

	gitHubData := GitHubData{
		WeekData:     scmActivity,                                                  // saving week wise pr, commit, loc data
		StarredRepos: fetchStarredRepos(),                                          // fetching list of starred repos
		Time:         time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)), // saving time for latest trigger
	}
	fmt.Println("Plugin execution completed!")
	return json.Marshal(gitHubData)
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
