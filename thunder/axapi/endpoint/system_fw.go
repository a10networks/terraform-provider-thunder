

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemFw struct {
	Inst struct {

    ApplicationFlow int `json:"application-flow"`
    ApplicationMempool int `json:"application-mempool"`
    BasicDpiEnable int `json:"basic-dpi-enable"`
    Uuid string `json:"uuid"`

	} `json:"fw"`
}

func (p *SystemFw) GetId() string{
    return "1"
}

func (p *SystemFw) getPath() string{
    return "system/fw"
}

func (p *SystemFw) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFw::Post")
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

func (p *SystemFw) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFw::Get")
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
func (p *SystemFw) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFw::Put")
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

func (p *SystemFw) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemFw::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
