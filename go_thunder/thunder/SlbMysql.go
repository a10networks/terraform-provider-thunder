package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Mysql struct {
	Counters1 MysqlInstance `json:"mysql,omitempty"`
}

type SamplingEnable46 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type MysqlInstance struct {
	Counters1 []SamplingEnable46 `json:"sampling-enable,omitempty"`
}

func PostSlbMysql(id string, inst Mysql, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbMysql")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/mysql", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Mysql
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbMysql REQ RES..........................", m)
			err := check_api_status("PostSlbMysql", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbMysql(id string, host string) (*Mysql, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbMysql")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/mysql/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Mysql
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbMysql REQ RES..........................", m)
			err := check_api_status("GetSlbMysql", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
