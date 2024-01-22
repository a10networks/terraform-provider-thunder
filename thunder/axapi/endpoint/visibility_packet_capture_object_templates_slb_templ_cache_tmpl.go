

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc2711 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate2712 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsSeverity2713 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"slb-templ-cache-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc2711 struct {
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate2712 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsSeverity2713 struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/slb-templ-cache-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
