package repositories

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://cornershopapp.com/api/v3/orders/"

func Get(id string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprint(url, id), nil)
	if err != nil {
		return "", errors.New("error creating a request")
	}

	req.Header.Add("cookie", "")
	req.Header.Add("accept", "application/json")
	//req.Header.Add("accept-encoding", "gzip, deflate, br")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New("error doing request")
	}

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprint("error code ", resp.StatusCode))
	}

	fmt.Println("Successful response ", resp.StatusCode)
	//var response map[string]interface{}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("error reading bytes")
	}

	//bytes = bytes.TrimPrefix(bytes, []byte("\xef\xbb\xbf"))
	/*err = json.Unmarshal(bytes, response)

	if err != nil {
		return nil, errors.New("error parsing json")
	}*/
	return string(bytes), nil
}
