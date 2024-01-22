

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayKerberosInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationRelayKerberosInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationRelayKerberosInstanceStats struct {
    DtAamAuthenticationRelayKerberosInstanceStats AamAuthenticationRelayKerberosInstanceStats `json:"instance"`
}


type AamAuthenticationRelayKerberosInstanceStatsStats struct {
    RequestSend int `json:"request-send"`
    ResponseReceive int `json:"response-receive"`
    CurrentRequestsOfUser int `json:"current-requests-of-user"`
    Tickets int `json:"tickets"`
}

func (p *AamAuthenticationRelayKerberosInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayKerberosInstanceStats) getPath() string{
    
    return "aam/authentication/relay/kerberos/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationRelayKerberosInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayKerberosInstanceStats,error) {
logger.Println("AamAuthenticationRelayKerberosInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayKerberosInstanceStats
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
