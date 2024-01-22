

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwPerInstanceStats struct {
    
    Stats FwPerInstanceStatsStats `json:"stats"`

}
type DataFwPerInstanceStats struct {
    DtFwPerInstanceStats FwPerInstanceStats `json:"per-instance"`
}


type FwPerInstanceStatsStats struct {
    KeyName string `json:"key-name"`
    Data_session_created int `json:"data_session_created"`
    Data_session_freed int `json:"data_session_freed"`
    Tcp_fullcone_created int `json:"tcp_fullcone_created"`
    Tcp_fullcone_freed int `json:"tcp_fullcone_freed"`
    Udp_fullcone_created int `json:"udp_fullcone_created"`
    Udp_fullcone_freed int `json:"udp_fullcone_freed"`
    Dyn_blist_sess_created int `json:"dyn_blist_sess_created"`
    Dyn_blist_sess_freed int `json:"dyn_blist_sess_freed"`
    Dyn_blist_pkt_drop int `json:"dyn_blist_pkt_drop"`
    Dyn_blist_version_mismatch int `json:"dyn_blist_version_mismatch"`
    Dyn_blist_pkt_rate_low int `json:"dyn_blist_pkt_rate_low"`
    Dyn_blist_pkt_rate_high int `json:"dyn_blist_pkt_rate_high"`
}

func (p *FwPerInstanceStats) GetId() string{
    return "1"
}

func (p *FwPerInstanceStats) getPath() string{
    return "fw/per-instance/stats"
}

func (p *FwPerInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwPerInstanceStats,error) {
logger.Println("FwPerInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwPerInstanceStats
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
