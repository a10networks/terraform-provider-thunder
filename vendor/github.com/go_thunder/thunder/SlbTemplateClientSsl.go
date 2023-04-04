package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateClientSsl struct {
	SlbTemplateClientSslInstanceName SlbTemplateClientSslInstance `json:"client-ssl,omitempty"`
}

type SlbTemplateClientSslInstance struct {
	SlbTemplateClientSslInstanceAdGroupList                                                        string                                                         `json:"ad-group-list,omitempty"`
	SlbTemplateClientSslInstanceAlertType                                                          string                                                         `json:"alert-type,omitempty"`
	SlbTemplateClientSslInstanceAuthSg                                                             string                                                         `json:"auth-sg,omitempty"`
	SlbTemplateClientSslInstanceAuthSgDn                                                           int                                                            `json:"auth-sg-dn,omitempty"`
	SlbTemplateClientSslInstanceAuthSgFilter                                                       string                                                         `json:"auth-sg-filter,omitempty"`
	SlbTemplateClientSslInstanceAuthUsername                                                       string                                                         `json:"auth-username,omitempty"`
	SlbTemplateClientSslInstanceAuthUsernameAttribute                                              string                                                         `json:"auth-username-attribute,omitempty"`
	SlbTemplateClientSslInstanceAuthenName                                                         string                                                         `json:"authen-name,omitempty"`
	SlbTemplateClientSslInstanceAuthorization                                                      int                                                            `json:"authorization,omitempty"`
	SlbTemplateClientSslInstanceBypassCertIssuerClassListName                                      string                                                         `json:"bypass-cert-issuer-class-list-name,omitempty"`
	SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListBypassCertIssuerMultiClassListName   []SlbTemplateClientSslInstanceBypassCertIssuerMultiClassList   `json:"bypass-cert-issuer-multi-class-list,omitempty"`
	SlbTemplateClientSslInstanceBypassCertSanClassListName                                         string                                                         `json:"bypass-cert-san-class-list-name,omitempty"`
	SlbTemplateClientSslInstanceBypassCertSanMultiClassListBypassCertSanMultiClassListName         []SlbTemplateClientSslInstanceBypassCertSanMultiClassList      `json:"bypass-cert-san-multi-class-list,omitempty"`
	SlbTemplateClientSslInstanceBypassCertSubjectClassListName                                     string                                                         `json:"bypass-cert-subject-class-list-name,omitempty"`
	SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListBypassCertSubjectMultiClassListName []SlbTemplateClientSslInstanceBypassCertSubjectMultiClassList  `json:"bypass-cert-subject-multi-class-list,omitempty"`
	SlbTemplateClientSslInstanceCaCertsCaCert                                                      []SlbTemplateClientSslInstanceCaCerts                          `json:"ca-certs,omitempty"`
	SlbTemplateClientSslInstanceCachePersistenceListName                                           string                                                         `json:"cache-persistence-list-name,omitempty"`
	SlbTemplateClientSslInstanceCaseInsensitive                                                    int                                                            `json:"case-insensitive,omitempty"`
	SlbTemplateClientSslInstanceCentralCertPinList                                                 int                                                            `json:"central-cert-pin-list,omitempty"`
	SlbTemplateClientSslInstanceCertRevokeAction                                                   string                                                         `json:"cert-revoke-action,omitempty"`
	SlbTemplateClientSslInstanceCertUnknownAction                                                  string                                                         `json:"cert-unknown-action,omitempty"`
	SlbTemplateClientSslInstanceCertificateIssuerContainsListCertificateIssuerContains             []SlbTemplateClientSslInstanceCertificateIssuerContainsList    `json:"certificate-issuer-contains-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCertificateIssuerEndsWith             []SlbTemplateClientSslInstanceCertificateIssuerEndsWithList    `json:"certificate-issuer-ends-with-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateIssuerEqualsListCertificateIssuerEquals                 []SlbTemplateClientSslInstanceCertificateIssuerEqualsList      `json:"certificate-issuer-equals-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCertificateIssuerStarts             []SlbTemplateClientSslInstanceCertificateIssuerStartsWithList  `json:"certificate-issuer-starts-with-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateListCert                                                []SlbTemplateClientSslInstanceCertificateList                  `json:"certificate-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSanContainsListCertificateSanContains                   []SlbTemplateClientSslInstanceCertificateSanContainsList       `json:"certificate-san-contains-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSanEndsWithListCertificateSanEndsWith                   []SlbTemplateClientSslInstanceCertificateSanEndsWithList       `json:"certificate-san-ends-with-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSanEqualsListCertificateSanEquals                       []SlbTemplateClientSslInstanceCertificateSanEqualsList         `json:"certificate-san-equals-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSanStartsWithListCertificateSanStarts                   []SlbTemplateClientSslInstanceCertificateSanStartsWithList     `json:"certificate-san-starts-with-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSubjectContainsListCertificateSubjectContains           []SlbTemplateClientSslInstanceCertificateSubjectContainsList   `json:"certificate-subject-contains-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCertificateSubjectEndsWith           []SlbTemplateClientSslInstanceCertificateSubjectEndsWithList   `json:"certificate-subject-ends-with-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSubjectEqualsListCertificateSubjectEquals               []SlbTemplateClientSslInstanceCertificateSubjectEqualsList     `json:"certificate-subject-equals-list,omitempty"`
	SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCertificateSubjectStarts           []SlbTemplateClientSslInstanceCertificateSubjectStartsWithList `json:"certificate-subject-starts-with-list,omitempty"`
	SlbTemplateClientSslInstanceChainCert                                                          string                                                         `json:"chain-cert,omitempty"`
	SlbTemplateClientSslInstanceChainCertSharedStr                                                 string                                                         `json:"chain-cert-shared-str,omitempty"`
	SlbTemplateClientSslInstanceCipherWithoutPrioListCipherWoPrio                                  []SlbTemplateClientSslInstanceCipherWithoutPrioList            `json:"cipher-without-prio-list,omitempty"`
	SlbTemplateClientSslInstanceClassListName                                                      string                                                         `json:"class-list-name,omitempty"`
	SlbTemplateClientSslInstanceClientAuthCaseInsensitive                                          int                                                            `json:"client-auth-case-insensitive,omitempty"`
	SlbTemplateClientSslInstanceClientAuthClassList                                                string                                                         `json:"client-auth-class-list,omitempty"`
	SlbTemplateClientSslInstanceClientAuthContainsListClientAuthContains                           []SlbTemplateClientSslInstanceClientAuthContainsList           `json:"client-auth-contains-list,omitempty"`
	SlbTemplateClientSslInstanceClientAuthEndsWithListClientAuthEndsWith                           []SlbTemplateClientSslInstanceClientAuthEndsWithList           `json:"client-auth-ends-with-list,omitempty"`
	SlbTemplateClientSslInstanceClientAuthEqualsListClientAuthEquals                               []SlbTemplateClientSslInstanceClientAuthEqualsList             `json:"client-auth-equals-list,omitempty"`
	SlbTemplateClientSslInstanceClientAuthStartsWithListClientAuthStartsWith                       []SlbTemplateClientSslInstanceClientAuthStartsWithList         `json:"client-auth-starts-with-list,omitempty"`
	SlbTemplateClientSslInstanceClientCertificate                                                  string                                                         `json:"client-certificate,omitempty"`
	SlbTemplateClientSslInstanceClientIpv4ListClientIpv4ListName                                   []SlbTemplateClientSslInstanceClientIpv4List                   `json:"client-ipv4-list,omitempty"`
	SlbTemplateClientSslInstanceClientIpv6ListClientIpv6ListName                                   []SlbTemplateClientSslInstanceClientIpv6List                   `json:"client-ipv6-list,omitempty"`
	SlbTemplateClientSslInstanceCloseNotify                                                        int                                                            `json:"close-notify,omitempty"`
	SlbTemplateClientSslInstanceContainsListContains                                               []SlbTemplateClientSslInstanceContainsList                     `json:"contains-list,omitempty"`
	SlbTemplateClientSslInstanceCrlCertsCrl                                                        []SlbTemplateClientSslInstanceCrlCerts                         `json:"crl-certs,omitempty"`
	SlbTemplateClientSslInstanceDgversion                                                          int                                                            `json:"dgversion,omitempty"`
	SlbTemplateClientSslInstanceDhType                                                             string                                                         `json:"dh-type,omitempty"`
	SlbTemplateClientSslInstanceDirectClientServerAuth                                             int                                                            `json:"direct-client-server-auth,omitempty"`
	SlbTemplateClientSslInstanceDisableSslv3                                                       int                                                            `json:"disable-sslv3,omitempty"`
	SlbTemplateClientSslInstanceEarlyData                                                          int                                                            `json:"early-data,omitempty"`
	SlbTemplateClientSslInstanceEcListEc                                                           []SlbTemplateClientSslInstanceEcList                           `json:"ec-list,omitempty"`
	SlbTemplateClientSslInstanceEnableSsliFtpAlg                                                   int                                                            `json:"enable-ssli-ftp-alg,omitempty"`
	SlbTemplateClientSslInstanceEnableTLSAlertLogging                                              int                                                            `json:"enable-tls-alert-logging,omitempty"`
	SlbTemplateClientSslInstanceEndsWithListEndsWith                                               []SlbTemplateClientSslInstanceEndsWithList                     `json:"ends-with-list,omitempty"`
	SlbTemplateClientSslInstanceEqualsListEquals                                                   []SlbTemplateClientSslInstanceEqualsList                       `json:"equals-list,omitempty"`
	SlbTemplateClientSslInstanceExceptionAdGroupList                                               string                                                         `json:"exception-ad-group-list,omitempty"`
	SlbTemplateClientSslInstanceExceptionCertificateIssuerClName                                   string                                                         `json:"exception-certificate-issuer-cl-name,omitempty"`
	SlbTemplateClientSslInstanceExceptionCertificateSanClName                                      string                                                         `json:"exception-certificate-san-cl-name,omitempty"`
	SlbTemplateClientSslInstanceExceptionCertificateSubjectClName                                  string                                                         `json:"exception-certificate-subject-cl-name,omitempty"`
	SlbTemplateClientSslInstanceExceptionClientIpv4ListExceptionClientIpv4ListName                 []SlbTemplateClientSslInstanceExceptionClientIpv4List          `json:"exception-client-ipv4-list,omitempty"`
	SlbTemplateClientSslInstanceExceptionClientIpv6ListExceptionClientIpv6ListName                 []SlbTemplateClientSslInstanceExceptionClientIpv6List          `json:"exception-client-ipv6-list,omitempty"`
	SlbTemplateClientSslInstanceExceptionServerIpv4ListExceptionServerIpv4ListName                 []SlbTemplateClientSslInstanceExceptionServerIpv4List          `json:"exception-server-ipv4-list,omitempty"`
	SlbTemplateClientSslInstanceExceptionServerIpv6ListExceptionServerIpv6ListName                 []SlbTemplateClientSslInstanceExceptionServerIpv6List          `json:"exception-server-ipv6-list,omitempty"`
	SlbTemplateClientSslInstanceExceptionSniClName                                                 string                                                         `json:"exception-sni-cl-name,omitempty"`
	SlbTemplateClientSslInstanceExceptionUserNameList                                              string                                                         `json:"exception-user-name-list,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionUncategorized                         SlbTemplateClientSslInstanceExceptionWebCategory               `json:"exception-web-category,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebReputationExceptionTrustworthy                         SlbTemplateClientSslInstanceExceptionWebReputation             `json:"exception-web-reputation,omitempty"`
	SlbTemplateClientSslInstanceExpireHours                                                        int                                                            `json:"expire-hours,omitempty"`
	SlbTemplateClientSslInstanceForwardPassphrase                                                  string                                                         `json:"forward-passphrase,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyAltSign                                                int                                                            `json:"forward-proxy-alt-sign,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyBlockMessage                                           string                                                         `json:"forward-proxy-block-message,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCaCert                                                 string                                                         `json:"forward-proxy-ca-cert,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCaKey                                                  string                                                         `json:"forward-proxy-ca-key,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCertCacheLimit                                         int                                                            `json:"forward-proxy-cert-cache-limit,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCertCacheTimeout                                       int                                                            `json:"forward-proxy-cert-cache-timeout,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCertExpiry                                             int                                                            `json:"forward-proxy-cert-expiry,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCertNotReadyAction                                     string                                                         `json:"forward-proxy-cert-not-ready-action,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCertRevokeAction                                       int                                                            `json:"forward-proxy-cert-revoke-action,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCertUnknownAction                                      int                                                            `json:"forward-proxy-cert-unknown-action,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyCrlDisable                                             int                                                            `json:"forward-proxy-crl-disable,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyDecryptedDscp                                          int                                                            `json:"forward-proxy-decrypted-dscp,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyDecryptedDscpBypass                                    int                                                            `json:"forward-proxy-decrypted-dscp-bypass,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyEnable                                                 int                                                            `json:"forward-proxy-enable,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyEsniAction                                             int                                                            `json:"forward-proxy-esni-action,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyFailsafeDisable                                        int                                                            `json:"forward-proxy-failsafe-disable,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyHashPersistenceInterval                                int                                                            `json:"forward-proxy-hash-persistence-interval,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyLogDisable                                             int                                                            `json:"forward-proxy-log-disable,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyNoSharedCipherAction                                   int                                                            `json:"forward-proxy-no-shared-cipher-action,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyNoSniAction                                            string                                                         `json:"forward-proxy-no-sni-action,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyOcspDisable                                            int                                                            `json:"forward-proxy-ocsp-disable,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyRequireSniCertMatched                                  string                                                         `json:"forward-proxy-require-sni-cert-matched,omitempty"`
	SlbTemplateClientSslInstanceForwardProxySelfsignRedir                                          int                                                            `json:"forward-proxy-selfsign-redir,omitempty"`
	SlbTemplateClientSslInstanceForwardProxySslVersion                                             int                                                            `json:"forward-proxy-ssl-version,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyTrustedCaListsForwardProxyTrustedCa                    []SlbTemplateClientSslInstanceForwardProxyTrustedCaLists       `json:"forward-proxy-trusted-ca-lists,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyVerifyCertFailAction                                   int                                                            `json:"forward-proxy-verify-cert-fail-action,omitempty"`
	SlbTemplateClientSslInstanceFpAltCert                                                          string                                                         `json:"fp-alt-cert,omitempty"`
	SlbTemplateClientSslInstanceFpAltChainCert                                                     string                                                         `json:"fp-alt-chain-cert,omitempty"`
	SlbTemplateClientSslInstanceFpAltKey                                                           string                                                         `json:"fp-alt-key,omitempty"`
	SlbTemplateClientSslInstanceFpAltPassphrase                                                    string                                                         `json:"fp-alt-passphrase,omitempty"`
	SlbTemplateClientSslInstanceFpAltShared                                                        int                                                            `json:"fp-alt-shared,omitempty"`
	SlbTemplateClientSslInstanceFpCaCertificate                                                    string                                                         `json:"fp-ca-certificate,omitempty"`
	SlbTemplateClientSslInstanceFpCaCertificateShared                                              int                                                            `json:"fp-ca-certificate-shared,omitempty"`
	SlbTemplateClientSslInstanceFpCaChainCert                                                      string                                                         `json:"fp-ca-chain-cert,omitempty"`
	SlbTemplateClientSslInstanceFpCaKey                                                            string                                                         `json:"fp-ca-key,omitempty"`
	SlbTemplateClientSslInstanceFpCaKeyPassphrase                                                  string                                                         `json:"fp-ca-key-passphrase,omitempty"`
	SlbTemplateClientSslInstanceFpCaKeyShared                                                      int                                                            `json:"fp-ca-key-shared,omitempty"`
	SlbTemplateClientSslInstanceFpCaShared                                                         int                                                            `json:"fp-ca-shared,omitempty"`
	SlbTemplateClientSslInstanceFpCertExtAiaCaIssuers                                              string                                                         `json:"fp-cert-ext-aia-ca-issuers,omitempty"`
	SlbTemplateClientSslInstanceFpCertExtAiaOcsp                                                   string                                                         `json:"fp-cert-ext-aia-ocsp,omitempty"`
	SlbTemplateClientSslInstanceFpCertExtCrldp                                                     string                                                         `json:"fp-cert-ext-crldp,omitempty"`
	SlbTemplateClientSslInstanceFpCertFetchAutonat                                                 string                                                         `json:"fp-cert-fetch-autonat,omitempty"`
	SlbTemplateClientSslInstanceFpCertFetchAutonatPrecedence                                       int                                                            `json:"fp-cert-fetch-autonat-precedence,omitempty"`
	SlbTemplateClientSslInstanceFpCertFetchNatpoolName                                             string                                                         `json:"fp-cert-fetch-natpool-name,omitempty"`
	SlbTemplateClientSslInstanceFpCertFetchNatpoolNameShared                                       string                                                         `json:"fp-cert-fetch-natpool-name-shared,omitempty"`
	SlbTemplateClientSslInstanceFpCertFetchNatpoolPrecedence                                       int                                                            `json:"fp-cert-fetch-natpool-precedence,omitempty"`
	SlbTemplateClientSslInstanceFpEsniAction                                                       string                                                         `json:"fp-esni-action,omitempty"`
	SlbTemplateClientSslInstanceHandshakeLoggingEnable                                             int                                                            `json:"handshake-logging-enable,omitempty"`
	SlbTemplateClientSslInstanceHsmType                                                            string                                                         `json:"hsm-type,omitempty"`
	SlbTemplateClientSslInstanceInspectCertificateIssuerClName                                     string                                                         `json:"inspect-certificate-issuer-cl-name,omitempty"`
	SlbTemplateClientSslInstanceInspectCertificateSanClName                                        string                                                         `json:"inspect-certificate-san-cl-name,omitempty"`
	SlbTemplateClientSslInstanceInspectCertificateSubjectClName                                    string                                                         `json:"inspect-certificate-subject-cl-name,omitempty"`
	SlbTemplateClientSslInstanceInspectListName                                                    string                                                         `json:"inspect-list-name,omitempty"`
	SlbTemplateClientSslInstanceJa3Enable                                                          int                                                            `json:"ja3-enable,omitempty"`
	SlbTemplateClientSslInstanceJa3InsertHTTPHeader                                                string                                                         `json:"ja3-insert-http-header,omitempty"`
	SlbTemplateClientSslInstanceJa3RejectClassList                                                 string                                                         `json:"ja3-reject-class-list,omitempty"`
	SlbTemplateClientSslInstanceJa3RejectMaxNumberPerHost                                          int                                                            `json:"ja3-reject-max-number-per-host,omitempty"`
	SlbTemplateClientSslInstanceJa3TTL                                                             int                                                            `json:"ja3-ttl,omitempty"`
	SlbTemplateClientSslInstanceLdapBaseDnFromCert                                                 int                                                            `json:"ldap-base-dn-from-cert,omitempty"`
	SlbTemplateClientSslInstanceLdapSearchFilter                                                   string                                                         `json:"ldap-search-filter,omitempty"`
	SlbTemplateClientSslInstanceLocalCertPinListLocalCertPinListBypassFailCount                    SlbTemplateClientSslInstanceLocalCertPinList                   `json:"local-cert-pin-list,omitempty"`
	SlbTemplateClientSslInstanceLocalLogging                                                       int                                                            `json:"local-logging,omitempty"`
	SlbTemplateClientSslInstanceMultiClassListMultiClistName                                       []SlbTemplateClientSslInstanceMultiClassList                   `json:"multi-class-list,omitempty"`
	SlbTemplateClientSslInstanceName                                                               string                                                         `json:"name,omitempty"`
	SlbTemplateClientSslInstanceNoAntiReplay                                                       int                                                            `json:"no-anti-replay,omitempty"`
	SlbTemplateClientSslInstanceNoSharedCipherAction                                               string                                                         `json:"no-shared-cipher-action,omitempty"`
	SlbTemplateClientSslInstanceNonSslBypassL4Session                                              int                                                            `json:"non-ssl-bypass-l4session,omitempty"`
	SlbTemplateClientSslInstanceNonSslBypassServiceGroup                                           string                                                         `json:"non-ssl-bypass-service-group,omitempty"`
	SlbTemplateClientSslInstanceNotafter                                                           int                                                            `json:"notafter,omitempty"`
	SlbTemplateClientSslInstanceNotafterday                                                        int                                                            `json:"notafterday,omitempty"`
	SlbTemplateClientSslInstanceNotaftermonth                                                      int                                                            `json:"notaftermonth,omitempty"`
	SlbTemplateClientSslInstanceNotafteryear                                                       int                                                            `json:"notafteryear,omitempty"`
	SlbTemplateClientSslInstanceNotbefore                                                          int                                                            `json:"notbefore,omitempty"`
	SlbTemplateClientSslInstanceNotbeforeday                                                       int                                                            `json:"notbeforeday,omitempty"`
	SlbTemplateClientSslInstanceNotbeforemonth                                                     int                                                            `json:"notbeforemonth,omitempty"`
	SlbTemplateClientSslInstanceNotbeforeyear                                                      int                                                            `json:"notbeforeyear,omitempty"`
	SlbTemplateClientSslInstanceOcspStapling                                                       int                                                            `json:"ocsp-stapling,omitempty"`
	SlbTemplateClientSslInstanceOcspstCaCert                                                       string                                                         `json:"ocspst-ca-cert,omitempty"`
	SlbTemplateClientSslInstanceOcspstOcsp                                                         int                                                            `json:"ocspst-ocsp,omitempty"`
	SlbTemplateClientSslInstanceOcspstSg                                                           string                                                         `json:"ocspst-sg,omitempty"`
	SlbTemplateClientSslInstanceOcspstSgDays                                                       int                                                            `json:"ocspst-sg-days,omitempty"`
	SlbTemplateClientSslInstanceOcspstSgHours                                                      int                                                            `json:"ocspst-sg-hours,omitempty"`
	SlbTemplateClientSslInstanceOcspstSgMinutes                                                    int                                                            `json:"ocspst-sg-minutes,omitempty"`
	SlbTemplateClientSslInstanceOcspstSgTimeout                                                    int                                                            `json:"ocspst-sg-timeout,omitempty"`
	SlbTemplateClientSslInstanceOcspstSrvr                                                         string                                                         `json:"ocspst-srvr,omitempty"`
	SlbTemplateClientSslInstanceOcspstSrvrDays                                                     int                                                            `json:"ocspst-srvr-days,omitempty"`
	SlbTemplateClientSslInstanceOcspstSrvrHours                                                    int                                                            `json:"ocspst-srvr-hours,omitempty"`
	SlbTemplateClientSslInstanceOcspstSrvrMinutes                                                  int                                                            `json:"ocspst-srvr-minutes,omitempty"`
	SlbTemplateClientSslInstanceOcspstSrvrTimeout                                                  int                                                            `json:"ocspst-srvr-timeout,omitempty"`
	SlbTemplateClientSslInstanceRenegotiationDisable                                               int                                                            `json:"renegotiation-disable,omitempty"`
	SlbTemplateClientSslInstanceReqCaListsClientCertificateRequestCA                               []SlbTemplateClientSslInstanceReqCaLists                       `json:"req-ca-lists,omitempty"`
	SlbTemplateClientSslInstanceRequireWebCategory                                                 int                                                            `json:"require-web-category,omitempty"`
	SlbTemplateClientSslInstanceServerIpv4ListServerIpv4ListName                                   []SlbTemplateClientSslInstanceServerIpv4List                   `json:"server-ipv4-list,omitempty"`
	SlbTemplateClientSslInstanceServerIpv6ListServerIpv6ListName                                   []SlbTemplateClientSslInstanceServerIpv6List                   `json:"server-ipv6-list,omitempty"`
	SlbTemplateClientSslInstanceServerNameAutoMap                                                  int                                                            `json:"server-name-auto-map,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerName                                           []SlbTemplateClientSslInstanceServerNameList                   `json:"server-name-list,omitempty"`
	SlbTemplateClientSslInstanceSessionCacheSize                                                   int                                                            `json:"session-cache-size,omitempty"`
	SlbTemplateClientSslInstanceSessionCacheTimeout                                                int                                                            `json:"session-cache-timeout,omitempty"`
	SlbTemplateClientSslInstanceSessionTicketDisable                                               int                                                            `json:"session-ticket-disable,omitempty"`
	SlbTemplateClientSslInstanceSessionTicketLifetime                                              int                                                            `json:"session-ticket-lifetime,omitempty"`
	SlbTemplateClientSslInstanceSharedPartitionCipherTemplate                                      int                                                            `json:"shared-partition-cipher-template,omitempty"`
	SlbTemplateClientSslInstanceSharedPartitionPool                                                int                                                            `json:"shared-partition-pool,omitempty"`
	SlbTemplateClientSslInstanceSniBypassEnableLog                                                 int                                                            `json:"sni-bypass-enable-log,omitempty"`
	SlbTemplateClientSslInstanceSniBypassExpiredCert                                               int                                                            `json:"sni-bypass-expired-cert,omitempty"`
	SlbTemplateClientSslInstanceSniBypassExplicitList                                              string                                                         `json:"sni-bypass-explicit-list,omitempty"`
	SlbTemplateClientSslInstanceSniBypassMissingCert                                               int                                                            `json:"sni-bypass-missing-cert,omitempty"`
	SlbTemplateClientSslInstanceSniEnableLog                                                       int                                                            `json:"sni-enable-log,omitempty"`
	SlbTemplateClientSslInstanceSslFalseStartDisable                                               int                                                            `json:"ssl-false-start-disable,omitempty"`
	SlbTemplateClientSslInstanceSsliLogging                                                        int                                                            `json:"ssli-logging,omitempty"`
	SlbTemplateClientSslInstanceSslilogging                                                        string                                                         `json:"sslilogging,omitempty"`
	SlbTemplateClientSslInstanceSslv2BypassServiceGroup                                            string                                                         `json:"sslv2-bypass-service-group,omitempty"`
	SlbTemplateClientSslInstanceStartsWithListStartsWith                                           []SlbTemplateClientSslInstanceStartsWithList                   `json:"starts-with-list,omitempty"`
	SlbTemplateClientSslInstanceTemplateCipher                                                     string                                                         `json:"template-cipher,omitempty"`
	SlbTemplateClientSslInstanceTemplateCipherShared                                               string                                                         `json:"template-cipher-shared,omitempty"`
	SlbTemplateClientSslInstanceTemplateHsm                                                        string                                                         `json:"template-hsm,omitempty"`
	SlbTemplateClientSslInstanceUUID                                                               string                                                         `json:"uuid,omitempty"`
	SlbTemplateClientSslInstanceUserNameList                                                       string                                                         `json:"user-name-list,omitempty"`
	SlbTemplateClientSslInstanceUserTag                                                            string                                                         `json:"user-tag,omitempty"`
	SlbTemplateClientSslInstanceVerifyCertFailAction                                               string                                                         `json:"verify-cert-fail-action,omitempty"`
	SlbTemplateClientSslInstanceVersion                                                            int                                                            `json:"version,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryUncategorized                                           SlbTemplateClientSslInstanceWebCategory                        `json:"web-category,omitempty"`
	SlbTemplateClientSslInstanceWebReputationBypassTrustworthy                                     SlbTemplateClientSslInstanceWebReputation                      `json:"web-reputation,omitempty"`
}

type SlbTemplateClientSslInstanceBypassCertIssuerMultiClassList struct {
	SlbTemplateClientSslInstanceBypassCertIssuerMultiClassListBypassCertIssuerMultiClassListName string `json:"bypass-cert-issuer-multi-class-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceBypassCertSanMultiClassList struct {
	SlbTemplateClientSslInstanceBypassCertSanMultiClassListBypassCertSanMultiClassListName string `json:"bypass-cert-san-multi-class-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceBypassCertSubjectMultiClassList struct {
	SlbTemplateClientSslInstanceBypassCertSubjectMultiClassListBypassCertSubjectMultiClassListName string `json:"bypass-cert-subject-multi-class-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceCaCerts struct {
	SlbTemplateClientSslInstanceCaCertsCaCert         string `json:"ca-cert,omitempty"`
	SlbTemplateClientSslInstanceCaCertsCaShared       int    `json:"ca-shared,omitempty"`
	SlbTemplateClientSslInstanceCaCertsClientOcsp     int    `json:"client-ocsp,omitempty"`
	SlbTemplateClientSslInstanceCaCertsClientOcspSg   string `json:"client-ocsp-sg,omitempty"`
	SlbTemplateClientSslInstanceCaCertsClientOcspSrvr string `json:"client-ocsp-srvr,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateIssuerContainsList struct {
	SlbTemplateClientSslInstanceCertificateIssuerContainsListCertificateIssuerContains string `json:"certificate-issuer-contains,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateIssuerEndsWithList struct {
	SlbTemplateClientSslInstanceCertificateIssuerEndsWithListCertificateIssuerEndsWith string `json:"certificate-issuer-ends-with,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateIssuerEqualsList struct {
	SlbTemplateClientSslInstanceCertificateIssuerEqualsListCertificateIssuerEquals string `json:"certificate-issuer-equals,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateIssuerStartsWithList struct {
	SlbTemplateClientSslInstanceCertificateIssuerStartsWithListCertificateIssuerStarts string `json:"certificate-issuer-starts,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateList struct {
	SlbTemplateClientSslInstanceCertificateListCert       string `json:"cert,omitempty"`
	SlbTemplateClientSslInstanceCertificateListChainCert  string `json:"chain-cert,omitempty"`
	SlbTemplateClientSslInstanceCertificateListKey        string `json:"key,omitempty"`
	SlbTemplateClientSslInstanceCertificateListPassphrase string `json:"passphrase,omitempty"`
	SlbTemplateClientSslInstanceCertificateListShared     int    `json:"shared,omitempty"`
	SlbTemplateClientSslInstanceCertificateListUUID       string `json:"uuid,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSanContainsList struct {
	SlbTemplateClientSslInstanceCertificateSanContainsListCertificateSanContains string `json:"certificate-san-contains,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSanEndsWithList struct {
	SlbTemplateClientSslInstanceCertificateSanEndsWithListCertificateSanEndsWith string `json:"certificate-san-ends-with,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSanEqualsList struct {
	SlbTemplateClientSslInstanceCertificateSanEqualsListCertificateSanEquals string `json:"certificate-san-equals,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSanStartsWithList struct {
	SlbTemplateClientSslInstanceCertificateSanStartsWithListCertificateSanStarts string `json:"certificate-san-starts,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSubjectContainsList struct {
	SlbTemplateClientSslInstanceCertificateSubjectContainsListCertificateSubjectContains string `json:"certificate-subject-contains,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSubjectEndsWithList struct {
	SlbTemplateClientSslInstanceCertificateSubjectEndsWithListCertificateSubjectEndsWith string `json:"certificate-subject-ends-with,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSubjectEqualsList struct {
	SlbTemplateClientSslInstanceCertificateSubjectEqualsListCertificateSubjectEquals string `json:"certificate-subject-equals,omitempty"`
}

type SlbTemplateClientSslInstanceCertificateSubjectStartsWithList struct {
	SlbTemplateClientSslInstanceCertificateSubjectStartsWithListCertificateSubjectStarts string `json:"certificate-subject-starts,omitempty"`
}

type SlbTemplateClientSslInstanceCipherWithoutPrioList struct {
	SlbTemplateClientSslInstanceCipherWithoutPrioListCipherWoPrio string `json:"cipher-wo-prio,omitempty"`
}

type SlbTemplateClientSslInstanceClientAuthContainsList struct {
	SlbTemplateClientSslInstanceClientAuthContainsListClientAuthContains string `json:"client-auth-contains,omitempty"`
}

type SlbTemplateClientSslInstanceClientAuthEndsWithList struct {
	SlbTemplateClientSslInstanceClientAuthEndsWithListClientAuthEndsWith string `json:"client-auth-ends-with,omitempty"`
}

type SlbTemplateClientSslInstanceClientAuthEqualsList struct {
	SlbTemplateClientSslInstanceClientAuthEqualsListClientAuthEquals string `json:"client-auth-equals,omitempty"`
}

type SlbTemplateClientSslInstanceClientAuthStartsWithList struct {
	SlbTemplateClientSslInstanceClientAuthStartsWithListClientAuthStartsWith string `json:"client-auth-starts-with,omitempty"`
}

type SlbTemplateClientSslInstanceClientIpv4List struct {
	SlbTemplateClientSslInstanceClientIpv4ListClientIpv4ListName string `json:"client-ipv4-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceClientIpv6List struct {
	SlbTemplateClientSslInstanceClientIpv6ListClientIpv6ListName string `json:"client-ipv6-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceContainsList struct {
	SlbTemplateClientSslInstanceContainsListContains string `json:"contains,omitempty"`
}

type SlbTemplateClientSslInstanceCrlCerts struct {
	SlbTemplateClientSslInstanceCrlCertsCrl       string `json:"crl,omitempty"`
	SlbTemplateClientSslInstanceCrlCertsCrlShared int    `json:"crl-shared,omitempty"`
}

type SlbTemplateClientSslInstanceEcList struct {
	SlbTemplateClientSslInstanceEcListEc string `json:"ec,omitempty"`
}

type SlbTemplateClientSslInstanceEndsWithList struct {
	SlbTemplateClientSslInstanceEndsWithListEndsWith string `json:"ends-with,omitempty"`
}

type SlbTemplateClientSslInstanceEqualsList struct {
	SlbTemplateClientSslInstanceEqualsListEquals string `json:"equals,omitempty"`
}

type SlbTemplateClientSslInstanceExceptionClientIpv4List struct {
	SlbTemplateClientSslInstanceExceptionClientIpv4ListExceptionClientIpv4ListName string `json:"exception-client-ipv4-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceExceptionClientIpv6List struct {
	SlbTemplateClientSslInstanceExceptionClientIpv6ListExceptionClientIpv6ListName string `json:"exception-client-ipv6-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceExceptionServerIpv4List struct {
	SlbTemplateClientSslInstanceExceptionServerIpv4ListExceptionServerIpv4ListName string `json:"exception-server-ipv4-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceExceptionServerIpv6List struct {
	SlbTemplateClientSslInstanceExceptionServerIpv6ListExceptionServerIpv6ListName string `json:"exception-server-ipv6-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceExceptionWebCategory struct {
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAbortion                    int `json:"exception-abortion,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAdultAndPornography         int `json:"exception-adult-and-pornography,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAlcoholAndTobacco           int `json:"exception-alcohol-and-tobacco,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionAuctions                    int `json:"exception-auctions,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionBotNets                     int `json:"exception-bot-nets,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionBusinessAndEconomy          int `json:"exception-business-and-economy,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionCdns                        int `json:"exception-cdns,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionCheating                    int `json:"exception-cheating,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionComputerAndInternetInfo     int `json:"exception-computer-and-internet-info,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionComputerAndInternetSecurity int `json:"exception-computer-and-internet-security,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionConfirmedSpamSources        int `json:"exception-confirmed-spam-sources,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionCultAndOccult               int `json:"exception-cult-and-occult,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDating                      int `json:"exception-dating,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDeadSites                   int `json:"exception-dead-sites,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDrugs                       int `json:"exception-drugs,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionDynamicComment              int `json:"exception-dynamic-comment,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionEducationalInstitutions     int `json:"exception-educational-institutions,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionEntertainmentAndArts        int `json:"exception-entertainment-and-arts,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionFashionAndBeauty            int `json:"exception-fashion-and-beauty,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionFinancialServices           int `json:"exception-financial-services,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionFoodAndDining               int `json:"exception-food-and-dining,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGambling                    int `json:"exception-gambling,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGames                       int `json:"exception-games,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGovernment                  int `json:"exception-government,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionGross                       int `json:"exception-gross,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHacking                     int `json:"exception-hacking,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHateAndRacism               int `json:"exception-hate-and-racism,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHealthAndMedicine           int `json:"exception-health-and-medicine,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHomeAndGarden               int `json:"exception-home-and-garden,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionHuntingAndFishing           int `json:"exception-hunting-and-fishing,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionIllegal                     int `json:"exception-illegal,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionIllegalPornography          int `json:"exception-illegal-pornography,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionImageAndVideoSearch         int `json:"exception-image-and-video-search,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionInternetCommunications      int `json:"exception-internet-communications,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionInternetPortals             int `json:"exception-internet-portals,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionJobSearch                   int `json:"exception-job-search,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionKeyloggersAndMonitoring     int `json:"exception-keyloggers-and-monitoring,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionKids                        int `json:"exception-kids,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionLegal                       int `json:"exception-legal,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionLocalInformation            int `json:"exception-local-information,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMalwareSites                int `json:"exception-malware-sites,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMarijuana                   int `json:"exception-marijuana,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMilitary                    int `json:"exception-military,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMotorVehicles               int `json:"exception-motor-vehicles,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionMusic                       int `json:"exception-music,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionNewsAndMedia                int `json:"exception-news-and-media,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionNudity                      int `json:"exception-nudity,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionNudityArtistic              int `json:"exception-nudity-artistic,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionOnlineGreetingCards         int `json:"exception-online-greeting-cards,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionOpenHTTPProxies             int `json:"exception-open-http-proxies,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionParkedDomains               int `json:"exception-parked-domains,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPayToSurf                   int `json:"exception-pay-to-surf,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPeerToPeer                  int `json:"exception-peer-to-peer,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPersonalSitesAndBlogs       int `json:"exception-personal-sites-and-blogs,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPersonalStorage             int `json:"exception-personal-storage,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPhilosophyAndPolitics       int `json:"exception-philosophy-and-politics,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPhishingAndOtherFraud       int `json:"exception-phishing-and-other-fraud,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionPrivateIPAddresses          int `json:"exception-private-ip-addresses,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionProxyAvoidAndAnonymizers    int `json:"exception-proxy-avoid-and-anonymizers,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionQuestionable                int `json:"exception-questionable,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionRealEstate                  int `json:"exception-real-estate,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionRecreationAndHobbies        int `json:"exception-recreation-and-hobbies,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionReferenceAndResearch        int `json:"exception-reference-and-research,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionReligion                    int `json:"exception-religion,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSearchEngines               int `json:"exception-search-engines,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSexEducation                int `json:"exception-sex-education,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSharewareAndFreeware        int `json:"exception-shareware-and-freeware,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionShopping                    int `json:"exception-shopping,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSocialNetwork               int `json:"exception-social-network,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSociety                     int `json:"exception-society,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSpamUrls                    int `json:"exception-spam-urls,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSports                      int `json:"exception-sports,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSpywareAndAdware            int `json:"exception-spyware-and-adware,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionStockAdviceAndTools         int `json:"exception-stock-advice-and-tools,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionStreamingMedia              int `json:"exception-streaming-media,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionSwimsuitsAndIntimateApparel int `json:"exception-swimsuits-and-intimate-apparel,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionTrainingAndTools            int `json:"exception-training-and-tools,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionTranslation                 int `json:"exception-translation,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionTravel                      int `json:"exception-travel,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionUncategorized               int `json:"exception-uncategorized,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionUnconfirmedSpamSources      int `json:"exception-unconfirmed-spam-sources,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionViolence                    int `json:"exception-violence,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWeapons                     int `json:"exception-weapons,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWebAdvertisements           int `json:"exception-web-advertisements,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWebBasedEmail               int `json:"exception-web-based-email,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebCategoryExceptionWebHostingSites             int `json:"exception-web-hosting-sites,omitempty"`
}

type SlbTemplateClientSslInstanceExceptionWebReputation struct {
	SlbTemplateClientSslInstanceExceptionWebReputationExceptionLowRisk      int `json:"exception-low-risk,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebReputationExceptionMalicious    int `json:"exception-malicious,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebReputationExceptionModerateRisk int `json:"exception-moderate-risk,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebReputationExceptionSuspicious   int `json:"exception-suspicious,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebReputationExceptionThreshold    int `json:"exception-threshold,omitempty"`
	SlbTemplateClientSslInstanceExceptionWebReputationExceptionTrustworthy  int `json:"exception-trustworthy,omitempty"`
}

type SlbTemplateClientSslInstanceForwardProxyTrustedCaLists struct {
	SlbTemplateClientSslInstanceForwardProxyTrustedCaListsForwardProxyTrustedCa string `json:"forward-proxy-trusted-ca,omitempty"`
	SlbTemplateClientSslInstanceForwardProxyTrustedCaListsFpTrustedCaShared     int    `json:"fp-trusted-ca-shared,omitempty"`
}

type SlbTemplateClientSslInstanceLocalCertPinList struct {
	SlbTemplateClientSslInstanceLocalCertPinListLocalCertPinListBypassFailCount int `json:"local-cert-pin-list-bypass-fail-count,omitempty"`
}

type SlbTemplateClientSslInstanceMultiClassList struct {
	SlbTemplateClientSslInstanceMultiClassListMultiClistName string `json:"multi-clist-name,omitempty"`
}

type SlbTemplateClientSslInstanceReqCaLists struct {
	SlbTemplateClientSslInstanceReqCaListsClientCertReqCaShared      int    `json:"client-cert-req-ca-shared,omitempty"`
	SlbTemplateClientSslInstanceReqCaListsClientCertificateRequestCA string `json:"client-certificate-Request-CA,omitempty"`
}

type SlbTemplateClientSslInstanceServerIpv4List struct {
	SlbTemplateClientSslInstanceServerIpv4ListServerIpv4ListName string `json:"server-ipv4-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceServerIpv6List struct {
	SlbTemplateClientSslInstanceServerIpv6ListServerIpv6ListName string `json:"server-ipv6-list-name,omitempty"`
}

type SlbTemplateClientSslInstanceServerNameList struct {
	SlbTemplateClientSslInstanceServerNameListServerCert               string `json:"server-cert,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerCertRegex          string `json:"server-cert-regex,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerChain              string `json:"server-chain,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerChainRegex         string `json:"server-chain-regex,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerKey                string `json:"server-key,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerKeyRegex           string `json:"server-key-regex,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerName               string `json:"server-name,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerNameAlternate      int    `json:"server-name-alternate,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerNameRegex          string `json:"server-name-regex,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerNameRegexAlternate int    `json:"server-name-regex-alternate,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerPassphrase         string `json:"server-passphrase,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerPassphraseRegex    string `json:"server-passphrase-regex,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerShared             int    `json:"server-shared,omitempty"`
	SlbTemplateClientSslInstanceServerNameListServerSharedRegex        int    `json:"server-shared-regex,omitempty"`
}

type SlbTemplateClientSslInstanceStartsWithList struct {
	SlbTemplateClientSslInstanceStartsWithListStartsWith string `json:"starts-with,omitempty"`
}

type SlbTemplateClientSslInstanceWebCategory struct {
	SlbTemplateClientSslInstanceWebCategoryAbortion                    int `json:"abortion,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryAdultAndPornography         int `json:"adult-and-pornography,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryAlcoholAndTobacco           int `json:"alcohol-and-tobacco,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryAuctions                    int `json:"auctions,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryBotNets                     int `json:"bot-nets,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryBusinessAndEconomy          int `json:"business-and-economy,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryCdns                        int `json:"cdns,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryCheating                    int `json:"cheating,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryComputerAndInternetInfo     int `json:"computer-and-internet-info,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryComputerAndInternetSecurity int `json:"computer-and-internet-security,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryConfirmedSpamSources        int `json:"confirmed-spam-sources,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryCultAndOccult               int `json:"cult-and-occult,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryDating                      int `json:"dating,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryDeadSites                   int `json:"dead-sites,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryDrugs                       int `json:"drugs,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryDynamicComment              int `json:"dynamic-comment,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryEducationalInstitutions     int `json:"educational-institutions,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryEntertainmentAndArts        int `json:"entertainment-and-arts,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryFashionAndBeauty            int `json:"fashion-and-beauty,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryFinancialServices           int `json:"financial-services,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryFoodAndDining               int `json:"food-and-dining,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryGambling                    int `json:"gambling,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryGames                       int `json:"games,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryGovernment                  int `json:"government,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryGross                       int `json:"gross,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryHacking                     int `json:"hacking,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryHateAndRacism               int `json:"hate-and-racism,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryHealthAndMedicine           int `json:"health-and-medicine,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryHomeAndGarden               int `json:"home-and-garden,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryHuntingAndFishing           int `json:"hunting-and-fishing,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryIllegal                     int `json:"illegal,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryIllegalPornography          int `json:"illegal-pornography,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryImageAndVideoSearch         int `json:"image-and-video-search,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryInternetCommunications      int `json:"internet-communications,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryInternetPortals             int `json:"internet-portals,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryJobSearch                   int `json:"job-search,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryKeyloggersAndMonitoring     int `json:"keyloggers-and-monitoring,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryKids                        int `json:"kids,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryLegal                       int `json:"legal,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryLocalInformation            int `json:"local-information,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryMalwareSites                int `json:"malware-sites,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryMarijuana                   int `json:"marijuana,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryMilitary                    int `json:"military,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryMotorVehicles               int `json:"motor-vehicles,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryMusic                       int `json:"music,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryNewsAndMedia                int `json:"news-and-media,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryNudity                      int `json:"nudity,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryNudityArtistic              int `json:"nudity-artistic,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryOnlineGreetingCards         int `json:"online-greeting-cards,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryOpenHTTPProxies             int `json:"open-http-proxies,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryParkedDomains               int `json:"parked-domains,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryPayToSurf                   int `json:"pay-to-surf,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryPeerToPeer                  int `json:"peer-to-peer,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryPersonalSitesAndBlogs       int `json:"personal-sites-and-blogs,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryPersonalStorage             int `json:"personal-storage,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryPhilosophyAndPolitics       int `json:"philosophy-and-politics,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryPhishingAndOtherFraud       int `json:"phishing-and-other-fraud,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryPrivateIPAddresses          int `json:"private-ip-addresses,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryProxyAvoidAndAnonymizers    int `json:"proxy-avoid-and-anonymizers,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryQuestionable                int `json:"questionable,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryRealEstate                  int `json:"real-estate,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryRecreationAndHobbies        int `json:"recreation-and-hobbies,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryReferenceAndResearch        int `json:"reference-and-research,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryReligion                    int `json:"religion,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySearchEngines               int `json:"search-engines,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySexEducation                int `json:"sex-education,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySharewareAndFreeware        int `json:"shareware-and-freeware,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryShopping                    int `json:"shopping,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySocialNetwork               int `json:"social-network,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySociety                     int `json:"society,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySpamUrls                    int `json:"spam-urls,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySports                      int `json:"sports,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySpywareAndAdware            int `json:"spyware-and-adware,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryStockAdviceAndTools         int `json:"stock-advice-and-tools,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryStreamingMedia              int `json:"streaming-media,omitempty"`
	SlbTemplateClientSslInstanceWebCategorySwimsuitsAndIntimateApparel int `json:"swimsuits-and-intimate-apparel,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryTrainingAndTools            int `json:"training-and-tools,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryTranslation                 int `json:"translation,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryTravel                      int `json:"travel,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryUncategorized               int `json:"uncategorized,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryUnconfirmedSpamSources      int `json:"unconfirmed-spam-sources,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryViolence                    int `json:"violence,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryWeapons                     int `json:"weapons,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryWebAdvertisements           int `json:"web-advertisements,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryWebBasedEmail               int `json:"web-based-email,omitempty"`
	SlbTemplateClientSslInstanceWebCategoryWebHostingSites             int `json:"web-hosting-sites,omitempty"`
}

type SlbTemplateClientSslInstanceWebReputation struct {
	SlbTemplateClientSslInstanceWebReputationBypassLowRisk      int `json:"bypass-low-risk,omitempty"`
	SlbTemplateClientSslInstanceWebReputationBypassMalicious    int `json:"bypass-malicious,omitempty"`
	SlbTemplateClientSslInstanceWebReputationBypassModerateRisk int `json:"bypass-moderate-risk,omitempty"`
	SlbTemplateClientSslInstanceWebReputationBypassSuspicious   int `json:"bypass-suspicious,omitempty"`
	SlbTemplateClientSslInstanceWebReputationBypassThreshold    int `json:"bypass-threshold,omitempty"`
	SlbTemplateClientSslInstanceWebReputationBypassTrustworthy  int `json:"bypass-trustworthy,omitempty"`
}

func PostSlbTemplateClientSsl(id string, inst SlbTemplateClientSsl, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateClientSsl")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/client-ssl", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateClientSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateClientSsl", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateClientSsl(id string, name1 string, host string) (*SlbTemplateClientSsl, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateClientSsl")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/client-ssl/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateClientSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateClientSsl", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateClientSsl(id string, name1 string, inst SlbTemplateClientSsl, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateClientSsl")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/client-ssl/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateClientSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateClientSsl", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateClientSsl(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateClientSsl")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/client-ssl/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateClientSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateClientSsl", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
