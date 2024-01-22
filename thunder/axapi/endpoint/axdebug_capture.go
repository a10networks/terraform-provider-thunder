

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AxdebugCapture struct {
	Inst struct {

    Brief int `json:"brief"`
    CurrentSlot int `json:"current-slot"`
    Detail int `json:"detail"`
    NoStop int `json:"no-stop"`
    Save string `json:"save"`

	} `json:"capture"`
}

func (p *AxdebugCapture) GetId() string{
    return "1"
}

func (p *AxdebugCapture) getPath() string{
    return "axdebug/capture"
}

func (p *AxdebugCapture) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugCapture::Post")
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

func (p *AxdebugCapture) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugCapture::Get")
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
func (p *AxdebugCapture) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugCapture::Put")
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

func (p *AxdebugCapture) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugCapture::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
