

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type TemplateLimitPolicy struct {
	Inst struct {

    LimitConcurrentSessions int `json:"limit-concurrent-sessions"`
    LimitCps TemplateLimitPolicyLimitCps1903 `json:"limit-cps"`
    LimitPps TemplateLimitPolicyLimitPps1904 `json:"limit-pps"`
    LimitScope string `json:"limit-scope" dval:"subscriber-ip"`
    LimitThroughput TemplateLimitPolicyLimitThroughput1905 `json:"limit-throughput"`
    Log int `json:"log"`
    MaxMinFair int `json:"max-min-fair"`
    Parent int `json:"parent"`
    PolicyNumber int `json:"policy-number"`
    PrefixLength int `json:"prefix-length"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"limit-policy"`
}


type TemplateLimitPolicyLimitCps1903 struct {
    Value int `json:"value"`
    Burstsize int `json:"burstsize"`
    Relaxed int `json:"relaxed"`
    Uuid string `json:"uuid"`
}


type TemplateLimitPolicyLimitPps1904 struct {
    Uplink int `json:"uplink"`
    UplinkBurstsize int `json:"uplink-burstsize"`
    UplinkRelaxed int `json:"uplink-relaxed"`
    Downlink int `json:"downlink"`
    DdosProtectionFactor int `json:"ddos-protection-factor"`
    DownlinkBurstsize int `json:"downlink-burstsize"`
    DownlinkRelaxed int `json:"downlink-relaxed"`
    Total int `json:"total"`
    TotalBurstsize int `json:"total-burstsize"`
    TotalRelaxed int `json:"total-relaxed"`
    Uuid string `json:"uuid"`
}


type TemplateLimitPolicyLimitThroughput1905 struct {
    Uplink int `json:"uplink"`
    UplinkBurstsize int `json:"uplink-burstsize"`
    UplinkRelaxed int `json:"uplink-relaxed"`
    Downlink int `json:"downlink"`
    DownlinkBurstsize int `json:"downlink-burstsize"`
    DownlinkRelaxed int `json:"downlink-relaxed"`
    Total int `json:"total"`
    TotalBurstsize int `json:"total-burstsize"`
    TotalRelaxed int `json:"total-relaxed"`
    Uuid string `json:"uuid"`
}

func (p *TemplateLimitPolicy) GetId() string{
    return strconv.Itoa(p.Inst.PolicyNumber)
}

func (p *TemplateLimitPolicy) getPath() string{
    return "template/limit-policy"
}

func (p *TemplateLimitPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicy::Post")
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

func (p *TemplateLimitPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicy::Get")
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
func (p *TemplateLimitPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicy::Put")
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

func (p *TemplateLimitPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateLimitPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
