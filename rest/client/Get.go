package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Get(url string, headers map[string]interface{}) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("error creating a request")
	}

	setHeaders(headers, req)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.New("error doing request")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprint("error code ", resp.StatusCode))
	}

	var response map[string]interface{}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return nil, errors.New("error parsing json")
	}
	return response, nil

}

func setHeaders(headers map[string]interface{}, req *http.Request) {
	var key string
	var val interface{}
	for key, val = range headers {
		req.Header.Add(key, val.(string))
	}
	//req.Header.Add("accept-encoding", "gzip, deflate, br")
}
