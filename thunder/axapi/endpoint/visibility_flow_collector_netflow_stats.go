

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFlowCollectorNetflowStats struct {
    
    Stats VisibilityFlowCollectorNetflowStatsStats `json:"stats"`
    Template VisibilityFlowCollectorNetflowStatsTemplate `json:"template"`

}
type DataVisibilityFlowCollectorNetflowStats struct {
    DtVisibilityFlowCollectorNetflowStats VisibilityFlowCollectorNetflowStats `json:"netflow"`
}


type VisibilityFlowCollectorNetflowStatsStats struct {
    PktsRcvd int `json:"pkts-rcvd"`
    V9TemplatesCreated int `json:"v9-templates-created"`
    V9TemplatesDeleted int `json:"v9-templates-deleted"`
    V10TemplatesCreated int `json:"v10-templates-created"`
    V10TemplatesDeleted int `json:"v10-templates-deleted"`
    TemplateDropExceeded int `json:"template-drop-exceeded"`
    TemplateDropOutOfMemory int `json:"template-drop-out-of-memory"`
    FragDropped int `json:"frag-dropped"`
    AgentNotFound int `json:"agent-not-found"`
    VersionNotSupported int `json:"version-not-supported"`
    UnknownDir int `json:"unknown-dir"`
}


type VisibilityFlowCollectorNetflowStatsTemplate struct {
    Stats VisibilityFlowCollectorNetflowStatsTemplateStats `json:"stats"`
}


type VisibilityFlowCollectorNetflowStatsTemplateStats struct {
    TemplatesAddedToDelq int `json:"templates-added-to-delq"`
    TemplatesRemovedFromDelq int `json:"templates-removed-from-delq"`
}

func (p *VisibilityFlowCollectorNetflowStats) GetId() string{
    return "1"
}

func (p *VisibilityFlowCollectorNetflowStats) getPath() string{
    return "visibility/flow-collector/netflow/stats"
}

func (p *VisibilityFlowCollectorNetflowStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityFlowCollectorNetflowStats,error) {
logger.Println("VisibilityFlowCollectorNetflowStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityFlowCollectorNetflowStats
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
