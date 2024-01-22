

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationOauthGlobalStats struct {
    
    Stats AamAuthenticationOauthGlobalStatsStats `json:"stats"`

}
type DataAamAuthenticationOauthGlobalStats struct {
    DtAamAuthenticationOauthGlobalStats AamAuthenticationOauthGlobalStats `json:"global"`
}


type AamAuthenticationOauthGlobalStatsStats struct {
    AuthReq int `json:"auth-req"`
    AuthSucc int `json:"auth-succ"`
    AuthFail int `json:"auth-fail"`
    AuthError int `json:"auth-error"`
    RelayReq int `json:"relay-req"`
    RelaySucc int `json:"relay-succ"`
    RelayFail int `json:"relay-fail"`
    OtherError int `json:"other-error"`
}

func (p *AamAuthenticationOauthGlobalStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationOauthGlobalStats) getPath() string{
    return "aam/authentication/oauth/global/stats"
}

func (p *AamAuthenticationOauthGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationOauthGlobalStats,error) {
logger.Println("AamAuthenticationOauthGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationOauthGlobalStats
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
