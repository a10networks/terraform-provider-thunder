

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpAddress struct {
	Inst struct {

    IpAddr string `json:"ip-addr"`
    IpMask string `json:"ip-mask"`
    Uuid string `json:"uuid"`

	} `json:"address"`
}

func (p *IpAddress) GetId() string{
    return "1"
}

func (p *IpAddress) getPath() string{
    return "ip/address"
}

func (p *IpAddress) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAddress::Post")
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

func (p *IpAddress) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAddress::Get")
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
func (p *IpAddress) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAddress::Put")
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

func (p *IpAddress) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAddress::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
