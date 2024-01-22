

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitor struct {
	Inst struct {

    DefaultStateUp int `json:"default-state-up"`
    DisableAfterDown int `json:"disable-after-down"`
    Dplane string `json:"dplane" dval:"auto"`
    DsrL2Strict int `json:"dsr-l2-strict"`
    HeaderInsert HealthMonitorHeaderInsert403 `json:"header-insert"`
    Interval int `json:"interval"`
    Method HealthMonitorMethod405 `json:"method"`
    Name string `json:"name"`
    OverrideIpv4 string `json:"override-ipv4"`
    OverrideIpv6 string `json:"override-ipv6"`
    OverridePort int `json:"override-port"`
    Passive int `json:"passive"`
    PassiveInterval int `json:"passive-interval" dval:"10"`
    ProxyHeader HealthMonitorProxyHeader437 `json:"proxy-header"`
    Retry int `json:"retry"`
    SampleThreshold int `json:"sample-threshold" dval:"50"`
    SslCiphers string `json:"ssl-ciphers" dval:"DEFAULT"`
    SslDgversion int `json:"ssl-dgversion" dval:"31"`
    SslTicket int `json:"ssl-ticket"`
    SslTicketLifetime int `json:"ssl-ticket-lifetime"`
    SslVersion int `json:"ssl-version" dval:"34"`
    StatusCode string `json:"status-code"`
    StrictRetryOnServerErrResp int `json:"strict-retry-on-server-err-resp"`
    TemplateServerSsl string `json:"template-server-ssl"`
    Threshold int `json:"threshold" dval:"75"`
    Timeout int `json:"timeout"`
    UpRetry int `json:"up-retry"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"monitor"`
}


type HealthMonitorHeaderInsert403 struct {
    InsertList []HealthMonitorHeaderInsertInsertList404 `json:"insert-list"`
    Uuid string `json:"uuid"`
}


type HealthMonitorHeaderInsertInsertList404 struct {
    InsertContent string `json:"insert-content"`
}


type HealthMonitorMethod405 struct {
    Icmp HealthMonitorMethodIcmp406 `json:"icmp"`
    Quic HealthMonitorMethodQuic407 `json:"quic"`
    Tcp HealthMonitorMethodTcp408 `json:"tcp"`
    Udp HealthMonitorMethodUdp410 `json:"udp"`
    Http HealthMonitorMethodHttp411 `json:"http"`
    Ftp HealthMonitorMethodFtp413 `json:"ftp"`
    Snmp HealthMonitorMethodSnmp414 `json:"snmp"`
    Smtp HealthMonitorMethodSmtp417 `json:"smtp"`
    Dns HealthMonitorMethodDns418 `json:"dns"`
    Pop3 HealthMonitorMethodPop3422 `json:"pop3"`
    Imap HealthMonitorMethodImap423 `json:"imap"`
    Sip HealthMonitorMethodSip424 `json:"sip"`
    Radius HealthMonitorMethodRadius425 `json:"radius"`
    Ldap HealthMonitorMethodLdap426 `json:"ldap"`
    Rtsp HealthMonitorMethodRtsp427 `json:"rtsp"`
    Database HealthMonitorMethodDatabase428 `json:"database"`
    External HealthMonitorMethodExternal429 `json:"external"`
    Ntp HealthMonitorMethodNtp430 `json:"ntp"`
    KerberosKdc HealthMonitorMethodKerberosKdc431 `json:"kerberos-kdc"`
    Https HealthMonitorMethodHttps433 `json:"https"`
    Tacplus HealthMonitorMethodTacplus435 `json:"tacplus"`
    Compound HealthMonitorMethodCompound436 `json:"compound"`
}


type HealthMonitorMethodIcmp406 struct {
    Icmp int `json:"icmp" dval:"1"`
    Transparent int `json:"transparent"`
    Ipv6 string `json:"ipv6"`
    Ip string `json:"ip"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodQuic407 struct {
    Quic int `json:"quic"`
    QuicPort int `json:"quic-port" dval:"443"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodTcp408 struct {
    MethodTcp int `json:"method-tcp"`
    TcpPort int `json:"tcp-port"`
    PortHalfopen int `json:"port-halfopen"`
    PortSend string `json:"port-send"`
    PortResp HealthMonitorMethodTcpPortResp409 `json:"port-resp"`
    Maintenance int `json:"maintenance"`
    MaintenanceText string `json:"maintenance-text"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodTcpPortResp409 struct {
    PortContains string `json:"port-contains"`
}


type HealthMonitorMethodUdp410 struct {
    Udp int `json:"udp"`
    UdpPort int `json:"udp-port"`
    ForceUpWithSingleHealthcheck int `json:"force-up-with-single-healthcheck"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodHttp411 struct {
    Http int `json:"http"`
    HttpPort int `json:"http-port" dval:"80"`
    Version2 int `json:"version2"`
    HttpExpect int `json:"http-expect"`
    HttpResponseCode string `json:"http-response-code"`
    ResponseCodeRegex string `json:"response-code-regex"`
    HttpText string `json:"http-text"`
    TextRegex string `json:"text-regex"`
    HttpHost string `json:"http-host"`
    HttpMaintenanceCode string `json:"http-maintenance-code"`
    HttpUrl int `json:"http-url"`
    UrlType string `json:"url-type"`
    Maintenance int `json:"maintenance"`
    MaintenanceText string `json:"maintenance-text"`
    MaintenanceTextRegex string `json:"maintenance-text-regex"`
    UrlPath string `json:"url-path"`
    PostPath string `json:"post-path"`
    PostType string `json:"post-type"`
    HttpPostdata string `json:"http-postdata"`
    HttpPostfile string `json:"http-postfile"`
    HttpUsername string `json:"http-username"`
    HttpPassword int `json:"http-password"`
    HttpPasswordString string `json:"http-password-string"`
    HttpEncrypted string `json:"http-encrypted"`
    HttpKerberosAuth int `json:"http-kerberos-auth"`
    HttpKerberosRealm string `json:"http-kerberos-realm"`
    HttpKerberosKdc HealthMonitorMethodHttpHttpKerberosKdc412 `json:"http-kerberos-kdc"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodHttpHttpKerberosKdc412 struct {
    HttpKerberosHostip string `json:"http-kerberos-hostip"`
    HttpKerberosHostipv6 string `json:"http-kerberos-hostipv6"`
    HttpKerberosPort int `json:"http-kerberos-port"`
    HttpKerberosPortv6 int `json:"http-kerberos-portv6"`
}


type HealthMonitorMethodFtp413 struct {
    Ftp int `json:"ftp"`
    FtpPort int `json:"ftp-port" dval:"21"`
    FtpUsername string `json:"ftp-username"`
    FtpPassword int `json:"ftp-password"`
    FtpPasswordString string `json:"ftp-password-string"`
    FtpEncrypted string `json:"ftp-encrypted"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodSnmp414 struct {
    Snmp int `json:"snmp"`
    SnmpPort int `json:"snmp-port" dval:"161"`
    Community string `json:"community" dval:"public"`
    Oid HealthMonitorMethodSnmpOid415 `json:"oid"`
    Operation HealthMonitorMethodSnmpOperation416 `json:"operation"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodSnmpOid415 struct {
    Mib string `json:"mib"`
    Asn string `json:"asn"`
}


type HealthMonitorMethodSnmpOperation416 struct {
    OperType string `json:"oper-type"`
}


type HealthMonitorMethodSmtp417 struct {
    Smtp int `json:"smtp"`
    SmtpDomain string `json:"smtp-domain"`
    SmtpPort int `json:"smtp-port" dval:"25"`
    SmtpStarttls int `json:"smtp-starttls"`
    MailFrom string `json:"mail-from"`
    RcptTo string `json:"rcpt-to"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodDns418 struct {
    Dns int `json:"dns"`
    DnsIpKey int `json:"dns-ip-key"`
    DnsIpv4Addr string `json:"dns-ipv4-addr"`
    DnsIpv6Addr string `json:"dns-ipv6-addr"`
    DnsIpv4Port int `json:"dns-ipv4-port" dval:"53"`
    DnsIpv4Expect HealthMonitorMethodDnsDnsIpv4Expect419 `json:"dns-ipv4-expect"`
    DnsIpv4Recurse string `json:"dns-ipv4-recurse" dval:"enabled"`
    DnsIpv4Tcp int `json:"dns-ipv4-tcp"`
    DnsIpv6Port int `json:"dns-ipv6-port" dval:"53"`
    DnsIpv6Expect HealthMonitorMethodDnsDnsIpv6Expect420 `json:"dns-ipv6-expect"`
    DnsIpv6Recurse string `json:"dns-ipv6-recurse" dval:"enabled"`
    DnsIpv6Tcp int `json:"dns-ipv6-tcp"`
    DnsDomain string `json:"dns-domain"`
    DnsDomainPort int `json:"dns-domain-port" dval:"53"`
    DnsDomainType string `json:"dns-domain-type" dval:"A"`
    DnsDomainExpect HealthMonitorMethodDnsDnsDomainExpect421 `json:"dns-domain-expect"`
    DnsDomainRecurse string `json:"dns-domain-recurse" dval:"enabled"`
    DnsDomainTcp int `json:"dns-domain-tcp"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodDnsDnsIpv4Expect419 struct {
    DnsIpv4Response string `json:"dns-ipv4-response"`
    DnsIpv4Fqdn string `json:"dns-ipv4-fqdn"`
}


type HealthMonitorMethodDnsDnsIpv6Expect420 struct {
    DnsIpv6Response string `json:"dns-ipv6-response"`
    DnsIpv6Fqdn string `json:"dns-ipv6-fqdn"`
}


type HealthMonitorMethodDnsDnsDomainExpect421 struct {
    DnsDomainResponse string `json:"dns-domain-response"`
    DnsDomainFqdn string `json:"dns-domain-fqdn"`
    DnsDomainIpv4 string `json:"dns-domain-ipv4"`
    DnsDomainIpv6 string `json:"dns-domain-ipv6"`
}


type HealthMonitorMethodPop3422 struct {
    Pop3 int `json:"pop3"`
    Pop3Username string `json:"pop3-username"`
    Pop3Password int `json:"pop3-password"`
    Pop3PasswordString string `json:"pop3-password-string"`
    Pop3Encrypted string `json:"pop3-encrypted"`
    Pop3Port int `json:"pop3-port" dval:"110"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodImap423 struct {
    Imap int `json:"imap"`
    ImapPort int `json:"imap-port" dval:"143"`
    ImapUsername string `json:"imap-username"`
    ImapPassword int `json:"imap-password"`
    ImapPasswordString string `json:"imap-password-string"`
    ImapEncrypted string `json:"imap-encrypted"`
    PwdAuth int `json:"pwd-auth"`
    ImapPlain int `json:"imap-plain"`
    ImapCramMd5 int `json:"imap-cram-md5"`
    ImapLogin int `json:"imap-login"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodSip424 struct {
    Sip int `json:"sip"`
    Register int `json:"register"`
    SipPort int `json:"sip-port" dval:"5060"`
    ExpectResponseCode string `json:"expect-response-code"`
    SipTcp int `json:"sip-tcp"`
    SipHostname string `json:"sip-hostname"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodRadius425 struct {
    Radius int `json:"radius"`
    RadiusUsername string `json:"radius-username"`
    RadiusPasswordString string `json:"radius-password-string"`
    RadiusEncrypted string `json:"radius-encrypted"`
    RadiusSecret string `json:"radius-secret"`
    RadiusSecretEncrypted string `json:"radius-secret-encrypted"`
    RadiusPort int `json:"radius-port" dval:"1812"`
    RadiusExpect int `json:"radius-expect"`
    RadiusResponseCode string `json:"radius-response-code"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodLdap426 struct {
    Ldap int `json:"ldap"`
    LdapPort int `json:"ldap-port" dval:"389"`
    LdapSecurity string `json:"ldap-security"`
    LdapBinddn string `json:"ldap-binddn"`
    LdapPassword int `json:"ldap-password"`
    LdapPasswordString string `json:"ldap-password-string"`
    LdapEncrypted string `json:"ldap-encrypted"`
    LdapRunSearch int `json:"ldap-run-search"`
    Basedn string `json:"BaseDN"`
    LdapQuery string `json:"ldap-query"`
    Acceptresref int `json:"AcceptResRef"`
    Acceptnotfound int `json:"AcceptNotFound"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodRtsp427 struct {
    Rtsp int `json:"rtsp"`
    Rtspurl string `json:"rtspurl"`
    RtspPort int `json:"rtsp-port" dval:"554"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodDatabase428 struct {
    Database int `json:"database"`
    DatabaseName string `json:"database-name"`
    DbName string `json:"db-name"`
    DbUsername string `json:"db-username"`
    DbPassword int `json:"db-password"`
    DbPasswordStr string `json:"db-password-str"`
    DbEncrypted string `json:"db-encrypted"`
    DbSend string `json:"db-send"`
    DbReceive string `json:"db-receive"`
    DbRow int `json:"db-row"`
    DbColumn int `json:"db-column"`
    DbReceiveInteger int `json:"db-receive-integer"`
    DbRowInteger int `json:"db-row-integer"`
    DbColumnInteger int `json:"db-column-integer"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodExternal429 struct {
    External int `json:"external"`
    ExtProgram string `json:"ext-program"`
    SharedPartitionProgram int `json:"shared-partition-program"`
    ExtProgramShared string `json:"ext-program-shared"`
    ExtPort int `json:"ext-port"`
    ExtArguments string `json:"ext-arguments"`
    ExtPreference int `json:"ext-preference"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodNtp430 struct {
    Ntp int `json:"ntp"`
    NtpPort int `json:"ntp-port" dval:"123"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodKerberosKdc431 struct {
    KerberosCfg HealthMonitorMethodKerberosKdcKerberosCfg432 `json:"kerberos-cfg"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodKerberosKdcKerberosCfg432 struct {
    Kinit int `json:"kinit"`
    KinitPricipalName string `json:"kinit-pricipal-name"`
    KinitPassword string `json:"kinit-password"`
    KinitEncrypted string `json:"kinit-encrypted"`
    KinitKdc string `json:"kinit-kdc"`
    TcpOnly int `json:"tcp-only"`
    Kadmin int `json:"kadmin"`
    KadminRealm string `json:"kadmin-realm"`
    KadminPricipalName string `json:"kadmin-pricipal-name"`
    KadminPassword string `json:"kadmin-password"`
    KadminEncrypted string `json:"kadmin-encrypted"`
    KadminServer string `json:"kadmin-server"`
    KadminKdc string `json:"kadmin-kdc"`
    Kpasswd int `json:"kpasswd"`
    KpasswdPricipalName string `json:"kpasswd-pricipal-name"`
    KpasswdPassword string `json:"kpasswd-password"`
    KpasswdEncrypted string `json:"kpasswd-encrypted"`
    KpasswdServer string `json:"kpasswd-server"`
    KpasswdKdc string `json:"kpasswd-kdc"`
}


type HealthMonitorMethodHttps433 struct {
    Https int `json:"https"`
    WebPort int `json:"web-port" dval:"443"`
    DisableSslv2hello int `json:"disable-sslv2hello"`
    HttpVersion string `json:"http-version"`
    HttpsHost string `json:"https-host"`
    Sni int `json:"sni"`
    HttpsExpect int `json:"https-expect"`
    HttpsResponseCode string `json:"https-response-code"`
    ResponseCodeRegex string `json:"response-code-regex"`
    HttpsText string `json:"https-text"`
    TextRegex string `json:"text-regex"`
    HttpsUrl int `json:"https-url"`
    UrlType string `json:"url-type"`
    UrlPath string `json:"url-path"`
    PostPath string `json:"post-path"`
    PostType string `json:"post-type"`
    HttpsPostdata string `json:"https-postdata"`
    HttpsPostfile string `json:"https-postfile"`
    HttpsMaintenanceCode string `json:"https-maintenance-code"`
    Maintenance int `json:"maintenance"`
    MaintenanceText string `json:"maintenance-text"`
    MaintenanceTextRegex string `json:"maintenance-text-regex"`
    HttpsUsername string `json:"https-username"`
    HttpsServerCertName string `json:"https-server-cert-name"`
    HttpsPassword int `json:"https-password"`
    HttpsPasswordString string `json:"https-password-string"`
    HttpsEncrypted string `json:"https-encrypted"`
    HttpsKerberosAuth int `json:"https-kerberos-auth"`
    HttpsKerberosRealm string `json:"https-kerberos-realm"`
    HttpsKerberosKdc HealthMonitorMethodHttpsHttpsKerberosKdc434 `json:"https-kerberos-kdc"`
    CertKeyShared int `json:"cert-key-shared"`
    Cert string `json:"cert"`
    Key string `json:"key"`
    KeyPassPhrase int `json:"key-pass-phrase"`
    KeyPhrase string `json:"key-phrase"`
    HttpsKeyEncrypted string `json:"https-key-encrypted"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodHttpsHttpsKerberosKdc434 struct {
    HttpsKerberosHostip string `json:"https-kerberos-hostip"`
    HttpsKerberosHostipv6 string `json:"https-kerberos-hostipv6"`
    HttpsKerberosPort int `json:"https-kerberos-port"`
    HttpsKerberosPortv6 int `json:"https-kerberos-portv6"`
}


type HealthMonitorMethodTacplus435 struct {
    Tacplus int `json:"tacplus"`
    TacplusUsername string `json:"tacplus-username"`
    TacplusPassword int `json:"tacplus-password"`
    TacplusPasswordString string `json:"tacplus-password-string"`
    TacplusEncrypted string `json:"tacplus-encrypted"`
    TacplusSecret int `json:"tacplus-secret"`
    TacplusSecretString string `json:"tacplus-secret-string"`
    SecretEncrypted string `json:"secret-encrypted"`
    TacplusPort int `json:"tacplus-port" dval:"49"`
    TacplusType string `json:"tacplus-type" dval:"inbound-ascii-login"`
    Uuid string `json:"uuid"`
}


type HealthMonitorMethodCompound436 struct {
    Compound int `json:"compound"`
    RpnString string `json:"rpn-string"`
    Uuid string `json:"uuid"`
}


type HealthMonitorProxyHeader437 struct {
    ProxyHeaderVer string `json:"proxy-header-ver"`
    Uuid string `json:"uuid"`
}

func (p *HealthMonitor) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *HealthMonitor) getPath() string{
    return "health/monitor"
}

func (p *HealthMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitor::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *HealthMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitor::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *HealthMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitor::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *HealthMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
