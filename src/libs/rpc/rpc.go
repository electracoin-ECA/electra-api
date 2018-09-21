package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Electra-project/electra-api/src/helpers"
)

func query(method string, params []string, response interface{}) error {
	daemonURI := "http://localhost:5788"

	var quotedParams []string
	for _, param := range params {
		quotedParams = append(quotedParams, fmt.Sprintf(`"%s"`, param))
	}

	reqData := fmt.Sprintf(`{"method":"%s","params":[%s]}`, method, strings.Join(quotedParams, ","))
	// helpers.Log(reqData)
	helpers.LogInfo("using the reqData " + reqData)
	reqDataBuffer := bytes.NewBuffer([]byte(reqData))
	req, err := http.NewRequest("POST", daemonURI, reqDataBuffer)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("user", "pass")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	helpers.Log(res.Status + " " + fmt.Sprint(res.StatusCode) + " " + string(bodyBytes))

	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return err
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return err
	}

	return nil
}

func QueryBytes(method string, params []string) ([]byte, error) {
	bytes, _, err := QueryBytesAndResp(method, params)
	return bytes, err
}

func QueryBytesAndResp(method string, params []string) ([]byte, *http.Response, error) {
	daemonURI := "http://localhost:5788"

	var quotedParams []string
	for _, param := range params {
		quotedParams = append(quotedParams, fmt.Sprintf(`"%s"`, param))
	}

	reqData := fmt.Sprintf(`{"method":"%s","params":[%s]}`, method, strings.Join(quotedParams, ","))
	// helpers.Log(reqData)
	helpers.LogInfo("using the reqData " + reqData)
	reqDataBuffer := bytes.NewBuffer([]byte(reqData))
	req, err := http.NewRequest("POST", daemonURI, reqDataBuffer)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("user", "pass")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return nil, nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	helpers.Log(res.Status + " " + fmt.Sprint(res.StatusCode) + " " + string(bodyBytes))

	if err != nil {
		return nil, nil, err
	}
	return bodyBytes, res, nil
}
