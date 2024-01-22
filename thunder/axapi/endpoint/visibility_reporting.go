

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityReporting struct {
	Inst struct {

    SamplingEnable []VisibilityReportingSamplingEnable `json:"sampling-enable"`
    SessionLogging string `json:"session-logging" dval:"disable"`
    TelemetryExportInterval VisibilityReportingTelemetryExportInterval3127 `json:"telemetry-export-interval"`
    Template VisibilityReportingTemplate3128 `json:"template"`
    Uuid string `json:"uuid"`

	} `json:"reporting"`
}


type VisibilityReportingSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VisibilityReportingTelemetryExportInterval3127 struct {
    Value int `json:"value" dval:"5"`
    Uuid string `json:"uuid"`
}


type VisibilityReportingTemplate3128 struct {
    Notification VisibilityReportingTemplateNotification3129 `json:"notification"`
}


type VisibilityReportingTemplateNotification3129 struct {
    TemplateNameList []VisibilityReportingTemplateNotificationTemplateNameList `json:"template-name-list"`
    Debug VisibilityReportingTemplateNotificationDebug3130 `json:"debug"`
}


type VisibilityReportingTemplateNotificationTemplateNameList struct {
    Name string `json:"name"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    HostName string `json:"host-name"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Protocol string `json:"protocol" dval:"https"`
    HttpPort int `json:"http-port" dval:"80"`
    HttpsPort int `json:"https-port" dval:"443"`
    RelativeUri string `json:"relative-uri" dval:"/"`
    Action string `json:"action" dval:"enable"`
    DebugMode int `json:"debug-mode"`
    TestConnectivity int `json:"test-connectivity"`
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable `json:"sampling-enable"`
    Authentication VisibilityReportingTemplateNotificationTemplateNameListAuthentication `json:"authentication"`
}


type VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VisibilityReportingTemplateNotificationTemplateNameListAuthentication struct {
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


type VisibilityReportingTemplateNotificationDebug3130 struct {
    Uuid string `json:"uuid"`
}

func (p *VisibilityReporting) GetId() string{
    return "1"
}

func (p *VisibilityReporting) getPath() string{
    return "visibility/reporting"
}

func (p *VisibilityReporting) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReporting::Post")
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

func (p *VisibilityReporting) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReporting::Get")
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
func (p *VisibilityReporting) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReporting::Put")
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

func (p *VisibilityReporting) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityReporting::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
