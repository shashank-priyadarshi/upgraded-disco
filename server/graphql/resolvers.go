package graphql

// func resolveTodos() []byte {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%v/list", config.FetchConfig().TODOORIGIN), nil)

// 	if err != nil {
// 		panic(fmt.Sprintf("error while creating http request to endpoint: %v", err))
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(fmt.Sprintf("error while sending http request to endpoint: %v", err))
// 	}
// 	logger.Info().Msg(resp.Status)
// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(fmt.Sprintf("error while reading response body: %v", err))
// 	}
// 	defer resp.Body.Close()
// 	return respBody
// }
