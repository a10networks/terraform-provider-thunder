

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpExtcommunityListExpanded struct {
	Inst struct {

    Expanded string `json:"expanded"`
    RulesList []IpExtcommunityListExpandedRulesList `json:"rules-list"`
    Uuid string `json:"uuid"`

	} `json:"expanded"`
}


type IpExtcommunityListExpandedRulesList struct {
    ExpandedAction string `json:"expanded-action"`
    ExpandedValue string `json:"expanded-value"`
}

func (p *IpExtcommunityListExpanded) GetId() string{
    return p.Inst.Expanded
}

func (p *IpExtcommunityListExpanded) getPath() string{
    return "ip/extcommunity-list/expanded"
}

func (p *IpExtcommunityListExpanded) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpanded::Post")
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

func (p *IpExtcommunityListExpanded) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpanded::Get")
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
func (p *IpExtcommunityListExpanded) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpanded::Put")
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

func (p *IpExtcommunityListExpanded) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpanded::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
