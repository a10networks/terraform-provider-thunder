

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationJwksOper struct {
    
    Oper AamAuthenticationJwksOperOper `json:"oper"`

}
type DataAamAuthenticationJwksOper struct {
    DtAamAuthenticationJwksOper AamAuthenticationJwksOper `json:"jwks"`
}


type AamAuthenticationJwksOperOper struct {
    JwkList []AamAuthenticationJwksOperOperJwkList `json:"jwk-list"`
}


type AamAuthenticationJwksOperOperJwkList struct {
    JwkName string `json:"jwk-name"`
    JwkSize int `json:"jwk-size"`
}

func (p *AamAuthenticationJwksOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationJwksOper) getPath() string{
    return "aam/authentication/jwks/oper"
}

func (p *AamAuthenticationJwksOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationJwksOper,error) {
logger.Println("AamAuthenticationJwksOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationJwksOper
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
