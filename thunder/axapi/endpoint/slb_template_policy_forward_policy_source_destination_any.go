

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationAny struct {
	Inst struct {

    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"any"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAny) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAny) getPath() string{
    return "slb/template/policy/" +p.Inst.Name + "/forward-policy/source/" +p.Inst.Name + "/destination/any"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAny) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAny::Post")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAny) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAny::Get")
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
func (p *SlbTemplatePolicyForwardPolicySourceDestinationAny) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAny::Put")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAny) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAny::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
