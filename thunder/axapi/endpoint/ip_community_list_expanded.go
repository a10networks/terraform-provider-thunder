

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpCommunityListExpanded struct {
	Inst struct {

    Expanded string `json:"expanded"`
    RulesList []IpCommunityListExpandedRulesList `json:"rules-list"`
    Uuid string `json:"uuid"`

	} `json:"expanded"`
}


type IpCommunityListExpandedRulesList struct {
    ExpandedAction string `json:"expanded-action"`
    ExpandedValue string `json:"expanded-value"`
}

func (p *IpCommunityListExpanded) GetId() string{
    return p.Inst.Expanded
}

func (p *IpCommunityListExpanded) getPath() string{
    return "ip/community-list/expanded"
}

func (p *IpCommunityListExpanded) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpanded::Post")
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

func (p *IpCommunityListExpanded) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpanded::Get")
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
func (p *IpCommunityListExpanded) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpanded::Put")
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

func (p *IpCommunityListExpanded) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpanded::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
