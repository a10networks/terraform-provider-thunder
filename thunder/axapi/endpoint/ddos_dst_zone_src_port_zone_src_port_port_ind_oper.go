

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortZoneSrcPortPortIndOper struct {
    
    Oper DdosDstZoneSrcPortZoneSrcPortPortIndOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZoneSrcPortZoneSrcPortPortIndOper struct {
    DtDdosDstZoneSrcPortZoneSrcPortPortIndOper DdosDstZoneSrcPortZoneSrcPortPortIndOper `json:"port-ind"`
}


type DdosDstZoneSrcPortZoneSrcPortPortIndOperOper struct {
    Indicators []DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    CurrentLevel string `json:"current-level"`
    Details int `json:"details"`
}


type DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    IndicatorCfg []DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
}


type DdosDstZoneSrcPortZoneSrcPortPortIndOperOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}

func (p *DdosDstZoneSrcPortZoneSrcPortPortIndOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneSrcPortZoneSrcPortPortIndOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/src-port/zone-src-port/" +p.PortNum + "+" +p.Protocol + "/port-ind/oper"
}

func (p *DdosDstZoneSrcPortZoneSrcPortPortIndOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneSrcPortZoneSrcPortPortIndOper,error) {
logger.Println("DdosDstZoneSrcPortZoneSrcPortPortIndOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneSrcPortZoneSrcPortPortIndOper
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
