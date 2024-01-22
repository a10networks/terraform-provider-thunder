

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameOper struct {
    
    IpFilteringPolicyOper DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    Oper DdosDstZoneIpProtoProtoNameOperOper `json:"oper"`
    PortInd DdosDstZoneIpProtoProtoNameOperPortInd `json:"port-ind"`
    ProgressionTracking DdosDstZoneIpProtoProtoNameOperProgressionTracking `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    TopkDestinations DdosDstZoneIpProtoProtoNameOperTopkDestinations `json:"topk-destinations"`
    TopkSources DdosDstZoneIpProtoProtoNameOperTopkSources `json:"topk-sources"`
    ZoneName string 

}
type DataDdosDstZoneIpProtoProtoNameOper struct {
    DtDdosDstZoneIpProtoProtoNameOper DdosDstZoneIpProtoProtoNameOper `json:"proto-name"`
}


type DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOper struct {
    Oper DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOper struct {
    RuleList []DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstZoneIpProtoProtoNameOperIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstZoneIpProtoProtoNameOperOper struct {
    Ddos_entry_list []DdosDstZoneIpProtoProtoNameOperOperDdos_entry_list `json:"ddos_entry_list"`
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


type DdosDstZoneIpProtoProtoNameOperOperDdos_entry_list struct {
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


type DdosDstZoneIpProtoProtoNameOperPortInd struct {
    Oper DdosDstZoneIpProtoProtoNameOperPortIndOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNameOperPortIndOper struct {
    SrcEntryList []DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList `json:"src-entry-list"`
    Indicators []DdosDstZoneIpProtoProtoNameOperPortIndOperIndicators `json:"indicators"`
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


type DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryList struct {
    SrcAddressStr string `json:"src-address-str"`
    Indicators []DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    TotalScore string `json:"total-score"`
    CurrentLevel string `json:"current-level"`
    SrcLevel string `json:"src-level"`
    EscalationTimestamp string `json:"escalation-timestamp"`
    InitialLearning string `json:"initial-learning"`
    ActiveTime int `json:"active-time"`
}


type DdosDstZoneIpProtoProtoNameOperPortIndOperSrcEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    SrcMaximum string `json:"src-maximum"`
    SrcMinimum string `json:"src-minimum"`
    SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
    SrcAverage string `json:"src-average"`
    Score string `json:"score"`
}


type DdosDstZoneIpProtoProtoNameOperPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    ZoneMaximum string `json:"zone-maximum"`
    ZoneMinimum string `json:"zone-minimum"`
    ZoneNonZeroMinimum string `json:"zone-non-zero-minimum"`
    ZoneAverage string `json:"zone-average"`
    ZoneAdaptiveThreshold string `json:"zone-adaptive-threshold"`
    SrcMaximum string `json:"src-maximum"`
    IndicatorCfg []DdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
    Score string `json:"score"`
}


type DdosDstZoneIpProtoProtoNameOperPortIndOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}


type DdosDstZoneIpProtoProtoNameOperProgressionTracking struct {
    Oper DdosDstZoneIpProtoProtoNameOperProgressionTrackingOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNameOperProgressionTrackingOper struct {
    Indicators []DdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators `json:"indicators"`
    LearningDetails int `json:"learning-details"`
    LearningBrief int `json:"learning-brief"`
    RecommendedTemplate int `json:"recommended-template"`
    TemplateDebugTable int `json:"template-debug-table"`
}


type DdosDstZoneIpProtoProtoNameOperProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstZoneIpProtoProtoNameOperTopkDestinations struct {
    Oper DdosDstZoneIpProtoProtoNameOperTopkDestinationsOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNameOperTopkDestinationsOper struct {
    Indicators []DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Destinations []DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}


type DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperIndicatorsDestinations struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNameOperTopkDestinationsOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}


type DdosDstZoneIpProtoProtoNameOperTopkSources struct {
    Oper DdosDstZoneIpProtoProtoNameOperTopkSourcesOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoNameOperTopkSourcesOper struct {
    Indicators []DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstZoneIpProtoProtoNameOperTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstZoneIpProtoProtoNameOperTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneIpProtoProtoNameOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNameOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-name/"+p.Protocol+"/oper"
}

func (p *DdosDstZoneIpProtoProtoNameOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNameOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNameOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNameOper
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
