

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunk struct {
	Inst struct {

    AccessList InterfaceTrunkAccessList `json:"access-list"`
    Action string `json:"action" dval:"enable"`
    Bfd InterfaceTrunkBfd780 `json:"bfd"`
    Ddos InterfaceTrunkDdos784 `json:"ddos"`
    DoAutoRecovery int `json:"do-auto-recovery"`
    GamingProtocolCompliance int `json:"gaming-protocol-compliance"`
    IcmpRateLimit InterfaceTrunkIcmpRateLimit `json:"icmp-rate-limit"`
    Icmpv6RateLimit InterfaceTrunkIcmpv6RateLimit `json:"icmpv6-rate-limit"`
    Ifnum int `json:"ifnum"`
    Ip InterfaceTrunkIp785 `json:"ip"`
    Ipv6 InterfaceTrunkIpv6810 `json:"ipv6"`
    Isis InterfaceTrunkIsis836 `json:"isis"`
    L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`
    Lw4o6 InterfaceTrunkLw4o6851 `json:"lw-4o6"`
    MacLearning string `json:"mac-learning"`
    Map InterfaceTrunkMap852 `json:"map"`
    Mtu int `json:"mtu"`
    Name string `json:"name"`
    Nptv6 InterfaceTrunkNptv6853 `json:"nptv6"`
    PortsThreshold int `json:"ports-threshold"`
    SamplingEnable []InterfaceTrunkSamplingEnable `json:"sampling-enable"`
    SpanningTree InterfaceTrunkSpanningTree854 `json:"spanning-tree"`
    SyncModifyDisable int `json:"sync-modify-disable"`
    Timer int `json:"timer" dval:"10"`
    TrapSource int `json:"trap-source"`
    UpdateL2Info int `json:"update-l2-info"`
    UseHwHash int `json:"use-hw-hash"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VirtualWire int `json:"virtual-wire"`
    VlanLearning string `json:"vlan-learning"`

	} `json:"trunk"`
}


type InterfaceTrunkAccessList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
}


type InterfaceTrunkBfd780 struct {
    Authentication InterfaceTrunkBfdAuthentication781 `json:"authentication"`
    Echo int `json:"echo"`
    Demand int `json:"demand"`
    IntervalCfg InterfaceTrunkBfdIntervalCfg782 `json:"interval-cfg"`
    Uuid string `json:"uuid"`
    PerMemberPort InterfaceTrunkBfdPerMemberPort783 `json:"per-member-port"`
}


type InterfaceTrunkBfdAuthentication781 struct {
    KeyId int `json:"key-id"`
    Method string `json:"method"`
    Password string `json:"password"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTrunkBfdIntervalCfg782 struct {
    Interval int `json:"interval"`
    MinRx int `json:"min-rx"`
    Multiplier int `json:"multiplier"`
}


type InterfaceTrunkBfdPerMemberPort783 struct {
    LocalAddress string `json:"local-address"`
    NeighborAddress string `json:"neighbor-address"`
    Ipv6Local string `json:"ipv6-local"`
    Ipv6Nbr string `json:"ipv6-nbr"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkDdos784 struct {
    Outside int `json:"outside"`
    Inside int `json:"inside"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIcmpRateLimit struct {
    Normal int `json:"normal"`
    Lockup int `json:"lockup"`
    LockupPeriod int `json:"lockup-period"`
}


type InterfaceTrunkIcmpv6RateLimit struct {
    NormalV6 int `json:"normal-v6"`
    LockupV6 int `json:"lockup-v6"`
    LockupPeriodV6 int `json:"lockup-period-v6"`
}


type InterfaceTrunkIp785 struct {
    Dhcp int `json:"dhcp"`
    AddressList []InterfaceTrunkIpAddressList786 `json:"address-list"`
    AllowPromiscuousVip int `json:"allow-promiscuous-vip"`
    Client int `json:"client"`
    Server int `json:"server"`
    CacheSpoofingPort int `json:"cache-spoofing-port"`
    HelperAddressList []InterfaceTrunkIpHelperAddressList787 `json:"helper-address-list"`
    Nat InterfaceTrunkIpNat788 `json:"nat"`
    TtlIgnore int `json:"ttl-ignore"`
    SynCookie int `json:"syn-cookie"`
    SlbPartitionRedirect int `json:"slb-partition-redirect"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    QueryInterval int `json:"query-interval" dval:"125"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Unnumbered int `json:"unnumbered"`
    Uuid string `json:"uuid"`
    StatefulFirewall InterfaceTrunkIpStatefulFirewall789 `json:"stateful-firewall"`
    Router InterfaceTrunkIpRouter790 `json:"router"`
    Rip InterfaceTrunkIpRip792 `json:"rip"`
    Ospf InterfaceTrunkIpOspf800 `json:"ospf"`
}


type InterfaceTrunkIpAddressList786 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceTrunkIpHelperAddressList787 struct {
    HelperAddress string `json:"helper-address"`
}


type InterfaceTrunkIpNat788 struct {
    Inside int `json:"inside"`
    Outside int `json:"outside"`
}


type InterfaceTrunkIpStatefulFirewall789 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclId int `json:"acl-id"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpRouter790 struct {
    Isis InterfaceTrunkIpRouterIsis791 `json:"isis"`
}


type InterfaceTrunkIpRouterIsis791 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpRip792 struct {
    Authentication InterfaceTrunkIpRipAuthentication793 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceTrunkIpRipSendCfg797 `json:"send-cfg"`
    ReceiveCfg InterfaceTrunkIpRipReceiveCfg798 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceTrunkIpRipSplitHorizonCfg799 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpRipAuthentication793 struct {
    Str InterfaceTrunkIpRipAuthenticationStr794 `json:"str"`
    Mode InterfaceTrunkIpRipAuthenticationMode795 `json:"mode"`
    KeyChain InterfaceTrunkIpRipAuthenticationKeyChain796 `json:"key-chain"`
}


type InterfaceTrunkIpRipAuthenticationStr794 struct {
    String string `json:"string"`
}


type InterfaceTrunkIpRipAuthenticationMode795 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceTrunkIpRipAuthenticationKeyChain796 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceTrunkIpRipSendCfg797 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceTrunkIpRipReceiveCfg798 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceTrunkIpRipSplitHorizonCfg799 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceTrunkIpOspf800 struct {
    OspfGlobal InterfaceTrunkIpOspfOspfGlobal801 `json:"ospf-global"`
    OspfIpList []InterfaceTrunkIpOspfOspfIpList808 `json:"ospf-ip-list"`
}


type InterfaceTrunkIpOspfOspfGlobal801 struct {
    AuthenticationCfg InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg802 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceTrunkIpOspfOspfGlobalBfdCfg803 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg804 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg805 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceTrunkIpOspfOspfGlobalNetwork807 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg802 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceTrunkIpOspfOspfGlobalBfdCfg803 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg804 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg805 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5806 `json:"md5"`
}


type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5806 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTrunkIpOspfOspfGlobalNetwork807 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceTrunkIpOspfOspfIpList808 struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTrunkIpOspfOspfIpListMessageDigestCfg809 `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpOspfOspfIpListMessageDigestCfg809 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTrunkIpv6810 struct {
    AddressList []InterfaceTrunkIpv6AddressList811 `json:"address-list"`
    Ipv6Enable int `json:"ipv6-enable"`
    AccessListCfg InterfaceTrunkIpv6AccessListCfg812 `json:"access-list-cfg"`
    Nat InterfaceTrunkIpv6Nat813 `json:"nat"`
    TtlIgnore int `json:"ttl-ignore"`
    RouterAdver InterfaceTrunkIpv6RouterAdver814 `json:"router-adver"`
    Uuid string `json:"uuid"`
    StatefulFirewall InterfaceTrunkIpv6StatefulFirewall818 `json:"stateful-firewall"`
    Router InterfaceTrunkIpv6Router819 `json:"router"`
    Rip InterfaceTrunkIpv6Rip824 `json:"rip"`
    Ospf InterfaceTrunkIpv6Ospf826 `json:"ospf"`
}


type InterfaceTrunkIpv6AddressList811 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
}


type InterfaceTrunkIpv6AccessListCfg812 struct {
    V6AclName string `json:"v6-acl-name"`
    Inbound int `json:"inbound"`
}


type InterfaceTrunkIpv6Nat813 struct {
    Inside int `json:"inside"`
    Outside int `json:"outside"`
}


type InterfaceTrunkIpv6RouterAdver814 struct {
    Action string `json:"action" dval:"disable"`
    DefaultLifetime int `json:"default-lifetime" dval:"1800"`
    HopLimit int `json:"hop-limit" dval:"255"`
    MaxInterval int `json:"max-interval" dval:"600"`
    MinInterval int `json:"min-interval" dval:"200"`
    RateLimit int `json:"rate-limit" dval:"100000"`
    ReachableTime int `json:"reachable-time"`
    RetransmitTimer int `json:"retransmit-timer"`
    Mtu InterfaceTrunkIpv6RouterAdverMtu815 `json:"mtu"`
    PrefixList []InterfaceTrunkIpv6RouterAdverPrefixList816 `json:"prefix-list"`
    ManagedConfigAction string `json:"managed-config-action" dval:"disable"`
    OtherConfigAction string `json:"other-config-action" dval:"disable"`
    Vrid InterfaceTrunkIpv6RouterAdverVrid817 `json:"vrid"`
}


type InterfaceTrunkIpv6RouterAdverMtu815 struct {
    AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
    AdverMtu int `json:"adver-mtu"`
}


type InterfaceTrunkIpv6RouterAdverPrefixList816 struct {
    Prefix string `json:"prefix"`
    NotAutonomous int `json:"not-autonomous"`
    NotOnLink int `json:"not-on-link"`
    PreferredLifetime int `json:"preferred-lifetime" dval:"604800"`
    ValidLifetime int `json:"valid-lifetime" dval:"2592000"`
}


type InterfaceTrunkIpv6RouterAdverVrid817 struct {
    AdverVrid int `json:"adver-vrid"`
    UseFloatingIp int `json:"use-floating-ip"`
    FloatingIp string `json:"floating-ip"`
    AdverVridDefault int `json:"adver-vrid-default"`
    UseFloatingIpDefaultVrid int `json:"use-floating-ip-default-vrid"`
    FloatingIpDefaultVrid string `json:"floating-ip-default-vrid"`
}


type InterfaceTrunkIpv6StatefulFirewall818 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclName string `json:"acl-name"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpv6Router819 struct {
    Ripng InterfaceTrunkIpv6RouterRipng820 `json:"ripng"`
    Ospf InterfaceTrunkIpv6RouterOspf821 `json:"ospf"`
    Isis InterfaceTrunkIpv6RouterIsis823 `json:"isis"`
}


type InterfaceTrunkIpv6RouterRipng820 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpv6RouterOspf821 struct {
    AreaList []InterfaceTrunkIpv6RouterOspfAreaList822 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpv6RouterOspfAreaList822 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIpv6RouterIsis823 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpv6Rip824 struct {
    SplitHorizonCfg InterfaceTrunkIpv6RipSplitHorizonCfg825 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpv6RipSplitHorizonCfg825 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceTrunkIpv6Ospf826 struct {
    NetworkList []InterfaceTrunkIpv6OspfNetworkList827 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceTrunkIpv6OspfCostCfg828 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceTrunkIpv6OspfDeadIntervalCfg829 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceTrunkIpv6OspfHelloIntervalCfg830 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceTrunkIpv6OspfMtuIgnoreCfg831 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceTrunkIpv6OspfNeighborCfg832 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceTrunkIpv6OspfPriorityCfg833 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceTrunkIpv6OspfRetransmitIntervalCfg834 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceTrunkIpv6OspfTransmitDelayCfg835 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpv6OspfNetworkList827 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceTrunkIpv6OspfCostCfg828 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIpv6OspfDeadIntervalCfg829 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIpv6OspfHelloIntervalCfg830 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIpv6OspfMtuIgnoreCfg831 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIpv6OspfNeighborCfg832 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceTrunkIpv6OspfPriorityCfg833 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIpv6OspfRetransmitIntervalCfg834 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIpv6OspfTransmitDelayCfg835 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTrunkIsis836 struct {
    Authentication InterfaceTrunkIsisAuthentication837 `json:"authentication"`
    BfdCfg InterfaceTrunkIsisBfdCfg841 `json:"bfd-cfg"`
    CircuitType string `json:"circuit-type" dval:"level-1-2"`
    CsnpIntervalList []InterfaceTrunkIsisCsnpIntervalList842 `json:"csnp-interval-list"`
    Padding int `json:"padding" dval:"1"`
    HelloIntervalList []InterfaceTrunkIsisHelloIntervalList843 `json:"hello-interval-list"`
    HelloIntervalMinimalList []InterfaceTrunkIsisHelloIntervalMinimalList844 `json:"hello-interval-minimal-list"`
    HelloMultiplierList []InterfaceTrunkIsisHelloMultiplierList845 `json:"hello-multiplier-list"`
    LspInterval int `json:"lsp-interval" dval:"33"`
    MeshGroup InterfaceTrunkIsisMeshGroup846 `json:"mesh-group"`
    MetricList []InterfaceTrunkIsisMetricList847 `json:"metric-list"`
    Network string `json:"network"`
    PasswordList []InterfaceTrunkIsisPasswordList848 `json:"password-list"`
    PriorityList []InterfaceTrunkIsisPriorityList849 `json:"priority-list"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    WideMetricList []InterfaceTrunkIsisWideMetricList850 `json:"wide-metric-list"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIsisAuthentication837 struct {
    SendOnlyList []InterfaceTrunkIsisAuthenticationSendOnlyList838 `json:"send-only-list"`
    ModeList []InterfaceTrunkIsisAuthenticationModeList839 `json:"mode-list"`
    KeyChainList []InterfaceTrunkIsisAuthenticationKeyChainList840 `json:"key-chain-list"`
}


type InterfaceTrunkIsisAuthenticationSendOnlyList838 struct {
    SendOnly int `json:"send-only"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisAuthenticationModeList839 struct {
    Mode string `json:"mode"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisAuthenticationKeyChainList840 struct {
    KeyChain string `json:"key-chain"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisBfdCfg841 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceTrunkIsisCsnpIntervalList842 struct {
    CsnpInterval int `json:"csnp-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisHelloIntervalList843 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisHelloIntervalMinimalList844 struct {
    HelloIntervalMinimal int `json:"hello-interval-minimal"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisHelloMultiplierList845 struct {
    HelloMultiplier int `json:"hello-multiplier" dval:"3"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisMeshGroup846 struct {
    Value int `json:"value"`
    Blocked int `json:"blocked"`
}


type InterfaceTrunkIsisMetricList847 struct {
    Metric int `json:"metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisPasswordList848 struct {
    Password string `json:"password"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisPriorityList849 struct {
    Priority int `json:"priority" dval:"64"`
    Level string `json:"level"`
}


type InterfaceTrunkIsisWideMetricList850 struct {
    WideMetric int `json:"wide-metric" dval:"10"`
    Level string `json:"level"`
}


type InterfaceTrunkLw4o6851 struct {
    Outside int `json:"outside"`
    Inside int `json:"inside"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkMap852 struct {
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    MapTInside int `json:"map-t-inside"`
    MapTOutside int `json:"map-t-outside"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkNptv6853 struct {
    DomainList []InterfaceTrunkNptv6DomainList `json:"domain-list"`
}


type InterfaceTrunkNptv6DomainList struct {
    DomainName string `json:"domain-name"`
    BindType string `json:"bind-type"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type InterfaceTrunkSpanningTree854 struct {
    AutoEdge int `json:"auto-edge" dval:"1"`
    AdminEdge int `json:"admin-edge"`
    InstanceList []InterfaceTrunkSpanningTreeInstanceList855 `json:"instance-list"`
    PathCost int `json:"path-cost"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkSpanningTreeInstanceList855 struct {
    InstanceStart int `json:"instance-start"`
    MstpPathCost int `json:"mstp-path-cost"`
}

func (p *InterfaceTrunk) GetId() string{
    return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceTrunk) getPath() string{
    return "interface/trunk"
}

func (p *InterfaceTrunk) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunk::Post")
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

func (p *InterfaceTrunk) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunk::Get")
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
func (p *InterfaceTrunk) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunk::Put")
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

func (p *InterfaceTrunk) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunk::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
