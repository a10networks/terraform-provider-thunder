

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper struct {
    
    Oper DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper `json:"oper"`
    PortOther string 
    ZoneName string 
    Protocol string 

}
type DataDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper struct {
    DtDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper `json:"port-ind"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOper struct {
    Indicators []DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    CurrentLevel string `json:"current-level"`
    Details int `json:"details"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    IndicatorCfg []DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
}


type DdosDstZoneSrcPortZoneSrcPortOtherPortIndOperOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}

func (p *DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/src-port/zone-src-port-other/" +p.PortOther + "+" +p.Protocol + "/port-ind/oper"
}

func (p *DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper,error) {
logger.Println("DdosDstZoneSrcPortZoneSrcPortOtherPortIndOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneSrcPortZoneSrcPortOtherPortIndOper
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
