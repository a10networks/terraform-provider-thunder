package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type AlternatePort struct {
	AlternateName       string `json:"alternate-name,omitempty"`
	Alternate           int    `json:"alternate,omitempty"`
	AlternateServerPort int    `json:"alternate-server-port,omitempty"`
}

type AuthCfgS struct {
	ServicePrincipalName string `json:"service-principal-name,omitempty"`
}

type PortLists struct {
	HealthCheckDisable     int              `json:"health-check-disable,omitempty"`
	Protocol               string           `json:"protocol,omitempty"`
	Weight                 int              `json:"weight,omitempty"`
	SharedRportHealthCheck int              `json:"shared-rport-health-check,omitempty"`
	StatsDataAction        string           `json:"stats-data-action,omitempty"`
	HealthCheckFollowPort  int              `json:"health-check-follow-port,omitempty"`
	TemplatePort           string           `json:"template-port,omitempty"`
	ConnLimit              int              `json:"conn-limit,omitempty"`
	UUID                   string           `json:"uuid,omitempty"`
	SupportHTTP2           int              `json:"support-http2,omitempty"`
	Counters1              []SamplingEnable `json:"sampling-enable,omitempty"`
	NoSsl                  int              `json:"no-ssl,omitempty"`
	FollowPortProtocol     string           `json:"follow-port-protocol,omitempty"`
	TemplateServerSsl      string           `json:"template-server-ssl,omitempty"`
	AlternateName          []AlternatePort  `json:"alternate-port,omitempty"`
	PortNumber             int              `json:"port-number,omitempty"`
	ExtendedStats          int              `json:"extended-stats,omitempty"`
	RportHealthCheckShared string           `json:"rport-health-check-shared,omitempty"`
	ConnResume             int              `json:"conn-resume,omitempty"`
	UserTag                string           `json:"user-tag,omitempty"`
	Range                  int              `json:"range,omitempty"`
	//ServicePrincipalName   AuthCfgS         `json:"auth-cfg,omitempty"`
	Action      string `json:"action,omitempty"`
	HealthCheck string `json:"health-check,omitempty"`
	NoLogging   int    `json:"no-logging,omitempty"`
}

type AlternateServer struct {
	AlternateName string `json:"alternate-name,omitempty"`
	Alternate     int    `json:"alternate,omitempty"`
}

type ServerInstance struct {
	HealthCheckDisable         int               `json:"health-check-disable,omitempty"`
	PortNumber                 []PortLists       `json:"port-list,omitempty"`
	StatsDataAction            string            `json:"stats-data-action,omitempty"`
	SlowStart                  int               `json:"slow-start,omitempty"`
	Weight                     int               `json:"weight,omitempty"`
	SpoofingCache              int               `json:"spoofing-cache,omitempty"`
	ResolveAs                  string            `json:"resolve-as,omitempty"`
	ConnLimit                  int               `json:"conn-limit,omitempty"`
	UUID                       string            `json:"uuid,omitempty"`
	FqdnName                   string            `json:"fqdn-name,omitempty"`
	ExternalIP                 string            `json:"external-ip,omitempty"`
	HealthCheckShared          string            `json:"health-check-shared,omitempty"`
	Ipv6                       string            `json:"ipv6,omitempty"`
	TemplateServer             string            `json:"template-server,omitempty"`
	ServerIpv6Addr             string            `json:"server-ipv6-addr,omitempty"`
	AlternateName              []AlternateServer `json:"alternate-server,omitempty"`
	SharedPartitionHealthCheck int               `json:"shared-partition-health-check,omitempty"`
	Host                       string            `json:"host,omitempty"`
	ExtendedStats              int               `json:"extended-stats,omitempty"`
	ConnResume                 int               `json:"conn-resume,omitempty"`
	Name                       string            `json:"name,omitempty"`
	UserTag                    string            `json:"user-tag,omitempty"`
	Counters1                  []SamplingEnable  `json:"sampling-enable,omitempty"`
	Action                     string            `json:"action,omitempty"`
	HealthCheck                string            `json:"health-check,omitempty"`
	NoLogging                  int               `json:"no-logging,omitempty"`
}

type Server struct {
	Name ServerInstance `json:"server,omitempty"`
}

func GetServer(id string, name string, host string) (*Server, error) {

	logger := util.GetLoggerInstance()
	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/server/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostServer(id string, sg Server, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("POSTServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func PutServer(id string, name string, sg Server, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	logger.Println("[INFO] received V  --" + sg.Name.Name)

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/server/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PUTServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteServer(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/server/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
