

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkIpRip struct {
	Inst struct {

    Authentication InterfaceTrunkIpRipAuthentication `json:"authentication"`
    ReceiveCfg InterfaceTrunkIpRipReceiveCfg `json:"receive-cfg"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceTrunkIpRipSendCfg `json:"send-cfg"`
    SendPacket int `json:"send-packet" dval:"1"`
    SplitHorizonCfg InterfaceTrunkIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"rip"`
}


type InterfaceTrunkIpRipAuthentication struct {
    Str InterfaceTrunkIpRipAuthenticationStr `json:"str"`
    Mode InterfaceTrunkIpRipAuthenticationMode `json:"mode"`
    KeyChain InterfaceTrunkIpRipAuthenticationKeyChain `json:"key-chain"`
}


type InterfaceTrunkIpRipAuthenticationStr struct {
    String string `json:"string"`
}


type InterfaceTrunkIpRipAuthenticationMode struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceTrunkIpRipAuthenticationKeyChain struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceTrunkIpRipReceiveCfg struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceTrunkIpRipSendCfg struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceTrunkIpRipSplitHorizonCfg struct {
    State string `json:"state" dval:"poisoned"`
}

func (p *InterfaceTrunkIpRip) GetId() string{
    return "1"
}

func (p *InterfaceTrunkIpRip) getPath() string{
    return "interface/trunk/" +p.Inst.Ifnum + "/ip/rip"
}

func (p *InterfaceTrunkIpRip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpRip::Post")
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

func (p *InterfaceTrunkIpRip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpRip::Get")
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
func (p *InterfaceTrunkIpRip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpRip::Put")
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

func (p *InterfaceTrunkIpRip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpRip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
