

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationAccountKerberosSpn struct {
	Inst struct {

    Account string `json:"account"`
    Encrypted string `json:"encrypted"`
    Name string `json:"name"`
    Password int `json:"password"`
    Realm string `json:"realm"`
    SecretString string `json:"secret-string"`
    ServicePrincipalName string `json:"service-principal-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"kerberos-spn"`
}

func (p *AamAuthenticationAccountKerberosSpn) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationAccountKerberosSpn) getPath() string{
    return "aam/authentication/account/kerberos-spn"
}

func (p *AamAuthenticationAccountKerberosSpn) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccountKerberosSpn::Post")
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

func (p *AamAuthenticationAccountKerberosSpn) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccountKerberosSpn::Get")
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
func (p *AamAuthenticationAccountKerberosSpn) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccountKerberosSpn::Put")
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

func (p *AamAuthenticationAccountKerberosSpn) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccountKerberosSpn::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
