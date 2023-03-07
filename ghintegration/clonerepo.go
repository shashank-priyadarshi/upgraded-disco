package ghintegration

import (
	"fmt"
	"os"
	"server/common"
	"strings"
)

func createCleanDir() (err error) {
	// if dir exists delete the directory
	// create clean dir
	err = os.RemoveAll("cloned-repos")
	if err != nil {
		return
	}
	err = os.Mkdir("cloned-repos", 0600)
	if err != nil {
		return
	}
	return
}

func getURLWithCredentials(repoURL string) (urlWithCredentials string) {
	split := strings.Split(repoURL, "://")
	urlWithCredentials = split[0] + "://" + fmt.Sprintf("%s:%s@", userName, authToken) + split[1] + ".git"
	return
}

func cloneRepo(reqURL, path string) (err error) {
	urlWithCredentials := getURLWithCredentials(reqURL)
	err = common.RunCommand(fmt.Sprintf("%v %v %v", "git", "clone", urlWithCredentials), path)
	return
}
