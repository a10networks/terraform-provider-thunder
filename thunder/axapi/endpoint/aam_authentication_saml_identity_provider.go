

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSamlIdentityProvider struct {
	Inst struct {

    Metadata string `json:"metadata"`
    Name string `json:"name"`
    ReloadInterval int `json:"reload-interval" dval:"28800"`
    ReloadMetadata int `json:"reload-metadata"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"identity-provider"`
}

func (p *AamAuthenticationSamlIdentityProvider) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationSamlIdentityProvider) getPath() string{
    return "aam/authentication/saml/identity-provider"
}

func (p *AamAuthenticationSamlIdentityProvider) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlIdentityProvider::Post")
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

func (p *AamAuthenticationSamlIdentityProvider) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlIdentityProvider::Get")
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
func (p *AamAuthenticationSamlIdentityProvider) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlIdentityProvider::Put")
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

func (p *AamAuthenticationSamlIdentityProvider) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlIdentityProvider::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
