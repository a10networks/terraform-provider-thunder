

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodHttp struct {
	Inst struct {

    Http int `json:"http"`
    HttpEncrypted string `json:"http-encrypted"`
    HttpExpect int `json:"http-expect"`
    HttpHost string `json:"http-host"`
    HttpKerberosAuth int `json:"http-kerberos-auth"`
    HttpKerberosKdc HealthMonitorMethodHttpHttpKerberosKdc `json:"http-kerberos-kdc"`
    HttpKerberosRealm string `json:"http-kerberos-realm"`
    HttpMaintenanceCode string `json:"http-maintenance-code"`
    HttpPassword int `json:"http-password"`
    HttpPasswordString string `json:"http-password-string"`
    HttpPort int `json:"http-port" dval:"80"`
    HttpPostdata string `json:"http-postdata"`
    HttpPostfile string `json:"http-postfile"`
    HttpResponseCode string `json:"http-response-code"`
    HttpText string `json:"http-text"`
    HttpUrl int `json:"http-url"`
    HttpUsername string `json:"http-username"`
    Maintenance int `json:"maintenance"`
    MaintenanceText string `json:"maintenance-text"`
    MaintenanceTextRegex string `json:"maintenance-text-regex"`
    PostPath string `json:"post-path"`
    PostType string `json:"post-type"`
    ResponseCodeRegex string `json:"response-code-regex"`
    TextRegex string `json:"text-regex"`
    UrlPath string `json:"url-path"`
    UrlType string `json:"url-type"`
    Uuid string `json:"uuid"`
    Version2 int `json:"version2"`
    Name string 

	} `json:"http"`
}


type HealthMonitorMethodHttpHttpKerberosKdc struct {
    HttpKerberosHostip string `json:"http-kerberos-hostip"`
    HttpKerberosHostipv6 string `json:"http-kerberos-hostipv6"`
    HttpKerberosPort int `json:"http-kerberos-port"`
    HttpKerberosPortv6 int `json:"http-kerberos-portv6"`
}

func (p *HealthMonitorMethodHttp) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodHttp) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/http"
}

func (p *HealthMonitorMethodHttp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttp::Post")
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

func (p *HealthMonitorMethodHttp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttp::Get")
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
func (p *HealthMonitorMethodHttp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttp::Put")
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

func (p *HealthMonitorMethodHttp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
