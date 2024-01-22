

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoName struct {
	Inst struct {

    Age int `json:"age" dval:"5"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    Deny int `json:"deny"`
    DropFragPkt int `json:"drop-frag-pkt"`
    DynamicEntryOverflowPolicyList []DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
    EnableClassListOverflow int `json:"enable-class-list-overflow"`
    EnableTopK int `json:"enable-top-k"`
    EnableTopKDestination int `json:"enable-top-k-destination"`
    FasterDeEscalation int `json:"faster-de-escalation"`
    GlidCfg DdosDstZoneIpProtoProtoNameGlidCfg `json:"glid-cfg"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    IpFilteringPolicyOper DdosDstZoneIpProtoProtoNameIpFilteringPolicyOper199 `json:"ip-filtering-policy-oper"`
    KeyCfg []DdosDstZoneIpProtoProtoNameKeyCfg `json:"key-cfg"`
    LevelList []DdosDstZoneIpProtoProtoNameLevelList `json:"level-list"`
    ManualModeEnable int `json:"manual-mode-enable"`
    ManualModeList []DdosDstZoneIpProtoProtoNameManualModeList `json:"manual-mode-list"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    PortInd DdosDstZoneIpProtoProtoNamePortInd200 `json:"port-ind"`
    ProgressionTracking DdosDstZoneIpProtoProtoNameProgressionTracking202 `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    SrcBasedPolicyList []DdosDstZoneIpProtoProtoNameSrcBasedPolicyList `json:"src-based-policy-list"`
    TopkDestinations DdosDstZoneIpProtoProtoNameTopkDestinations203 `json:"topk-destinations"`
    TopkDstNumRecords int `json:"topk-dst-num-records" dval:"20"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    TopkSources DdosDstZoneIpProtoProtoNameTopkSources204 `json:"topk-sources"`
    TunnelDecap int `json:"tunnel-decap"`
    TunnelRateLimit int `json:"tunnel-rate-limit"`
    UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"proto-name"`
}


type DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNameDynamicEntryOverflowPolicyListZoneTemplate struct {
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNameGlidCfg struct {
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    ActionList string `json:"action-list"`
    PerAddrGlid string `json:"per-addr-glid"`
}


type DdosDstZoneIpProtoProtoNameIpFilteringPolicyOper199 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneIpProtoProtoNameKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstZoneIpProtoProtoNameLevelList struct {
    LevelNum string `json:"level-num"`
    SrcDefaultGlid string `json:"src-default-glid"`
    GlidAction string `json:"glid-action"`
    ZoneEscalationScore int `json:"zone-escalation-score"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    SrcEscalationScore int `json:"src-escalation-score"`
    SrcViolationActions string `json:"src-violation-actions"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameLevelListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosDstZoneIpProtoProtoNameLevelListIndicatorList `json:"indicator-list"`
}


type DdosDstZoneIpProtoProtoNameLevelListZoneTemplate struct {
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNameLevelListIndicatorList struct {
    Type string `json:"type"`
    DataPacketSize int `json:"data-packet-size"`
    Score int `json:"score"`
    SrcThresholdNum int `json:"src-threshold-num"`
    SrcThresholdLargeNum int `json:"src-threshold-large-num"`
    SrcThresholdStr string `json:"src-threshold-str"`
    SrcViolationActions string `json:"src-violation-actions"`
    ZoneThresholdNum int `json:"zone-threshold-num"`
    ZoneThresholdLargeNum int `json:"zone-threshold-large-num"`
    ZoneThresholdStr string `json:"zone-threshold-str"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNameManualModeList struct {
    Config string `json:"config"`
    SrcDefaultGlid string `json:"src-default-glid"`
    GlidAction string `json:"glid-action"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameManualModeListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNameManualModeListZoneTemplate struct {
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNamePortInd200 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstZoneIpProtoProtoNamePortIndSamplingEnable201 `json:"sampling-enable"`
}


type DdosDstZoneIpProtoProtoNamePortIndSamplingEnable201 struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZoneIpProtoProtoNameProgressionTracking202 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyList struct {
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PolicyClassListList []DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListSamplingEnable `json:"sampling-enable"`
    ClassListOverflowPolicyList []DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
    Logging string `json:"logging"`
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
    IcmpV4 string `json:"icmp-v4"`
    IcmpV6 string `json:"icmp-v6"`
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNameTopkDestinations203 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneIpProtoProtoNameTopkSources204 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDstZoneIpProtoProtoName) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosDstZoneIpProtoProtoName) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-name"
}

func (p *DdosDstZoneIpProtoProtoName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoName::Post")
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

func (p *DdosDstZoneIpProtoProtoName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoName::Get")
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
func (p *DdosDstZoneIpProtoProtoName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoName::Put")
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

func (p *DdosDstZoneIpProtoProtoName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
