

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HsmTemplate struct {
	Inst struct {

    Encrypted string `json:"encrypted"`
    EnrollTimeout int `json:"enroll-timeout"`
    HealthCheckInterval int `json:"health-check-interval"`
    HsmDev []HsmTemplateHsmDev `json:"hsm-dev"`
    Password int `json:"password"`
    PasswordString string `json:"password-string"`
    Protection int `json:"protection"`
    ProtectionModule int `json:"protection-module"`
    ProtectionOcs int `json:"protection-ocs"`
    ProtectionSoftcardHash string `json:"protection-softcard-hash"`
    RfsIp string `json:"rfs-ip"`
    RfsPort int `json:"rfs-port"`
    SecWorld string `json:"sec-world"`
    Softcard int `json:"softcard"`
    SofthsmEnum string `json:"softhsm-enum"`
    TemplateName string `json:"template-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Worker int `json:"worker"`

	} `json:"template"`
}


type HsmTemplateHsmDev struct {
    HsmIp string `json:"hsm-ip"`
    HsmPort int `json:"hsm-port"`
    HsmPriority int `json:"hsm-priority"`
}

func (p *HsmTemplate) GetId() string{
    return p.Inst.TemplateName
}

func (p *HsmTemplate) getPath() string{
    return "hsm/template"
}

func (p *HsmTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HsmTemplate::Post")
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

func (p *HsmTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HsmTemplate::Get")
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
func (p *HsmTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HsmTemplate::Put")
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

func (p *HsmTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HsmTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
