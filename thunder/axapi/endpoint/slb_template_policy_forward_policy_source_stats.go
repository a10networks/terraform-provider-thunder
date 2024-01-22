

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplatePolicyForwardPolicySourceStatsStats `json:"stats"`

}
type DataSlbTemplatePolicyForwardPolicySourceStats struct {
    DtSlbTemplatePolicyForwardPolicySourceStats SlbTemplatePolicyForwardPolicySourceStats `json:"source"`
}


type SlbTemplatePolicyForwardPolicySourceStatsStats struct {
    Hits int `json:"hits"`
    DestinationMatchNotFound int `json:"destination-match-not-found"`
    NoHostInfo int `json:"no-host-info"`
}

func (p *SlbTemplatePolicyForwardPolicySourceStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceStats) getPath() string{
    
    return "slb/template/policy/"+p.Name+"/forward-policy/source/"+p.Name+"/stats"
}

func (p *SlbTemplatePolicyForwardPolicySourceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicySourceStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicySourceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicySourceStats
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
