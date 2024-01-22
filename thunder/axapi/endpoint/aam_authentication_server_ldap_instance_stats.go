

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerLdapInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationServerLdapInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationServerLdapInstanceStats struct {
    DtAamAuthenticationServerLdapInstanceStats AamAuthenticationServerLdapInstanceStats `json:"instance"`
}


type AamAuthenticationServerLdapInstanceStatsStats struct {
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

func (p *AamAuthenticationServerLdapInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerLdapInstanceStats) getPath() string{
    
    return "aam/authentication/server/ldap/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationServerLdapInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerLdapInstanceStats,error) {
logger.Println("AamAuthenticationServerLdapInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerLdapInstanceStats
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
