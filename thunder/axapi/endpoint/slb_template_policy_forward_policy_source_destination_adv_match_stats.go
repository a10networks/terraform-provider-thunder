

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats struct {
    
    Priority int `json:"priority"`
    Stats SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats `json:"stats"`
    Name string 

}
type DataSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats struct {
    DtSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats `json:"adv-match"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats struct {
    Hits int `json:"hits"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats) getPath() string{
    
    return "slb/template/policy/" +p.Name + "/forward-policy/source/" +p.Name + "/destination/adv-match/" +strconv.Itoa(p.Priority)+"/stats"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats
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
