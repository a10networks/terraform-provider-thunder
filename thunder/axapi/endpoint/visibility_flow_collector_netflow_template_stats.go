

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFlowCollectorNetflowTemplateStats struct {
    
    Stats VisibilityFlowCollectorNetflowTemplateStatsStats `json:"stats"`

}
type DataVisibilityFlowCollectorNetflowTemplateStats struct {
    DtVisibilityFlowCollectorNetflowTemplateStats VisibilityFlowCollectorNetflowTemplateStats `json:"template"`
}


type VisibilityFlowCollectorNetflowTemplateStatsStats struct {
    TemplatesAddedToDelq int `json:"templates-added-to-delq"`
    TemplatesRemovedFromDelq int `json:"templates-removed-from-delq"`
}

func (p *VisibilityFlowCollectorNetflowTemplateStats) GetId() string{
    return "1"
}

func (p *VisibilityFlowCollectorNetflowTemplateStats) getPath() string{
    return "visibility/flow-collector/netflow/template/stats"
}

func (p *VisibilityFlowCollectorNetflowTemplateStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityFlowCollectorNetflowTemplateStats,error) {
logger.Println("VisibilityFlowCollectorNetflowTemplateStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityFlowCollectorNetflowTemplateStats
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
