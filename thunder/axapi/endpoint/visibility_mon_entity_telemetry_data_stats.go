

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonEntityTelemetryDataStats struct {
    
    Stats VisibilityMonEntityTelemetryDataStatsStats `json:"stats"`

}
type DataVisibilityMonEntityTelemetryDataStats struct {
    DtVisibilityMonEntityTelemetryDataStats VisibilityMonEntityTelemetryDataStats `json:"mon-entity-telemetry-data"`
}


type VisibilityMonEntityTelemetryDataStatsStats struct {
    In_pkts int `json:"in_pkts"`
    Out_pkts int `json:"out_pkts"`
    In_bytes int `json:"in_bytes"`
    Out_bytes int `json:"out_bytes"`
    Errors int `json:"errors"`
    In_small_pkt int `json:"in_small_pkt"`
    In_frag int `json:"in_frag"`
    Out_small_pkt int `json:"out_small_pkt"`
    Out_frag int `json:"out_frag"`
    NewConn int `json:"new-conn"`
    Avg_data_cpu_util int `json:"avg_data_cpu_util"`
    Outside_intf_util int `json:"outside_intf_util"`
    ConcurrentConn int `json:"concurrent-conn"`
    In_bytes_per_out_bytes int `json:"in_bytes_per_out_bytes"`
    Drop_pkts_per_pkts int `json:"drop_pkts_per_pkts"`
    Tcp_in_syn int `json:"tcp_in_syn"`
    Tcp_out_syn int `json:"tcp_out_syn"`
    Tcp_in_fin int `json:"tcp_in_fin"`
    Tcp_out_fin int `json:"tcp_out_fin"`
    Tcp_in_payload int `json:"tcp_in_payload"`
    Tcp_out_payload int `json:"tcp_out_payload"`
    Tcp_in_rexmit int `json:"tcp_in_rexmit"`
    Tcp_out_rexmit int `json:"tcp_out_rexmit"`
    Tcp_in_rst int `json:"tcp_in_rst"`
    Tcp_out_rst int `json:"tcp_out_rst"`
    Tcp_in_empty_ack int `json:"tcp_in_empty_ack"`
    Tcp_out_empty_ack int `json:"tcp_out_empty_ack"`
    Tcp_in_zero_wnd int `json:"tcp_in_zero_wnd"`
    Tcp_out_zero_wnd int `json:"tcp_out_zero_wnd"`
    Tcp_conn_miss int `json:"tcp_conn_miss"`
    Tcp_fwd_syn_per_fin int `json:"tcp_fwd_syn_per_fin"`
}

func (p *VisibilityMonEntityTelemetryDataStats) GetId() string{
    return "1"
}

func (p *VisibilityMonEntityTelemetryDataStats) getPath() string{
    return "visibility/mon-entity-telemetry-data/stats"
}

func (p *VisibilityMonEntityTelemetryDataStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityMonEntityTelemetryDataStats,error) {
logger.Println("VisibilityMonEntityTelemetryDataStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityMonEntityTelemetryDataStats
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
