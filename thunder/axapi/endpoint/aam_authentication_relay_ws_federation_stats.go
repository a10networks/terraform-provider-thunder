

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayWsFederationStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationRelayWsFederationStatsStats `json:"stats"`

}
type DataAamAuthenticationRelayWsFederationStats struct {
    DtAamAuthenticationRelayWsFederationStats AamAuthenticationRelayWsFederationStats `json:"ws-federation"`
}


type AamAuthenticationRelayWsFederationStatsStats struct {
    Request int `json:"request"`
    Success int `json:"success"`
    Failure int `json:"failure"`
}

func (p *AamAuthenticationRelayWsFederationStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayWsFederationStats) getPath() string{
    
    return "aam/authentication/relay/ws-federation/"+p.Name+"/stats"
}

func (p *AamAuthenticationRelayWsFederationStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayWsFederationStats,error) {
logger.Println("AamAuthenticationRelayWsFederationStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayWsFederationStats
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
