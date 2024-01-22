

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbQuic struct {
	Inst struct {

    SamplingEnable []SlbQuicSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"quic"`
}


type SlbQuicSamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
}

func (p *SlbQuic) GetId() string{
    return "1"
}

func (p *SlbQuic) getPath() string{
    return "slb/quic"
}

func (p *SlbQuic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbQuic::Post")
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

func (p *SlbQuic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbQuic::Get")
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
func (p *SlbQuic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbQuic::Put")
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

func (p *SlbQuic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbQuic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
