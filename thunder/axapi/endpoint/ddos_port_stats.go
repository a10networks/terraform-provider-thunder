

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosPortStats struct {
    
    Stats DdosPortStatsStats `json:"stats"`

}
type DataDdosPortStats struct {
    DtDdosPortStats DdosPortStats `json:"port"`
}


type DdosPortStatsStats struct {
    Dst_port_undef_drop int `json:"dst_port_undef_drop"`
    Dst_port_bl int `json:"dst_port_bl"`
    Dst_port_exceed_drop_any int `json:"dst_port_exceed_drop_any"`
    Dst_port_pkt_rate_exceed int `json:"dst_port_pkt_rate_exceed"`
    Dst_port_kbit_rate_exceed int `json:"dst_port_kbit_rate_exceed"`
    Dst_port_frag_rate_exceed int `json:"dst_port_frag_rate_exceed"`
    Dst_port_conn_limit_exceed int `json:"dst_port_conn_limit_exceed"`
    Dst_port_conn_rate_exceed int `json:"dst_port_conn_rate_exceed"`
    Dst_sport_bl int `json:"dst_sport_bl"`
    Dst_sport_exceed_drop_any int `json:"dst_sport_exceed_drop_any"`
    Dst_sport_pkt_rate_exceed int `json:"dst_sport_pkt_rate_exceed"`
    Dst_sport_kbit_rate_exceed int `json:"dst_sport_kbit_rate_exceed"`
    Dst_sport_frag_rate_exceed int `json:"dst_sport_frag_rate_exceed"`
    Dst_sport_conn_limit_exceed int `json:"dst_sport_conn_limit_exceed"`
    Dst_sport_conn_rate_exceed int `json:"dst_sport_conn_rate_exceed"`
}

func (p *DdosPortStats) GetId() string{
    return "1"
}

func (p *DdosPortStats) getPath() string{
    return "ddos/port/stats"
}

func (p *DdosPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosPortStats,error) {
logger.Println("DdosPortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosPortStats
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
