package timeular

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// API represents a client to the timeular API.
type API struct {
	url    *url.URL
	client *http.Client
	auth   *Auth
}

// Params is an convenience alias for URL query values as used with OAuth.
type Params map[string]string

// NewAPI creates a new timeular API client.
func NewAPI(baseurl, key, secret string) (*API, error) {
	parsed, err := url.Parse(baseurl)
	if err != nil {
		return nil, err
	}
	baseurl += "/developer/sign-in"

	client := &http.Client{}
	var jsonStr = []byte(`{"apiKey": "` + key + `", "apiSecret": "` + secret + `"}`)
	req, _ := http.NewRequest("POST", baseurl, bytes.NewBuffer(jsonStr))
	req.Header.Set("accept", "application/json;charset=UTF-8")
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dst := Auth{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&dst)
	if err != nil {
		return nil, err
	}

	return &API{
		url:    parsed,
		client: client,
		auth:   &dst,
	}, nil
}

// BuildURL builds the fully qualified url to the timeular API.
func (a *API) BuildURL(relpath string) string {
	if url, err := a.url.Parse(a.url.String() + relpath); err == nil {
		return url.String()
	}

	return ""
}

// try is used to encapsulate a HTTP operation and retrieve the optional
// timeular error if one happened.
func try(resp *http.Response, err error) (*http.Response, error) {
	if resp == nil {
		return resp, err
	}

	if resp.StatusCode < 300 {
		return resp, err
	}

	msg := Message{} // fetch Error Message
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&msg)
	if err != nil {
		return nil, err
	}
	if msg.Message != "" {
		return resp, errors.New(msg.Message)
	}

	return resp, err
}

// Get is a helper for all GET request with json payload.
func (a *API) Get(url string, dst interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+a.auth.Token)
	req.Header.Set("accept", "application/json;charset=UTF-8")

	resp, err := try(a.client.Do(req))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// no destination, so caller was only interested in the
	// side effects.
	if dst == nil {
		return nil
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("DEBUG: response Body:", string(body))

	bodybytes := []byte(body)
	json.Unmarshal(bodybytes, &dst)

	return nil
}

// Post is a helper for all POST request with json payload.
func (a *API) Post(url string, src interface{}, dst interface{}) error {
	payload, err := json.Marshal(src)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "Bearer "+a.auth.Token)
	req.Header.Set("accept", "application/json;charset=UTF-8")
	req.Header.Set("content-type", "application/json")

	resp, err := try(a.client.Do(req))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// no destination, so caller was only interested in the
	// side effects.
	if dst == nil {
		return nil
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&dst)
	if err != nil {
		return err
	}

	return nil
}
