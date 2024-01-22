

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodHttps struct {
	Inst struct {

    Cert string `json:"cert"`
    CertKeyShared int `json:"cert-key-shared"`
    DisableSslv2hello int `json:"disable-sslv2hello"`
    HttpVersion string `json:"http-version"`
    Https int `json:"https"`
    HttpsEncrypted string `json:"https-encrypted"`
    HttpsExpect int `json:"https-expect"`
    HttpsHost string `json:"https-host"`
    HttpsKerberosAuth int `json:"https-kerberos-auth"`
    HttpsKerberosKdc HealthMonitorMethodHttpsHttpsKerberosKdc `json:"https-kerberos-kdc"`
    HttpsKerberosRealm string `json:"https-kerberos-realm"`
    HttpsKeyEncrypted string `json:"https-key-encrypted"`
    HttpsMaintenanceCode string `json:"https-maintenance-code"`
    HttpsPassword int `json:"https-password"`
    HttpsPasswordString string `json:"https-password-string"`
    HttpsPostdata string `json:"https-postdata"`
    HttpsPostfile string `json:"https-postfile"`
    HttpsResponseCode string `json:"https-response-code"`
    HttpsServerCertName string `json:"https-server-cert-name"`
    HttpsText string `json:"https-text"`
    HttpsUrl int `json:"https-url"`
    HttpsUsername string `json:"https-username"`
    Key string `json:"key"`
    KeyPassPhrase int `json:"key-pass-phrase"`
    KeyPhrase string `json:"key-phrase"`
    Maintenance int `json:"maintenance"`
    MaintenanceText string `json:"maintenance-text"`
    MaintenanceTextRegex string `json:"maintenance-text-regex"`
    PostPath string `json:"post-path"`
    PostType string `json:"post-type"`
    ResponseCodeRegex string `json:"response-code-regex"`
    Sni int `json:"sni"`
    TextRegex string `json:"text-regex"`
    UrlPath string `json:"url-path"`
    UrlType string `json:"url-type"`
    Uuid string `json:"uuid"`
    WebPort int `json:"web-port" dval:"443"`
    Name string 

	} `json:"https"`
}


type HealthMonitorMethodHttpsHttpsKerberosKdc struct {
    HttpsKerberosHostip string `json:"https-kerberos-hostip"`
    HttpsKerberosHostipv6 string `json:"https-kerberos-hostipv6"`
    HttpsKerberosPort int `json:"https-kerberos-port"`
    HttpsKerberosPortv6 int `json:"https-kerberos-portv6"`
}

func (p *HealthMonitorMethodHttps) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodHttps) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/https"
}

func (p *HealthMonitorMethodHttps) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttps::Post")
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

func (p *HealthMonitorMethodHttps) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttps::Get")
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
func (p *HealthMonitorMethodHttps) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttps::Put")
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

func (p *HealthMonitorMethodHttps) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodHttps::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
