

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopbackIp struct {
	Inst struct {

    AddressList []InterfaceLoopbackIpAddressList `json:"address-list"`
    Ospf InterfaceLoopbackIpOspf652 `json:"ospf"`
    Rip InterfaceLoopbackIpRip659 `json:"rip"`
    Router InterfaceLoopbackIpRouter667 `json:"router"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ip"`
}


type InterfaceLoopbackIpAddressList struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
}


type InterfaceLoopbackIpOspf652 struct {
    OspfGlobal InterfaceLoopbackIpOspfOspfGlobal653 `json:"ospf-global"`
    OspfIpList []InterfaceLoopbackIpOspfOspfIpList `json:"ospf-ip-list"`
}


type InterfaceLoopbackIpOspfOspfGlobal653 struct {
    AuthenticationCfg InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg654 `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceLoopbackIpOspfOspfGlobalBfdCfg655 `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg656 `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg657 `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg654 struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceLoopbackIpOspfOspfGlobalBfdCfg655 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg656 struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg657 struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5658 `json:"md5"`
}


type InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5658 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLoopbackIpOspfOspfIpList struct {
    IpAddr string `json:"ip-addr"`
    Authentication int `json:"authentication"`
    Value string `json:"value"`
    AuthenticationKey string `json:"authentication-key"`
    Cost int `json:"cost"`
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg `json:"message-digest-cfg"`
    MtuIgnore int `json:"mtu-ignore"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLoopbackIpRip659 struct {
    Authentication InterfaceLoopbackIpRipAuthentication660 `json:"authentication"`
    SendPacket int `json:"send-packet" dval:"1"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceLoopbackIpRipSendCfg664 `json:"send-cfg"`
    ReceiveCfg InterfaceLoopbackIpRipReceiveCfg665 `json:"receive-cfg"`
    SplitHorizonCfg InterfaceLoopbackIpRipSplitHorizonCfg666 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpRipAuthentication660 struct {
    Str InterfaceLoopbackIpRipAuthenticationStr661 `json:"str"`
    Mode InterfaceLoopbackIpRipAuthenticationMode662 `json:"mode"`
    KeyChain InterfaceLoopbackIpRipAuthenticationKeyChain663 `json:"key-chain"`
}


type InterfaceLoopbackIpRipAuthenticationStr661 struct {
    String string `json:"string"`
}


type InterfaceLoopbackIpRipAuthenticationMode662 struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceLoopbackIpRipAuthenticationKeyChain663 struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceLoopbackIpRipSendCfg664 struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceLoopbackIpRipReceiveCfg665 struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceLoopbackIpRipSplitHorizonCfg666 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceLoopbackIpRouter667 struct {
    Isis InterfaceLoopbackIpRouterIsis668 `json:"isis"`
}


type InterfaceLoopbackIpRouterIsis668 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceLoopbackIp) GetId() string{
    return "1"
}

func (p *InterfaceLoopbackIp) getPath() string{
    return "interface/loopback/" +p.Inst.Ifnum + "/ip"
}

func (p *InterfaceLoopbackIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIp::Post")
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

func (p *InterfaceLoopbackIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIp::Get")
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
func (p *InterfaceLoopbackIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIp::Put")
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

func (p *InterfaceLoopbackIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
