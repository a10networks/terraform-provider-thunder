

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberOper struct {
    
    IpFilteringPolicyOper DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    Oper DdosDstZoneIpProtoProtoNumberOperOper `json:"oper"`
    PortInd DdosDstZoneIpProtoProtoNumberOperPortInd `json:"port-ind"`
    ProgressionTracking DdosDstZoneIpProtoProtoNumberOperProgressionTracking `json:"progression-tracking"`
    ProtocolNum int `json:"protocol-num"`
    TopkDestinations DdosDstZoneIpProtoProtoNumberOperTopkDestinations `json:"topk-destinations"`
    TopkSources DdosDstZoneIpProtoProtoNumberOperTopkSources `json:"topk-sources"`
    ZoneName string 

}
type DataDdosDstZoneIpProtoProtoNumberOper struct {
    DtDdosDstZoneIpProtoProtoNumberOper DdosDstZoneIpProtoProtoNumberOper `json:"proto-number"`
}


type DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOper struct {
    Oper DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOper struct {
    RuleList []DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstZoneIpProtoProtoNumberOperIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstZoneIpProtoProtoNumberOperOper struct {
    Ddos_entry_list []DdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list `json:"ddos_entry_list"`
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
    HwBlacklisted int `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstZoneIpProtoProtoNumberOperOperDdos_entry_list struct {
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
    Age int `json:"age"`
    LockupTime int `json:"lockup-time"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId int `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}


type DdosDstZoneIpProtoProtoNumberOperPortInd struct {
    Oper DdosDstZoneIpProtoProtoNumberOperPortIndOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNumberOperPortIndOper struct {
    SrcEntryList []DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList `json:"src-entry-list"`
    Indicators []DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators `json:"indicators"`
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


type DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryList struct {
    SrcAddressStr string `json:"src-address-str"`
    Indicators []DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    TotalScore string `json:"total-score"`
    CurrentLevel string `json:"current-level"`
    SrcLevel string `json:"src-level"`
    EscalationTimestamp string `json:"escalation-timestamp"`
    InitialLearning string `json:"initial-learning"`
    ActiveTime int `json:"active-time"`
}


type DdosDstZoneIpProtoProtoNumberOperPortIndOperSrcEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    SrcMaximum string `json:"src-maximum"`
    SrcMinimum string `json:"src-minimum"`
    SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
    SrcAverage string `json:"src-average"`
    Score string `json:"score"`
}


type DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    ZoneMaximum string `json:"zone-maximum"`
    ZoneMinimum string `json:"zone-minimum"`
    ZoneNonZeroMinimum string `json:"zone-non-zero-minimum"`
    ZoneAverage string `json:"zone-average"`
    ZoneAdaptiveThreshold string `json:"zone-adaptive-threshold"`
    SrcMaximum string `json:"src-maximum"`
    IndicatorCfg []DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
    Score string `json:"score"`
}


type DdosDstZoneIpProtoProtoNumberOperPortIndOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}


type DdosDstZoneIpProtoProtoNumberOperProgressionTracking struct {
    Oper DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOper struct {
    Indicators []DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZoneIpProtoProtoNumberOperProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkDestinations struct {
    Oper DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOper struct {
    Indicators []DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkDestinationsOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkSources struct {
    Oper DdosDstZoneIpProtoProtoNumberOperTopkSourcesOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkSourcesOper struct {
    Indicators []DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNumberOperTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneIpProtoProtoNumberOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-number/" +strconv.Itoa(p.ProtocolNum)+"/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNumberOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNumberOper
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
