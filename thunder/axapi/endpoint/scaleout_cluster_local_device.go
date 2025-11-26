package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ScaleoutClusterLocalDevice struct {
	Inst struct {
		Action string `json:"action" dval:"enable"`

		ClusterMode string `json:"cluster-mode" dval:"layer-2"`

		ExcludeInterfaces ScaleoutClusterLocalDeviceExcludeInterfaces1427 `json:"exclude-interfaces"`

		Id1 int `json:"id"`

		L2Redirect ScaleoutClusterLocalDeviceL2Redirect1432 `json:"l2-redirect"`

		Priority int `json:"priority"`

		SessionSync ScaleoutClusterLocalDeviceSessionSync1433 `json:"session-sync"`

		StartDelay int `json:"start-delay"`

		TrackingTemplate ScaleoutClusterLocalDeviceTrackingTemplate1440 `json:"tracking-template"`

		TrafficRedirection ScaleoutClusterLocalDeviceTrafficRedirection1441 `json:"traffic-redirection"`

		Uuid string `json:"uuid"`

		ClusterId string
	} `json:"local-device"`
}

type ScaleoutClusterLocalDeviceExcludeInterfaces1427 struct {
	EthCfg      []ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1428      `json:"eth-cfg"`
	TrunkCfg    []ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1429    `json:"trunk-cfg"`
	VeCfg       []ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1430       `json:"ve-cfg"`
	LoopbackCfg []ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1431 `json:"loopback-cfg"`
	Uuid        string                                                       `json:"uuid"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1428 struct {
	Ethernet int `json:"ethernet"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1429 struct {
	Trunk int `json:"trunk"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1430 struct {
	Ve int `json:"ve"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1431 struct {
	Loopback int `json:"loopback"`
}

type ScaleoutClusterLocalDeviceL2Redirect1432 struct {
	RedirectEth   int    `json:"redirect-eth"`
	EthernetVlan  int    `json:"ethernet-vlan"`
	RedirectTrunk int    `json:"redirect-trunk"`
	TrunkVlan     int    `json:"trunk-vlan"`
	Uuid          string `json:"uuid"`
}

type ScaleoutClusterLocalDeviceSessionSync1433 struct {
	FollowShared        int                                                          `json:"follow-shared"`
	Uuid                string                                                       `json:"uuid"`
	Interfaces          ScaleoutClusterLocalDeviceSessionSyncInterfaces1434          `json:"interfaces"`
	ReachabilityOptions ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1439 `json:"reachability-options"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfaces1434 struct {
	EthCfg      []ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1435      `json:"eth-cfg"`
	TrunkCfg    []ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1436    `json:"trunk-cfg"`
	VeCfg       []ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1437       `json:"ve-cfg"`
	LoopbackCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1438 `json:"loopback-cfg"`
	Uuid        string                                                           `json:"uuid"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1435 struct {
	Ethernet int `json:"ethernet"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1436 struct {
	Trunk int `json:"trunk"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1437 struct {
	Ve int `json:"ve"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1438 struct {
	Loopback int `json:"loopback"`
}

type ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1439 struct {
	SkipDefaultRoute int    `json:"skip-default-route"`
	Uuid             string `json:"uuid"`
}

type ScaleoutClusterLocalDeviceTrackingTemplate1440 struct {
	TemplateList      []ScaleoutClusterLocalDeviceTrackingTemplateTemplateList      `json:"template-list"`
	MultiTemplateList []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList `json:"multi-template-list"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateTemplateList struct {
	Template     string                                                               `json:"template"`
	IpVersion    string                                                               `json:"ip-version"`
	ThresholdCfg []ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg `json:"threshold-cfg"`
	Uuid         string                                                               `json:"uuid"`
	UserTag      string                                                               `json:"user-tag"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg struct {
	Threshold int    `json:"threshold"`
	Action    string `json:"action"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList struct {
	MultiTemplate string                                                                `json:"multi-template"`
	Template      []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate `json:"template"`
	Threshold     int                                                                   `json:"threshold"`
	Action        string                                                                `json:"action"`
	IpVersion     string                                                                `json:"ip-version"`
	Uuid          string                                                                `json:"uuid"`
	UserTag       string                                                                `json:"user-tag"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate struct {
	TemplateName  string `json:"template-name"`
	PartitionName string `json:"partition-name"`
}

type ScaleoutClusterLocalDeviceTrafficRedirection1441 struct {
	FollowShared        int                                                                 `json:"follow-shared"`
	Uuid                string                                                              `json:"uuid"`
	Interfaces          ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1442          `json:"interfaces"`
	ReachabilityOptions ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1447 `json:"reachability-options"`
	Encap               ScaleoutClusterLocalDeviceTrafficRedirectionEncap1448               `json:"encap"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1442 struct {
	EthCfg      []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1443      `json:"eth-cfg"`
	TrunkCfg    []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1444    `json:"trunk-cfg"`
	VeCfg       []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1445       `json:"ve-cfg"`
	LoopbackCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1446 `json:"loopback-cfg"`
	Uuid        string                                                                  `json:"uuid"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1443 struct {
	Ethernet int `json:"ethernet"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1444 struct {
	Trunk int `json:"trunk"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1445 struct {
	Ve int `json:"ve"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1446 struct {
	Loopback int `json:"loopback"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1447 struct {
	SkipDefaultRoute int    `json:"skip-default-route"`
	Uuid             string `json:"uuid"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionEncap1448 struct {
	Type       string `json:"type" dval:"vxlan"`
	UseV4Vxlan int    `json:"use-v4-vxlan"`
	Uuid       string `json:"uuid"`
}

func (p *ScaleoutClusterLocalDevice) GetId() string {
	return "1"
}

func (p *ScaleoutClusterLocalDevice) getPath() string {
	return "scaleout/cluster/" + p.Inst.ClusterId + "/local-device"
}

func (p *ScaleoutClusterLocalDevice) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDevice::Post")
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

func (p *ScaleoutClusterLocalDevice) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDevice::Get")
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
func (p *ScaleoutClusterLocalDevice) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDevice::Put")
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

func (p *ScaleoutClusterLocalDevice) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutClusterLocalDevice::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
