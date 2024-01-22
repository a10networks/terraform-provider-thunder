

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTunnelIp struct {
	Inst struct {

    Address InterfaceTunnelIpAddress `json:"address"`
    GenerateMembershipQuery int `json:"generate-membership-query"`
    GenerateMembershipQueryVal int `json:"generate-membership-query-val" dval:"125"`
    Inside int `json:"inside"`
    MaxRespTime int `json:"max-resp-time" dval:"100"`
    Ospf InterfaceTunnelIpOspf856 `json:"ospf"`
    Outside int `json:"outside"`
    Rip InterfaceTunnelIpRip864 `json:"rip"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ip"`
}


type InterfaceTunnelIpAddress struct {
    Dhcp int `json:"dhcp"`
    IpCfg []InterfaceTunnelIpAddressIpCfg `json:"ip-cfg"`
}


type InterfaceTunnelIpAddressIpCfg struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceTunnelIpOspf856 struct {
    OspfGlobal InterfaceTunnelIpOspfOspfGlobal857 `json:"ospf-global"`
    OspfIpList []InterfaceTunnelIpOspfOspfIpList `json:"ospf-ip-list"`
}


type InterfaceTunnelIpOspfOspfGlobal857 struct {
    AuthenticationCfg InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg858 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceTunnelIpOspfOspfGlobalBfdCfg859 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg860 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg861 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceTunnelIpOspfOspfGlobalNetwork863 `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg858 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceTunnelIpOspfOspfGlobalBfdCfg859 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg860 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg861 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5862 `json:"md5"`
}


type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5862 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTunnelIpOspfOspfGlobalNetwork863 struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}


type InterfaceTunnelIpOspfOspfIpList struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTunnelIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpOspfOspfIpListMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTunnelIpRip864 struct {
    Authentication InterfaceTunnelIpRipAuthentication865 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceTunnelIpRipSendCfg869 `json:"send-cfg"`
    ReceiveCfg InterfaceTunnelIpRipReceiveCfg870 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceTunnelIpRipSplitHorizonCfg871 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpRipAuthentication865 struct {
    Str InterfaceTunnelIpRipAuthenticationStr866 `json:"str"`
    Mode InterfaceTunnelIpRipAuthenticationMode867 `json:"mode"`
    KeyChain InterfaceTunnelIpRipAuthenticationKeyChain868 `json:"key-chain"`
}


type InterfaceTunnelIpRipAuthenticationStr866 struct {
    String string `json:"string"`
}


type InterfaceTunnelIpRipAuthenticationMode867 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceTunnelIpRipAuthenticationKeyChain868 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceTunnelIpRipSendCfg869 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceTunnelIpRipReceiveCfg870 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceTunnelIpRipSplitHorizonCfg871 struct {
    State string `json:"state" dval:"poisoned"`
}

func (p *InterfaceTunnelIp) GetId() string{
    return "1"
}

func (p *InterfaceTunnelIp) getPath() string{
    return "interface/tunnel/" +p.Inst.Ifnum + "/ip"
}

func (p *InterfaceTunnelIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIp::Post")
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

func (p *InterfaceTunnelIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIp::Get")
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
func (p *InterfaceTunnelIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIp::Put")
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

func (p *InterfaceTunnelIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
