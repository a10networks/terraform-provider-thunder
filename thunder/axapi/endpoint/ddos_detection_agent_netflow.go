

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionAgentNetflow struct {
	Inst struct {

    ActiveTimeout int `json:"active-timeout"`
    InactiveTimeout int `json:"inactive-timeout"`
    NetflowSamplesCollection string `json:"netflow-samples-collection" dval:"enable"`
    NetflowSamplingRate int `json:"netflow-sampling-rate" dval:"1"`
    Uuid string `json:"uuid"`
    AgentName string 

	} `json:"netflow"`
}

func (p *DdosDetectionAgentNetflow) GetId() string{
    return "1"
}

func (p *DdosDetectionAgentNetflow) getPath() string{
    return "ddos/detection/agent/" +p.Inst.AgentName + "/netflow"
}

func (p *DdosDetectionAgentNetflow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgentNetflow::Post")
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

func (p *DdosDetectionAgentNetflow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgentNetflow::Get")
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
func (p *DdosDetectionAgentNetflow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgentNetflow::Put")
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

func (p *DdosDetectionAgentNetflow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionAgentNetflow::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
