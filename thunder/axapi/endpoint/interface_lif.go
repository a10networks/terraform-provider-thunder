

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLif struct {
	Inst struct {

    AccessList InterfaceLifAccessList `json:"access-list"`
    Action string `json:"action" dval:"enable"`
    Bfd InterfaceLifBfd593 `json:"bfd"`
    Encapsulation InterfaceLifEncapsulation596 `json:"encapsulation"`
    Ifname string `json:"ifname"`
    Ip InterfaceLifIp598 `json:"ip"`
    Ipv6 InterfaceLifIpv6620 `json:"ipv6"`
    Isis InterfaceLifIsis637 `json:"isis"`
    Mtu int `json:"mtu"`
    Name string `json:"name"`
    SamplingEnable []InterfaceLifSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"lif"`
}


type InterfaceLifAccessList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
}


type InterfaceLifBfd593 struct {
    Authentication InterfaceLifBfdAuthentication594 `json:"authentication"`
    Echo int `json:"echo"`
    Demand int `json:"demand"`
    IntervalCfg InterfaceLifBfdIntervalCfg595 `json:"interval-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLifBfdAuthentication594 struct {
    KeyId int `json:"key-id"`
    Method string `json:"method"`
    Password string `json:"password"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLifBfdIntervalCfg595 struct {
    Interval int `json:"interval"`
    MinRx int `json:"min-rx"`
    Multiplier int `json:"multiplier"`
}


type InterfaceLifEncapsulation596 struct {
    Dot1q InterfaceLifEncapsulationDot1q597 `json:"dot1q"`
}


type InterfaceLifEncapsulationDot1q597 struct {
    Tag int `json:"tag"`
    Ethernet int `json:"ethernet"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIp598 struct {
    Dhcp int `json:"dhcp"`
    AddressList []InterfaceLifIpAddressList599 `json:"address-list"`
    AllowPromiscuousVip int `json:"allow-promiscuous-vip"`
    CacheSpoofingPort int `json:"cache-spoofing-port"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    QueryInterval int `json:"query-interval" dval:"125"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Unnumbered int `json:"unnumbered"`
    Uuid string `json:"uuid"`
    Router InterfaceLifIpRouter600 `json:"router"`
    Rip InterfaceLifIpRip602 `json:"rip"`
    Ospf InterfaceLifIpOspf610 `json:"ospf"`
}


type InterfaceLifIpAddressList599 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceLifIpRouter600 struct {
    Isis InterfaceLifIpRouterIsis601 `json:"isis"`
}


type InterfaceLifIpRouterIsis601 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpRip602 struct {
    Authentication InterfaceLifIpRipAuthentication603 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceLifIpRipSendCfg607 `json:"send-cfg"`
    ReceiveCfg InterfaceLifIpRipReceiveCfg608 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceLifIpRipSplitHorizonCfg609 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpRipAuthentication603 struct {
    Str InterfaceLifIpRipAuthenticationStr604 `json:"str"`
    Mode InterfaceLifIpRipAuthenticationMode605 `json:"mode"`
    KeyChain InterfaceLifIpRipAuthenticationKeyChain606 `json:"key-chain"`
}


type InterfaceLifIpRipAuthenticationStr604 struct {
    String string `json:"string"`
}


type InterfaceLifIpRipAuthenticationMode605 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceLifIpRipAuthenticationKeyChain606 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceLifIpRipSendCfg607 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceLifIpRipReceiveCfg608 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceLifIpRipSplitHorizonCfg609 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceLifIpOspf610 struct {
    OspfGlobal InterfaceLifIpOspfOspfGlobal611 `json:"ospf-global"`
    OspfIpList []InterfaceLifIpOspfOspfIpList618 `json:"ospf-ip-list"`
}


type InterfaceLifIpOspfOspfGlobal611 struct {
    AuthenticationCfg InterfaceLifIpOspfOspfGlobalAuthenticationCfg612 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceLifIpOspfOspfGlobalBfdCfg613 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg614 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceLifIpOspfOspfGlobalMessageDigestCfg615 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceLifIpOspfOspfGlobalNetwork617 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpOspfOspfGlobalAuthenticationCfg612 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceLifIpOspfOspfGlobalBfdCfg613 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg614 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceLifIpOspfOspfGlobalMessageDigestCfg615 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5616 `json:"md5"`
}


type InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5616 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLifIpOspfOspfGlobalNetwork617 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceLifIpOspfOspfIpList618 struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceLifIpOspfOspfIpListMessageDigestCfg619 `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpOspfOspfIpListMessageDigestCfg619 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLifIpv6620 struct {
    AddressList []InterfaceLifIpv6AddressList621 `json:"address-list"`
    Ipv6Enable int `json:"ipv6-enable"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    Uuid string `json:"uuid"`
    Router InterfaceLifIpv6Router622 `json:"router"`
    Ospf InterfaceLifIpv6Ospf627 `json:"ospf"`
}


type InterfaceLifIpv6AddressList621 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    Anycast int `json:"anycast"`
    LinkLocal int `json:"link-local"`
}


type InterfaceLifIpv6Router622 struct {
    Ripng InterfaceLifIpv6RouterRipng623 `json:"ripng"`
    Ospf InterfaceLifIpv6RouterOspf624 `json:"ospf"`
    Isis InterfaceLifIpv6RouterIsis626 `json:"isis"`
}


type InterfaceLifIpv6RouterRipng623 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpv6RouterOspf624 struct {
    AreaList []InterfaceLifIpv6RouterOspfAreaList625 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpv6RouterOspfAreaList625 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6RouterIsis626 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpv6Ospf627 struct {
    NetworkList []InterfaceLifIpv6OspfNetworkList628 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceLifIpv6OspfCostCfg629 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceLifIpv6OspfDeadIntervalCfg630 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceLifIpv6OspfHelloIntervalCfg631 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceLifIpv6OspfMtuIgnoreCfg632 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceLifIpv6OspfNeighborCfg633 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceLifIpv6OspfPriorityCfg634 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceLifIpv6OspfRetransmitIntervalCfg635 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceLifIpv6OspfTransmitDelayCfg636 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpv6OspfNetworkList628 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceLifIpv6OspfCostCfg629 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfDeadIntervalCfg630 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfHelloIntervalCfg631 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfMtuIgnoreCfg632 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfNeighborCfg633 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceLifIpv6OspfPriorityCfg634 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfRetransmitIntervalCfg635 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfTransmitDelayCfg636 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIsis637 struct {
    Authentication InterfaceLifIsisAuthentication638 `json:"authentication"`
    BfdCfg InterfaceLifIsisBfdCfg642 `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceLifIsisCsnpIntervalList643 `json:"csnp-interval-list"`
    Padding int `json:"padding" dval:"1"`
    HelloIntervalList []InterfaceLifIsisHelloIntervalList644 `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceLifIsisHelloIntervalMinimalList645 `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceLifIsisHelloMultiplierList646 `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceLifIsisMeshGroup647 `json:"mesh-group"`
    MetricList []InterfaceLifIsisMetricList648 `json:"metric-list"`
    Network string `json:"network"`
    PasswordList []InterfaceLifIsisPasswordList649 `json:"password-list"`
    PriorityList []InterfaceLifIsisPriorityList650 `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    WideMetricList []InterfaceLifIsisWideMetricList651 `json:"wide-metric-list"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIsisAuthentication638 struct {
    SendOnlyList []InterfaceLifIsisAuthenticationSendOnlyList639 `json:"send-only-list"`
    ModeList []InterfaceLifIsisAuthenticationModeList640 `json:"mode-list"`
    KeyChainList []InterfaceLifIsisAuthenticationKeyChainList641 `json:"key-chain-list"`
}


type InterfaceLifIsisAuthenticationSendOnlyList639 struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceLifIsisAuthenticationModeList640 struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceLifIsisAuthenticationKeyChainList641 struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceLifIsisBfdCfg642 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLifIsisCsnpIntervalList643 struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLifIsisHelloIntervalList644 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLifIsisHelloIntervalMinimalList645 struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceLifIsisHelloMultiplierList646 struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceLifIsisMeshGroup647 struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceLifIsisMetricList648 struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLifIsisPasswordList649 struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceLifIsisPriorityList650 struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceLifIsisWideMetricList651 struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLifSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *InterfaceLif) GetId() string{
    return p.Inst.Ifname
}

func (p *InterfaceLif) getPath() string{
    return "interface/lif"
}

func (p *InterfaceLif) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLif::Post")
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

func (p *InterfaceLif) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLif::Get")
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
func (p *InterfaceLif) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLif::Put")
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

func (p *InterfaceLif) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLif::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
