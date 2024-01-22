

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcBasedPolicy struct {
	Inst struct {

    Name string `json:"name"`
    PolicyClassListList []DdosSrcBasedPolicyPolicyClassListList `json:"policy-class-list-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"src-based-policy"`
}


type DdosSrcBasedPolicyPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Uuid string `json:"uuid"`
}

func (p *DdosSrcBasedPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosSrcBasedPolicy) getPath() string{
    return "ddos/src-based-policy"
}

func (p *DdosSrcBasedPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicy::Post")
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

func (p *DdosSrcBasedPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicy::Get")
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
func (p *DdosSrcBasedPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicy::Put")
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

func (p *DdosSrcBasedPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcBasedPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
