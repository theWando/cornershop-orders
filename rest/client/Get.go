package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func Get(url string, headers map[string]interface{}, str interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.New("error creating a request")
	}

	setHeaders(headers, req)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return errors.New("error doing request")
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprint("error code ", resp.StatusCode))
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&str); err != nil {
		log.Fatal(err)
		return errors.New("error parsing json")
	}
	return nil

}

func setHeaders(headers map[string]interface{}, req *http.Request) {
	var key string
	var val interface{}
	for key, val = range headers {
		req.Header.Add(key, val.(string))
	}
	//req.Header.Add("accept-encoding", "gzip, deflate, br")
}
