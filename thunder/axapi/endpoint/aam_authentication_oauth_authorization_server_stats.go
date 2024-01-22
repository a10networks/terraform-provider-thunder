

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationOauthAuthorizationServerStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationOauthAuthorizationServerStatsStats `json:"stats"`

}
type DataAamAuthenticationOauthAuthorizationServerStats struct {
    DtAamAuthenticationOauthAuthorizationServerStats AamAuthenticationOauthAuthorizationServerStats `json:"authorization-server"`
}


type AamAuthenticationOauthAuthorizationServerStatsStats struct {
    AuthReq int `json:"auth-req"`
    AuthSucc int `json:"auth-succ"`
    AuthFail int `json:"auth-fail"`
    AuthError int `json:"auth-error"`
    OtherError int `json:"other-error"`
}

func (p *AamAuthenticationOauthAuthorizationServerStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationOauthAuthorizationServerStats) getPath() string{
    
    return "aam/authentication/oauth/authorization-server/"+p.Name+"/stats"
}

func (p *AamAuthenticationOauthAuthorizationServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationOauthAuthorizationServerStats,error) {
logger.Println("AamAuthenticationOauthAuthorizationServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationOauthAuthorizationServerStats
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
