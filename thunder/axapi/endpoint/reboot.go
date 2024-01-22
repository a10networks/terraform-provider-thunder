

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Reboot struct {
	Inst struct {

    All int `json:"all"`
    At int `json:"at"`
    Cancel int `json:"cancel"`
    DayOfMonth int `json:"day-of-month"`
    DayOfMonth2 int `json:"day-of-month-2"`
    Device int `json:"device"`
    In string `json:"in"`
    Month string `json:"month"`
    Month2 string `json:"month-2"`
    Reason string `json:"reason"`
    Reason2 string `json:"reason-2"`
    Reason3 string `json:"reason-3"`
    Time string `json:"time"`

	} `json:"reboot"`
}

func (p *Reboot) GetId() string{
    return "1"
}

func (p *Reboot) getPath() string{
    return "reboot"
}

func (p *Reboot) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Reboot::Post")
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

func (p *Reboot) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Reboot::Get")
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
func (p *Reboot) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Reboot::Put")
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

func (p *Reboot) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Reboot::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
