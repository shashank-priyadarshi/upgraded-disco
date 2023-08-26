package ghintegration

import (
	"encoding/json"
	"server/config"
	"time"

	mongoconnection "server/db/mongo"

	logger "github.com/rs/zerolog/log"
)

func main() {
	pluginTriggeredRecently <- true
	rawGitHubData, rawGraphData, err := triggerIntegration() // github integration core logic to fetch data
	if err != nil {
		logger.Info().Err(err).Msg("error while fetching integration data: ")
	}

	githubData, graphData := GitHubData{}, GraphData{}
	err = json.Unmarshal(rawGitHubData, &githubData)
	if err != nil {
		logger.Info().Err(err).Msg("error while unmarshalling github data: ")
	}

	err = json.Unmarshal(rawGraphData, &graphData)
	if err != nil {
		logger.Info().Err(err).Msg("error while unmarshalling graph data: ")
	}

	// getting collection names from environment
	gitHubDataCollection, graphCollection := config.FetchConfig().Collections.GITHUBDATA, config.FetchConfig().GRAPHDATA

	// writing githubdata data to mongodb
	err = mongoconnection.WriteDataToCollection(gitHubDataCollection, githubData)
	if err != nil {
		logger.Info().Err(err).Msg("error while writing github data to mongodb: ")
	}

	// writing graph data to mongodb
	err = mongoconnection.WriteDataToCollection(graphCollection, graphData)
	if err != nil {
		logger.Info().Err(err).Msg("error while writing graph data to mongodb: ")
	}
}

// Fetching & saving raw data for past 30 days
func triggerIntegration() (gitHubData, graphData []byte, err error) {
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
