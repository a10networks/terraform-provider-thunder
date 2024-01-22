

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosLocalAddressIp struct {
	Inst struct {

    IpAddr string `json:"ip-addr"`
    Uuid string `json:"uuid"`

	} `json:"ip"`
}

func (p *DdosLocalAddressIp) GetId() string{
    return p.Inst.IpAddr
}

func (p *DdosLocalAddressIp) getPath() string{
    return "ddos/local-address/ip"
}

func (p *DdosLocalAddressIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosLocalAddressIp::Post")
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

func (p *DdosLocalAddressIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosLocalAddressIp::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *DdosLocalAddressIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosLocalAddressIp::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *DdosLocalAddressIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosLocalAddressIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
