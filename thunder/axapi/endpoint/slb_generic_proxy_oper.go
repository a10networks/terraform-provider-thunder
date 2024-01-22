

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbGenericProxyOper struct {
    
    Oper SlbGenericProxyOperOper `json:"oper"`

}
type DataSlbGenericProxyOper struct {
    DtSlbGenericProxyOper SlbGenericProxyOper `json:"generic-proxy"`
}


type SlbGenericProxyOperOper struct {
    GenericProxyCpuList []SlbGenericProxyOperOperGenericProxyCpuList `json:"generic-proxy-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbGenericProxyOperOperGenericProxyCpuList struct {
    Curr_proxy_conns int `json:"curr_proxy_conns"`
    Total_proxy_conns int `json:"total_proxy_conns"`
    Total_http_conn_generic_proxy int `json:"total_http_conn_generic_proxy"`
    Client_fail int `json:"client_fail"`
    Server_fail int `json:"server_fail"`
    Server_selection_fail int `json:"server_selection_fail"`
    No_route_fail int `json:"no_route_fail"`
    Source_nat_fail int `json:"source_nat_fail"`
    User_session string `json:"user_session"`
    Acr_out int `json:"acr_out"`
    Acr_in int `json:"acr_in"`
    Aca_out int `json:"aca_out"`
    Aca_in int `json:"aca_in"`
    Dpr_out int `json:"dpr_out"`
    Dpr_in int `json:"dpr_in"`
    Dpa_out int `json:"dpa_out"`
    Dpa_in int `json:"dpa_in"`
    Cea_out int `json:"cea_out"`
    Cea_in int `json:"cea_in"`
    Cer_out int `json:"cer_out"`
    Cer_in int `json:"cer_in"`
    Dwa_out int `json:"dwa_out"`
    Dwa_in int `json:"dwa_in"`
    Dwr_out int `json:"dwr_out"`
    Dwr_in int `json:"dwr_in"`
    Str_out int `json:"str_out"`
    Str_in int `json:"str_in"`
    Sta_out int `json:"sta_out"`
    Sta_in int `json:"sta_in"`
    Asr_out int `json:"asr_out"`
    Asr_in int `json:"asr_in"`
    Asa_out int `json:"asa_out"`
    Asa_in int `json:"asa_in"`
    Other_out int `json:"other_out"`
    Other_in int `json:"other_in"`
    Mismatch_fwd_id int `json:"mismatch_fwd_id"`
    Mismatch_rev_id int `json:"mismatch_rev_id"`
    Unkwn_cmd_code int `json:"unkwn_cmd_code"`
    No_session_id int `json:"no_session_id"`
    No_fwd_tuple int `json:"no_fwd_tuple"`
    No_rev_tuple int `json:"no_rev_tuple"`
    Dcmsg_fwd_in int `json:"dcmsg_fwd_in"`
    Dcmsg_fwd_out int `json:"dcmsg_fwd_out"`
    Dcmsg_rev_in int `json:"dcmsg_rev_in"`
    Dcmsg_rev_out int `json:"dcmsg_rev_out"`
    Dcmsg_error int `json:"dcmsg_error"`
    Retry_client_request int `json:"retry_client_request"`
    Retry_client_request_fail int `json:"retry_client_request_fail"`
    Reply_unknown_session_id int `json:"reply_unknown_session_id"`
    Ccr_out int `json:"ccr_out"`
    Ccr_in int `json:"ccr_in"`
    Cca_out int `json:"cca_out"`
    Cca_in int `json:"cca_in"`
    Ccr_i int `json:"ccr_i"`
    Ccr_u int `json:"ccr_u"`
    Ccr_t int `json:"ccr_t"`
    Cca_t int `json:"cca_t"`
    Terminate_on_cca_t int `json:"terminate_on_cca_t"`
    Forward_unknown_session_id int `json:"forward_unknown_session_id"`
    Update_latest_server int `json:"update_latest_server"`
    Client_select_fail int `json:"client_select_fail"`
    Close_conn_when_vport_down int `json:"close_conn_when_vport_down"`
    Invalid_avp int `json:"invalid_avp"`
    Reselect_fwd_tuple int `json:"reselect_fwd_tuple"`
    Reselect_fwd_tuple_other_cpu int `json:"reselect_fwd_tuple_other_cpu"`
    Reselect_rev_tuple int `json:"reselect_rev_tuple"`
    Conn_closed_by_client int `json:"conn_closed_by_client"`
    Conn_closed_by_server int `json:"conn_closed_by_server"`
    Reply_invalid_avp_value int `json:"reply_invalid_avp_value"`
    Reply_unable_to_deliver int `json:"reply_unable_to_deliver"`
    Reply_error_info_fail int `json:"reply_error_info_fail"`
}

func (p *SlbGenericProxyOper) GetId() string{
    return "1"
}

func (p *SlbGenericProxyOper) getPath() string{
    return "slb/generic-proxy/oper"
}

func (p *SlbGenericProxyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbGenericProxyOper,error) {
logger.Println("SlbGenericProxyOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbGenericProxyOper
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
