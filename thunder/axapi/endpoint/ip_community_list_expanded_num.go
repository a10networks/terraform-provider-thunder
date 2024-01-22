

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpCommunityListExpandedNum struct {
	Inst struct {

    ExtListNum int `json:"ext-list-num"`
    RulesList []IpCommunityListExpandedNumRulesList `json:"rules-list"`
    Uuid string `json:"uuid"`

	} `json:"expanded-num"`
}


type IpCommunityListExpandedNumRulesList struct {
    ExtListAction string `json:"ext-list-action"`
    ExtListValue string `json:"ext-list-value"`
}

func (p *IpCommunityListExpandedNum) GetId() string{
    return strconv.Itoa(p.Inst.ExtListNum)
}

func (p *IpCommunityListExpandedNum) getPath() string{
    return "ip/community-list/expanded-num"
}

func (p *IpCommunityListExpandedNum) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpandedNum::Post")
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

func (p *IpCommunityListExpandedNum) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpandedNum::Get")
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
func (p *IpCommunityListExpandedNum) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpandedNum::Put")
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

func (p *IpCommunityListExpandedNum) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListExpandedNum::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
