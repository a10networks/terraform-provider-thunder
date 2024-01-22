

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerRadiusInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationServerRadiusInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationServerRadiusInstanceStats struct {
    DtAamAuthenticationServerRadiusInstanceStats AamAuthenticationServerRadiusInstanceStats `json:"instance"`
}


type AamAuthenticationServerRadiusInstanceStatsStats struct {
    Authen_success int `json:"authen_success"`
    Authen_failure int `json:"authen_failure"`
    Authorize_success int `json:"authorize_success"`
    Authorize_failure int `json:"authorize_failure"`
    Access_challenge int `json:"access_challenge"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    Request int `json:"request"`
    AccountingRequestSent int `json:"accounting-request-sent"`
    AccountingSuccess int `json:"accounting-success"`
    AccountingFailure int `json:"accounting-failure"`
}

func (p *AamAuthenticationServerRadiusInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerRadiusInstanceStats) getPath() string{
    
    return "aam/authentication/server/radius/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationServerRadiusInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerRadiusInstanceStats,error) {
logger.Println("AamAuthenticationServerRadiusInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerRadiusInstanceStats
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
