

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpExtcommunityListStandard struct {
	Inst struct {

    RulesList []IpExtcommunityListStandardRulesList `json:"rules-list"`
    Standard string `json:"standard"`
    Uuid string `json:"uuid"`

	} `json:"standard"`
}


type IpExtcommunityListStandardRulesList struct {
    StandardAction string `json:"standard-action"`
    StandardValue string `json:"standard-value"`
}

func (p *IpExtcommunityListStandard) GetId() string{
    return p.Inst.Standard
}

func (p *IpExtcommunityListStandard) getPath() string{
    return "ip/extcommunity-list/standard"
}

func (p *IpExtcommunityListStandard) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandard::Post")
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

func (p *IpExtcommunityListStandard) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandard::Get")
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
func (p *IpExtcommunityListStandard) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandard::Put")
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

func (p *IpExtcommunityListStandard) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpExtcommunityListStandard::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
