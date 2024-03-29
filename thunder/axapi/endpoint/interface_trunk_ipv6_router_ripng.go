

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkIpv6RouterRipng struct {
	Inst struct {

    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ripng"`
}

func (p *InterfaceTrunkIpv6RouterRipng) GetId() string{
    return "1"
}

func (p *InterfaceTrunkIpv6RouterRipng) getPath() string{
    return "interface/trunk/" +p.Inst.Ifnum + "/ipv6/router/ripng"
}

func (p *InterfaceTrunkIpv6RouterRipng) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpv6RouterRipng::Post")
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

func (p *InterfaceTrunkIpv6RouterRipng) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpv6RouterRipng::Get")
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
func (p *InterfaceTrunkIpv6RouterRipng) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpv6RouterRipng::Put")
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

func (p *InterfaceTrunkIpv6RouterRipng) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkIpv6RouterRipng::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
