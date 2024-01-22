

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2043 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2044 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"slb-imap-proxy"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2043 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Auth_unsupported int `json:"auth_unsupported"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2044 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Auth_unsupported int `json:"auth_unsupported"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-imap-proxy"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
