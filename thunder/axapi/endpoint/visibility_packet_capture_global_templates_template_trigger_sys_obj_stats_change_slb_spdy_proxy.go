

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2075 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2076 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"slb-spdy-proxy"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2075 struct {
    Tcp_err int `json:"tcp_err"`
    Stream_not_found int `json:"stream_not_found"`
    Stream_err int `json:"stream_err"`
    Session_err int `json:"session_err"`
    Data_no_stream int `json:"data_no_stream"`
    Data_no_stream_no_goaway int `json:"data_no_stream_no_goaway"`
    Data_no_stream_goaway_close int `json:"data_no_stream_goaway_close"`
    Est_cb_no_tuple int `json:"est_cb_no_tuple"`
    Data_cb_no_tuple int `json:"data_cb_no_tuple"`
    Ctx_alloc_fail int `json:"ctx_alloc_fail"`
    Stream_alloc_fail int `json:"stream_alloc_fail"`
    Http_conn_alloc_fail int `json:"http_conn_alloc_fail"`
    Request_header_alloc_fail int `json:"request_header_alloc_fail"`
    Decompress_fail int `json:"decompress_fail"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Invalid_version int `json:"invalid_version"`
    Compress_ctx_alloc_fail int `json:"compress_ctx_alloc_fail"`
    Header_compress_fail int `json:"header_compress_fail"`
    Http_err_stream_closed int `json:"http_err_stream_closed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2076 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Tcp_err int `json:"tcp_err"`
    Stream_not_found int `json:"stream_not_found"`
    Stream_err int `json:"stream_err"`
    Session_err int `json:"session_err"`
    Data_no_stream int `json:"data_no_stream"`
    Data_no_stream_no_goaway int `json:"data_no_stream_no_goaway"`
    Data_no_stream_goaway_close int `json:"data_no_stream_goaway_close"`
    Est_cb_no_tuple int `json:"est_cb_no_tuple"`
    Data_cb_no_tuple int `json:"data_cb_no_tuple"`
    Ctx_alloc_fail int `json:"ctx_alloc_fail"`
    Stream_alloc_fail int `json:"stream_alloc_fail"`
    Http_conn_alloc_fail int `json:"http_conn_alloc_fail"`
    Request_header_alloc_fail int `json:"request_header_alloc_fail"`
    Decompress_fail int `json:"decompress_fail"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Invalid_version int `json:"invalid_version"`
    Compress_ctx_alloc_fail int `json:"compress_ctx_alloc_fail"`
    Header_compress_fail int `json:"header_compress_fail"`
    Http_err_stream_closed int `json:"http_err_stream_closed"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-spdy-proxy"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
