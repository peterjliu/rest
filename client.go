package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Make a http json-encoded POST using input struct as data and parses response into output struct.
// HTTP headers, hdrs, are optional.
func Post(url string, hdrs map[string]string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range hdrs {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return parseResp(resp, &output)
}

// Make a GET request parses response into output struct.
// HTTP headers, hdrs, are optional.
func Get(url string, hdrs map[string]string, output interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	for k, v := range hdrs {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	return parseResp(resp, &output)
}

// Parse a json http response into a struc and closes the Body.
func parseResp(resp *http.Response, output interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Http error status code %d", resp.StatusCode))
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
