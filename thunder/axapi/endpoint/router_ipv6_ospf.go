

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6Ospf struct {
	Inst struct {

    AbrTypeOption string `json:"abr-type-option" dval:"cisco"`
    AreaList []RouterIpv6OspfAreaList `json:"area-list"`
    AutoCostReferenceBandwidth int `json:"auto-cost-reference-bandwidth" dval:"10000"`
    BfdAllInterfaces int `json:"bfd-all-interfaces"`
    DefaultInformation RouterIpv6OspfDefaultInformation1248 `json:"default-information"`
    DefaultMetric int `json:"default-metric" dval:"20"`
    DistributeInternalList []RouterIpv6OspfDistributeInternalList `json:"distribute-internal-list"`
    DistributeList RouterIpv6OspfDistributeList `json:"distribute-list"`
    HaStandbyExtraCost []RouterIpv6OspfHaStandbyExtraCost `json:"ha-standby-extra-cost"`
    LogAdjacencyChanges string `json:"log-adjacency-changes"`
    MaxConcurrentDd int `json:"max-concurrent-dd" dval:"5"`
    PassiveInterface RouterIpv6OspfPassiveInterface `json:"passive-interface"`
    ProcessId string `json:"process-id"`
    Redistribute RouterIpv6OspfRedistribute1249 `json:"redistribute"`
    RouterId string `json:"router-id"`
    Timers RouterIpv6OspfTimers `json:"timers"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ospf"`
}


type RouterIpv6OspfAreaList struct {
    AreaIpv4 string `json:"area-ipv4"`
    AreaNum int `json:"area-num"`
    DefaultCost int `json:"default-cost" dval:"1"`
    RangeList []RouterIpv6OspfAreaListRangeList `json:"range-list"`
    Stub int `json:"stub"`
    NoSummary int `json:"no-summary"`
    VirtualLinkList []RouterIpv6OspfAreaListVirtualLinkList `json:"virtual-link-list"`
    Uuid string `json:"uuid"`
}


type RouterIpv6OspfAreaListRangeList struct {
    Value string `json:"value"`
    Option string `json:"option"`
}


type RouterIpv6OspfAreaListVirtualLinkList struct {
    Value string `json:"value"`
    DeadInterval int `json:"dead-interval"`
    Bfd int `json:"bfd"`
    HelloInterval int `json:"hello-interval"`
    RetransmitInterval int `json:"retransmit-interval"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type RouterIpv6OspfDefaultInformation1248 struct {
    Originate int `json:"originate"`
    Always int `json:"always"`
    Metric int `json:"metric"`
    MetricType int `json:"metric-type" dval:"2"`
    RouteMap string `json:"route-map"`
    Uuid string `json:"uuid"`
}


type RouterIpv6OspfDistributeInternalList struct {
    Type string `json:"type"`
    AreaIpv4 string `json:"area-ipv4" dval:"255.255.255.255"`
    AreaNum int `json:"area-num"`
    Cost int `json:"cost" dval:"1"`
}


type RouterIpv6OspfDistributeList struct {
    PrefixList RouterIpv6OspfDistributeListPrefixList `json:"prefix-list"`
}


type RouterIpv6OspfDistributeListPrefixList struct {
    Value string `json:"value"`
    Direction string `json:"direction"`
}


type RouterIpv6OspfHaStandbyExtraCost struct {
    ExtraCost int `json:"extra-cost"`
    Group int `json:"group"`
}


type RouterIpv6OspfPassiveInterface struct {
    LoopbackCfg []RouterIpv6OspfPassiveInterfaceLoopbackCfg `json:"loopback-cfg"`
    TrunkCfg []RouterIpv6OspfPassiveInterfaceTrunkCfg `json:"trunk-cfg"`
    VeCfg []RouterIpv6OspfPassiveInterfaceVeCfg `json:"ve-cfg"`
    TunnelCfg []RouterIpv6OspfPassiveInterfaceTunnelCfg `json:"tunnel-cfg"`
    EthCfg []RouterIpv6OspfPassiveInterfaceEthCfg `json:"eth-cfg"`
}


type RouterIpv6OspfPassiveInterfaceLoopbackCfg struct {
    Loopback int `json:"loopback"`
}


type RouterIpv6OspfPassiveInterfaceTrunkCfg struct {
    Trunk int `json:"trunk"`
}


type RouterIpv6OspfPassiveInterfaceVeCfg struct {
    Ve int `json:"ve"`
}


type RouterIpv6OspfPassiveInterfaceTunnelCfg struct {
    Tunnel int `json:"tunnel"`
}


type RouterIpv6OspfPassiveInterfaceEthCfg struct {
    Ethernet int `json:"ethernet"`
}


type RouterIpv6OspfRedistribute1249 struct {
    RedistList []RouterIpv6OspfRedistributeRedistList1250 `json:"redist-list"`
    OspfList []RouterIpv6OspfRedistributeOspfList1251 `json:"ospf-list"`
    IpNat int `json:"ip-nat"`
    MetricIpNat int `json:"metric-ip-nat"`
    MetricTypeIpNat string `json:"metric-type-ip-nat" dval:"2"`
    RouteMapIpNat string `json:"route-map-ip-nat"`
    IpNatFloatingList []RouterIpv6OspfRedistributeIpNatFloatingList1252 `json:"ip-nat-floating-list"`
    VipList []RouterIpv6OspfRedistributeVipList1253 `json:"vip-list"`
    VipFloatingList []RouterIpv6OspfRedistributeVipFloatingList1254 `json:"vip-floating-list"`
    Uuid string `json:"uuid"`
}


type RouterIpv6OspfRedistributeRedistList1250 struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"2"`
    RouteMap string `json:"route-map"`
}


type RouterIpv6OspfRedistributeOspfList1251 struct {
    Ospf int `json:"ospf"`
    ProcessId string `json:"process-id"`
    MetricOspf int `json:"metric-ospf"`
    MetricTypeOspf string `json:"metric-type-ospf" dval:"2"`
    RouteMapOspf string `json:"route-map-ospf"`
}


type RouterIpv6OspfRedistributeIpNatFloatingList1252 struct {
    IpNatPrefix string `json:"ip-nat-prefix"`
    IpNatFloatingIpForward string `json:"ip-nat-floating-IP-forward"`
}


type RouterIpv6OspfRedistributeVipList1253 struct {
    TypeVip string `json:"type-vip"`
    MetricVip int `json:"metric-vip"`
    MetricTypeVip string `json:"metric-type-vip" dval:"2"`
    RouteMapVip string `json:"route-map-vip"`
}


type RouterIpv6OspfRedistributeVipFloatingList1254 struct {
    VipAddress string `json:"vip-address"`
    VipFloatingIpForward string `json:"vip-floating-IP-forward"`
}


type RouterIpv6OspfTimers struct {
    Spf RouterIpv6OspfTimersSpf `json:"spf"`
}


type RouterIpv6OspfTimersSpf struct {
    Exp RouterIpv6OspfTimersSpfExp `json:"exp"`
}


type RouterIpv6OspfTimersSpfExp struct {
    MinDelay int `json:"min-delay" dval:"500"`
    MaxDelay int `json:"max-delay" dval:"50000"`
}

func (p *RouterIpv6Ospf) GetId() string{
    return p.Inst.ProcessId
}

func (p *RouterIpv6Ospf) getPath() string{
    return "router/ipv6/ospf"
}

func (p *RouterIpv6Ospf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Ospf::Post")
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

func (p *RouterIpv6Ospf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Ospf::Get")
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
func (p *RouterIpv6Ospf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Ospf::Put")
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

func (p *RouterIpv6Ospf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6Ospf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
