

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumber struct {
	Inst struct {

    Age int `json:"age" dval:"5"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    Deny int `json:"deny"`
    DropFragPkt int `json:"drop-frag-pkt"`
    DynamicEntryOverflowPolicyList []DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
    EnableClassListOverflow int `json:"enable-class-list-overflow"`
    EnableTopK int `json:"enable-top-k"`
    EnableTopKDestination int `json:"enable-top-k-destination"`
    EspInspect DdosDstZoneIpProtoProtoNumberEspInspect `json:"esp-inspect"`
    FasterDeEscalation int `json:"faster-de-escalation"`
    GlidCfg DdosDstZoneIpProtoProtoNumberGlidCfg `json:"glid-cfg"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    IpFilteringPolicyOper DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOper205 `json:"ip-filtering-policy-oper"`
    LevelList []DdosDstZoneIpProtoProtoNumberLevelList `json:"level-list"`
    ManualModeEnable int `json:"manual-mode-enable"`
    ManualModeList []DdosDstZoneIpProtoProtoNumberManualModeList `json:"manual-mode-list"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    PortInd DdosDstZoneIpProtoProtoNumberPortInd206 `json:"port-ind"`
    ProgressionTracking DdosDstZoneIpProtoProtoNumberProgressionTracking208 `json:"progression-tracking"`
    ProtocolNum int `json:"protocol-num"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    SrcBasedPolicyList []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyList `json:"src-based-policy-list"`
    TopkDestinations DdosDstZoneIpProtoProtoNumberTopkDestinations209 `json:"topk-destinations"`
    TopkDstNumRecords int `json:"topk-dst-num-records" dval:"20"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    TopkSources DdosDstZoneIpProtoProtoNumberTopkSources210 `json:"topk-sources"`
    UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"proto-number"`
}


type DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNumberDynamicEntryOverflowPolicyListZoneTemplate struct {
    IpProto string `json:"ip-proto"`
}


type DdosDstZoneIpProtoProtoNumberEspInspect struct {
    AuthAlgorithm string `json:"auth-algorithm"`
    EncryptAlgorithm string `json:"encrypt-algorithm"`
    Mode string `json:"mode"`
}


type DdosDstZoneIpProtoProtoNumberGlidCfg struct {
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    ActionList string `json:"action-list"`
    PerAddrGlid string `json:"per-addr-glid"`
}


type DdosDstZoneIpProtoProtoNumberIpFilteringPolicyOper205 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneIpProtoProtoNumberLevelList struct {
    LevelNum string `json:"level-num"`
    SrcDefaultGlid string `json:"src-default-glid"`
    GlidAction string `json:"glid-action"`
    ZoneEscalationScore int `json:"zone-escalation-score"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    SrcEscalationScore int `json:"src-escalation-score"`
    SrcViolationActions string `json:"src-violation-actions"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberLevelListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosDstZoneIpProtoProtoNumberLevelListIndicatorList `json:"indicator-list"`
}


type DdosDstZoneIpProtoProtoNumberLevelListZoneTemplate struct {
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNumberLevelListIndicatorList struct {
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


type DdosDstZoneIpProtoProtoNumberManualModeList struct {
    Config string `json:"config"`
    SrcDefaultGlid string `json:"src-default-glid"`
    GlidAction string `json:"glid-action"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberManualModeListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNumberManualModeListZoneTemplate struct {
    IpProto string `json:"ip-proto"`
    Encap string `json:"encap"`
}


type DdosDstZoneIpProtoProtoNumberPortInd206 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstZoneIpProtoProtoNumberPortIndSamplingEnable207 `json:"sampling-enable"`
}


type DdosDstZoneIpProtoProtoNumberPortIndSamplingEnable207 struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZoneIpProtoProtoNumberProgressionTracking208 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyList struct {
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PolicyClassListList []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListSamplingEnable `json:"sampling-enable"`
    ClassListOverflowPolicyList []DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
    Logging string `json:"logging"`
    IpProto string `json:"ip-proto"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
    IpProto string `json:"ip-proto"`
}


type DdosDstZoneIpProtoProtoNumberTopkDestinations209 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZoneIpProtoProtoNumberTopkSources210 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDstZoneIpProtoProtoNumber) GetId() string{
    return strconv.Itoa(p.Inst.ProtocolNum)
}

func (p *DdosDstZoneIpProtoProtoNumber) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/ip-proto/proto-number"
}

func (p *DdosDstZoneIpProtoProtoNumber) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumber::Post")
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

func (p *DdosDstZoneIpProtoProtoNumber) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumber::Get")
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
func (p *DdosDstZoneIpProtoProtoNumber) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumber::Put")
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

func (p *DdosDstZoneIpProtoProtoNumber) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneIpProtoProtoNumber::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
