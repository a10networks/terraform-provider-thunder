

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSystemDefault struct {
	Inst struct {

    LimitList []DdosSystemDefaultLimitList `json:"limit-list"`
    Uuid string `json:"uuid"`

	} `json:"system-default"`
}


type DdosSystemDefaultLimitList struct {
    LimitType string `json:"limit-type"`
    DefaultOverLimitAction DdosSystemDefaultLimitListDefaultOverLimitAction `json:"default-over-limit-action"`
    DefaultPktRateLimit int `json:"default-pkt-rate-limit"`
    DefaultBitRateLimit int `json:"default-bit-rate-limit"`
    DefaultFragPktRateLimit int `json:"default-frag-pkt-rate-limit"`
    DefaultConnLimit int `json:"default-conn-limit"`
    DefaultConnRateLimit int `json:"default-conn-rate-limit"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSystemDefaultLimitListDefaultOverLimitAction struct {
    Drop int `json:"drop"`
}

func (p *DdosSystemDefault) GetId() string{
    return "1"
}

func (p *DdosSystemDefault) getPath() string{
    return "ddos/system-default"
}

func (p *DdosSystemDefault) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefault::Post")
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

func (p *DdosSystemDefault) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefault::Get")
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
func (p *DdosSystemDefault) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefault::Put")
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

func (p *DdosSystemDefault) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSystemDefault::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
