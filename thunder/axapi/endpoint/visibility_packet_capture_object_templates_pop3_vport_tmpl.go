

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesPop3VportTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesPop3VportTmplTriggerStatsInc2702 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesPop3VportTmplTriggerStatsRate2703 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesPop3VportTmplTriggerStatsSeverity2704 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"pop3-vport-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplTriggerStatsInc2702 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplTriggerStatsRate2703 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplTriggerStatsSeverity2704 struct {
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

func (p *VisibilityPacketCaptureObjectTemplatesPop3VportTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesPop3VportTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/pop3-vport-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesPop3VportTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesPop3VportTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesPop3VportTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesPop3VportTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesPop3VportTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesPop3VportTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesPop3VportTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesPop3VportTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
