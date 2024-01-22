

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TemplateLimitPolicyLimitPps struct {
	Inst struct {

    DdosProtectionFactor int `json:"ddos-protection-factor"`
    Downlink int `json:"downlink"`
    DownlinkBurstsize int `json:"downlink-burstsize"`
    DownlinkRelaxed int `json:"downlink-relaxed"`
    Total int `json:"total"`
    TotalBurstsize int `json:"total-burstsize"`
    TotalRelaxed int `json:"total-relaxed"`
    Uplink int `json:"uplink"`
    UplinkBurstsize int `json:"uplink-burstsize"`
    UplinkRelaxed int `json:"uplink-relaxed"`
    Uuid string `json:"uuid"`
    PolicyNumber string 

	} `json:"limit-pps"`
}

func (p *TemplateLimitPolicyLimitPps) GetId() string{
    return "1"
}

func (p *TemplateLimitPolicyLimitPps) getPath() string{
    return "template/limit-policy/" +p.Inst.PolicyNumber + "/limit-pps"
}

func (p *TemplateLimitPolicyLimitPps) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitPps::Post")
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

func (p *TemplateLimitPolicyLimitPps) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitPps::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *TemplateLimitPolicyLimitPps) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitPps::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *TemplateLimitPolicyLimitPps) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitPps::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
