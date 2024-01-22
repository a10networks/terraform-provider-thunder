

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeIp struct {
	Inst struct {

    AddressList []InterfaceVeIpAddressList `json:"address-list"`
    AllowPromiscuousVip int `json:"allow-promiscuous-vip"`
    Client int `json:"client"`
    Dhcp int `json:"dhcp"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    HelperAddressList []InterfaceVeIpHelperAddressList `json:"helper-address-list"`
    Inside int `json:"inside"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Ospf InterfaceVeIpOspf925 `json:"ospf"`
    Outside int `json:"outside"`
    QueryInterval int `json:"query-interval" dval:"125"`
    Rip InterfaceVeIpRip933 `json:"rip"`
    Router InterfaceVeIpRouter941 `json:"router"`
    Server int `json:"server"`
    SlbPartitionRedirect int `json:"slb-partition-redirect"`
    StatefulFirewall InterfaceVeIpStatefulFirewall943 `json:"stateful-firewall"`
    SynCookie int `json:"syn-cookie"`
    TtlIgnore int `json:"ttl-ignore"`
    Unnumbered int `json:"unnumbered"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ip"`
}


type InterfaceVeIpAddressList struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceVeIpHelperAddressList struct {
    HelperAddress string `json:"helper-address"`
}


type InterfaceVeIpOspf925 struct {
    OspfGlobal InterfaceVeIpOspfOspfGlobal926 `json:"ospf-global"`
    OspfIpList []InterfaceVeIpOspfOspfIpList `json:"ospf-ip-list"`
}


type InterfaceVeIpOspfOspfGlobal926 struct {
    AuthenticationCfg InterfaceVeIpOspfOspfGlobalAuthenticationCfg927 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceVeIpOspfOspfGlobalBfdCfg928 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg929 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceVeIpOspfOspfGlobalMessageDigestCfg930 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceVeIpOspfOspfGlobalNetwork932 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpOspfOspfGlobalAuthenticationCfg927 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceVeIpOspfOspfGlobalBfdCfg928 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceVeIpOspfOspfGlobalDatabaseFilterCfg929 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceVeIpOspfOspfGlobalMessageDigestCfg930 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5931 `json:"md5"`
}


type InterfaceVeIpOspfOspfGlobalMessageDigestCfgMd5931 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceVeIpOspfOspfGlobalNetwork932 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceVeIpOspfOspfIpList struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceVeIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpOspfOspfIpListMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceVeIpRip933 struct {
    Authentication InterfaceVeIpRipAuthentication934 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceVeIpRipSendCfg938 `json:"send-cfg"`
    ReceiveCfg InterfaceVeIpRipReceiveCfg939 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceVeIpRipSplitHorizonCfg940 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpRipAuthentication934 struct {
    Str InterfaceVeIpRipAuthenticationStr935 `json:"str"`
    Mode InterfaceVeIpRipAuthenticationMode936 `json:"mode"`
    KeyChain InterfaceVeIpRipAuthenticationKeyChain937 `json:"key-chain"`
}


type InterfaceVeIpRipAuthenticationStr935 struct {
    String string `json:"string"`
}


type InterfaceVeIpRipAuthenticationMode936 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceVeIpRipAuthenticationKeyChain937 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceVeIpRipSendCfg938 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceVeIpRipReceiveCfg939 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceVeIpRipSplitHorizonCfg940 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceVeIpRouter941 struct {
    Isis InterfaceVeIpRouterIsis942 `json:"isis"`
}


type InterfaceVeIpRouterIsis942 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpStatefulFirewall943 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclId int `json:"acl-id"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceVeIp) GetId() string{
    return "1"
}

func (p *InterfaceVeIp) getPath() string{
    return "interface/ve/" +p.Inst.Ifnum + "/ip"
}

func (p *InterfaceVeIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIp::Post")
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

func (p *InterfaceVeIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIp::Get")
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
func (p *InterfaceVeIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIp::Put")
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

func (p *InterfaceVeIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
