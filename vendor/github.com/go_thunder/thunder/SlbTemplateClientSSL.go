package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type ClientSSL struct {
	UUID ClientSSLInstance `json:"client-ssl,omitempty"`
}

type BypassCertSubjectMultiClassList struct {
	BypassCertSubjectMultiClassListName string `json:"bypass-cert-subject-multi-class-list-name,omitempty"`
}
type CertificateSanContainsList struct {
	CertificateSanContains string `json:"certificate-san-contains,omitempty"`
}
type EqualsList struct {
	Equals string `json:"equals,omitempty"`
}
type ForwardProxyTrustedCaLists struct {
	ForwardProxyTrustedCa string `json:"forward-proxy-trusted-ca,omitempty"`
}
type ContainsList struct {
	Contains string `json:"contains,omitempty"`
}
type EndsWithList struct {
	EndsWith string `json:"ends-with,omitempty"`
}
type CaCerts2 struct {
	ClientOcspSrvr string `json:"client-ocsp-srvr,omitempty"`
	CaCert         string `json:"ca-cert,omitempty"`
	ClientOcspSg   string `json:"client-ocsp-sg,omitempty"`
	ClientOcsp     int    `json:"client-ocsp,omitempty"`
	CaShared       int    `json:"ca-shared,omitempty"`
}
type ClientAuthContainsList struct {
	ClientAuthContains string `json:"client-auth-contains,omitempty"`
}
type CertificateSubjectContainsList struct {
	CertificateSubjectContains string `json:"certificate-subject-contains,omitempty"`
}
type ReqCaLists struct {
	ClientCertificateRequestCA string `json:"client-certificate-Request-CA,omitempty"`
}

// type SamplingEnable5 struct {
// 	Counters1 string `json:"counters1,omitempty"`
// }
type BypassCertIssuerMultiClassList struct {
	BypassCertIssuerMultiClassListName string `json:"bypass-cert-issuer-multi-class-list-name,omitempty"`
}
type ClientAuthEqualsList struct {
	ClientAuthEquals string `json:"client-auth-equals,omitempty"`
}
type CertificateIssuerEqualsList struct {
	CertificateIssuerEquals string `json:"certificate-issuer-equals,omitempty"`
}
type CertificateSubjectStartsWithList struct {
	CertificateSubjectStarts string `json:"certificate-subject-starts,omitempty"`
}
type CertificateSanEndsWithList struct {
	CertificateSanEndsWith string `json:"certificate-san-ends-with,omitempty"`
}
type CrlCerts2 struct {
	CrlShared int    `json:"crl-shared,omitempty"`
	Crl       string `json:"crl,omitempty"`
}
type MultiClassList struct {
	MultiClistName string `json:"multi-clist-name,omitempty"`
}
type CertificateIssuerEndsWithList struct {
	CertificateIssuerEndsWith string `json:"certificate-issuer-ends-with,omitempty"`
}
type WebCategory struct {
	PhilosophyAndPolitics       int `json:"philosophy-and-politics,omitempty"`
	StockAdviceAndTools         int `json:"stock-advice-and-tools,omitempty"`
	NewsAndMedia                int `json:"news-and-media,omitempty"`
	BusinessAndEconomy          int `json:"business-and-economy,omitempty"`
	PeerToPeer                  int `json:"peer-to-peer,omitempty"`
	PhishingAndOtherFraud       int `json:"phishing-and-other-fraud,omitempty"`
	Nudity                      int `json:"nudity,omitempty"`
	Weapons                     int `json:"weapons,omitempty"`
	HealthAndMedicine           int `json:"health-and-medicine,omitempty"`
	Marijuana                   int `json:"marijuana,omitempty"`
	HomeAndGarden               int `json:"home-and-garden,omitempty"`
	CultAndOccult               int `json:"cult-and-occult,omitempty"`
	Society                     int `json:"society,omitempty"`
	PersonalStorage             int `json:"personal-storage,omitempty"`
	ComputerAndInternetSecurity int `json:"computer-and-internet-security,omitempty"`
	FoodAndDining               int `json:"food-and-dining,omitempty"`
	MotorVehicles               int `json:"motor-vehicles,omitempty"`
	SwimsuitsAndIntimateApparel int `json:"swimsuits-and-intimate-apparel,omitempty"`
	DeadSites                   int `json:"dead-sites,omitempty"`
	Translation                 int `json:"translation,omitempty"`
	ProxyAvoidAndAnonymizers    int `json:"proxy-avoid-and-anonymizers,omitempty"`
	FinancialServices           int `json:"financial-services,omitempty"`
	Gross                       int `json:"gross,omitempty"`
	Cheating                    int `json:"cheating,omitempty"`
	EntertainmentAndArts        int `json:"entertainment-and-arts,omitempty"`
	SexEducation                int `json:"sex-education,omitempty"`
	Illegal                     int `json:"illegal,omitempty"`
	Travel                      int `json:"travel,omitempty"`
	Cdns                        int `json:"cdns,omitempty"`
	LocalInformation            int `json:"local-information,omitempty"`
	Legal                       int `json:"legal,omitempty"`
	Sports                      int `json:"sports,omitempty"`
	BotNets                     int `json:"bot-nets,omitempty"`
	Religion                    int `json:"religion,omitempty"`
	PrivateIPAddresses          int `json:"private-ip-addresses,omitempty"`
	Music                       int `json:"music,omitempty"`
	HateAndRacism               int `json:"hate-and-racism,omitempty"`
	OpenHTTPProxies             int `json:"open-http-proxies,omitempty"`
	InternetCommunications      int `json:"internet-communications,omitempty"`
	SharewareAndFreeware        int `json:"shareware-and-freeware,omitempty"`
	Dating                      int `json:"dating,omitempty"`
	SpywareAndAdware            int `json:"spyware-and-adware,omitempty"`
	Uncategorized               int `json:"uncategorized,omitempty"`
	Questionable                int `json:"questionable,omitempty"`
	ReferenceAndResearch        int `json:"reference-and-research,omitempty"`
	WebAdvertisements           int `json:"web-advertisements,omitempty"`
	StreamingMedia              int `json:"streaming-media,omitempty"`
	SocialNetwork               int `json:"social-network,omitempty"`
	Government                  int `json:"government,omitempty"`
	Drugs                       int `json:"drugs,omitempty"`
	WebHostingSites             int `json:"web-hosting-sites,omitempty"`
	MalwareSites                int `json:"malware-sites,omitempty"`
	PayToSurf                   int `json:"pay-to-surf,omitempty"`
	SpamUrls                    int `json:"spam-urls,omitempty"`
	Kids                        int `json:"kids,omitempty"`
	Gambling                    int `json:"gambling,omitempty"`
	OnlineGreetingCards         int `json:"online-greeting-cards,omitempty"`
	ConfirmedSpamSources        int `json:"confirmed-spam-sources,omitempty"`
	ImageAndVideoSearch         int `json:"image-and-video-search,omitempty"`
	EducationalInstitutions     int `json:"educational-institutions,omitempty"`
	KeyloggersAndMonitoring     int `json:"keyloggers-and-monitoring,omitempty"`
	HuntingAndFishing           int `json:"hunting-and-fishing,omitempty"`
	SearchEngines               int `json:"search-engines,omitempty"`
	FashionAndBeauty            int `json:"fashion-and-beauty,omitempty"`
	DynamicComment              int `json:"dynamic-comment,omitempty"`
	ComputerAndInternetInfo     int `json:"computer-and-internet-info,omitempty"`
	RealEstate                  int `json:"real-estate,omitempty"`
	InternetPortals             int `json:"internet-portals,omitempty"`
	Shopping                    int `json:"shopping,omitempty"`
	Violence                    int `json:"violence,omitempty"`
	Abortion                    int `json:"abortion,omitempty"`
	TrainingAndTools            int `json:"training-and-tools,omitempty"`
	WebBasedEmail               int `json:"web-based-email,omitempty"`
	PersonalSitesAndBlogs       int `json:"personal-sites-and-blogs,omitempty"`
	UnconfirmedSpamSources      int `json:"unconfirmed-spam-sources,omitempty"`
	Games                       int `json:"games,omitempty"`
	ParkedDomains               int `json:"parked-domains,omitempty"`
	Auctions                    int `json:"auctions,omitempty"`
	JobSearch                   int `json:"job-search,omitempty"`
	RecreationAndHobbies        int `json:"recreation-and-hobbies,omitempty"`
	Hacking                     int `json:"hacking,omitempty"`
	AlcoholAndTobacco           int `json:"alcohol-and-tobacco,omitempty"`
	AdultAndPornography         int `json:"adult-and-pornography,omitempty"`
	Military                    int `json:"military,omitempty"`
}
type CertificateSanEqualsList struct {
	CertificateSanEquals string `json:"certificate-san-equals,omitempty"`
}
type CertificateIssuerContainsList struct {
	CertificateIssuerContains string `json:"certificate-issuer-contains,omitempty"`
}
type BypassCertSanMultiClassList struct {
	BypassCertSanMultiClassListName string `json:"bypass-cert-san-multi-class-list-name,omitempty"`
}
type ClientAuthStartsWithList struct {
	ClientAuthStartsWith string `json:"client-auth-starts-with,omitempty"`
}
type CertificateSubjectEndsWithList struct {
	CertificateSubjectEndsWith string `json:"certificate-subject-ends-with,omitempty"`
}
type EcList2 struct {
	Ec string `json:"ec,omitempty"`
}
type ServerNameList struct {
	ServerShared             int    `json:"server-shared,omitempty"`
	ServerPassphraseRegex    string `json:"server-passphrase-regex,omitempty"`
	ServerChain              string `json:"server-chain,omitempty"`
	ServerCertRegex          string `json:"server-cert-regex,omitempty"`
	ServerName               string `json:"server-name,omitempty"`
	ServerKeyRegex           string `json:"server-key-regex,omitempty"`
	ServerNameRegexAlternate int    `json:"server-name-regex-alternate,omitempty"`
	ServerEncryptedRegex     string `json:"server-encrypted-regex,omitempty"`
	ServerSharedRegex        int    `json:"server-shared-regex,omitempty"`
	ServerNameRegex          string `json:"server-name-regex,omitempty"`
	ServerPassphrase         string `json:"server-passphrase,omitempty"`
	ServerKey                string `json:"server-key,omitempty"`
	ServerChainRegex         string `json:"server-chain-regex,omitempty"`
	ServerNameAlternate      int    `json:"server-name-alternate,omitempty"`
	ServerEncrypted          string `json:"server-encrypted,omitempty"`
	ServerCert               string `json:"server-cert,omitempty"`
}
type CertificateIssuerStartsWithList struct {
	CertificateIssuerStarts string `json:"certificate-issuer-starts,omitempty"`
}
type CertificateSanStartsWithList struct {
	CertificateSanStarts string `json:"certificate-san-starts,omitempty"`
}
type ClientAuthEndsWithList struct {
	ClientAuthEndsWith string `json:"client-auth-ends-with,omitempty"`
}
type CertificateSubjectEqualsList struct {
	CertificateSubjectEquals string `json:"certificate-subject-equals,omitempty"`
}
type CipherWithoutPrioList2 struct {
	CipherWoPrio string `json:"cipher-wo-prio,omitempty"`
}
type StartsWithList struct {
	StartsWith string `json:"starts-with,omitempty"`
}

type SlbTemplateCertificateList struct {
	Cert         string `json:"cert,omitempty"`
	ChainCert    string `json:"chain-cert,omitempty"`
	Key          string `json:"key,omitempty"`
	KeyEncrypted string `json:"key-encrypted,omitempty"`
	Passphrase   string `json:"passphrase,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

type ClientSSLInstance struct {
	BypassCertSubjectMultiClassListName []BypassCertSubjectMultiClassList `json:"bypass-cert-subject-multi-class-list,omitempty"`
	VerifyCertFailAction                string                            `json:"verify-cert-fail-action,omitempty"`
	InspectCertificateIssuerClName      string                            `json:"inspect-certificate-issuer-cl-name,omitempty"`
	CertificateSanContains              []CertificateSanContainsList      `json:"certificate-san-contains-list,omitempty"`
	ForwardProxyBlockMessage            string                            `json:"forward-proxy-block-message,omitempty"`
	DirectClientServerAuth              int                               `json:"direct-client-server-auth,omitempty"`
	OcspstSgHours                       int                               `json:"ocspst-sg-hours,omitempty"`
	NoSharedCipherAction                string                            `json:"no-shared-cipher-action,omitempty"`
	FpCertFetchAutonat                  string                            `json:"fp-cert-fetch-autonat,omitempty"`
	Equals                              []EqualsList                      `json:"equals-list,omitempty"`
	Passphrase                          []SlbTemplateCertificateList      `json:"certificate-list,omitempty"`
	ExceptionCertificateSubjectClName   string                            `json:"exception-certificate-subject-cl-name,omitempty"`
	UUID                                string                            `json:"uuid,omitempty"`
	ForwardProxyTrustedCa               []ForwardProxyTrustedCaLists      `json:"forward-proxy-trusted-ca-lists,omitempty"`
	TemplateCipherShared                string                            `json:"template-cipher-shared,omitempty"`
	ForwardProxyCaCert                  string                            `json:"forward-proxy-ca-cert,omitempty"`
	SslFalseStartDisable                int                               `json:"ssl-false-start-disable,omitempty"`
	Dgversion                           int                               `json:"dgversion,omitempty"`
	ClientAuthClassList                 string                            `json:"client-auth-class-list,omitempty"`
	KeyEncrypted                        string                            `json:"key-encrypted,omitempty"`
	Notafteryear                        int                               `json:"notafteryear,omitempty"`
	ForwardProxyAltSign                 int                               `json:"forward-proxy-alt-sign,omitempty"`
	TemplateHsm                         string                            `json:"template-hsm,omitempty"`
	ForwardPassphrase                   string                            `json:"forward-passphrase,omitempty"`
	ExceptionCertificateIssuerClName    string                            `json:"exception-certificate-issuer-cl-name,omitempty"`
	Contains                            []ContainsList                    `json:"contains-list,omitempty"`
	ForwardProxyCaKey                   string                            `json:"forward-proxy-ca-key,omitempty"`
	Notbefore                           int                               `json:"notbefore,omitempty"`
	EndsWith                            []EndsWithList                    `json:"ends-with-list,omitempty"`
	BypassCertSubjectClassListName      string                            `json:"bypass-cert-subject-class-list-name,omitempty"`
	Notafter                            int                               `json:"notafter,omitempty"`
	ClassListName                       string                            `json:"class-list-name,omitempty"`
	OcspstOcsp                          int                               `json:"ocspst-ocsp,omitempty"`
	Notbeforeday                        int                               `json:"notbeforeday,omitempty"`
	KeyAltPassphrase                    string                            `json:"key-alt-passphrase,omitempty"`
	ForwardProxySslVersion              int                               `json:"forward-proxy-ssl-version,omitempty"`
	ClientOcspSrvr                      []CaCerts2                        `json:"ca-certs,omitempty"`
	ForwardProxyCrlDisable              int                               `json:"forward-proxy-crl-disable,omitempty"`
	ClientAuthContains                  []ClientAuthContainsList          `json:"client-auth-contains-list,omitempty"`
	CertificateSubjectContains          []CertificateSubjectContainsList  `json:"certificate-subject-contains-list,omitempty"`
	Name                                string                            `json:"name,omitempty"`
	Key                                 string                            `json:"key,omitempty"`
	Cert                                string                            `json:"cert,omitempty"`
	ForwardProxyCertRevokeAction        int                               `json:"forward-proxy-cert-revoke-action,omitempty"`
	FpCertExtAiaOcsp                    string                            `json:"fp-cert-ext-aia-ocsp,omitempty"`
	ClientCertificateRequestCA          []ReqCaLists                      `json:"req-ca-lists,omitempty"`
	UserTag                             string                            `json:"user-tag,omitempty"`
	CertUnknownAction                   string                            `json:"cert-unknown-action,omitempty"`
	//Counters1                           []SamplingEnable5                  `json:"sampling-enable,omitempty"`
	RenegotiationDisable               int                                `json:"renegotiation-disable,omitempty"`
	ExceptionAdGroupList               string                             `json:"exception-ad-group-list,omitempty"`
	KeyAlternate                       string                             `json:"key-alternate,omitempty"`
	FpAltKey                           string                             `json:"fp-alt-key,omitempty"`
	ServerNameAutoMap                  int                                `json:"server-name-auto-map,omitempty"`
	DisableSslv3                       int                                `json:"disable-sslv3,omitempty"`
	BypassCertIssuerMultiClassListName []BypassCertIssuerMultiClassList   `json:"bypass-cert-issuer-multi-class-list,omitempty"`
	ClientAuthEquals                   []ClientAuthEqualsList             `json:"client-auth-equals-list,omitempty"`
	ForwardProxyNoSniAction            string                             `json:"forward-proxy-no-sni-action,omitempty"`
	CertificateIssuerEquals            []CertificateIssuerEqualsList      `json:"certificate-issuer-equals-list,omitempty"`
	FpAltPassphrase                    string                             `json:"fp-alt-passphrase,omitempty"`
	CertificateSubjectStarts           []CertificateSubjectStartsWithList `json:"certificate-subject-starts-with-list,omitempty"`
	CertificateSanEndsWith             []CertificateSanEndsWithList       `json:"certificate-san-ends-with-list,omitempty"`
	ForwardProxyCertCacheTimeout       int                                `json:"forward-proxy-cert-cache-timeout,omitempty"`
	FpCertFetchNatpoolNameShared       string                             `json:"fp-cert-fetch-natpool-name-shared,omitempty"`
	CrlShared                          []CrlCerts2                        `json:"crl-certs,omitempty"`
	Notafterday                        int                                `json:"notafterday,omitempty"`
	OcspstSrvrHours                    int                                `json:"ocspst-srvr-hours,omitempty"`
	LocalLogging                       int                                `json:"local-logging,omitempty"`
	FpCertFetchAutonatPrecedence       int                                `json:"fp-cert-fetch-autonat-precedence,omitempty"`
	//CertStr                             string                             `json:"cert-str,omitempty"`
	CertSharedStr                    string                            `json:"cert-shared-str,omitempty"`
	CertRevokeAction                 string                            `json:"cert-revoke-action,omitempty"`
	Version                          int                               `json:"version,omitempty"`
	MultiClistName                   []MultiClassList                  `json:"multi-class-list,omitempty"`
	UserNameList                     string                            `json:"user-name-list,omitempty"`
	SessionTicketLifetime            int                               `json:"session-ticket-lifetime,omitempty"`
	CertificateIssuerEndsWith        []CertificateIssuerEndsWithList   `json:"certificate-issuer-ends-with-list,omitempty"`
	SsliLogging                      int                               `json:"ssli-logging,omitempty"`
	SessionCacheSize                 int                               `json:"session-cache-size,omitempty"`
	HandshakeLoggingEnable           int                               `json:"handshake-logging-enable,omitempty"`
	NonSslBypassServiceGroup         string                            `json:"non-ssl-bypass-service-group,omitempty"`
	ForwardProxyFailsafeDisable      int                               `json:"forward-proxy-failsafe-disable,omitempty"`
	SessionCacheTimeout              int                               `json:"session-cache-timeout,omitempty"`
	Sslv2BypassServiceGroup          string                            `json:"sslv2-bypass-service-group,omitempty"`
	ForwardProxyDecryptedDscp        int                               `json:"forward-proxy-decrypted-dscp,omitempty"`
	AuthSg                           string                            `json:"auth-sg,omitempty"`
	OcspstCaCert                     string                            `json:"ocspst-ca-cert,omitempty"`
	ForwardProxySelfsignRedir        int                               `json:"forward-proxy-selfsign-redir,omitempty"`
	AuthSgDn                         int                               `json:"auth-sg-dn,omitempty"`
	HsmType                          string                            `json:"hsm-type,omitempty"`
	ForwardProxyLogDisable           int                               `json:"forward-proxy-log-disable,omitempty"`
	FpAltEncrypted                   string                            `json:"fp-alt-encrypted,omitempty"`
	InspectCertificateSanClName      string                            `json:"inspect-certificate-san-cl-name,omitempty"`
	PhilosophyAndPolitics            WebCategory                       `json:"web-category,omitempty"`
	CertificateSanEquals             []CertificateSanEqualsList        `json:"certificate-san-equals-list,omitempty"`
	TemplateCipher                   string                            `json:"template-cipher,omitempty"`
	Notbeforemonth                   int                               `json:"notbeforemonth,omitempty"`
	BypassCertSanClassListName       string                            `json:"bypass-cert-san-class-list-name,omitempty"`
	ChainCert                        string                            `json:"chain-cert,omitempty"`
	ForwardProxyCertUnknownAction    int                               `json:"forward-proxy-cert-unknown-action,omitempty"`
	ExceptionCertificateSanClName    string                            `json:"exception-certificate-san-cl-name,omitempty"`
	OcspstSg                         string                            `json:"ocspst-sg,omitempty"`
	KeyAltEncrypted                  string                            `json:"key-alt-encrypted,omitempty"`
	FpCertExtAiaCaIssuers            string                            `json:"fp-cert-ext-aia-ca-issuers,omitempty"`
	AuthenName                       string                            `json:"authen-name,omitempty"`
	ExpireHours                      int                               `json:"expire-hours,omitempty"`
	ClientAuthCaseInsensitive        int                               `json:"client-auth-case-insensitive,omitempty"`
	OcspStapling                     int                               `json:"ocsp-stapling,omitempty"`
	Notbeforeyear                    int                               `json:"notbeforeyear,omitempty"`
	ForwardEncrypted                 string                            `json:"forward-encrypted,omitempty"`
	SniEnableLog                     int                               `json:"sni-enable-log,omitempty"`
	KeySharedStr                     string                            `json:"key-shared-str,omitempty"`
	Notaftermonth                    int                               `json:"notaftermonth,omitempty"`
	CachePersistenceListName         string                            `json:"cache-persistence-list-name,omitempty"`
	OcspstSgTimeout                  int                               `json:"ocspst-sg-timeout,omitempty"`
	KeyPassphrase                    string                            `json:"key-passphrase,omitempty"`
	OcspstSrvr                       string                            `json:"ocspst-srvr,omitempty"`
	OcspstSrvrMinutes                int                               `json:"ocspst-srvr-minutes,omitempty"`
	CertificateIssuerContains        []CertificateIssuerContainsList   `json:"certificate-issuer-contains-list,omitempty"`
	RequireWebCategory               int                               `json:"require-web-category,omitempty"`
	BypassCertSanMultiClassListName  []BypassCertSanMultiClassList     `json:"bypass-cert-san-multi-class-list,omitempty"`
	ClientAuthStartsWith             []ClientAuthStartsWithList        `json:"client-auth-starts-with-list,omitempty"`
	CertificateSubjectEndsWith       []CertificateSubjectEndsWithList  `json:"certificate-subject-ends-with-list,omitempty"`
	Authorization                    int                               `json:"authorization,omitempty"`
	ForwardProxyVerifyCertFailAction int                               `json:"forward-proxy-verify-cert-fail-action,omitempty"`
	OcspstSrvrDays                   int                               `json:"ocspst-srvr-days,omitempty"`
	Ec                               []EcList2                         `json:"ec-list,omitempty"`
	ForwardProxyDecryptedDscpBypass  int                               `json:"forward-proxy-decrypted-dscp-bypass,omitempty"`
	AlertType                        string                            `json:"alert-type,omitempty"`
	ForwardProxyCertNotReadyAction   string                            `json:"forward-proxy-cert-not-ready-action,omitempty"`
	ServerShared                     []ServerNameList                  `json:"server-name-list,omitempty"`
	BypassCertIssuerClassListName    string                            `json:"bypass-cert-issuer-class-list-name,omitempty"`
	FpCertExtCrldp                   string                            `json:"fp-cert-ext-crldp,omitempty"`
	SharedPartitionCipherTemplate    int                               `json:"shared-partition-cipher-template,omitempty"`
	FpCertFetchNatpoolPrecedence     int                               `json:"fp-cert-fetch-natpool-precedence,omitempty"`
	CertAlternate                    string                            `json:"cert-alternate,omitempty"`
	ForwardProxyCertCacheLimit       int                               `json:"forward-proxy-cert-cache-limit,omitempty"`
	NonSslBypassL4Session            int                               `json:"non-ssl-bypass-l4session,omitempty"`
	CertificateIssuerStarts          []CertificateIssuerStartsWithList `json:"certificate-issuer-starts-with-list,omitempty"`
	CertificateSanStarts             []CertificateSanStartsWithList    `json:"certificate-san-starts-with-list,omitempty"`
	ClientAuthEndsWith               []ClientAuthEndsWithList          `json:"client-auth-ends-with-list,omitempty"`
	CloseNotify                      int                               `json:"close-notify,omitempty"`
	ForwardProxyNoSharedCipherAction int                               `json:"forward-proxy-no-shared-cipher-action,omitempty"`
	ForwardProxyOcspDisable          int                               `json:"forward-proxy-ocsp-disable,omitempty"`
	Sslilogging                      string                            `json:"sslilogging,omitempty"`
	AuthUsername                     string                            `json:"auth-username,omitempty"`
	ExceptionUserNameList            string                            `json:"exception-user-name-list,omitempty"`
	OcspstSgDays                     int                               `json:"ocspst-sg-days,omitempty"`
	//KeyStr                              string                             `json:"key-str,omitempty"`
	InspectListName                 string                         `json:"inspect-list-name,omitempty"`
	AuthUsernameAttribute           string                         `json:"auth-username-attribute,omitempty"`
	FpCertFetchNatpoolName          string                         `json:"fp-cert-fetch-natpool-name,omitempty"`
	ExceptionSniClName              string                         `json:"exception-sni-cl-name,omitempty"`
	InspectCertificateSubjectClName string                         `json:"inspect-certificate-subject-cl-name,omitempty"`
	LdapBaseDnFromCert              int                            `json:"ldap-base-dn-from-cert,omitempty"`
	AdGroupList                     string                         `json:"ad-group-list,omitempty"`
	ClientCertificate               string                         `json:"client-certificate,omitempty"`
	ForwardProxyCertExpiry          int                            `json:"forward-proxy-cert-expiry,omitempty"`
	ForwardProxyEnable              int                            `json:"forward-proxy-enable,omitempty"`
	SharedPartitionPool             int                            `json:"shared-partition-pool,omitempty"`
	LdapSearchFilter                string                         `json:"ldap-search-filter,omitempty"`
	KeySharedEncrypted              string                         `json:"key-shared-encrypted,omitempty"`
	AuthSgFilter                    string                         `json:"auth-sg-filter,omitempty"`
	OcspstSrvrTimeout               int                            `json:"ocspst-srvr-timeout,omitempty"`
	CertificateSubjectEquals        []CertificateSubjectEqualsList `json:"certificate-subject-equals-list,omitempty"`
	ChainCertSharedStr              string                         `json:"chain-cert-shared-str,omitempty"`
	EnableTLSAlertLogging           int                            `json:"enable-tls-alert-logging,omitempty"`
	DhType                          string                         `json:"dh-type,omitempty"`
	FpAltCert                       string                         `json:"fp-alt-cert,omitempty"`
	CaseInsensitive                 int                            `json:"case-insensitive,omitempty"`
	CipherWoPrio                    []CipherWithoutPrioList2       `json:"cipher-without-prio-list,omitempty"`
	OcspstSgMinutes                 int                            `json:"ocspst-sg-minutes,omitempty"`
	StartsWith                      []StartsWithList               `json:"starts-with-list,omitempty"`
	KeySharedPassphrase             string                         `json:"key-shared-passphrase,omitempty"`
}

func PostSlbTemplateClientSSL(id string, inst ClientSSL, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateClientSSL")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/client-ssl", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSL
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateClientSSL REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateClientSSL", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateClientSSL(id string, name string, host string) (*ClientSSL, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateClientSSL")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/client-ssl/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSL
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateClientSSL REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateClientSSL", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateClientSSL(id string, name string, inst ClientSSL, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateClientSSL")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/client-ssl/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSL
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateClientSSL REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateClientSSL", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateClientSSL(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateClientSSL")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/client-ssl/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSL
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
