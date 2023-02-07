package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Diameter struct {
	UUID DiameterInstance `json:"diameter,omitempty"`
}

type MessageCodeList struct {
	MessageCode int `json:"message-code,omitempty"`
}
type AvpList struct {
	Int32     int    `json:"int32,omitempty"`
	Mandatory int    `json:"mandatory,omitempty"`
	String    string `json:"string,omitempty"`
	Avp       int    `json:"avp,omitempty"`
	Int64     int    `json:"int64,omitempty"`
}
type OriginHost struct {
	UUID           string `json:"uuid,omitempty"`
	OriginHostName string `json:"origin-host-name,omitempty"`
}
type DiameterInstance struct {
	AvpString               string            `json:"avp-string,omitempty"`
	TerminateOnCcaT         int               `json:"terminate-on-cca-t,omitempty"`
	MessageCode             []MessageCodeList `json:"message-code-list,omitempty"`
	Avp                     []AvpList         `json:"avp-list,omitempty"`
	ServiceGroupName        string            `json:"service-group-name,omitempty"`
	UUID                    string            `json:"uuid,omitempty"`
	IdleTimeout             int               `json:"idle-timeout,omitempty"`
	CustomizeCea            int               `json:"customize-cea,omitempty"`
	ProductName             string            `json:"product-name,omitempty"`
	DwrUpRetry              int               `json:"dwr-up-retry,omitempty"`
	ForwardUnknownSessionID int               `json:"forward-unknown-session-id,omitempty"`
	AvpCode                 int               `json:"avp-code,omitempty"`
	VendorID                int               `json:"vendor-id,omitempty"`
	SessionAge              int               `json:"session-age,omitempty"`
	LoadBalanceOnSessionID  int               `json:"load-balance-on-session-id,omitempty"`
	Name                    string            `json:"name,omitempty"`
	DwrTime                 int               `json:"dwr-time,omitempty"`
	UserTag                 string            `json:"user-tag,omitempty"`
	OriginRealm             string            `json:"origin-realm,omitempty"`
	OriginHostName          OriginHost        `json:"origin-host,omitempty"`
	MultipleOriginHost      int               `json:"multiple-origin-host,omitempty"`
	ForwardToLatestServer   int               `json:"forward-to-latest-server,omitempty"`
}

func PostSlbTemplateDiameter(id string, inst Diameter, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateDiameter")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/diameter", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Diameter
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateDiameter REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateDiameter", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateDiameter(id string, name string, host string) (*Diameter, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateDiameter")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/diameter/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Diameter
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateDiameter REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateDiameter", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateDiameter(id string, name string, inst Diameter, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateDiameter")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/diameter/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Diameter
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateDiameter REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateDiameter", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateDiameter(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateDiameter")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/diameter/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Diameter
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
