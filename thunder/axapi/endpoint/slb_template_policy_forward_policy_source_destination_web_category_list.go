

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList struct {
	Inst struct {

    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Priority int `json:"priority"`
    Type string `json:"type"`
    Uuid string `json:"uuid"`
    WebCategoryList string `json:"web-category-list"`
    Name string 

	} `json:"web-category-list"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList) GetId() string{
    return url.QueryEscape(p.Inst.WebCategoryList)
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList) getPath() string{
    return "slb/template/policy/" +p.Inst.Name + "/forward-policy/source/" +p.Inst.Name + "/destination/web-category-list"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList::Post")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList::Get")
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
func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList::Put")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
