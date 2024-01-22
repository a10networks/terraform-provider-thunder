

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FailSafe struct {
	Inst struct {

    Config FailSafeConfig349 `json:"config"`
    DataplaneRecoveryTimeout int `json:"dataplane-recovery-timeout" dval:"5"`
    DisableFailsafe FailSafeDisableFailsafe350 `json:"disable-failsafe"`
    FpgaBuffRecoveryThreshold int `json:"fpga-buff-recovery-threshold" dval:"2"`
    FpgaMonitorEnable int `json:"fpga-monitor-enable"`
    FpgaMonitorForcedReboot int `json:"fpga-monitor-forced-reboot"`
    FpgaMonitorInterval int `json:"fpga-monitor-interval" dval:"1"`
    FpgaMonitorThreshold int `json:"fpga-monitor-threshold" dval:"30"`
    HwErrorMonitor string `json:"hw-error-monitor" dval:"hw-error-monitor-enable"`
    HwErrorRecoveryTimeout int `json:"hw-error-recovery-timeout"`
    HwSslTimeoutMonitor string `json:"hw-ssl-timeout-monitor" dval:"hw-ssl-timeout-monitor-enable"`
    Kill int `json:"kill"`
    Log int `json:"log"`
    SessionMemRecoveryThreshold int `json:"session-mem-recovery-threshold" dval:"30"`
    SwErrorMonitorEnable int `json:"sw-error-monitor-enable"`
    SwErrorRecoveryTimeout int `json:"sw-error-recovery-timeout" dval:"3"`
    TotalMemorySizeCheck int `json:"total-memory-size-check"`
    Uuid string `json:"uuid"`

	} `json:"fail-safe"`
}


type FailSafeConfig349 struct {
    Uuid string `json:"uuid"`
}


type FailSafeDisableFailsafe350 struct {
    Action string `json:"action" dval:"all"`
    Uuid string `json:"uuid"`
}

func (p *FailSafe) GetId() string{
    return "1"
}

func (p *FailSafe) getPath() string{
    return "fail-safe"
}

func (p *FailSafe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FailSafe::Post")
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

func (p *FailSafe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FailSafe::Get")
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
func (p *FailSafe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FailSafe::Put")
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

func (p *FailSafe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FailSafe::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
