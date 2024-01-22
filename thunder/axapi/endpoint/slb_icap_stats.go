

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbIcapStats struct {
    
    Stats SlbIcapStatsStats `json:"stats"`

}
type DataSlbIcapStats struct {
    DtSlbIcapStats SlbIcapStats `json:"icap"`
}


type SlbIcapStatsStats struct {
    Reqmod_request int `json:"reqmod_request"`
    Respmod_request int `json:"respmod_request"`
    Reqmod_request_after_100 int `json:"reqmod_request_after_100"`
    Respmod_request_after_100 int `json:"respmod_request_after_100"`
    Reqmod_response int `json:"reqmod_response"`
    Respmod_response int `json:"respmod_response"`
    Reqmod_response_after_100 int `json:"reqmod_response_after_100"`
    Respmod_response_after_100 int `json:"respmod_response_after_100"`
    Chunk_no_allow_204 int `json:"chunk_no_allow_204"`
    Len_exceed_no_allow_204 int `json:"len_exceed_no_allow_204"`
    Result_continue int `json:"result_continue"`
    Result_icap_response int `json:"result_icap_response"`
    Result_100_continue int `json:"result_100_continue"`
    Result_other int `json:"result_other"`
    Status_2xx int `json:"status_2xx"`
    Status_200 int `json:"status_200"`
    Status_201 int `json:"status_201"`
    Status_202 int `json:"status_202"`
    Status_203 int `json:"status_203"`
    Status_204 int `json:"status_204"`
    Status_205 int `json:"status_205"`
    Status_206 int `json:"status_206"`
    Status_207 int `json:"status_207"`
    Status_1xx int `json:"status_1xx"`
    Status_100 int `json:"status_100"`
    Status_101 int `json:"status_101"`
    Status_102 int `json:"status_102"`
    Status_3xx int `json:"status_3xx"`
    Status_300 int `json:"status_300"`
    Status_301 int `json:"status_301"`
    Status_302 int `json:"status_302"`
    Status_303 int `json:"status_303"`
    Status_304 int `json:"status_304"`
    Status_305 int `json:"status_305"`
    Status_306 int `json:"status_306"`
    Status_307 int `json:"status_307"`
    Status_4xx int `json:"status_4xx"`
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
    Status_419 int `json:"status_419"`
    Status_420 int `json:"status_420"`
    Status_422 int `json:"status_422"`
    Status_423 int `json:"status_423"`
    Status_424 int `json:"status_424"`
    Status_425 int `json:"status_425"`
    Status_426 int `json:"status_426"`
    Status_449 int `json:"status_449"`
    Status_450 int `json:"status_450"`
    Status_5xx int `json:"status_5xx"`
    Status_500 int `json:"status_500"`
    Status_501 int `json:"status_501"`
    Status_502 int `json:"status_502"`
    Status_503 int `json:"status_503"`
    Status_504 int `json:"status_504"`
    Status_505 int `json:"status_505"`
    Status_506 int `json:"status_506"`
    Status_507 int `json:"status_507"`
    Status_508 int `json:"status_508"`
    Status_509 int `json:"status_509"`
    Status_510 int `json:"status_510"`
    Status_6xx int `json:"status_6xx"`
    Status_unknown int `json:"status_unknown"`
    Send_option_req int `json:"send_option_req"`
    App_serv_conn_no_pcb_err int `json:"app_serv_conn_no_pcb_err"`
    App_serv_conn_err int `json:"app_serv_conn_err"`
    Chunk1_hdr_err int `json:"chunk1_hdr_err"`
    Chunk2_hdr_err int `json:"chunk2_hdr_err"`
    Chunk_bad_trail_err int `json:"chunk_bad_trail_err"`
    No_payload_next_buff_err int `json:"no_payload_next_buff_err"`
    No_payload_buff_err int `json:"no_payload_buff_err"`
    Resp_hdr_incomplete_err int `json:"resp_hdr_incomplete_err"`
    Serv_sel_fail_err int `json:"serv_sel_fail_err"`
    Start_icap_conn_fail_err int `json:"start_icap_conn_fail_err"`
    Prep_req_fail_err int `json:"prep_req_fail_err"`
    Icap_ver_err int `json:"icap_ver_err"`
    Icap_line_err int `json:"icap_line_err"`
    Encap_hdr_incomplete_err int `json:"encap_hdr_incomplete_err"`
    No_icap_resp_err int `json:"no_icap_resp_err"`
    Resp_line_read_err int `json:"resp_line_read_err"`
    Resp_line_parse_err int `json:"resp_line_parse_err"`
    Resp_hdr_err int `json:"resp_hdr_err"`
    Req_hdr_incomplete_err int `json:"req_hdr_incomplete_err"`
    No_status_code_err int `json:"no_status_code_err"`
    Http_resp_line_read_err int `json:"http_resp_line_read_err"`
    Http_resp_line_parse_err int `json:"http_resp_line_parse_err"`
    Http_resp_hdr_err int `json:"http_resp_hdr_err"`
    Recv_option_resp int `json:"recv_option_resp"`
    Http_connect_reqmod_request int `json:"http_connect_reqmod_request"`
}

func (p *SlbIcapStats) GetId() string{
    return "1"
}

func (p *SlbIcapStats) getPath() string{
    return "slb/icap/stats"
}

func (p *SlbIcapStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbIcapStats,error) {
logger.Println("SlbIcapStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbIcapStats
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
