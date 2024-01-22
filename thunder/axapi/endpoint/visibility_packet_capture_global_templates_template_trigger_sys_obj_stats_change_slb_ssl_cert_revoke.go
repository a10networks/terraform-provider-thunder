

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2079 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2080 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"slb-ssl-cert-revoke"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2079 struct {
    Ocsp_chain_status_revoked int `json:"ocsp_chain_status_revoked"`
    Ocsp_chain_status_unknown int `json:"ocsp_chain_status_unknown"`
    Ocsp_connection_error int `json:"ocsp_connection_error"`
    Ocsp_uri_not_found int `json:"ocsp_uri_not_found"`
    Ocsp_uri_https int `json:"ocsp_uri_https"`
    Ocsp_uri_unsupported int `json:"ocsp_uri_unsupported"`
    Ocsp_response_status_revoked int `json:"ocsp_response_status_revoked"`
    Ocsp_response_status_unknown int `json:"ocsp_response_status_unknown"`
    Ocsp_cache_status_revoked int `json:"ocsp_cache_status_revoked"`
    Ocsp_cache_miss int `json:"ocsp_cache_miss"`
    Ocsp_other_error int `json:"ocsp_other_error"`
    Ocsp_response_no_nonce int `json:"ocsp_response_no_nonce"`
    Ocsp_response_nonce_error int `json:"ocsp_response_nonce_error"`
    Crl_connection_error int `json:"crl_connection_error"`
    Crl_uri_not_found int `json:"crl_uri_not_found"`
    Crl_uri_https int `json:"crl_uri_https"`
    Crl_uri_unsupported int `json:"crl_uri_unsupported"`
    Crl_response_status_revoked int `json:"crl_response_status_revoked"`
    Crl_response_status_unknown int `json:"crl_response_status_unknown"`
    Crl_cache_status_revoked int `json:"crl_cache_status_revoked"`
    Crl_other_error int `json:"crl_other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2080 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ocsp_chain_status_revoked int `json:"ocsp_chain_status_revoked"`
    Ocsp_chain_status_unknown int `json:"ocsp_chain_status_unknown"`
    Ocsp_connection_error int `json:"ocsp_connection_error"`
    Ocsp_uri_not_found int `json:"ocsp_uri_not_found"`
    Ocsp_uri_https int `json:"ocsp_uri_https"`
    Ocsp_uri_unsupported int `json:"ocsp_uri_unsupported"`
    Ocsp_response_status_revoked int `json:"ocsp_response_status_revoked"`
    Ocsp_response_status_unknown int `json:"ocsp_response_status_unknown"`
    Ocsp_cache_status_revoked int `json:"ocsp_cache_status_revoked"`
    Ocsp_cache_miss int `json:"ocsp_cache_miss"`
    Ocsp_other_error int `json:"ocsp_other_error"`
    Ocsp_response_no_nonce int `json:"ocsp_response_no_nonce"`
    Ocsp_response_nonce_error int `json:"ocsp_response_nonce_error"`
    Crl_connection_error int `json:"crl_connection_error"`
    Crl_uri_not_found int `json:"crl_uri_not_found"`
    Crl_uri_https int `json:"crl_uri_https"`
    Crl_uri_unsupported int `json:"crl_uri_unsupported"`
    Crl_response_status_revoked int `json:"crl_response_status_revoked"`
    Crl_response_status_unknown int `json:"crl_response_status_unknown"`
    Crl_cache_status_revoked int `json:"crl_cache_status_revoked"`
    Crl_other_error int `json:"crl_other_error"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-ssl-cert-revoke"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
