

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate struct {
	Inst struct {

    Alloc_fail_total int `json:"alloc_fail_total"`
    Bad_connection_preface int `json:"bad_connection_preface"`
    Bad_frame_type_for_stream_state int `json:"bad_frame_type_for_stream_state"`
    Buff_alloc_error int `json:"buff_alloc_error"`
    Cancel int `json:"cancel"`
    Cant_allocate_control_frame int `json:"cant_allocate_control_frame"`
    Cant_allocate_goaway_frame int `json:"cant_allocate_goaway_frame"`
    Cant_allocate_ping_frame int `json:"cant_allocate_ping_frame"`
    Cant_allocate_rst_frame int `json:"cant_allocate_rst_frame"`
    Cant_allocate_settings_frame int `json:"cant_allocate_settings_frame"`
    Cant_allocate_stream int `json:"cant_allocate_stream"`
    Cant_allocate_window_frame int `json:"cant_allocate_window_frame"`
    Closed_state_unexpected_frame int `json:"closed_state_unexpected_frame"`
    Compression_error int `json:"compression_error"`
    Connect_error int `json:"connect_error"`
    Continuation_before_headers int `json:"continuation_before_headers"`
    Data_no_stream int `json:"data_no_stream"`
    Data_queue_alloc_error int `json:"data_queue_alloc_error"`
    Deflate_alloc_fail int `json:"deflate_alloc_fail"`
    Duration int `json:"duration" dval:"60"`
    Enhance_your_calm int `json:"enhance_your_calm"`
    Err_rcvd_total int `json:"err_rcvd_total"`
    Err_sent_cancel int `json:"err_sent_cancel"`
    Err_sent_compression_err int `json:"err_sent_compression_err"`
    Err_sent_connect_err int `json:"err_sent_connect_err"`
    Err_sent_flow_control int `json:"err_sent_flow_control"`
    Err_sent_frame_size_err int `json:"err_sent_frame_size_err"`
    Err_sent_http11_required int `json:"err_sent_http11_required"`
    Err_sent_inadequate_security int `json:"err_sent_inadequate_security"`
    Err_sent_internal_err int `json:"err_sent_internal_err"`
    Err_sent_proto_err int `json:"err_sent_proto_err"`
    Err_sent_refused_stream int `json:"err_sent_refused_stream"`
    Err_sent_setting_timeout int `json:"err_sent_setting_timeout"`
    Err_sent_stream_closed int `json:"err_sent_stream_closed"`
    Err_sent_total int `json:"err_sent_total"`
    Err_sent_your_calm int `json:"err_sent_your_calm"`
    Error_max_invalid_stream int `json:"error_max_invalid_stream"`
    Exceeds_max_window_size_stream int `json:"exceeds_max_window_size_stream"`
    Flow_control_error int `json:"flow_control_error"`
    Frame_size_error int `json:"frame_size_error"`
    Half_closed_remote_state_unexpected_fra int `json:"half_closed_remote_state_unexpected_fra"`
    Header_no_stream int `json:"header_no_stream"`
    Header_padlen_gt_frame_payload int `json:"header_padlen_gt_frame_payload"`
    Headers_after_continuation int `json:"headers_after_continuation"`
    Headers_interleaved int `json:"headers_interleaved"`
    Http_1_1_required int `json:"http_1_1_required"`
    Idle_state_unexpected_frame int `json:"idle_state_unexpected_frame"`
    Inadequate_security int `json:"inadequate_security"`
    Inflate_alloc_fail int `json:"inflate_alloc_fail"`
    Inflate_header_fail int `json:"inflate_header_fail"`
    Internal_error int `json:"internal_error"`
    Invalid_frame_during_headers int `json:"invalid_frame_during_headers"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Invalid_push_promise int `json:"invalid_push_promise"`
    Invalid_setting_value int `json:"invalid_setting_value"`
    Invalid_stream_id int `json:"invalid_stream_id"`
    Invalid_window_update int `json:"invalid_window_update"`
    Protocol_error int `json:"protocol_error"`
    Proxy_alloc_error int `json:"proxy_alloc_error"`
    Refused_stream int `json:"refused_stream"`
    Reserved_local_state_unexpected_frame int `json:"reserved_local_state_unexpected_frame"`
    Reserved_remote_state_unexpected_frame int `json:"reserved_remote_state_unexpected_frame"`
    Settings_timeout int `json:"settings_timeout"`
    Split_buff_fail int `json:"split_buff_fail"`
    Streams_gt_max_concur_streams int `json:"streams_gt_max_concur_streams"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Trailers_no_end_stream int `json:"trailers_no_end_stream"`
    Uuid string `json:"uuid"`
    Wrong_stream_state int `json:"wrong_stream_state"`
    Zero_window_size_on_stream int `json:"zero_window_size_on_stream"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-http2/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
