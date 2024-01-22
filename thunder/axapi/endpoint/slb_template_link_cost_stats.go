

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateLinkCostStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplateLinkCostStatsStats `json:"stats"`

}
type DataSlbTemplateLinkCostStats struct {
    DtSlbTemplateLinkCostStats SlbTemplateLinkCostStats `json:"link-cost"`
}


type SlbTemplateLinkCostStatsStats struct {
    Link_total_fwd_bytes int `json:"link_total_fwd_bytes"`
    Interval_fwd_bytes int `json:"interval_fwd_bytes"`
    Link_total_conn int `json:"link_total_conn"`
    Interval_avg_throughput int `json:"interval_avg_throughput"`
    Interval_max_throughput int `json:"interval_max_throughput"`
}

func (p *SlbTemplateLinkCostStats) GetId() string{
    return "1"
}

func (p *SlbTemplateLinkCostStats) getPath() string{
    
    return "slb/template/link-cost/"+p.Name+"/stats"
}

func (p *SlbTemplateLinkCostStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplateLinkCostStats,error) {
logger.Println("SlbTemplateLinkCostStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplateLinkCostStats
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
