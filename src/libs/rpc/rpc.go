package rpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Electra-project/electra-api/src/helpers"
)

func query(method string, params []string, response interface{}) error {
	daemonURI := "http://127.0.0.1:5788"

	reqParams, err := json.Marshal(params)
	if err != nil {
		helpers.LogErr("Error: " + err.Error())

		return err
	}

	reqData := `{"method":"` + method + `","params":` + string(reqParams) + `}`
	// helpers.Log(reqData)
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
