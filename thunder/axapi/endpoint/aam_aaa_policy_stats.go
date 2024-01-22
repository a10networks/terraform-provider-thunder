

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAaaPolicyStats struct {
    
    AaaRuleList []AamAaaPolicyStatsAaaRuleList `json:"aaa-rule-list"`
    Name string `json:"name"`
    Stats AamAaaPolicyStatsStats `json:"stats"`

}
type DataAamAaaPolicyStats struct {
    DtAamAaaPolicyStats AamAaaPolicyStats `json:"aaa-policy"`
}


type AamAaaPolicyStatsAaaRuleList struct {
    Index int `json:"index"`
    Stats AamAaaPolicyStatsAaaRuleListStats `json:"stats"`
}


type AamAaaPolicyStatsAaaRuleListStats struct {
    Total_count int `json:"total_count"`
    Hit_deny int `json:"hit_deny"`
    Hit_auth int `json:"hit_auth"`
    Hit_bypass int `json:"hit_bypass"`
    Failure_bypass int `json:"failure_bypass"`
}


type AamAaaPolicyStatsStats struct {
    Req int `json:"req"`
    ReqReject int `json:"req-reject"`
    ReqAuth int `json:"req-auth"`
    ReqBypass int `json:"req-bypass"`
    ReqSkip int `json:"req-skip"`
    Error int `json:"error"`
    FailureBypass int `json:"failure-bypass"`
}

func (p *AamAaaPolicyStats) GetId() string{
    return "1"
}

func (p *AamAaaPolicyStats) getPath() string{
    
    return "aam/aaa-policy/"+p.Name+"/stats"
}

func (p *AamAaaPolicyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAaaPolicyStats,error) {
logger.Println("AamAaaPolicyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAaaPolicyStats
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
