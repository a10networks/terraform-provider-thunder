package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type WebCategoryStatistics struct {
	UUID WebCategoryStatisticsInstance `json:"statistics,omitempty"`
}

type WebCategoryStatisticsInstance struct {
	Counters1 []WebCategoryStatisticsSamplingEnable `json:"sampling-enable,omitempty"`
	UUID      string                                `json:"uuid,omitempty"`
}

type WebCategoryStatisticsSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostWebCategoryStatistics(id string, inst WebCategoryStatistics, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostWebCategoryStatistics")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/web-category/statistics", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryStatistics
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostWebCategoryStatistics", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetWebCategoryStatistics(id string, host string) (*WebCategoryStatistics, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetWebCategoryStatistics")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/web-category/statistics", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryStatistics
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetWebCategoryStatistics", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
