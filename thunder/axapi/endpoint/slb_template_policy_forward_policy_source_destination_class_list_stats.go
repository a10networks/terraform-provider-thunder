

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationClassListStats struct {
    
    DestClassList string `json:"dest-class-list"`
    Stats SlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats `json:"stats"`
    Name string 

}
type DataSlbTemplatePolicyForwardPolicySourceDestinationClassListStats struct {
    DtSlbTemplatePolicyForwardPolicySourceDestinationClassListStats SlbTemplatePolicyForwardPolicySourceDestinationClassListStats `json:"class-list"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats struct {
    Hits int `json:"hits"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassListStats) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassListStats) getPath() string{
    
    return "slb/template/policy/" +p.Name + "/forward-policy/source/" +p.Name + "/destination/class-list/"+p.DestClassList+"/stats"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationClassListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplatePolicyForwardPolicySourceDestinationClassListStats,error) {
logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationClassListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplatePolicyForwardPolicySourceDestinationClassListStats
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
