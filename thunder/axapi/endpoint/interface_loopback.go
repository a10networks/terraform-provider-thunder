

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopback struct {
	Inst struct {

    Ifnum int `json:"ifnum"`
    Ip InterfaceLoopbackIp684 `json:"ip"`
    Ipv6 InterfaceLoopbackIpv6705 `json:"ipv6"`
    Isis InterfaceLoopbackIsis722 `json:"isis"`
    Name string `json:"name"`
    SnmpServer InterfaceLoopbackSnmpServer `json:"snmp-server"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"loopback"`
}


type InterfaceLoopbackIp684 struct {
    AddressList []InterfaceLoopbackIpAddressList685 `json:"address-list"`
    Uuid string `json:"uuid"`
    Router InterfaceLoopbackIpRouter686 `json:"router"`
    Rip InterfaceLoopbackIpRip688 `json:"rip"`
    Ospf InterfaceLoopbackIpOspf696 `json:"ospf"`
}


type InterfaceLoopbackIpAddressList685 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceLoopbackIpRouter686 struct {
    Isis InterfaceLoopbackIpRouterIsis687 `json:"isis"`
}


type InterfaceLoopbackIpRouterIsis687 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpRip688 struct {
    Authentication InterfaceLoopbackIpRipAuthentication689 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceLoopbackIpRipSendCfg693 `json:"send-cfg"`
    ReceiveCfg InterfaceLoopbackIpRipReceiveCfg694 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceLoopbackIpRipSplitHorizonCfg695 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpRipAuthentication689 struct {
    Str InterfaceLoopbackIpRipAuthenticationStr690 `json:"str"`
    Mode InterfaceLoopbackIpRipAuthenticationMode691 `json:"mode"`
    KeyChain InterfaceLoopbackIpRipAuthenticationKeyChain692 `json:"key-chain"`
}


type InterfaceLoopbackIpRipAuthenticationStr690 struct {
    String string `json:"string"`
}


type InterfaceLoopbackIpRipAuthenticationMode691 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceLoopbackIpRipAuthenticationKeyChain692 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceLoopbackIpRipSendCfg693 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceLoopbackIpRipReceiveCfg694 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceLoopbackIpRipSplitHorizonCfg695 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceLoopbackIpOspf696 struct {
    OspfGlobal InterfaceLoopbackIpOspfOspfGlobal697 `json:"ospf-global"`
    OspfIpList []InterfaceLoopbackIpOspfOspfIpList703 `json:"ospf-ip-list"`
}


type InterfaceLoopbackIpOspfOspfGlobal697 struct {
    AuthenticationCfg InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg698 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceLoopbackIpOspfOspfGlobalBfdCfg699 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg700 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg701 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg698 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceLoopbackIpOspfOspfGlobalBfdCfg699 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg700 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg701 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5702 `json:"md5"`
}


type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5702 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLoopbackIpOspfOspfIpList703 struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg704 `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg704 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLoopbackIpv6705 struct {
    AddressList []InterfaceLoopbackIpv6AddressList706 `json:"address-list"`
    Ipv6Enable int `json:"ipv6-enable"`
    Uuid string `json:"uuid"`
    Router InterfaceLoopbackIpv6Router707 `json:"router"`
    Rip InterfaceLoopbackIpv6Rip712 `json:"rip"`
    Ospf InterfaceLoopbackIpv6Ospf714 `json:"ospf"`
}


type InterfaceLoopbackIpv6AddressList706 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    Anycast int `json:"anycast"`
    LinkLocal int `json:"link-local"`
}


type InterfaceLoopbackIpv6Router707 struct {
    Ripng InterfaceLoopbackIpv6RouterRipng708 `json:"ripng"`
    Ospf InterfaceLoopbackIpv6RouterOspf709 `json:"ospf"`
    Isis InterfaceLoopbackIpv6RouterIsis711 `json:"isis"`
}


type InterfaceLoopbackIpv6RouterRipng708 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6RouterOspf709 struct {
    AreaList []InterfaceLoopbackIpv6RouterOspfAreaList710 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6RouterOspfAreaList710 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6RouterIsis711 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6Rip712 struct {
    SplitHorizonCfg InterfaceLoopbackIpv6RipSplitHorizonCfg713 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6RipSplitHorizonCfg713 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceLoopbackIpv6Ospf714 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceLoopbackIpv6OspfCostCfg715 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceLoopbackIpv6OspfDeadIntervalCfg716 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceLoopbackIpv6OspfHelloIntervalCfg717 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceLoopbackIpv6OspfMtuIgnoreCfg718 `json:"mtu-ignore-cfg"`
    PriorityCfg []InterfaceLoopbackIpv6OspfPriorityCfg719 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceLoopbackIpv6OspfRetransmitIntervalCfg720 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceLoopbackIpv6OspfTransmitDelayCfg721 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6OspfCostCfg715 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfDeadIntervalCfg716 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfHelloIntervalCfg717 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfMtuIgnoreCfg718 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfPriorityCfg719 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfRetransmitIntervalCfg720 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfTransmitDelayCfg721 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIsis722 struct {
    Authentication InterfaceLoopbackIsisAuthentication723 `json:"authentication"`
    BfdCfg InterfaceLoopbackIsisBfdCfg727 `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceLoopbackIsisCsnpIntervalList728 `json:"csnp-interval-list"`
    Padding int `json:"padding" dval:"1"`
    HelloIntervalList []InterfaceLoopbackIsisHelloIntervalList729 `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceLoopbackIsisHelloIntervalMinimalList730 `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceLoopbackIsisHelloMultiplierList731 `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceLoopbackIsisMeshGroup732 `json:"mesh-group"`
    MetricList []InterfaceLoopbackIsisMetricList733 `json:"metric-list"`
    PasswordList []InterfaceLoopbackIsisPasswordList734 `json:"password-list"`
    PriorityList []InterfaceLoopbackIsisPriorityList735 `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    WideMetricList []InterfaceLoopbackIsisWideMetricList736 `json:"wide-metric-list"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIsisAuthentication723 struct {
    SendOnlyList []InterfaceLoopbackIsisAuthenticationSendOnlyList724 `json:"send-only-list"`
    ModeList []InterfaceLoopbackIsisAuthenticationModeList725 `json:"mode-list"`
    KeyChainList []InterfaceLoopbackIsisAuthenticationKeyChainList726 `json:"key-chain-list"`
}


type InterfaceLoopbackIsisAuthenticationSendOnlyList724 struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisAuthenticationModeList725 struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisAuthenticationKeyChainList726 struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisBfdCfg727 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLoopbackIsisCsnpIntervalList728 struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisHelloIntervalList729 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisHelloIntervalMinimalList730 struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisHelloMultiplierList731 struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisMeshGroup732 struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceLoopbackIsisMetricList733 struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisPasswordList734 struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisPriorityList735 struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceLoopbackIsisWideMetricList736 struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceLoopbackSnmpServer struct {
    TrapSource int `json:"trap-source"`
}

func (p *InterfaceLoopback) GetId() string{
    return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceLoopback) getPath() string{
    return "interface/loopback"
}

func (p *InterfaceLoopback) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopback::Post")
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

func (p *InterfaceLoopback) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopback::Get")
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
func (p *InterfaceLoopback) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopback::Put")
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

func (p *InterfaceLoopback) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopback::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
