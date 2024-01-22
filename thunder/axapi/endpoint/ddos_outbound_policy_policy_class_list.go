

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosOutboundPolicyPolicyClassList struct {
	Inst struct {

    ClassListGlid string `json:"class-list-glid"`
    ClassListName string `json:"class-list-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"policy-class-list"`
}

func (p *DdosOutboundPolicyPolicyClassList) GetId() string{
    return url.QueryEscape(p.Inst.ClassListName)
}

func (p *DdosOutboundPolicyPolicyClassList) getPath() string{
    return "ddos/outbound-policy/" +p.Inst.Name + "/policy-class-list"
}

func (p *DdosOutboundPolicyPolicyClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyClassList::Post")
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

func (p *DdosOutboundPolicyPolicyClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyClassList::Get")
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
func (p *DdosOutboundPolicyPolicyClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyClassList::Put")
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

func (p *DdosOutboundPolicyPolicyClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
