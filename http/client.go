package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ContentJson = "application/json"
)

type Client struct {
	BaseUrl string
}

func (c *Client) Post(fullPath string, param, header map[string]string, body map[string]interface{}) (response []byte, err error) {
	fullPath = fmt.Sprintf("%v%v", c.BaseUrl, fullPath)
	requestBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	request, err := http.NewRequest(http.MethodPost, fullPath, bytes.NewReader(requestBody))
	if err != nil {
		return
	}
	c.SetHeader(request, header)
	parmValue := url.Values{}
	for key, value := range param {
		parmValue.Set(key, value)
	}
	request.URL.RawQuery = parmValue.Encode()
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	response, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return
}

func (c *Client) Get(fullPath string, param, header map[string]string) (response []byte, err error) {
	fullPath = fmt.Sprintf("%v%v", c.BaseUrl, fullPath)
	request, err := http.NewRequest(http.MethodGet, fullPath, nil)
	if err != nil {
		return
	}
	c.SetHeader(request, header)
	parmValue := url.Values{}
	for key, value := range param {
		parmValue.Set(key, value)
	}
	request.URL.RawQuery = parmValue.Encode()
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	response, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return
}

// 设置请求头，
func (c *Client) SetHeader(request *http.Request, header map[string]string) {

	for key, value := range header {
		request.Header.Set(key, value)
	}

}
