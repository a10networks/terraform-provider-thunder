

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityReportingTemplateNotificationTemplateName struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    Authentication VisibilityReportingTemplateNotificationTemplateNameAuthentication3126 `json:"authentication"`
    DebugMode int `json:"debug-mode"`
    HostName string `json:"host-name"`
    HttpPort int `json:"http-port" dval:"80"`
    HttpsPort int `json:"https-port" dval:"443"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    Name string `json:"name"`
    Protocol string `json:"protocol" dval:"https"`
    RelativeUri string `json:"relative-uri" dval:"/"`
    SamplingEnable []VisibilityReportingTemplateNotificationTemplateNameSamplingEnable `json:"sampling-enable"`
    TestConnectivity int `json:"test-connectivity"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"template-name"`
}


type VisibilityReportingTemplateNotificationTemplateNameAuthentication3126 struct {
    RelativeLoginUri string `json:"relative-login-uri"`
    RelativeLogoffUri string `json:"relative-logoff-uri"`
    AuthUsername string `json:"auth-username"`
    AuthPassword int `json:"auth-password"`
    AuthPasswordString string `json:"auth-password-string"`
    Encrypted string `json:"encrypted"`
    ApiKey int `json:"api-key"`
    ApiKeyString string `json:"api-key-string"`
    ApiKeyEncrypted string `json:"api-key-encrypted"`
    Uuid string `json:"uuid"`
}


type VisibilityReportingTemplateNotificationTemplateNameSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *VisibilityReportingTemplateNotificationTemplateName) GetId() string{
    return p.Inst.Name
}

func (p *VisibilityReportingTemplateNotificationTemplateName) getPath() string{
    return "visibility/reporting/template/notification/template-name"
}

func (p *VisibilityReportingTemplateNotificationTemplateName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateName::Post")
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

func (p *VisibilityReportingTemplateNotificationTemplateName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateName::Get")
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
func (p *VisibilityReportingTemplateNotificationTemplateName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateName::Put")
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

func (p *VisibilityReportingTemplateNotificationTemplateName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReportingTemplateNotificationTemplateName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
