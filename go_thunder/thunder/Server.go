package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"net/url"
	"util"
)

type Server struct {
	ServerInstanceName ServerInstance `json:"server,omitempty"`
}

type ServerInstance struct {
	ServerInstanceAction                        string                          `json:"action,omitempty"`
	ServerInstanceAlternateServerAlternate      []ServerInstanceAlternateServer `json:"alternate-server,omitempty"`
	ServerInstanceConnLimit                     int                             `json:"conn-limit,omitempty"`
	ServerInstanceConnResume                    int                             `json:"conn-resume,omitempty"`
	ServerInstanceEthernet                      int                             `json:"ethernet,omitempty"`
	ServerInstanceExtendedStats                 int                             `json:"extended-stats,omitempty"`
	ServerInstanceExternalIP                    string                          `json:"external-ip,omitempty"`
	ServerInstanceFqdnName                      string                          `json:"fqdn-name,omitempty"`
	ServerInstanceHealthCheck                   string                          `json:"health-check,omitempty"`
	ServerInstanceHealthCheckDisable            int                             `json:"health-check-disable,omitempty"`
	ServerInstanceHealthCheckShared             string                          `json:"health-check-shared,omitempty"`
	ServerInstanceHost                          string                          `json:"host,omitempty"`
	ServerInstanceIpv6                          string                          `json:"ipv6,omitempty"`
	ServerInstanceL2HealthCheckPath             string                          `json:"l2-health-check-path,omitempty"`
	ServerInstanceName                          string                          `json:"name,omitempty"`
	ServerInstanceNoLogging                     int                             `json:"no-logging,omitempty"`
	ServerInstancePortListPortNumber            []ServerInstancePortList        `json:"port-list,omitempty"`
	ServerInstanceResolveAs                     string                          `json:"resolve-as,omitempty"`
	ServerInstanceSamplingEnableCounters1       []ServerInstanceSamplingEnable  `json:"sampling-enable,omitempty"`
	ServerInstanceServerIpv6Addr                string                          `json:"server-ipv6-addr,omitempty"`
	ServerInstanceSharedPartitionHealthCheck    int                             `json:"shared-partition-health-check,omitempty"`
	ServerInstanceSharedPartitionServerTemplate int                             `json:"shared-partition-server-template,omitempty"`
	ServerInstanceSlowStart                     int                             `json:"slow-start,omitempty"`
	ServerInstanceSpoofingCache                 int                             `json:"spoofing-cache,omitempty"`
	ServerInstanceStatsDataAction               string                          `json:"stats-data-action,omitempty"`
	ServerInstanceTemplateLinkCost              string                          `json:"template-link-cost,omitempty"`
	ServerInstanceTemplateServer                string                          `json:"template-server,omitempty"`
	ServerInstanceTemplateServerShared          string                          `json:"template-server-shared,omitempty"`
	ServerInstanceTrunk                         int                             `json:"trunk,omitempty"`
	ServerInstanceUUID                          string                          `json:"uuid,omitempty"`
	ServerInstanceUseAamServer                  int                             `json:"use-aam-server,omitempty"`
	ServerInstanceUserTag                       string                          `json:"user-tag,omitempty"`
	ServerInstanceWeight                        int                             `json:"weight,omitempty"`
}

type ServerInstanceAlternateServer struct {
	ServerInstanceAlternateServerAlternate     int    `json:"alternate,omitempty"`
	ServerInstanceAlternateServerAlternateName string `json:"alternate-name,omitempty"`
}

type ServerInstancePortList struct {
	ServerInstancePortListAction                      string                                 `json:"action,omitempty"`
	ServerInstancePortListAlternatePortAlternate      []ServerInstancePortListAlternatePort  `json:"alternate-port,omitempty"`
	ServerInstancePortListAuthCfgServicePrincipalName ServerInstancePortListAuthCfg          `json:"auth-cfg,omitempty"`
	ServerInstancePortListConnLimit                   int                                    `json:"conn-limit,omitempty"`
	ServerInstancePortListConnResume                  int                                    `json:"conn-resume,omitempty"`
	ServerInstancePortListExtendedStats               int                                    `json:"extended-stats,omitempty"`
	ServerInstancePortListFollowPortProtocol          string                                 `json:"follow-port-protocol,omitempty"`
	ServerInstancePortListHealthCheck                 string                                 `json:"health-check,omitempty"`
	ServerInstancePortListHealthCheckDisable          int                                    `json:"health-check-disable,omitempty"`
	ServerInstancePortListHealthCheckFollowPort       int                                    `json:"health-check-follow-port,omitempty"`
	ServerInstancePortListNoLogging                   int                                    `json:"no-logging,omitempty"`
	ServerInstancePortListNoSsl                       int                                    `json:"no-ssl,omitempty"`
	ServerInstancePortListPacketCaptureTemplate       string                                 `json:"packet-capture-template,omitempty"`
	ServerInstancePortListPortNumber                  int                                    `json:"port-number,omitempty"`
	ServerInstancePortListProtocol                    string                                 `json:"protocol,omitempty"`
	ServerInstancePortListRange                       int                                    `json:"range,omitempty"`
	ServerInstancePortListRportHealthCheckShared      string                                 `json:"rport-health-check-shared,omitempty"`
	ServerInstancePortListSamplingEnableCounters1     []ServerInstancePortListSamplingEnable `json:"sampling-enable,omitempty"`
	ServerInstancePortListSharedPartitionPortTemplate int                                    `json:"shared-partition-port-template,omitempty"`
	ServerInstancePortListSharedRportHealthCheck      int                                    `json:"shared-rport-health-check,omitempty"`
	ServerInstancePortListStatsDataAction             string                                 `json:"stats-data-action,omitempty"`
	ServerInstancePortListSupportHTTP2                int                                    `json:"support-http2,omitempty"`
	ServerInstancePortListTemplatePort                string                                 `json:"template-port,omitempty"`
	ServerInstancePortListTemplatePortShared          string                                 `json:"template-port-shared,omitempty"`
	ServerInstancePortListTemplateServerSsl           string                                 `json:"template-server-ssl,omitempty"`
	ServerInstancePortListUUID                        string                                 `json:"uuid,omitempty"`
	ServerInstancePortListUserTag                     string                                 `json:"user-tag,omitempty"`
	ServerInstancePortListWeight                      int                                    `json:"weight,omitempty"`
}

type ServerInstanceSamplingEnable struct {
	ServerInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type ServerInstancePortListAlternatePort struct {
	ServerInstancePortListAlternatePortAlternate           int    `json:"alternate,omitempty"`
	ServerInstancePortListAlternatePortAlternateName       string `json:"alternate-name,omitempty"`
	ServerInstancePortListAlternatePortAlternateServerPort int    `json:"alternate-server-port,omitempty"`
}

type ServerInstancePortListAuthCfg struct {
	ServerInstancePortListAuthCfgServicePrincipalName string `json:"service-principal-name,omitempty"`
}

type ServerInstancePortListSamplingEnable struct {
	ServerInstancePortListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

func PostServer(id string, inst Server, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetServer(id string, name1 string, host string) (*Server, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetServer")
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/server/"+nameEncode, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutServer(id string, name1 string, inst Server, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/server/"+nameEncode, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteServer(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteServer")
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/server/"+nameEncode, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Server
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
