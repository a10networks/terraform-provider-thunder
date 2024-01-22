

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpExtcommunityListExpandedNum struct {
	Inst struct {

    ExtListNum int `json:"ext-list-num"`
    RulesList []IpExtcommunityListExpandedNumRulesList `json:"rules-list"`
    Uuid string `json:"uuid"`

	} `json:"expanded-num"`
}


type IpExtcommunityListExpandedNumRulesList struct {
    ExtListAction string `json:"ext-list-action"`
    ExtListValue string `json:"ext-list-value"`
}

func (p *IpExtcommunityListExpandedNum) GetId() string{
    return strconv.Itoa(p.Inst.ExtListNum)
}

func (p *IpExtcommunityListExpandedNum) getPath() string{
    return "ip/extcommunity-list/expanded-num"
}

func (p *IpExtcommunityListExpandedNum) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpandedNum::Post")
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

func (p *IpExtcommunityListExpandedNum) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpandedNum::Get")
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
func (p *IpExtcommunityListExpandedNum) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpandedNum::Put")
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

func (p *IpExtcommunityListExpandedNum) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListExpandedNum::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
