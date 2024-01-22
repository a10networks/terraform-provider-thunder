

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpExtcommunityListStandardNum struct {
	Inst struct {

    RulesList []IpExtcommunityListStandardNumRulesList `json:"rules-list"`
    StdListNum int `json:"std-list-num"`
    Uuid string `json:"uuid"`

	} `json:"standard-num"`
}


type IpExtcommunityListStandardNumRulesList struct {
    StdListAction string `json:"std-list-action"`
    StdListValue string `json:"std-list-value"`
}

func (p *IpExtcommunityListStandardNum) GetId() string{
    return strconv.Itoa(p.Inst.StdListNum)
}

func (p *IpExtcommunityListStandardNum) getPath() string{
    return "ip/extcommunity-list/standard-num"
}

func (p *IpExtcommunityListStandardNum) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandardNum::Post")
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

func (p *IpExtcommunityListStandardNum) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandardNum::Get")
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
func (p *IpExtcommunityListStandardNum) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandardNum::Put")
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

func (p *IpExtcommunityListStandardNum) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandardNum::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
