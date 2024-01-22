

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateFix struct {
	Inst struct {

    InsertClientIp int `json:"insert-client-ip"`
    Logging string `json:"logging"`
    Name string `json:"name"`
    TagSwitching []SlbTemplateFixTagSwitching `json:"tag-switching"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"fix"`
}


type SlbTemplateFixTagSwitching struct {
    SwitchingType string `json:"switching-type"`
    Equals string `json:"equals"`
    ServiceGroup string `json:"service-group"`
}

func (p *SlbTemplateFix) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateFix) getPath() string{
    return "slb/template/fix"
}

func (p *SlbTemplateFix) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateFix::Post")
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

func (p *SlbTemplateFix) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateFix::Get")
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
func (p *SlbTemplateFix) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateFix::Put")
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

func (p *SlbTemplateFix) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateFix::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
