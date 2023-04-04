package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Mssql struct {
	Counters1 MssqlInstance `json:"mssql,omitempty"`
}
type SamplingEnable45 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type MssqlInstance struct {
	Counters1 []SamplingEnable45 `json:"sampling-enable,omitempty"`
}

func PostSlbMssql(id string, inst Mssql, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbMssql")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/mssql", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Mssql
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbMssql REQ RES..........................", m)
			err := check_api_status("PostSlbMssql", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbMssql(id string, host string) (*Mssql, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbMssql")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/mssql/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Mssql
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbMssql REQ RES..........................", m)
			err := check_api_status("GetSlbMssql", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
