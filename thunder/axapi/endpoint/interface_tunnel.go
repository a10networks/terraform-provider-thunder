

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTunnel struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    Ifnum int `json:"ifnum"`
    Ip InterfaceTunnelIp886 `json:"ip"`
    Ipv6 InterfaceTunnelIpv6907 `json:"ipv6"`
    LoadInterval int `json:"load-interval" dval:"300"`
    Lw4o6 InterfaceTunnelLw4o6923 `json:"lw-4o6"`
    Map InterfaceTunnelMap924 `json:"map"`
    Mtu int `json:"mtu"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    SamplingEnable []InterfaceTunnelSamplingEnable `json:"sampling-enable"`
    SecurityLevel int `json:"security-level"`
    Speed int `json:"speed" dval:"10"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"tunnel"`
}


type InterfaceTunnelIp886 struct {
    Address InterfaceTunnelIpAddress887 `json:"address"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    GenerateMembershipQueryVal int `json:"generate-membership-query-val" dval:"125"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    Uuid string `json:"uuid"`
    Rip InterfaceTunnelIpRip889 `json:"rip"`
    Ospf InterfaceTunnelIpOspf897 `json:"ospf"`
}


type InterfaceTunnelIpAddress887 struct {
    Dhcp int `json:"dhcp"`
    IpCfg []InterfaceTunnelIpAddressIpCfg888 `json:"ip-cfg"`
}


type InterfaceTunnelIpAddressIpCfg888 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceTunnelIpRip889 struct {
    Authentication InterfaceTunnelIpRipAuthentication890 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceTunnelIpRipSendCfg894 `json:"send-cfg"`
    ReceiveCfg InterfaceTunnelIpRipReceiveCfg895 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceTunnelIpRipSplitHorizonCfg896 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpRipAuthentication890 struct {
    Str InterfaceTunnelIpRipAuthenticationStr891 `json:"str"`
    Mode InterfaceTunnelIpRipAuthenticationMode892 `json:"mode"`
    KeyChain InterfaceTunnelIpRipAuthenticationKeyChain893 `json:"key-chain"`
}


type InterfaceTunnelIpRipAuthenticationStr891 struct {
    String string `json:"string"`
}


type InterfaceTunnelIpRipAuthenticationMode892 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceTunnelIpRipAuthenticationKeyChain893 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceTunnelIpRipSendCfg894 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceTunnelIpRipReceiveCfg895 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceTunnelIpRipSplitHorizonCfg896 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceTunnelIpOspf897 struct {
    OspfGlobal InterfaceTunnelIpOspfOspfGlobal898 `json:"ospf-global"`
    OspfIpList []InterfaceTunnelIpOspfOspfIpList905 `json:"ospf-ip-list"`
}


type InterfaceTunnelIpOspfOspfGlobal898 struct {
    AuthenticationCfg InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg899 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceTunnelIpOspfOspfGlobalBfdCfg900 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg901 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg902 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceTunnelIpOspfOspfGlobalNetwork904 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg899 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceTunnelIpOspfOspfGlobalBfdCfg900 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg901 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg902 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5903 `json:"md5"`
}


type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5903 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTunnelIpOspfOspfGlobalNetwork904 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceTunnelIpOspfOspfIpList905 struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTunnelIpOspfOspfIpListMessageDigestCfg906 `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpOspfOspfIpListMessageDigestCfg906 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTunnelIpv6907 struct {
    AddressCfg []InterfaceTunnelIpv6AddressCfg908 `json:"address-cfg"`
    Ipv6Enable int `json:"ipv6-enable"`
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    Uuid string `json:"uuid"`
    Router InterfaceTunnelIpv6Router909 `json:"router"`
    Ospf InterfaceTunnelIpv6Ospf913 `json:"ospf"`
}


type InterfaceTunnelIpv6AddressCfg908 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
}


type InterfaceTunnelIpv6Router909 struct {
    Ripng InterfaceTunnelIpv6RouterRipng910 `json:"ripng"`
    Ospf InterfaceTunnelIpv6RouterOspf911 `json:"ospf"`
}


type InterfaceTunnelIpv6RouterRipng910 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpv6RouterOspf911 struct {
    AreaList []InterfaceTunnelIpv6RouterOspfAreaList912 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpv6RouterOspfAreaList912 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6Ospf913 struct {
    NetworkList []InterfaceTunnelIpv6OspfNetworkList914 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceTunnelIpv6OspfCostCfg915 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceTunnelIpv6OspfDeadIntervalCfg916 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceTunnelIpv6OspfHelloIntervalCfg917 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceTunnelIpv6OspfMtuIgnoreCfg918 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceTunnelIpv6OspfNeighborCfg919 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceTunnelIpv6OspfPriorityCfg920 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceTunnelIpv6OspfRetransmitIntervalCfg921 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceTunnelIpv6OspfTransmitDelayCfg922 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpv6OspfNetworkList914 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceTunnelIpv6OspfCostCfg915 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfDeadIntervalCfg916 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfHelloIntervalCfg917 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfMtuIgnoreCfg918 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfNeighborCfg919 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceTunnelIpv6OspfPriorityCfg920 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfRetransmitIntervalCfg921 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfTransmitDelayCfg922 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelLw4o6923 struct {
    Outside int `json:"outside"`
    Inside int `json:"inside"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelMap924 struct {
    Inside int `json:"inside"`
    Outside int `json:"outside"`
    MapTInside int `json:"map-t-inside"`
    MapTOutside int `json:"map-t-outside"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *InterfaceTunnel) GetId() string{
    return strconv.Itoa(p.Inst.Ifnum)
}

func (p *InterfaceTunnel) getPath() string{
    return "interface/tunnel"
}

func (p *InterfaceTunnel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnel::Post")
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

func (p *InterfaceTunnel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnel::Get")
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
func (p *InterfaceTunnel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnel::Put")
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

func (p *InterfaceTunnel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnel::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
