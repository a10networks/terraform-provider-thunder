

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcBasedPolicyPolicyClassList struct {
	Inst struct {

    ClassListName string `json:"class-list-name"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"policy-class-list"`
}

func (p *DdosSrcBasedPolicyPolicyClassList) GetId() string{
    return url.QueryEscape(p.Inst.ClassListName)
}

func (p *DdosSrcBasedPolicyPolicyClassList) getPath() string{
    return "ddos/src-based-policy/" +p.Inst.Name + "/policy-class-list"
}

func (p *DdosSrcBasedPolicyPolicyClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicyPolicyClassList::Post")
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

func (p *DdosSrcBasedPolicyPolicyClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicyPolicyClassList::Get")
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
func (p *DdosSrcBasedPolicyPolicyClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicyPolicyClassList::Put")
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

func (p *DdosSrcBasedPolicyPolicyClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicyPolicyClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
