

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AamAaaPolicyAaaRuleStats struct {
    
    Index int `json:"index"`
    Stats AamAaaPolicyAaaRuleStatsStats `json:"stats"`
    Name string 

}
type DataAamAaaPolicyAaaRuleStats struct {
    DtAamAaaPolicyAaaRuleStats AamAaaPolicyAaaRuleStats `json:"aaa-rule"`
}


type AamAaaPolicyAaaRuleStatsStats struct {
    Total_count int `json:"total_count"`
    Hit_deny int `json:"hit_deny"`
    Hit_auth int `json:"hit_auth"`
    Hit_bypass int `json:"hit_bypass"`
    Failure_bypass int `json:"failure_bypass"`
}

func (p *AamAaaPolicyAaaRuleStats) GetId() string{
    return "1"
}

func (p *AamAaaPolicyAaaRuleStats) getPath() string{
    
    return "aam/aaa-policy/" +p.Name + "/aaa-rule/" +strconv.Itoa(p.Index)+"/stats"
}

func (p *AamAaaPolicyAaaRuleStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAaaPolicyAaaRuleStats,error) {
logger.Println("AamAaaPolicyAaaRuleStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAaaPolicyAaaRuleStats
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
