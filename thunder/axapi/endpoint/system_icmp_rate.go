

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIcmpRate struct {
	Inst struct {

    SamplingEnable []SystemIcmpRateSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"icmp-rate"`
}


type SystemIcmpRateSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemIcmpRate) GetId() string{
    return "1"
}

func (p *SystemIcmpRate) getPath() string{
    return "system/icmp-rate"
}

func (p *SystemIcmpRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmpRate::Post")
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

func (p *SystemIcmpRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmpRate::Get")
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
func (p *SystemIcmpRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmpRate::Put")
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

func (p *SystemIcmpRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIcmpRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
