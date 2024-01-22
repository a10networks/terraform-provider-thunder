

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTelemetryLogPartitionMetricsOper struct {
    
    Oper SystemTelemetryLogPartitionMetricsOperOper `json:"oper"`

}
type DataSystemTelemetryLogPartitionMetricsOper struct {
    DtSystemTelemetryLogPartitionMetricsOper SystemTelemetryLogPartitionMetricsOper `json:"partition-metrics"`
}


type SystemTelemetryLogPartitionMetricsOperOper struct {
    Data_cpu_usage int `json:"data_cpu_usage"`
}

func (p *SystemTelemetryLogPartitionMetricsOper) GetId() string{
    return "1"
}

func (p *SystemTelemetryLogPartitionMetricsOper) getPath() string{
    return "system/telemetry-log/partition-metrics/oper"
}

func (p *SystemTelemetryLogPartitionMetricsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTelemetryLogPartitionMetricsOper,error) {
logger.Println("SystemTelemetryLogPartitionMetricsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTelemetryLogPartitionMetricsOper
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
