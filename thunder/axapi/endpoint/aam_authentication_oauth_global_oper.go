

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationOauthGlobalOper struct {
    
    Oper AamAuthenticationOauthGlobalOperOper `json:"oper"`

}
type DataAamAuthenticationOauthGlobalOper struct {
    DtAamAuthenticationOauthGlobalOper AamAuthenticationOauthGlobalOper `json:"global"`
}


type AamAuthenticationOauthGlobalOperOper struct {
    StatsClearType string `json:"stats-clear-type"`
}

func (p *AamAuthenticationOauthGlobalOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationOauthGlobalOper) getPath() string{
    return "aam/authentication/oauth/global/oper"
}

func (p *AamAuthenticationOauthGlobalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationOauthGlobalOper,error) {
logger.Println("AamAuthenticationOauthGlobalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationOauthGlobalOper
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
