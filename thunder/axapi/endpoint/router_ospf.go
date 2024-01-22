

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type RouterOspf struct {
	Inst struct {

    AreaList []RouterOspfAreaList `json:"area-list"`
    AutoCostReferenceBandwidth int `json:"auto-cost-reference-bandwidth" dval:"10000"`
    BfdAllInterfaces int `json:"bfd-all-interfaces"`
    DefaultInformation RouterOspfDefaultInformation1297 `json:"default-information"`
    DefaultMetric int `json:"default-metric" dval:"20"`
    Distance RouterOspfDistance `json:"distance"`
    DistributeInternalList []RouterOspfDistributeInternalList `json:"distribute-internal-list"`
    DistributeLists []RouterOspfDistributeLists `json:"distribute-lists"`
    ExternLsaEquivalenceCheck int `json:"extern-lsa-equivalence-check"`
    HaStandbyExtraCost []RouterOspfHaStandbyExtraCost `json:"ha-standby-extra-cost"`
    HostList []RouterOspfHostList `json:"host-list"`
    LogAdjacencyChangesCfg RouterOspfLogAdjacencyChangesCfg `json:"log-adjacency-changes-cfg"`
    MaxConcurrentDd int `json:"max-concurrent-dd" dval:"5"`
    MaximumArea int `json:"maximum-area"`
    NeighborList []RouterOspfNeighborList `json:"neighbor-list"`
    NetworkList []RouterOspfNetworkList `json:"network-list"`
    Ospf1 RouterOspfOspf1 `json:"ospf-1"`
    Overflow RouterOspfOverflow `json:"overflow"`
    PassiveInterface RouterOspfPassiveInterface `json:"passive-interface"`
    ProcessId int `json:"process-id"`
    Redistribute RouterOspfRedistribute1298 `json:"redistribute"`
    Rfc1583Compatible int `json:"rfc1583-compatible"`
    RouterId RouterOspfRouterId `json:"router-id"`
    SummaryAddressList []RouterOspfSummaryAddressList `json:"summary-address-list"`
    Timers RouterOspfTimers `json:"timers"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ospf"`
}


type RouterOspfAreaList struct {
    AreaIpv4 string `json:"area-ipv4"`
    AreaNum int `json:"area-num"`
    AuthCfg RouterOspfAreaListAuthCfg `json:"auth-cfg"`
    FilterLists []RouterOspfAreaListFilterLists `json:"filter-lists"`
    NssaCfg RouterOspfAreaListNssaCfg `json:"nssa-cfg"`
    DefaultCost int `json:"default-cost" dval:"1"`
    RangeList []RouterOspfAreaListRangeList `json:"range-list"`
    Shortcut string `json:"shortcut" dval:"default"`
    StubCfg RouterOspfAreaListStubCfg `json:"stub-cfg"`
    VirtualLinkList []RouterOspfAreaListVirtualLinkList `json:"virtual-link-list"`
    Uuid string `json:"uuid"`
}


type RouterOspfAreaListAuthCfg struct {
    Authentication int `json:"authentication"`
    MessageDigest int `json:"message-digest"`
}


type RouterOspfAreaListFilterLists struct {
    FilterList int `json:"filter-list"`
    AclName string `json:"acl-name"`
    AclDirection string `json:"acl-direction"`
    PlistName string `json:"plist-name"`
    PlistDirection string `json:"plist-direction"`
}


type RouterOspfAreaListNssaCfg struct {
    Nssa int `json:"nssa"`
    NoRedistribution int `json:"no-redistribution"`
    NoSummary int `json:"no-summary"`
    TranslatorRole string `json:"translator-role" dval:"candidate"`
    DefaultInformationOriginate int `json:"default-information-originate"`
    Metric int `json:"metric" dval:"1"`
    MetricType int `json:"metric-type" dval:"2"`
}


type RouterOspfAreaListRangeList struct {
    AreaRangePrefix string `json:"area-range-prefix"`
    Option string `json:"option" dval:"advertise"`
}


type RouterOspfAreaListStubCfg struct {
    Stub int `json:"stub"`
    NoSummary int `json:"no-summary"`
}


type RouterOspfAreaListVirtualLinkList struct {
    VirtualLinkIpAddr string `json:"virtual-link-ip-addr"`
    Bfd int `json:"bfd"`
    HelloInterval int `json:"hello-interval"`
    DeadInterval int `json:"dead-interval"`
    RetransmitInterval int `json:"retransmit-interval"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    VirtualLinkAuthentication int `json:"virtual-link-authentication"`
    VirtualLinkAuthType string `json:"virtual-link-auth-type"`
    AuthenticationKey string `json:"authentication-key"`
    MessageDigestKey int `json:"message-digest-key"`
    Md5 string `json:"md5"`
}


type RouterOspfDefaultInformation1297 struct {
    Originate int `json:"originate"`
    Always int `json:"always"`
    Metric int `json:"metric"`
    MetricType int `json:"metric-type" dval:"2"`
    RouteMap string `json:"route-map"`
    Uuid string `json:"uuid"`
}


type RouterOspfDistance struct {
    DistanceValue int `json:"distance-value" dval:"110"`
    DistanceOspf RouterOspfDistanceDistanceOspf `json:"distance-ospf"`
}


type RouterOspfDistanceDistanceOspf struct {
    DistanceExternal int `json:"distance-external"`
    DistanceInterArea int `json:"distance-inter-area"`
    DistanceIntraArea int `json:"distance-intra-area"`
}


type RouterOspfDistributeInternalList struct {
    DiType string `json:"di-type"`
    DiAreaIpv4 string `json:"di-area-ipv4"`
    DiAreaNum int `json:"di-area-num"`
    DiCost int `json:"di-cost" dval:"1"`
}


type RouterOspfDistributeLists struct {
    Value string `json:"value"`
    Direction string `json:"direction"`
    Protocol string `json:"protocol"`
    OspfId int `json:"ospf-id"`
    Option string `json:"option"`
}


type RouterOspfHaStandbyExtraCost struct {
    ExtraCost int `json:"extra-cost"`
    Group int `json:"group"`
}


type RouterOspfHostList struct {
    HostAddress string `json:"host-address"`
    AreaCfg RouterOspfHostListAreaCfg `json:"area-cfg"`
}


type RouterOspfHostListAreaCfg struct {
    AreaIpv4 string `json:"area-ipv4"`
    AreaNum int `json:"area-num"`
    Cost int `json:"cost"`
}


type RouterOspfLogAdjacencyChangesCfg struct {
    State string `json:"state"`
}


type RouterOspfNeighborList struct {
    Address string `json:"address"`
    Cost int `json:"cost"`
    PollInterval int `json:"poll-interval"`
    Priority int `json:"priority"`
}


type RouterOspfNetworkList struct {
    NetworkIpv4 string `json:"network-ipv4"`
    NetworkIpv4Mask string `json:"network-ipv4-mask"`
    NetworkIpv4Cidr string `json:"network-ipv4-cidr"`
    NetworkArea RouterOspfNetworkListNetworkArea `json:"network-area"`
}


type RouterOspfNetworkListNetworkArea struct {
    NetworkAreaIpv4 string `json:"network-area-ipv4"`
    NetworkAreaNum int `json:"network-area-num"`
    InstanceValue int `json:"instance-value"`
}


type RouterOspfOspf1 struct {
    AbrType RouterOspfOspf1AbrType `json:"abr-type"`
}


type RouterOspfOspf1AbrType struct {
    Option string `json:"option" dval:"cisco"`
}


type RouterOspfOverflow struct {
    Database RouterOspfOverflowDatabase `json:"database"`
}


type RouterOspfOverflowDatabase struct {
    Count1 int `json:"count1" dval:"4294967294"`
    Limit string `json:"limit" dval:"hard"`
    DbExternal int `json:"db-external" dval:"2147483647"`
    RecoveryTime int `json:"recovery-time"`
}


type RouterOspfPassiveInterface struct {
    LoopbackCfg []RouterOspfPassiveInterfaceLoopbackCfg `json:"loopback-cfg"`
    TrunkCfg []RouterOspfPassiveInterfaceTrunkCfg `json:"trunk-cfg"`
    VeCfg []RouterOspfPassiveInterfaceVeCfg `json:"ve-cfg"`
    TunnelCfg []RouterOspfPassiveInterfaceTunnelCfg `json:"tunnel-cfg"`
    LifCfg []RouterOspfPassiveInterfaceLifCfg `json:"lif-cfg"`
    EthCfg []RouterOspfPassiveInterfaceEthCfg `json:"eth-cfg"`
}


type RouterOspfPassiveInterfaceLoopbackCfg struct {
    Loopback int `json:"loopback"`
    LoopbackAddress string `json:"loopback-address"`
}


type RouterOspfPassiveInterfaceTrunkCfg struct {
    Trunk int `json:"trunk"`
    TrunkAddress string `json:"trunk-address"`
}


type RouterOspfPassiveInterfaceVeCfg struct {
    Ve int `json:"ve"`
    VeAddress string `json:"ve-address"`
}


type RouterOspfPassiveInterfaceTunnelCfg struct {
    Tunnel int `json:"tunnel"`
    TunnelAddress string `json:"tunnel-address"`
}


type RouterOspfPassiveInterfaceLifCfg struct {
    Lif string `json:"lif"`
    LifAddress string `json:"lif-address"`
}


type RouterOspfPassiveInterfaceEthCfg struct {
    Ethernet int `json:"ethernet"`
    EthAddress string `json:"eth-address"`
}


type RouterOspfRedistribute1298 struct {
    RedistList []RouterOspfRedistributeRedistList1299 `json:"redist-list"`
    OspfList []RouterOspfRedistributeOspfList1300 `json:"ospf-list"`
    IpNat int `json:"ip-nat"`
    MetricIpNat int `json:"metric-ip-nat"`
    MetricTypeIpNat string `json:"metric-type-ip-nat" dval:"2"`
    RouteMapIpNat string `json:"route-map-ip-nat"`
    TagIpNat int `json:"tag-ip-nat"`
    IpNatFloatingList []RouterOspfRedistributeIpNatFloatingList1301 `json:"ip-nat-floating-list"`
    VipList []RouterOspfRedistributeVipList1302 `json:"vip-list"`
    VipFloatingList []RouterOspfRedistributeVipFloatingList1303 `json:"vip-floating-list"`
    Uuid string `json:"uuid"`
}


type RouterOspfRedistributeRedistList1299 struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"2"`
    RouteMap string `json:"route-map"`
    Tag int `json:"tag"`
}


type RouterOspfRedistributeOspfList1300 struct {
    Ospf int `json:"ospf"`
    ProcessId int `json:"process-id"`
    MetricOspf int `json:"metric-ospf"`
    MetricTypeOspf string `json:"metric-type-ospf" dval:"2"`
    RouteMapOspf string `json:"route-map-ospf"`
    TagOspf int `json:"tag-ospf"`
}


type RouterOspfRedistributeIpNatFloatingList1301 struct {
    IpNatPrefix string `json:"ip-nat-prefix"`
    IpNatFloatingIpForward string `json:"ip-nat-floating-IP-forward"`
}


type RouterOspfRedistributeVipList1302 struct {
    TypeVip string `json:"type-vip"`
    MetricVip int `json:"metric-vip"`
    MetricTypeVip string `json:"metric-type-vip" dval:"2"`
    RouteMapVip string `json:"route-map-vip"`
    TagVip int `json:"tag-vip"`
}


type RouterOspfRedistributeVipFloatingList1303 struct {
    VipAddress string `json:"vip-address"`
    VipFloatingIpForward string `json:"vip-floating-IP-forward"`
}


type RouterOspfRouterId struct {
    Value string `json:"value"`
}


type RouterOspfSummaryAddressList struct {
    SummaryAddress string `json:"summary-address"`
    NotAdvertise int `json:"not-advertise"`
    Tag int `json:"tag"`
}


type RouterOspfTimers struct {
    Spf RouterOspfTimersSpf `json:"spf"`
}


type RouterOspfTimersSpf struct {
    Exp RouterOspfTimersSpfExp `json:"exp"`
}


type RouterOspfTimersSpfExp struct {
    MinDelay int `json:"min-delay" dval:"500"`
    MaxDelay int `json:"max-delay" dval:"50000"`
}

func (p *RouterOspf) GetId() string{
    return strconv.Itoa(p.Inst.ProcessId)
}

func (p *RouterOspf) getPath() string{
    return "router/ospf"
}

func (p *RouterOspf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspf::Post")
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

func (p *RouterOspf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspf::Get")
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
func (p *RouterOspf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspf::Put")
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

func (p *RouterOspf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
