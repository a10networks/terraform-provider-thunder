

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityFlowCollectorSflowStats struct {
    
    Stats VisibilityFlowCollectorSflowStatsStats `json:"stats"`

}
type DataVisibilityFlowCollectorSflowStats struct {
    DtVisibilityFlowCollectorSflowStats VisibilityFlowCollectorSflowStats `json:"sflow"`
}


type VisibilityFlowCollectorSflowStatsStats struct {
    PktsReceived int `json:"pkts-received"`
    FragDropped int `json:"frag-dropped"`
    AgentNotFound int `json:"agent-not-found"`
    VersionNotSupported int `json:"version-not-supported"`
    UnknownDir int `json:"unknown-dir"`
}

func (p *VisibilityFlowCollectorSflowStats) GetId() string{
    return "1"
}

func (p *VisibilityFlowCollectorSflowStats) getPath() string{
    return "visibility/flow-collector/sflow/stats"
}

func (p *VisibilityFlowCollectorSflowStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityFlowCollectorSflowStats,error) {
logger.Println("VisibilityFlowCollectorSflowStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityFlowCollectorSflowStats
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
