

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Key struct {
	Inst struct {

    KeyChainFlag int `json:"key-chain-flag"`
    KeyChainName string `json:"key-chain-name"`
    KeyList []KeyKeyList `json:"key-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"key"`
}


type KeyKeyList struct {
    KeyNumber int `json:"key-number"`
    KeyString string `json:"key-string"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *Key) GetId() string{
    return strconv.Itoa(p.Inst.KeyChainFlag)+"+"+p.Inst.KeyChainName
}

func (p *Key) getPath() string{
    return "key"
}

func (p *Key) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Key::Post")
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

func (p *Key) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Key::Get")
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
func (p *Key) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Key::Put")
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

func (p *Key) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Key::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
