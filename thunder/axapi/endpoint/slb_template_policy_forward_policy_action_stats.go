

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicyActionStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplatePolicyForwardPolicyActionStatsStats `json:"stats"`

}
type DataSlbTemplatePolicyForwardPolicyActionStats struct {
    DtSlbTemplatePolicyForwardPolicyActionStats SlbTemplatePolicyForwardPolicyActionStats `json:"action"`
}


type SlbTemplatePolicyForwardPolicyActionStatsStats struct {
    Hits int `json:"hits"`
}

func (p *SlbTemplatePolicyForwardPolicyActionStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicyActionStats) getPath() string{
    
    return "slb/template/policy/"+p.Name+"/forward-policy/action/"+p.Name+"/stats"
}

func (p *SlbTemplatePolicyForwardPolicyActionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicyActionStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicyActionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicyActionStats
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
