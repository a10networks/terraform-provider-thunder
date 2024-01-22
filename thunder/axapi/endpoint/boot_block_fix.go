

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type BootBlockFix struct {
	Inst struct {

    Cf int `json:"cf"`
    Hd int `json:"hd"`

	} `json:"boot-block-fix"`
}

func (p *BootBlockFix) GetId() string{
    return "1"
}

func (p *BootBlockFix) getPath() string{
    return "boot-block-fix"
}

func (p *BootBlockFix) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("BootBlockFix::Post")
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

func (p *BootBlockFix) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("BootBlockFix::Get")
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
func (p *BootBlockFix) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("BootBlockFix::Put")
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

func (p *BootBlockFix) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("BootBlockFix::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
