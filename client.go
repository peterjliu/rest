package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Make a http json-encoded post using input struct as body and parses response into output struct
func Post(url string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return parseResp(resp, &output)
}

func Get(url string, output interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	return parseResp(resp, &output)
}

// parse a json http response into a struct
func parseResp(resp *http.Response, output interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("http error status code %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}
	return nil
}
