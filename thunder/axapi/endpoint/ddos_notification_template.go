

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosNotificationTemplate struct {
	Inst struct {

    Api DdosNotificationTemplateApi285 `json:"api"`
    DebugMode int `json:"debug-mode"`
    Disable int `json:"disable"`
    Name string `json:"name"`
    TestConnectivity int `json:"test-connectivity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Verbose int `json:"verbose"`

	} `json:"notification-template"`
}


type DdosNotificationTemplateApi285 struct {
    HostIpv4Address string `json:"host-ipv4-address"`
    HostIpv6Address string `json:"host-ipv6-address"`
    Hostname string `json:"hostname"`
    HttpProtocol string `json:"http-protocol" dval:"https"`
    HttpPort int `json:"http-port" dval:"80"`
    HttpsPort int `json:"https-port" dval:"443"`
    Timeout int `json:"timeout" dval:"10"`
    RelativeUri string `json:"relative-uri" dval:"/"`
    DisableAuthentication int `json:"disable-authentication"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`
    Authentication DdosNotificationTemplateApiAuthentication286 `json:"authentication"`
}


type DdosNotificationTemplateApiAuthentication286 struct {
    RelativeLoginUri string `json:"relative-login-uri"`
    RelativeLogoffUri string `json:"relative-logoff-uri"`
    AuthUsername string `json:"auth-username"`
    AuthPassword int `json:"auth-password"`
    AuthPasswordVal string `json:"auth-password-val"`
    Encrypted string `json:"encrypted"`
    ApiKey int `json:"api-key"`
    ApiKeyString string `json:"api-key-string"`
    ApiKeyEncrypted string `json:"api-key-encrypted"`
    Uuid string `json:"uuid"`
}

func (p *DdosNotificationTemplate) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosNotificationTemplate) getPath() string{
    return "ddos/notification-template"
}

func (p *DdosNotificationTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplate::Post")
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

func (p *DdosNotificationTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplate::Get")
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
func (p *DdosNotificationTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplate::Put")
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

func (p *DdosNotificationTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
