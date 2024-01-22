

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpLargeCommunityListStandard struct {
	Inst struct {

    RulesList []IpLargeCommunityListStandardRulesList `json:"rules-list"`
    Standard string `json:"standard"`
    Uuid string `json:"uuid"`

	} `json:"standard"`
}


type IpLargeCommunityListStandardRulesList struct {
    StandardLcomAction string `json:"standard-lcom-action"`
    StandardLcommValue string `json:"standard-lcomm-value"`
}

func (p *IpLargeCommunityListStandard) GetId() string{
    return p.Inst.Standard
}

func (p *IpLargeCommunityListStandard) getPath() string{
    return "ip/large-community-list/standard"
}

func (p *IpLargeCommunityListStandard) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandard::Post")
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

func (p *IpLargeCommunityListStandard) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandard::Get")
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
func (p *IpLargeCommunityListStandard) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandard::Put")
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

func (p *IpLargeCommunityListStandard) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpLargeCommunityListStandard::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
