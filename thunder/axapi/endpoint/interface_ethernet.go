

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernet struct {
	Inst struct {

    AccessList InterfaceEthernetAccessList `json:"access-list"`
    Action string `json:"action" dval:"disable"`
    AutoNegEnable int `json:"auto-neg-enable"`
    Bfd InterfaceEthernetBfd484 `json:"bfd"`
    CpuProcess int `json:"cpu-process"`
    CpuProcessDir string `json:"cpu-process-dir"`
    Ddos InterfaceEthernetDdos487 `json:"ddos"`
    Duplexity string `json:"duplexity" dval:"auto"`
    FecForcedOff int `json:"fec-forced-off"`
    FecForcedOn int `json:"fec-forced-on"`
    FlowControl int `json:"flow-control"`
    GamingProtocolCompliance int `json:"gaming-protocol-compliance"`
    IcmpRateLimit InterfaceEthernetIcmpRateLimit `json:"icmp-rate-limit"`
    Icmpv6RateLimit InterfaceEthernetIcmpv6RateLimit `json:"icmpv6-rate-limit"`
    Ifnum int `json:"ifnum"`
    Ip InterfaceEthernetIp488 `json:"ip"`
    IpgBitTime int `json:"ipg-bit-time" dval:"96"`
    Ipv6 InterfaceEthernetIpv6512 `json:"ipv6"`
    Isis InterfaceEthernetIsis535 `json:"isis"`
    L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`
    Lldp InterfaceEthernetLldp550 `json:"lldp"`
    LoadInterval int `json:"load-interval" dval:"300"`
    Lw4o6 InterfaceEthernetLw4o6555 `json:"lw-4o6"`
    MacLearning string `json:"mac-learning"`
    Map InterfaceEthernetMap556 `json:"map"`
    MediaTypeCopper int `json:"media-type-copper"`
    MonitorList []InterfaceEthernetMonitorList `json:"monitor-list"`
    Mtu int `json:"mtu"`
    Name string `json:"name"`
    Nptv6 InterfaceEthernetNptv6557 `json:"nptv6"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    PingSweepDetection string `json:"ping-sweep-detection" dval:"disable"`
    PortBreakout string `json:"port-breakout"`
    PortScanDetection string `json:"port-scan-detection" dval:"disable"`
    RemoveVlanTag int `json:"remove-vlan-tag"`
    SamplingEnable []InterfaceEthernetSamplingEnable `json:"sampling-enable"`
    SpanningTree InterfaceEthernetSpanningTree558 `json:"spanning-tree"`
    Speed string `json:"speed" dval:"auto"`
    SpeedForced10g int `json:"speed-forced-10g"`
    SpeedForced1g int `json:"speed-forced-1g"`
    SpeedForced40g int `json:"speed-forced-40g"`
    TrafficDistributionMode string `json:"traffic-distribution-mode"`
    TrapSource int `json:"trap-source"`
    TrunkGroupList []InterfaceEthernetTrunkGroupList `json:"trunk-group-list"`
    UpdateL2Info int `json:"update-l2-info"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VirtualWire int `json:"virtual-wire"`
    VlanLearning string `json:"vlan-learning"`

	} `json:"ethernet"`
}


type InterfaceEthernetAccessList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
}


type InterfaceEthernetBfd484 struct {
    Authentication InterfaceEthernetBfdAuthentication485 `json:"authentication"`
    Echo int `json:"echo"`
    Demand int `json:"demand"`
    IntervalCfg InterfaceEthernetBfdIntervalCfg486 `json:"interval-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetBfdAuthentication485 struct {
    KeyId int `json:"key-id"`
    Method string `json:"method"`
    Password string `json:"password"`
    Encrypted string `json:"encrypted"`
}


type InterfaceEthernetBfdIntervalCfg486 struct {
    Interval int `json:"interval"`
    MinRx int `json:"min-rx"`
    Multiplier int `json:"multiplier"`
}


type InterfaceEthernetDdos487 struct {
    Outside int `json:"outside"`
    Inside int `json:"inside"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIcmpRateLimit struct {
    Normal int `json:"normal"`
    Lockup int `json:"lockup"`
    LockupPeriod int `json:"lockup-period"`
}


type InterfaceEthernetIcmpv6RateLimit struct {
    NormalV6 int `json:"normal-v6"`
    LockupV6 int `json:"lockup-v6"`
    LockupPeriodV6 int `json:"lockup-period-v6"`
}


type InterfaceEthernetIp488 struct {
    Dhcp int `json:"dhcp"`
    AddressList []InterfaceEthernetIpAddressList489 `json:"address-list"`
    AllowPromiscuousVip int `json:"allow-promiscuous-vip"`
    CacheSpoofingPort int `json:"cache-spoofing-port"`
    HelperAddressList []InterfaceEthernetIpHelperAddressList490 `json:"helper-address-list"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    TtlIgnore int `json:"ttl-ignore"`
    SynCookie int `json:"syn-cookie"`
    SlbPartitionRedirect int `json:"slb-partition-redirect"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    QueryInterval int `json:"query-interval" dval:"125"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Client int `json:"client"`
    Server int `json:"server"`
    Unnumbered int `json:"unnumbered"`
    Uuid string `json:"uuid"`
    StatefulFirewall InterfaceEthernetIpStatefulFirewall491 `json:"stateful-firewall"`
    Router InterfaceEthernetIpRouter492 `json:"router"`
    Rip InterfaceEthernetIpRip494 `json:"rip"`
    Ospf InterfaceEthernetIpOspf502 `json:"ospf"`
}


type InterfaceEthernetIpAddressList489 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceEthernetIpHelperAddressList490 struct {
    HelperAddress string `json:"helper-address"`
}


type InterfaceEthernetIpStatefulFirewall491 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclId int `json:"acl-id"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpRouter492 struct {
    Isis InterfaceEthernetIpRouterIsis493 `json:"isis"`
}


type InterfaceEthernetIpRouterIsis493 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpRip494 struct {
    Authentication InterfaceEthernetIpRipAuthentication495 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceEthernetIpRipSendCfg499 `json:"send-cfg"`
    ReceiveCfg InterfaceEthernetIpRipReceiveCfg500 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceEthernetIpRipSplitHorizonCfg501 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpRipAuthentication495 struct {
    Str InterfaceEthernetIpRipAuthenticationStr496 `json:"str"`
    Mode InterfaceEthernetIpRipAuthenticationMode497 `json:"mode"`
    KeyChain InterfaceEthernetIpRipAuthenticationKeyChain498 `json:"key-chain"`
}


type InterfaceEthernetIpRipAuthenticationStr496 struct {
    String string `json:"string"`
}


type InterfaceEthernetIpRipAuthenticationMode497 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceEthernetIpRipAuthenticationKeyChain498 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceEthernetIpRipSendCfg499 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceEthernetIpRipReceiveCfg500 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceEthernetIpRipSplitHorizonCfg501 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceEthernetIpOspf502 struct {
    OspfGlobal InterfaceEthernetIpOspfOspfGlobal503 `json:"ospf-global"`
    OspfIpList []InterfaceEthernetIpOspfOspfIpList510 `json:"ospf-ip-list"`
}


type InterfaceEthernetIpOspfOspfGlobal503 struct {
    AuthenticationCfg InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg504 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceEthernetIpOspfOspfGlobalBfdCfg505 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg506 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg507 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceEthernetIpOspfOspfGlobalNetwork509 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg504 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceEthernetIpOspfOspfGlobalBfdCfg505 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg506 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg507 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5508 `json:"md5"`
}


type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5508 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceEthernetIpOspfOspfGlobalNetwork509 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceEthernetIpOspfOspfIpList510 struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceEthernetIpOspfOspfIpListMessageDigestCfg511 `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpOspfOspfIpListMessageDigestCfg511 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceEthernetIpv6512 struct {
    AddressList []InterfaceEthernetIpv6AddressList513 `json:"address-list"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    Ipv6Enable int `json:"ipv6-enable"`
    TtlIgnore int `json:"ttl-ignore"`
    AccessListCfg InterfaceEthernetIpv6AccessListCfg514 `json:"access-list-cfg"`
    RouterAdver InterfaceEthernetIpv6RouterAdver515 `json:"router-adver"`
    Uuid string `json:"uuid"`
    StatefulFirewall InterfaceEthernetIpv6StatefulFirewall517 `json:"stateful-firewall"`
    Router InterfaceEthernetIpv6Router518 `json:"router"`
    Rip InterfaceEthernetIpv6Rip523 `json:"rip"`
    Ospf InterfaceEthernetIpv6Ospf525 `json:"ospf"`
}


type InterfaceEthernetIpv6AddressList513 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
}


type InterfaceEthernetIpv6AccessListCfg514 struct {
    V6AclName string `json:"v6-acl-name"`
    Inbound int `json:"inbound"`
}


type InterfaceEthernetIpv6RouterAdver515 struct {
    Action string `json:"action" dval:"disable"`
    HopLimit int `json:"hop-limit" dval:"255"`
    MaxInterval int `json:"max-interval" dval:"600"`
    MinInterval int `json:"min-interval" dval:"200"`
    DefaultLifetime int `json:"default-lifetime" dval:"1800"`
    RateLimit int `json:"rate-limit" dval:"100000"`
    ReachableTime int `json:"reachable-time"`
    RetransmitTimer int `json:"retransmit-timer"`
    AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
    AdverMtu int `json:"adver-mtu"`
    PrefixList []InterfaceEthernetIpv6RouterAdverPrefixList516 `json:"prefix-list"`
    ManagedConfigAction string `json:"managed-config-action" dval:"disable"`
    OtherConfigAction string `json:"other-config-action" dval:"disable"`
    AdverVrid int `json:"adver-vrid"`
    UseFloatingIp int `json:"use-floating-ip"`
    FloatingIp string `json:"floating-ip"`
    AdverVridDefault int `json:"adver-vrid-default"`
    UseFloatingIpDefaultVrid int `json:"use-floating-ip-default-vrid"`
    FloatingIpDefaultVrid string `json:"floating-ip-default-vrid"`
}


type InterfaceEthernetIpv6RouterAdverPrefixList516 struct {
    Prefix string `json:"prefix"`
    NotAutonomous int `json:"not-autonomous"`
    NotOnLink int `json:"not-on-link"`
    PreferredLifetime int `json:"preferred-lifetime" dval:"604800"`
    ValidLifetime int `json:"valid-lifetime" dval:"2592000"`
}


type InterfaceEthernetIpv6StatefulFirewall517 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclName string `json:"acl-name"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6Router518 struct {
    Ripng InterfaceEthernetIpv6RouterRipng519 `json:"ripng"`
    Ospf InterfaceEthernetIpv6RouterOspf520 `json:"ospf"`
    Isis InterfaceEthernetIpv6RouterIsis522 `json:"isis"`
}


type InterfaceEthernetIpv6RouterRipng519 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6RouterOspf520 struct {
    AreaList []InterfaceEthernetIpv6RouterOspfAreaList521 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6RouterOspfAreaList521 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6RouterIsis522 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6Rip523 struct {
    SplitHorizonCfg InterfaceEthernetIpv6RipSplitHorizonCfg524 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6RipSplitHorizonCfg524 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceEthernetIpv6Ospf525 struct {
    NetworkList []InterfaceEthernetIpv6OspfNetworkList526 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceEthernetIpv6OspfCostCfg527 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceEthernetIpv6OspfDeadIntervalCfg528 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceEthernetIpv6OspfHelloIntervalCfg529 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceEthernetIpv6OspfMtuIgnoreCfg530 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceEthernetIpv6OspfNeighborCfg531 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceEthernetIpv6OspfPriorityCfg532 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceEthernetIpv6OspfRetransmitIntervalCfg533 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceEthernetIpv6OspfTransmitDelayCfg534 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6OspfNetworkList526 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceEthernetIpv6OspfCostCfg527 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfDeadIntervalCfg528 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfHelloIntervalCfg529 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfMtuIgnoreCfg530 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfNeighborCfg531 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceEthernetIpv6OspfPriorityCfg532 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfRetransmitIntervalCfg533 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfTransmitDelayCfg534 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIsis535 struct {
    Authentication InterfaceEthernetIsisAuthentication536 `json:"authentication"`
    BfdCfg InterfaceEthernetIsisBfdCfg540 `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceEthernetIsisCsnpIntervalList541 `json:"csnp-interval-list"`
    Padding int `json:"padding" dval:"1"`
    HelloIntervalList []InterfaceEthernetIsisHelloIntervalList542 `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceEthernetIsisHelloIntervalMinimalList543 `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceEthernetIsisHelloMultiplierList544 `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceEthernetIsisMeshGroup545 `json:"mesh-group"`
    MetricList []InterfaceEthernetIsisMetricList546 `json:"metric-list"`
    Network string `json:"network"`
    PasswordList []InterfaceEthernetIsisPasswordList547 `json:"password-list"`
    PriorityList []InterfaceEthernetIsisPriorityList548 `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    WideMetricList []InterfaceEthernetIsisWideMetricList549 `json:"wide-metric-list"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIsisAuthentication536 struct {
    SendOnlyList []InterfaceEthernetIsisAuthenticationSendOnlyList537 `json:"send-only-list"`
    ModeList []InterfaceEthernetIsisAuthenticationModeList538 `json:"mode-list"`
    KeyChainList []InterfaceEthernetIsisAuthenticationKeyChainList539 `json:"key-chain-list"`
}


type InterfaceEthernetIsisAuthenticationSendOnlyList537 struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisAuthenticationModeList538 struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisAuthenticationKeyChainList539 struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisBfdCfg540 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceEthernetIsisCsnpIntervalList541 struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisHelloIntervalList542 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisHelloIntervalMinimalList543 struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisHelloMultiplierList544 struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisMeshGroup545 struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceEthernetIsisMetricList546 struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisPasswordList547 struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisPriorityList548 struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceEthernetIsisWideMetricList549 struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceEthernetLldp550 struct {
    EnableCfg InterfaceEthernetLldpEnableCfg551 `json:"enable-cfg"`
    NotificationCfg InterfaceEthernetLldpNotificationCfg552 `json:"notification-cfg"`
    TxDot1Cfg InterfaceEthernetLldpTxDot1Cfg553 `json:"tx-dot1-cfg"`
    TxTlvsCfg InterfaceEthernetLldpTxTlvsCfg554 `json:"tx-tlvs-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetLldpEnableCfg551 struct {
    RtEnable int `json:"rt-enable"`
    Rx int `json:"rx"`
    Tx int `json:"tx"`
}


type InterfaceEthernetLldpNotificationCfg552 struct {
    Notification int `json:"notification"`
    NotifEnable int `json:"notif-enable"`
}


type InterfaceEthernetLldpTxDot1Cfg553 struct {
    TxDot1Tlvs int `json:"tx-dot1-tlvs"`
    LinkAggregation int `json:"link-aggregation"`
    Vlan int `json:"vlan"`
}


type InterfaceEthernetLldpTxTlvsCfg554 struct {
    TxTlvs int `json:"tx-tlvs"`
    Exclude int `json:"exclude"`
    ManagementAddress int `json:"management-address"`
    PortDescription int `json:"port-description"`
    SystemCapabilities int `json:"system-capabilities"`
    SystemDescription int `json:"system-description"`
    SystemName int `json:"system-name"`
}


type InterfaceEthernetLw4o6555 struct {
    Outside int `json:"outside"`
    Inside int `json:"inside"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetMap556 struct {
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    MapTInside int `json:"map-t-inside"`
    MapTOutside int `json:"map-t-outside"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetMonitorList struct {
    Monitor string `json:"monitor"`
    MirrorIndex int `json:"mirror-index"`
    MonitorVlan int `json:"monitor-vlan"`
}


type InterfaceEthernetNptv6557 struct {
    DomainList []InterfaceEthernetNptv6DomainList `json:"domain-list"`
}


type InterfaceEthernetNptv6DomainList struct {
    DomainName string `json:"domain-name"`
    BindType string `json:"bind-type"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type InterfaceEthernetSpanningTree558 struct {
    AutoEdge int `json:"auto-edge" dval:"1"`
    AdminEdge int `json:"admin-edge"`
    InstanceList []InterfaceEthernetSpanningTreeInstanceList559 `json:"instance-list"`
    PathCost int `json:"path-cost"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetSpanningTreeInstanceList559 struct {
    InstanceStart int `json:"instance-start"`
    MstpPathCost int `json:"mstp-path-cost"`
}


type InterfaceEthernetTrunkGroupList struct {
    TrunkNumber int `json:"trunk-number"`
    Type string `json:"type" dval:"static"`
    AdminKey int `json:"admin-key"`
    PortPriority int `json:"port-priority"`
    UdldTimeoutCfg InterfaceEthernetTrunkGroupListUdldTimeoutCfg `json:"udld-timeout-cfg"`
    Mode string `json:"mode" dval:"active"`
    Timeout string `json:"timeout" dval:"long"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type InterfaceEthernetTrunkGroupListUdldTimeoutCfg struct {
    Fast int `json:"fast" dval:"1000"`
    Slow int `json:"slow"`
}

func (p *InterfaceEthernet) GetId() string{
    return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceEthernet) getPath() string{
    return "interface/ethernet"
}

func (p *InterfaceEthernet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernet::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath()+"/"+strconv.Itoa(p.Inst.Ifnum), payloadBytes, headers, logger)
    return err
}

func (p *InterfaceEthernet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernet::Get")
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
func (p *InterfaceEthernet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernet::Put")
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

func (p *InterfaceEthernet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
