

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetRuleStats struct {
    
    Name string `json:"name"`
    Stats RuleSetRuleStatsStats `json:"stats"`

}
type DataRuleSetRuleStats struct {
    DtRuleSetRuleStats RuleSetRuleStats `json:"rule"`
}


type RuleSetRuleStatsStats struct {
    HitCount int `json:"hit-count"`
    PermitBytes int `json:"permit-bytes"`
    DenyBytes int `json:"deny-bytes"`
    ResetBytes int `json:"reset-bytes"`
    PermitPackets int `json:"permit-packets"`
    DenyPackets int `json:"deny-packets"`
    ResetPackets int `json:"reset-packets"`
    ActiveSessionTcp int `json:"active-session-tcp"`
    ActiveSessionUdp int `json:"active-session-udp"`
    ActiveSessionIcmp int `json:"active-session-icmp"`
    ActiveSessionOther int `json:"active-session-other"`
    SessionTcp int `json:"session-tcp"`
    SessionUdp int `json:"session-udp"`
    SessionIcmp int `json:"session-icmp"`
    SessionOther int `json:"session-other"`
    ActiveSessionSctp int `json:"active-session-sctp"`
    SessionSctp int `json:"session-sctp"`
    HitcountTimestamp int `json:"hitcount-timestamp"`
    RateLimitDrops int `json:"rate-limit-drops"`
}

func (p *RuleSetRuleStats) GetId() string{
    return "1"
}

func (p *RuleSetRuleStats) getPath() string{
    
    return "rule-set/"+p.Name+"/rule/"+p.Name+"/stats"
}

func (p *RuleSetRuleStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetRuleStats,error) {
logger.Println("RuleSetRuleStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetRuleStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
