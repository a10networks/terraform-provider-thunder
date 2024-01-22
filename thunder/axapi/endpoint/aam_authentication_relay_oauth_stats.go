

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayOauthStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationRelayOauthStatsStats `json:"stats"`

}
type DataAamAuthenticationRelayOauthStats struct {
    DtAamAuthenticationRelayOauthStats AamAuthenticationRelayOauthStats `json:"oauth"`
}


type AamAuthenticationRelayOauthStatsStats struct {
    RelayReq int `json:"relay-req"`
    RelaySucc int `json:"relay-succ"`
    RelayFail int `json:"relay-fail"`
}

func (p *AamAuthenticationRelayOauthStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayOauthStats) getPath() string{
    
    return "aam/authentication/relay/oauth/"+p.Name+"/stats"
}

func (p *AamAuthenticationRelayOauthStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayOauthStats,error) {
logger.Println("AamAuthenticationRelayOauthStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayOauthStats
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
