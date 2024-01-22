

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslForwardProxyCertOper struct {
    
    Oper SlbSslForwardProxyCertOperOper `json:"oper"`

}
type DataSlbSslForwardProxyCertOper struct {
    DtSlbSslForwardProxyCertOper SlbSslForwardProxyCertOper `json:"ssl-forward-proxy-cert"`
}


type SlbSslForwardProxyCertOperOper struct {
    Vserver string `json:"vserver"`
    Port int `json:"port"`
    ServerIp string `json:"server-ip"`
    ServerPort int `json:"server-port"`
    ServerName string `json:"server-name"`
    HashedCertificate []SlbSslForwardProxyCertOperOperHashedCertificate `json:"hashed-certificate"`
    HashedPersistEntries int `json:"hashed-persist-entries"`
}


type SlbSslForwardProxyCertOperOperHashedCertificate struct {
    ServerIp string `json:"server-ip"`
    ServerIpv6 string `json:"server-ipv6"`
    RealPort int `json:"real-port"`
    Protocol string `json:"protocol"`
    ServerName string `json:"server-name"`
    State string `json:"state"`
    HitTimes int `json:"hit-times"`
    IdleTime int `json:"idle-time"`
    TimeoutAfter int `json:"timeout-after"`
    ExpiresAfter int `json:"expires-after"`
    AlpnProtocol int `json:"alpn-protocol"`
    Version int `json:"version"`
    Subject string `json:"subject"`
    CommonName string `json:"common-name"`
    Division string `json:"division"`
    Locality string `json:"locality"`
    StateProvince string `json:"state-province"`
    Country string `json:"country"`
    SubjectAltName string `json:"subject-alt-name"`
    Email string `json:"email"`
    StartTime string `json:"start-time"`
    ExpireTime string `json:"expire-time"`
    Issuer string `json:"issuer"`
    Serial string `json:"serial"`
}

func (p *SlbSslForwardProxyCertOper) GetId() string{
    return "1"
}

func (p *SlbSslForwardProxyCertOper) getPath() string{
    return "slb/ssl-forward-proxy-cert/oper"
}

func (p *SlbSslForwardProxyCertOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslForwardProxyCertOper,error) {
logger.Println("SlbSslForwardProxyCertOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslForwardProxyCertOper
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
