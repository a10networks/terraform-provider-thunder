

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpDnsSuffix struct {
	Inst struct {

    DomainName string `json:"domain-name"`
    Uuid string `json:"uuid"`

	} `json:"suffix"`
}

func (p *IpDnsSuffix) GetId() string{
    return "1"
}

func (p *IpDnsSuffix) getPath() string{
    return "ip/dns/suffix"
}

func (p *IpDnsSuffix) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpDnsSuffix::Post")
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

func (p *IpDnsSuffix) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpDnsSuffix::Get")
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
func (p *IpDnsSuffix) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpDnsSuffix::Put")
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

func (p *IpDnsSuffix) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpDnsSuffix::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
