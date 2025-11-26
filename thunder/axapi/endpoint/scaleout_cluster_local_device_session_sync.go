package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ScaleoutClusterLocalDeviceSessionSync struct {
	Inst struct {
		FollowShared int `json:"follow-shared"`

		Interfaces ScaleoutClusterLocalDeviceSessionSyncInterfaces1414 `json:"interfaces"`

		ReachabilityOptions ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1419 `json:"reachability-options"`

		Uuid string `json:"uuid"`

		ClusterId string
	} `json:"session-sync"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfaces1414 struct {
	EthCfg      []ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1415      `json:"eth-cfg"`
	TrunkCfg    []ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1416    `json:"trunk-cfg"`
	VeCfg       []ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1417       `json:"ve-cfg"`
	LoopbackCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1418 `json:"loopback-cfg"`
	Uuid        string                                                           `json:"uuid"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1415 struct {
	Ethernet int `json:"ethernet"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1416 struct {
	Trunk int `json:"trunk"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1417 struct {
	Ve int `json:"ve"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1418 struct {
	Loopback int `json:"loopback"`
}

type ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1419 struct {
	SkipDefaultRoute int    `json:"skip-default-route"`
	Uuid             string `json:"uuid"`
}

func (p *ScaleoutClusterLocalDeviceSessionSync) GetId() string {
	return "1"
}

func (p *ScaleoutClusterLocalDeviceSessionSync) getPath() string {
	return "scaleout/cluster/" + p.Inst.ClusterId + "/local-device/session-sync"
}

func (p *ScaleoutClusterLocalDeviceSessionSync) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDeviceSessionSync::Post")
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

func (p *ScaleoutClusterLocalDeviceSessionSync) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDeviceSessionSync::Get")
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
func (p *ScaleoutClusterLocalDeviceSessionSync) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDeviceSessionSync::Put")
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

func (p *ScaleoutClusterLocalDeviceSessionSync) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDeviceSessionSync::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
