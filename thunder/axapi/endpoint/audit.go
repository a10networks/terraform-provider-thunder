

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Audit struct {
	Inst struct {

    Enable int `json:"enable" dval:"1"`
    Privilege int `json:"privilege"`
    Size int `json:"size"`
    Uuid string `json:"uuid"`

	} `json:"audit"`
}

func (p *Audit) GetId() string{
    return "1"
}

func (p *Audit) getPath() string{
    return "audit"
}

func (p *Audit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Audit::Post")
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

func (p *Audit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Audit::Get")
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
func (p *Audit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Audit::Put")
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

func (p *Audit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Audit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
