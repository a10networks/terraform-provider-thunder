package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 5_2_1-P4_90
type HealthMonitor struct {
	Inst struct {
		DefaultStateUp             int                 `json:"default-state-up"`
		DisableAfterDown           int                 `json:"disable-after-down"`
		DsrL2Strict                int                 `json:"dsr-l2-strict"`
		Interval                   int                 `json:"interval"`
		Method                     HealthMonitorMethod `json:"method"`
		Name                       string              `json:"name"`
		OverrideIpv4               string              `json:"override-ipv4"`
		OverrideIpv6               string              `json:"override-ipv6"`
		OverridePort               int                 `json:"override-port"`
		Passive                    int                 `json:"passive"`
		PassiveInterval            int                 `json:"passive-interval" dval:"10"`
		Retry                      int                 `json:"retry"`
		SampleThreshold            int                 `json:"sample-threshold" dval:"50"`
		SslCiphers                 string              `json:"ssl-ciphers" dval:"DEFAULT"`
		SslDgversion               int                 `json:"ssl-dgversion" dval:"-1"` //ssl-dgversion is mandatory when ssl-version is configured
		SslTicket                  int                 `json:"ssl-ticket"`
		SslTicketLifetime          int                 `json:"ssl-ticket-lifetime"`
		SslVersion                 int                 `json:"ssl-version" dval:"-1"`
		StatusCode                 string              `json:"status-code"`
		StrictRetryOnServerErrResp int                 `json:"strict-retry-on-server-err-resp"`
		Threshold                  int                 `json:"threshold" dval:"75"`
		Timeout                    int                 `json:"timeout"`
		UpRetry                    int                 `json:"up-retry"`
		UserTag                    string              `json:"user-tag"`
		Uuid                       string              `json:"uuid"`
	} `json:"monitor"`
}

type HealthMonitorMethod struct {
	Icmp        HealthMonitorMethodIcmp        `json:"icmp"`
	Tcp         HealthMonitorMethodTcp         `json:"tcp"`
	Udp         HealthMonitorMethodUdp         `json:"udp"`
	Http        HealthMonitorMethodHttp        `json:"http"`
	Ftp         HealthMonitorMethodFtp         `json:"ftp"`
	Snmp        HealthMonitorMethodSnmp        `json:"snmp"`
	Smtp        HealthMonitorMethodSmtp        `json:"smtp"`
	Dns         HealthMonitorMethodDns         `json:"dns"`
	Pop3        HealthMonitorMethodPop3        `json:"pop3"`
	Imap        HealthMonitorMethodImap        `json:"imap"`
	Sip         HealthMonitorMethodSip         `json:"sip"`
	Radius      HealthMonitorMethodRadius      `json:"radius"`
	Ldap        HealthMonitorMethodLdap        `json:"ldap"`
	Rtsp        HealthMonitorMethodRtsp        `json:"rtsp"`
	Database    HealthMonitorMethodDatabase    `json:"database"`
	External    HealthMonitorMethodExternal    `json:"external"`
	Ntp         HealthMonitorMethodNtp         `json:"ntp"`
	KerberosKdc HealthMonitorMethodKerberosKdc `json:"kerberos-kdc"`
	Https       HealthMonitorMethodHttps       `json:"https"`
	Tacplus     HealthMonitorMethodTacplus     `json:"tacplus"`
	Compound    HealthMonitorMethodCompound    `json:"compound"`
}

type HealthMonitorMethodIcmp struct {
	Icmp        int    `json:"icmp" dval:"1"`
	Transparent int    `json:"transparent"`
	Ipv6        string `json:"ipv6"`
	Ip          string `json:"ip"`
	Uuid        string `json:"uuid"`
}

type HealthMonitorMethodTcp struct {
	MethodTcp       int                            `json:"method-tcp"`
	TcpPort         int                            `json:"tcp-port"`
	PortHalfopen    int                            `json:"port-halfopen"`
	PortSend        string                         `json:"port-send"`
	PortResp        HealthMonitorMethodTcpPortResp `json:"port-resp"`
	Maintenance     int                            `json:"maintenance"`
	MaintenanceText string                         `json:"maintenance-text"`
	Uuid            string                         `json:"uuid"`
}

type HealthMonitorMethodTcpPortResp struct {
	PortContains string `json:"port-contains"`
}

type HealthMonitorMethodUdp struct {
	Udp                          int    `json:"udp"`
	UdpPort                      int    `json:"udp-port"`
	ForceUpWithSingleHealthcheck int    `json:"force-up-with-single-healthcheck"`
	Uuid                         string `json:"uuid"`
}

type HealthMonitorMethodHttp struct {
	Http                 int                                    `json:"http"`
	HttpPort             int                                    `json:"http-port" dval:"80"`
	HttpExpect           int                                    `json:"http-expect"`
	HttpResponseCode     string                                 `json:"http-response-code"`
	ResponseCodeRegex    string                                 `json:"response-code-regex"`
	HttpText             string                                 `json:"http-text"`
	TextRegex            string                                 `json:"text-regex"`
	HttpHost             string                                 `json:"http-host"`
	HttpMaintenanceCode  string                                 `json:"http-maintenance-code"`
	HttpUrl              int                                    `json:"http-url"`
	UrlType              string                                 `json:"url-type"`
	Maintenance          int                                    `json:"maintenance"`
	MaintenanceText      string                                 `json:"maintenance-text"`
	MaintenanceTextRegex string                                 `json:"maintenance-text-regex"`
	UrlPath              string                                 `json:"url-path"`
	PostPath             string                                 `json:"post-path"`
	PostType             string                                 `json:"post-type"`
	HttpPostdata         string                                 `json:"http-postdata"`
	HttpPostfile         string                                 `json:"http-postfile"`
	HttpUsername         string                                 `json:"http-username"`
	HttpPassword         int                                    `json:"http-password"`
	HttpPasswordString   string                                 `json:"http-password-string"`
	HttpEncrypted        string                                 `json:"http-encrypted"`
	HttpKerberosAuth     int                                    `json:"http-kerberos-auth"`
	HttpKerberosRealm    string                                 `json:"http-kerberos-realm"`
	HttpKerberosKdc      HealthMonitorMethodHttpHttpKerberosKdc `json:"http-kerberos-kdc"`
	Uuid                 string                                 `json:"uuid"`
}

type HealthMonitorMethodHttpHttpKerberosKdc struct {
	HttpKerberosHostip   string `json:"http-kerberos-hostip"`
	HttpKerberosHostipv6 string `json:"http-kerberos-hostipv6"`
	HttpKerberosPort     int    `json:"http-kerberos-port"`
	HttpKerberosPortv6   int    `json:"http-kerberos-portv6"`
}

type HealthMonitorMethodFtp struct {
	Ftp               int    `json:"ftp"`
	FtpPort           int    `json:"ftp-port" dval:"21"`
	FtpUsername       string `json:"ftp-username"`
	FtpPassword       int    `json:"ftp-password"`
	FtpPasswordString string `json:"ftp-password-string"`
	FtpEncrypted      string `json:"ftp-encrypted"`
	Uuid              string `json:"uuid"`
}

type HealthMonitorMethodSnmp struct {
	Snmp      int                              `json:"snmp"`
	SnmpPort  int                              `json:"snmp-port" dval:"161"`
	Community string                           `json:"community" dval:"public"`
	Oid       HealthMonitorMethodSnmpOid       `json:"oid"`
	Operation HealthMonitorMethodSnmpOperation `json:"operation"`
	Uuid      string                           `json:"uuid"`
}

type HealthMonitorMethodSnmpOid struct {
	Mib string `json:"mib"`
	Asn string `json:"asn"`
}

type HealthMonitorMethodSnmpOperation struct {
	OperType string `json:"oper-type"`
}

type HealthMonitorMethodSmtp struct {
	Smtp         int    `json:"smtp"`
	SmtpDomain   string `json:"smtp-domain"`
	SmtpPort     int    `json:"smtp-port" dval:"25"`
	SmtpStarttls int    `json:"smtp-starttls"`
	MailFrom     string `json:"mail-from"`
	RcptTo       string `json:"rcpt-to"`
	Uuid         string `json:"uuid"`
}

type HealthMonitorMethodDns struct {
	Dns              int                                   `json:"dns"`
	DnsIpKey         int                                   `json:"dns-ip-key"`
	DnsIpv4Addr      string                                `json:"dns-ipv4-addr"`
	DnsIpv6Addr      string                                `json:"dns-ipv6-addr"`
	DnsIpv4Port      int                                   `json:"dns-ipv4-port" dval:"53"`
	DnsIpv4Expect    HealthMonitorMethodDnsDnsIpv4Expect   `json:"dns-ipv4-expect"`
	DnsIpv4Recurse   string                                `json:"dns-ipv4-recurse" dval:"enabled"`
	DnsIpv4Tcp       int                                   `json:"dns-ipv4-tcp"`
	DnsIpv6Port      int                                   `json:"dns-ipv6-port" dval:"53"`
	DnsIpv6Expect    HealthMonitorMethodDnsDnsIpv6Expect   `json:"dns-ipv6-expect"`
	DnsIpv6Recurse   string                                `json:"dns-ipv6-recurse" dval:"enabled"`
	DnsIpv6Tcp       int                                   `json:"dns-ipv6-tcp"`
	DnsDomain        string                                `json:"dns-domain"`
	DnsDomainPort    int                                   `json:"dns-domain-port" dval:"53"`
	DnsDomainType    string                                `json:"dns-domain-type" dval:"A"`
	DnsDomainExpect  HealthMonitorMethodDnsDnsDomainExpect `json:"dns-domain-expect"`
	DnsDomainRecurse string                                `json:"dns-domain-recurse" dval:"enabled"`
	DnsDomainTcp     int                                   `json:"dns-domain-tcp"`
	Uuid             string                                `json:"uuid"`
}

type HealthMonitorMethodDnsDnsIpv4Expect struct {
	DnsIpv4Response string `json:"dns-ipv4-response"`
	DnsIpv4Fqdn     string `json:"dns-ipv4-fqdn"`
}

type HealthMonitorMethodDnsDnsIpv6Expect struct {
	DnsIpv6Response string `json:"dns-ipv6-response"`
	DnsIpv6Fqdn     string `json:"dns-ipv6-fqdn"`
}

type HealthMonitorMethodDnsDnsDomainExpect struct {
	DnsDomainResponse string `json:"dns-domain-response"`
	DnsDomainFqdn     string `json:"dns-domain-fqdn"`
	DnsDomainIpv4     string `json:"dns-domain-ipv4"`
	DnsDomainIpv6     string `json:"dns-domain-ipv6"`
}

type HealthMonitorMethodPop3 struct {
	Pop3               int    `json:"pop3"`
	Pop3Username       string `json:"pop3-username"`
	Pop3Password       int    `json:"pop3-password"`
	Pop3PasswordString string `json:"pop3-password-string"`
	Pop3Encrypted      string `json:"pop3-encrypted"`
	Pop3Port           int    `json:"pop3-port" dval:"110"`
	Uuid               string `json:"uuid"`
}

type HealthMonitorMethodImap struct {
	Imap               int    `json:"imap"`
	ImapPort           int    `json:"imap-port" dval:"143"`
	ImapUsername       string `json:"imap-username"`
	ImapPassword       int    `json:"imap-password"`
	ImapPasswordString string `json:"imap-password-string"`
	ImapEncrypted      string `json:"imap-encrypted"`
	PwdAuth            int    `json:"pwd-auth"`
	ImapPlain          int    `json:"imap-plain"`
	ImapCramMd5        int    `json:"imap-cram-md5"`
	ImapLogin          int    `json:"imap-login"`
	Uuid               string `json:"uuid"`
}

type HealthMonitorMethodSip struct {
	Sip                int    `json:"sip"`
	Register           int    `json:"register"`
	SipPort            int    `json:"sip-port" dval:"5060"`
	ExpectResponseCode string `json:"expect-response-code"`
	SipTcp             int    `json:"sip-tcp"`
	Uuid               string `json:"uuid"`
}

type HealthMonitorMethodRadius struct {
	Radius                int    `json:"radius"`
	RadiusUsername        string `json:"radius-username"`
	RadiusPasswordString  string `json:"radius-password-string"`
	RadiusEncrypted       string `json:"radius-encrypted"`
	RadiusSecret          string `json:"radius-secret"`
	RadiusSecretEncrypted string `json:"radius-secret-encrypted"`
	RadiusPort            int    `json:"radius-port" dval:"1812"`
	RadiusExpect          int    `json:"radius-expect"`
	RadiusResponseCode    string `json:"radius-response-code"`
	Uuid                  string `json:"uuid"`
}

type HealthMonitorMethodLdap struct {
	Ldap               int    `json:"ldap"`
	LdapPort           int    `json:"ldap-port" dval:"389"`
	LdapSecurity       string `json:"ldap-security"`
	LdapBinddn         string `json:"ldap-binddn"`
	LdapPassword       int    `json:"ldap-password"`
	LdapPasswordString string `json:"ldap-password-string"`
	LdapEncrypted      string `json:"ldap-encrypted"`
	LdapRunSearch      int    `json:"ldap-run-search"`
	Basedn             string `json:"BaseDN"`
	LdapQuery          string `json:"ldap-query"`
	Acceptresref       int    `json:"AcceptResRef"`
	Acceptnotfound     int    `json:"AcceptNotFound"`
	Uuid               string `json:"uuid"`
}

type HealthMonitorMethodRtsp struct {
	Rtsp     int    `json:"rtsp"`
	Rtspurl  string `json:"rtspurl"`
	RtspPort int    `json:"rtsp-port" dval:"554"`
	Uuid     string `json:"uuid"`
}

type HealthMonitorMethodDatabase struct {
	Database         int    `json:"database"`
	DatabaseName     string `json:"database-name"`
	DbName           string `json:"db-name"`
	DbUsername       string `json:"db-username"`
	DbPassword       int    `json:"db-password"`
	DbPasswordStr    string `json:"db-password-str"`
	DbEncrypted      string `json:"db-encrypted"`
	DbSend           string `json:"db-send"`
	DbReceive        string `json:"db-receive"`
	DbRow            int    `json:"db-row"`
	DbColumn         int    `json:"db-column"`
	DbReceiveInteger int    `json:"db-receive-integer"`
	DbRowInteger     int    `json:"db-row-integer"`
	DbColumnInteger  int    `json:"db-column-integer"`
	Uuid             string `json:"uuid"`
}

type HealthMonitorMethodExternal struct {
	External               int    `json:"external"`
	ExtProgram             string `json:"ext-program"`
	SharedPartitionProgram int    `json:"shared-partition-program"`
	ExtProgramShared       string `json:"ext-program-shared"`
	ExtPort                int    `json:"ext-port"`
	ExtArguments           string `json:"ext-arguments"`
	ExtPreference          int    `json:"ext-preference"`
	Uuid                   string `json:"uuid"`
}

type HealthMonitorMethodNtp struct {
	Ntp     int    `json:"ntp"`
	NtpPort int    `json:"ntp-port" dval:"123"`
	Uuid    string `json:"uuid"`
}

type HealthMonitorMethodKerberosKdc struct {
	KerberosCfg HealthMonitorMethodKerberosKdcKerberosCfg `json:"kerberos-cfg"`
	Uuid        string                                    `json:"uuid"`
}

type HealthMonitorMethodKerberosKdcKerberosCfg struct {
	Kinit               int    `json:"kinit"`
	KinitPricipalName   string `json:"kinit-pricipal-name"`
	KinitPassword       string `json:"kinit-password"`
	KinitEncrypted      string `json:"kinit-encrypted"`
	KinitKdc            string `json:"kinit-kdc"`
	TcpOnly             int    `json:"tcp-only"`
	Kadmin              int    `json:"kadmin"`
	KadminRealm         string `json:"kadmin-realm"`
	KadminPricipalName  string `json:"kadmin-pricipal-name"`
	KadminPassword      string `json:"kadmin-password"`
	KadminEncrypted     string `json:"kadmin-encrypted"`
	KadminServer        string `json:"kadmin-server"`
	KadminKdc           string `json:"kadmin-kdc"`
	Kpasswd             int    `json:"kpasswd"`
	KpasswdPricipalName string `json:"kpasswd-pricipal-name"`
	KpasswdPassword     string `json:"kpasswd-password"`
	KpasswdEncrypted    string `json:"kpasswd-encrypted"`
	KpasswdServer       string `json:"kpasswd-server"`
	KpasswdKdc          string `json:"kpasswd-kdc"`
}

type HealthMonitorMethodHttps struct {
	Https                int                                      `json:"https"`
	WebPort              int                                      `json:"web-port" dval:"443"`
	DisableSslv2hello    int                                      `json:"disable-sslv2hello"`
	HttpsHost            string                                   `json:"https-host"`
	Sni                  int                                      `json:"sni"`
	HttpsExpect          int                                      `json:"https-expect"`
	HttpsResponseCode    string                                   `json:"https-response-code"`
	ResponseCodeRegex    string                                   `json:"response-code-regex"`
	HttpsText            string                                   `json:"https-text"`
	TextRegex            string                                   `json:"text-regex"`
	HttpsUrl             int                                      `json:"https-url"`
	UrlType              string                                   `json:"url-type"`
	UrlPath              string                                   `json:"url-path"`
	PostPath             string                                   `json:"post-path"`
	PostType             string                                   `json:"post-type"`
	HttpsPostdata        string                                   `json:"https-postdata"`
	HttpsPostfile        string                                   `json:"https-postfile"`
	HttpsMaintenanceCode string                                   `json:"https-maintenance-code"`
	Maintenance          int                                      `json:"maintenance"`
	MaintenanceText      string                                   `json:"maintenance-text"`
	MaintenanceTextRegex string                                   `json:"maintenance-text-regex"`
	HttpsUsername        string                                   `json:"https-username"`
	HttpsServerCertName  string                                   `json:"https-server-cert-name"`
	HttpsPassword        int                                      `json:"https-password"`
	HttpsPasswordString  string                                   `json:"https-password-string"`
	HttpsEncrypted       string                                   `json:"https-encrypted"`
	HttpsKerberosAuth    int                                      `json:"https-kerberos-auth"`
	HttpsKerberosRealm   string                                   `json:"https-kerberos-realm"`
	HttpsKerberosKdc     HealthMonitorMethodHttpsHttpsKerberosKdc `json:"https-kerberos-kdc"`
	CertKeyShared        int                                      `json:"cert-key-shared"`
	Cert                 string                                   `json:"cert"`
	Key                  string                                   `json:"key"`
	KeyPassPhrase        int                                      `json:"key-pass-phrase"`
	KeyPhrase            string                                   `json:"key-phrase"`
	HttpsKeyEncrypted    string                                   `json:"https-key-encrypted"`
	Uuid                 string                                   `json:"uuid"`
}

type HealthMonitorMethodHttpsHttpsKerberosKdc struct {
	HttpsKerberosHostip   string `json:"https-kerberos-hostip"`
	HttpsKerberosHostipv6 string `json:"https-kerberos-hostipv6"`
	HttpsKerberosPort     int    `json:"https-kerberos-port"`
	HttpsKerberosPortv6   int    `json:"https-kerberos-portv6"`
}

type HealthMonitorMethodTacplus struct {
	Tacplus               int    `json:"tacplus"`
	TacplusUsername       string `json:"tacplus-username"`
	TacplusPassword       int    `json:"tacplus-password"`
	TacplusPasswordString string `json:"tacplus-password-string"`
	TacplusEncrypted      string `json:"tacplus-encrypted"`
	TacplusSecret         int    `json:"tacplus-secret"`
	TacplusSecretString   string `json:"tacplus-secret-string"`
	SecretEncrypted       string `json:"secret-encrypted"`
	TacplusPort           int    `json:"tacplus-port" dval:"49"`
	TacplusType           string `json:"tacplus-type" dval:"inbound-ascii-login"`
	Uuid                  string `json:"uuid"`
}

type HealthMonitorMethodCompound struct {
	Compound  int    `json:"compound"`
	RpnString string `json:"rpn-string"`
	Uuid      string `json:"uuid"`
}

func (p *HealthMonitor) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *HealthMonitor) getPath() string {
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
		err = json.Unmarshal(axResp, &p)
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
