

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbIcap struct {
	Inst struct {

    SamplingEnable []SlbIcapSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"icap"`
}


type SlbIcapSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbIcap) GetId() string{
    return "1"
}

func (p *SlbIcap) getPath() string{
    return "slb/icap"
}

func (p *SlbIcap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIcap::Post")
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

func (p *SlbIcap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIcap::Get")
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
func (p *SlbIcap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIcap::Put")
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

func (p *SlbIcap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbIcap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
