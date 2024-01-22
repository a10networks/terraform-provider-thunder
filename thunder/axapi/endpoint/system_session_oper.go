

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSessionOper struct {
    
    Oper SystemSessionOperOper `json:"oper"`

}
type DataSystemSessionOper struct {
    DtSystemSessionOper SystemSessionOper `json:"session"`
}


type SystemSessionOperOper struct {
    All int `json:"all"`
    FilterName string `json:"filter-name"`
    Ipv4 int `json:"ipv4"`
    SourceV4Addr string `json:"source-v4-addr"`
    DestV4Addr string `json:"dest-v4-addr"`
    SourcePort int `json:"source-port"`
    DestPort int `json:"dest-port"`
    Ipv6 int `json:"ipv6"`
    SourceV6Addr string `json:"source-v6-addr"`
    DestV6Addr string `json:"dest-v6-addr"`
    V6SourcePort int `json:"v6-source-port"`
    V6DestPort int `json:"v6-dest-port"`
    Persist int `json:"persist"`
    Uie int `json:"uie"`
    PersistType string `json:"persist-type"`
    PersistSourceV4Addr string `json:"persist-source-v4-addr"`
    PersistDestV4Addr string `json:"persist-dest-v4-addr"`
    PersistSourceV6Addr string `json:"persist-source-v6-addr"`
    PersistDestV6Addr string `json:"persist-dest-v6-addr"`
    PersistSourcePort int `json:"persist-source-port"`
    PersistDestPort int `json:"persist-dest-port"`
    PersistIpv6 int `json:"persist-ipv6"`
    PersistIpv6Type string `json:"persist-ipv6-type"`
    PersistV6SourceAddr string `json:"persist-v6-source-addr"`
    PersistV6DestAddr string `json:"persist-v6-dest-addr"`
    PersistV6SourcePort int `json:"persist-v6-source-port"`
    PersistV6DestPort int `json:"persist-v6-dest-port"`
    Fw int `json:"fw"`
    HelperSessions int `json:"helper-sessions"`
    FwIpv4 int `json:"fw-ipv4"`
    FwIpv6 int `json:"fw-ipv6"`
    Sip int `json:"sip"`
    SipSourceV4Addr string `json:"sip-source-v4-addr"`
    SipDestV4Addr string `json:"sip-dest-v4-addr"`
    SipSourceV6Addr string `json:"sip-source-v6-addr"`
    SipDestV6Addr string `json:"sip-dest-v6-addr"`
    SipSourcePort int `json:"sip-source-port"`
    SipDestPort int `json:"sip-dest-port"`
}

func (p *SystemSessionOper) GetId() string{
    return "1"
}

func (p *SystemSessionOper) getPath() string{
    return "system/session/oper"
}

func (p *SystemSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemSessionOper,error) {
logger.Println("SystemSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemSessionOper
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
