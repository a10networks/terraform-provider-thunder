

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPasswordPolicy struct {
	Inst struct {

    Aging string `json:"aging"`
    Complexity string `json:"complexity"`
    ForbidConsecutiveCharacter string `json:"forbid-consecutive-character" dval:"0"`
    History string `json:"history"`
    MinPswdLen int `json:"min-pswd-len"`
    RepeatCharacterCheck string `json:"repeat-character-check" dval:"disable"`
    UsernameCheck string `json:"username-check" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"password-policy"`
}

func (p *SystemPasswordPolicy) GetId() string{
    return "1"
}

func (p *SystemPasswordPolicy) getPath() string{
    return "system/password-policy"
}

func (p *SystemPasswordPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPasswordPolicy::Post")
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

func (p *SystemPasswordPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPasswordPolicy::Get")
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
func (p *SystemPasswordPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPasswordPolicy::Put")
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

func (p *SystemPasswordPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemPasswordPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
