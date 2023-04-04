package common

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	logger "github.com/rs/zerolog/log"
)

// func PaginatedAPICall(reqURL, authToken string, unmarshalTo []interface{}, queryParams ...string) ([]interface{}, error) {
// 	index := len(queryParams) - 1
// 	page := 0
// 	for {
// 		var tempRepoList []interface{}
// 		page++
// 		queryParams[index] = fmt.Sprintf("page %v", page)
// 		rawData, statusCode := BearerAuthAPICall(reqURL, authToken, queryParams...)
// 		err := json.Unmarshal(rawData, &tempRepoList)
// 		if err != nil {
// 			logger.Info().Msg("Unable to unmarshal raw repo response: ", err)
// 			return unmarshalTo, err
// 		}
// 		if statusCode != http.StatusOK {
// 			log.Printf("Status code while making API call to %v: %v", reqURL, statusCode)
// 			return unmarshalTo, nil
// 		} else if err == nil {
// 			unmarshalTo = append(unmarshalTo, tempRepoList...)
// 		}
// 	}
// }

func BearerAuthAPICall(reqURL, authToken string, queryParams ...string) ([]byte, int) {
	client := http.Client{}
	request, err := http.NewRequest("GET", reqURL, bytes.NewBuffer([]byte("")))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authToken))
	if err != nil {
		logger.Info().Err(err).Msg(fmt.Sprintf("err in making request to %v: %v", reqURL, err))
		return []byte{}, 503
	}

	if len(queryParams) > 0 {
		request.URL.RawQuery = addQueryParameters(request.URL.Query(), queryParams...)
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		logger.Info().Err(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
		return []byte{}, resp.StatusCode
	}

	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, resp.StatusCode
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, 503
	}

	if strings.EqualFold(string(respBody), "[]") {
		return []byte{}, 503
	}
	return respBody, resp.StatusCode
}

func NoAuthAPICall(reqURL, origin string, reqBody []byte, queryParams ...string) ([]byte, int) {
	timeOut := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeOut,
	}

	request, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", origin)
	if err != nil {
		logger.Info().Err(err).Msg("err in creating new request: ")
		return []byte{}, 503
	}

	if len(queryParams) > 0 {
		request.URL.RawQuery = addQueryParameters(request.URL.Query(), queryParams...)
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		logger.Info().Err(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
		return []byte{}, resp.StatusCode
	}

	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, resp.StatusCode
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Info().Err(err).Msg("err in reading req response: ")
		return []byte{}, 503
	}

	return respBody, resp.StatusCode
}

func addQueryParameters(q url.Values, queryParams ...string) string {
	for _, param := range queryParams {
		keyVal := strings.SplitN(param, " ", 2)
		q.Add(keyVal[0], keyVal[1])
	}
	return q.Encode()
}
