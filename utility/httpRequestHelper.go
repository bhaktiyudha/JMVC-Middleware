package utilities

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestClient struct {
	Timeout time.Duration
}

type Request interface {
	Send(method string, url string, body map[string]interface{}, header map[string]string) ([]byte, int, error)
}

func (rc RequestClient) Send(method string, url string, body map[string]interface{}, header map[string]string) ([]byte, int, error) {
	client := &http.Client{
		Timeout: rc.Timeout,
	}

	bodyByte, _ := json.Marshal(body)

	request, err := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))

	if err != nil {
		return nil, 0, err
	}

	for key, value := range header {
		if request.Header.Get(key) == "" {
			request.Header.Add(key, value)
		}
	}

	res, err := client.Do(request)

	if err != nil {
		return nil, 0, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, 0, err
	}

	return resBody, res.StatusCode, nil
}