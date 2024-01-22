

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicyDualStackActionStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplatePolicyForwardPolicyDualStackActionStatsStats `json:"stats"`

}
type DataSlbTemplatePolicyForwardPolicyDualStackActionStats struct {
    DtSlbTemplatePolicyForwardPolicyDualStackActionStats SlbTemplatePolicyForwardPolicyDualStackActionStats `json:"dual-stack-action"`
}


type SlbTemplatePolicyForwardPolicyDualStackActionStatsStats struct {
    Hits int `json:"hits"`
}

func (p *SlbTemplatePolicyForwardPolicyDualStackActionStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicyDualStackActionStats) getPath() string{
    
    return "slb/template/policy/"+p.Name+"/forward-policy/dual-stack-action/"+p.Name+"/stats"
}

func (p *SlbTemplatePolicyForwardPolicyDualStackActionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicyDualStackActionStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicyDualStackActionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicyDualStackActionStats
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
