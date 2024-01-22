

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetIp struct {
	Inst struct {

    AddressList []InterfaceEthernetIpAddressList `json:"address-list"`
    AllowPromiscuousVip int `json:"allow-promiscuous-vip"`
    CacheSpoofingPort int `json:"cache-spoofing-port"`
    Client int `json:"client"`
    Dhcp int `json:"dhcp"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    HelperAddressList []InterfaceEthernetIpHelperAddressList `json:"helper-address-list"`
    Inside int `json:"inside"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Ospf InterfaceEthernetIpOspf447 `json:"ospf"`
    Outside int `json:"outside"`
    QueryInterval int `json:"query-interval" dval:"125"`
    Rip InterfaceEthernetIpRip455 `json:"rip"`
    Router InterfaceEthernetIpRouter463 `json:"router"`
    Server int `json:"server"`
    SlbPartitionRedirect int `json:"slb-partition-redirect"`
    StatefulFirewall InterfaceEthernetIpStatefulFirewall465 `json:"stateful-firewall"`
    SynCookie int `json:"syn-cookie"`
    TtlIgnore int `json:"ttl-ignore"`
    Unnumbered int `json:"unnumbered"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ip"`
}


type InterfaceEthernetIpAddressList struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceEthernetIpHelperAddressList struct {
    HelperAddress string `json:"helper-address"`
}


type InterfaceEthernetIpOspf447 struct {
    OspfGlobal InterfaceEthernetIpOspfOspfGlobal448 `json:"ospf-global"`
    OspfIpList []InterfaceEthernetIpOspfOspfIpList `json:"ospf-ip-list"`
}


type InterfaceEthernetIpOspfOspfGlobal448 struct {
    AuthenticationCfg InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg449 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceEthernetIpOspfOspfGlobalBfdCfg450 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg451 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg452 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceEthernetIpOspfOspfGlobalNetwork454 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg449 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceEthernetIpOspfOspfGlobalBfdCfg450 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg451 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg452 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5453 `json:"md5"`
}


type InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5453 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceEthernetIpOspfOspfGlobalNetwork454 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceEthernetIpOspfOspfIpList struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceEthernetIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpOspfOspfIpListMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceEthernetIpRip455 struct {
    Authentication InterfaceEthernetIpRipAuthentication456 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceEthernetIpRipSendCfg460 `json:"send-cfg"`
    ReceiveCfg InterfaceEthernetIpRipReceiveCfg461 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceEthernetIpRipSplitHorizonCfg462 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpRipAuthentication456 struct {
    Str InterfaceEthernetIpRipAuthenticationStr457 `json:"str"`
    Mode InterfaceEthernetIpRipAuthenticationMode458 `json:"mode"`
    KeyChain InterfaceEthernetIpRipAuthenticationKeyChain459 `json:"key-chain"`
}


type InterfaceEthernetIpRipAuthenticationStr457 struct {
    String string `json:"string"`
}


type InterfaceEthernetIpRipAuthenticationMode458 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceEthernetIpRipAuthenticationKeyChain459 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceEthernetIpRipSendCfg460 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceEthernetIpRipReceiveCfg461 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceEthernetIpRipSplitHorizonCfg462 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceEthernetIpRouter463 struct {
    Isis InterfaceEthernetIpRouterIsis464 `json:"isis"`
}


type InterfaceEthernetIpRouterIsis464 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpStatefulFirewall465 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclId int `json:"acl-id"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceEthernetIp) GetId() string{
    return "1"
}

func (p *InterfaceEthernetIp) getPath() string{
    return "interface/ethernet/" +p.Inst.Ifnum + "/ip"
}

func (p *InterfaceEthernetIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIp::Post")
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

func (p *InterfaceEthernetIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIp::Get")
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
func (p *InterfaceEthernetIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIp::Put")
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

func (p *InterfaceEthernetIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
