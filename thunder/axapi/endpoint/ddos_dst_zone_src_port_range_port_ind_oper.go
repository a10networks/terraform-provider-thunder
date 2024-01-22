

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneSrcPortRangePortIndOper struct {
    
    Oper DdosDstZoneSrcPortRangePortIndOperOper `json:"oper"`
    SrcPortRangeEnd string 
    ZoneName string 
    SrcPortRangeStart string 
    Protocol string 

}
type DataDdosDstZoneSrcPortRangePortIndOper struct {
    DtDdosDstZoneSrcPortRangePortIndOper DdosDstZoneSrcPortRangePortIndOper `json:"port-ind"`
}


type DdosDstZoneSrcPortRangePortIndOperOper struct {
    Indicators []DdosDstZoneSrcPortRangePortIndOperOperIndicators `json:"indicators"`
    DetectionDataSource string `json:"detection-data-source"`
    CurrentLevel string `json:"current-level"`
    Details int `json:"details"`
}


type DdosDstZoneSrcPortRangePortIndOperOperIndicators struct {
    IndicatorName string `json:"indicator-name"`
    IndicatorIndex int `json:"indicator-index"`
    Rate string `json:"rate"`
    IndicatorCfg []DdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
}


type DdosDstZoneSrcPortRangePortIndOperOperIndicatorsIndicatorCfg struct {
    Level int `json:"level"`
    ZoneThreshold string `json:"zone-threshold"`
    SourceThreshold string `json:"source-threshold"`
}

func (p *DdosDstZoneSrcPortRangePortIndOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneSrcPortRangePortIndOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/src-port-range/" +p.SrcPortRangeStart + "+" +p.SrcPortRangeEnd + "+" +p.Protocol + "/port-ind/oper"
}

func (p *DdosDstZoneSrcPortRangePortIndOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneSrcPortRangePortIndOper,error) {
logger.Println("DdosDstZoneSrcPortRangePortIndOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneSrcPortRangePortIndOper
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
