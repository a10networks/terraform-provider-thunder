

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTunnelIpRip struct {
	Inst struct {

    Authentication InterfaceTunnelIpRipAuthentication `json:"authentication"`
    ReceiveCfg InterfaceTunnelIpRipReceiveCfg `json:"receive-cfg"`
    ReceivePacket int `json:"receive-packet" dval:"1"`
    SendCfg InterfaceTunnelIpRipSendCfg `json:"send-cfg"`
    SendPacket int `json:"send-packet" dval:"1"`
    SplitHorizonCfg InterfaceTunnelIpRipSplitHorizonCfg `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"rip"`
}


type InterfaceTunnelIpRipAuthentication struct {
    Str InterfaceTunnelIpRipAuthenticationStr `json:"str"`
    Mode InterfaceTunnelIpRipAuthenticationMode `json:"mode"`
    KeyChain InterfaceTunnelIpRipAuthenticationKeyChain `json:"key-chain"`
}


type InterfaceTunnelIpRipAuthenticationStr struct {
    String string `json:"string"`
}


type InterfaceTunnelIpRipAuthenticationMode struct {
    Mode string `json:"mode" dval:"text"`
}


type InterfaceTunnelIpRipAuthenticationKeyChain struct {
    KeyChain string `json:"key-chain"`
}


type InterfaceTunnelIpRipReceiveCfg struct {
    Receive int `json:"receive"`
    Version string `json:"version"`
}


type InterfaceTunnelIpRipSendCfg struct {
    Send int `json:"send"`
    Version string `json:"version"`
}


type InterfaceTunnelIpRipSplitHorizonCfg struct {
    State string `json:"state" dval:"poisoned"`
}

func (p *InterfaceTunnelIpRip) GetId() string{
    return "1"
}

func (p *InterfaceTunnelIpRip) getPath() string{
    return "interface/tunnel/" +p.Inst.Ifnum + "/ip/rip"
}

func (p *InterfaceTunnelIpRip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpRip::Post")
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

func (p *InterfaceTunnelIpRip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpRip::Get")
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
func (p *InterfaceTunnelIpRip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpRip::Put")
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

func (p *InterfaceTunnelIpRip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpRip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
