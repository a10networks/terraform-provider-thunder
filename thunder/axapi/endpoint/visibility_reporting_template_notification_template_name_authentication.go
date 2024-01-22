

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityReportingTemplateNotificationTemplateNameAuthentication struct {
	Inst struct {

    ApiKey int `json:"api-key"`
    ApiKeyEncrypted string `json:"api-key-encrypted"`
    ApiKeyString string `json:"api-key-string"`
    AuthPassword int `json:"auth-password"`
    AuthPasswordString string `json:"auth-password-string"`
    AuthUsername string `json:"auth-username"`
    Encrypted string `json:"encrypted"`
    RelativeLoginUri string `json:"relative-login-uri"`
    RelativeLogoffUri string `json:"relative-logoff-uri"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"authentication"`
}

func (p *VisibilityReportingTemplateNotificationTemplateNameAuthentication) GetId() string{
    return "1"
}

func (p *VisibilityReportingTemplateNotificationTemplateNameAuthentication) getPath() string{
    return "visibility/reporting/template/notification/template-name/" +p.Inst.Name + "/authentication"
}

func (p *VisibilityReportingTemplateNotificationTemplateNameAuthentication) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateNameAuthentication::Post")
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

func (p *VisibilityReportingTemplateNotificationTemplateNameAuthentication) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateNameAuthentication::Get")
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
func (p *VisibilityReportingTemplateNotificationTemplateNameAuthentication) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateNameAuthentication::Put")
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

func (p *VisibilityReportingTemplateNotificationTemplateNameAuthentication) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateNameAuthentication::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
