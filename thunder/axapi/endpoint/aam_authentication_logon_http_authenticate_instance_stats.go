

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationLogonHttpAuthenticateInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationLogonHttpAuthenticateInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationLogonHttpAuthenticateInstanceStats struct {
    DtAamAuthenticationLogonHttpAuthenticateInstanceStats AamAuthenticationLogonHttpAuthenticateInstanceStats `json:"instance"`
}


type AamAuthenticationLogonHttpAuthenticateInstanceStatsStats struct {
    Spn_krb_request int `json:"spn_krb_request"`
    Spn_krb_success int `json:"spn_krb_success"`
    Spn_krb_faiure int `json:"spn_krb_faiure"`
}

func (p *AamAuthenticationLogonHttpAuthenticateInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationLogonHttpAuthenticateInstanceStats) getPath() string{
    
    return "aam/authentication/logon/http-authenticate/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationLogonHttpAuthenticateInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationLogonHttpAuthenticateInstanceStats,error) {
logger.Println("AamAuthenticationLogonHttpAuthenticateInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationLogonHttpAuthenticateInstanceStats
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
