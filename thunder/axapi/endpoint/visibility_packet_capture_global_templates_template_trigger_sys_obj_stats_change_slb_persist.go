

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2059 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2060 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"slb-persist"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2059 struct {
    Hash_tbl_trylock_fail int `json:"hash_tbl_trylock_fail"`
    Hash_tbl_create_fail int `json:"hash_tbl_create_fail"`
    Hash_tbl_rst_updown int `json:"hash_tbl_rst_updown"`
    Hash_tbl_rst_adddel int `json:"hash_tbl_rst_adddel"`
    Url_hash_fail int `json:"url_hash_fail"`
    Header_hash_fail int `json:"header_hash_fail"`
    Src_ip_fail int `json:"src_ip_fail"`
    Src_ip_new_sess_cache_fail int `json:"src_ip_new_sess_cache_fail"`
    Src_ip_new_sess_sel_fail int `json:"src_ip_new_sess_sel_fail"`
    Src_ip_hash_fail int `json:"src_ip_hash_fail"`
    Dst_ip_fail int `json:"dst_ip_fail"`
    Dst_ip_new_sess_cache_fail int `json:"dst_ip_new_sess_cache_fail"`
    Dst_ip_new_sess_sel_fail int `json:"dst_ip_new_sess_sel_fail"`
    Dst_ip_hash_fail int `json:"dst_ip_hash_fail"`
    Cssl_sid_not_found int `json:"cssl_sid_not_found"`
    Cssl_sid_not_match int `json:"cssl_sid_not_match"`
    Sssl_sid_not_found int `json:"sssl_sid_not_found"`
    Sssl_sid_not_match int `json:"sssl_sid_not_match"`
    Ssl_sid_persist_fail int `json:"ssl_sid_persist_fail"`
    Ssl_sid_session_fail int `json:"ssl_sid_session_fail"`
    Cookie_persist_fail int `json:"cookie_persist_fail"`
    Cookie_not_found int `json:"cookie_not_found"`
    Cookie_invalid int `json:"cookie_invalid"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2060 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Hash_tbl_trylock_fail int `json:"hash_tbl_trylock_fail"`
    Hash_tbl_create_fail int `json:"hash_tbl_create_fail"`
    Hash_tbl_rst_updown int `json:"hash_tbl_rst_updown"`
    Hash_tbl_rst_adddel int `json:"hash_tbl_rst_adddel"`
    Url_hash_fail int `json:"url_hash_fail"`
    Header_hash_fail int `json:"header_hash_fail"`
    Src_ip_fail int `json:"src_ip_fail"`
    Src_ip_new_sess_cache_fail int `json:"src_ip_new_sess_cache_fail"`
    Src_ip_new_sess_sel_fail int `json:"src_ip_new_sess_sel_fail"`
    Src_ip_hash_fail int `json:"src_ip_hash_fail"`
    Dst_ip_fail int `json:"dst_ip_fail"`
    Dst_ip_new_sess_cache_fail int `json:"dst_ip_new_sess_cache_fail"`
    Dst_ip_new_sess_sel_fail int `json:"dst_ip_new_sess_sel_fail"`
    Dst_ip_hash_fail int `json:"dst_ip_hash_fail"`
    Cssl_sid_not_found int `json:"cssl_sid_not_found"`
    Cssl_sid_not_match int `json:"cssl_sid_not_match"`
    Sssl_sid_not_found int `json:"sssl_sid_not_found"`
    Sssl_sid_not_match int `json:"sssl_sid_not_match"`
    Ssl_sid_persist_fail int `json:"ssl_sid_persist_fail"`
    Ssl_sid_session_fail int `json:"ssl_sid_session_fail"`
    Cookie_persist_fail int `json:"cookie_persist_fail"`
    Cookie_not_found int `json:"cookie_not_found"`
    Cookie_invalid int `json:"cookie_invalid"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-persist"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
