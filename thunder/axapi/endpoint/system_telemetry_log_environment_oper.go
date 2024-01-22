

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTelemetryLogEnvironmentOper struct {
    
    Oper SystemTelemetryLogEnvironmentOperOper `json:"oper"`

}
type DataSystemTelemetryLogEnvironmentOper struct {
    DtSystemTelemetryLogEnvironmentOper SystemTelemetryLogEnvironmentOper `json:"environment"`
}


type SystemTelemetryLogEnvironmentOperOper struct {
    PhysicalSystemTemperature int `json:"physical-system-temperature"`
    HwFanSetting string `json:"hw-fan-setting"`
    FanReport string `json:"fan-report"`
    FanValue int `json:"fan-value"`
    SystemVoltage string `json:"system-voltage"`
    PowerUnit string `json:"power-unit"`
}

func (p *SystemTelemetryLogEnvironmentOper) GetId() string{
    return "1"
}

func (p *SystemTelemetryLogEnvironmentOper) getPath() string{
    return "system/telemetry-log/environment/oper"
}

func (p *SystemTelemetryLogEnvironmentOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTelemetryLogEnvironmentOper,error) {
logger.Println("SystemTelemetryLogEnvironmentOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTelemetryLogEnvironmentOper
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
