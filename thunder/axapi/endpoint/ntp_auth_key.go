

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NtpAuthKey struct {
	Inst struct {

    AlgType string `json:"alg-type"`
    AscKey string `json:"asc-key"`
    Encrypted string `json:"encrypted"`
    HexEncrypted string `json:"hex-encrypted"`
    HexKey string `json:"hex-key"`
    Key int `json:"key"`
    KeyType string `json:"key-type"`
    Uuid string `json:"uuid"`

	} `json:"auth-key"`
}

func (p *NtpAuthKey) GetId() string{
    return strconv.Itoa(p.Inst.Key)
}

func (p *NtpAuthKey) getPath() string{
    return "ntp/auth-key"
}

func (p *NtpAuthKey) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NtpAuthKey::Post")
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

func (p *NtpAuthKey) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NtpAuthKey::Get")
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
func (p *NtpAuthKey) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NtpAuthKey::Put")
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

func (p *NtpAuthKey) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NtpAuthKey::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
