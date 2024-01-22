

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTelemetryLogTopKAppSvcListOper struct {
    
    Oper SystemTelemetryLogTopKAppSvcListOperOper `json:"oper"`

}
type DataSystemTelemetryLogTopKAppSvcListOper struct {
    DtSystemTelemetryLogTopKAppSvcListOper SystemTelemetryLogTopKAppSvcListOper `json:"top-k-app-svc-list"`
}


type SystemTelemetryLogTopKAppSvcListOperOper struct {
    TopKAppSvcList []SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList `json:"top-k-app-svc-list"`
}


type SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcList struct {
    Indicator string `json:"indicator"`
    IndicatorList []SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList `json:"indicator-list"`
}


type SystemTelemetryLogTopKAppSvcListOperOperTopKAppSvcListIndicatorList struct {
    AppSvcUuid string `json:"app-svc-uuid"`
    AppSvcName string `json:"app-svc-name"`
    IndicatorValue int `json:"indicator-value"`
}

func (p *SystemTelemetryLogTopKAppSvcListOper) GetId() string{
    return "1"
}

func (p *SystemTelemetryLogTopKAppSvcListOper) getPath() string{
    return "system/telemetry-log/top-k-app-svc-list/oper"
}

func (p *SystemTelemetryLogTopKAppSvcListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTelemetryLogTopKAppSvcListOper,error) {
logger.Println("SystemTelemetryLogTopKAppSvcListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTelemetryLogTopKAppSvcListOper
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
