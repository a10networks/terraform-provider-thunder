

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuHyperThread struct {
	Inst struct {

    Disable int `json:"disable"`
    Enable int `json:"enable"`

	} `json:"cpu-hyper-thread"`
}

func (p *SystemCpuHyperThread) GetId() string{
    return "1"
}

func (p *SystemCpuHyperThread) getPath() string{
    return "system/cpu-hyper-thread"
}

func (p *SystemCpuHyperThread) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuHyperThread::Post")
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

func (p *SystemCpuHyperThread) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuHyperThread::Get")
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
func (p *SystemCpuHyperThread) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuHyperThread::Put")
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

func (p *SystemCpuHyperThread) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemCpuHyperThread::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
