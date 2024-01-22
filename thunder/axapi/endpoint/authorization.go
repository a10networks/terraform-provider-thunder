

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Authorization struct {
	Inst struct {

    Commands int `json:"commands"`
    Debug int `json:"debug"`
    Method AuthorizationMethod `json:"method"`
    Uuid string `json:"uuid"`

	} `json:"authorization"`
}


type AuthorizationMethod struct {
    Tacplus int `json:"tacplus"`
    None int `json:"none"`
}

func (p *Authorization) GetId() string{
    return "1"
}

func (p *Authorization) getPath() string{
    return "authorization"
}

func (p *Authorization) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Authorization::Post")
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

func (p *Authorization) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Authorization::Get")
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
func (p *Authorization) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Authorization::Put")
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

func (p *Authorization) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Authorization::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
