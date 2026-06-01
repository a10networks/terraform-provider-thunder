package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type ScaleoutCluster struct {
	Inst struct {
		ClusterDevices ScaleoutClusterClusterDevices1449 `json:"cluster-devices"`

		ClusterId int `json:"cluster-id"`

		DbConfig ScaleoutClusterDbConfig1453 `json:"db-config"`

		DeviceGroups ScaleoutClusterDeviceGroups1454 `json:"device-groups"`

		LocalDevice ScaleoutClusterLocalDevice1457 `json:"local-device"`

		ServiceConfig ScaleoutClusterServiceConfig1484 `json:"service-config"`

		SlogLevel int `json:"slog-level" dval:"5"`

		TrackingTemplate ScaleoutClusterTrackingTemplate1486 `json:"tracking-template"`

		Uuid string `json:"uuid"`
	} `json:"cluster"`
}

type ScaleoutClusterClusterDevices1449 struct {
	Enable                  int                                                      `json:"enable"`
	Uuid                    string                                                   `json:"uuid"`
	MinimumNodes            ScaleoutClusterClusterDevicesMinimumNodes1450            `json:"minimum-nodes"`
	ClusterDiscoveryTimeout ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1451 `json:"cluster-discovery-timeout"`
	DeviceIdList            []ScaleoutClusterClusterDevicesDeviceIdList1452          `json:"device-id-list"`
}

type ScaleoutClusterClusterDevicesMinimumNodes1450 struct {
	MinimumNodesNum int    `json:"minimum-nodes-num"`
	Uuid            string `json:"uuid"`
}

type ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1451 struct {
	Uuid string `json:"uuid"`
}

type ScaleoutClusterClusterDevicesDeviceIdList1452 struct {
	Ip     string `json:"ip"`
	Action string `json:"action" dval:"enable"`
	Uuid   string `json:"uuid"`
}

type ScaleoutClusterDbConfig1453 struct {
	Ticktime            int    `json:"tickTime"`
	Initlimit           int    `json:"initLimit"`
	Synclimit           int    `json:"syncLimit"`
	Minsessiontimeout   int    `json:"minSessionTimeout" dval:"100"`
	Maxsessiontimeout   int    `json:"maxSessionTimeout" dval:"30000"`
	ClientRecvTimeout   int    `json:"client-recv-timeout" dval:"13000"`
	Clientport          int    `json:"clientPort"`
	LoopbackIntfSupport int    `json:"loopback-intf-support" dval:"1"`
	BrokenDetectTimeout int    `json:"broken-detect-timeout" dval:"12000"`
	MoreElectionPacket  int    `json:"more-election-packet" dval:"1"`
	ElectConnTimeout    int    `json:"elect-conn-timeout" dval:"1200"`
	Uuid                string `json:"uuid"`
}

type ScaleoutClusterDeviceGroups1454 struct {
	Enable          int                                              `json:"enable"`
	Uuid            string                                           `json:"uuid"`
	DeviceGroupList []ScaleoutClusterDeviceGroupsDeviceGroupList1455 `json:"device-group-list"`
}

type ScaleoutClusterDeviceGroupsDeviceGroupList1455 struct {
	DeviceGroup  int                                                          `json:"device-group"`
	DeviceIdList []ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1456 `json:"device-id-list"`
	Uuid         string                                                       `json:"uuid"`
	UserTag      string                                                       `json:"user-tag"`
}

type ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1456 struct {
	DeviceIdStart int `json:"device-id-start"`
	DeviceIdEnd   int `json:"device-id-end"`
}

type ScaleoutClusterLocalDevice1457 struct {
	Priority            int                                              `json:"priority"`
	Id1                 int                                              `json:"id1"`
	Action              string                                           `json:"action" dval:"enable"`
	FailureDomain       int                                              `json:"failure-domain"`
	FailureDomainString string                                           `json:"failure-domain-string"`
	StartDelay          int                                              `json:"start-delay"`
	ClusterMode         string                                           `json:"cluster-mode" dval:"layer-2"`
	Uuid                string                                           `json:"uuid"`
	L2Redirect          ScaleoutClusterLocalDeviceL2Redirect1458         `json:"l2-redirect"`
	TrafficRedirection  ScaleoutClusterLocalDeviceTrafficRedirection1459 `json:"traffic-redirection"`
	SessionSync         ScaleoutClusterLocalDeviceSessionSync1467        `json:"session-sync"`
	ExcludeInterfaces   ScaleoutClusterLocalDeviceExcludeInterfaces1474  `json:"exclude-interfaces"`
	TrackingTemplate    ScaleoutClusterLocalDeviceTrackingTemplate1479   `json:"tracking-template"`
}

type ScaleoutClusterLocalDeviceL2Redirect1458 struct {
	RedirectEth   int    `json:"redirect-eth"`
	EthernetVlan  int    `json:"ethernet-vlan"`
	RedirectTrunk int    `json:"redirect-trunk"`
	TrunkVlan     int    `json:"trunk-vlan"`
	Uuid          string `json:"uuid"`
}

type ScaleoutClusterLocalDeviceTrafficRedirection1459 struct {
	FollowShared        int                                                                 `json:"follow-shared"`
	Uuid                string                                                              `json:"uuid"`
	Interfaces          ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1460          `json:"interfaces"`
	ReachabilityOptions ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1465 `json:"reachability-options"`
	Encap               ScaleoutClusterLocalDeviceTrafficRedirectionEncap1466               `json:"encap"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1460 struct {
	EthCfg      []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1461      `json:"eth-cfg"`
	TrunkCfg    []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1462    `json:"trunk-cfg"`
	VeCfg       []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1463       `json:"ve-cfg"`
	LoopbackCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1464 `json:"loopback-cfg"`
	Uuid        string                                                                  `json:"uuid"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1461 struct {
	Ethernet int `json:"ethernet"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1462 struct {
	Trunk int `json:"trunk"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1463 struct {
	Ve int `json:"ve"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1464 struct {
	Loopback int `json:"loopback"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1465 struct {
	SkipDefaultRoute int    `json:"skip-default-route"`
	Uuid             string `json:"uuid"`
}

type ScaleoutClusterLocalDeviceTrafficRedirectionEncap1466 struct {
	Type       string `json:"type" dval:"vxlan"`
	UseV4Vxlan int    `json:"use-v4-vxlan"`
	Uuid       string `json:"uuid"`
}

type ScaleoutClusterLocalDeviceSessionSync1467 struct {
	FollowShared        int                                                          `json:"follow-shared"`
	Uuid                string                                                       `json:"uuid"`
	Interfaces          ScaleoutClusterLocalDeviceSessionSyncInterfaces1468          `json:"interfaces"`
	ReachabilityOptions ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1473 `json:"reachability-options"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfaces1468 struct {
	EthCfg      []ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1469      `json:"eth-cfg"`
	TrunkCfg    []ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1470    `json:"trunk-cfg"`
	VeCfg       []ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1471       `json:"ve-cfg"`
	LoopbackCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1472 `json:"loopback-cfg"`
	Uuid        string                                                           `json:"uuid"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1469 struct {
	Ethernet int `json:"ethernet"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1470 struct {
	Trunk int `json:"trunk"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1471 struct {
	Ve int `json:"ve"`
}

type ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1472 struct {
	Loopback int `json:"loopback"`
}

type ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1473 struct {
	SkipDefaultRoute int    `json:"skip-default-route"`
	Uuid             string `json:"uuid"`
}

type ScaleoutClusterLocalDeviceExcludeInterfaces1474 struct {
	EthCfg      []ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1475      `json:"eth-cfg"`
	TrunkCfg    []ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1476    `json:"trunk-cfg"`
	VeCfg       []ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1477       `json:"ve-cfg"`
	LoopbackCfg []ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1478 `json:"loopback-cfg"`
	Uuid        string                                                       `json:"uuid"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1475 struct {
	Ethernet int `json:"ethernet"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1476 struct {
	Trunk int `json:"trunk"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1477 struct {
	Ve int `json:"ve"`
}

type ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1478 struct {
	Loopback int `json:"loopback"`
}

type ScaleoutClusterLocalDeviceTrackingTemplate1479 struct {
	TemplateList      []ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1480      `json:"template-list"`
	MultiTemplateList []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1482 `json:"multi-template-list"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1480 struct {
	Template     string                                                                   `json:"template"`
	IpVersion    string                                                                   `json:"ip-version"`
	ThresholdCfg []ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1481 `json:"threshold-cfg"`
	Uuid         string                                                                   `json:"uuid"`
	UserTag      string                                                                   `json:"user-tag"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1481 struct {
	Threshold int    `json:"threshold"`
	Action    string `json:"action"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1482 struct {
	MultiTemplate string                                                                    `json:"multi-template"`
	Template      []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1483 `json:"template"`
	Threshold     int                                                                       `json:"threshold"`
	Action        string                                                                    `json:"action"`
	IpVersion     string                                                                    `json:"ip-version"`
	Uuid          string                                                                    `json:"uuid"`
	UserTag       string                                                                    `json:"user-tag"`
}

type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1483 struct {
	TemplateName  string `json:"template-name"`
	PartitionName string `json:"partition-name"`
}

type ScaleoutClusterServiceConfig1484 struct {
	DefaultUserGroupCount int                                            `json:"default-user-group-count"`
	Enable                int                                            `json:"enable"`
	Uuid                  string                                         `json:"uuid"`
	TemplateList          []ScaleoutClusterServiceConfigTemplateList1485 `json:"template-list"`
}

type ScaleoutClusterServiceConfigTemplateList1485 struct {
	Name           string `json:"name"`
	UserGroupCount int    `json:"user-group-count"`
	Uuid           string `json:"uuid"`
	UserTag        string `json:"user-tag"`
}

type ScaleoutClusterTrackingTemplate1486 struct {
	TemplateList []ScaleoutClusterTrackingTemplateTemplateList `json:"template-list"`
}

type ScaleoutClusterTrackingTemplateTemplateList struct {
	Template     string                                                    `json:"template"`
	ThresholdCfg []ScaleoutClusterTrackingTemplateTemplateListThresholdCfg `json:"threshold-cfg"`
	Uuid         string                                                    `json:"uuid"`
	UserTag      string                                                    `json:"user-tag"`
}

type ScaleoutClusterTrackingTemplateTemplateListThresholdCfg struct {
	Threshold int    `json:"threshold"`
	Action    string `json:"action"`
}

func (p *ScaleoutCluster) GetId() string {
	return strconv.Itoa(p.Inst.ClusterId)
}

func (p *ScaleoutCluster) getPath() string {
	return "scaleout/cluster"
}

func (p *ScaleoutCluster) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutCluster::Post")
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

func (p *ScaleoutCluster) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutCluster::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *ScaleoutCluster) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutCluster::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *ScaleoutCluster) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutCluster::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
