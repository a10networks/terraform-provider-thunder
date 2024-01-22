

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats struct {
    
    Stats SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats `json:"stats"`
    WebReputationScope string `json:"web-reputation-scope"`
    Name string 

}
type DataSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats struct {
    DtSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats `json:"web-reputation-scope"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats struct {
    Hits int `json:"hits"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats) getPath() string{
    
    return "slb/template/policy/" +p.Name + "/forward-policy/source/" +p.Name + "/destination/web-reputation-scope/"+p.WebReputationScope+"/stats"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats
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
