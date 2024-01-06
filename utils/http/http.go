package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HttpClient(method, url, body, authMechanism, auth string) (resBody []byte, err error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return
	}

	switch authMechanism {
	case "Authorization":
		req.Header.Add(authMechanism, auth)
	default:
	}

	response, err := client.Do(req)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("status code: %v", response.StatusCode)
		return
	}

	defer response.Body.Close()

	resBody, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}

	return
}
