

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthCheckDetailsOper struct {
    
    Oper SlbHealthCheckDetailsOperOper `json:"oper"`

}
type DataSlbHealthCheckDetailsOper struct {
    DtSlbHealthCheckDetailsOper SlbHealthCheckDetailsOper `json:"health-check-details"`
}


type SlbHealthCheckDetailsOperOper struct {
    PinId int `json:"pin-id"`
    ProcessIndex int `json:"process-index"`
    HealthState string `json:"health-state"`
    StateReason string `json:"state-reason"`
    MonitorName string `json:"monitor-name"`
    ReceivedSuccess int `json:"received-success"`
    ReceivedFail int `json:"received-fail"`
    ResponseTimeout int `json:"response-timeout"`
    CurrInterval int `json:"curr-interval"`
    Method string `json:"method"`
    AttrAliasAddr string `json:"attr-alias-addr"`
    AttrPort int `json:"attr-port"`
    HalfOpen int `json:"half-open"`
    Send string `json:"send"`
    RespCont string `json:"resp-cont"`
    ForceUp int `json:"force-up"`
    Url string `json:"url"`
    ExpectText string `json:"expect-text"`
    ExpectRespCode string `json:"expect-resp-code"`
    ExpectTextRegex string `json:"expect-text-regex"`
    ExpectRespRegexCode string `json:"expect-resp-regex-code"`
    MaintenanceCode string `json:"maintenance-code"`
    User string `json:"user"`
    Pass string `json:"pass"`
    Postdata string `json:"postdata"`
    Host string `json:"host"`
    KerberosRealm string `json:"kerberos-realm"`
    KerberosKdc string `json:"kerberos-kdc"`
    KerberosPort int `json:"kerberos-port"`
    SnmpOperation int `json:"snmp-operation"`
    Community string `json:"community"`
    Oid string `json:"oid"`
    Domain string `json:"domain"`
    Starttls int `json:"starttls"`
    MailFrom string `json:"mail-from"`
    RcptTo string `json:"rcpt-to"`
    Ipaddr string `json:"ipaddr"`
    DnsQtype int `json:"dns-qtype"`
    DnsRecurse int `json:"dns-recurse"`
    DnsExpectType int `json:"dns-expect-type"`
    DnsExpect string `json:"dns-expect"`
    TransportProto int `json:"transport-proto"`
    SipRegister int `json:"sip-register"`
    Secret string `json:"secret"`
    Query string `json:"query"`
    BaseDn string `json:"base-dn"`
    LdapSsl int `json:"ldap-ssl"`
    LdapTls int `json:"ldap-tls"`
    AttrType string `json:"attr-type"`
    DbName string `json:"db-name"`
    Receive string `json:"receive"`
    RcvInteger int `json:"rcv-integer"`
    DbRow int `json:"db-row"`
    DbColumn int `json:"db-column"`
    Pname string `json:"pname"`
    TcpOnly int `json:"tcp-only"`
    AttrProgram string `json:"attr-program"`
    Arguments string `json:"arguments"`
    AttrRpn string `json:"attr-rpn"`
    HttpWaitResp int `json:"http-wait-resp"`
    L4ConnNum int `json:"l4-conn-num"`
    L4Errors int `json:"l4-errors"`
    AvgRtt int `json:"avg-rtt"`
    CurrRtt int `json:"curr-rtt"`
    AvgTcpRtt int `json:"avg-tcp-rtt"`
    CurrTcpRtt int `json:"curr-tcp-rtt"`
    StatusCodeRcv int `json:"status-code-rcv"`
    HttpReqSent int `json:"http-req-sent"`
    HttpErrors int `json:"http-errors"`
    MacAddr string `json:"mac-addr"`
}

func (p *SlbHealthCheckDetailsOper) GetId() string{
    return "1"
}

func (p *SlbHealthCheckDetailsOper) getPath() string{
    return "slb/health-check-details/oper"
}

func (p *SlbHealthCheckDetailsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthCheckDetailsOper,error) {
logger.Println("SlbHealthCheckDetailsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthCheckDetailsOper
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
