

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayKerberos struct {
	Inst struct {

    InstanceList []AamAuthenticationRelayKerberosInstanceList `json:"instance-list"`
    SamplingEnable []AamAuthenticationRelayKerberosSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"kerberos"`
}


type AamAuthenticationRelayKerberosInstanceList struct {
    Name string `json:"name"`
    KerberosRealm string `json:"kerberos-realm"`
    KerberosKdc string `json:"kerberos-kdc"`
    KerberosKdcServiceGroup string `json:"kerberos-kdc-service-group"`
    KerberosAccount string `json:"kerberos-account"`
    Password int `json:"password"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Port int `json:"port" dval:"88"`
    Timeout int `json:"timeout" dval:"10"`
    Uuid string `json:"uuid"`
    SamplingEnable []AamAuthenticationRelayKerberosInstanceListSamplingEnable `json:"sampling-enable"`
}


type AamAuthenticationRelayKerberosInstanceListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAuthenticationRelayKerberosSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationRelayKerberos) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayKerberos) getPath() string{
    return "aam/authentication/relay/kerberos"
}

func (p *AamAuthenticationRelayKerberos) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberos::Post")
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

func (p *AamAuthenticationRelayKerberos) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberos::Get")
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
func (p *AamAuthenticationRelayKerberos) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberos::Put")
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

func (p *AamAuthenticationRelayKerberos) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberos::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
