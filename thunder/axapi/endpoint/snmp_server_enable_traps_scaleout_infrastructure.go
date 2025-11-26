package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SnmpServerEnableTrapsScaleoutInfrastructure struct {
	Inst struct {
		All int `json:"all"`

		Cluster SnmpServerEnableTrapsScaleoutInfrastructureCluster1581 `json:"cluster"`

		MasterNode SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1582 `json:"master-node"`

		ServiceNode SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1583 `json:"service-node"`

		TestSendAllTraps int `json:"test-send-all-traps"`

		Uuid string `json:"uuid"`
	} `json:"infrastructure"`
}

type SnmpServerEnableTrapsScaleoutInfrastructureCluster1581 struct {
	Election                int    `json:"election"`
	MasterCallingReElection int    `json:"master-calling-re-election"`
	NodeStatus              int    `json:"node-status"`
	Uuid                    string `json:"uuid"`
}

type SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1582 struct {
	TrafficMapDistribution  int    `json:"traffic-map-distribution"`
	VserverTrafficMapUpdate int    `json:"vserver-traffic-map-update"`
	Uuid                    string `json:"uuid"`
}

type SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1583 struct {
	LocalDeviceDisabled int    `json:"local-device-disabled"`
	ServiceMaster       int    `json:"service-master"`
	TrafficMapUpdate    int    `json:"traffic-map-update"`
	Uuid                string `json:"uuid"`
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructure) GetId() string {
	return "1"
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructure) getPath() string {
	return "snmp-server/enable/traps/scaleout/infrastructure"
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructure) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerEnableTrapsScaleoutInfrastructure::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructure) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerEnableTrapsScaleoutInfrastructure::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *SnmpServerEnableTrapsScaleoutInfrastructure) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerEnableTrapsScaleoutInfrastructure::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructure) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SnmpServerEnableTrapsScaleoutInfrastructure::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
