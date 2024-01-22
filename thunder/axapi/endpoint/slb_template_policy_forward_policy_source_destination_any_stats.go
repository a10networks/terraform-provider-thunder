

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationAnyStats struct {
    
    Stats SlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats `json:"stats"`
    Name string 

}
type DataSlbTemplatePolicyForwardPolicySourceDestinationAnyStats struct {
    DtSlbTemplatePolicyForwardPolicySourceDestinationAnyStats SlbTemplatePolicyForwardPolicySourceDestinationAnyStats `json:"any"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats struct {
    Hits int `json:"hits"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAnyStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAnyStats) getPath() string{
    
    return "slb/template/policy/" +p.Name + "/forward-policy/source/" +p.Name + "/destination/any/stats"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAnyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicySourceDestinationAnyStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAnyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicySourceDestinationAnyStats
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
