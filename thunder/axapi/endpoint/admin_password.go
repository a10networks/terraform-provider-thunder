

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AdminPassword struct {
	Inst struct {

    EncryptedInModule string `json:"encrypted-in-module"`
    PasswordInModule string `json:"password-in-module"`
    Uuid string `json:"uuid"`
    User string 

	} `json:"password"`
}

func (p *AdminPassword) GetId() string{
    return "1"
}

func (p *AdminPassword) getPath() string{
    return "admin/" +p.Inst.User + "/password"
}

func (p *AdminPassword) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AdminPassword::Post")
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

func (p *AdminPassword) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AdminPassword::Get")
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
func (p *AdminPassword) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AdminPassword::Put")
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

func (p *AdminPassword) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AdminPassword::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
