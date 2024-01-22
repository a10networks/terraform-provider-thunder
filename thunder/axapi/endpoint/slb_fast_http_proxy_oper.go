

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFastHttpProxyOper struct {
    
    Oper SlbFastHttpProxyOperOper `json:"oper"`

}
type DataSlbFastHttpProxyOper struct {
    DtSlbFastHttpProxyOper SlbFastHttpProxyOper `json:"fast-http-proxy"`
}


type SlbFastHttpProxyOperOper struct {
    FastHttpProxyCpuList []SlbFastHttpProxyOperOperFastHttpProxyCpuList `json:"fast-http-proxy-cpu-list"`
    CpuCount int `json:"cpu-count"`
    Debug_fields int `json:"debug_fields"`
}


type SlbFastHttpProxyOperOperFastHttpProxyCpuList struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Req int `json:"req"`
    Req_succ int `json:"req_succ"`
    Noproxy int `json:"noproxy"`
    Client_rst int `json:"client_rst"`
    Server_rst int `json:"server_rst"`
    Notuple int `json:"notuple"`
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreq_fail_buff int `json:"fwdreq_fail_buff"`
    Fwdreq_fail_rport int `json:"fwdreq_fail_rport"`
    Fwdreq_fail_route int `json:"fwdreq_fail_route"`
    Fwdreq_fail_persist int `json:"fwdreq_fail_persist"`
    Fwdreq_fail_server int `json:"fwdreq_fail_server"`
    Fwdreq_fail_tuple int `json:"fwdreq_fail_tuple"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Req_retran int `json:"req_retran"`
    Req_ofo int `json:"req_ofo"`
    Server_resel int `json:"server_resel"`
    Svr_prem_close int `json:"svr_prem_close"`
    New_svrconn int `json:"new_svrconn"`
    Snat_fail int `json:"snat_fail"`
    Tcpoutrst int `json:"tcpoutrst"`
    Full_proxy int `json:"full_proxy"`
    Full_proxy_post int `json:"full_proxy_post"`
    Full_proxy_put int `json:"full_proxy_put"`
    Full_proxy_pipeline int `json:"full_proxy_pipeline"`
    Full_proxy_fpga_err int `json:"full_proxy_fpga_err"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    L4_switching int `json:"l4_switching"`
    Cookie_switching int `json:"cookie_switching"`
    Aflex_switching int `json:"aflex_switching"`
    Url_switching int `json:"url_switching"`
    Host_switching int `json:"host_switching"`
    Lb_switching int `json:"lb_switching"`
    L4_switching_ok int `json:"l4_switching_ok"`
    Cookie_switching_ok int `json:"cookie_switching_ok"`
    Aflex_switching_ok int `json:"aflex_switching_ok"`
    Url_switching_ok int `json:"url_switching_ok"`
    Host_switching_ok int `json:"host_switching_ok"`
    Lb_switching_ok int `json:"lb_switching_ok"`
    L4_switching_enqueue int `json:"l4_switching_enqueue"`
    Cookie_switching_enqueue int `json:"cookie_switching_enqueue"`
    Aflex_switching_enqueue int `json:"aflex_switching_enqueue"`
    Url_switching_enqueue int `json:"url_switching_enqueue"`
    Host_switching_enqueue int `json:"host_switching_enqueue"`
    Lb_switching_enqueue int `json:"lb_switching_enqueue"`
    Retry_503 int `json:"retry_503"`
    Aflex_retry int `json:"aflex_retry"`
    Aflex_lb_reselect int `json:"aflex_lb_reselect"`
    Aflex_lb_reselect_ok int `json:"aflex_lb_reselect_ok"`
    Client_rst_request int `json:"client_rst_request"`
    Client_rst_connecting int `json:"client_rst_connecting"`
    Client_rst_connected int `json:"client_rst_connected"`
    Client_rst_response int `json:"client_rst_response"`
    Server_rst_request int `json:"server_rst_request"`
    Server_rst_connecting int `json:"server_rst_connecting"`
    Server_rst_connected int `json:"server_rst_connected"`
    Server_rst_response int `json:"server_rst_response"`
    Invalid_header int `json:"invalid_header"`
    Too_many_headers int `json:"too_many_headers"`
    Line_too_long int `json:"line_too_long"`
    Header_name_too_long int `json:"header_name_too_long"`
    Wrong_resp_header int `json:"wrong_resp_header"`
    Header_insert int `json:"header_insert"`
    Header_delete int `json:"header_delete"`
    Insert_client_ip int `json:"insert_client_ip"`
    Negative_req_remain int `json:"negative_req_remain"`
    Negative_resp_remain int `json:"negative_resp_remain"`
    Large_cookie int `json:"large_cookie"`
    Large_cookie_header int `json:"large_cookie_header"`
    Huge_cookie int `json:"huge_cookie"`
    Huge_cookie_header int `json:"huge_cookie_header"`
    Parse_cookie_fail int `json:"parse_cookie_fail"`
    Parse_setcookie_fail int `json:"parse_setcookie_fail"`
    Asm_cookie_fail int `json:"asm_cookie_fail"`
    Asm_cookie_header_fail int `json:"asm_cookie_header_fail"`
    Asm_setcookie_fail int `json:"asm_setcookie_fail"`
    Asm_setcookie_header_fail int `json:"asm_setcookie_header_fail"`
    Client_req_unexp_flag int `json:"client_req_unexp_flag"`
    Connecting_fin int `json:"connecting_fin"`
    Connecting_fin_retrans int `json:"connecting_fin_retrans"`
    Connecting_fin_ofo int `json:"connecting_fin_ofo"`
    Connecting_rst int `json:"connecting_rst"`
    Connecting_rst_retrans int `json:"connecting_rst_retrans"`
    Connecting_rst_ofo int `json:"connecting_rst_ofo"`
    Connecting_ack int `json:"connecting_ack"`
    Pkts_ofo int `json:"pkts_ofo"`
    Pkts_retrans int `json:"pkts_retrans"`
    Pkts_retrans_ack_finwait int `json:"pkts_retrans_ack_finwait"`
    Pkts_retrans_fin int `json:"pkts_retrans_fin"`
    Pkts_retrans_rst int `json:"pkts_retrans_rst"`
    Pkts_retrans_push int `json:"pkts_retrans_push"`
    Stale_sess int `json:"stale_sess"`
    Server_resel_failed int `json:"server_resel_failed"`
    Response_1xx int `json:"response_1xx"`
    Response_100 int `json:"response_100"`
    Response_101 int `json:"response_101"`
    Response_102 int `json:"response_102"`
    Response_2xx int `json:"response_2xx"`
    Response_200 int `json:"response_200"`
    Response_201 int `json:"response_201"`
    Response_202 int `json:"response_202"`
    Response_203 int `json:"response_203"`
    Response_204 int `json:"response_204"`
    Response_205 int `json:"response_205"`
    Response_206 int `json:"response_206"`
    Response_207 int `json:"response_207"`
    Response_3xx int `json:"response_3xx"`
    Response_300 int `json:"response_300"`
    Response_301 int `json:"response_301"`
    Response_302 int `json:"response_302"`
    Response_303 int `json:"response_303"`
    Response_304 int `json:"response_304"`
    Response_305 int `json:"response_305"`
    Response_306 int `json:"response_306"`
    Response_307 int `json:"response_307"`
    Response_4xx int `json:"response_4xx"`
    Response_400 int `json:"response_400"`
    Response_401 int `json:"response_401"`
    Response_402 int `json:"response_402"`
    Response_403 int `json:"response_403"`
    Response_404 int `json:"response_404"`
    Response_405 int `json:"response_405"`
    Response_406 int `json:"response_406"`
    Response_407 int `json:"response_407"`
    Response_408 int `json:"response_408"`
    Response_409 int `json:"response_409"`
    Response_410 int `json:"response_410"`
    Response_411 int `json:"response_411"`
    Response_412 int `json:"response_412"`
    Response_413 int `json:"response_413"`
    Response_414 int `json:"response_414"`
    Response_415 int `json:"response_415"`
    Response_416 int `json:"response_416"`
    Response_417 int `json:"response_417"`
    Response_418 int `json:"response_418"`
    Response_422 int `json:"response_422"`
    Response_423 int `json:"response_423"`
    Response_424 int `json:"response_424"`
    Response_425 int `json:"response_425"`
    Response_426 int `json:"response_426"`
    Response_449 int `json:"response_449"`
    Response_450 int `json:"response_450"`
    Response_5xx int `json:"response_5xx"`
    Response_500 int `json:"response_500"`
    Response_501 int `json:"response_501"`
    Response_502 int `json:"response_502"`
    Response_503 int `json:"response_503"`
    Response_504 int `json:"response_504"`
    Response_505 int `json:"response_505"`
    Response_506 int `json:"response_506"`
    Response_507 int `json:"response_507"`
    Response_508 int `json:"response_508"`
    Response_509 int `json:"response_509"`
    Response_510 int `json:"response_510"`
    Response_6xx int `json:"response_6xx"`
    Response_unknown int `json:"response_unknown"`
    Req_http10 int `json:"req_http10"`
    Req_http11 int `json:"req_http11"`
    Response_http10 int `json:"response_http10"`
    Response_http11 int `json:"response_http11"`
    Req_get int `json:"req_get"`
    Req_head int `json:"req_head"`
    Req_put int `json:"req_put"`
    Req_post int `json:"req_post"`
    Req_trace int `json:"req_trace"`
    Req_options int `json:"req_options"`
    Req_connect int `json:"req_connect"`
    Req_delete int `json:"req_delete"`
    Req_unknown int `json:"req_unknown"`
    Req_content_len int `json:"req_content_len"`
    Rsp_content_len int `json:"rsp_content_len"`
    Rsp_chunk int `json:"rsp_chunk"`
    Req_chunk int `json:"req_chunk"`
    Compress_rsp int `json:"compress_rsp"`
    Compress_rsp_br int `json:"compress_rsp_br"`
    Compress_rsp_total int `json:"compress_rsp_total"`
    Compress_del_accept_enc int `json:"compress_del_accept_enc"`
    Compress_resp_already_compressed int `json:"compress_resp_already_compressed"`
    Compress_content_type_excluded int `json:"compress_content_type_excluded"`
    Compress_no_content_type int `json:"compress_no_content_type"`
    Compress_resp_lt_min int `json:"compress_resp_lt_min"`
    Compress_resp_no_cl_or_ce int `json:"compress_resp_no_cl_or_ce"`
    Compress_ratio_too_high int `json:"compress_ratio_too_high"`
    Cache_rsp int `json:"cache_rsp"`
    Close_on_ddos int `json:"close_on_ddos"`
    Req_http10_keepalive int `json:"req_http10_keepalive"`
    Req_sz_1k int `json:"req_sz_1k"`
    Req_sz_2k int `json:"req_sz_2k"`
    Req_sz_4k int `json:"req_sz_4k"`
    Req_sz_8k int `json:"req_sz_8k"`
    Req_sz_16k int `json:"req_sz_16k"`
    Req_sz_32k int `json:"req_sz_32k"`
    Req_sz_64k int `json:"req_sz_64k"`
    Req_sz_256k int `json:"req_sz_256k"`
    Req_sz_gt_256k int `json:"req_sz_gt_256k"`
    Rsp_sz_1k int `json:"rsp_sz_1k"`
    Rsp_sz_2k int `json:"rsp_sz_2k"`
    Rsp_sz_4k int `json:"rsp_sz_4k"`
    Rsp_sz_8k int `json:"rsp_sz_8k"`
    Rsp_sz_16k int `json:"rsp_sz_16k"`
    Rsp_sz_32k int `json:"rsp_sz_32k"`
    Rsp_sz_64k int `json:"rsp_sz_64k"`
    Rsp_sz_256k int `json:"rsp_sz_256k"`
    Rsp_sz_gt_256k int `json:"rsp_sz_gt_256k"`
    Chunk_sz_512 int `json:"chunk_sz_512"`
    Chunk_sz_1k int `json:"chunk_sz_1k"`
    Chunk_sz_2k int `json:"chunk_sz_2k"`
    Chunk_sz_4k int `json:"chunk_sz_4k"`
    Chunk_sz_gt_4k int `json:"chunk_sz_gt_4k"`
    Pconn_connecting int `json:"pconn_connecting"`
    Pconn_connected int `json:"pconn_connected"`
    Pconn_connecting_failed int `json:"pconn_connecting_failed"`
    Chunk_bad int `json:"chunk_bad"`
    Req_10u int `json:"req_10u"`
    Req_20u int `json:"req_20u"`
    Req_50u int `json:"req_50u"`
    Req_100u int `json:"req_100u"`
    Req_200u int `json:"req_200u"`
    Req_500u int `json:"req_500u"`
    Req_1m int `json:"req_1m"`
    Req_2m int `json:"req_2m"`
    Req_5m int `json:"req_5m"`
    Req_10m int `json:"req_10m"`
    Req_20m int `json:"req_20m"`
    Req_50m int `json:"req_50m"`
    Req_100m int `json:"req_100m"`
    Req_200m int `json:"req_200m"`
    Req_500m int `json:"req_500m"`
    Req_1s int `json:"req_1s"`
    Req_2s int `json:"req_2s"`
    Req_5s int `json:"req_5s"`
    Req_over_5s int `json:"req_over_5s"`
    Insert_client_port int `json:"insert_client_port"`
    Req_track int `json:"req_track"`
    Non_http_bypass int `json:"non_http_bypass"`
    Skip_insert_client_ip int `json:"skip_insert_client_ip"`
    Skip_insert_client_port int `json:"skip_insert_client_port"`
    Http_pkts_ofo int `json:"http_pkts_ofo"`
}

func (p *SlbFastHttpProxyOper) GetId() string{
    return "1"
}

func (p *SlbFastHttpProxyOper) getPath() string{
    return "slb/fast-http-proxy/oper"
}

func (p *SlbFastHttpProxyOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFastHttpProxyOper,error) {
logger.Println("SlbFastHttpProxyOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFastHttpProxyOper
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
