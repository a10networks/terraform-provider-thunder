package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Monitor struct {
	UUID MonitorInstance `json:"monitor,omitempty"`
}

type ClearCfg struct {
	ClearSequence    int    `json:"clear-sequence,omitempty"`
	ClearAllSequence int    `json:"clear-all-sequence,omitempty"`
	Sessions         string `json:"sessions,omitempty"`
}
type LinkEnableCfg struct {
	EnaSequence int `json:"ena-sequence,omitempty"`
	Enaeth      int `json:"enaeth,omitempty"`
}
type LinkUpCfg struct {
	LinkupEthernet3 int `json:"linkup-ethernet3,omitempty"`
	LinkupEthernet2 int `json:"linkup-ethernet2,omitempty"`
	LinkupEthernet1 int `json:"linkup-ethernet1,omitempty"`
	LinkUpSequence1 int `json:"link-up-sequence1,omitempty"`
	LinkUpSequence3 int `json:"link-up-sequence3,omitempty"`
	LinkUpSequence2 int `json:"link-up-sequence2,omitempty"`
}
type LinkDownCfg struct {
	LinkDownSequence1 int `json:"link-down-sequence1,omitempty"`
	LinkDownSequence2 int `json:"link-down-sequence2,omitempty"`
	LinkDownSequence3 int `json:"link-down-sequence3,omitempty"`
	LinkdownEthernet2 int `json:"linkdown-ethernet2,omitempty"`
	LinkdownEthernet3 int `json:"linkdown-ethernet3,omitempty"`
	LinkdownEthernet1 int `json:"linkdown-ethernet1,omitempty"`
}
type LinkDisableCfg struct {
	DisSequence int `json:"dis-sequence,omitempty"`
	Diseth      int `json:"diseth,omitempty"`
}
type MonitorInstance struct {
	ClearSequence     []ClearCfg       `json:"clear-cfg,omitempty"`
	UUID              string           `json:"uuid,omitempty"`
	EnaSequence       []LinkEnableCfg  `json:"link-enable-cfg,omitempty"`
	LinkupEthernet3   []LinkUpCfg      `json:"link-up-cfg,omitempty"`
	LinkDownSequence1 []LinkDownCfg    `json:"link-down-cfg,omitempty"`
	UserTag           string           `json:"user-tag,omitempty"`
	DisSequence       []LinkDisableCfg `json:"link-disable-cfg,omitempty"`
	MonitorRelation   string           `json:"monitor-relation,omitempty"`
	ID                int              `json:"id,omitempty"`
}

func PostSlbTemplateMonitor(id string, inst Monitor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateMonitor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/monitor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Monitor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("POSTSlbTemplateMonitor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateMonitor(id string, name string, host string) (*Monitor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateMonitor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/monitor/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Monitor
		logger.Println("data -->  ", data)
		//return err
		logger.Printf("type  data   %T\n", data)
		//return err
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateMonitor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateMonitor(id string, name string, inst Monitor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateMonitor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/monitor/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Monitor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PUT REQ RES..........................", m)
			err := check_api_status("PUTSlbTemplateMonitor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateMonitor(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateMonitor")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/monitor/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Monitor
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
