package ghintegration

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
)

func createCleanDir() (err error) {
	// if dir exists delete the directory
	// create clean dir
	err = os.RemoveAll("./../cloned-repos")
	if err != nil {
		return
	}
	err = os.Mkdir("./../cloned-repos", 0600)
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

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL:      urlWithCredentials,
		Progress: os.Stdout,
	})

	// err = common.RunCommand(fmt.Sprintf("%s %s", "git clone", urlWithCredentials), path)
	return
}
