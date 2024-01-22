

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationGlobalStats struct {
    
    Stats AamAuthenticationGlobalStatsStats `json:"stats"`

}
type DataAamAuthenticationGlobalStats struct {
    DtAamAuthenticationGlobalStats AamAuthenticationGlobalStats `json:"global"`
}


type AamAuthenticationGlobalStatsStats struct {
    Requests int `json:"requests"`
    Responses int `json:"responses"`
    Misses int `json:"misses"`
    OcspStaplingRequestsToA10authd int `json:"ocsp-stapling-requests-to-a10authd"`
    OcspStaplingResponsesFromA10authd int `json:"ocsp-stapling-responses-from-a10authd"`
    OpenedSocket int `json:"opened-socket"`
    OpenSocketFailed int `json:"open-socket-failed"`
    Connect int `json:"connect"`
    ConnectFailed int `json:"connect-failed"`
    CreatedTimer int `json:"created-timer"`
    CreateTimerFailed int `json:"create-timer-failed"`
    TotalRequest int `json:"total-request"`
    GetSocketOptionFailed int `json:"get-socket-option-failed"`
    AflexAuthzSucc int `json:"aflex-authz-succ"`
    AflexAuthzFail int `json:"aflex-authz-fail"`
    AuthnSuccess int `json:"authn-success"`
    AuthnFailure int `json:"authn-failure"`
    AuthzSuccess int `json:"authz-success"`
    AuthzFailure int `json:"authz-failure"`
    ActiveSession int `json:"active-session"`
    ActiveUser int `json:"active-user"`
    DnsResolveFailed int `json:"dns-resolve-failed"`
    DomainWlistMatch int `json:"domain-wlist-match"`
    DomainWlistUnmatch int `json:"domain-wlist-unmatch"`
    Auth_ctx_num int `json:"auth_ctx_num"`
}

func (p *AamAuthenticationGlobalStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationGlobalStats) getPath() string{
    return "aam/authentication/global/stats"
}

func (p *AamAuthenticationGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationGlobalStats,error) {
logger.Println("AamAuthenticationGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationGlobalStats
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
