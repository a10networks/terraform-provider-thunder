

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerLdapInstance struct {
	Inst struct {

    AdminDn string `json:"admin-dn"`
    AdminSecret int `json:"admin-secret"`
    AuthType string `json:"auth-type"`
    Base string `json:"base"`
    BindWithDn int `json:"bind-with-dn"`
    CaCert string `json:"ca-cert"`
    DefaultDomain string `json:"default-domain"`
    DeriveBindDn AamAuthenticationServerLdapInstanceDeriveBindDn `json:"derive-bind-dn"`
    DnAttribute string `json:"dn-attribute" dval:"cn"`
    Encrypted string `json:"encrypted"`
    HealthCheck int `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckString string `json:"health-check-string"`
    Host AamAuthenticationServerLdapInstanceHost `json:"host"`
    LdapsConnReuseIdleTimeout int `json:"ldaps-conn-reuse-idle-timeout"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Port int `json:"port" dval:"389"`
    PortHm string `json:"port-hm"`
    PortHmDisable int `json:"port-hm-disable"`
    PromptPwChangeBeforeExp int `json:"prompt-pw-change-before-exp"`
    Protocol string `json:"protocol" dval:"ldap"`
    Pwdmaxage int `json:"pwdmaxage"`
    SamplingEnable []AamAuthenticationServerLdapInstanceSamplingEnable `json:"sampling-enable"`
    SecretString string `json:"secret-string"`
    Timeout int `json:"timeout" dval:"10"`
    Uuid string `json:"uuid"`

	} `json:"instance"`
}


type AamAuthenticationServerLdapInstanceDeriveBindDn struct {
    UsernameAttr string `json:"username-attr"`
}


type AamAuthenticationServerLdapInstanceHost struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerLdapInstanceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServerLdapInstance) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationServerLdapInstance) getPath() string{
    return "aam/authentication/server/ldap/instance"
}

func (p *AamAuthenticationServerLdapInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdapInstance::Post")
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

func (p *AamAuthenticationServerLdapInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdapInstance::Get")
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
func (p *AamAuthenticationServerLdapInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdapInstance::Put")
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

func (p *AamAuthenticationServerLdapInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdapInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
