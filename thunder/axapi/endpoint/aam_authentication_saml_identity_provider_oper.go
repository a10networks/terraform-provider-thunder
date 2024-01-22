

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSamlIdentityProviderOper struct {
    
    Name string `json:"name"`
    Oper AamAuthenticationSamlIdentityProviderOperOper `json:"oper"`

}
type DataAamAuthenticationSamlIdentityProviderOper struct {
    DtAamAuthenticationSamlIdentityProviderOper AamAuthenticationSamlIdentityProviderOper `json:"identity-provider"`
}


type AamAuthenticationSamlIdentityProviderOperOper struct {
    Md string `json:"md"`
    Cert string `json:"cert"`
    EntityId string `json:"entity-id"`
    SsoList []AamAuthenticationSamlIdentityProviderOperOperSsoList `json:"sso-list"`
    SloList []AamAuthenticationSamlIdentityProviderOperOperSloList `json:"slo-list"`
    ArsList []AamAuthenticationSamlIdentityProviderOperOperArsList `json:"ars-list"`
    AqsList []AamAuthenticationSamlIdentityProviderOperOperAqsList `json:"aqs-list"`
}


type AamAuthenticationSamlIdentityProviderOperOperSsoList struct {
    SsoLocation string `json:"sso-location"`
    SsoBinding string `json:"sso-binding"`
}


type AamAuthenticationSamlIdentityProviderOperOperSloList struct {
    SloLocation string `json:"slo-location"`
    SloBinding string `json:"slo-binding"`
}


type AamAuthenticationSamlIdentityProviderOperOperArsList struct {
    ArsIndex int `json:"ars-index"`
    ArsLocation string `json:"ars-location"`
    ArsBinding string `json:"ars-binding"`
}


type AamAuthenticationSamlIdentityProviderOperOperAqsList struct {
    AqsLocation string `json:"aqs-location"`
    AqsBinding string `json:"aqs-binding"`
}

func (p *AamAuthenticationSamlIdentityProviderOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationSamlIdentityProviderOper) getPath() string{
    
    return "aam/authentication/saml/identity-provider/"+p.Name+"/oper"
}

func (p *AamAuthenticationSamlIdentityProviderOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationSamlIdentityProviderOper,error) {
logger.Println("AamAuthenticationSamlIdentityProviderOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationSamlIdentityProviderOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
