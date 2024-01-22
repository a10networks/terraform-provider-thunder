

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayHttpBasicInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationRelayHttpBasicInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationRelayHttpBasicInstanceStats struct {
    DtAamAuthenticationRelayHttpBasicInstanceStats AamAuthenticationRelayHttpBasicInstanceStats `json:"instance"`
}


type AamAuthenticationRelayHttpBasicInstanceStatsStats struct {
    Success int `json:"success"`
    NoCreds int `json:"no-creds"`
    BadReq int `json:"bad-req"`
    Unauth int `json:"unauth"`
    Forbidden int `json:"forbidden"`
    NotFound int `json:"not-found"`
    ServerError int `json:"server-error"`
    Unavailable int `json:"unavailable"`
}

func (p *AamAuthenticationRelayHttpBasicInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayHttpBasicInstanceStats) getPath() string{
    
    return "aam/authentication/relay/http-basic/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationRelayHttpBasicInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayHttpBasicInstanceStats,error) {
logger.Println("AamAuthenticationRelayHttpBasicInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayHttpBasicInstanceStats
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
