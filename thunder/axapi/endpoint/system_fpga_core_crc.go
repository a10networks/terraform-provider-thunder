

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemFpgaCoreCrc struct {
	Inst struct {

    MonitorDisable int `json:"monitor-disable"`
    RebootEnable int `json:"reboot-enable"`
    Uuid string `json:"uuid"`

	} `json:"fpga-core-crc"`
}

func (p *SystemFpgaCoreCrc) GetId() string{
    return "1"
}

func (p *SystemFpgaCoreCrc) getPath() string{
    return "system/fpga-core-crc"
}

func (p *SystemFpgaCoreCrc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFpgaCoreCrc::Post")
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

func (p *SystemFpgaCoreCrc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFpgaCoreCrc::Get")
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
func (p *SystemFpgaCoreCrc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFpgaCoreCrc::Put")
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

func (p *SystemFpgaCoreCrc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFpgaCoreCrc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
