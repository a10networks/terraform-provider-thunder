

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemDelPort struct {
	Inst struct {

    PortIndex int `json:"port-index"`

	} `json:"del-port"`
}

func (p *SystemDelPort) GetId() string{
    return "1"
}

func (p *SystemDelPort) getPath() string{
    return "system/del-port"
}

func (p *SystemDelPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDelPort::Post")
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

func (p *SystemDelPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDelPort::Get")
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
func (p *SystemDelPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDelPort::Put")
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

func (p *SystemDelPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemDelPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
