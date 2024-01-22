

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTelemetryLogTopKSourceListOper struct {
    
    Oper SystemTelemetryLogTopKSourceListOperOper `json:"oper"`

}
type DataSystemTelemetryLogTopKSourceListOper struct {
    DtSystemTelemetryLogTopKSourceListOper SystemTelemetryLogTopKSourceListOper `json:"top-k-source-list"`
}


type SystemTelemetryLogTopKSourceListOperOper struct {
    TopKSourceList []SystemTelemetryLogTopKSourceListOperOperTopKSourceList `json:"top-k-source-list"`
}


type SystemTelemetryLogTopKSourceListOperOperTopKSourceList struct {
    Indicator string `json:"indicator"`
    IndicatorList []SystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList `json:"indicator-list"`
}


type SystemTelemetryLogTopKSourceListOperOperTopKSourceListIndicatorList struct {
    IpAddress string `json:"ip-address"`
    IndicatorValue int `json:"indicator-value"`
}

func (p *SystemTelemetryLogTopKSourceListOper) GetId() string{
    return "1"
}

func (p *SystemTelemetryLogTopKSourceListOper) getPath() string{
    return "system/telemetry-log/top-k-source-list/oper"
}

func (p *SystemTelemetryLogTopKSourceListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTelemetryLogTopKSourceListOper,error) {
logger.Println("SystemTelemetryLogTopKSourceListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTelemetryLogTopKSourceListOper
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
