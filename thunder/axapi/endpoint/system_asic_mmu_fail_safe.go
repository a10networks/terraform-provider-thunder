

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemAsicMmuFailSafe struct {
	Inst struct {

    InjectError int `json:"inject-error"`
    MonitorDisable int `json:"monitor-disable"`
    MonitorInterval int `json:"monitor-interval" dval:"60"`
    RebootDisable int `json:"reboot-disable"`
    RecoveryThreshold int `json:"recovery-threshold" dval:"2"`
    TestPatternType string `json:"test-pattern-type" dval:"lcb"`
    Uuid string `json:"uuid"`

	} `json:"asic-mmu-fail-safe"`
}

func (p *SystemAsicMmuFailSafe) GetId() string{
    return "1"
}

func (p *SystemAsicMmuFailSafe) getPath() string{
    return "system/asic-mmu-fail-safe"
}

func (p *SystemAsicMmuFailSafe) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAsicMmuFailSafe::Post")
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

func (p *SystemAsicMmuFailSafe) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAsicMmuFailSafe::Get")
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
func (p *SystemAsicMmuFailSafe) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAsicMmuFailSafe::Put")
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

func (p *SystemAsicMmuFailSafe) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemAsicMmuFailSafe::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
