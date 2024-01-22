

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServer struct {
	Inst struct {

    Ldap AamAuthenticationServerLdap24 `json:"ldap"`
    Ocsp AamAuthenticationServerOcsp30 `json:"ocsp"`
    Radius AamAuthenticationServerRadius34 `json:"radius"`
    Uuid string `json:"uuid"`
    Windows AamAuthenticationServerWindows39 `json:"windows"`

	} `json:"server"`
}


type AamAuthenticationServerLdap24 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerLdapSamplingEnable25 `json:"sampling-enable"`
    InstanceList []AamAuthenticationServerLdapInstanceList26 `json:"instance-list"`
}


type AamAuthenticationServerLdapSamplingEnable25 struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerLdapInstanceList26 struct {
    Name string `json:"name"`
    Host AamAuthenticationServerLdapInstanceListHost27 `json:"host"`
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
    DeriveBindDn AamAuthenticationServerLdapInstanceListDeriveBindDn28 `json:"derive-bind-dn"`
    HealthCheck int `json:"health-check"`
    HealthCheckString string `json:"health-check-string"`
    HealthCheckDisable int `json:"health-check-disable"`
    Protocol string `json:"protocol" dval:"ldap"`
    CaCert string `json:"ca-cert"`
    LdapsConnReuseIdleTimeout int `json:"ldaps-conn-reuse-idle-timeout"`
    AuthType string `json:"auth-type"`
    PromptPwChangeBeforeExp int `json:"prompt-pw-change-before-exp"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerLdapInstanceListSamplingEnable29 `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServerLdapInstanceListHost27 struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerLdapInstanceListDeriveBindDn28 struct {
    UsernameAttr string `json:"username-attr"`
}


type AamAuthenticationServerLdapInstanceListSamplingEnable29 struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerOcsp30 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerOcspSamplingEnable31 `json:"sampling-enable"`
    InstanceList []AamAuthenticationServerOcspInstanceList32 `json:"instance-list"`
}


type AamAuthenticationServerOcspSamplingEnable31 struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerOcspInstanceList32 struct {
    Name string `json:"name"`
    Url string `json:"url"`
    ResponderCa string `json:"responder-ca"`
    ResponderCert string `json:"responder-cert"`
    HealthCheck int `json:"health-check"`
    HealthCheckString string `json:"health-check-string"`
    HealthCheckDisable int `json:"health-check-disable"`
    PortHealthCheck string `json:"port-health-check"`
    PortHealthCheckDisable int `json:"port-health-check-disable"`
    HttpVersion int `json:"http-version"`
    VersionType string `json:"version-type"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerOcspInstanceListSamplingEnable33 `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServerOcspInstanceListSamplingEnable33 struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerRadius34 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerRadiusSamplingEnable35 `json:"sampling-enable"`
    InstanceList []AamAuthenticationServerRadiusInstanceList36 `json:"instance-list"`
}


type AamAuthenticationServerRadiusSamplingEnable35 struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerRadiusInstanceList36 struct {
    Name string `json:"name"`
    Host AamAuthenticationServerRadiusInstanceListHost37 `json:"host"`
    Secret int `json:"secret"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Port int `json:"port" dval:"1812"`
    PortHm string `json:"port-hm"`
    PortHmDisable int `json:"port-hm-disable"`
    Interval int `json:"interval" dval:"3"`
    Retry int `json:"retry" dval:"5"`
    HealthCheck int `json:"health-check"`
    HealthCheckString string `json:"health-check-string"`
    HealthCheckDisable int `json:"health-check-disable"`
    AccountingPort int `json:"accounting-port" dval:"1813"`
    AcctPortHm string `json:"acct-port-hm"`
    AcctPortHmDisable int `json:"acct-port-hm-disable"`
    AuthType string `json:"auth-type"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerRadiusInstanceListSamplingEnable38 `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServerRadiusInstanceListHost37 struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerRadiusInstanceListSamplingEnable38 struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerWindows39 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerWindowsSamplingEnable `json:"sampling-enable"`
    InstanceList []AamAuthenticationServerWindowsInstanceList `json:"instance-list"`
}


type AamAuthenticationServerWindowsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationServerWindowsInstanceList struct {
    Name string `json:"name"`
    Host AamAuthenticationServerWindowsInstanceListHost `json:"host"`
    Timeout int `json:"timeout" dval:"10"`
    AuthProtocol AamAuthenticationServerWindowsInstanceListAuthProtocol `json:"auth-protocol"`
    Realm string `json:"realm"`
    SupportApachedsKdc int `json:"support-apacheds-kdc"`
    HealthCheck int `json:"health-check"`
    HealthCheckString string `json:"health-check-string"`
    HealthCheckDisable int `json:"health-check-disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationServerWindowsInstanceListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type AamAuthenticationServerWindowsInstanceListHost struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerWindowsInstanceListAuthProtocol struct {
    NtlmDisable int `json:"ntlm-disable"`
    NtlmVersion int `json:"ntlm-version" dval:"2"`
    NtlmHealthCheck string `json:"ntlm-health-check"`
    NtlmHealthCheckDisable int `json:"ntlm-health-check-disable"`
    KerberosDisable int `json:"kerberos-disable"`
    KerberosPort int `json:"kerberos-port" dval:"88"`
    KportHm string `json:"kport-hm"`
    KportHmDisable int `json:"kport-hm-disable"`
    KerberosPasswordChangePort int `json:"kerberos-password-change-port" dval:"464"`
    KdcValidate int `json:"kdc-validate"`
    KerberosKdcValidation AamAuthenticationServerWindowsInstanceListAuthProtocolKerberosKdcValidation `json:"kerberos-kdc-validation"`
}


type AamAuthenticationServerWindowsInstanceListAuthProtocolKerberosKdcValidation struct {
    KdcSpn string `json:"kdc-spn"`
    KdcAccount string `json:"kdc-account"`
    KdcPassword int `json:"kdc-password"`
    KdcPwd string `json:"kdc-pwd"`
    Encrypted string `json:"encrypted"`
}


type AamAuthenticationServerWindowsInstanceListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServer) GetId() string{
    return "1"
}

func (p *AamAuthenticationServer) getPath() string{
    return "aam/authentication/server"
}

func (p *AamAuthenticationServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServer::Post")
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

func (p *AamAuthenticationServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServer::Get")
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
func (p *AamAuthenticationServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServer::Put")
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

func (p *AamAuthenticationServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
