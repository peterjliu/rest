package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Httpmethod int

const (
	GET Httpmethod = iota
	POST
	PUT
	DELETE
)

// An HTTP request
type Request struct {
	Method  Httpmethod
	Headers map[string]string
	Url     string
	Data    []byte
}

func (r *Request) AddHeader(k, v string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[k] = v
}

// Send Request
func (r Request) Do(v interface{}) error {
	var req *http.Request
	var err error
	switch r.Method {
	case GET:
		req, err = http.NewRequest("GET", r.Url, nil)
	case POST:
		req, err = http.NewRequest("POST", r.Url, bytes.NewBuffer(r.Data))
	case PUT:
		req, err = http.NewRequest("PUT", r.Url, bytes.NewBuffer(r.Data))
	case DELETE:
		req, err = http.NewRequest("DELETE", r.Url, nil)
	}
	if err != nil {
		return err
	}
	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	return parseResp(resp, &v)
}

// Parse a json http response into a struc and closes the Body.
func parseResp(resp *http.Response, output interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		var errResp interface{}
		json.Unmarshal(body, &errResp)
		return errors.New(fmt.Sprintf("Http error status code %d\n%s\n",
			resp.StatusCode, errResp))
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

// Make a http json-encoded POST using input struct as data and parses response into output struct.
// Input struct is encoded as JSON.
// HTTP headers, hdrs, are optional.
func Post(url string, hdrs map[string]string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		log.Printf("Failed to marshal data %s", data)
		return err
	}
	r := Request{
		Method:  POST,
		Headers: hdrs,
		Url:     url,
		Data:    data,
	}
	r.AddHeader("Content-Type", "application/json")
	return r.Do(&output)
}

// Make a http json-encoded PUT using input struct as data and parses response into output struct.
// Input struct is encoded as JSON.
// HTTP headers, hdrs, are optional.
func Put(url string, hdrs map[string]string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		log.Printf("Failed to marshal data %s", data)
		return err
	}
	r := Request{
		Method:  PUT,
		Headers: hdrs,
		Url:     url,
		Data:    data,
	}
	r.AddHeader("Content-Type", "application/json")
	return r.Do(&output)
}

// Make a GET request parses response into output struct.
// HTTP headers, hdrs, are optional.
func Get(url string, hdrs map[string]string, output interface{}) error {
	r := Request{
		Method:  GET,
		Headers: hdrs,
		Url:     url,
	}
	return r.Do(&output)
}
