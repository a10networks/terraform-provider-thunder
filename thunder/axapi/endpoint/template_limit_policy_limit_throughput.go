

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TemplateLimitPolicyLimitThroughput struct {
	Inst struct {

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

	} `json:"limit-throughput"`
}

func (p *TemplateLimitPolicyLimitThroughput) GetId() string{
    return "1"
}

func (p *TemplateLimitPolicyLimitThroughput) getPath() string{
    return "template/limit-policy/" +p.Inst.PolicyNumber + "/limit-throughput"
}

func (p *TemplateLimitPolicyLimitThroughput) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitThroughput::Post")
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

func (p *TemplateLimitPolicyLimitThroughput) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitThroughput::Get")
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
func (p *TemplateLimitPolicyLimitThroughput) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitThroughput::Put")
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

func (p *TemplateLimitPolicyLimitThroughput) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicyLimitThroughput::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
