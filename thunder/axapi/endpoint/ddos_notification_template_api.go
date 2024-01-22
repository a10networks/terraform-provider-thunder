

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosNotificationTemplateApi struct {
	Inst struct {

    Authentication DdosNotificationTemplateApiAuthentication284 `json:"authentication"`
    DisableAuthentication int `json:"disable-authentication"`
    HostIpv4Address string `json:"host-ipv4-address"`
    HostIpv6Address string `json:"host-ipv6-address"`
    Hostname string `json:"hostname"`
    HttpPort int `json:"http-port" dval:"80"`
    HttpProtocol string `json:"http-protocol" dval:"https"`
    HttpsPort int `json:"https-port" dval:"443"`
    RelativeUri string `json:"relative-uri" dval:"/"`
    Timeout int `json:"timeout" dval:"10"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"api"`
}


type DdosNotificationTemplateApiAuthentication284 struct {
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

func (p *DdosNotificationTemplateApi) GetId() string{
    return "1"
}

func (p *DdosNotificationTemplateApi) getPath() string{
    return "ddos/notification-template/" +p.Inst.Name + "/api"
}

func (p *DdosNotificationTemplateApi) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateApi::Post")
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

func (p *DdosNotificationTemplateApi) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateApi::Get")
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
func (p *DdosNotificationTemplateApi) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateApi::Put")
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

func (p *DdosNotificationTemplateApi) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNotificationTemplateApi::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
