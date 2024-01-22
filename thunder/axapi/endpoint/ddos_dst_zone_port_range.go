

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRange struct {
	Inst struct {

    Age int `json:"age" dval:"5"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    DefaultActionList string `json:"default-action-list"`
    Deny int `json:"deny"`
    DynamicEntryOverflowPolicyList []DdosDstZonePortRangeDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
    EnableClassListOverflow int `json:"enable-class-list-overflow"`
    EnableTopK int `json:"enable-top-k"`
    EnableTopKDestination int `json:"enable-top-k-destination"`
    FasterDeEscalation int `json:"faster-de-escalation"`
    GlidCfg DdosDstZonePortRangeGlidCfg `json:"glid-cfg"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    IpFilteringPolicyOper DdosDstZonePortRangeIpFilteringPolicyOper212 `json:"ip-filtering-policy-oper"`
    Ips DdosDstZonePortRangeIps213 `json:"ips"`
    LevelList []DdosDstZonePortRangeLevelList `json:"level-list"`
    ManualModeEnable int `json:"manual-mode-enable"`
    ManualModeList []DdosDstZonePortRangeManualModeList `json:"manual-mode-list"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    OutboundOnly int `json:"outbound-only"`
    PatternRecognition DdosDstZonePortRangePatternRecognition215 `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstZonePortRangePatternRecognitionPuDetails216 `json:"pattern-recognition-pu-details"`
    PortInd DdosDstZonePortRangePortInd217 `json:"port-ind"`
    PortRangeEnd int `json:"port-range-end"`
    PortRangeStart int `json:"port-range-start"`
    ProgressionTracking DdosDstZonePortRangeProgressionTracking219 `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    SflowCommon int `json:"sflow-common"`
    SflowHttp int `json:"sflow-http"`
    SflowPackets int `json:"sflow-packets"`
    SflowTcp DdosDstZonePortRangeSflowTcp `json:"sflow-tcp"`
    SrcBasedPolicyList []DdosDstZonePortRangeSrcBasedPolicyList `json:"src-based-policy-list"`
    Stateful int `json:"stateful"`
    TopkDestinations DdosDstZonePortRangeTopkDestinations220 `json:"topk-destinations"`
    TopkDstNumRecords int `json:"topk-dst-num-records" dval:"20"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    TopkSources DdosDstZonePortRangeTopkSources221 `json:"topk-sources"`
    UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"port-range"`
}


type DdosDstZonePortRangeDynamicEntryOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZonePortRangeDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZonePortRangeDynamicEntryOverflowPolicyListZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}


type DdosDstZonePortRangeGlidCfg struct {
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    ActionList string `json:"action-list"`
    PerAddrGlid string `json:"per-addr-glid"`
}


type DdosDstZonePortRangeIpFilteringPolicyOper212 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZonePortRangeIps213 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstZonePortRangeIpsSamplingEnable214 `json:"sampling-enable"`
}


type DdosDstZonePortRangeIpsSamplingEnable214 struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZonePortRangeLevelList struct {
    LevelNum string `json:"level-num"`
    SrcDefaultGlid string `json:"src-default-glid"`
    GlidAction string `json:"glid-action"`
    ZoneEscalationScore int `json:"zone-escalation-score"`
    ZoneViolationActions string `json:"zone-violation-actions"`
    SrcEscalationScore int `json:"src-escalation-score"`
    SrcViolationActions string `json:"src-violation-actions"`
    ZoneTemplate DdosDstZonePortRangeLevelListZoneTemplate `json:"zone-template"`
    CloseSessionsForUnauthSources int `json:"close-sessions-for-unauth-sources"`
    StartPatternRecognition int `json:"start-pattern-recognition"`
    ApplyExtractedFilters int `json:"apply-extracted-filters"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IndicatorList []DdosDstZonePortRangeLevelListIndicatorList `json:"indicator-list"`
}


type DdosDstZonePortRangeLevelListZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
}


type DdosDstZonePortRangeLevelListIndicatorList struct {
    Type string `json:"type"`
    TcpWindowSize int `json:"tcp-window-size"`
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


type DdosDstZonePortRangeManualModeList struct {
    Config string `json:"config"`
    SrcDefaultGlid string `json:"src-default-glid"`
    GlidAction string `json:"glid-action"`
    ZoneTemplate DdosDstZonePortRangeManualModeListZoneTemplate `json:"zone-template"`
    CloseSessionsForUnauthSources int `json:"close-sessions-for-unauth-sources"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZonePortRangeManualModeListZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
}


type DdosDstZonePortRangePatternRecognition215 struct {
    Algorithm string `json:"algorithm"`
    Mode string `json:"mode"`
    Sensitivity string `json:"sensitivity"`
    FilterThreshold int `json:"filter-threshold"`
    FilterInactiveThreshold int `json:"filter-inactive-threshold"`
    TriggeredBy string `json:"triggered-by"`
    CaptureTraffic string `json:"capture-traffic"`
    AppPayloadOffset int `json:"app-payload-offset"`
    Uuid string `json:"uuid"`
}


type DdosDstZonePortRangePatternRecognitionPuDetails216 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZonePortRangePortInd217 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstZonePortRangePortIndSamplingEnable218 `json:"sampling-enable"`
}


type DdosDstZonePortRangePortIndSamplingEnable218 struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZonePortRangeProgressionTracking219 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZonePortRangeSflowTcp struct {
    SflowTcpBasic int `json:"sflow-tcp-basic"`
    SflowTcpStateful int `json:"sflow-tcp-stateful"`
}


type DdosDstZonePortRangeSrcBasedPolicyList struct {
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PolicyClassListList []DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}


type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Glid string `json:"glid"`
    GlidAction string `json:"glid-action"`
    Action string `json:"action"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ZoneTemplate DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListSamplingEnable `json:"sampling-enable"`
    ClassListOverflowPolicyList []DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}


type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    Action string `json:"action"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    ZoneTemplate DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
    Quic string `json:"quic"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}


type DdosDstZonePortRangeTopkDestinations220 struct {
    Uuid string `json:"uuid"`
}


type DdosDstZonePortRangeTopkSources221 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDstZonePortRange) GetId() string{
    return strconv.Itoa(p.Inst.PortRangeStart)+"+"+strconv.Itoa(p.Inst.PortRangeEnd)+"+"+p.Inst.Protocol
}

func (p *DdosDstZonePortRange) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port-range"
}

func (p *DdosDstZonePortRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRange::Post")
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

func (p *DdosDstZonePortRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRange::Get")
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
func (p *DdosDstZonePortRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRange::Put")
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

func (p *DdosDstZonePortRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortRange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
