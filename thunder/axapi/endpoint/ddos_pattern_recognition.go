

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosPatternRecognition struct {
	Inst struct {

    CaptureBackup int `json:"capture-backup"`
    CapturingTimeout int `json:"capturing-timeout"`
    Cpu DdosPatternRecognitionCpu290 `json:"cpu"`
    CpuLimit int `json:"cpu-limit"`
    DedicatedCpus int `json:"dedicated-cpus"`
    DisableAppPayloadAll int `json:"disable-app-payload-all"`
    ErrorTimeout int `json:"error-timeout"`
    ExtractingTimeout int `json:"extracting-timeout"`
    HardwareFilter string `json:"hardware-filter" dval:"disable"`
    SampleSize int `json:"sample-size"`
    SchedulingTimeout int `json:"scheduling-timeout"`
    Sensitivity string `json:"sensitivity"`
    SflowEventPeriodicInterval int `json:"sflow-event-periodic-interval" dval:"5"`
    Toggle string `json:"toggle" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"pattern-recognition"`
}


type DdosPatternRecognitionCpu290 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosPatternRecognition) GetId() string{
    return "1"
}

func (p *DdosPatternRecognition) getPath() string{
    return "ddos/pattern-recognition"
}

func (p *DdosPatternRecognition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosPatternRecognition::Post")
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

func (p *DdosPatternRecognition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosPatternRecognition::Get")
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
func (p *DdosPatternRecognition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosPatternRecognition::Put")
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

func (p *DdosPatternRecognition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosPatternRecognition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
