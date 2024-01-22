

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetRuleMoveRule struct {
	Inst struct {

    Location string `json:"location" dval:"bottom"`
    TargetRule string `json:"target-rule"`
    Name string 

	} `json:"move-rule"`
}

func (p *RuleSetRuleMoveRule) GetId() string{
    return "1"
}

func (p *RuleSetRuleMoveRule) getPath() string{
    return "rule-set/" +p.Inst.Name + "/rule/" +p.Inst.Name + "/move-rule"
}

func (p *RuleSetRuleMoveRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleMoveRule::Post")
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

func (p *RuleSetRuleMoveRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleMoveRule::Get")
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
func (p *RuleSetRuleMoveRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleMoveRule::Put")
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

func (p *RuleSetRuleMoveRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleMoveRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
