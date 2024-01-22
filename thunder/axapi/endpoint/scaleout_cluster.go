

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutCluster struct {
	Inst struct {

    ClusterDevices ScaleoutClusterClusterDevices1360 `json:"cluster-devices"`
    ClusterId int `json:"cluster-id"`
    DbConfig ScaleoutClusterDbConfig1364 `json:"db-config"`
    DeviceGroups ScaleoutClusterDeviceGroups1365 `json:"device-groups"`
    LocalDevice ScaleoutClusterLocalDevice1368 `json:"local-device"`
    ServiceConfig ScaleoutClusterServiceConfig1394 `json:"service-config"`
    SlogLevel int `json:"slog-level" dval:"5"`
    TrackingTemplate ScaleoutClusterTrackingTemplate1396 `json:"tracking-template"`
    Uuid string `json:"uuid"`

	} `json:"cluster"`
}


type ScaleoutClusterClusterDevices1360 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
    MinimumNodes ScaleoutClusterClusterDevicesMinimumNodes1361 `json:"minimum-nodes"`
    ClusterDiscoveryTimeout ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1362 `json:"cluster-discovery-timeout"`
    DeviceIdList []ScaleoutClusterClusterDevicesDeviceIdList1363 `json:"device-id-list"`
}


type ScaleoutClusterClusterDevicesMinimumNodes1361 struct {
    MinimumNodesNum int `json:"minimum-nodes-num"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1362 struct {
    Uuid string `json:"uuid"`
}


type ScaleoutClusterClusterDevicesDeviceIdList1363 struct {
    Ip string `json:"ip"`
    Action string `json:"action" dval:"enable"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterDbConfig1364 struct {
    Ticktime int `json:"tickTime"`
    Initlimit int `json:"initLimit"`
    Synclimit int `json:"syncLimit"`
    Minsessiontimeout int `json:"minSessionTimeout" dval:"100"`
    Maxsessiontimeout int `json:"maxSessionTimeout" dval:"30000"`
    ClientRecvTimeout int `json:"client-recv-timeout" dval:"13000"`
    Clientport int `json:"clientPort"`
    LoopbackIntfSupport int `json:"loopback-intf-support" dval:"1"`
    BrokenDetectTimeout int `json:"broken-detect-timeout" dval:"12000"`
    MoreElectionPacket int `json:"more-election-packet" dval:"1"`
    ElectConnTimeout int `json:"elect-conn-timeout" dval:"1200"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterDeviceGroups1365 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
    DeviceGroupList []ScaleoutClusterDeviceGroupsDeviceGroupList1366 `json:"device-group-list"`
}


type ScaleoutClusterDeviceGroupsDeviceGroupList1366 struct {
    DeviceGroup int `json:"device-group"`
    DeviceIdList []ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1367 `json:"device-id-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1367 struct {
    DeviceIdStart int `json:"device-id-start"`
    DeviceIdEnd int `json:"device-id-end"`
}


type ScaleoutClusterLocalDevice1368 struct {
    Priority int `json:"priority"`
    Id1 int `json:"id1"`
    Action string `json:"action" dval:"enable"`
    StartDelay int `json:"start-delay"`
    ClusterMode string `json:"cluster-mode" dval:"layer-2"`
    Uuid string `json:"uuid"`
    L2Redirect ScaleoutClusterLocalDeviceL2Redirect1369 `json:"l2-redirect"`
    TrafficRedirection ScaleoutClusterLocalDeviceTrafficRedirection1370 `json:"traffic-redirection"`
    SessionSync ScaleoutClusterLocalDeviceSessionSync1377 `json:"session-sync"`
    ExcludeInterfaces ScaleoutClusterLocalDeviceExcludeInterfaces1384 `json:"exclude-interfaces"`
    TrackingTemplate ScaleoutClusterLocalDeviceTrackingTemplate1389 `json:"tracking-template"`
}


type ScaleoutClusterLocalDeviceL2Redirect1369 struct {
    RedirectEth int `json:"redirect-eth"`
    EthernetVlan int `json:"ethernet-vlan"`
    RedirectTrunk int `json:"redirect-trunk"`
    TrunkVlan int `json:"trunk-vlan"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceTrafficRedirection1370 struct {
    FollowShared int `json:"follow-shared"`
    Uuid string `json:"uuid"`
    Interfaces ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1371 `json:"interfaces"`
    ReachabilityOptions ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1376 `json:"reachability-options"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1371 struct {
    EthCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1372 `json:"eth-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1373 `json:"trunk-cfg"`
    VeCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1374 `json:"ve-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1375 `json:"loopback-cfg"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1372 struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1373 struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1374 struct {
    Ve int `json:"ve"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1375 struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1376 struct {
    SkipDefaultRoute int `json:"skip-default-route"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceSessionSync1377 struct {
    FollowShared int `json:"follow-shared"`
    Uuid string `json:"uuid"`
    Interfaces ScaleoutClusterLocalDeviceSessionSyncInterfaces1378 `json:"interfaces"`
    ReachabilityOptions ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1383 `json:"reachability-options"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfaces1378 struct {
    EthCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1379 `json:"eth-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1380 `json:"trunk-cfg"`
    VeCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1381 `json:"ve-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1382 `json:"loopback-cfg"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1379 struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1380 struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1381 struct {
    Ve int `json:"ve"`
}


type ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1382 struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1383 struct {
    SkipDefaultRoute int `json:"skip-default-route"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceExcludeInterfaces1384 struct {
    EthCfg []ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1385 `json:"eth-cfg"`
    TrunkCfg []ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1386 `json:"trunk-cfg"`
    VeCfg []ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1387 `json:"ve-cfg"`
    LoopbackCfg []ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1388 `json:"loopback-cfg"`
    Uuid string `json:"uuid"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1385 struct {
    Ethernet int `json:"ethernet"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1386 struct {
    Trunk int `json:"trunk"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1387 struct {
    Ve int `json:"ve"`
}


type ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1388 struct {
    Loopback int `json:"loopback"`
}


type ScaleoutClusterLocalDeviceTrackingTemplate1389 struct {
    TemplateList []ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1390 `json:"template-list"`
    MultiTemplateList []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1392 `json:"multi-template-list"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1390 struct {
    Template string `json:"template"`
    ThresholdCfg []ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1391 `json:"threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1391 struct {
    Threshold int `json:"threshold"`
    Action string `json:"action"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1392 struct {
    MultiTemplate string `json:"multi-template"`
    Template []ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1393 `json:"template"`
    Threshold int `json:"threshold"`
    Action string `json:"action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1393 struct {
    TemplateName string `json:"template-name"`
    PartitionName string `json:"partition-name"`
}


type ScaleoutClusterServiceConfig1394 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
    TemplateList []ScaleoutClusterServiceConfigTemplateList1395 `json:"template-list"`
}


type ScaleoutClusterServiceConfigTemplateList1395 struct {
    Name string `json:"name"`
    BucketCount int `json:"bucket-count" dval:"256"`
    DeviceGroup int `json:"device-group"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterTrackingTemplate1396 struct {
    TemplateList []ScaleoutClusterTrackingTemplateTemplateList `json:"template-list"`
}


type ScaleoutClusterTrackingTemplateTemplateList struct {
    Template string `json:"template"`
    ThresholdCfg []ScaleoutClusterTrackingTemplateTemplateListThresholdCfg `json:"threshold-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type ScaleoutClusterTrackingTemplateTemplateListThresholdCfg struct {
    Threshold int `json:"threshold"`
    Action string `json:"action"`
}

func (p *ScaleoutCluster) GetId() string{
    return strconv.Itoa(p.Inst.ClusterId)
}

func (p *ScaleoutCluster) getPath() string{
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
        if len(axResp) > 0{
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
