

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelaySamlStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationRelaySamlStatsStats `json:"stats"`

}
type DataAamAuthenticationRelaySamlStats struct {
    DtAamAuthenticationRelaySamlStats AamAuthenticationRelaySamlStats `json:"saml"`
}


type AamAuthenticationRelaySamlStatsStats struct {
    Request int `json:"request"`
    Success int `json:"success"`
    Failure int `json:"failure"`
    Error int `json:"error"`
}

func (p *AamAuthenticationRelaySamlStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelaySamlStats) getPath() string{
    
    return "aam/authentication/relay/saml/"+p.Name+"/stats"
}

func (p *AamAuthenticationRelaySamlStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelaySamlStats,error) {
logger.Println("AamAuthenticationRelaySamlStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelaySamlStats
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
