

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbLinkProbeEntryStats struct {
    
    Stats SlbLinkProbeEntryStatsStats `json:"stats"`

}
type DataSlbLinkProbeEntryStats struct {
    DtSlbLinkProbeEntryStats SlbLinkProbeEntryStats `json:"entry"`
}


type SlbLinkProbeEntryStatsStats struct {
    Curr_entries int `json:"curr_entries"`
    Total_created int `json:"total_created"`
    Total_inserted int `json:"total_inserted"`
    Total_ready_to_free int `json:"total_ready_to_free"`
    Total_freed int `json:"total_freed"`
    Err_entry_create_failed int `json:"err_entry_create_failed"`
    Err_entry_create_oom int `json:"err_entry_create_oom"`
    Err_entry_insert_failed int `json:"err_entry_insert_failed"`
    Err_tmpl_probe_create_failed int `json:"err_tmpl_probe_create_failed"`
    Err_tmpl_probe_create_oom int `json:"err_tmpl_probe_create_oom"`
    Total_http_probes_sent int `json:"total_http_probes_sent"`
    Total_http_response_received int `json:"total_http_response_received"`
    Total_http_response_good int `json:"total_http_response_good"`
    Total_http_response_bad int `json:"total_http_response_bad"`
    Total_tcp_err int `json:"total_tcp_err"`
    Err_smart_nat_alloc int `json:"err_smart_nat_alloc"`
    Err_smart_nat_port_alloc int `json:"err_smart_nat_port_alloc"`
    Err_l4_sess_alloc int `json:"err_l4_sess_alloc"`
    Err_probe_tcp_conn_send int `json:"err_probe_tcp_conn_send"`
    Probe_tcp_conn_sent int `json:"probe_tcp_conn_sent"`
}

func (p *SlbLinkProbeEntryStats) GetId() string{
    return "1"
}

func (p *SlbLinkProbeEntryStats) getPath() string{
    return "slb/link-probe/entry/stats"
}

func (p *SlbLinkProbeEntryStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbLinkProbeEntryStats,error) {
logger.Println("SlbLinkProbeEntryStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbLinkProbeEntryStats
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
