

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTso struct {
	Inst struct {

    Disable int `json:"disable"`
    Enable int `json:"enable"`

	} `json:"tso"`
}

func (p *SystemTso) GetId() string{
    return "1"
}

func (p *SystemTso) getPath() string{
    return "system/tso"
}

func (p *SystemTso) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTso::Post")
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

func (p *SystemTso) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTso::Get")
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
func (p *SystemTso) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTso::Put")
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

func (p *SystemTso) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTso::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
