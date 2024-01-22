

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc1969 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate1970 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"cgnv6-fixed-nat-global"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc1969 struct {
    NatPortUnavailableTcp int `json:"nat-port-unavailable-tcp"`
    NatPortUnavailableUdp int `json:"nat-port-unavailable-udp"`
    NatPortUnavailableIcmp int `json:"nat-port-unavailable-icmp"`
    SessionUserQuotaExceeded int `json:"session-user-quota-exceeded"`
    FullconeFailure int `json:"fullcone-failure"`
    Nat44InboundFiltered int `json:"nat44-inbound-filtered"`
    Nat64InboundFiltered int `json:"nat64-inbound-filtered"`
    DsliteInboundFiltered int `json:"dslite-inbound-filtered"`
    Nat44EifLimitExceeded int `json:"nat44-eif-limit-exceeded"`
    Nat64EifLimitExceeded int `json:"nat64-eif-limit-exceeded"`
    DsliteEifLimitExceeded int `json:"dslite-eif-limit-exceeded"`
    StandbyDrop int `json:"standby-drop"`
    FixedNatFullconeSelfHairpinningDro int `json:"fixed-nat-fullcone-self-hairpinning-dro"`
    SixrdDrop int `json:"sixrd-drop"`
    DestRlistDrop int `json:"dest-rlist-drop"`
    DestRlistPassThrough int `json:"dest-rlist-pass-through"`
    DestRlistSnatDrop int `json:"dest-rlist-snat-drop"`
    ConfigNotFound int `json:"config-not-found"`
    PortOverloadFailed int `json:"port-overload-failed"`
    HaSessionUserQuotaExceeded int `json:"ha-session-user-quota-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate1970 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NatPortUnavailableTcp int `json:"nat-port-unavailable-tcp"`
    NatPortUnavailableUdp int `json:"nat-port-unavailable-udp"`
    NatPortUnavailableIcmp int `json:"nat-port-unavailable-icmp"`
    SessionUserQuotaExceeded int `json:"session-user-quota-exceeded"`
    FullconeFailure int `json:"fullcone-failure"`
    Nat44InboundFiltered int `json:"nat44-inbound-filtered"`
    Nat64InboundFiltered int `json:"nat64-inbound-filtered"`
    DsliteInboundFiltered int `json:"dslite-inbound-filtered"`
    Nat44EifLimitExceeded int `json:"nat44-eif-limit-exceeded"`
    Nat64EifLimitExceeded int `json:"nat64-eif-limit-exceeded"`
    DsliteEifLimitExceeded int `json:"dslite-eif-limit-exceeded"`
    StandbyDrop int `json:"standby-drop"`
    FixedNatFullconeSelfHairpinningDro int `json:"fixed-nat-fullcone-self-hairpinning-dro"`
    SixrdDrop int `json:"sixrd-drop"`
    DestRlistDrop int `json:"dest-rlist-drop"`
    DestRlistPassThrough int `json:"dest-rlist-pass-through"`
    DestRlistSnatDrop int `json:"dest-rlist-snat-drop"`
    ConfigNotFound int `json:"config-not-found"`
    PortOverloadFailed int `json:"port-overload-failed"`
    HaSessionUserQuotaExceeded int `json:"ha-session-user-quota-exceeded"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/cgnv6-fixed-nat-global"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
