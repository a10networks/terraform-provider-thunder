

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeIpRip struct {
	Inst struct {

    Authentication InterfaceVeIpRipAuthentication `json:"authentication"`
    ReceiveCfg InterfaceVeIpRipReceiveCfg `json:"receive-cfg"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceVeIpRipSendCfg `json:"send-cfg"`
    SendPacket int `json:"send-packet" dval:"1"`
    SplitHorizonCfg InterfaceVeIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"rip"`
}


type InterfaceVeIpRipAuthentication struct {
    Str InterfaceVeIpRipAuthenticationStr `json:"str"`
    Mode InterfaceVeIpRipAuthenticationMode `json:"mode"`
    KeyChain InterfaceVeIpRipAuthenticationKeyChain `json:"key-chain"`
}


type InterfaceVeIpRipAuthenticationStr struct {
    String string `json:"string"`
}


type InterfaceVeIpRipAuthenticationMode struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceVeIpRipAuthenticationKeyChain struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceVeIpRipReceiveCfg struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceVeIpRipSendCfg struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceVeIpRipSplitHorizonCfg struct {
    State string `json:"state" dval:"poisoned"`
}

func (p *InterfaceVeIpRip) GetId() string{
    return "1"
}

func (p *InterfaceVeIpRip) getPath() string{
    return "interface/ve/" +p.Inst.Ifnum + "/ip/rip"
}

func (p *InterfaceVeIpRip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpRip::Post")
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

func (p *InterfaceVeIpRip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpRip::Get")
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
func (p *InterfaceVeIpRip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpRip::Put")
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

func (p *InterfaceVeIpRip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpRip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
