

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SflowCollectorCustom struct {
	Inst struct {

    Ipv4_addr string `json:"ipv4_addr"`
    Ipv6_addr string `json:"ipv6_addr"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"custom"`
}

func (p *SflowCollectorCustom) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SflowCollectorCustom) getPath() string{
    return "sflow/collector/custom"
}

func (p *SflowCollectorCustom) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorCustom::Post")
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

func (p *SflowCollectorCustom) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorCustom::Get")
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
func (p *SflowCollectorCustom) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorCustom::Put")
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

func (p *SflowCollectorCustom) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowCollectorCustom::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
