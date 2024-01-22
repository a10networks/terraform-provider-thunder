

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryOper struct {
    
    DstEntryName string `json:"dst-entry-name"`
    IpProtoList []DdosDstEntryOperIpProtoList `json:"ip-proto-list"`
    L4TypeList []DdosDstEntryOperL4TypeList `json:"l4-type-list"`
    Oper DdosDstEntryOperOper `json:"oper"`
    PortList []DdosDstEntryOperPortList `json:"port-list"`
    PortRangeList []DdosDstEntryOperPortRangeList `json:"port-range-list"`
    SrcPortList []DdosDstEntryOperSrcPortList `json:"src-port-list"`
    SrcPortRangeList []DdosDstEntryOperSrcPortRangeList `json:"src-port-range-list"`
    TopkDestinations DdosDstEntryOperTopkDestinations `json:"topk-destinations"`

}
type DataDdosDstEntryOper struct {
    DtDdosDstEntryOper DdosDstEntryOper `json:"entry"`
}


type DdosDstEntryOperIpProtoList struct {
    PortNum int `json:"port-num"`
    Oper DdosDstEntryOperIpProtoListOper `json:"oper"`
    IpFilteringPolicyOper DdosDstEntryOperIpProtoListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
}


type DdosDstEntryOperIpProtoListOper struct {
    Ddos_entry_list []DdosDstEntryOperIpProtoListOperDdos_entry_list `json:"ddos_entry_list"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    AllIpProtos int `json:"all-ip-protos"`
    PortProtocol string `json:"port-protocol"`
    AppStat int `json:"app-stat"`
    SflowSourceId int `json:"sflow-source-id"`
    HwBlacklisted string `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstEntryOperIpProtoListOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryOperIpProtoListIpFilteringPolicyOper struct {
    Oper DdosDstEntryOperIpProtoListIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstEntryOperIpProtoListIpFilteringPolicyOperOper struct {
    RuleList []DdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryOperIpProtoListIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstEntryOperL4TypeList struct {
    Protocol string `json:"protocol"`
    Oper DdosDstEntryOperL4TypeListOper `json:"oper"`
    IpFilteringPolicyOper DdosDstEntryOperL4TypeListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    PortInd DdosDstEntryOperL4TypeListPortInd `json:"port-ind"`
    TopkSources DdosDstEntryOperL4TypeListTopkSources `json:"topk-sources"`
    ProgressionTracking DdosDstEntryOperL4TypeListProgressionTracking `json:"progression-tracking"`
}


type DdosDstEntryOperL4TypeListOper struct {
    Ddos_entry_list []DdosDstEntryOperL4TypeListOperDdos_entry_list `json:"ddos_entry_list"`
    UndefinedPortHitStatsWellknown []DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown `json:"undefined-port-hit-stats-wellknown"`
    UndefinedPortHitStatsNonWellknown []DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown `json:"undefined-port-hit-stats-non-wellknown"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    UndefinedPortHitStatistics int `json:"undefined-port-hit-statistics"`
    UndefinedStatsPortNum int `json:"undefined-stats-port-num"`
    AllL4Types int `json:"all-l4-types"`
    HwBlacklisted string `json:"hw-blacklisted"`
}


type DdosDstEntryOperL4TypeListOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsWellknown struct {
    Port string `json:"port"`
    Counter string `json:"counter"`
}


type DdosDstEntryOperL4TypeListOperUndefinedPortHitStatsNonWellknown struct {
    PortStart string `json:"port-start"`
    PortEnd string `json:"port-end"`
    Status string `json:"status"`
}


type DdosDstEntryOperL4TypeListIpFilteringPolicyOper struct {
    Oper DdosDstEntryOperL4TypeListIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstEntryOperL4TypeListIpFilteringPolicyOperOper struct {
    RuleList []DdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryOperL4TypeListIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstEntryOperL4TypeListPortInd struct {
    Oper DdosDstEntryOperL4TypeListPortIndOper `json:"oper"`
}


type DdosDstEntryOperL4TypeListPortIndOper struct {
    Indicators []DdosDstEntryOperL4TypeListPortIndOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryOperL4TypeListPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}


type DdosDstEntryOperL4TypeListTopkSources struct {
    Oper DdosDstEntryOperL4TypeListTopkSourcesOper `json:"oper"`
}


type DdosDstEntryOperL4TypeListTopkSourcesOper struct {
    Indicators []DdosDstEntryOperL4TypeListTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryOperL4TypeListTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryOperL4TypeListTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryOperL4TypeListTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryOperL4TypeListTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryOperL4TypeListTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}


type DdosDstEntryOperL4TypeListProgressionTracking struct {
    Oper DdosDstEntryOperL4TypeListProgressionTrackingOper `json:"oper"`
}


type DdosDstEntryOperL4TypeListProgressionTrackingOper struct {
    Indicators []DdosDstEntryOperL4TypeListProgressionTrackingOperIndicators `json:"indicators"`
}


type DdosDstEntryOperL4TypeListProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstEntryOperOper struct {
    Ddos_entry_list []DdosDstEntryOperOperDdos_entry_list `json:"ddos_entry_list"`
    EntryAddressStr string `json:"entry-address-str"`
    TotalDynamicEntryCount string `json:"total-dynamic-entry-count"`
    TotalDynamicEntryLimit string `json:"total-dynamic-entry-limit"`
    UdpDynamicEntryCount string `json:"udp-dynamic-entry-count"`
    UdpDynamicEntryLimit string `json:"udp-dynamic-entry-limit"`
    TcpDynamicEntryCount string `json:"tcp-dynamic-entry-count"`
    TcpDynamicEntryLimit string `json:"tcp-dynamic-entry-limit"`
    IcmpDynamicEntryCount string `json:"icmp-dynamic-entry-count"`
    IcmpDynamicEntryLimit string `json:"icmp-dynamic-entry-limit"`
    OtherDynamicEntryCount string `json:"other-dynamic-entry-count"`
    OtherDynamicEntryLimit string `json:"other-dynamic-entry-limit"`
    OperationalMode string `json:"operational-mode"`
    TrafficDistributionStatus []DdosDstEntryOperOperTrafficDistributionStatus `json:"traffic-distribution-status"`
    DstEntryName string `json:"dst-entry-name"`
    SourceEntryLimit string `json:"source-entry-limit"`
    SourceEntryAlloc string `json:"source-entry-alloc"`
    SourceEntryRemain string `json:"source-entry-remain"`
    DstServiceLimit string `json:"dst-service-limit"`
    DstServiceAlloc string `json:"dst-service-alloc"`
    DstServiceRemain string `json:"dst-service-remain"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    NoT2IdxPortCount int `json:"no-t2-idx-port-count"`
    DstAllEntries int `json:"dst-all-entries"`
    Sources int `json:"sources"`
    SourcesAllEntries int `json:"sources-all-entries"`
    OverflowPolicy int `json:"overflow-policy"`
    EntryCount int `json:"entry-count"`
    SflowSourceId int `json:"sflow-source-id"`
    Ipv6 string `json:"ipv6"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    L4TypeStr string `json:"l4-type-str"`
    AppType string `json:"app-type"`
    Exceeded int `json:"exceeded"`
    BlackListed int `json:"black-listed"`
    WhiteListed int `json:"white-listed"`
    Authenticated int `json:"authenticated"`
    ClassList string `json:"class-list"`
    IpProtoNum int `json:"ip-proto-num"`
    PortNum int `json:"port-num"`
    PortRangeStart int `json:"port-range-start"`
    PortRangeEnd int `json:"port-range-end"`
    SrcPortNum int `json:"src-port-num"`
    SrcPortRangeStart int `json:"src-port-range-start"`
    SrcPortRangeEnd int `json:"src-port-range-end"`
    Protocol string `json:"protocol"`
    OptProtocol string `json:"opt-protocol"`
    SportProtocol string `json:"sport-protocol"`
    OptSportProtocol string `json:"opt-sport-protocol"`
    AppStat int `json:"app-stat"`
    PortAppStat int `json:"port-app-stat"`
    AllIpProtos int `json:"all-ip-protos"`
    AllL4Types int `json:"all-l4-types"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    BlackHoled int `json:"black-holed"`
    ResourceUsage int `json:"resource-usage"`
    DisplayTrafficDistributionStatus int `json:"display-traffic-distribution-status"`
    EntryStatus int `json:"entry-status"`
    L4ExtRate int `json:"l4-ext-rate"`
    HwBlacklisted string `json:"hw-blacklisted"`
}


type DdosDstEntryOperOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryOperOperTrafficDistributionStatus struct {
    MasterPu string `json:"master-pu"`
    ActivePu []DdosDstEntryOperOperTrafficDistributionStatusActivePu `json:"active-pu"`
}


type DdosDstEntryOperOperTrafficDistributionStatusActivePu struct {
    PuId string `json:"pu-id"`
}


type DdosDstEntryOperPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Oper DdosDstEntryOperPortListOper `json:"oper"`
    PortInd DdosDstEntryOperPortListPortInd `json:"port-ind"`
    IpFilteringPolicyOper DdosDstEntryOperPortListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    TopkSources DdosDstEntryOperPortListTopkSources `json:"topk-sources"`
    ProgressionTracking DdosDstEntryOperPortListProgressionTracking `json:"progression-tracking"`
    PatternRecognition DdosDstEntryOperPortListPatternRecognition `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstEntryOperPortListPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
}


type DdosDstEntryOperPortListOper struct {
    Ddos_entry_list []DdosDstEntryOperPortListOperDdos_entry_list `json:"ddos_entry_list"`
    ResourceLimitConfig string `json:"resource-limit-config"`
    ReourceLimitAlloc string `json:"reource-limit-alloc"`
    ResourceLimitRemain string `json:"resource-limit-remain"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    AllIpProtos int `json:"all-ip-protos"`
    PortProtocol string `json:"port-protocol"`
    AppStat int `json:"app-stat"`
    SflowSourceId int `json:"sflow-source-id"`
    ResourceUsage int `json:"resource-usage"`
    L4ExtRate int `json:"l4-ext-rate"`
    HwBlacklisted string `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstEntryOperPortListOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryOperPortListPortInd struct {
    Oper DdosDstEntryOperPortListPortIndOper `json:"oper"`
}


type DdosDstEntryOperPortListPortIndOper struct {
    Indicators []DdosDstEntryOperPortListPortIndOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryOperPortListPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}


type DdosDstEntryOperPortListIpFilteringPolicyOper struct {
    Oper DdosDstEntryOperPortListIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstEntryOperPortListIpFilteringPolicyOperOper struct {
    RuleList []DdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryOperPortListIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstEntryOperPortListTopkSources struct {
    Oper DdosDstEntryOperPortListTopkSourcesOper `json:"oper"`
}


type DdosDstEntryOperPortListTopkSourcesOper struct {
    Indicators []DdosDstEntryOperPortListTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryOperPortListTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryOperPortListTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryOperPortListTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryOperPortListTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryOperPortListTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryOperPortListTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryOperPortListTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}


type DdosDstEntryOperPortListProgressionTracking struct {
    Oper DdosDstEntryOperPortListProgressionTrackingOper `json:"oper"`
}


type DdosDstEntryOperPortListProgressionTrackingOper struct {
    Indicators []DdosDstEntryOperPortListProgressionTrackingOperIndicators `json:"indicators"`
}


type DdosDstEntryOperPortListProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstEntryOperPortListPatternRecognition struct {
    Oper DdosDstEntryOperPortListPatternRecognitionOper `json:"oper"`
}


type DdosDstEntryOperPortListPatternRecognitionOper struct {
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryOperPortListPatternRecognitionOperFilterList `json:"filter-list"`
}


type DdosDstEntryOperPortListPatternRecognitionOperFilterList struct {
    ProcessingUnit string `json:"processing-unit"`
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstEntryOperPortListPatternRecognitionPuDetails struct {
    Oper DdosDstEntryOperPortListPatternRecognitionPuDetailsOper `json:"oper"`
}


type DdosDstEntryOperPortListPatternRecognitionPuDetailsOper struct {
    AllFilters []DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}


type DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstEntryOperPortListPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstEntryOperPortRangeList struct {
    PortRangeStart int `json:"port-range-start"`
    PortRangeEnd int `json:"port-range-end"`
    Protocol string `json:"protocol"`
    Oper DdosDstEntryOperPortRangeListOper `json:"oper"`
    IpFilteringPolicyOper DdosDstEntryOperPortRangeListIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    PortInd DdosDstEntryOperPortRangeListPortInd `json:"port-ind"`
    TopkSources DdosDstEntryOperPortRangeListTopkSources `json:"topk-sources"`
    ProgressionTracking DdosDstEntryOperPortRangeListProgressionTracking `json:"progression-tracking"`
    PatternRecognition DdosDstEntryOperPortRangeListPatternRecognition `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstEntryOperPortRangeListPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
}


type DdosDstEntryOperPortRangeListOper struct {
    Ddos_entry_list []DdosDstEntryOperPortRangeListOperDdos_entry_list `json:"ddos_entry_list"`
    ResourceLimitConfig string `json:"resource-limit-config"`
    ReourceLimitAlloc string `json:"reource-limit-alloc"`
    ResourceLimitRemain string `json:"resource-limit-remain"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    AllIpProtos int `json:"all-ip-protos"`
    PortProtocol string `json:"port-protocol"`
    AppStat int `json:"app-stat"`
    SflowSourceId int `json:"sflow-source-id"`
    ResourceUsage int `json:"resource-usage"`
    L4ExtRate int `json:"l4-ext-rate"`
    HwBlacklisted string `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstEntryOperPortRangeListOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryOperPortRangeListIpFilteringPolicyOper struct {
    Oper DdosDstEntryOperPortRangeListIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstEntryOperPortRangeListIpFilteringPolicyOperOper struct {
    RuleList []DdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryOperPortRangeListIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstEntryOperPortRangeListPortInd struct {
    Oper DdosDstEntryOperPortRangeListPortIndOper `json:"oper"`
}


type DdosDstEntryOperPortRangeListPortIndOper struct {
    Indicators []DdosDstEntryOperPortRangeListPortIndOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryOperPortRangeListPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}


type DdosDstEntryOperPortRangeListTopkSources struct {
    Oper DdosDstEntryOperPortRangeListTopkSourcesOper `json:"oper"`
}


type DdosDstEntryOperPortRangeListTopkSourcesOper struct {
    Indicators []DdosDstEntryOperPortRangeListTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryOperPortRangeListTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryOperPortRangeListTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryOperPortRangeListTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryOperPortRangeListTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryOperPortRangeListTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}


type DdosDstEntryOperPortRangeListProgressionTracking struct {
    Oper DdosDstEntryOperPortRangeListProgressionTrackingOper `json:"oper"`
}


type DdosDstEntryOperPortRangeListProgressionTrackingOper struct {
    Indicators []DdosDstEntryOperPortRangeListProgressionTrackingOperIndicators `json:"indicators"`
}


type DdosDstEntryOperPortRangeListProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstEntryOperPortRangeListPatternRecognition struct {
    Oper DdosDstEntryOperPortRangeListPatternRecognitionOper `json:"oper"`
}


type DdosDstEntryOperPortRangeListPatternRecognitionOper struct {
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryOperPortRangeListPatternRecognitionOperFilterList `json:"filter-list"`
}


type DdosDstEntryOperPortRangeListPatternRecognitionOperFilterList struct {
    ProcessingUnit string `json:"processing-unit"`
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstEntryOperPortRangeListPatternRecognitionPuDetails struct {
    Oper DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper `json:"oper"`
}


type DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOper struct {
    AllFilters []DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}


type DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstEntryOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstEntryOperSrcPortList struct {
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Oper DdosDstEntryOperSrcPortListOper `json:"oper"`
}


type DdosDstEntryOperSrcPortListOper struct {
    Ddos_entry_list []DdosDstEntryOperSrcPortListOperDdos_entry_list `json:"ddos_entry_list"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    AllIpProtos int `json:"all-ip-protos"`
    PortProtocol string `json:"port-protocol"`
    AppStat int `json:"app-stat"`
    SflowSourceId int `json:"sflow-source-id"`
    HwBlacklisted string `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstEntryOperSrcPortListOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryOperSrcPortRangeList struct {
    SrcPortRangeStart int `json:"src-port-range-start"`
    SrcPortRangeEnd int `json:"src-port-range-end"`
    Protocol string `json:"protocol"`
    Oper DdosDstEntryOperSrcPortRangeListOper `json:"oper"`
}


type DdosDstEntryOperSrcPortRangeListOper struct {
    Ddos_entry_list []DdosDstEntryOperSrcPortRangeListOperDdos_entry_list `json:"ddos_entry_list"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    AllIpProtos int `json:"all-ip-protos"`
    PortProtocol string `json:"port-protocol"`
    AppStat int `json:"app-stat"`
    SflowSourceId int `json:"sflow-source-id"`
    HwBlacklisted string `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstEntryOperSrcPortRangeListOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstEntryOperTopkDestinations struct {
    Oper DdosDstEntryOperTopkDestinationsOper `json:"oper"`
}


type DdosDstEntryOperTopkDestinationsOper struct {
    Indicators []DdosDstEntryOperTopkDestinationsOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryOperTopkDestinationsOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryOperTopkDestinationsOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstEntryOperTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstEntryOperTopkDestinationsOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryOperTopkDestinationsOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryOperTopkDestinationsOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryOperTopkDestinationsOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstEntryOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryOper) getPath() string{
    
    return "ddos/dst/entry/"+p.DstEntryName+"/oper"
}

func (p *DdosDstEntryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryOper,error) {
logger.Println("DdosDstEntryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
