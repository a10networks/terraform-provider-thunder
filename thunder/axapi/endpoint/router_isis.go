

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIsis struct {
	Inst struct {

    AddressFamily RouterIsisAddressFamily1276 `json:"address-family"`
    AdjacencyCheck int `json:"adjacency-check" dval:"1"`
    AreaPasswordCfg RouterIsisAreaPasswordCfg `json:"area-password-cfg"`
    Authentication RouterIsisAuthentication `json:"authentication"`
    Bfd string `json:"bfd"`
    DefaultInformation string `json:"default-information"`
    DistanceList []RouterIsisDistanceList `json:"distance-list"`
    DomainPasswordCfg RouterIsisDomainPasswordCfg `json:"domain-password-cfg"`
    HaStandbyExtraCost []RouterIsisHaStandbyExtraCost `json:"ha-standby-extra-cost"`
    IgnoreLspErrors int `json:"ignore-lsp-errors"`
    IsType string `json:"is-type" dval:"level-1"`
    LogAdjacencyChangesCfg RouterIsisLogAdjacencyChangesCfg `json:"log-adjacency-changes-cfg"`
    LspGenIntervalList []RouterIsisLspGenIntervalList `json:"lsp-gen-interval-list"`
    LspRefreshInterval int `json:"lsp-refresh-interval" dval:"900"`
    MaxLspLifetime int `json:"max-lsp-lifetime" dval:"1200"`
    MetricStyleList []RouterIsisMetricStyleList `json:"metric-style-list"`
    NetList []RouterIsisNetList `json:"net-list"`
    PassiveInterfaceList []RouterIsisPassiveInterfaceList `json:"passive-interface-list"`
    ProtocolList []RouterIsisProtocolList `json:"protocol-list"`
    Redistribute RouterIsisRedistribute1288 `json:"redistribute"`
    SetOverloadBitCfg RouterIsisSetOverloadBitCfg `json:"set-overload-bit-cfg"`
    SpfIntervalExpList []RouterIsisSpfIntervalExpList `json:"spf-interval-exp-list"`
    SummaryAddressList []RouterIsisSummaryAddressList `json:"summary-address-list"`
    Tag string `json:"tag"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"isis"`
}


type RouterIsisAddressFamily1276 struct {
    Ipv6 RouterIsisAddressFamilyIpv61277 `json:"ipv6"`
}


type RouterIsisAddressFamilyIpv61277 struct {
    DefaultInformation string `json:"default-information"`
    AdjacencyCheck int `json:"adjacency-check" dval:"1"`
    Distance int `json:"distance" dval:"115"`
    MultiTopologyCfg RouterIsisAddressFamilyIpv6MultiTopologyCfg1278 `json:"multi-topology-cfg"`
    SummaryPrefixList []RouterIsisAddressFamilyIpv6SummaryPrefixList1279 `json:"summary-prefix-list"`
    Uuid string `json:"uuid"`
    Redistribute RouterIsisAddressFamilyIpv6Redistribute1280 `json:"redistribute"`
}


type RouterIsisAddressFamilyIpv6MultiTopologyCfg1278 struct {
    MultiTopology int `json:"multi-topology"`
    Level string `json:"level"`
    Transition int `json:"transition"`
    LevelTransition int `json:"level-transition"`
}


type RouterIsisAddressFamilyIpv6SummaryPrefixList1279 struct {
    Prefix string `json:"prefix"`
    Level string `json:"level" dval:"level-2"`
}


type RouterIsisAddressFamilyIpv6Redistribute1280 struct {
    RedistList []RouterIsisAddressFamilyIpv6RedistributeRedistList1281 `json:"redist-list"`
    VipList []RouterIsisAddressFamilyIpv6RedistributeVipList1282 `json:"vip-list"`
    Isis RouterIsisAddressFamilyIpv6RedistributeIsis1283 `json:"isis"`
    Uuid string `json:"uuid"`
}


type RouterIsisAddressFamilyIpv6RedistributeRedistList1281 struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"internal"`
    RouteMap string `json:"route-map"`
    Level string `json:"level" dval:"level-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeVipList1282 struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
    VipMetricType string `json:"vip-metric-type" dval:"internal"`
    VipLevel string `json:"vip-level" dval:"level-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsis1283 struct {
    Level1From RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1284 `json:"level-1-from"`
    Level2From RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1286 `json:"level-2-from"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1284 struct {
    Into1 RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11285 `json:"into-1"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11285 struct {
    Level2 int `json:"level-2"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1286 struct {
    Into2 RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21287 `json:"into-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21287 struct {
    Level1 int `json:"level-1"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisAreaPasswordCfg struct {
    Password string `json:"password"`
    Authenticate RouterIsisAreaPasswordCfgAuthenticate `json:"authenticate"`
}


type RouterIsisAreaPasswordCfgAuthenticate struct {
    Snp string `json:"snp"`
}


type RouterIsisAuthentication struct {
    SendOnlyList []RouterIsisAuthenticationSendOnlyList `json:"send-only-list"`
    ModeList []RouterIsisAuthenticationModeList `json:"mode-list"`
    KeyChainList []RouterIsisAuthenticationKeyChainList `json:"key-chain-list"`
}


type RouterIsisAuthenticationSendOnlyList struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type RouterIsisAuthenticationModeList struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type RouterIsisAuthenticationKeyChainList struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type RouterIsisDistanceList struct {
    Distance int `json:"distance" dval:"115"`
    SystemId string `json:"System-ID"`
    Acl string `json:"acl"`
}


type RouterIsisDomainPasswordCfg struct {
    Password string `json:"password"`
    Authenticate RouterIsisDomainPasswordCfgAuthenticate `json:"authenticate"`
}


type RouterIsisDomainPasswordCfgAuthenticate struct {
    Snp string `json:"snp"`
}


type RouterIsisHaStandbyExtraCost struct {
    ExtraCost int `json:"extra-cost"`
    Group int `json:"group"`
}


type RouterIsisLogAdjacencyChangesCfg struct {
    State string `json:"state"`
}


type RouterIsisLspGenIntervalList struct {
    Interval int `json:"interval" dval:"30"`
    Level string `json:"level"`
}


type RouterIsisMetricStyleList struct {
    Type string `json:"type" dval:"narrow"`
    Level string `json:"level" dval:"level-1-2"`
}


type RouterIsisNetList struct {
    Net string `json:"net"`
}


type RouterIsisPassiveInterfaceList struct {
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Lif string `json:"lif"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
}


type RouterIsisProtocolList struct {
    ProtocolTopology int `json:"protocol-topology"`
}


type RouterIsisRedistribute1288 struct {
    RedistList []RouterIsisRedistributeRedistList1289 `json:"redist-list"`
    VipList []RouterIsisRedistributeVipList1290 `json:"vip-list"`
    Isis RouterIsisRedistributeIsis1291 `json:"isis"`
    Uuid string `json:"uuid"`
}


type RouterIsisRedistributeRedistList1289 struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"internal"`
    RouteMap string `json:"route-map"`
    Level string `json:"level" dval:"level-2"`
}


type RouterIsisRedistributeVipList1290 struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
    VipMetricType string `json:"vip-metric-type" dval:"internal"`
    VipLevel string `json:"vip-level" dval:"level-2"`
}


type RouterIsisRedistributeIsis1291 struct {
    Level1From RouterIsisRedistributeIsisLevel1From1292 `json:"level-1-from"`
    Level2From RouterIsisRedistributeIsisLevel2From1294 `json:"level-2-from"`
}


type RouterIsisRedistributeIsisLevel1From1292 struct {
    Into1 RouterIsisRedistributeIsisLevel1FromInto11293 `json:"into-1"`
}


type RouterIsisRedistributeIsisLevel1FromInto11293 struct {
    Level2 int `json:"level-2"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisRedistributeIsisLevel2From1294 struct {
    Into2 RouterIsisRedistributeIsisLevel2FromInto21295 `json:"into-2"`
}


type RouterIsisRedistributeIsisLevel2FromInto21295 struct {
    Level1 int `json:"level-1"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisSetOverloadBitCfg struct {
    SetOverloadBit int `json:"set-overload-bit"`
    OnStartup RouterIsisSetOverloadBitCfgOnStartup `json:"on-startup"`
    SuppressCfg RouterIsisSetOverloadBitCfgSuppressCfg `json:"suppress-cfg"`
}


type RouterIsisSetOverloadBitCfgOnStartup struct {
    Delay int `json:"delay"`
    WaitForBgp int `json:"wait-for-bgp"`
}


type RouterIsisSetOverloadBitCfgSuppressCfg struct {
    External int `json:"external"`
    Interlevel int `json:"interlevel"`
}


type RouterIsisSpfIntervalExpList struct {
    Min int `json:"min" dval:"500"`
    Max int `json:"max" dval:"50000"`
    Level string `json:"level"`
}


type RouterIsisSummaryAddressList struct {
    Prefix string `json:"prefix"`
    Level string `json:"level" dval:"level-2"`
}

func (p *RouterIsis) GetId() string{
    return p.Inst.Tag
}

func (p *RouterIsis) getPath() string{
    return "router/isis"
}

func (p *RouterIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsis::Post")
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

func (p *RouterIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsis::Get")
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
func (p *RouterIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsis::Put")
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

func (p *RouterIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
