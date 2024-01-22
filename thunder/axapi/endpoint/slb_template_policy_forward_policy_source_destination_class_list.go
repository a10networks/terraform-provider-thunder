

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationClassList struct {
	Inst struct {

    Action string `json:"action"`
    DestClassList string `json:"dest-class-list"`
    DualStackAction string `json:"dual-stack-action"`
    Priority int `json:"priority"`
    Type string `json:"type"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"class-list"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassList) GetId() string{
    return url.QueryEscape(p.Inst.DestClassList)
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassList) getPath() string{
    return "slb/template/policy/" +p.Inst.Name + "/forward-policy/source/" +p.Inst.Name + "/destination/class-list"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationClassList::Post")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationClassList::Get")
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
func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationClassList::Put")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
