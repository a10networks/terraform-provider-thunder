

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSamlSessionOper struct {
    
    Oper AamAuthenticationSamlSessionOperOper `json:"oper"`

}
type DataAamAuthenticationSamlSessionOper struct {
    DtAamAuthenticationSamlSessionOper AamAuthenticationSamlSessionOper `json:"session"`
}


type AamAuthenticationSamlSessionOperOper struct {
    SpList []AamAuthenticationSamlSessionOperOperSpList `json:"sp-list"`
    Name string `json:"name"`
}


type AamAuthenticationSamlSessionOperOperSpList struct {
    SpId string `json:"sp-id"`
    SessionCount int `json:"session-count"`
    SessionList []AamAuthenticationSamlSessionOperOperSpListSessionList `json:"session-list"`
}


type AamAuthenticationSamlSessionOperOperSpListSessionList struct {
    Nameid string `json:"nameid"`
    ClientAddr string `json:"client-addr"`
    IdProvider string `json:"id-provider"`
    AuthInstant string `json:"auth-instant"`
    ExpireTime string `json:"expire-time"`
}

func (p *AamAuthenticationSamlSessionOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationSamlSessionOper) getPath() string{
    return "aam/authentication/saml/session/oper"
}

func (p *AamAuthenticationSamlSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationSamlSessionOper,error) {
logger.Println("AamAuthenticationSamlSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationSamlSessionOper
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
