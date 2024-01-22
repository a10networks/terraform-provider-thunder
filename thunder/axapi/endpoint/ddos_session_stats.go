

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSessionStats struct {
    
    Stats DdosSessionStatsStats `json:"stats"`

}
type DataDdosSessionStats struct {
    DtDdosSessionStats DdosSessionStats `json:"session"`
}


type DdosSessionStatsStats struct {
    Sess_create_ip int `json:"sess_create_ip"`
    Sess_create_ip6 int `json:"sess_create_ip6"`
    Sess_create_tcp int `json:"sess_create_tcp"`
    Conn_tcp_est_w_syn int `json:"conn_tcp_est_w_syn"`
    Conn_tcp_est_w_ack int `json:"conn_tcp_est_w_ack"`
    Conn_tcp_est int `json:"conn_tcp_est"`
    Conn_tcp_close_w_rst int `json:"conn_tcp_close_w_rst"`
    Conn_tcp_close_w_fin int `json:"conn_tcp_close_w_fin"`
    Conn_tcp_close_w_idle int `json:"conn_tcp_close_w_idle"`
    Conn_tcp_close_w_half_open int `json:"conn_tcp_close_w_half_open"`
    Conn_tcp_close int `json:"conn_tcp_close"`
    Sess_create_udp int `json:"sess_create_udp"`
    Conn_udp_est int `json:"conn_udp_est"`
    Conn_udp_close int `json:"conn_udp_close"`
    Sess_aged_out int `json:"sess_aged_out"`
    Conn_entry_mismatch int `json:"conn_entry_mismatch"`
    Conn_tcp_create_from_syn int `json:"conn_tcp_create_from_syn"`
    Conn_tcp_create_from_ack int `json:"conn_tcp_create_from_ack"`
    Sess_oom int `json:"sess_oom"`
    Sess_create_udp_auth int `json:"sess_create_udp_auth"`
    Sess_aged_udp_auth int `json:"sess_aged_udp_auth"`
    Sess_snat_failed int `json:"sess_snat_failed"`
}

func (p *DdosSessionStats) GetId() string{
    return "1"
}

func (p *DdosSessionStats) getPath() string{
    return "ddos/session/stats"
}

func (p *DdosSessionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosSessionStats,error) {
logger.Println("DdosSessionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosSessionStats
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
