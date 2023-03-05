package ghintegration

import (
	"server/config"
	"time"
)

var (
	authToken = config.FetchConfig().GITHUBTOKEN
	userName  = config.FetchConfig().GITHUBUSERNAME
	days      = int(time.Now().AddDate(0, 0, -31).Unix() / (24 * 60 * 60)) // adding -31d to today's date, converting to days, gives days elapsed from Jan 1, 1970
)
