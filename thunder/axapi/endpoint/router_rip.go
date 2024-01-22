

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterRip struct {
	Inst struct {

    CiscoMetricBehavior string `json:"cisco-metric-behavior" dval:"disable"`
    DefaultInformation string `json:"default-information"`
    DefaultMetric int `json:"default-metric" dval:"1"`
    DistanceListCfg []RouterRipDistanceListCfg `json:"distance-list-cfg"`
    DistributeList RouterRipDistributeList1306 `json:"distribute-list"`
    Neighbor []RouterRipNeighbor `json:"neighbor"`
    NetworkAddresses []RouterRipNetworkAddresses `json:"network-addresses"`
    NetworkInterfaceListCfg []RouterRipNetworkInterfaceListCfg `json:"network-interface-list-cfg"`
    OffsetList RouterRipOffsetList1310 `json:"offset-list"`
    PassiveInterfaceList []RouterRipPassiveInterfaceList `json:"passive-interface-list"`
    RecvBufferSize int `json:"recv-buffer-size"`
    Redistribute RouterRipRedistribute1312 `json:"redistribute"`
    RipMaximumPrefixCfg RouterRipRipMaximumPrefixCfg `json:"rip-maximum-prefix-cfg"`
    RouteCfg []RouterRipRouteCfg `json:"route-cfg"`
    Timers RouterRipTimers `json:"timers"`
    Uuid string `json:"uuid"`
    Version int `json:"version" dval:"2"`

	} `json:"rip"`
}


type RouterRipDistanceListCfg struct {
    Distance int `json:"distance" dval:"120"`
    DistanceIpv4Mask string `json:"distance-ipv4-mask"`
    DistanceAcl string `json:"distance-acl"`
}


type RouterRipDistributeList1306 struct {
    AclCfg []RouterRipDistributeListAclCfg1307 `json:"acl-cfg"`
    Uuid string `json:"uuid"`
    Prefix RouterRipDistributeListPrefix1308 `json:"prefix"`
}


type RouterRipDistributeListAclCfg1307 struct {
    Acl string `json:"acl"`
    AclDirection string `json:"acl-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterRipDistributeListPrefix1308 struct {
    PrefixCfg []RouterRipDistributeListPrefixPrefixCfg1309 `json:"prefix-cfg"`
    Uuid string `json:"uuid"`
}


type RouterRipDistributeListPrefixPrefixCfg1309 struct {
    PrefixList string `json:"prefix-list"`
    PrefixListDirection string `json:"prefix-list-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterRipNeighbor struct {
    Value string `json:"value"`
}


type RouterRipNetworkAddresses struct {
    NetworkIpv4Mask string `json:"network-ipv4-mask"`
}


type RouterRipNetworkInterfaceListCfg struct {
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterRipOffsetList1310 struct {
    AclCfg []RouterRipOffsetListAclCfg1311 `json:"acl-cfg"`
    Uuid string `json:"uuid"`
}


type RouterRipOffsetListAclCfg1311 struct {
    Acl string `json:"acl"`
    OffsetListDirection string `json:"offset-list-direction"`
    Metric int `json:"metric"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterRipPassiveInterfaceList struct {
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterRipRedistribute1312 struct {
    RedistList []RouterRipRedistributeRedistList1313 `json:"redist-list"`
    VipList []RouterRipRedistributeVipList1314 `json:"vip-list"`
    Uuid string `json:"uuid"`
}


type RouterRipRedistributeRedistList1313 struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    RouteMap string `json:"route-map"`
}


type RouterRipRedistributeVipList1314 struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
}


type RouterRipRipMaximumPrefixCfg struct {
    MaximumPrefix int `json:"maximum-prefix"`
    MaximumPrefixThres int `json:"maximum-prefix-thres" dval:"75"`
}


type RouterRipRouteCfg struct {
    Route string `json:"route"`
}


type RouterRipTimers struct {
    TimersCfg RouterRipTimersTimersCfg `json:"timers-cfg"`
}


type RouterRipTimersTimersCfg struct {
    Basic int `json:"basic"`
    Val2 int `json:"val-2"`
    Val3 int `json:"val-3"`
}

func (p *RouterRip) GetId() string{
    return "1"
}

func (p *RouterRip) getPath() string{
    return "router/rip"
}

func (p *RouterRip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRip::Post")
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

func (p *RouterRip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRip::Get")
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
func (p *RouterRip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRip::Put")
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

func (p *RouterRip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterRip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
