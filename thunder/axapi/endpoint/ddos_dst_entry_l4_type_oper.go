

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryL4TypeOper struct {
    
    IpFilteringPolicyOper DdosDstEntryL4TypeOperIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    Oper DdosDstEntryL4TypeOperOper `json:"oper"`
    PortInd DdosDstEntryL4TypeOperPortInd `json:"port-ind"`
    ProgressionTracking DdosDstEntryL4TypeOperProgressionTracking `json:"progression-tracking"`
    Protocol string `json:"protocol"`
    TopkSources DdosDstEntryL4TypeOperTopkSources `json:"topk-sources"`
    DstEntryName string 

}
type DataDdosDstEntryL4TypeOper struct {
    DtDdosDstEntryL4TypeOper DdosDstEntryL4TypeOper `json:"l4-type"`
}


type DdosDstEntryL4TypeOperIpFilteringPolicyOper struct {
    Oper DdosDstEntryL4TypeOperIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstEntryL4TypeOperIpFilteringPolicyOperOper struct {
    RuleList []DdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryL4TypeOperIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstEntryL4TypeOperOper struct {
    Ddos_entry_list []DdosDstEntryL4TypeOperOperDdos_entry_list `json:"ddos_entry_list"`
    UndefinedPortHitStatsWellknown []DdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown `json:"undefined-port-hit-stats-wellknown"`
    UndefinedPortHitStatsNonWellknown []DdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown `json:"undefined-port-hit-stats-non-wellknown"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    UndefinedPortHitStatistics int `json:"undefined-port-hit-statistics"`
    UndefinedStatsPortNum int `json:"undefined-stats-port-num"`
    AllL4Types int `json:"all-l4-types"`
    HwBlacklisted string `json:"hw-blacklisted"`
}


type DdosDstEntryL4TypeOperOperDdos_entry_list struct {
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


type DdosDstEntryL4TypeOperOperUndefinedPortHitStatsWellknown struct {
    Port string `json:"port"`
    Counter string `json:"counter"`
}


type DdosDstEntryL4TypeOperOperUndefinedPortHitStatsNonWellknown struct {
    PortStart string `json:"port-start"`
    PortEnd string `json:"port-end"`
    Status string `json:"status"`
}


type DdosDstEntryL4TypeOperPortInd struct {
    Oper DdosDstEntryL4TypeOperPortIndOper `json:"oper"`
}


type DdosDstEntryL4TypeOperPortIndOper struct {
    Indicators []DdosDstEntryL4TypeOperPortIndOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
}


type DdosDstEntryL4TypeOperPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    EntryMaximum string `json:"entry-maximum"`
    EntryMinimum string `json:"entry-minimum"`
    EntryNonZeroMinimum string `json:"entry-non-zero-minimum"`
    EntryAverage string `json:"entry-average"`
    SrcMaximum string `json:"src-maximum"`
}


type DdosDstEntryL4TypeOperProgressionTracking struct {
    Oper DdosDstEntryL4TypeOperProgressionTrackingOper `json:"oper"`
}


type DdosDstEntryL4TypeOperProgressionTrackingOper struct {
    Indicators []DdosDstEntryL4TypeOperProgressionTrackingOperIndicators `json:"indicators"`
}


type DdosDstEntryL4TypeOperProgressionTrackingOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    NumSample int `json:"num-sample"`
    Average string `json:"average"`
    Maximum string `json:"maximum"`
    Minimum string `json:"minimum"`
    StandardDeviation string `json:"standard-deviation"`
}


type DdosDstEntryL4TypeOperTopkSources struct {
    Oper DdosDstEntryL4TypeOperTopkSourcesOper `json:"oper"`
}


type DdosDstEntryL4TypeOperTopkSourcesOper struct {
    Indicators []DdosDstEntryL4TypeOperTopkSourcesOperIndicators `json:"indicators"`
    EntryList []DdosDstEntryL4TypeOperTopkSourcesOperEntryList `json:"entry-list"`
    NextIndicator int `json:"next-indicator"`
    Finished int `json:"finished"`
    Details int `json:"details"`
    TopKKey string `json:"top-k-key"`
}


type DdosDstEntryL4TypeOperTopkSourcesOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Sources []DdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources `json:"sources"`
}


type DdosDstEntryL4TypeOperTopkSourcesOperIndicatorsSources struct {
    Address string `json:"address"`
    Rate string `json:"rate"`
}


type DdosDstEntryL4TypeOperTopkSourcesOperEntryList struct {
    AddressStr string `json:"address-str"`
    Indicators []DdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators `json:"indicators"`
}


type DdosDstEntryL4TypeOperTopkSourcesOperEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    MaxPeak string `json:"max-peak"`
    PsdWdwCnt int `json:"psd-wdw-cnt"`
}

func (p *DdosDstEntryL4TypeOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryL4TypeOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/l4-type/"+p.Protocol+"/oper"
}

func (p *DdosDstEntryL4TypeOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryL4TypeOper,error) {
logger.Println("DdosDstEntryL4TypeOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryL4TypeOper
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
