

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHttpProxyStats struct {
    
    Stats SlbHttpProxyStatsStats `json:"stats"`

}
type DataSlbHttpProxyStats struct {
    DtSlbHttpProxyStats SlbHttpProxyStats `json:"http-proxy"`
}


type SlbHttpProxyStatsStats struct {
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
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Req_retran int `json:"req_retran"`
    Req_ofo int `json:"req_ofo"`
    Server_resel int `json:"server_resel"`
    Svr_prem_close int `json:"svr_prem_close"`
    New_svrconn int `json:"new_svrconn"`
    Snat_fail int `json:"snat_fail"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Compression_before int `json:"compression_before"`
    Compression_after int `json:"compression_after"`
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
    Cache_rsp int `json:"cache_rsp"`
    Close_on_ddos int `json:"close_on_ddos"`
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
    Req_track int `json:"req_track"`
    Connect_req int `json:"connect_req"`
    Req_enter_ssli int `json:"req_enter_ssli"`
    Decompression_before int `json:"decompression_before"`
    Decompression_after int `json:"decompression_after"`
    Req_http2 int `json:"req_http2"`
    Response_http2 int `json:"response_http2"`
    Doh_req int `json:"doh_req"`
    Doh_req_get int `json:"doh_req_get"`
    Doh_req_post int `json:"doh_req_post"`
    Doh_non_doh_req int `json:"doh_non_doh_req"`
    Doh_non_doh_req_get int `json:"doh_non_doh_req_get"`
    Doh_non_doh_req_post int `json:"doh_non_doh_req_post"`
    Doh_resp int `json:"doh_resp"`
    Doh_tc_resp int `json:"doh_tc_resp"`
    Doh_udp_dns_req int `json:"doh_udp_dns_req"`
    Doh_udp_dns_resp int `json:"doh_udp_dns_resp"`
    Doh_tcp_dns_req int `json:"doh_tcp_dns_req"`
    Doh_tcp_dns_resp int `json:"doh_tcp_dns_resp"`
    Doh_req_send_failed int `json:"doh_req_send_failed"`
    Doh_resp_send_failed int `json:"doh_resp_send_failed"`
    Doh_malloc_fail int `json:"doh_malloc_fail"`
    Doh_req_udp_retry int `json:"doh_req_udp_retry"`
    Doh_req_udp_retry_fail int `json:"doh_req_udp_retry_fail"`
    Doh_req_tcp_retry int `json:"doh_req_tcp_retry"`
    Doh_req_tcp_retry_fail int `json:"doh_req_tcp_retry_fail"`
    Doh_snat_failed int `json:"doh_snat_failed"`
    Doh_path_not_found int `json:"doh_path_not_found"`
    Doh_get_dns_arg_failed int `json:"doh_get_dns_arg_failed"`
    Doh_get_base64_decode_failed int `json:"doh_get_base64_decode_failed"`
    Doh_post_content_type_mismatch int `json:"doh_post_content_type_mismatch"`
    Doh_post_payload_not_found int `json:"doh_post_payload_not_found"`
    Doh_post_payload_extract_failed int `json:"doh_post_payload_extract_failed"`
    Doh_non_doh_method int `json:"doh_non_doh_method"`
    Doh_tcp_send_failed int `json:"doh_tcp_send_failed"`
    Doh_udp_send_failed int `json:"doh_udp_send_failed"`
    Doh_query_time_out int `json:"doh_query_time_out"`
    Doh_dns_query_type_a int `json:"doh_dns_query_type_a"`
    Doh_dns_query_type_aaaa int `json:"doh_dns_query_type_aaaa"`
    Doh_dns_query_type_ns int `json:"doh_dns_query_type_ns"`
    Doh_dns_query_type_cname int `json:"doh_dns_query_type_cname"`
    Doh_dns_query_type_any int `json:"doh_dns_query_type_any"`
    Doh_dns_query_type_srv int `json:"doh_dns_query_type_srv"`
    Doh_dns_query_type_mx int `json:"doh_dns_query_type_mx"`
    Doh_dns_query_type_soa int `json:"doh_dns_query_type_soa"`
    Doh_dns_query_type_others int `json:"doh_dns_query_type_others"`
    Doh_resp_setup_failed int `json:"doh_resp_setup_failed"`
    Doh_resp_header_alloc_failed int `json:"doh_resp_header_alloc_failed"`
    Doh_resp_que_failed int `json:"doh_resp_que_failed"`
    Doh_resp_udp_frags int `json:"doh_resp_udp_frags"`
    Doh_resp_tcp_frags int `json:"doh_resp_tcp_frags"`
    Doh_serv_sel_failed int `json:"doh_serv_sel_failed"`
    Doh_retry_w_tcp int `json:"doh_retry_w_tcp"`
    Doh_get_uri_too_long int `json:"doh_get_uri_too_long"`
    Doh_post_payload_too_large int `json:"doh_post_payload_too_large"`
    Doh_dns_malformed_query int `json:"doh_dns_malformed_query"`
    Doh_dns_resp_rcode_err_format int `json:"doh_dns_resp_rcode_err_format"`
    Doh_dns_resp_rcode_err_server int `json:"doh_dns_resp_rcode_err_server"`
    Doh_dns_resp_rcode_err_name int `json:"doh_dns_resp_rcode_err_name"`
    Doh_dns_resp_rcode_err_type int `json:"doh_dns_resp_rcode_err_type"`
    Doh_dns_resp_rcode_refuse int `json:"doh_dns_resp_rcode_refuse"`
    Doh_dns_resp_rcode_yxdomain int `json:"doh_dns_resp_rcode_yxdomain"`
    Doh_dns_resp_rcode_yxrrset int `json:"doh_dns_resp_rcode_yxrrset"`
    Doh_dns_resp_rcode_nxrrset int `json:"doh_dns_resp_rcode_nxrrset"`
    Doh_dns_resp_rcode_notauth int `json:"doh_dns_resp_rcode_notauth"`
    Doh_dns_resp_rcode_notzone int `json:"doh_dns_resp_rcode_notzone"`
    Doh_dns_resp_rcode_other int `json:"doh_dns_resp_rcode_other"`
    Compression_before_br int `json:"compression_before_br"`
    Compression_after_br int `json:"compression_after_br"`
    Compression_before_total int `json:"compression_before_total"`
    Compression_after_total int `json:"compression_after_total"`
    Decompression_before_br int `json:"decompression_before_br"`
    Decompression_after_br int `json:"decompression_after_br"`
    Decompression_before_total int `json:"decompression_before_total"`
    Decompression_after_total int `json:"decompression_after_total"`
    Req_http3 int `json:"req_http3"`
    Response_http3 int `json:"response_http3"`
}

func (p *SlbHttpProxyStats) GetId() string{
    return "1"
}

func (p *SlbHttpProxyStats) getPath() string{
    return "slb/http-proxy/stats"
}

func (p *SlbHttpProxyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHttpProxyStats,error) {
logger.Println("SlbHttpProxyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHttpProxyStats
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
