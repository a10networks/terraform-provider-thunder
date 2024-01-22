

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpLargeCommunityListStandardNum struct {
	Inst struct {

    RulesList []IpLargeCommunityListStandardNumRulesList `json:"rules-list"`
    StdListNum int `json:"std-list-num"`
    Uuid string `json:"uuid"`

	} `json:"standard-num"`
}


type IpLargeCommunityListStandardNumRulesList struct {
    StdListLcomAction string `json:"std-list-lcom-action"`
    StdListLcommValue string `json:"std-list-lcomm-value"`
}

func (p *IpLargeCommunityListStandardNum) GetId() string{
    return strconv.Itoa(p.Inst.StdListNum)
}

func (p *IpLargeCommunityListStandardNum) getPath() string{
    return "ip/large-community-list/standard-num"
}

func (p *IpLargeCommunityListStandardNum) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandardNum::Post")
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

func (p *IpLargeCommunityListStandardNum) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandardNum::Get")
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
func (p *IpLargeCommunityListStandardNum) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandardNum::Put")
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

func (p *IpLargeCommunityListStandardNum) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandardNum::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
