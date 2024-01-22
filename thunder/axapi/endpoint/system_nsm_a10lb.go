

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemNsmA10lb struct {
	Inst struct {

    Kill int `json:"kill"`
    Uuid string `json:"uuid"`

	} `json:"nsm-a10lb"`
}

func (p *SystemNsmA10lb) GetId() string{
    return "1"
}

func (p *SystemNsmA10lb) getPath() string{
    return "system/nsm-a10lb"
}

func (p *SystemNsmA10lb) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemNsmA10lb::Post")
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

func (p *SystemNsmA10lb) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemNsmA10lb::Get")
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
func (p *SystemNsmA10lb) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemNsmA10lb::Put")
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

func (p *SystemNsmA10lb) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemNsmA10lb::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
