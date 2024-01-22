

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSessionOper struct {
    
    Oper AamAuthenticationSessionOperOper `json:"oper"`

}
type DataAamAuthenticationSessionOper struct {
    DtAamAuthenticationSessionOper AamAuthenticationSessionOper `json:"session"`
}


type AamAuthenticationSessionOperOper struct {
    SessionList []AamAuthenticationSessionOperOperSessionList `json:"session-list"`
    CmdType string `json:"cmd-type"`
    Username string `json:"username"`
    Vserver string `json:"vserver"`
    Ipv4 string `json:"ipv4"`
    Ipv6 string `json:"ipv6"`
    Partition string `json:"partition"`
}


type AamAuthenticationSessionOperOperSessionList struct {
    SessionId int `json:"Session-id"`
    Type string `json:"Type"`
    Vip string `json:"VIP"`
    Vport string `json:"VPort"`
    User string `json:"User"`
    ClientIp string `json:"Client-IP"`
    Domain string `json:"Domain"`
    DomainGroup string `json:"Domain-Group"`
    CreatedTime string `json:"Created-Time"`
    TtlInSeconds int `json:"TTL-in-seconds"`
    TokenLifetime string `json:"Token-Lifetime"`
}

func (p *AamAuthenticationSessionOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationSessionOper) getPath() string{
    return "aam/authentication/session/oper"
}

func (p *AamAuthenticationSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationSessionOper,error) {
logger.Println("AamAuthenticationSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationSessionOper
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
