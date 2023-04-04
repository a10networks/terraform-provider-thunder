package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwRadiusServer struct {
	AccountingStart FwRadiusServerInstance `json:"server,omitempty"`
}

type FwRadiusServerInstance struct {
	AccountingInterimUpdate string                         `json:"accounting-interim-update,omitempty"`
	AccountingOn            string                         `json:"accounting-on,omitempty"`
	AccountingStart         string                         `json:"accounting-start,omitempty"`
	AccountingStop          string                         `json:"accounting-stop,omitempty"`
	PrefixNumber            []FwRadiusServerAttribute      `json:"attribute,omitempty"`
	AttributeName           string                         `json:"attribute-name,omitempty"`
	CustomAttributeName     string                         `json:"custom-attribute-name,omitempty"`
	Encrypted               string                         `json:"encrypted,omitempty"`
	ListenPort              int                            `json:"listen-port,omitempty"`
	IPList                  FwRadiusServerRemote           `json:"remote,omitempty"`
	Counters1               []FwRadiusServerSamplingEnable `json:"sampling-enable,omitempty"`
	Secret                  int                            `json:"secret,omitempty"`
	SecretString            string                         `json:"secret-string,omitempty"`
	UUID                    string                         `json:"uuid,omitempty"`
	Vrid                    int                            `json:"vrid,omitempty"`
}

type FwRadiusServerAttribute struct {
	AttributeValue string `json:"attribute-value,omitempty"`
	CustomNumber   int    `json:"custom-number,omitempty"`
	CustomVendor   int    `json:"custom-vendor,omitempty"`
	Name           string `json:"name,omitempty"`
	Number         int    `json:"number,omitempty"`
	PrefixLength   string `json:"prefix-length,omitempty"`
	PrefixNumber   int    `json:"prefix-number,omitempty"`
	PrefixVendor   int    `json:"prefix-vendor,omitempty"`
	Value          string `json:"value,omitempty"`
	Vendor         int    `json:"vendor,omitempty"`
}

type FwRadiusServerRemote struct {
	IPListName []FwRadiusServerIPList `json:"ip-list,omitempty"`
}

type FwRadiusServerSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

type FwRadiusServerIPList struct {
	IPListEncrypted    string `json:"ip-list-encrypted,omitempty"`
	IPListName         string `json:"ip-list-name,omitempty"`
	IPListSecret       int    `json:"ip-list-secret,omitempty"`
	IPListSecretString string `json:"ip-list-secret-string,omitempty"`
}

func PostFwRadiusServer(id string, inst FwRadiusServer, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwRadiusServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/radius/server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwRadiusServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwRadiusServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwRadiusServer(id string, host string) (*FwRadiusServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwRadiusServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/radius/server/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwRadiusServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwRadiusServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
