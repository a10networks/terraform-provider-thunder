package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Banner struct {
	BannerInstanceExecBannerCfg BannerInstance `json:"banner,omitempty"`
}

type BannerInstance struct {
	BannerInstanceExecBannerCfgExec   BannerInstanceExecBannerCfg  `json:"exec-banner-cfg,omitempty"`
	BannerInstanceLoginBannerCfgLogin BannerInstanceLoginBannerCfg `json:"login-banner-cfg,omitempty"`
	BannerInstanceUUID                string                       `json:"uuid,omitempty"`
}

type BannerInstanceExecBannerCfg struct {
	BannerInstanceExecBannerCfgExec       int    `json:"exec,omitempty"`
	BannerInstanceExecBannerCfgExecBanner string `json:"exec-banner,omitempty"`
}

type BannerInstanceLoginBannerCfg struct {
	BannerInstanceLoginBannerCfgLogin       int    `json:"login,omitempty"`
	BannerInstanceLoginBannerCfgLoginBanner string `json:"login-banner,omitempty"`
}

func PostBanner(id string, inst Banner, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostBanner")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/banner", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Banner
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostBanner", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func PutBanner(id string, inst Banner, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutBanner")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/banner", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Banner
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutBanner", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetBanner(id string, host string) (*Banner, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetBanner")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/banner", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Banner
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetBanner", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func DeleteBanner(id string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteBanner")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/banner", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Banner
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
