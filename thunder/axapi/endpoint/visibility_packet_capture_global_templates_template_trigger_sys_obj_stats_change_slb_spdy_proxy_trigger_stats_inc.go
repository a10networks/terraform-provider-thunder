

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc struct {
	Inst struct {

    Compress_ctx_alloc_fail int `json:"compress_ctx_alloc_fail"`
    Ctx_alloc_fail int `json:"ctx_alloc_fail"`
    Data_cb_no_tuple int `json:"data_cb_no_tuple"`
    Data_no_stream int `json:"data_no_stream"`
    Data_no_stream_goaway_close int `json:"data_no_stream_goaway_close"`
    Data_no_stream_no_goaway int `json:"data_no_stream_no_goaway"`
    Decompress_fail int `json:"decompress_fail"`
    Est_cb_no_tuple int `json:"est_cb_no_tuple"`
    Header_compress_fail int `json:"header_compress_fail"`
    Http_conn_alloc_fail int `json:"http_conn_alloc_fail"`
    Http_err_stream_closed int `json:"http_err_stream_closed"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Invalid_version int `json:"invalid_version"`
    Request_header_alloc_fail int `json:"request_header_alloc_fail"`
    Session_err int `json:"session_err"`
    Stream_alloc_fail int `json:"stream_alloc_fail"`
    Stream_err int `json:"stream_err"`
    Stream_not_found int `json:"stream_not_found"`
    Tcp_err int `json:"tcp_err"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-spdy-proxy/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
