package exporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type QueryResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  [2]interface{}    `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func GetMetricsWithQuery(endpoint string, headers map[string]string, rule string) (*QueryResponse, error) {
	url := fmt.Sprintf("%s/api/v1/query?query=%s", endpoint, rule)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var queryResp QueryResponse
	err = json.Unmarshal(body, &queryResp)
	if err != nil {
		return nil, err
	}

	return &queryResp, nil
}
