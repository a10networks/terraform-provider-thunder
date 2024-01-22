

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSystemDefaultLimit struct {
	Inst struct {

    DefaultBitRateLimit int `json:"default-bit-rate-limit"`
    DefaultConnLimit int `json:"default-conn-limit"`
    DefaultConnRateLimit int `json:"default-conn-rate-limit"`
    DefaultFragPktRateLimit int `json:"default-frag-pkt-rate-limit"`
    DefaultOverLimitAction DdosSystemDefaultLimitDefaultOverLimitAction `json:"default-over-limit-action"`
    DefaultPktRateLimit int `json:"default-pkt-rate-limit"`
    LimitType string `json:"limit-type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"limit"`
}


type DdosSystemDefaultLimitDefaultOverLimitAction struct {
    Drop int `json:"drop"`
}

func (p *DdosSystemDefaultLimit) GetId() string{
    return p.Inst.LimitType
}

func (p *DdosSystemDefaultLimit) getPath() string{
    return "ddos/system-default/limit"
}

func (p *DdosSystemDefaultLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefaultLimit::Post")
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

func (p *DdosSystemDefaultLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefaultLimit::Get")
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
func (p *DdosSystemDefaultLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefaultLimit::Put")
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

func (p *DdosSystemDefaultLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefaultLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
