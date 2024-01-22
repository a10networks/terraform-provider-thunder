

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpCommunityListStandard struct {
	Inst struct {

    RulesList []IpCommunityListStandardRulesList `json:"rules-list"`
    Standard string `json:"standard"`
    Uuid string `json:"uuid"`

	} `json:"standard"`
}


type IpCommunityListStandardRulesList struct {
    StandardAction string `json:"standard-action"`
    StandardCommValue string `json:"standard-comm-value"`
}

func (p *IpCommunityListStandard) GetId() string{
    return p.Inst.Standard
}

func (p *IpCommunityListStandard) getPath() string{
    return "ip/community-list/standard"
}

func (p *IpCommunityListStandard) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListStandard::Post")
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

func (p *IpCommunityListStandard) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListStandard::Get")
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
func (p *IpCommunityListStandard) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListStandard::Put")
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

func (p *IpCommunityListStandard) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpCommunityListStandard::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
