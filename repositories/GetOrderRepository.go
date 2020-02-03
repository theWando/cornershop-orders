package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const url = "https://cornershopapp.com/api/v3/orders/"

func Get(id string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", fmt.Sprint(url, id), nil)
	if err != nil {
		return nil, errors.New("error creating a request")
	}

	req.Header.Add("cookie", os.Getenv("CORNERSHOP_COOKIE"))
	req.Header.Add("accept", "application/json")
	//req.Header.Add("accept-encoding", "gzip, deflate, br")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.New("error doing request")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprint("error code ", resp.StatusCode))
	}

	fmt.Println("Successful response ", resp.StatusCode)
	var response map[string]interface{}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return nil, errors.New("error parsing json")
	}
	return response, nil
}
