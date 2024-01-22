

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope struct {
	Inst struct {

    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Priority int `json:"priority"`
    Type string `json:"type"`
    Uuid string `json:"uuid"`
    WebReputationScope string `json:"web-reputation-scope"`
    Name string 

	} `json:"web-reputation-scope"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope) GetId() string{
    return url.QueryEscape(p.Inst.WebReputationScope)
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope) getPath() string{
    return "slb/template/policy/" +p.Inst.Name + "/forward-policy/source/" +p.Inst.Name + "/destination/web-reputation-scope"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope::Post")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
