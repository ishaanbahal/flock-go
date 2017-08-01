package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const contentType = "application/json"

// HTTPRequestPost -> Send post HTTP request to flock
func HTTPRequestPost(endpoint string, requestData interface{}) ([]byte, error) {
	url := BaseURL + endpoint
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	httpRes, err := http.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()
	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode != 200 {
		var errResponse ErrorResponse
		_ = json.Unmarshal(body, &errResponse)
		str := fmt.Sprintf("Error: %s | Param: %s \nDescriptin: %s", errResponse.ErrorName, errResponse.Parameter, errResponse.Description)
		return nil, errors.New(str)
	}
	return body, nil
}
