

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewMemoryView struct {
	Inst struct {

    SamplingEnable []SystemViewMemoryViewSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"memory-view"`
}


type SystemViewMemoryViewSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemViewMemoryView) GetId() string{
    return "1"
}

func (p *SystemViewMemoryView) getPath() string{
    return "system-view/memory-view"
}

func (p *SystemViewMemoryView) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemViewMemoryView::Post")
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

func (p *SystemViewMemoryView) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemViewMemoryView::Get")
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
func (p *SystemViewMemoryView) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemViewMemoryView::Put")
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

func (p *SystemViewMemoryView) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemViewMemoryView::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
