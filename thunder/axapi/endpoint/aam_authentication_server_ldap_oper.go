

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerLdapOper struct {
    
    Oper AamAuthenticationServerLdapOperOper `json:"oper"`

}
type DataAamAuthenticationServerLdapOper struct {
    DtAamAuthenticationServerLdapOper AamAuthenticationServerLdapOper `json:"ldap"`
}


type AamAuthenticationServerLdapOperOper struct {
    LdapsServerList []AamAuthenticationServerLdapOperOperLdapsServerList `json:"ldaps-server-list"`
}


type AamAuthenticationServerLdapOperOperLdapsServerList struct {
    LdapUri string `json:"ldap-uri"`
    LdapsIdleConnNum int `json:"ldaps-idle-conn-num"`
    LdapsIdleConnFdList string `json:"ldaps-idle-conn-fd-list"`
    LdapsInuseConnNum int `json:"ldaps-inuse-conn-num"`
    LdapsInuseConnFdList string `json:"ldaps-inuse-conn-fd-list"`
}

func (p *AamAuthenticationServerLdapOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerLdapOper) getPath() string{
    return "aam/authentication/server/ldap/oper"
}

func (p *AamAuthenticationServerLdapOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerLdapOper,error) {
logger.Println("AamAuthenticationServerLdapOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerLdapOper
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
