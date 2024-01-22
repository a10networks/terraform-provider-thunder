

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerLdapStats struct {
    
    InstanceList []AamAuthenticationServerLdapStatsInstanceList `json:"instance-list"`
    Stats AamAuthenticationServerLdapStatsStats `json:"stats"`

}
type DataAamAuthenticationServerLdapStats struct {
    DtAamAuthenticationServerLdapStats AamAuthenticationServerLdapStats `json:"ldap"`
}


type AamAuthenticationServerLdapStatsInstanceList struct {
    Name string `json:"name"`
    Stats AamAuthenticationServerLdapStatsInstanceListStats `json:"stats"`
}


type AamAuthenticationServerLdapStatsInstanceListStats struct {
    AdminBindSuccess int `json:"admin-bind-success"`
    AdminBindFailure int `json:"admin-bind-failure"`
    BindSuccess int `json:"bind-success"`
    BindFailure int `json:"bind-failure"`
    SearchSuccess int `json:"search-success"`
    SearchFailure int `json:"search-failure"`
    AuthorizeSuccess int `json:"authorize-success"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    Request int `json:"request"`
    SslSessionCreated int `json:"ssl-session-created"`
    SslSessionFailure int `json:"ssl-session-failure"`
    Pw_expiry int `json:"pw_expiry"`
    Pw_change_success int `json:"pw_change_success"`
    Pw_change_failure int `json:"pw_change_failure"`
}


type AamAuthenticationServerLdapStatsStats struct {
    AdminBindSuccess int `json:"admin-bind-success"`
    AdminBindFailure int `json:"admin-bind-failure"`
    BindSuccess int `json:"bind-success"`
    BindFailure int `json:"bind-failure"`
    SearchSuccess int `json:"search-success"`
    SearchFailure int `json:"search-failure"`
    AuthorizeSuccess int `json:"authorize-success"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
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
    SslSessionCreated int `json:"ssl-session-created"`
    SslSessionFailure int `json:"ssl-session-failure"`
    LdapsIdleConnNum int `json:"ldaps-idle-conn-num"`
    LdapsInuseConnNum int `json:"ldaps-inuse-conn-num"`
    PwExpiry int `json:"pw-expiry"`
    PwChangeSuccess int `json:"pw-change-success"`
    PwChangeFailure int `json:"pw-change-failure"`
}

func (p *AamAuthenticationServerLdapStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerLdapStats) getPath() string{
    return "aam/authentication/server/ldap/stats"
}

func (p *AamAuthenticationServerLdapStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerLdapStats,error) {
logger.Println("AamAuthenticationServerLdapStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerLdapStats
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
