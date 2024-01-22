

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbProxyStats struct {
    
    Stats SlbProxyStatsStats `json:"stats"`

}
type DataSlbProxyStats struct {
    DtSlbProxyStats SlbProxyStats `json:"proxy"`
}


type SlbProxyStatsStats struct {
    Tcp_event int `json:"tcp_event"`
    Est_event int `json:"est_event"`
    Data_event int `json:"data_event"`
    Client_fin int `json:"client_fin"`
    Server_fin int `json:"server_fin"`
    Wbuf_event int `json:"wbuf_event"`
    Err_event int `json:"err_event"`
    No_mem int `json:"no_mem"`
    Client_rst int `json:"client_rst"`
    Server_rst int `json:"server_rst"`
    Queue_depth_over_limit int `json:"queue_depth_over_limit"`
    Event_failed int `json:"event_failed"`
    Conn_not_exist int `json:"conn_not_exist"`
    Service_alloc_cb int `json:"service_alloc_cb"`
    Service_alloc_cb_failed int `json:"service_alloc_cb_failed"`
    Service_free_cb int `json:"service_free_cb"`
    Service_free_cb_failed int `json:"service_free_cb_failed"`
    Est_cb_failed int `json:"est_cb_failed"`
    Data_cb_failed int `json:"data_cb_failed"`
    Wbuf_cb_failed int `json:"wbuf_cb_failed"`
    Err_cb_failed int `json:"err_cb_failed"`
    Start_server_conn int `json:"start_server_conn"`
    Start_server_conn_succ int `json:"start_server_conn_succ"`
    Start_server_conn_no_route int `json:"start_server_conn_no_route"`
    Start_server_conn_fail_mem int `json:"start_server_conn_fail_mem"`
    Start_server_conn_fail_snat int `json:"start_server_conn_fail_snat"`
    Start_server_conn_fail_persist int `json:"start_server_conn_fail_persist"`
    Start_server_conn_fail_server int `json:"start_server_conn_fail_server"`
    Start_server_conn_fail_tuple int `json:"start_server_conn_fail_tuple"`
    Line_too_long int `json:"line_too_long"`
}

func (p *SlbProxyStats) GetId() string{
    return "1"
}

func (p *SlbProxyStats) getPath() string{
    return "slb/proxy/stats"
}

func (p *SlbProxyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbProxyStats,error) {
logger.Println("SlbProxyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbProxyStats
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
