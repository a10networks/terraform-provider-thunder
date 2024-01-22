

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGlid struct {
	Inst struct {

    GlidId string `json:"glid-id"`
    NonShared int `json:"non-shared"`
    Uuid string `json:"uuid"`

	} `json:"glid"`
}

func (p *SystemGlid) GetId() string{
    return "1"
}

func (p *SystemGlid) getPath() string{
    return "system/glid"
}

func (p *SystemGlid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGlid::Post")
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

func (p *SystemGlid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGlid::Get")
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
func (p *SystemGlid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGlid::Put")
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

func (p *SystemGlid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemGlid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
