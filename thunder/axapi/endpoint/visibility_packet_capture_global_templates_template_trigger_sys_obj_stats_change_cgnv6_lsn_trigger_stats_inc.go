

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc struct {
	Inst struct {

    Adc_port_allocation_failed int `json:"adc_port_allocation_failed"`
    Data_sesn_user_quota_exceeded int `json:"data_sesn_user_quota_exceeded"`
    Fullcone_ext_mem_alloc_failure int `json:"fullcone_ext_mem_alloc_failure"`
    Fullcone_ext_mem_alloc_init_faulure int `json:"fullcone_ext_mem_alloc_init_faulure"`
    Fullcone_failure int `json:"fullcone_failure"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    H323_alg_alloc_single_port_failure int `json:"h323_alg_alloc_single_port_failure"`
    H323_alg_create_rtcp_fullcone_failure int `json:"h323_alg_create_rtcp_fullcone_failure"`
    H323_alg_create_rtp_fullcone_failure int `json:"h323_alg_create_rtp_fullcone_failure"`
    H323_alg_create_single_fullcone_failure int `json:"h323_alg_create_single_fullcone_failure"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Mgcp_alg_create_rtcp_fullcone_failure int `json:"mgcp_alg_create_rtcp_fullcone_failure"`
    Mgcp_alg_create_rtp_fullcone_failure int `json:"mgcp_alg_create_rtp_fullcone_failure"`
    Mgcp_alg_port_pair_alloc_from_quota_par int `json:"mgcp_alg_port_pair_alloc_from_quota_par"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Port_overloading_inc_overflow int `json:"port_overloading_inc_overflow"`
    Port_overloading_out_of_memory int `json:"port_overloading_out_of_memory"`
    Sip_alg_alloc_rtp_rtcp_port_failure int `json:"sip_alg_alloc_rtp_rtcp_port_failure"`
    Sip_alg_alloc_single_port_failure int `json:"sip_alg_alloc_single_port_failure"`
    Sip_alg_create_rtcp_fullcone_failure int `json:"sip_alg_create_rtcp_fullcone_failure"`
    Sip_alg_create_rtp_fullcone_failure int `json:"sip_alg_create_rtp_fullcone_failure"`
    Sip_alg_create_single_fullcone_failure int `json:"sip_alg_create_single_fullcone_failure"`
    Sip_alg_quota_inc_failure int `json:"sip_alg_quota_inc_failure"`
    User_quota_failure int `json:"user_quota_failure"`
    User_quota_unusable int `json:"user_quota_unusable"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/cgnv6-lsn/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
