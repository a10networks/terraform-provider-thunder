package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SlbServerPort struct {
	HealthCheckDisable SlbServerPortInstance `json:"port,omitempty"`
}

type SlbServerPortInstance struct {
	Action                      string                    `json:"action,omitempty"`
	AlternateName               []SlbServerAlternatePort  `json:"alternate-port,omitempty"`
	ServicePrincipalName        SlbServerAuthCfg          `json:"auth-cfg,omitempty"`
	ConnLimit                   int                       `json:"conn-limit,omitempty"`
	ConnResume                  int                       `json:"conn-resume,omitempty"`
	ExtendedStats               int                       `json:"extended-stats,omitempty"`
	FollowPortProtocol          string                    `json:"follow-port-protocol,omitempty"`
	HealthCheck                 string                    `json:"health-check,omitempty"`
	HealthCheckDisable          int                       `json:"health-check-disable,omitempty"`
	HealthCheckFollowPort       int                       `json:"health-check-follow-port,omitempty"`
	NoLogging                   int                       `json:"no-logging,omitempty"`
	NoSsl                       int                       `json:"no-ssl,omitempty"`
	PortNumber                  int                       `json:"port-number,omitempty"`
	Protocol                    string                    `json:"protocol,omitempty"`
	Range                       int                       `json:"range,omitempty"`
	RportHealthCheckShared      string                    `json:"rport-health-check-shared,omitempty"`
	Counters1                   []SlbServerSamplingEnable `json:"sampling-enable,omitempty"`
	SharedPartitionPortTemplate int                       `json:"shared-partition-port-template,omitempty"`
	SharedRportHealthCheck      int                       `json:"shared-rport-health-check,omitempty"`
	StatsDataAction             string                    `json:"stats-data-action,omitempty"`
	SupportHTTP2                int                       `json:"support-http2,omitempty"`
	TemplatePort                string                    `json:"template-port,omitempty"`
	TemplatePortShared          string                    `json:"template-port-shared,omitempty"`
	TemplateServerSsl           string                    `json:"template-server-ssl,omitempty"`
	UUID                        string                    `json:"uuid,omitempty"`
	UserTag                     string                    `json:"user-tag,omitempty"`
	Weight                      int                       `json:"weight,omitempty"`
}

type SlbServerAlternatePort struct {
	Alternate           int    `json:"alternate,omitempty"`
	AlternateName       string `json:"alternate-name,omitempty"`
	AlternateServerPort int    `json:"alternate-server-port,omitempty"`
}

type SlbServerAuthCfg struct {
	ServicePrincipalName string `json:"service-principal-name,omitempty"`
}

type SlbServerSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostSlbServerPort(id string, name string, inst SlbServerPort, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbServerPort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/server/"+name+"/port", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbServerPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbServerPort REQ RES..........................", m)
			check_api_status("PostSlbServerPort", data)

		}
	}

}

func GetSlbServerPort(id string, name1 string, name2 string, name3 string, host string) (*SlbServerPort, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbServerPort")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/server/"+name1+"/port/"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbServerPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbServerPort REQ RES..........................", m)
			check_api_status("GetSlbServerPort", data)
			return &m, nil
		}
	}

}

func PutSlbServerPort(id string, name1 string, name2 string, name3 string, inst SlbServerPort, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbServerPort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/server/"+name1+"/port/"+name2+"+"+name3, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbServerPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbServerPort REQ RES..........................", m)
			check_api_status("PutSlbServerPort", data)

		}
	}

}

func DeleteSlbServerPort(id string, name1 string, name2 string, name3 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbServerPort")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/server/"+name1+"/port/"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbServerPort
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
