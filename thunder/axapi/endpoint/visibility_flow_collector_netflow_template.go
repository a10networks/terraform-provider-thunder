

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFlowCollectorNetflowTemplate struct {
	Inst struct {

    Detail VisibilityFlowCollectorNetflowTemplateDetail1910 `json:"detail"`
    SamplingEnable []VisibilityFlowCollectorNetflowTemplateSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"template"`
}


type VisibilityFlowCollectorNetflowTemplateDetail1910 struct {
    Uuid string `json:"uuid"`
}


type VisibilityFlowCollectorNetflowTemplateSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *VisibilityFlowCollectorNetflowTemplate) GetId() string{
    return "1"
}

func (p *VisibilityFlowCollectorNetflowTemplate) getPath() string{
    return "visibility/flow-collector/netflow/template"
}

func (p *VisibilityFlowCollectorNetflowTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflowTemplate::Post")
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

func (p *VisibilityFlowCollectorNetflowTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflowTemplate::Get")
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
func (p *VisibilityFlowCollectorNetflowTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflowTemplate::Put")
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

func (p *VisibilityFlowCollectorNetflowTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflowTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
