

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamJwtAuthorizationStats struct {
    
    Name string `json:"name"`
    Stats AamJwtAuthorizationStatsStats `json:"stats"`

}
type DataAamJwtAuthorizationStats struct {
    DtAamJwtAuthorizationStats AamJwtAuthorizationStats `json:"jwt-authorization"`
}


type AamJwtAuthorizationStatsStats struct {
    JwtRequest int `json:"jwt-request"`
    JwtAuthorizeSuccess int `json:"jwt-authorize-success"`
    JwtAuthorizeFailure int `json:"jwt-authorize-failure"`
    JwtMissingToken int `json:"jwt-missing-token"`
    JwtMissingClaim int `json:"jwt-missing-claim"`
    JwtTokenExpired int `json:"jwt-token-expired"`
    JwtSignatureFailure int `json:"jwt-signature-failure"`
    JwtOtherError int `json:"jwt-other-error"`
}

func (p *AamJwtAuthorizationStats) GetId() string{
    return "1"
}

func (p *AamJwtAuthorizationStats) getPath() string{
    
    return "aam/jwt-authorization/"+p.Name+"/stats"
}

func (p *AamJwtAuthorizationStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamJwtAuthorizationStats,error) {
logger.Println("AamJwtAuthorizationStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamJwtAuthorizationStats
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
