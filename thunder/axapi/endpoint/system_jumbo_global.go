

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemJumboGlobal struct {
	Inst struct {

    EnableJumbo int `json:"enable-jumbo"`
    Uuid string `json:"uuid"`

	} `json:"system-jumbo-global"`
}

func (p *SystemJumboGlobal) GetId() string{
    return "1"
}

func (p *SystemJumboGlobal) getPath() string{
    return "system-jumbo-global"
}

func (p *SystemJumboGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJumboGlobal::Post")
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

func (p *SystemJumboGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJumboGlobal::Get")
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
func (p *SystemJumboGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJumboGlobal::Put")
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

func (p *SystemJumboGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemJumboGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
