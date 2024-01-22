

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RbaRole struct {
	Inst struct {

    DefaultPrivilege string `json:"default-privilege" dval:"no-access"`
    Name string `json:"name"`
    PartitionOnly int `json:"partition-only"`
    RuleList []RbaRoleRuleList `json:"rule-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"role"`
}


type RbaRoleRuleList struct {
    Object string `json:"object"`
    Operation string `json:"operation"`
}

func (p *RbaRole) GetId() string{
    return p.Inst.Name
}

func (p *RbaRole) getPath() string{
    return "rba/role"
}

func (p *RbaRole) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaRole::Post")
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

func (p *RbaRole) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaRole::Get")
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
func (p *RbaRole) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RbaRole::Put")
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

func (p *RbaRole) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RbaRole::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
