

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTelemetryLogDeviceStatusOper struct {
    
    Oper SystemTelemetryLogDeviceStatusOperOper `json:"oper"`

}
type DataSystemTelemetryLogDeviceStatusOper struct {
    DtSystemTelemetryLogDeviceStatusOper SystemTelemetryLogDeviceStatusOper `json:"device-status"`
}


type SystemTelemetryLogDeviceStatusOperOper struct {
    Memory_usage string `json:"memory_usage"`
    Control_cpu_usage int `json:"control_cpu_usage"`
    Cpu_usage_overall int `json:"cpu_usage_overall"`
    Ratio_session_count string `json:"ratio_session_count"`
    Ratio_buffer_count string `json:"ratio_buffer_count"`
    Total_bytes_in int `json:"total_bytes_in"`
    Total_bytes_out int `json:"total_bytes_out"`
}

func (p *SystemTelemetryLogDeviceStatusOper) GetId() string{
    return "1"
}

func (p *SystemTelemetryLogDeviceStatusOper) getPath() string{
    return "system/telemetry-log/device-status/oper"
}

func (p *SystemTelemetryLogDeviceStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTelemetryLogDeviceStatusOper,error) {
logger.Println("SystemTelemetryLogDeviceStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTelemetryLogDeviceStatusOper
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
