

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationAccount struct {
	Inst struct {

    KerberosSpnList []AamAuthenticationAccountKerberosSpnList `json:"kerberos-spn-list"`
    SamplingEnable []AamAuthenticationAccountSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"account"`
}


type AamAuthenticationAccountKerberosSpnList struct {
    Name string `json:"name"`
    Realm string `json:"realm"`
    Account string `json:"account"`
    ServicePrincipalName string `json:"service-principal-name"`
    Password int `json:"password"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type AamAuthenticationAccountSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationAccount) GetId() string{
    return "1"
}

func (p *AamAuthenticationAccount) getPath() string{
    return "aam/authentication/account"
}

func (p *AamAuthenticationAccount) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccount::Post")
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

func (p *AamAuthenticationAccount) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccount::Get")
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
func (p *AamAuthenticationAccount) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccount::Put")
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

func (p *AamAuthenticationAccount) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationAccount::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
