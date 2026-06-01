package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type ControllerProfileOper struct {
	Oper ControllerProfileOperOper `json:"oper"`
}
type DataControllerProfileOper struct {
	DtControllerProfileOper ControllerProfileOper `json:"profile"`
}

type ControllerProfileOperOper struct {
	OverallStatus                     string `json:"overall-status"`
	HeartbeatStatus                   string `json:"heartbeat-status"`
	HeartbeatErrorMessage             string `json:"heartbeat-error-message"`
	ServiceRegistry                   string `json:"service-registry"`
	ServiceRegistryErrorMessage       string `json:"service-registry-error-message"`
	RegistrationStatus                string `json:"registration-status"`
	RegistrationStatusCode            int    `json:"registration-status-code"`
	RegistrationErrorMessage          string `json:"registration-error-message"`
	DeregistrationStatus              string `json:"deregistration-status"`
	DeregistrationStatusCode          int    `json:"deregistration-status-code"`
	DeregistrationErrorMessage        string `json:"deregistration-error-message"`
	SchemaRegistryStatus              string `json:"schema-registry-status"`
	Broker_info                       string `json:"broker_info"`
	KafkaBrokerState                  string `json:"kafka-broker-state"`
	NumberOfOrgunitMappedPartitions   int    `json:"Number-of-orgunit-mapped-partitions"`
	NumberOfOrgunitUnmappedPartitions int    `json:"Number-of-orgunit-unmapped-partitions"`
	TunnelStatus                      string `json:"tunnel-status"`
	TunnelErrorMessage                string `json:"tunnel-error-message"`
	PeerDeviceInfo                    string `json:"peer-device-info"`
}

func (p *ControllerProfileOper) GetId() string {
	return "1"
}

func (p *ControllerProfileOper) getPath() string {
	return "controller/profile/oper"
}

func (p *ControllerProfileOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataControllerProfileOper, error) {
	logger.Println("ControllerProfileOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataControllerProfileOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
