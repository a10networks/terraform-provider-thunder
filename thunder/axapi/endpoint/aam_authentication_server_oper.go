

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerOper struct {
    
    Ldap AamAuthenticationServerOperLdap `json:"ldap"`
    Ocsp AamAuthenticationServerOperOcsp `json:"ocsp"`
    Oper AamAuthenticationServerOperOper `json:"oper"`
    Windows AamAuthenticationServerOperWindows `json:"windows"`

}
type DataAamAuthenticationServerOper struct {
    DtAamAuthenticationServerOper AamAuthenticationServerOper `json:"server"`
}


type AamAuthenticationServerOperLdap struct {
    Oper AamAuthenticationServerOperLdapOper `json:"oper"`
}


type AamAuthenticationServerOperLdapOper struct {
    LdapsServerList []AamAuthenticationServerOperLdapOperLdapsServerList `json:"ldaps-server-list"`
}


type AamAuthenticationServerOperLdapOperLdapsServerList struct {
    LdapUri string `json:"ldap-uri"`
    LdapsIdleConnNum int `json:"ldaps-idle-conn-num"`
    LdapsIdleConnFdList string `json:"ldaps-idle-conn-fd-list"`
    LdapsInuseConnNum int `json:"ldaps-inuse-conn-num"`
    LdapsInuseConnFdList string `json:"ldaps-inuse-conn-fd-list"`
}


type AamAuthenticationServerOperOcsp struct {
    Oper AamAuthenticationServerOperOcspOper `json:"oper"`
}


type AamAuthenticationServerOperOcspOper struct {
    StatsClearType string `json:"stats-clear-type"`
    Name string `json:"name"`
}


type AamAuthenticationServerOperOper struct {
    RserverCount int `json:"rserver-count"`
    RportCount int `json:"rport-count"`
    RserverList []AamAuthenticationServerOperOperRserverList `json:"rserver-list"`
    Name string `json:"name"`
    PartId int `json:"part-id"`
    GetCount string `json:"get-count"`
}


type AamAuthenticationServerOperOperRserverList struct {
    ServerName string `json:"server-name"`
    Host string `json:"host"`
    Ip string `json:"ip"`
    Hm string `json:"hm"`
    Status string `json:"status"`
    MaxConn int `json:"max-conn"`
    Weight int `json:"weight"`
    RportList []AamAuthenticationServerOperOperRserverListRportList `json:"rport-list"`
}


type AamAuthenticationServerOperOperRserverListRportList struct {
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    PortState string `json:"port-state"`
    PortHm string `json:"port-hm"`
    PortStatus string `json:"port-status"`
    PortMaxConn int `json:"port-max-conn"`
    PortWeight int `json:"port-weight"`
    SgList []AamAuthenticationServerOperOperRserverListRportListSgList `json:"sg-list"`
}


type AamAuthenticationServerOperOperRserverListRportListSgList struct {
    SgName string `json:"sg-name"`
    SgState string `json:"sg-state"`
}


type AamAuthenticationServerOperWindows struct {
    Oper AamAuthenticationServerOperWindowsOper `json:"oper"`
}


type AamAuthenticationServerOperWindowsOper struct {
    StatsClearType string `json:"stats-clear-type"`
    Name string `json:"name"`
}

func (p *AamAuthenticationServerOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerOper) getPath() string{
    return "aam/authentication/server/oper"
}

func (p *AamAuthenticationServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerOper,error) {
logger.Println("AamAuthenticationServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerOper
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
