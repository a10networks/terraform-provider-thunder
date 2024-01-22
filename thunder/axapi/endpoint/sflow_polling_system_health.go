

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SflowPollingSystemHealth struct {
	Inst struct {

    LicenseStatistics string `json:"license-statistics" dval:"disable"`
    PerControlCpuUsage string `json:"per-control-cpu-usage" dval:"disable"`
    PerDataCpuUsage string `json:"per-data-cpu-usage" dval:"disable"`
    SystemHealthUsage string `json:"system-health-usage" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"system-health"`
}

func (p *SflowPollingSystemHealth) GetId() string{
    return "1"
}

func (p *SflowPollingSystemHealth) getPath() string{
    return "sflow/polling/system-health"
}

func (p *SflowPollingSystemHealth) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowPollingSystemHealth::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *SflowPollingSystemHealth) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowPollingSystemHealth::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SflowPollingSystemHealth) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowPollingSystemHealth::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SflowPollingSystemHealth) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowPollingSystemHealth::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
