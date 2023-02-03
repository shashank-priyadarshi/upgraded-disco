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
	rawData, err := triggerIntegration()
	if err != nil {
		panic(err)
	}

	var githubData GitHubData
	err = json.Unmarshal(rawData, &githubData)
	if err != nil {
		panic(err)
	}

	gitHubDataCollection, issueCollection := config.FetchConfig().Collections.GITHUBDATA, config.FetchConfig().Collections.TODOS

	err = mongoconnection.WriteDataToCollection(gitHubDataCollection, githubData)
	if err != nil {
		panic(err)
	}
	err = mongoconnection.WriteDataToCollection(issueCollection, struct {
		Issues []string `json:"issues"`
	}{Issues: githubData.Issues})
	if err != nil {
		panic(err)
	}
	fmt.Println("Data written to MongoDB successfully!")
}

func triggerIntegration() ([]byte, error) {
	fmt.Println("Triggered Integration...")

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println(time.LoadLocation("Asia/Calcutta"))
	}
	fmt.Println("Location: ", loc)
	fmt.Println("Time: ", time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)))

	repoCount, repoList, scmActivity := fetchRepoWiseData()
	scmActivity, issues := getIssueData(scmActivity)

	gitHubData := GitHubData{
		Repos: Repo{
			List:  repoList,
			Count: repoCount,
		},
		WeekData:     scmActivity,
		Issues:       issues,
		StarredRepos: fetchStarredRepos(),
		Time:         time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)),
	}
	fmt.Println("Plugin execution completed!")
	return json.Marshal(gitHubData)
}

func addDate() []SCMActivity {
	scmActivity := []SCMActivity{}

	num := 0
	for num < 7 {
		scmActivity = append(scmActivity, SCMActivity{
			Date: time.Now().AddDate(0, 0, -num).String(),
		})
		num++
	}

	return scmActivity
}
