

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherOper struct {
    
    IpFilteringPolicyOper DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    Oper DdosDstZonePortZoneServiceOtherOperOper `json:"oper"`
    PatternRecognition DdosDstZonePortZoneServiceOtherOperPatternRecognition `json:"pattern-recognition"`
    PatternRecognitionPuDetails DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
    PortInd DdosDstZonePortZoneServiceOtherOperPortInd `json:"port-ind"`
    PortOther string `json:"port-other"`
    ProgressionTracking DdosDstZonePortZoneServiceOtherOperProgressionTracking `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    TopkDestinations DdosDstZonePortZoneServiceOtherOperTopkDestinations `json:"topk-destinations"`
    TopkSources DdosDstZonePortZoneServiceOtherOperTopkSources `json:"topk-sources"`
    ZoneName string 

}
type DataDdosDstZonePortZoneServiceOtherOper struct {
    DtDdosDstZonePortZoneServiceOtherOper DdosDstZonePortZoneServiceOtherOper `json:"zone-service-other"`
}


type DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOper struct {
    Oper DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOper struct {
    RuleList []DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstZonePortZoneServiceOtherOperIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstZonePortZoneServiceOtherOperOper struct {
    Ddos_entry_list []DdosDstZonePortZoneServiceOtherOperOperDdos_entry_list `json:"ddos_entry_list"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    Sources int `json:"sources"`
    OverflowPolicy int `json:"overflow-policy"`
    SourcesAllEntries int `json:"sources-all-entries"`
    ClassList string `json:"class-list"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    Ipv6 string `json:"ipv6"`
    Exceeded int `json:"exceeded"`
    BlackListed int `json:"black-listed"`
    WhiteListed int `json:"white-listed"`
    Authenticated int `json:"authenticated"`
    Level int `json:"level"`
    AppStat int `json:"app-stat"`
    Indicators int `json:"indicators"`
    IndicatorDetail int `json:"indicator-detail"`
    L4ExtRate int `json:"l4-ext-rate"`
    HwBlacklisted int `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstZonePortZoneServiceOtherOperOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    BwState string `json:"bw-state"`
    Is_auth_passed string `json:"is_auth_passed"`
    Level int `json:"level"`
    BlReasoningRcode string `json:"bl-reasoning-rcode"`
    BlReasoningTimestamp string `json:"bl-reasoning-timestamp"`
    CurrentConnections string `json:"current-connections"`
    IsConnectionsExceed int `json:"is-connections-exceed"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    IsConnectionRateExceed int `json:"is-connection-rate-exceed"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    IsPacketRateExceed int `json:"is-packet-rate-exceed"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    IsKbitRateExceed int `json:"is-kBit-rate-exceed"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    IsFragPacketRateExceed int `json:"is-frag-packet-rate-exceed"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    IsAppStat1Exceed int `json:"is-app-stat1-exceed"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    IsAppStat2Exceed int `json:"is-app-stat2-exceed"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    IsAppStat3Exceed int `json:"is-app-stat3-exceed"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    IsAppStat4Exceed int `json:"is-app-stat4-exceed"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    IsAppStat5Exceed int `json:"is-app-stat5-exceed"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    IsAppStat6Exceed int `json:"is-app-stat6-exceed"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    IsAppStat7Exceed int `json:"is-app-stat7-exceed"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    IsAppStat8Exceed int `json:"is-app-stat8-exceed"`
    AppStat8Limit string `json:"app-stat8-limit"`
    Age int `json:"age"`
    LockupTime int `json:"lockup-time"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId int `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstZonePortZoneServiceOtherOperPatternRecognition struct {
    Oper DdosDstZonePortZoneServiceOtherOperPatternRecognitionOper `json:"oper"`
}


type DdosDstZonePortZoneServiceOtherOperPatternRecognitionOper struct {
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList `json:"filter-list"`
}


type DdosDstZonePortZoneServiceOtherOperPatternRecognitionOperFilterList struct {
    ProcessingUnit string `json:"processing-unit"`
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetails struct {
    Oper DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper `json:"oper"`
}


type DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOper struct {
    AllFilters []DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}


type DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFilters struct {
    ProcessingUnit string `json:"processing-unit"`
    State string `json:"state"`
    Timestamp string `json:"timestamp"`
    PeacePktCount int `json:"peace-pkt-count"`
    WarPktCount int `json:"war-pkt-count"`
    WarPktPercentage int `json:"war-pkt-percentage"`
    FilterThreshold int `json:"filter-threshold"`
    FilterCount int `json:"filter-count"`
    FilterList []DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}


type DdosDstZonePortZoneServiceOtherOperPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
    FilterEnabled int `json:"filter-enabled"`
    HardwareFilter int `json:"hardware-filter"`
    FilterExpr string `json:"filter-expr"`
    FilterDesc string `json:"filter-desc"`
    SampleRatio int `json:"sample-ratio"`
}


type DdosDstZonePortZoneServiceOtherOperPortInd struct {
    Oper DdosDstZonePortZoneServiceOtherOperPortIndOper `json:"oper"`
}


type DdosDstZonePortZoneServiceOtherOperPortIndOper struct {
    SrcEntryList []DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList `json:"src-entry-list"`
    Indicators []DdosDstZonePortZoneServiceOtherOperPortIndOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    TotalScore string `json:"total-score"`
    CurrentLevel string `json:"current-level"`
    EscalationTimestamp string `json:"escalation-timestamp"`
    InitialLearning string `json:"initial-learning"`
    ActiveTime int `json:"active-time"`
    SourcesAllEntries int `json:"sources-all-entries"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    Ipv6 string `json:"ipv6"`
    Details int `json:"details"`
    Sources int `json:"sources"`
}


type DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryList struct {
    SrcAddressStr string `json:"src-address-str"`
    Indicators []DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    TotalScore string `json:"total-score"`
    CurrentLevel string `json:"current-level"`
    SrcLevel string `json:"src-level"`
    EscalationTimestamp string `json:"escalation-timestamp"`
    InitialLearning string `json:"initial-learning"`
    ActiveTime int `json:"active-time"`
}


type DdosDstZonePortZoneServiceOtherOperPortIndOperSrcEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    SrcMaximum string `json:"src-maximum"`
    SrcMinimum string `json:"src-minimum"`
    SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
    SrcAverage string `json:"src-average"`
    Score string `json:"score"`
}


type DdosDstZonePortZoneServiceOtherOperPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    ZoneMaximum string `json:"zone-maximum"`
    ZoneMinimum string `json:"zone-minimum"`
    ZoneNonZeroMinimum string `json:"zone-non-zero-minimum"`
    ZoneAverage string `json:"zone-average"`
    ZoneAdaptiveThreshold string `json:"zone-adaptive-threshold"`
    SrcMaximum string `json:"src-maximum"`
    IndicatorCfg []DdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
    Score string `json:"score"`
}


type DdosDstZonePortZoneServiceOtherOperPortIndOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}


type DdosDstZonePortZoneServiceOtherOperProgressionTracking struct {
    Oper DdosDstZonePortZoneServiceOtherOperProgressionTrackingOper `json:"oper"`
}


type DdosDstZonePortZoneServiceOtherOperProgressionTrackingOper struct {
    Indicators []DdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZonePortZoneServiceOtherOperProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstZonePortZoneServiceOtherOperTopkDestinations struct {
    Oper DdosDstZonePortZoneServiceOtherOperTopkDestinationsOper `json:"oper"`
}


type DdosDstZonePortZoneServiceOtherOperTopkDestinationsOper struct {
    Indicators []DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortZoneServiceOtherOperTopkDestinationsOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}


type DdosDstZonePortZoneServiceOtherOperTopkSources struct {
    Oper DdosDstZonePortZoneServiceOtherOperTopkSourcesOper `json:"oper"`
}


type DdosDstZonePortZoneServiceOtherOperTopkSourcesOper struct {
    Indicators []DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstZonePortZoneServiceOtherOperTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstZonePortZoneServiceOtherOperTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZonePortZoneServiceOtherOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceOtherOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service-other/"+p.PortOther+"+"+p.Protocol+"/oper"
}

func (p *DdosDstZonePortZoneServiceOtherOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceOtherOper,error) {
logger.Println("DdosDstZonePortZoneServiceOtherOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceOtherOper
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
