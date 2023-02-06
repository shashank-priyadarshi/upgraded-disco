package ghintegration

import (
	"server/config"
	"time"
)

var (
	authToken = config.FetchConfig().GITHUBTOKEN
	days      = int(time.Now().AddDate(0, 0, -7).Unix() / (24 * 60 * 60))
)
