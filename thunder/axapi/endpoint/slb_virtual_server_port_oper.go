

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerPortOper struct {
    
    Oper SlbVirtualServerPortOperOper `json:"oper"`
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Name string 

}
type DataSlbVirtualServerPortOper struct {
    DtSlbVirtualServerPortOper SlbVirtualServerPortOper `json:"port"`
}


type SlbVirtualServerPortOperOper struct {
    State string `json:"state"`
    Real_curr_conn int `json:"real_curr_conn"`
    Int_curr_conn int `json:"int_curr_conn"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Print_extended_stats int `json:"print_extended_stats"`
    Loc_list string `json:"loc_list"`
    Geo_location string `json:"geo_location"`
    Level_str string `json:"level_str"`
    Group_id int `json:"group_id"`
    Loc_max_depth int `json:"loc_max_depth"`
    Loc_success int `json:"loc_success"`
    Loc_error int `json:"loc_error"`
    Loc_override int `json:"loc_override"`
    Loc_last string `json:"loc_last"`
    HttpHitsList []SlbVirtualServerPortOperOperHttpHitsList `json:"http-hits-list"`
    HttpVportCpuList []SlbVirtualServerPortOperOperHttpVportCpuList `json:"http-vport-cpu-list"`
    CpuCount int `json:"cpu-count"`
    Http_host_hits int `json:"http_host_hits"`
    Http_url_hits int `json:"http_url_hits"`
    Http_vport int `json:"http_vport"`
    Clear_curr_conn int `json:"clear_curr_conn"`
}


type SlbVirtualServerPortOperOperHttpHitsList struct {
    Name string `json:"name"`
    HitsCount int `json:"hits-count"`
}


type SlbVirtualServerPortOperOperHttpVportCpuList struct {
    Status_200 int `json:"status_200"`
    Status_201 int `json:"status_201"`
    Status_202 int `json:"status_202"`
    Status_203 int `json:"status_203"`
    Status_204 int `json:"status_204"`
    Status_205 int `json:"status_205"`
    Status_206 int `json:"status_206"`
    Status_207 int `json:"status_207"`
    Status_100 int `json:"status_100"`
    Status_101 int `json:"status_101"`
    Status_102 int `json:"status_102"`
    Status_300 int `json:"status_300"`
    Status_301 int `json:"status_301"`
    Status_302 int `json:"status_302"`
    Status_303 int `json:"status_303"`
    Status_304 int `json:"status_304"`
    Status_305 int `json:"status_305"`
    Status_306 int `json:"status_306"`
    Status_307 int `json:"status_307"`
    Status_400 int `json:"status_400"`
    Status_401 int `json:"status_401"`
    Status_402 int `json:"status_402"`
    Status_403 int `json:"status_403"`
    Status_404 int `json:"status_404"`
    Status_405 int `json:"status_405"`
    Status_406 int `json:"status_406"`
    Status_407 int `json:"status_407"`
    Status_408 int `json:"status_408"`
    Status_409 int `json:"status_409"`
    Status_410 int `json:"status_410"`
    Status_411 int `json:"status_411"`
    Status_412 int `json:"status_412"`
    Status_413 int `json:"status_413"`
    Status_414 int `json:"status_414"`
    Status_415 int `json:"status_415"`
    Status_416 int `json:"status_416"`
    Status_417 int `json:"status_417"`
    Status_418 int `json:"status_418"`
    Status_422 int `json:"status_422"`
    Status_423 int `json:"status_423"`
    Status_424 int `json:"status_424"`
    Status_425 int `json:"status_425"`
    Status_426 int `json:"status_426"`
    Status_449 int `json:"status_449"`
    Status_450 int `json:"status_450"`
    Status_500 int `json:"status_500"`
    Status_501 int `json:"status_501"`
    Status_502 int `json:"status_502"`
    Status_503 int `json:"status_503"`
    Status_504 int `json:"status_504"`
    Status_504_ax int `json:"status_504_ax"`
    Status_505 int `json:"status_505"`
    Status_506 int `json:"status_506"`
    Status_507 int `json:"status_507"`
    Status_508 int `json:"status_508"`
    Status_509 int `json:"status_509"`
    Status_510 int `json:"status_510"`
    Status_1xx int `json:"status_1xx"`
    Status_2xx int `json:"status_2xx"`
    Status_3xx int `json:"status_3xx"`
    Status_4xx int `json:"status_4xx"`
    Status_5xx int `json:"status_5xx"`
    Status_6xx int `json:"status_6xx"`
    Status_unknown int `json:"status_unknown"`
    Ws_handshake_request int `json:"ws_handshake_request"`
    Ws_handshake_success int `json:"ws_handshake_success"`
    Ws_client_switch int `json:"ws_client_switch"`
    Ws_server_switch int `json:"ws_server_switch"`
    Req_10u int `json:"REQ_10u"`
    Req_20u int `json:"REQ_20u"`
    Req_50u int `json:"REQ_50u"`
    Req_100u int `json:"REQ_100u"`
    Req_200u int `json:"REQ_200u"`
    Req_500u int `json:"REQ_500u"`
    Req_1m int `json:"REQ_1m"`
    Req_2m int `json:"REQ_2m"`
    Req_5m int `json:"REQ_5m"`
    Req_10m int `json:"REQ_10m"`
    Req_20m int `json:"REQ_20m"`
    Req_50m int `json:"REQ_50m"`
    Req_100m int `json:"REQ_100m"`
    Req_200m int `json:"REQ_200m"`
    Req_500m int `json:"REQ_500m"`
    Req_1s int `json:"REQ_1s"`
    Req_2s int `json:"REQ_2s"`
    Req_5s int `json:"REQ_5s"`
    Req_over_5s int `json:"REQ_OVER_5s"`
    Curr_http2_conn int `json:"curr_http2_conn"`
    Total_http2_conn int `json:"total_http2_conn"`
    Peak_http2_conn int `json:"peak_http2_conn"`
    Total_http2_bytes int `json:"total_http2_bytes"`
    Http2_control_bytes int `json:"http2_control_bytes"`
    Http2_header_bytes int `json:"http2_header_bytes"`
    Http2_data_bytes int `json:"http2_data_bytes"`
    Http2_reset_received int `json:"http2_reset_received"`
    Http2_reset_sent int `json:"http2_reset_sent"`
    Http2_goaway_received int `json:"http2_goaway_received"`
    Http2_goaway_sent int `json:"http2_goaway_sent"`
    Stream_closed int `json:"stream_closed"`
    Jsi_requests int `json:"jsi_requests"`
    Jsi_responses int `json:"jsi_responses"`
    Jsi_pri_requests int `json:"jsi_pri_requests"`
    Jsi_api_requests int `json:"jsi_api_requests"`
    Jsi_api_responses int `json:"jsi_api_responses"`
    Jsi_api_no_auth_hdr int `json:"jsi_api_no_auth_hdr"`
    Jsi_api_no_token int `json:"jsi_api_no_token"`
    Jsi_skip_no_fi int `json:"jsi_skip_no_fi"`
    Jsi_skip_no_ua int `json:"jsi_skip_no_ua"`
    Jsi_skip_not_browser int `json:"jsi_skip_not_browser"`
    Jsi_hash_add_fails int `json:"jsi_hash_add_fails"`
    Jsi_hash_lookup_fails int `json:"jsi_hash_lookup_fails"`
    Header_length_long int `json:"header_length_long"`
    Req_get int `json:"req_get"`
    Req_head int `json:"req_head"`
    Req_put int `json:"req_put"`
    Req_post int `json:"req_post"`
    Req_trace int `json:"req_trace"`
    Req_options int `json:"req_options"`
    Req_connect int `json:"req_connect"`
    Req_delete int `json:"req_delete"`
    Req_unknown int `json:"req_unknown"`
    Req_track int `json:"req_track"`
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
    Req_sz_1k int `json:"req_sz_1k"`
    Req_sz_2k int `json:"req_sz_2k"`
    Req_sz_4k int `json:"req_sz_4k"`
    Req_sz_8k int `json:"req_sz_8k"`
    Req_sz_16k int `json:"req_sz_16k"`
    Req_sz_32k int `json:"req_sz_32k"`
    Req_sz_64k int `json:"req_sz_64k"`
    Req_sz_256k int `json:"req_sz_256k"`
    Req_sz_gt_256k int `json:"req_sz_gt_256k"`
    Req_content_len int `json:"req_content_len"`
    Rsp_chunk int `json:"rsp_chunk"`
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
}

func (p *SlbVirtualServerPortOper) GetId() string{
    return "1"
}

func (p *SlbVirtualServerPortOper) getPath() string{
    
    return "slb/virtual-server/" +p.Name + "/port/" +strconv.Itoa(p.PortNumber)+"+"+p.Protocol+"/oper"
}

func (p *SlbVirtualServerPortOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbVirtualServerPortOper,error) {
logger.Println("SlbVirtualServerPortOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbVirtualServerPortOper
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
