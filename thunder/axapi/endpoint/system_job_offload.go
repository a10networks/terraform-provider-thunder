

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemJobOffload struct {
	Inst struct {

    SamplingEnable []SystemJobOffloadSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"job-offload"`
}


type SystemJobOffloadSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemJobOffload) GetId() string{
    return "1"
}

func (p *SystemJobOffload) getPath() string{
    return "system/job-offload"
}

func (p *SystemJobOffload) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJobOffload::Post")
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

func (p *SystemJobOffload) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJobOffload::Get")
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
func (p *SystemJobOffload) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJobOffload::Put")
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

func (p *SystemJobOffload) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJobOffload::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
