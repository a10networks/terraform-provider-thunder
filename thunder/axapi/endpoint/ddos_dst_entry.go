

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntry struct {
	Inst struct {

    AdvertisedEnable int `json:"advertised-enable"`
    BlackholeOnGlidExceed int `json:"blackhole-on-glid-exceed"`
    CaptureConfigList []DdosDstEntryCaptureConfigList `json:"capture-config-list"`
    Description string `json:"description"`
    DestNatIp string `json:"dest-nat-ip"`
    DestNatIpv6 string `json:"dest-nat-ipv6"`
    DropDisable int `json:"drop-disable"`
    DropDisableFwdImmediate int `json:"drop-disable-fwd-immediate"`
    DropFragPkt int `json:"drop-frag-pkt"`
    DropOnNoSrcDstDefault int `json:"drop-on-no-src-dst-default"`
    DstEntryName string `json:"dst-entry-name"`
    DynamicEntryOverflowPolicyList []DdosDstEntryDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
    EnableTopK []DdosDstEntryEnableTopK `json:"enable-top-k"`
    ExceedLogCfg DdosDstEntryExceedLogCfg `json:"exceed-log-cfg"`
    ExceedLogDepCfg DdosDstEntryExceedLogDepCfg `json:"exceed-log-dep-cfg"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryGlidExceedAction `json:"glid-exceed-action"`
    HwBlacklistBlocking DdosDstEntryHwBlacklistBlocking178 `json:"hw-blacklist-blocking"`
    InboundForwardDscp int `json:"inbound-forward-dscp"`
    IpAddr string `json:"ip-addr"`
    IpProtoList []DdosDstEntryIpProtoList `json:"ip-proto-list"`
    Ipv6Addr string `json:"ipv6-addr"`
    L4TypeList []DdosDstEntryL4TypeList `json:"l4-type-list"`
    LogPeriodic int `json:"log-periodic"`
    OperationalMode string `json:"operational-mode" dval:"protection"`
    OutboundForwardDscp int `json:"outbound-forward-dscp"`
    PatternRecognitionHwFilterEnable int `json:"pattern-recognition-hw-filter-enable"`
    PatternRecognitionSensitivity string `json:"pattern-recognition-sensitivity"`
    PortList []DdosDstEntryPortList `json:"port-list"`
    PortRangeList []DdosDstEntryPortRangeList `json:"port-range-list"`
    ReportingDisabled int `json:"reporting-disabled"`
    SamplingEnable []DdosDstEntrySamplingEnable `json:"sampling-enable"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Sflow DdosDstEntrySflow `json:"sflow"`
    SourceNatPool string `json:"source-nat-pool"`
    SrcDstPair DdosDstEntrySrcDstPair179 `json:"src-dst-pair"`
    SrcDstPairClassListList []DdosDstEntrySrcDstPairClassListList `json:"src-dst-pair-class-list-list"`
    SrcDstPairPolicyList []DdosDstEntrySrcDstPairPolicyList `json:"src-dst-pair-policy-list"`
    SrcDstPairSettingsList []DdosDstEntrySrcDstPairSettingsList `json:"src-dst-pair-settings-list"`
    SrcPortList []DdosDstEntrySrcPortList `json:"src-port-list"`
    SrcPortRangeList []DdosDstEntrySrcPortRangeList `json:"src-port-range-list"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    Template DdosDstEntryTemplate `json:"template"`
    TopkDestinations DdosDstEntryTopkDestinations186 `json:"topk-destinations"`
    TrafficDistributionMode string `json:"traffic-distribution-mode" dval:"default"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"entry"`
}


type DdosDstEntryCaptureConfigList struct {
    Name string `json:"name"`
    Mode string `json:"mode"`
    Uuid string `json:"uuid"`
}


type DdosDstEntryDynamicEntryOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Bypass int `json:"bypass"`
    ExceedLogCfg DdosDstEntryDynamicEntryOverflowPolicyListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntryDynamicEntryOverflowPolicyListTemplate `json:"template"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstList []DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    AppTypeSrcDstList []DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstList `json:"app-type-src-dst-list"`
}


type DdosDstEntryDynamicEntryOverflowPolicyListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntryDynamicEntryOverflowPolicyListTemplate struct {
    Logging string `json:"logging"`
}


type DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntryDynamicEntryOverflowPolicyListL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntryDynamicEntryOverflowPolicyListAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntryEnableTopK struct {
    TopkType string `json:"topk-type"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
}


type DdosDstEntryExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
    LogWithSflow int `json:"log-with-sflow"`
    LogHighFrequency int `json:"log-high-frequency"`
    RateLimit int `json:"rate-limit" dval:"1"`
}


type DdosDstEntryExceedLogDepCfg struct {
    ExceedLogEnable int `json:"exceed-log-enable"`
    LogWithSflowDep int `json:"log-with-sflow-dep"`
}


type DdosDstEntryGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryHwBlacklistBlocking178 struct {
    DstEnable int `json:"dst-enable"`
    SrcEnable int `json:"src-enable"`
    Uuid string `json:"uuid"`
}


type DdosDstEntryIpProtoList struct {
    PortNum int `json:"port-num"`
    Deny int `json:"deny"`
    EspInspect DdosDstEntryIpProtoListEspInspect `json:"esp-inspect"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryIpProtoListGlidExceedAction `json:"glid-exceed-action"`
    Template DdosDstEntryIpProtoListTemplate `json:"template"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IpFilteringPolicyOper DdosDstEntryIpProtoListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
}


type DdosDstEntryIpProtoListEspInspect struct {
    AuthAlgorithm string `json:"auth-algorithm"`
    EncryptAlgorithm string `json:"encrypt-algorithm"`
    Mode string `json:"mode"`
}


type DdosDstEntryIpProtoListGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryIpProtoListGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryIpProtoListGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryIpProtoListTemplate struct {
    Other string `json:"other"`
}


type DdosDstEntryIpProtoListIpFilteringPolicyOper struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryL4TypeList struct {
    Protocol string `json:"protocol"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryL4TypeListGlidExceedAction `json:"glid-exceed-action"`
    Deny int `json:"deny"`
    MaxRexmitSynPerFlow int `json:"max-rexmit-syn-per-flow"`
    MaxRexmitSynPerFlowExceedAction string `json:"max-rexmit-syn-per-flow-exceed-action"`
    DisableSynAuth int `json:"disable-syn-auth"`
    SynAuth string `json:"syn-auth" dval:"send-rst"`
    SynCookie int `json:"syn-cookie"`
    TcpResetClient int `json:"tcp-reset-client"`
    TcpResetServer int `json:"tcp-reset-server"`
    DropOnNoPortMatch string `json:"drop-on-no-port-match" dval:"enable"`
    Stateful int `json:"stateful"`
    TunnelDecap DdosDstEntryL4TypeListTunnelDecap `json:"tunnel-decap"`
    TunnelRateLimit DdosDstEntryL4TypeListTunnelRateLimit `json:"tunnel-rate-limit"`
    DropFragPkt int `json:"drop-frag-pkt"`
    UndefinedPortHitStatistics DdosDstEntryL4TypeListUndefinedPortHitStatistics `json:"undefined-port-hit-statistics"`
    Template DdosDstEntryL4TypeListTemplate `json:"template"`
    DetectionEnable int `json:"detection-enable"`
    EnableTopK int `json:"enable-top-k"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IpFilteringPolicyOper DdosDstEntryL4TypeListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    PortInd DdosDstEntryL4TypeListPortInd `json:"port-ind"`
    TopkSources DdosDstEntryL4TypeListTopkSources `json:"topk-sources"`
    ProgressionTracking DdosDstEntryL4TypeListProgressionTracking `json:"progression-tracking"`
}


type DdosDstEntryL4TypeListGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryL4TypeListGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryL4TypeListGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryL4TypeListTunnelDecap struct {
    IpDecap int `json:"ip-decap"`
    GreDecap int `json:"gre-decap"`
    KeyCfg []DdosDstEntryL4TypeListTunnelDecapKeyCfg `json:"key-cfg"`
}


type DdosDstEntryL4TypeListTunnelDecapKeyCfg struct {
    Key string `json:"key"`
}


type DdosDstEntryL4TypeListTunnelRateLimit struct {
    IpRateLimit int `json:"ip-rate-limit"`
    GreRateLimit int `json:"gre-rate-limit"`
}


type DdosDstEntryL4TypeListUndefinedPortHitStatistics struct {
    UndefinedPortHitStatistics int `json:"undefined-port-hit-statistics"`
    ResetInterval int `json:"reset-interval" dval:"60"`
}


type DdosDstEntryL4TypeListTemplate struct {
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntryL4TypeListIpFilteringPolicyOper struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryL4TypeListPortInd struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstEntryL4TypeListPortIndSamplingEnable `json:"sampling-enable"`
}


type DdosDstEntryL4TypeListPortIndSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntryL4TypeListTopkSources struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryL4TypeListProgressionTracking struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    DetectionEnable int `json:"detection-enable"`
    EnableTopK int `json:"enable-top-k"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryPortListGlidExceedAction `json:"glid-exceed-action"`
    DnsCache string `json:"dns-cache"`
    Template DdosDstEntryPortListTemplate `json:"template"`
    Sflow DdosDstEntryPortListSflow `json:"sflow"`
    CaptureConfig DdosDstEntryPortListCaptureConfig `json:"capture-config"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PortInd DdosDstEntryPortListPortInd `json:"port-ind"`
    IpFilteringPolicyOper DdosDstEntryPortListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    TopkSources DdosDstEntryPortListTopkSources `json:"topk-sources"`
    ProgressionTracking DdosDstEntryPortListProgressionTracking `json:"progression-tracking"`
    SignatureExtraction DdosDstEntryPortListSignatureExtraction `json:"signature-extraction"`
    PatternRecognition DdosDstEntryPortListPatternRecognition `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstEntryPortListPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
}


type DdosDstEntryPortListGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryPortListGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryPortListGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryPortListTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
}


type DdosDstEntryPortListSflow struct {
    Polling DdosDstEntryPortListSflowPolling `json:"polling"`
}


type DdosDstEntryPortListSflowPolling struct {
    SflowPackets int `json:"sflow-packets"`
    SflowTcp DdosDstEntryPortListSflowPollingSflowTcp `json:"sflow-tcp"`
    SflowHttp int `json:"sflow-http"`
}


type DdosDstEntryPortListSflowPollingSflowTcp struct {
    SflowTcpBasic int `json:"sflow-tcp-basic"`
    SflowTcpStateful int `json:"sflow-tcp-stateful"`
}


type DdosDstEntryPortListCaptureConfig struct {
    CaptureConfigName string `json:"capture-config-name"`
    CaptureConfigMode string `json:"capture-config-mode"`
}


type DdosDstEntryPortListPortInd struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstEntryPortListPortIndSamplingEnable `json:"sampling-enable"`
}


type DdosDstEntryPortListPortIndSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntryPortListIpFilteringPolicyOper struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortListTopkSources struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortListProgressionTracking struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortListSignatureExtraction struct {
    Algorithm string `json:"algorithm"`
    ManualMode int `json:"manual-mode"`
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortListPatternRecognition struct {
    Algorithm string `json:"algorithm"`
    Mode string `json:"mode"`
    Sensitivity string `json:"sensitivity"`
    FilterThreshold int `json:"filter-threshold"`
    FilterInactiveThreshold int `json:"filter-inactive-threshold"`
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortListPatternRecognitionPuDetails struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangeList struct {
    PortRangeStart int `json:"port-range-start"`
    PortRangeEnd int `json:"port-range-end"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    DetectionEnable int `json:"detection-enable"`
    EnableTopK int `json:"enable-top-k"`
    TopkNumRecords int `json:"topk-num-records" dval:"20"`
    Glid string `json:"glid"`
    GlidExceedAction DdosDstEntryPortRangeListGlidExceedAction `json:"glid-exceed-action"`
    Template DdosDstEntryPortRangeListTemplate `json:"template"`
    Sflow DdosDstEntryPortRangeListSflow `json:"sflow"`
    CaptureConfig DdosDstEntryPortRangeListCaptureConfig `json:"capture-config"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    IpFilteringPolicy string `json:"ip-filtering-policy"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IpFilteringPolicyOper DdosDstEntryPortRangeListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    PortInd DdosDstEntryPortRangeListPortInd `json:"port-ind"`
    TopkSources DdosDstEntryPortRangeListTopkSources `json:"topk-sources"`
    ProgressionTracking DdosDstEntryPortRangeListProgressionTracking `json:"progression-tracking"`
    PatternRecognition DdosDstEntryPortRangeListPatternRecognition `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstEntryPortRangeListPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
}


type DdosDstEntryPortRangeListGlidExceedAction struct {
    StatelessEncapActionCfg DdosDstEntryPortRangeListGlidExceedActionStatelessEncapActionCfg `json:"stateless-encap-action-cfg"`
}


type DdosDstEntryPortRangeListGlidExceedActionStatelessEncapActionCfg struct {
    StatelessEncapAction string `json:"stateless-encap-action"`
    EncapTemplate string `json:"encap-template"`
}


type DdosDstEntryPortRangeListTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
}


type DdosDstEntryPortRangeListSflow struct {
    Polling DdosDstEntryPortRangeListSflowPolling `json:"polling"`
}


type DdosDstEntryPortRangeListSflowPolling struct {
    SflowPackets int `json:"sflow-packets"`
    SflowTcp DdosDstEntryPortRangeListSflowPollingSflowTcp `json:"sflow-tcp"`
    SflowHttp int `json:"sflow-http"`
}


type DdosDstEntryPortRangeListSflowPollingSflowTcp struct {
    SflowTcpBasic int `json:"sflow-tcp-basic"`
    SflowTcpStateful int `json:"sflow-tcp-stateful"`
}


type DdosDstEntryPortRangeListCaptureConfig struct {
    CaptureConfigName string `json:"capture-config-name"`
    CaptureConfigMode string `json:"capture-config-mode"`
}


type DdosDstEntryPortRangeListIpFilteringPolicyOper struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangeListPortInd struct {
    Uuid string `json:"uuid"`
    SamplingEnable []DdosDstEntryPortRangeListPortIndSamplingEnable `json:"sampling-enable"`
}


type DdosDstEntryPortRangeListPortIndSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntryPortRangeListTopkSources struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangeListProgressionTracking struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangeListPatternRecognition struct {
    Algorithm string `json:"algorithm"`
    Mode string `json:"mode"`
    Sensitivity string `json:"sensitivity"`
    FilterThreshold int `json:"filter-threshold"`
    FilterInactiveThreshold int `json:"filter-inactive-threshold"`
    Uuid string `json:"uuid"`
}


type DdosDstEntryPortRangeListPatternRecognitionPuDetails struct {
    Uuid string `json:"uuid"`
}


type DdosDstEntrySamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
    Counters3 string `json:"counters3"`
}


type DdosDstEntrySflow struct {
    Polling DdosDstEntrySflowPolling `json:"polling"`
    Collector []DdosDstEntrySflowCollector `json:"collector"`
}


type DdosDstEntrySflowPolling struct {
    SflowPackets int `json:"sflow-packets"`
    SflowLayer4 int `json:"sflow-layer-4"`
    SflowTcp DdosDstEntrySflowPollingSflowTcp `json:"sflow-tcp"`
    SflowHttp int `json:"sflow-http"`
    SflowUndefPortHitStats int `json:"sflow-undef-port-hit-stats"`
    SflowUndefPortHitStatsBrief int `json:"sflow-undef-port-hit-stats-brief"`
}


type DdosDstEntrySflowPollingSflowTcp struct {
    SflowTcpBasic int `json:"sflow-tcp-basic"`
    SflowTcpStateful int `json:"sflow-tcp-stateful"`
}


type DdosDstEntrySflowCollector struct {
    SflowName string `json:"sflow-name"`
}


type DdosDstEntrySrcDstPair179 struct {
    Default int `json:"default"`
    Bypass int `json:"bypass"`
    ExceedLogCfg DdosDstEntrySrcDstPairExceedLogCfg180 `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairTemplate181 `json:"template"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairL4TypeSrcDstList182 `json:"l4-type-src-dst-list"`
    AppTypeSrcDstList []DdosDstEntrySrcDstPairAppTypeSrcDstList184 `json:"app-type-src-dst-list"`
}


type DdosDstEntrySrcDstPairExceedLogCfg180 struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairTemplate181 struct {
    Logging string `json:"logging"`
}


type DdosDstEntrySrcDstPairL4TypeSrcDstList182 struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate183 `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate183 struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairAppTypeSrcDstList184 struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate185 `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate185 struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairClassListList struct {
    ClassListName string `json:"class-list-name"`
    ExceedLogCfg DdosDstEntrySrcDstPairClassListListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairClassListListL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    AppTypeSrcDstList []DdosDstEntrySrcDstPairClassListListAppTypeSrcDstList `json:"app-type-src-dst-list"`
    CidList []DdosDstEntrySrcDstPairClassListListCidList `json:"cid-list"`
}


type DdosDstEntrySrcDstPairClassListListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairClassListListL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairClassListListL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListListL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairClassListListAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairClassListListAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListListAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairClassListListCidList struct {
    CidNum int `json:"cid-num"`
    ExceedLogCfg DdosDstEntrySrcDstPairClassListListCidListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstCidList []DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidList `json:"l4-type-src-dst-cid-list"`
    AppTypeSrcDstCidList []DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidList `json:"app-type-src-dst-cid-list"`
}


type DdosDstEntrySrcDstPairClassListListCidListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListListCidListL4TypeSrcDstCidListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairClassListListCidListAppTypeSrcDstCidListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairPolicyList struct {
    SrcBasedPolicyName string `json:"src-based-policy-name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    PolicyClassListList []DdosDstEntrySrcDstPairPolicyListPolicyClassListList `json:"policy-class-list-list"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    Bypass int `json:"bypass"`
    ExceedLogCfg DdosDstEntrySrcDstPairPolicyListPolicyClassListListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairPolicyListPolicyClassListListTemplate `json:"template"`
    Glid string `json:"glid"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDstEntrySrcDstPairPolicyListPolicyClassListListSamplingEnable `json:"sampling-enable"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstList `json:"l4-type-src-dst-list"`
    AppTypeSrcDstList []DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstList `json:"app-type-src-dst-list"`
    ClassListOverflowPolicyList []DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListTemplate struct {
    Logging string `json:"logging"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListL4TypeSrcDstListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListAppTypeSrcDstListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyList struct {
    DummyName string `json:"dummy-name"`
    Bypass int `json:"bypass"`
    ExceedLogCfg DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListExceedLogCfg `json:"exceed-log-cfg"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListTemplate `json:"template"`
    Glid string `json:"glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList `json:"l4-type-src-dst-overflow-list"`
    AppTypeSrcDstOverflowList []DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList `json:"app-type-src-dst-overflow-list"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListTemplate struct {
    Logging string `json:"logging"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList struct {
    Protocol string `json:"protocol"`
    Template DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcDstPairPolicyListPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate struct {
    SslL4 string `json:"ssl-l4"`
    Dns string `json:"dns"`
    Http string `json:"http"`
    Sip string `json:"sip"`
}


type DdosDstEntrySrcDstPairSettingsList struct {
    AllTypes string `json:"all-types"`
    Age int `json:"age" dval:"5"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`
    EnableClassListOverflow int `json:"enable-class-list-overflow"`
    SrcPrefixLen int `json:"src-prefix-len"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    L4TypeSrcDstList []DdosDstEntrySrcDstPairSettingsListL4TypeSrcDstList `json:"l4-type-src-dst-list"`
}


type DdosDstEntrySrcDstPairSettingsListL4TypeSrcDstList struct {
    Protocol string `json:"protocol"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    OutboundSrcTracking string `json:"outbound-src-tracking" dval:"disable"`
    Template DdosDstEntrySrcPortListTemplate `json:"template"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcPortListTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
    SrcDns string `json:"src-dns"`
}


type DdosDstEntrySrcPortRangeList struct {
    SrcPortRangeStart int `json:"src-port-range-start"`
    SrcPortRangeEnd int `json:"src-port-range-end"`
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosDstEntrySrcPortRangeListTemplate `json:"template"`
    SetCounterBaseVal int `json:"set-counter-base-val"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosDstEntrySrcPortRangeListTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
}


type DdosDstEntryTemplate struct {
    Logging string `json:"logging"`
}


type DdosDstEntryTopkDestinations186 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDstEntry) GetId() string{
    return url.QueryEscape(p.Inst.DstEntryName)
}

func (p *DdosDstEntry) getPath() string{
    return "ddos/dst/entry"
}

func (p *DdosDstEntry) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntry::Post")
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

func (p *DdosDstEntry) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntry::Get")
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
func (p *DdosDstEntry) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntry::Put")
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

func (p *DdosDstEntry) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntry::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
