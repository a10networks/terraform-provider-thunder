package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type HealthMonitor struct {
	HealthMonitorInstanceName HealthMonitorInstance `json:"monitor,omitempty"`
}

type HealthMonitorInstance struct {
	HealthMonitorInstanceDisableAfterDown           int                         `json:"disable-after-down,omitempty"`
	HealthMonitorInstanceDsrL2Strict                int                         `json:"dsr-l2-strict,omitempty"`
	HealthMonitorInstanceInterval                   int                         `json:"interval,omitempty"`
	HealthMonitorInstanceMethodIcmp                 HealthMonitorInstanceMethod `json:"method,omitempty"`
	HealthMonitorInstanceName                       string                      `json:"name,omitempty"`
	HealthMonitorInstanceOverrideIpv4               string                      `json:"override-ipv4,omitempty"`
	HealthMonitorInstanceOverrideIpv6               string                      `json:"override-ipv6,omitempty"`
	HealthMonitorInstanceOverridePort               int                         `json:"override-port,omitempty"`
	HealthMonitorInstancePassive                    int                         `json:"passive,omitempty"`
	HealthMonitorInstancePassiveInterval            int                         `json:"passive-interval,omitempty"`
	HealthMonitorInstanceRetry                      int                         `json:"retry,omitempty"`
	HealthMonitorInstanceSampleThreshold            int                         `json:"sample-threshold,omitempty"`
	HealthMonitorInstanceSslCiphers                 string                      `json:"ssl-ciphers,omitempty"`
	HealthMonitorInstanceSslDgversion               int                         `json:"ssl-dgversion,omitempty"`
	HealthMonitorInstanceSslTicket                  int                         `json:"ssl-ticket,omitempty"`
	HealthMonitorInstanceSslTicketLifetime          int                         `json:"ssl-ticket-lifetime,omitempty"`
	HealthMonitorInstanceSslVersion                 int                         `json:"ssl-version,omitempty"`
	HealthMonitorInstanceStatusCode                 string                      `json:"status-code,omitempty"`
	HealthMonitorInstanceStrictRetryOnServerErrResp int                         `json:"strict-retry-on-server-err-resp,omitempty"`
	HealthMonitorInstanceThreshold                  int                         `json:"threshold,omitempty"`
	HealthMonitorInstanceTimeout                    int                         `json:"timeout,omitempty"`
	HealthMonitorInstanceUUID                       string                      `json:"uuid,omitempty"`
	HealthMonitorInstanceUpRetry                    int                         `json:"up-retry,omitempty"`
	HealthMonitorInstanceUserTag                    string                      `json:"user-tag,omitempty"`
}

type HealthMonitorInstanceMethod struct {
	HealthMonitorInstanceMethodCompoundCompound       HealthMonitorInstanceMethodCompound    `json:"compound,omitempty"`
	HealthMonitorInstanceMethodDNSDNS                 HealthMonitorInstanceMethodDNS         `json:"dns,omitempty"`
	HealthMonitorInstanceMethodDatabaseDatabase       HealthMonitorInstanceMethodDatabase    `json:"database,omitempty"`
	HealthMonitorInstanceMethodExternalExternal       HealthMonitorInstanceMethodExternal    `json:"external,omitempty"`
	HealthMonitorInstanceMethodFtpFtp                 HealthMonitorInstanceMethodFtp         `json:"ftp,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTP               HealthMonitorInstanceMethodHTTP        `json:"http,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPS             HealthMonitorInstanceMethodHTTPS       `json:"https,omitempty"`
	HealthMonitorInstanceMethodIcmpIcmp               HealthMonitorInstanceMethodIcmp        `json:"icmp,omitempty"`
	HealthMonitorInstanceMethodImapImap               HealthMonitorInstanceMethodImap        `json:"imap,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfg HealthMonitorInstanceMethodKerberosKdc `json:"kerberos-kdc,omitempty"`
	HealthMonitorInstanceMethodLdapLdap               HealthMonitorInstanceMethodLdap        `json:"ldap,omitempty"`
	HealthMonitorInstanceMethodNtpNtp                 HealthMonitorInstanceMethodNtp         `json:"ntp,omitempty"`
	HealthMonitorInstanceMethodPop3Pop3               HealthMonitorInstanceMethodPop3        `json:"pop3,omitempty"`
	HealthMonitorInstanceMethodRadiusRadius           HealthMonitorInstanceMethodRadius      `json:"radius,omitempty"`
	HealthMonitorInstanceMethodRtspRtsp               HealthMonitorInstanceMethodRtsp        `json:"rtsp,omitempty"`
	HealthMonitorInstanceMethodSMTPSMTP               HealthMonitorInstanceMethodSMTP        `json:"smtp,omitempty"`
	HealthMonitorInstanceMethodSipSip                 HealthMonitorInstanceMethodSip         `json:"sip,omitempty"`
	HealthMonitorInstanceMethodSnmpSnmp               HealthMonitorInstanceMethodSnmp        `json:"snmp,omitempty"`
	HealthMonitorInstanceMethodTCPMethodTCP           HealthMonitorInstanceMethodTCP         `json:"tcp,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplus         HealthMonitorInstanceMethodTacplus     `json:"tacplus,omitempty"`
	HealthMonitorInstanceMethodUDPUDP                 HealthMonitorInstanceMethodUDP         `json:"udp,omitempty"`
}

type HealthMonitorInstanceMethodCompound struct {
	HealthMonitorInstanceMethodCompoundCompound  int    `json:"compound,omitempty"`
	HealthMonitorInstanceMethodCompoundRpnString string `json:"rpn-string,omitempty"`
	HealthMonitorInstanceMethodCompoundUUID      string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodDNS struct {
	HealthMonitorInstanceMethodDNSDNS                              int                                           `json:"dns,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomain                        string                                        `json:"dns-domain,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainResponse HealthMonitorInstanceMethodDNSDNSDomainExpect `json:"dns-domain-expect,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainPort                    int                                           `json:"dns-domain-port,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainRecurse                 string                                        `json:"dns-domain-recurse,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainTCP                     int                                           `json:"dns-domain-tcp,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainType                    string                                        `json:"dns-domain-type,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIPKey                         int                                           `json:"dns-ip-key,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv4Addr                      string                                        `json:"dns-ipv4-addr,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv4ExpectDNSIpv4Response     HealthMonitorInstanceMethodDNSDNSIpv4Expect   `json:"dns-ipv4-expect,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv4Port                      int                                           `json:"dns-ipv4-port,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv4Recurse                   string                                        `json:"dns-ipv4-recurse,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv4TCP                       int                                           `json:"dns-ipv4-tcp,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv6Addr                      string                                        `json:"dns-ipv6-addr,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv6ExpectDNSIpv6Response     HealthMonitorInstanceMethodDNSDNSIpv6Expect   `json:"dns-ipv6-expect,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv6Port                      int                                           `json:"dns-ipv6-port,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv6Recurse                   string                                        `json:"dns-ipv6-recurse,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv6TCP                       int                                           `json:"dns-ipv6-tcp,omitempty"`
	HealthMonitorInstanceMethodDNSUUID                             string                                        `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodDatabase struct {
	HealthMonitorInstanceMethodDatabaseDatabase         int    `json:"database,omitempty"`
	HealthMonitorInstanceMethodDatabaseDatabaseName     string `json:"database-name,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbColumn         int    `json:"db-column,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbColumnInteger  int    `json:"db-column-integer,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbName           string `json:"db-name,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbPassword       int    `json:"db-password,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbPasswordStr    string `json:"db-password-str,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbReceive        string `json:"db-receive,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbReceiveInteger int    `json:"db-receive-integer,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbRow            int    `json:"db-row,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbRowInteger     int    `json:"db-row-integer,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbSend           string `json:"db-send,omitempty"`
	HealthMonitorInstanceMethodDatabaseDbUsername       string `json:"db-username,omitempty"`
	HealthMonitorInstanceMethodDatabaseUUID             string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodExternal struct {
	HealthMonitorInstanceMethodExternalExtArguments           string `json:"ext-arguments,omitempty"`
	HealthMonitorInstanceMethodExternalExtPort                int    `json:"ext-port,omitempty"`
	HealthMonitorInstanceMethodExternalExtPreference          int    `json:"ext-preference,omitempty"`
	HealthMonitorInstanceMethodExternalExtProgram             string `json:"ext-program,omitempty"`
	HealthMonitorInstanceMethodExternalExtProgramShared       string `json:"ext-program-shared,omitempty"`
	HealthMonitorInstanceMethodExternalExternal               int    `json:"external,omitempty"`
	HealthMonitorInstanceMethodExternalSharedPartitionProgram int    `json:"shared-partition-program,omitempty"`
	HealthMonitorInstanceMethodExternalUUID                   string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodFtp struct {
	HealthMonitorInstanceMethodFtpFtp               int    `json:"ftp,omitempty"`
	HealthMonitorInstanceMethodFtpFtpPassword       int    `json:"ftp-password,omitempty"`
	HealthMonitorInstanceMethodFtpFtpPasswordString string `json:"ftp-password-string,omitempty"`
	HealthMonitorInstanceMethodFtpFtpPort           int    `json:"ftp-port,omitempty"`
	HealthMonitorInstanceMethodFtpFtpUsername       string `json:"ftp-username,omitempty"`
	HealthMonitorInstanceMethodFtpUUID              string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodHTTP struct {
	HealthMonitorInstanceMethodHTTPHTTP                              int                                            `json:"http,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPExpect                        int                                            `json:"http-expect,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPHost                          string                                         `json:"http-host,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPKerberosAuth                  int                                            `json:"http-kerberos-auth,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosHostip HealthMonitorInstanceMethodHTTPHTTPKerberosKdc `json:"http-kerberos-kdc,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPKerberosRealm                 string                                         `json:"http-kerberos-realm,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPMaintenanceCode               string                                         `json:"http-maintenance-code,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPPassword                      int                                            `json:"http-password,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPPasswordString                string                                         `json:"http-password-string,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPPort                          int                                            `json:"http-port,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPPostdata                      string                                         `json:"http-postdata,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPPostfile                      string                                         `json:"http-postfile,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPResponseCode                  string                                         `json:"http-response-code,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPText                          string                                         `json:"http-text,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPURL                           int                                            `json:"http-url,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPUsername                      string                                         `json:"http-username,omitempty"`
	HealthMonitorInstanceMethodHTTPMaintenance                       int                                            `json:"maintenance,omitempty"`
	HealthMonitorInstanceMethodHTTPMaintenanceText                   string                                         `json:"maintenance-text,omitempty"`
	HealthMonitorInstanceMethodHTTPMaintenanceTextRegex              string                                         `json:"maintenance-text-regex,omitempty"`
	HealthMonitorInstanceMethodHTTPPostPath                          string                                         `json:"post-path,omitempty"`
	HealthMonitorInstanceMethodHTTPPostType                          string                                         `json:"post-type,omitempty"`
	HealthMonitorInstanceMethodHTTPResponseCodeRegex                 string                                         `json:"response-code-regex,omitempty"`
	HealthMonitorInstanceMethodHTTPTextRegex                         string                                         `json:"text-regex,omitempty"`
	HealthMonitorInstanceMethodHTTPURLPath                           string                                         `json:"url-path,omitempty"`
	HealthMonitorInstanceMethodHTTPURLType                           string                                         `json:"url-type,omitempty"`
	HealthMonitorInstanceMethodHTTPUUID                              string                                         `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodHTTPS struct {
	HealthMonitorInstanceMethodHTTPSCert                                string                                           `json:"cert,omitempty"`
	HealthMonitorInstanceMethodHTTPSCertKeyShared                       int                                              `json:"cert-key-shared,omitempty"`
	HealthMonitorInstanceMethodHTTPSDisableSslv2Hello                   int                                              `json:"disable-sslv2hello,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPS                               int                                              `json:"https,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSExpect                         int                                              `json:"https-expect,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSHost                           string                                           `json:"https-host,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSKerberosAuth                   int                                              `json:"https-kerberos-auth,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosHostip HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdc `json:"https-kerberos-kdc,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSKerberosRealm                  string                                           `json:"https-kerberos-realm,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSMaintenanceCode                string                                           `json:"https-maintenance-code,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSPassword                       int                                              `json:"https-password,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSPasswordString                 string                                           `json:"https-password-string,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSPostdata                       string                                           `json:"https-postdata,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSPostfile                       string                                           `json:"https-postfile,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSResponseCode                   string                                           `json:"https-response-code,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSServerCertName                 string                                           `json:"https-server-cert-name,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSText                           string                                           `json:"https-text,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSURL                            int                                              `json:"https-url,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSUsername                       string                                           `json:"https-username,omitempty"`
	HealthMonitorInstanceMethodHTTPSKey                                 string                                           `json:"key,omitempty"`
	HealthMonitorInstanceMethodHTTPSKeyPassPhrase                       int                                              `json:"key-pass-phrase,omitempty"`
	HealthMonitorInstanceMethodHTTPSKeyPhrase                           string                                           `json:"key-phrase,omitempty"`
	HealthMonitorInstanceMethodHTTPSMaintenance                         int                                              `json:"maintenance,omitempty"`
	HealthMonitorInstanceMethodHTTPSMaintenanceText                     string                                           `json:"maintenance-text,omitempty"`
	HealthMonitorInstanceMethodHTTPSMaintenanceTextRegex                string                                           `json:"maintenance-text-regex,omitempty"`
	HealthMonitorInstanceMethodHTTPSPostPath                            string                                           `json:"post-path,omitempty"`
	HealthMonitorInstanceMethodHTTPSPostType                            string                                           `json:"post-type,omitempty"`
	HealthMonitorInstanceMethodHTTPSResponseCodeRegex                   string                                           `json:"response-code-regex,omitempty"`
	HealthMonitorInstanceMethodHTTPSSni                                 int                                              `json:"sni,omitempty"`
	HealthMonitorInstanceMethodHTTPSTextRegex                           string                                           `json:"text-regex,omitempty"`
	HealthMonitorInstanceMethodHTTPSURLPath                             string                                           `json:"url-path,omitempty"`
	HealthMonitorInstanceMethodHTTPSURLType                             string                                           `json:"url-type,omitempty"`
	HealthMonitorInstanceMethodHTTPSUUID                                string                                           `json:"uuid,omitempty"`
	HealthMonitorInstanceMethodHTTPSWebPort                             int                                              `json:"web-port,omitempty"`
}

type HealthMonitorInstanceMethodIcmp struct {
	HealthMonitorInstanceMethodIcmpIP          string `json:"ip,omitempty"`
	HealthMonitorInstanceMethodIcmpIcmp        int    `json:"icmp,omitempty"`
	HealthMonitorInstanceMethodIcmpIpv6        string `json:"ipv6,omitempty"`
	HealthMonitorInstanceMethodIcmpTransparent int    `json:"transparent,omitempty"`
	HealthMonitorInstanceMethodIcmpUUID        string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodImap struct {
	HealthMonitorInstanceMethodImapImap               int    `json:"imap,omitempty"`
	HealthMonitorInstanceMethodImapImapCramMd5        int    `json:"imap-cram-md5,omitempty"`
	HealthMonitorInstanceMethodImapImapLogin          int    `json:"imap-login,omitempty"`
	HealthMonitorInstanceMethodImapImapPassword       int    `json:"imap-password,omitempty"`
	HealthMonitorInstanceMethodImapImapPasswordString string `json:"imap-password-string,omitempty"`
	HealthMonitorInstanceMethodImapImapPlain          int    `json:"imap-plain,omitempty"`
	HealthMonitorInstanceMethodImapImapPort           int    `json:"imap-port,omitempty"`
	HealthMonitorInstanceMethodImapImapUsername       string `json:"imap-username,omitempty"`
	HealthMonitorInstanceMethodImapPwdAuth            int    `json:"pwd-auth,omitempty"`
	HealthMonitorInstanceMethodImapUUID               string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodKerberosKdc struct {
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinit HealthMonitorInstanceMethodKerberosKdcKerberosCfg `json:"kerberos-cfg,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcUUID             string                                            `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodLdap struct {
	HealthMonitorInstanceMethodLdapAcceptNotFound     int    `json:"AcceptNotFound,omitempty"`
	HealthMonitorInstanceMethodLdapAcceptResRef       int    `json:"AcceptResRef,omitempty"`
	HealthMonitorInstanceMethodLdapBaseDN             string `json:"BaseDN,omitempty"`
	HealthMonitorInstanceMethodLdapLdap               int    `json:"ldap,omitempty"`
	HealthMonitorInstanceMethodLdapLdapBinddn         string `json:"ldap-binddn,omitempty"`
	HealthMonitorInstanceMethodLdapLdapPassword       int    `json:"ldap-password,omitempty"`
	HealthMonitorInstanceMethodLdapLdapPasswordString string `json:"ldap-password-string,omitempty"`
	HealthMonitorInstanceMethodLdapLdapPort           int    `json:"ldap-port,omitempty"`
	HealthMonitorInstanceMethodLdapLdapQuery          string `json:"ldap-query,omitempty"`
	HealthMonitorInstanceMethodLdapLdapRunSearch      int    `json:"ldap-run-search,omitempty"`
	HealthMonitorInstanceMethodLdapLdapSecurity       string `json:"ldap-security,omitempty"`
	HealthMonitorInstanceMethodLdapUUID               string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodNtp struct {
	HealthMonitorInstanceMethodNtpNtp     int    `json:"ntp,omitempty"`
	HealthMonitorInstanceMethodNtpNtpPort int    `json:"ntp-port,omitempty"`
	HealthMonitorInstanceMethodNtpUUID    string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodPop3 struct {
	HealthMonitorInstanceMethodPop3Pop3               int    `json:"pop3,omitempty"`
	HealthMonitorInstanceMethodPop3Pop3Password       int    `json:"pop3-password,omitempty"`
	HealthMonitorInstanceMethodPop3Pop3PasswordString string `json:"pop3-password-string,omitempty"`
	HealthMonitorInstanceMethodPop3Pop3Port           int    `json:"pop3-port,omitempty"`
	HealthMonitorInstanceMethodPop3Pop3Username       string `json:"pop3-username,omitempty"`
	HealthMonitorInstanceMethodPop3UUID               string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodRadius struct {
	HealthMonitorInstanceMethodRadiusRadius               int    `json:"radius,omitempty"`
	HealthMonitorInstanceMethodRadiusRadiusExpect         int    `json:"radius-expect,omitempty"`
	HealthMonitorInstanceMethodRadiusRadiusPasswordString string `json:"radius-password-string,omitempty"`
	HealthMonitorInstanceMethodRadiusRadiusPort           int    `json:"radius-port,omitempty"`
	HealthMonitorInstanceMethodRadiusRadiusResponseCode   string `json:"radius-response-code,omitempty"`
	HealthMonitorInstanceMethodRadiusRadiusSecret         string `json:"radius-secret,omitempty"`
	HealthMonitorInstanceMethodRadiusRadiusUsername       string `json:"radius-username,omitempty"`
	HealthMonitorInstanceMethodRadiusUUID                 string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodRtsp struct {
	HealthMonitorInstanceMethodRtspRtsp     int    `json:"rtsp,omitempty"`
	HealthMonitorInstanceMethodRtspRtspPort int    `json:"rtsp-port,omitempty"`
	HealthMonitorInstanceMethodRtspRtspurl  string `json:"rtspurl,omitempty"`
	HealthMonitorInstanceMethodRtspUUID     string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodSMTP struct {
	HealthMonitorInstanceMethodSMTPMailFrom     string `json:"mail-from,omitempty"`
	HealthMonitorInstanceMethodSMTPRcptTo       string `json:"rcpt-to,omitempty"`
	HealthMonitorInstanceMethodSMTPSMTP         int    `json:"smtp,omitempty"`
	HealthMonitorInstanceMethodSMTPSMTPDomain   string `json:"smtp-domain,omitempty"`
	HealthMonitorInstanceMethodSMTPSMTPPort     int    `json:"smtp-port,omitempty"`
	HealthMonitorInstanceMethodSMTPSMTPStarttls int    `json:"smtp-starttls,omitempty"`
	HealthMonitorInstanceMethodSMTPUUID         string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodSip struct {
	HealthMonitorInstanceMethodSipExpectResponseCode string `json:"expect-response-code,omitempty"`
	HealthMonitorInstanceMethodSipRegister           int    `json:"register,omitempty"`
	HealthMonitorInstanceMethodSipSip                int    `json:"sip,omitempty"`
	HealthMonitorInstanceMethodSipSipPort            int    `json:"sip-port,omitempty"`
	HealthMonitorInstanceMethodSipSipTCP             int    `json:"sip-tcp,omitempty"`
	HealthMonitorInstanceMethodSipUUID               string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodSnmp struct {
	HealthMonitorInstanceMethodSnmpCommunity         string                                   `json:"community,omitempty"`
	HealthMonitorInstanceMethodSnmpOidMib            HealthMonitorInstanceMethodSnmpOid       `json:"oid,omitempty"`
	HealthMonitorInstanceMethodSnmpOperationOperType HealthMonitorInstanceMethodSnmpOperation `json:"operation,omitempty"`
	HealthMonitorInstanceMethodSnmpSnmp              int                                      `json:"snmp,omitempty"`
	HealthMonitorInstanceMethodSnmpSnmpPort          int                                      `json:"snmp-port,omitempty"`
	HealthMonitorInstanceMethodSnmpUUID              string                                   `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodTCP struct {
	HealthMonitorInstanceMethodTCPMaintenance          int                                    `json:"maintenance,omitempty"`
	HealthMonitorInstanceMethodTCPMaintenanceText      string                                 `json:"maintenance-text,omitempty"`
	HealthMonitorInstanceMethodTCPMethodTCP            int                                    `json:"method-tcp,omitempty"`
	HealthMonitorInstanceMethodTCPPortHalfopen         int                                    `json:"port-halfopen,omitempty"`
	HealthMonitorInstanceMethodTCPPortRespPortContains HealthMonitorInstanceMethodTCPPortResp `json:"port-resp,omitempty"`
	HealthMonitorInstanceMethodTCPPortSend             string                                 `json:"port-send,omitempty"`
	HealthMonitorInstanceMethodTCPTCPPort              int                                    `json:"tcp-port,omitempty"`
	HealthMonitorInstanceMethodTCPUUID                 string                                 `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodTacplus struct {
	HealthMonitorInstanceMethodTacplusTacplus               int    `json:"tacplus,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplusPassword       int    `json:"tacplus-password,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplusPasswordString string `json:"tacplus-password-string,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplusPort           int    `json:"tacplus-port,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplusSecret         int    `json:"tacplus-secret,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplusSecretString   string `json:"tacplus-secret-string,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplusType           string `json:"tacplus-type,omitempty"`
	HealthMonitorInstanceMethodTacplusTacplusUsername       string `json:"tacplus-username,omitempty"`
	HealthMonitorInstanceMethodTacplusUUID                  string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodUDP struct {
	HealthMonitorInstanceMethodUDPForceUpWithSingleHealthcheck int    `json:"force-up-with-single-healthcheck,omitempty"`
	HealthMonitorInstanceMethodUDPUDP                          int    `json:"udp,omitempty"`
	HealthMonitorInstanceMethodUDPUDPPort                      int    `json:"udp-port,omitempty"`
	HealthMonitorInstanceMethodUDPUUID                         string `json:"uuid,omitempty"`
}

type HealthMonitorInstanceMethodDNSDNSDomainExpect struct {
	HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainFqdn     string `json:"dns-domain-fqdn,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainIpv4     string `json:"dns-domain-ipv4,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainIpv6     string `json:"dns-domain-ipv6,omitempty"`
	HealthMonitorInstanceMethodDNSDNSDomainExpectDNSDomainResponse string `json:"dns-domain-response,omitempty"`
}

type HealthMonitorInstanceMethodDNSDNSIpv4Expect struct {
	HealthMonitorInstanceMethodDNSDNSIpv4ExpectDNSIpv4Fqdn     string `json:"dns-ipv4-fqdn,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv4ExpectDNSIpv4Response string `json:"dns-ipv4-response,omitempty"`
}

type HealthMonitorInstanceMethodDNSDNSIpv6Expect struct {
	HealthMonitorInstanceMethodDNSDNSIpv6ExpectDNSIpv6Fqdn     string `json:"dns-ipv6-fqdn,omitempty"`
	HealthMonitorInstanceMethodDNSDNSIpv6ExpectDNSIpv6Response string `json:"dns-ipv6-response,omitempty"`
}

type HealthMonitorInstanceMethodHTTPHTTPKerberosKdc struct {
	HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosHostip   string `json:"http-kerberos-hostip,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosHostipv6 string `json:"http-kerberos-hostipv6,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosPort     int    `json:"http-kerberos-port,omitempty"`
	HealthMonitorInstanceMethodHTTPHTTPKerberosKdcHTTPKerberosPortv6   int    `json:"http-kerberos-portv6,omitempty"`
}

type HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdc struct {
	HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosHostip   string `json:"https-kerberos-hostip,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosHostipv6 string `json:"https-kerberos-hostipv6,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosPort     int    `json:"https-kerberos-port,omitempty"`
	HealthMonitorInstanceMethodHTTPSHTTPSKerberosKdcHTTPSKerberosPortv6   int    `json:"https-kerberos-portv6,omitempty"`
}

type HealthMonitorInstanceMethodKerberosKdcKerberosCfg struct {
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadmin              int    `json:"kadmin,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminKdc           string `json:"kadmin-kdc,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminPassword      string `json:"kadmin-password,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminPricipalName  string `json:"kadmin-pricipal-name,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminRealm         string `json:"kadmin-realm,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKadminServer        string `json:"kadmin-server,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinit               int    `json:"kinit,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinitKdc            string `json:"kinit-kdc,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinitPassword       string `json:"kinit-password,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKinitPricipalName   string `json:"kinit-pricipal-name,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswd             int    `json:"kpasswd,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdKdc          string `json:"kpasswd-kdc,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdPassword     string `json:"kpasswd-password,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdPricipalName string `json:"kpasswd-pricipal-name,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgKpasswdServer       string `json:"kpasswd-server,omitempty"`
	HealthMonitorInstanceMethodKerberosKdcKerberosCfgTCPOnly             int    `json:"tcp-only,omitempty"`
}

type HealthMonitorInstanceMethodSnmpOid struct {
	HealthMonitorInstanceMethodSnmpOidAsn string `json:"asn,omitempty"`
	HealthMonitorInstanceMethodSnmpOidMib string `json:"mib,omitempty"`
}

type HealthMonitorInstanceMethodSnmpOperation struct {
	HealthMonitorInstanceMethodSnmpOperationOperType string `json:"oper-type,omitempty"`
}

type HealthMonitorInstanceMethodTCPPortResp struct {
	HealthMonitorInstanceMethodTCPPortRespPortContains string `json:"port-contains,omitempty"`
}

func PostHealthMonitor(id string, inst HealthMonitor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostHealthMonitor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/health/monitor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HealthMonitor
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostHealthMonitor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetHealthMonitor(id string, name1 string, host string) (*HealthMonitor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetHealthMonitor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/health/monitor/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HealthMonitor
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetHealthMonitor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutHealthMonitor(id string, name1 string, inst HealthMonitor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutHealthMonitor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/health/monitor/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HealthMonitor
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutHealthMonitor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteHealthMonitor(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteHealthMonitor")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/health/monitor/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HealthMonitor
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteHealthMonitor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
