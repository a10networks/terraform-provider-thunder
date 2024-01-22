

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerWindowsInstance struct {
	Inst struct {

    AuthProtocol AamAuthenticationServerWindowsInstanceAuthProtocol `json:"auth-protocol"`
    HealthCheck int `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    HealthCheckString string `json:"health-check-string"`
    Host AamAuthenticationServerWindowsInstanceHost `json:"host"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Realm string `json:"realm"`
    SamplingEnable []AamAuthenticationServerWindowsInstanceSamplingEnable `json:"sampling-enable"`
    SupportApachedsKdc int `json:"support-apacheds-kdc"`
    Timeout int `json:"timeout" dval:"10"`
    Uuid string `json:"uuid"`

	} `json:"instance"`
}


type AamAuthenticationServerWindowsInstanceAuthProtocol struct {
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
    KerberosKdcValidation AamAuthenticationServerWindowsInstanceAuthProtocolKerberosKdcValidation `json:"kerberos-kdc-validation"`
}


type AamAuthenticationServerWindowsInstanceAuthProtocolKerberosKdcValidation struct {
    KdcSpn string `json:"kdc-spn"`
    KdcAccount string `json:"kdc-account"`
    KdcPassword int `json:"kdc-password"`
    KdcPwd string `json:"kdc-pwd"`
    Encrypted string `json:"encrypted"`
}


type AamAuthenticationServerWindowsInstanceHost struct {
    Hostip string `json:"hostip"`
    Hostipv6 string `json:"hostipv6"`
}


type AamAuthenticationServerWindowsInstanceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationServerWindowsInstance) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationServerWindowsInstance) getPath() string{
    return "aam/authentication/server/windows/instance"
}

func (p *AamAuthenticationServerWindowsInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerWindowsInstance::Post")
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

func (p *AamAuthenticationServerWindowsInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerWindowsInstance::Get")
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
func (p *AamAuthenticationServerWindowsInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerWindowsInstance::Put")
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

func (p *AamAuthenticationServerWindowsInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationServerWindowsInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
