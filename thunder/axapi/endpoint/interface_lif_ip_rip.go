

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifIpRip struct {
	Inst struct {

    Authentication InterfaceLifIpRipAuthentication `json:"authentication"`
    ReceiveCfg InterfaceLifIpRipReceiveCfg `json:"receive-cfg"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceLifIpRipSendCfg `json:"send-cfg"`
    SendPacket int `json:"send-packet" dval:"1"`
    SplitHorizonCfg InterfaceLifIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
    Ifname string 

	} `json:"rip"`
}


type InterfaceLifIpRipAuthentication struct {
    Str InterfaceLifIpRipAuthenticationStr `json:"str"`
    Mode InterfaceLifIpRipAuthenticationMode `json:"mode"`
    KeyChain InterfaceLifIpRipAuthenticationKeyChain `json:"key-chain"`
}


type InterfaceLifIpRipAuthenticationStr struct {
    String string `json:"string"`
}


type InterfaceLifIpRipAuthenticationMode struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceLifIpRipAuthenticationKeyChain struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceLifIpRipReceiveCfg struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceLifIpRipSendCfg struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceLifIpRipSplitHorizonCfg struct {
    State string `json:"state" dval:"poisoned"`
}

func (p *InterfaceLifIpRip) GetId() string{
    return "1"
}

func (p *InterfaceLifIpRip) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/ip/rip"
}

func (p *InterfaceLifIpRip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpRip::Post")
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

func (p *InterfaceLifIpRip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpRip::Get")
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
func (p *InterfaceLifIpRip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpRip::Put")
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

func (p *InterfaceLifIpRip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpRip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
