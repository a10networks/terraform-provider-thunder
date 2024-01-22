

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type KeyKey struct {
	Inst struct {

    KeyNumber int `json:"key-number"`
    KeyString string `json:"key-string"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    KeyChainName string 
    KeyChainFlag string 

	} `json:"key"`
}

func (p *KeyKey) GetId() string{
    return strconv.Itoa(p.Inst.KeyNumber)
}

func (p *KeyKey) getPath() string{
    return "key/" +p.Inst.KeyChainFlag + "+" +p.Inst.KeyChainName + "/key"
}

func (p *KeyKey) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("KeyKey::Post")
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

func (p *KeyKey) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("KeyKey::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *KeyKey) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("KeyKey::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *KeyKey) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("KeyKey::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
