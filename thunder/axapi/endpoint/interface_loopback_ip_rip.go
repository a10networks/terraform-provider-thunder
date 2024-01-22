

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopbackIpRip struct {
	Inst struct {

    Authentication InterfaceLoopbackIpRipAuthentication `json:"authentication"`
    ReceiveCfg InterfaceLoopbackIpRipReceiveCfg `json:"receive-cfg"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceLoopbackIpRipSendCfg `json:"send-cfg"`
    SendPacket int `json:"send-packet" dval:"1"`
    SplitHorizonCfg InterfaceLoopbackIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"rip"`
}


type InterfaceLoopbackIpRipAuthentication struct {
    Str InterfaceLoopbackIpRipAuthenticationStr `json:"str"`
    Mode InterfaceLoopbackIpRipAuthenticationMode `json:"mode"`
    KeyChain InterfaceLoopbackIpRipAuthenticationKeyChain `json:"key-chain"`
}


type InterfaceLoopbackIpRipAuthenticationStr struct {
    String string `json:"string"`
}


type InterfaceLoopbackIpRipAuthenticationMode struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceLoopbackIpRipAuthenticationKeyChain struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceLoopbackIpRipReceiveCfg struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceLoopbackIpRipSendCfg struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceLoopbackIpRipSplitHorizonCfg struct {
    State string `json:"state" dval:"poisoned"`
}

func (p *InterfaceLoopbackIpRip) GetId() string{
    return "1"
}

func (p *InterfaceLoopbackIpRip) getPath() string{
    return "interface/loopback/" +p.Inst.Ifnum + "/ip/rip"
}

func (p *InterfaceLoopbackIpRip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRip::Post")
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

func (p *InterfaceLoopbackIpRip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRip::Get")
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
func (p *InterfaceLoopbackIpRip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRip::Put")
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

func (p *InterfaceLoopbackIpRip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
