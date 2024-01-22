

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerRadiusStats struct {
    
    InstanceList []AamAuthenticationServerRadiusStatsInstanceList `json:"instance-list"`
    Stats AamAuthenticationServerRadiusStatsStats `json:"stats"`

}
type DataAamAuthenticationServerRadiusStats struct {
    DtAamAuthenticationServerRadiusStats AamAuthenticationServerRadiusStats `json:"radius"`
}


type AamAuthenticationServerRadiusStatsInstanceList struct {
    Name string `json:"name"`
    Stats AamAuthenticationServerRadiusStatsInstanceListStats `json:"stats"`
}


type AamAuthenticationServerRadiusStatsInstanceListStats struct {
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


type AamAuthenticationServerRadiusStatsStats struct {
    Authen_success int `json:"authen_success"`
    Authen_failure int `json:"authen_failure"`
    Authorize_success int `json:"authorize_success"`
    Authorize_failure int `json:"authorize_failure"`
    Access_challenge int `json:"access_challenge"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    Request int `json:"request"`
    RequestNormal int `json:"request-normal"`
    RequestDropped int `json:"request-dropped"`
    ResponseSuccess int `json:"response-success"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    AccountingRequestSent int `json:"accounting-request-sent"`
    AccountingSuccess int `json:"accounting-success"`
    AccountingFailure int `json:"accounting-failure"`
}

func (p *AamAuthenticationServerRadiusStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerRadiusStats) getPath() string{
    return "aam/authentication/server/radius/stats"
}

func (p *AamAuthenticationServerRadiusStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerRadiusStats,error) {
logger.Println("AamAuthenticationServerRadiusStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerRadiusStats
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
