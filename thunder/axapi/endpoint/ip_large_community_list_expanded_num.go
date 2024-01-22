

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpLargeCommunityListExpandedNum struct {
	Inst struct {

    ExtListNum int `json:"ext-list-num"`
    RulesList []IpLargeCommunityListExpandedNumRulesList `json:"rules-list"`
    Uuid string `json:"uuid"`

	} `json:"expanded-num"`
}


type IpLargeCommunityListExpandedNumRulesList struct {
    ExtListLcomAction string `json:"ext-list-lcom-action"`
    ExtListLcomValue string `json:"ext-list-lcom-value"`
}

func (p *IpLargeCommunityListExpandedNum) GetId() string{
    return strconv.Itoa(p.Inst.ExtListNum)
}

func (p *IpLargeCommunityListExpandedNum) getPath() string{
    return "ip/large-community-list/expanded-num"
}

func (p *IpLargeCommunityListExpandedNum) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListExpandedNum::Post")
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

func (p *IpLargeCommunityListExpandedNum) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListExpandedNum::Get")
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
func (p *IpLargeCommunityListExpandedNum) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListExpandedNum::Put")
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

func (p *IpLargeCommunityListExpandedNum) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListExpandedNum::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
