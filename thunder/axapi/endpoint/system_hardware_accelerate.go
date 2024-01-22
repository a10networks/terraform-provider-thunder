

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemHardwareAccelerate struct {
	Inst struct {

    SamplingEnable []SystemHardwareAccelerateSamplingEnable `json:"sampling-enable"`
    SessionForwarding int `json:"session-forwarding"`
    Slb SystemHardwareAccelerateSlb1574 `json:"slb"`
    Uuid string `json:"uuid"`

	} `json:"hardware-accelerate"`
}


type SystemHardwareAccelerateSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SystemHardwareAccelerateSlb1574 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemHardwareAccelerateSlbSamplingEnable1575 `json:"sampling-enable"`
}


type SystemHardwareAccelerateSlbSamplingEnable1575 struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemHardwareAccelerate) GetId() string{
    return "1"
}

func (p *SystemHardwareAccelerate) getPath() string{
    return "system/hardware-accelerate"
}

func (p *SystemHardwareAccelerate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHardwareAccelerate::Post")
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

func (p *SystemHardwareAccelerate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHardwareAccelerate::Get")
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
func (p *SystemHardwareAccelerate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHardwareAccelerate::Put")
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

func (p *SystemHardwareAccelerate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemHardwareAccelerate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
