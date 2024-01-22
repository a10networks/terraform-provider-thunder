

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdPerformanceOper struct {
    
    Oper RrdPerformanceOperOper `json:"oper"`

}
type DataRrdPerformanceOper struct {
    DtRrdPerformanceOper RrdPerformanceOper `json:"performance"`
}


type RrdPerformanceOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    PerformanceData []RrdPerformanceOperOperPerformanceData `json:"performance-data"`
}


type RrdPerformanceOperOperPerformanceData struct {
    Time int `json:"time"`
    Ov_puts int `json:"ov_puts"`
    Ov_cur_conns int `json:"ov_cur_conns"`
    Ov_new_conn_l4 int `json:"ov_new_conn_l4"`
    Ov_new_conn_l7 int `json:"ov_new_conn_l7"`
    Ov_new_conn_ipnat int `json:"ov_new_conn_ipnat"`
    Ov_new_conn_ssl int `json:"ov_new_conn_ssl"`
    Ov_new_conn_srv_ssl int `json:"ov_new_conn_srv_ssl"`
    Ov_new_conn_tot int `json:"ov_new_conn_tot"`
    Ov_l7_req int `json:"ov_l7_req"`
    Rus_c_conns_tot int `json:"rus_c_conns_tot"`
    Rus_c_conns_cur int `json:"rus_c_conns_cur"`
    Rus_s_conns_tot int `json:"rus_s_conns_tot"`
    Rus_s_conns_cur int `json:"rus_s_conns_cur"`
    Rus_s_conns_act int `json:"rus_s_conns_act"`
    Syn_recv int `json:"syn_recv"`
    Syn_fail int `json:"syn_fail"`
    Lsn_cps int `json:"lsn_cps"`
    Lsn_used_sess int `json:"lsn_used_sess"`
    Lsn_avail_sess int `json:"lsn_avail_sess"`
    Lsn_tcp_port_used int `json:"lsn_tcp_port_used"`
    Lsn_tcp_port_avail int `json:"lsn_tcp_port_avail"`
    Lsn_udp_port_used int `json:"lsn_udp_port_used"`
    Lsn_udp_port_avail int `json:"lsn_udp_port_avail"`
    Ov_tcp_cur_conns int `json:"ov_tcp_cur_conns"`
    Ov_sctp_cur_conns int `json:"ov_sctp_cur_conns"`
    Ov_udp_cur_conns int `json:"ov_udp_cur_conns"`
    Ov_ip_cur_conns int `json:"ov_ip_cur_conns"`
    Ov_other_cur_conns int `json:"ov_other_cur_conns"`
}

func (p *RrdPerformanceOper) GetId() string{
    return "1"
}

func (p *RrdPerformanceOper) getPath() string{
    return "rrd/performance/oper"
}

func (p *RrdPerformanceOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdPerformanceOper,error) {
logger.Println("RrdPerformanceOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdPerformanceOper
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
