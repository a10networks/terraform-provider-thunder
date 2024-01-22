

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutClusterLocalDevice struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    ClusterMode string `json:"cluster-mode" dval:"layer-2"`
    ExcludeInterfaces ScaleoutClusterLocalDeviceExcludeInterfaces1339 `json:"exclude-interfaces"`
    Id1 int `json:"id1"`
    L2Redirect ScaleoutClusterLocalDeviceL2Redirect1344 `json:"l2-redirect"`
    Priority int `json:"priority"`
    SessionSync ScaleoutClusterLocalDeviceSessionSync1345 `json:"session-sync"`
    StartDelay int `json:"start-delay"`
    TrackingTemplate ScaleoutClusterLocalDeviceTrackingTemplate1352 `json:"tracking-template"`
    TrafficRedirection ScaleoutClusterLocalDeviceTrafficRedirection1353 `json:"traffic-redirection"`
    Uuid string `json:"uuid"`
    ClusterId string 

	} `json:"local-device"`
}


type ScaleoutClusterLocalDeviceExcludeInterfaces1339 struct {
    EthCfg []ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1340 `json:"eth-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1341 `json:"trunk-cfg"`
    VeCfg []ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1342 `json:"ve-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1343 `json:"loopback-cfg"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1340 struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1341 struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1342 struct {
    Ve int `json:"ve"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1343 struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceL2Redirect1344 struct {
    RedirectEth int `json:"redirect-eth"`
    EthernetVlan int `json:"ethernet-vlan"`
    RedirectTrunk int `json:"redirect-trunk"`
    TrunkVlan int `json:"trunk-vlan"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceSessionSync1345 struct {
    FollowShared int `json:"follow-shared"`
    Uuid string `json:"uuid"`
    Interfaces ScaleoutClusterLocalDeviceSessionSyncInterfaces1346 `json:"interfaces"`
    ReachabilityOptions ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1351 `json:"reachability-options"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfaces1346 struct {
    EthCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1347 `json:"eth-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1348 `json:"trunk-cfg"`
    VeCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1349 `json:"ve-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1350 `json:"loopback-cfg"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1347 struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1348 struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1349 struct {
    Ve int `json:"ve"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1350 struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1351 struct {
    SkipDefaultRoute int `json:"skip-default-route"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceTrackingTemplate1352 struct {
    TemplateList []ScaleoutClusterLocalDeviceTrackingTemplateTemplateList `json:"template-list"`
    MultiTemplateList []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList `json:"multi-template-list"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateTemplateList struct {
    Template string `json:"template"`
    ThresholdCfg []ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg `json:"threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg struct {
    Threshold int `json:"threshold"`
    Action string `json:"action"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList struct {
    MultiTemplate string `json:"multi-template"`
    Template []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate `json:"template"`
    Threshold int `json:"threshold"`
    Action string `json:"action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate struct {
    TemplateName string `json:"template-name"`
    PartitionName string `json:"partition-name"`
}


type ScaleoutClusterLocalDeviceTrafficRedirection1353 struct {
    FollowShared int `json:"follow-shared"`
    Uuid string `json:"uuid"`
    Interfaces ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1354 `json:"interfaces"`
    ReachabilityOptions ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1359 `json:"reachability-options"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1354 struct {
    EthCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1355 `json:"eth-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1356 `json:"trunk-cfg"`
    VeCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1357 `json:"ve-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1358 `json:"loopback-cfg"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1355 struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1356 struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1357 struct {
    Ve int `json:"ve"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1358 struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1359 struct {
    SkipDefaultRoute int `json:"skip-default-route"`
    Uuid string `json:"uuid"`
}

func (p *ScaleoutClusterLocalDevice) GetId() string{
    return "1"
}

func (p *ScaleoutClusterLocalDevice) getPath() string{
    return "scaleout/cluster/" +p.Inst.ClusterId + "/local-device"
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
        if len(axResp) > 0{
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
