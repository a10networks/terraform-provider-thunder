

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerLdap struct {
	Inst struct {

    InstanceList []AamAuthenticationServerLdapInstanceList `json:"instance-list"`
    SamplingEnable []AamAuthenticationServerLdapSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"ldap"`
}


type AamAuthenticationServerLdapInstanceList struct {
    Name string `json:"name"`
    Host AamAuthenticationServerLdapInstanceListHost `json:"host"`
    Base string `json:"base"`
    Port int `json:"port" dval:"389"`
    PortHm string `json:"port-hm"`
    PortHmDisable int `json:"port-hm-disable"`
    Pwdmaxage int `json:"pwdmaxage"`
    AdminDn string `json:"admin-dn"`
    AdminSecret int `json:"admin-secret"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Timeout int `json:"timeout" dval:"10"`
    DnAttribute string `json:"dn-attribute" dval:"cn"`
    DefaultDomain string `json:"default-domain"`
    BindWithDn int `json:"bind-with-dn"`
    DeriveBindDn AamAuthenticationServerLdapInstanceListDeriveBindDn `json:"derive-bind-dn"`
    HealthCheck int `json:"health-check"`
    HealthCheckString string `json:"health-check-string"`
    HealthCheckDisable int `json:"health-check-disable"`
    Protocol string `json:"protocol" dval:"ldap"`
    CaCert string `json:"ca-cert"`
    LdapsConnReuseIdleTimeout int `json:"ldaps-conn-reuse-idle-timeout"`
    AuthType string `json:"auth-type"`
    PromptPwChangeBeforeExp int `json:"prompt-pw-change-before-exp"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerLdapInstanceListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServerLdapInstanceListHost struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerLdapInstanceListDeriveBindDn struct {
    UsernameAttr string `json:"username-attr"`
}


type AamAuthenticationServerLdapInstanceListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerLdapSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServerLdap) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerLdap) getPath() string{
    return "aam/authentication/server/ldap"
}

func (p *AamAuthenticationServerLdap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdap::Post")
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

func (p *AamAuthenticationServerLdap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdap::Get")
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
func (p *AamAuthenticationServerLdap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdap::Put")
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

func (p *AamAuthenticationServerLdap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerLdap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
