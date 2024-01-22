

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPassthroughStats struct {
    
    Stats SlbPassthroughStatsStats `json:"stats"`

}
type DataSlbPassthroughStats struct {
    DtSlbPassthroughStats SlbPassthroughStats `json:"passthrough"`
}


type SlbPassthroughStatsStats struct {
    Curr_conn int `json:"curr_conn"`
    Total_conn int `json:"total_conn"`
    Total_fwd_bytes int `json:"total_fwd_bytes"`
    Total_fwd_packets int `json:"total_fwd_packets"`
    Total_rev_bytes int `json:"total_rev_bytes"`
    Total_rev_packets int `json:"total_rev_packets"`
    Curr_pconn int `json:"curr_pconn"`
}

func (p *SlbPassthroughStats) GetId() string{
    return "1"
}

func (p *SlbPassthroughStats) getPath() string{
    return "slb/passthrough/stats"
}

func (p *SlbPassthroughStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbPassthroughStats,error) {
logger.Println("SlbPassthroughStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbPassthroughStats
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
