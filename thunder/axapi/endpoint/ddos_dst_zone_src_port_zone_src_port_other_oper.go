

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortZoneSrcPortOtherOper struct {
    
    Oper DdosDstZoneSrcPortZoneSrcPortOtherOperOper `json:"oper"`
    PortInd DdosDstZoneSrcPortZoneSrcPortOtherOperPortInd `json:"port-ind"`
    PortOther string `json:"port-other"`
    Protocol string `json:"protocol"`
    ZoneName string 

}
type DataDdosDstZoneSrcPortZoneSrcPortOtherOper struct {
    DtDdosDstZoneSrcPortZoneSrcPortOtherOper DdosDstZoneSrcPortZoneSrcPortOtherOper `json:"zone-src-port-other"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherOperOper struct {
    Ddos_entry_list []DdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list `json:"ddos_entry_list"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    Sources int `json:"sources"`
    SourcesAllEntries int `json:"sources-all-entries"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    Ipv6 string `json:"ipv6"`
    HwBlacklisted int `json:"hw-blacklisted"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherOperOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    BwState string `json:"bw-state"`
    Is_auth_passed string `json:"is_auth_passed"`
    Level int `json:"level"`
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


type DdosDstZoneSrcPortZoneSrcPortOtherOperPortInd struct {
    Oper DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper `json:"oper"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOper struct {
    Indicators []DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    CurrentLevel string `json:"current-level"`
    Details int `json:"details"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    IndicatorCfg []DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherOperPortIndOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}

func (p *DdosDstZoneSrcPortZoneSrcPortOtherOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneSrcPortZoneSrcPortOtherOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/src-port/zone-src-port-other/"+p.PortOther+"+"+p.Protocol+"/oper"
}

func (p *DdosDstZoneSrcPortZoneSrcPortOtherOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneSrcPortZoneSrcPortOtherOper,error) {
logger.Println("DdosDstZoneSrcPortZoneSrcPortOtherOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneSrcPortZoneSrcPortOtherOper
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
