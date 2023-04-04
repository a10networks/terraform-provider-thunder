package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type CommonConnRateLimitSrcIP struct {
	UUID CommonConnRateLimitSrcIPInstance `json:"src-ip,omitempty"`
}

type CommonConnRateLimitSrcIPInstance struct {
	Protocol     string `json:"protocol,omitempty"`
	Log          int    `json:"log,omitempty"`
	LockOut      int    `json:"lock-out,omitempty"`
	LimitPeriod  string `json:"limit-period,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	ExceedAction int    `json:"exceed-action,omitempty"`
	Shared       int    `json:"shared,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

func PostSlbCommonConnRateLimitSrcIP(id string, inst CommonConnRateLimitSrcIP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbCommonConnRateLimitSrcIP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/common/conn-rate-limit/src-ip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CommonConnRateLimitSrcIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbCommonConnRateLimitSrcIP REQ RES..........................", m)
			err := check_api_status("PostSlbCommonConnRateLimitSrcIP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbCommonConnRateLimitSrcIP(id string, name string, host string) (*CommonConnRateLimitSrcIP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbCommonConnRateLimitSrcIP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/common/conn-rate-limit/src-ip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CommonConnRateLimitSrcIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbCommonConnRateLimitSrcIP REQ RES..........................", m)
			err := check_api_status("GetSlbCommonConnRateLimitSrcIP", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbCommonConnRateLimitSrcIP(id string, name string, inst CommonConnRateLimitSrcIP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbCommonConnRateLimitSrcIP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/common/conn-rate-limit/src-ip/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CommonConnRateLimitSrcIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbCommonConnRateLimitSrcIP REQ RES..........................", m)
			err := check_api_status("PutSlbCommonConnRateLimitSrcIP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbCommonConnRateLimitSrcIP(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbCommonConnRateLimitSrcIP")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/common/conn-rate-limit/src-ip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CommonConnRateLimitSrcIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
