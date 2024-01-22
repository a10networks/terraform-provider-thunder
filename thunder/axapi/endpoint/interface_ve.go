

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVe struct {
	Inst struct {

    AccessList InterfaceVeAccessList `json:"access-list"`
    Action string `json:"action" dval:"enable"`
    Bfd InterfaceVeBfd962 `json:"bfd"`
    Ddos InterfaceVeDdos965 `json:"ddos"`
    GamingProtocolCompliance int `json:"gaming-protocol-compliance"`
    IcmpRateLimit InterfaceVeIcmpRateLimit `json:"icmp-rate-limit"`
    Icmpv6RateLimit InterfaceVeIcmpv6RateLimit `json:"icmpv6-rate-limit"`
    Ifnum int `json:"ifnum"`
    Ip InterfaceVeIp966 `json:"ip"`
    Ipv6 InterfaceVeIpv6990 `json:"ipv6"`
    Isis InterfaceVeIsis1012 `json:"isis"`
    L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`
    Lw4o6 InterfaceVeLw4o61027 `json:"lw-4o6"`
    Map InterfaceVeMap1028 `json:"map"`
    Mtu int `json:"mtu"`
    Name string `json:"name"`
    Nptv6 InterfaceVeNptv61029 `json:"nptv6"`
    PingSweepDetection string `json:"ping-sweep-detection" dval:"disable"`
    PortScanDetection string `json:"port-scan-detection" dval:"disable"`
    SamplingEnable []InterfaceVeSamplingEnable `json:"sampling-enable"`
    TrapSource int `json:"trap-source"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ve"`
}


type InterfaceVeAccessList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
}


type InterfaceVeBfd962 struct {
    Authentication InterfaceVeBfdAuthentication963 `json:"authentication"`
    Echo int `json:"echo"`
    Demand int `json:"demand"`
    IntervalCfg InterfaceVeBfdIntervalCfg964 `json:"interval-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceVeBfdAuthentication963 struct {
    KeyId int `json:"key-id"`
    Method string `json:"method"`
    Password string `json:"password"`
    Encrypted string `json:"encrypted"`
}


type InterfaceVeBfdIntervalCfg964 struct {
    Interval int `json:"interval"`
    MinRx int `json:"min-rx"`
    Multiplier int `json:"multiplier"`
}


type InterfaceVeDdos965 struct {
    Outside int `json:"outside"`
    Inside int `json:"inside"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIcmpRateLimit struct {
    Normal int `json:"normal"`
    Lockup int `json:"lockup"`
    LockupPeriod int `json:"lockup-period"`
}


type InterfaceVeIcmpv6RateLimit struct {
    NormalV6 int `json:"normal-v6"`
    LockupV6 int `json:"lockup-v6"`
    LockupPeriodV6 int `json:"lockup-period-v6"`
}


type InterfaceVeIp966 struct {
    Dhcp int `json:"dhcp"`
    AddressList []InterfaceVeIpAddressList967 `json:"address-list"`
    AllowPromiscuousVip int `json:"allow-promiscuous-vip"`
    Client int `json:"client"`
    Server int `json:"server"`
    HelperAddressList []InterfaceVeIpHelperAddressList968 `json:"helper-address-list"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    TtlIgnore int `json:"ttl-ignore"`
    SynCookie int `json:"syn-cookie"`
    SlbPartitionRedirect int `json:"slb-partition-redirect"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    QueryInterval int `json:"query-interval" dval:"125"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Unnumbered int `json:"unnumbered"`
    Uuid string `json:"uuid"`
    StatefulFirewall InterfaceVeIpStatefulFirewall969 `json:"stateful-firewall"`
    Router InterfaceVeIpRouter970 `json:"router"`
    Rip InterfaceVeIpRip972 `json:"rip"`
    Ospf InterfaceVeIpOspf980 `json:"ospf"`
}


type InterfaceVeIpAddressList967 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceVeIpHelperAddressList968 struct {
    HelperAddress string `json:"helper-address"`
}


type InterfaceVeIpStatefulFirewall969 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclId int `json:"acl-id"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpRouter970 struct {
    Isis InterfaceVeIpRouterIsis971 `json:"isis"`
}


type InterfaceVeIpRouterIsis971 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpRip972 struct {
    Authentication InterfaceVeIpRipAuthentication973 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceVeIpRipSendCfg977 `json:"send-cfg"`
    ReceiveCfg InterfaceVeIpRipReceiveCfg978 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceVeIpRipSplitHorizonCfg979 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpRipAuthentication973 struct {
    Str InterfaceVeIpRipAuthenticationStr974 `json:"str"`
    Mode InterfaceVeIpRipAuthenticationMode975 `json:"mode"`
    KeyChain InterfaceVeIpRipAuthenticationKeyChain976 `json:"key-chain"`
}


type InterfaceVeIpRipAuthenticationStr974 struct {
    String string `json:"string"`
}


type InterfaceVeIpRipAuthenticationMode975 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceVeIpRipAuthenticationKeyChain976 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceVeIpRipSendCfg977 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceVeIpRipReceiveCfg978 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceVeIpRipSplitHorizonCfg979 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceVeIpOspf980 struct {
    OspfGlobal InterfaceVeIpOspfOspfGlobal981 `json:"ospf-global"`
    OspfIpList []InterfaceVeIpOspfOspfIpList988 `json:"ospf-ip-list"`
}


type InterfaceVeIpOspfOspfGlobal981 struct {
    AuthenticationCfg InterfaceVeIpOspfOspfGlobalAuthenticationCfg982 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceVeIpOspfOspfGlobalBfdCfg983 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg984 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceVeIpOspfOspfGlobalMessageDigestCfg985 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceVeIpOspfOspfGlobalNetwork987 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpOspfOspfGlobalAuthenticationCfg982 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceVeIpOspfOspfGlobalBfdCfg983 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg984 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceVeIpOspfOspfGlobalMessageDigestCfg985 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5986 `json:"md5"`
}


type InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5986 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceVeIpOspfOspfGlobalNetwork987 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceVeIpOspfOspfIpList988 struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceVeIpOspfOspfIpListMessageDigestCfg989 `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpOspfOspfIpListMessageDigestCfg989 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceVeIpv6990 struct {
    AddressList []InterfaceVeIpv6AddressList991 `json:"address-list"`
    Ipv6Enable int `json:"ipv6-enable"`
    V6AclName string `json:"v6-acl-name"`
    Inbound int `json:"inbound"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    TtlIgnore int `json:"ttl-ignore"`
    RouterAdver InterfaceVeIpv6RouterAdver992 `json:"router-adver"`
    Uuid string `json:"uuid"`
    StatefulFirewall InterfaceVeIpv6StatefulFirewall994 `json:"stateful-firewall"`
    Router InterfaceVeIpv6Router995 `json:"router"`
    Rip InterfaceVeIpv6Rip1000 `json:"rip"`
    Ospf InterfaceVeIpv6Ospf1002 `json:"ospf"`
}


type InterfaceVeIpv6AddressList991 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
}


type InterfaceVeIpv6RouterAdver992 struct {
    Action string `json:"action" dval:"disable"`
    DefaultLifetime int `json:"default-lifetime" dval:"1800"`
    HopLimit int `json:"hop-limit" dval:"255"`
    MaxInterval int `json:"max-interval" dval:"600"`
    MinInterval int `json:"min-interval" dval:"200"`
    RateLimit int `json:"rate-limit" dval:"100000"`
    ReachableTime int `json:"reachable-time"`
    RetransmitTimer int `json:"retransmit-timer"`
    AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
    AdverMtu int `json:"adver-mtu"`
    PrefixList []InterfaceVeIpv6RouterAdverPrefixList993 `json:"prefix-list"`
    ManagedConfigAction string `json:"managed-config-action" dval:"disable"`
    OtherConfigAction string `json:"other-config-action" dval:"disable"`
    AdverVrid int `json:"adver-vrid"`
    UseFloatingIp int `json:"use-floating-ip"`
    FloatingIp string `json:"floating-ip"`
    AdverVridDefault int `json:"adver-vrid-default"`
    UseFloatingIpDefaultVrid int `json:"use-floating-ip-default-vrid"`
    FloatingIpDefaultVrid string `json:"floating-ip-default-vrid"`
}


type InterfaceVeIpv6RouterAdverPrefixList993 struct {
    Prefix string `json:"prefix"`
    NotAutonomous int `json:"not-autonomous"`
    NotOnLink int `json:"not-on-link"`
    PreferredLifetime int `json:"preferred-lifetime" dval:"604800"`
    ValidLifetime int `json:"valid-lifetime" dval:"2592000"`
}


type InterfaceVeIpv6StatefulFirewall994 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclName string `json:"acl-name"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6Router995 struct {
    Ripng InterfaceVeIpv6RouterRipng996 `json:"ripng"`
    Ospf InterfaceVeIpv6RouterOspf997 `json:"ospf"`
    Isis InterfaceVeIpv6RouterIsis999 `json:"isis"`
}


type InterfaceVeIpv6RouterRipng996 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6RouterOspf997 struct {
    AreaList []InterfaceVeIpv6RouterOspfAreaList998 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6RouterOspfAreaList998 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6RouterIsis999 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6Rip1000 struct {
    SplitHorizonCfg InterfaceVeIpv6RipSplitHorizonCfg1001 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6RipSplitHorizonCfg1001 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceVeIpv6Ospf1002 struct {
    NetworkList []InterfaceVeIpv6OspfNetworkList1003 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceVeIpv6OspfCostCfg1004 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceVeIpv6OspfDeadIntervalCfg1005 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceVeIpv6OspfHelloIntervalCfg1006 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceVeIpv6OspfMtuIgnoreCfg1007 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceVeIpv6OspfNeighborCfg1008 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceVeIpv6OspfPriorityCfg1009 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceVeIpv6OspfRetransmitIntervalCfg1010 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceVeIpv6OspfTransmitDelayCfg1011 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6OspfNetworkList1003 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceVeIpv6OspfCostCfg1004 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfDeadIntervalCfg1005 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfHelloIntervalCfg1006 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfMtuIgnoreCfg1007 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfNeighborCfg1008 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceVeIpv6OspfPriorityCfg1009 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfRetransmitIntervalCfg1010 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfTransmitDelayCfg1011 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIsis1012 struct {
    Authentication InterfaceVeIsisAuthentication1013 `json:"authentication"`
    BfdCfg InterfaceVeIsisBfdCfg1017 `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceVeIsisCsnpIntervalList1018 `json:"csnp-interval-list"`
    Padding int `json:"padding" dval:"1"`
    HelloIntervalList []InterfaceVeIsisHelloIntervalList1019 `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceVeIsisHelloIntervalMinimalList1020 `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceVeIsisHelloMultiplierList1021 `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceVeIsisMeshGroup1022 `json:"mesh-group"`
    MetricList []InterfaceVeIsisMetricList1023 `json:"metric-list"`
    Network string `json:"network"`
    PasswordList []InterfaceVeIsisPasswordList1024 `json:"password-list"`
    PriorityList []InterfaceVeIsisPriorityList1025 `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    WideMetricList []InterfaceVeIsisWideMetricList1026 `json:"wide-metric-list"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIsisAuthentication1013 struct {
    SendOnlyList []InterfaceVeIsisAuthenticationSendOnlyList1014 `json:"send-only-list"`
    ModeList []InterfaceVeIsisAuthenticationModeList1015 `json:"mode-list"`
    KeyChainList []InterfaceVeIsisAuthenticationKeyChainList1016 `json:"key-chain-list"`
}


type InterfaceVeIsisAuthenticationSendOnlyList1014 struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceVeIsisAuthenticationModeList1015 struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceVeIsisAuthenticationKeyChainList1016 struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceVeIsisBfdCfg1017 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceVeIsisCsnpIntervalList1018 struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceVeIsisHelloIntervalList1019 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceVeIsisHelloIntervalMinimalList1020 struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceVeIsisHelloMultiplierList1021 struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceVeIsisMeshGroup1022 struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceVeIsisMetricList1023 struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceVeIsisPasswordList1024 struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceVeIsisPriorityList1025 struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceVeIsisWideMetricList1026 struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceVeLw4o61027 struct {
    Outside int `json:"outside"`
    Inside int `json:"inside"`
    Uuid string `json:"uuid"`
}


type InterfaceVeMap1028 struct {
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    MapTInside int `json:"map-t-inside"`
    MapTOutside int `json:"map-t-outside"`
    Uuid string `json:"uuid"`
}


type InterfaceVeNptv61029 struct {
    DomainList []InterfaceVeNptv6DomainList `json:"domain-list"`
}


type InterfaceVeNptv6DomainList struct {
    DomainName string `json:"domain-name"`
    BindType string `json:"bind-type"`
    Uuid string `json:"uuid"`
}


type InterfaceVeSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *InterfaceVe) GetId() string{
    return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceVe) getPath() string{
    return "interface/ve"
}

func (p *InterfaceVe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVe::Post")
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

func (p *InterfaceVe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVe::Get")
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
func (p *InterfaceVe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVe::Put")
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

func (p *InterfaceVe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVe::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
