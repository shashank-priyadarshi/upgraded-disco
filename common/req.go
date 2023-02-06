package common

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func BearerAuthAPICall(reqURL, authToken string) ([]byte, int) {
	timeOut := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeOut,
	}

	request, err := http.NewRequest("GET", reqURL, bytes.NewBuffer([]byte("")))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authToken))
	if err != nil {
		log.Println("err in creating new request: ", err)
		return []byte{}, 503
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
		return []byte{}, resp.StatusCode
	}

	if err != nil {
		log.Println("err in making bearerAuth req: ", err)
		return []byte{}, resp.StatusCode
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("err in reading req response: ", err)
		return []byte{}, 503
	}

	return respBody, resp.StatusCode
}

func NoAuthAPICall(reqURL, origin string, reqBody []byte) ([]byte, int) {
	timeOut := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeOut,
	}

	request, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Origin", origin)
	if err != nil {
		log.Println("err in creating new request: ", err)
		return []byte{}, 503
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
		return []byte{}, resp.StatusCode
	}

	if err != nil {
		log.Println("err in making bearerAuth req: ", err)
		return []byte{}, resp.StatusCode
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("err in reading req response: ", err)
		return []byte{}, 503
	}

	log.Println(string(respBody))
	return respBody, resp.StatusCode
}
