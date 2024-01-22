

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc1997 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate1998 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"cgnv6-nat64-global"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc1997 struct {
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    New_user_resource_unavailable int `json:"new_user_resource_unavailable"`
    Fullcone_failure int `json:"fullcone_failure"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    Eif_limit_exceeded int `json:"eif_limit_exceeded"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    No_radius_profile_match int `json:"no_radius_profile_match"`
    No_class_list_match int `json:"no_class_list_match"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    User_quota_unusable int `json:"user_quota_unusable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate1998 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    New_user_resource_unavailable int `json:"new_user_resource_unavailable"`
    Fullcone_failure int `json:"fullcone_failure"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    Eif_limit_exceeded int `json:"eif_limit_exceeded"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    No_radius_profile_match int `json:"no_radius_profile_match"`
    No_class_list_match int `json:"no_class_list_match"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    User_quota_unusable int `json:"user_quota_unusable"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/cgnv6-nat64-global"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
