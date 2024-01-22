

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberPortIndOper struct {
    
    Oper DdosDstZoneIpProtoProtoNumberPortIndOperOper `json:"oper"`
    ZoneName string 
    ProtocolNum string 

}
type DataDdosDstZoneIpProtoProtoNumberPortIndOper struct {
    DtDdosDstZoneIpProtoProtoNumberPortIndOper DdosDstZoneIpProtoProtoNumberPortIndOper `json:"port-ind"`
}


type DdosDstZoneIpProtoProtoNumberPortIndOperOper struct {
    SrcEntryList []DdosDstZoneIpProtoProtoNumberPortIndOperOperSrcEntryList `json:"src-entry-list"`
    Indicators []DdosDstZoneIpProtoProtoNumberPortIndOperOperIndicators `json:"indicators"`
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


type DdosDstZoneIpProtoProtoNumberPortIndOperOperSrcEntryList struct {
    SrcAddressStr string `json:"src-address-str"`
    Indicators []DdosDstZoneIpProtoProtoNumberPortIndOperOperSrcEntryListIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    TotalScore string `json:"total-score"`
    CurrentLevel string `json:"current-level"`
    SrcLevel string `json:"src-level"`
    EscalationTimestamp string `json:"escalation-timestamp"`
    InitialLearning string `json:"initial-learning"`
    ActiveTime int `json:"active-time"`
}


type DdosDstZoneIpProtoProtoNumberPortIndOperOperSrcEntryListIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    SrcMaximum string `json:"src-maximum"`
    SrcMinimum string `json:"src-minimum"`
    SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
    SrcAverage string `json:"src-average"`
    Score string `json:"score"`
}


type DdosDstZoneIpProtoProtoNumberPortIndOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    ZoneMaximum string `json:"zone-maximum"`
    ZoneMinimum string `json:"zone-minimum"`
    ZoneNonZeroMinimum string `json:"zone-non-zero-minimum"`
    ZoneAverage string `json:"zone-average"`
    ZoneAdaptiveThreshold string `json:"zone-adaptive-threshold"`
    SrcMaximum string `json:"src-maximum"`
    IndicatorCfg []DdosDstZoneIpProtoProtoNumberPortIndOperOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
    Score string `json:"score"`
}


type DdosDstZoneIpProtoProtoNumberPortIndOperOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}

func (p *DdosDstZoneIpProtoProtoNumberPortIndOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberPortIndOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-number/" +p.ProtocolNum + "/port-ind/oper"
}

func (p *DdosDstZoneIpProtoProtoNumberPortIndOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberPortIndOper,error) {
logger.Println("DdosDstZoneIpProtoProtoNumberPortIndOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNumberPortIndOper
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
