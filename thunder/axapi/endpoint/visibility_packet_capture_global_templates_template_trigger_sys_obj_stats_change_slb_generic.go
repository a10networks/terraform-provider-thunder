

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2033 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2034 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"slb-generic"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2033 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Client_fail int `json:"client_fail"`
    Server_fail int `json:"server_fail"`
    Mismatch_fwd_id int `json:"mismatch_fwd_id"`
    Mismatch_rev_id int `json:"mismatch_rev_id"`
    Unkwn_cmd_code int `json:"unkwn_cmd_code"`
    No_session_id int `json:"no_session_id"`
    No_fwd_tuple int `json:"no_fwd_tuple"`
    No_rev_tuple int `json:"no_rev_tuple"`
    Dcmsg_error int `json:"dcmsg_error"`
    Retry_client_request_fail int `json:"retry_client_request_fail"`
    Reply_unknown_session_id int `json:"reply_unknown_session_id"`
    Client_select_fail int `json:"client_select_fail"`
    Invalid_avp int `json:"invalid_avp"`
    Reply_error_info_fail int `json:"reply_error_info_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2034 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Client_fail int `json:"client_fail"`
    Server_fail int `json:"server_fail"`
    Mismatch_fwd_id int `json:"mismatch_fwd_id"`
    Mismatch_rev_id int `json:"mismatch_rev_id"`
    Unkwn_cmd_code int `json:"unkwn_cmd_code"`
    No_session_id int `json:"no_session_id"`
    No_fwd_tuple int `json:"no_fwd_tuple"`
    No_rev_tuple int `json:"no_rev_tuple"`
    Dcmsg_error int `json:"dcmsg_error"`
    Retry_client_request_fail int `json:"retry_client_request_fail"`
    Reply_unknown_session_id int `json:"reply_unknown_session_id"`
    Client_select_fail int `json:"client_select_fail"`
    Invalid_avp int `json:"invalid_avp"`
    Reply_error_info_fail int `json:"reply_error_info_fail"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-generic"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
