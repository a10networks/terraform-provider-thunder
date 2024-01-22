

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFlowCollectorNetflow struct {
	Inst struct {

    SamplingEnable []VisibilityFlowCollectorNetflowSamplingEnable `json:"sampling-enable"`
    Template VisibilityFlowCollectorNetflowTemplate1911 `json:"template"`
    Uuid string `json:"uuid"`

	} `json:"netflow"`
}


type VisibilityFlowCollectorNetflowSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VisibilityFlowCollectorNetflowTemplate1911 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityFlowCollectorNetflowTemplateSamplingEnable1912 `json:"sampling-enable"`
    Detail VisibilityFlowCollectorNetflowTemplateDetail1913 `json:"detail"`
}


type VisibilityFlowCollectorNetflowTemplateSamplingEnable1912 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityFlowCollectorNetflowTemplateDetail1913 struct {
    Uuid string `json:"uuid"`
}

func (p *VisibilityFlowCollectorNetflow) GetId() string{
    return "1"
}

func (p *VisibilityFlowCollectorNetflow) getPath() string{
    return "visibility/flow-collector/netflow"
}

func (p *VisibilityFlowCollectorNetflow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflow::Post")
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

func (p *VisibilityFlowCollectorNetflow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflow::Get")
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
func (p *VisibilityFlowCollectorNetflow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflow::Put")
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

func (p *VisibilityFlowCollectorNetflow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityFlowCollectorNetflow::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
