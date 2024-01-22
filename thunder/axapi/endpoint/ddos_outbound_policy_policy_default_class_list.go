

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosOutboundPolicyPolicyDefaultClassList struct {
	Inst struct {

    ClassListGlid string `json:"class-list-glid"`
    Configuration int `json:"configuration"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"policy-default-class-list"`
}

func (p *DdosOutboundPolicyPolicyDefaultClassList) GetId() string{
    return "1"
}

func (p *DdosOutboundPolicyPolicyDefaultClassList) getPath() string{
    return "ddos/outbound-policy/" +p.Inst.Name + "/policy-default-class-list"
}

func (p *DdosOutboundPolicyPolicyDefaultClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyDefaultClassList::Post")
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

func (p *DdosOutboundPolicyPolicyDefaultClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyDefaultClassList::Get")
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
func (p *DdosOutboundPolicyPolicyDefaultClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyDefaultClassList::Put")
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

func (p *DdosOutboundPolicyPolicyDefaultClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyPolicyDefaultClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
