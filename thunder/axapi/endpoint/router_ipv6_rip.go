

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6Rip struct {
	Inst struct {

    AggregateAddressCfg []RouterIpv6RipAggregateAddressCfg `json:"aggregate-address-cfg"`
    CiscoMetricBehavior string `json:"cisco-metric-behavior" dval:"disable"`
    DefaultInformation string `json:"default-information"`
    DefaultMetric int `json:"default-metric" dval:"1"`
    DistributeList RouterIpv6RipDistributeList1257 `json:"distribute-list"`
    OffsetList RouterIpv6RipOffsetList1261 `json:"offset-list"`
    PassiveInterfaceList []RouterIpv6RipPassiveInterfaceList `json:"passive-interface-list"`
    RecvBufferSize int `json:"recv-buffer-size"`
    Redistribute RouterIpv6RipRedistribute1263 `json:"redistribute"`
    RipngNeighbor RouterIpv6RipRipngNeighbor `json:"ripng-neighbor"`
    RouteCfg []RouterIpv6RipRouteCfg `json:"route-cfg"`
    RouteMap RouterIpv6RipRouteMap1266 `json:"route-map"`
    Timers RouterIpv6RipTimers `json:"timers"`
    Uuid string `json:"uuid"`

	} `json:"rip"`
}


type RouterIpv6RipAggregateAddressCfg struct {
    AggregateAddress string `json:"aggregate-address"`
}


type RouterIpv6RipDistributeList1257 struct {
    AclCfg []RouterIpv6RipDistributeListAclCfg1258 `json:"acl-cfg"`
    Uuid string `json:"uuid"`
    Prefix RouterIpv6RipDistributeListPrefix1259 `json:"prefix"`
}


type RouterIpv6RipDistributeListAclCfg1258 struct {
    Acl string `json:"acl"`
    AclDirection string `json:"acl-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterIpv6RipDistributeListPrefix1259 struct {
    PrefixCfg []RouterIpv6RipDistributeListPrefixPrefixCfg1260 `json:"prefix-cfg"`
    Uuid string `json:"uuid"`
}


type RouterIpv6RipDistributeListPrefixPrefixCfg1260 struct {
    PrefixList string `json:"prefix-list"`
    PrefixListDirection string `json:"prefix-list-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterIpv6RipOffsetList1261 struct {
    AclCfg []RouterIpv6RipOffsetListAclCfg1262 `json:"acl-cfg"`
    Uuid string `json:"uuid"`
}


type RouterIpv6RipOffsetListAclCfg1262 struct {
    Acl string `json:"acl"`
    OffsetListDirection string `json:"offset-list-direction"`
    Metric int `json:"metric"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterIpv6RipPassiveInterfaceList struct {
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterIpv6RipRedistribute1263 struct {
    RedistList []RouterIpv6RipRedistributeRedistList1264 `json:"redist-list"`
    VipList []RouterIpv6RipRedistributeVipList1265 `json:"vip-list"`
    Uuid string `json:"uuid"`
}


type RouterIpv6RipRedistributeRedistList1264 struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    RouteMap string `json:"route-map"`
}


type RouterIpv6RipRedistributeVipList1265 struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
}


type RouterIpv6RipRipngNeighbor struct {
    RipngNeighborCfg []RouterIpv6RipRipngNeighborRipngNeighborCfg `json:"ripng-neighbor-cfg"`
}


type RouterIpv6RipRipngNeighborRipngNeighborCfg struct {
    NeighborLinkLocalAddr string `json:"neighbor-link-local-addr"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterIpv6RipRouteCfg struct {
    Route string `json:"route"`
}


type RouterIpv6RipRouteMap1266 struct {
    MapCfg []RouterIpv6RipRouteMapMapCfg1267 `json:"map-cfg"`
    Uuid string `json:"uuid"`
}


type RouterIpv6RipRouteMapMapCfg1267 struct {
    Map string `json:"map"`
    RouteMapDirection string `json:"route-map-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}


type RouterIpv6RipTimers struct {
    TimersCfg RouterIpv6RipTimersTimersCfg `json:"timers-cfg"`
}


type RouterIpv6RipTimersTimersCfg struct {
    Basic int `json:"basic"`
    Val2 int `json:"val-2"`
    Val3 int `json:"val-3"`
}

func (p *RouterIpv6Rip) GetId() string{
    return "1"
}

func (p *RouterIpv6Rip) getPath() string{
    return "router/ipv6/rip"
}

func (p *RouterIpv6Rip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Rip::Post")
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

func (p *RouterIpv6Rip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Rip::Get")
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
func (p *RouterIpv6Rip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Rip::Put")
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

func (p *RouterIpv6Rip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Rip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
