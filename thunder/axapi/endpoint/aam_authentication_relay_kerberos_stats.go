

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayKerberosStats struct {
    
    InstanceList []AamAuthenticationRelayKerberosStatsInstanceList `json:"instance-list"`
    Stats AamAuthenticationRelayKerberosStatsStats `json:"stats"`

}
type DataAamAuthenticationRelayKerberosStats struct {
    DtAamAuthenticationRelayKerberosStats AamAuthenticationRelayKerberosStats `json:"kerberos"`
}


type AamAuthenticationRelayKerberosStatsInstanceList struct {
    Name string `json:"name"`
    Stats AamAuthenticationRelayKerberosStatsInstanceListStats `json:"stats"`
}


type AamAuthenticationRelayKerberosStatsInstanceListStats struct {
    RequestSend int `json:"request-send"`
    ResponseReceive int `json:"response-receive"`
    CurrentRequestsOfUser int `json:"current-requests-of-user"`
    Tickets int `json:"tickets"`
}


type AamAuthenticationRelayKerberosStatsStats struct {
    RequestSend int `json:"request-send"`
    ResponseGet int `json:"response-get"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    RequestNormal int `json:"request-normal"`
    RequestDropped int `json:"request-dropped"`
    ResponseSuccess int `json:"response-success"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
}

func (p *AamAuthenticationRelayKerberosStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayKerberosStats) getPath() string{
    return "aam/authentication/relay/kerberos/stats"
}

func (p *AamAuthenticationRelayKerberosStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayKerberosStats,error) {
logger.Println("AamAuthenticationRelayKerberosStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayKerberosStats
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
