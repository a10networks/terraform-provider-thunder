

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTokenAuthStats struct {
    
    Stats DdosTokenAuthStatsStats `json:"stats"`

}
type DataDdosTokenAuthStats struct {
    DtDdosTokenAuthStats DdosTokenAuthStats `json:"token-auth"`
}


type DdosTokenAuthStatsStats struct {
    Token_authentication_matched int `json:"token_authentication_matched"`
    Token_authentication_mismatched int `json:"token_authentication_mismatched"`
    Token_authentication_invalid int `json:"token_authentication_invalid"`
}

func (p *DdosTokenAuthStats) GetId() string{
    return "1"
}

func (p *DdosTokenAuthStats) getPath() string{
    return "ddos/token-auth/stats"
}

func (p *DdosTokenAuthStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosTokenAuthStats,error) {
logger.Println("DdosTokenAuthStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosTokenAuthStats
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
