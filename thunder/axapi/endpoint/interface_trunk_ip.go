

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkIp struct {
	Inst struct {

    AddressList []InterfaceTrunkIpAddressList `json:"address-list"`
    AllowPromiscuousVip int `json:"allow-promiscuous-vip"`
    CacheSpoofingPort int `json:"cache-spoofing-port"`
    Client int `json:"client"`
    Dhcp int `json:"dhcp"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    HelperAddressList []InterfaceTrunkIpHelperAddressList `json:"helper-address-list"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Nat InterfaceTrunkIpNat `json:"nat"`
    Ospf InterfaceTrunkIpOspf743 `json:"ospf"`
    QueryInterval int `json:"query-interval" dval:"125"`
    Rip InterfaceTrunkIpRip751 `json:"rip"`
    Router InterfaceTrunkIpRouter759 `json:"router"`
    Server int `json:"server"`
    SlbPartitionRedirect int `json:"slb-partition-redirect"`
    StatefulFirewall InterfaceTrunkIpStatefulFirewall761 `json:"stateful-firewall"`
    SynCookie int `json:"syn-cookie"`
    TtlIgnore int `json:"ttl-ignore"`
    Unnumbered int `json:"unnumbered"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ip"`
}


type InterfaceTrunkIpAddressList struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceTrunkIpHelperAddressList struct {
    HelperAddress string `json:"helper-address"`
}


type InterfaceTrunkIpNat struct {
    Inside int `json:"inside"`
    Outside int `json:"outside"`
}


type InterfaceTrunkIpOspf743 struct {
    OspfGlobal InterfaceTrunkIpOspfOspfGlobal744 `json:"ospf-global"`
    OspfIpList []InterfaceTrunkIpOspfOspfIpList `json:"ospf-ip-list"`
}


type InterfaceTrunkIpOspfOspfGlobal744 struct {
    AuthenticationCfg InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg745 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceTrunkIpOspfOspfGlobalBfdCfg746 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg747 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg748 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceTrunkIpOspfOspfGlobalNetwork750 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg745 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceTrunkIpOspfOspfGlobalBfdCfg746 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg747 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg748 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5749 `json:"md5"`
}


type InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5749 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTrunkIpOspfOspfGlobalNetwork750 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceTrunkIpOspfOspfIpList struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTrunkIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpOspfOspfIpListMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTrunkIpRip751 struct {
    Authentication InterfaceTrunkIpRipAuthentication752 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceTrunkIpRipSendCfg756 `json:"send-cfg"`
    ReceiveCfg InterfaceTrunkIpRipReceiveCfg757 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceTrunkIpRipSplitHorizonCfg758 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpRipAuthentication752 struct {
    Str InterfaceTrunkIpRipAuthenticationStr753 `json:"str"`
    Mode InterfaceTrunkIpRipAuthenticationMode754 `json:"mode"`
    KeyChain InterfaceTrunkIpRipAuthenticationKeyChain755 `json:"key-chain"`
}


type InterfaceTrunkIpRipAuthenticationStr753 struct {
    String string `json:"string"`
}


type InterfaceTrunkIpRipAuthenticationMode754 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceTrunkIpRipAuthenticationKeyChain755 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceTrunkIpRipSendCfg756 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceTrunkIpRipReceiveCfg757 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceTrunkIpRipSplitHorizonCfg758 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceTrunkIpRouter759 struct {
    Isis InterfaceTrunkIpRouterIsis760 `json:"isis"`
}


type InterfaceTrunkIpRouterIsis760 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceTrunkIpStatefulFirewall761 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclId int `json:"acl-id"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceTrunkIp) GetId() string{
    return "1"
}

func (p *InterfaceTrunkIp) getPath() string{
    return "interface/trunk/" +p.Inst.Ifnum + "/ip"
}

func (p *InterfaceTrunkIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIp::Post")
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

func (p *InterfaceTrunkIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIp::Get")
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
func (p *InterfaceTrunkIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIp::Put")
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

func (p *InterfaceTrunkIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
