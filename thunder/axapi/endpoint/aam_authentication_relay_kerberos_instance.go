

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayKerberosInstance struct {
	Inst struct {

    Encrypted string `json:"encrypted"`
    KerberosAccount string `json:"kerberos-account"`
    KerberosKdc string `json:"kerberos-kdc"`
    KerberosKdcServiceGroup string `json:"kerberos-kdc-service-group"`
    KerberosRealm string `json:"kerberos-realm"`
    Name string `json:"name"`
    Password int `json:"password"`
    Port int `json:"port" dval:"88"`
    SamplingEnable []AamAuthenticationRelayKerberosInstanceSamplingEnable `json:"sampling-enable"`
    SecretString string `json:"secret-string"`
    Timeout int `json:"timeout" dval:"10"`
    Uuid string `json:"uuid"`

	} `json:"instance"`
}


type AamAuthenticationRelayKerberosInstanceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationRelayKerberosInstance) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationRelayKerberosInstance) getPath() string{
    return "aam/authentication/relay/kerberos/instance"
}

func (p *AamAuthenticationRelayKerberosInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberosInstance::Post")
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

func (p *AamAuthenticationRelayKerberosInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberosInstance::Get")
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
func (p *AamAuthenticationRelayKerberosInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberosInstance::Put")
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

func (p *AamAuthenticationRelayKerberosInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayKerberosInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
