

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnablePassword struct {
	Inst struct {

    Encrypted string `json:"encrypted"`
    Password string `json:"password"`

	} `json:"enable-password"`
}

func (p *EnablePassword) GetId() string{
    return "1"
}

func (p *EnablePassword) getPath() string{
    return "enable-password"
}

func (p *EnablePassword) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnablePassword::Post")
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

func (p *EnablePassword) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnablePassword::Get")
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
func (p *EnablePassword) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnablePassword::Put")
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

func (p *EnablePassword) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnablePassword::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
