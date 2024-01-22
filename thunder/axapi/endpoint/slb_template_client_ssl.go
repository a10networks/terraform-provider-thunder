

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateClientSsl struct {
	Inst struct {

    AdGroupList string `json:"ad-group-list"`
    AlertType string `json:"alert-type"`
    AuthSg string `json:"auth-sg"`
    AuthSgDn int `json:"auth-sg-dn"`
    AuthSgFilter string `json:"auth-sg-filter"`
    AuthUsername string `json:"auth-username"`
    AuthUsernameAttribute string `json:"auth-username-attribute"`
    AuthenName string `json:"authen-name"`
    Authorization int `json:"authorization"`
    BypassCertIssuerClassListName string `json:"bypass-cert-issuer-class-list-name"`
    BypassCertIssuerMultiClassList []SlbTemplateClientSslBypassCertIssuerMultiClassList `json:"bypass-cert-issuer-multi-class-list"`
    BypassCertSanClassListName string `json:"bypass-cert-san-class-list-name"`
    BypassCertSanMultiClassList []SlbTemplateClientSslBypassCertSanMultiClassList `json:"bypass-cert-san-multi-class-list"`
    BypassCertSubjectClassListName string `json:"bypass-cert-subject-class-list-name"`
    BypassCertSubjectMultiClassList []SlbTemplateClientSslBypassCertSubjectMultiClassList `json:"bypass-cert-subject-multi-class-list"`
    CaCerts []SlbTemplateClientSslCaCerts `json:"ca-certs"`
    CachePersistenceListName string `json:"cache-persistence-list-name"`
    CaseInsensitive int `json:"case-insensitive"`
    CentralCertPinList int `json:"central-cert-pin-list"`
    CertRevokeAction string `json:"cert-revoke-action" dval:"bypass"`
    CertUnknownAction string `json:"cert-unknown-action" dval:"bypass"`
    CertificateIssuerContainsList []SlbTemplateClientSslCertificateIssuerContainsList `json:"certificate-issuer-contains-list"`
    CertificateIssuerEndsWithList []SlbTemplateClientSslCertificateIssuerEndsWithList `json:"certificate-issuer-ends-with-list"`
    CertificateIssuerEqualsList []SlbTemplateClientSslCertificateIssuerEqualsList `json:"certificate-issuer-equals-list"`
    CertificateIssuerStartsWithList []SlbTemplateClientSslCertificateIssuerStartsWithList `json:"certificate-issuer-starts-with-list"`
    CertificateList []SlbTemplateClientSslCertificateList `json:"certificate-list"`
    CertificateSanContainsList []SlbTemplateClientSslCertificateSanContainsList `json:"certificate-san-contains-list"`
    CertificateSanEndsWithList []SlbTemplateClientSslCertificateSanEndsWithList `json:"certificate-san-ends-with-list"`
    CertificateSanEqualsList []SlbTemplateClientSslCertificateSanEqualsList `json:"certificate-san-equals-list"`
    CertificateSanStartsWithList []SlbTemplateClientSslCertificateSanStartsWithList `json:"certificate-san-starts-with-list"`
    CertificateSubjectContainsList []SlbTemplateClientSslCertificateSubjectContainsList `json:"certificate-subject-contains-list"`
    CertificateSubjectEndsWithList []SlbTemplateClientSslCertificateSubjectEndsWithList `json:"certificate-subject-ends-with-list"`
    CertificateSubjectEqualsList []SlbTemplateClientSslCertificateSubjectEqualsList `json:"certificate-subject-equals-list"`
    CertificateSubjectStartsWithList []SlbTemplateClientSslCertificateSubjectStartsWithList `json:"certificate-subject-starts-with-list"`
    ChainCert string `json:"chain-cert"`
    ChainCertSharedStr string `json:"chain-cert-shared-str"`
    CipherWithoutPrioList []SlbTemplateClientSslCipherWithoutPrioList `json:"cipher-without-prio-list"`
    ClassListName string `json:"class-list-name"`
    ClientAuthCaseInsensitive int `json:"client-auth-case-insensitive"`
    ClientAuthClassList string `json:"client-auth-class-list"`
    ClientAuthContainsList []SlbTemplateClientSslClientAuthContainsList `json:"client-auth-contains-list"`
    ClientAuthEndsWithList []SlbTemplateClientSslClientAuthEndsWithList `json:"client-auth-ends-with-list"`
    ClientAuthEqualsList []SlbTemplateClientSslClientAuthEqualsList `json:"client-auth-equals-list"`
    ClientAuthStartsWithList []SlbTemplateClientSslClientAuthStartsWithList `json:"client-auth-starts-with-list"`
    ClientCertificate string `json:"client-certificate" dval:"Ignore"`
    ClientIpv4List []SlbTemplateClientSslClientIpv4List `json:"client-ipv4-list"`
    ClientIpv6List []SlbTemplateClientSslClientIpv6List `json:"client-ipv6-list"`
    CloseNotify int `json:"close-notify"`
    ContainsList []SlbTemplateClientSslContainsList `json:"contains-list"`
    CrlCerts []SlbTemplateClientSslCrlCerts `json:"crl-certs"`
    Dgversion int `json:"dgversion" dval:"31"`
    DhType string `json:"dh-type"`
    DirectClientServerAuth int `json:"direct-client-server-auth"`
    DisableSslv3 int `json:"disable-sslv3"`
    EarlyData int `json:"early-data"`
    EcList []SlbTemplateClientSslEcList `json:"ec-list"`
    EnableSsliFtpAlg int `json:"enable-ssli-ftp-alg"`
    EnableTlsAlertLogging int `json:"enable-tls-alert-logging"`
    EndsWithList []SlbTemplateClientSslEndsWithList `json:"ends-with-list"`
    EqualsList []SlbTemplateClientSslEqualsList `json:"equals-list"`
    ExceptionAdGroupList string `json:"exception-ad-group-list"`
    ExceptionCertificateIssuerClName string `json:"exception-certificate-issuer-cl-name"`
    ExceptionCertificateSanClName string `json:"exception-certificate-san-cl-name"`
    ExceptionCertificateSubjectClName string `json:"exception-certificate-subject-cl-name"`
    ExceptionClientIpv4List []SlbTemplateClientSslExceptionClientIpv4List `json:"exception-client-ipv4-list"`
    ExceptionClientIpv6List []SlbTemplateClientSslExceptionClientIpv6List `json:"exception-client-ipv6-list"`
    ExceptionServerIpv4List []SlbTemplateClientSslExceptionServerIpv4List `json:"exception-server-ipv4-list"`
    ExceptionServerIpv6List []SlbTemplateClientSslExceptionServerIpv6List `json:"exception-server-ipv6-list"`
    ExceptionSniClName string `json:"exception-sni-cl-name"`
    ExceptionUserNameList string `json:"exception-user-name-list"`
    ExceptionWebCategory SlbTemplateClientSslExceptionWebCategory `json:"exception-web-category"`
    ExceptionWebReputation SlbTemplateClientSslExceptionWebReputation `json:"exception-web-reputation"`
    ExpireHours int `json:"expire-hours"`
    ForwardEncrypted string `json:"forward-encrypted"`
    ForwardPassphrase string `json:"forward-passphrase"`
    ForwardProxyAltSign int `json:"forward-proxy-alt-sign"`
    ForwardProxyBlockMessage string `json:"forward-proxy-block-message"`
    ForwardProxyCaCert string `json:"forward-proxy-ca-cert"`
    ForwardProxyCaKey string `json:"forward-proxy-ca-key"`
    ForwardProxyCertCacheLimit int `json:"forward-proxy-cert-cache-limit" dval:"524288"`
    ForwardProxyCertCacheTimeout int `json:"forward-proxy-cert-cache-timeout" dval:"3600"`
    ForwardProxyCertExpiry int `json:"forward-proxy-cert-expiry"`
    ForwardProxyCertNotReadyAction string `json:"forward-proxy-cert-not-ready-action" dval:"bypass"`
    ForwardProxyCertRevokeAction int `json:"forward-proxy-cert-revoke-action" dval:"1"`
    ForwardProxyCertUnknownAction int `json:"forward-proxy-cert-unknown-action" dval:"1"`
    ForwardProxyCrlDisable int `json:"forward-proxy-crl-disable"`
    ForwardProxyDecryptedDscp int `json:"forward-proxy-decrypted-dscp"`
    ForwardProxyDecryptedDscpBypass int `json:"forward-proxy-decrypted-dscp-bypass"`
    ForwardProxyEnable int `json:"forward-proxy-enable"`
    ForwardProxyEsniAction int `json:"forward-proxy-esni-action"`
    ForwardProxyFailsafeDisable int `json:"forward-proxy-failsafe-disable"`
    ForwardProxyHashPersistenceInterval int `json:"forward-proxy-hash-persistence-interval" dval:"30"`
    ForwardProxyLogDisable int `json:"forward-proxy-log-disable"`
    ForwardProxyNoSharedCipherAction int `json:"forward-proxy-no-shared-cipher-action" dval:"1"`
    ForwardProxyNoSniAction string `json:"forward-proxy-no-sni-action" dval:"intercept"`
    ForwardProxyOcspDisable int `json:"forward-proxy-ocsp-disable"`
    ForwardProxyRequireSniCertMatched string `json:"forward-proxy-require-sni-cert-matched"`
    ForwardProxySelfsignRedir int `json:"forward-proxy-selfsign-redir"`
    ForwardProxySslVersion int `json:"forward-proxy-ssl-version" dval:"33"`
    ForwardProxyTrustedCaLists []SlbTemplateClientSslForwardProxyTrustedCaLists `json:"forward-proxy-trusted-ca-lists"`
    ForwardProxyVerifyCertFailAction int `json:"forward-proxy-verify-cert-fail-action" dval:"1"`
    FpAltCert string `json:"fp-alt-cert"`
    FpAltChainCert string `json:"fp-alt-chain-cert"`
    FpAltEncrypted string `json:"fp-alt-encrypted"`
    FpAltKey string `json:"fp-alt-key"`
    FpAltPassphrase string `json:"fp-alt-passphrase"`
    FpAltShared int `json:"fp-alt-shared"`
    FpCaCertificate string `json:"fp-ca-certificate"`
    FpCaCertificateShared int `json:"fp-ca-certificate-shared"`
    FpCaChainCert string `json:"fp-ca-chain-cert"`
    FpCaKey string `json:"fp-ca-key"`
    FpCaKeyEncrypted string `json:"fp-ca-key-encrypted"`
    FpCaKeyPassphrase string `json:"fp-ca-key-passphrase"`
    FpCaKeyShared int `json:"fp-ca-key-shared"`
    FpCaShared int `json:"fp-ca-shared"`
    FpCertExtAiaCaIssuers string `json:"fp-cert-ext-aia-ca-issuers"`
    FpCertExtAiaOcsp string `json:"fp-cert-ext-aia-ocsp"`
    FpCertExtCrldp string `json:"fp-cert-ext-crldp"`
    FpCertFetchAutonat string `json:"fp-cert-fetch-autonat"`
    FpCertFetchAutonatPrecedence int `json:"fp-cert-fetch-autonat-precedence"`
    FpCertFetchNatpoolName string `json:"fp-cert-fetch-natpool-name"`
    FpCertFetchNatpoolNameShared string `json:"fp-cert-fetch-natpool-name-shared"`
    FpCertFetchNatpoolPrecedence int `json:"fp-cert-fetch-natpool-precedence"`
    FpEsniAction string `json:"fp-esni-action" dval:"bypass"`
    HandshakeLoggingEnable int `json:"handshake-logging-enable"`
    HsmType string `json:"hsm-type"`
    InspectCertificateIssuerClName string `json:"inspect-certificate-issuer-cl-name"`
    InspectCertificateSanClName string `json:"inspect-certificate-san-cl-name"`
    InspectCertificateSubjectClName string `json:"inspect-certificate-subject-cl-name"`
    InspectListName string `json:"inspect-list-name"`
    Ja3Enable int `json:"ja3-enable"`
    Ja3InsertHttpHeader string `json:"ja3-insert-http-header"`
    Ja3RejectClassList string `json:"ja3-reject-class-list"`
    Ja3RejectMaxNumberPerHost int `json:"ja3-reject-max-number-per-host"`
    Ja3Ttl int `json:"ja3-ttl" dval:"600"`
    LdapBaseDnFromCert int `json:"ldap-base-dn-from-cert"`
    LdapSearchFilter string `json:"ldap-search-filter"`
    LocalCertPinList SlbTemplateClientSslLocalCertPinList `json:"local-cert-pin-list"`
    LocalLogging int `json:"local-logging"`
    MultiClassList []SlbTemplateClientSslMultiClassList `json:"multi-class-list"`
    Name string `json:"name"`
    NoAntiReplay int `json:"no-anti-replay"`
    NoSharedCipherAction string `json:"no-shared-cipher-action" dval:"drop"`
    NonSslBypassL4session int `json:"non-ssl-bypass-l4session"`
    NonSslBypassServiceGroup string `json:"non-ssl-bypass-service-group"`
    Notafter int `json:"notafter"`
    Notafterday int `json:"notafterday"`
    Notaftermonth int `json:"notaftermonth"`
    Notafteryear int `json:"notafteryear"`
    Notbefore int `json:"notbefore"`
    Notbeforeday int `json:"notbeforeday"`
    Notbeforemonth int `json:"notbeforemonth"`
    Notbeforeyear int `json:"notbeforeyear"`
    OcspStapling int `json:"ocsp-stapling"`
    OcspstCaCert string `json:"ocspst-ca-cert"`
    OcspstOcsp int `json:"ocspst-ocsp"`
    OcspstSg string `json:"ocspst-sg"`
    OcspstSgDays int `json:"ocspst-sg-days"`
    OcspstSgHours int `json:"ocspst-sg-hours" dval:"1"`
    OcspstSgMinutes int `json:"ocspst-sg-minutes"`
    OcspstSgTimeout int `json:"ocspst-sg-timeout" dval:"30"`
    OcspstSrvr string `json:"ocspst-srvr"`
    OcspstSrvrDays int `json:"ocspst-srvr-days"`
    OcspstSrvrHours int `json:"ocspst-srvr-hours" dval:"1"`
    OcspstSrvrMinutes int `json:"ocspst-srvr-minutes"`
    OcspstSrvrTimeout int `json:"ocspst-srvr-timeout" dval:"30"`
    RenegotiationDisable int `json:"renegotiation-disable"`
    ReqCaLists []SlbTemplateClientSslReqCaLists `json:"req-ca-lists"`
    RequireWebCategory int `json:"require-web-category"`
    ServerIpv4List []SlbTemplateClientSslServerIpv4List `json:"server-ipv4-list"`
    ServerIpv6List []SlbTemplateClientSslServerIpv6List `json:"server-ipv6-list"`
    ServerNameAutoMap int `json:"server-name-auto-map"`
    ServerNameList []SlbTemplateClientSslServerNameList `json:"server-name-list"`
    SessionCacheSize int `json:"session-cache-size"`
    SessionCacheTimeout int `json:"session-cache-timeout"`
    SessionTicketDisable int `json:"session-ticket-disable"`
    SessionTicketLifetime int `json:"session-ticket-lifetime"`
    SharedPartitionCipherTemplate int `json:"shared-partition-cipher-template"`
    SharedPartitionPool int `json:"shared-partition-pool"`
    SniBypassEnableLog int `json:"sni-bypass-enable-log"`
    SniBypassExpiredCert int `json:"sni-bypass-expired-cert"`
    SniBypassExplicitList string `json:"sni-bypass-explicit-list"`
    SniBypassMissingCert int `json:"sni-bypass-missing-cert"`
    SniEnableLog int `json:"sni-enable-log"`
    SslFalseStartDisable int `json:"ssl-false-start-disable"`
    SsliInboundEnable int `json:"ssli-inbound-enable"`
    SsliLogging int `json:"ssli-logging"`
    Sslilogging string `json:"sslilogging"`
    Sslv2BypassServiceGroup string `json:"sslv2-bypass-service-group"`
    StartsWithList []SlbTemplateClientSslStartsWithList `json:"starts-with-list"`
    TemplateCipher string `json:"template-cipher"`
    TemplateCipherShared string `json:"template-cipher-shared"`
    TemplateHsm string `json:"template-hsm"`
    UserNameList string `json:"user-name-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VerifyCertFailAction string `json:"verify-cert-fail-action" dval:"drop"`
    Version int `json:"version"`
    WebCategory SlbTemplateClientSslWebCategory `json:"web-category"`
    WebReputation SlbTemplateClientSslWebReputation `json:"web-reputation"`

	} `json:"client-ssl"`
}


type SlbTemplateClientSslBypassCertIssuerMultiClassList struct {
    BypassCertIssuerMultiClassListName string `json:"bypass-cert-issuer-multi-class-list-name"`
}


type SlbTemplateClientSslBypassCertSanMultiClassList struct {
    BypassCertSanMultiClassListName string `json:"bypass-cert-san-multi-class-list-name"`
}


type SlbTemplateClientSslBypassCertSubjectMultiClassList struct {
    BypassCertSubjectMultiClassListName string `json:"bypass-cert-subject-multi-class-list-name"`
}


type SlbTemplateClientSslCaCerts struct {
    CaCert string `json:"ca-cert"`
    CaShared int `json:"ca-shared"`
    ClientOcsp int `json:"client-ocsp"`
    ClientOcspSrvr string `json:"client-ocsp-srvr"`
    ClientOcspSg string `json:"client-ocsp-sg"`
}


type SlbTemplateClientSslCertificateIssuerContainsList struct {
    CertificateIssuerContains string `json:"certificate-issuer-contains"`
}


type SlbTemplateClientSslCertificateIssuerEndsWithList struct {
    CertificateIssuerEndsWith string `json:"certificate-issuer-ends-with"`
}


type SlbTemplateClientSslCertificateIssuerEqualsList struct {
    CertificateIssuerEquals string `json:"certificate-issuer-equals"`
}


type SlbTemplateClientSslCertificateIssuerStartsWithList struct {
    CertificateIssuerStarts string `json:"certificate-issuer-starts"`
}


type SlbTemplateClientSslCertificateList struct {
    Cert string `json:"cert"`
    Key string `json:"key"`
    Passphrase string `json:"passphrase"`
    KeyEncrypted string `json:"key-encrypted"`
    ChainCert string `json:"chain-cert"`
    Shared int `json:"shared"`
    Uuid string `json:"uuid"`
}


type SlbTemplateClientSslCertificateSanContainsList struct {
    CertificateSanContains string `json:"certificate-san-contains"`
}


type SlbTemplateClientSslCertificateSanEndsWithList struct {
    CertificateSanEndsWith string `json:"certificate-san-ends-with"`
}


type SlbTemplateClientSslCertificateSanEqualsList struct {
    CertificateSanEquals string `json:"certificate-san-equals"`
}


type SlbTemplateClientSslCertificateSanStartsWithList struct {
    CertificateSanStarts string `json:"certificate-san-starts"`
}


type SlbTemplateClientSslCertificateSubjectContainsList struct {
    CertificateSubjectContains string `json:"certificate-subject-contains"`
}


type SlbTemplateClientSslCertificateSubjectEndsWithList struct {
    CertificateSubjectEndsWith string `json:"certificate-subject-ends-with"`
}


type SlbTemplateClientSslCertificateSubjectEqualsList struct {
    CertificateSubjectEquals string `json:"certificate-subject-equals"`
}


type SlbTemplateClientSslCertificateSubjectStartsWithList struct {
    CertificateSubjectStarts string `json:"certificate-subject-starts"`
}


type SlbTemplateClientSslCipherWithoutPrioList struct {
    CipherWoPrio string `json:"cipher-wo-prio"`
}


type SlbTemplateClientSslClientAuthContainsList struct {
    ClientAuthContains string `json:"client-auth-contains"`
}


type SlbTemplateClientSslClientAuthEndsWithList struct {
    ClientAuthEndsWith string `json:"client-auth-ends-with"`
}


type SlbTemplateClientSslClientAuthEqualsList struct {
    ClientAuthEquals string `json:"client-auth-equals"`
}


type SlbTemplateClientSslClientAuthStartsWithList struct {
    ClientAuthStartsWith string `json:"client-auth-starts-with"`
}


type SlbTemplateClientSslClientIpv4List struct {
    ClientIpv4ListName string `json:"client-ipv4-list-name"`
}


type SlbTemplateClientSslClientIpv6List struct {
    ClientIpv6ListName string `json:"client-ipv6-list-name"`
}


type SlbTemplateClientSslContainsList struct {
    Contains string `json:"contains"`
}


type SlbTemplateClientSslCrlCerts struct {
    Crl string `json:"crl"`
    CrlShared int `json:"crl-shared"`
}


type SlbTemplateClientSslEcList struct {
    Ec string `json:"ec"`
}


type SlbTemplateClientSslEndsWithList struct {
    EndsWith string `json:"ends-with"`
}


type SlbTemplateClientSslEqualsList struct {
    Equals string `json:"equals"`
}


type SlbTemplateClientSslExceptionClientIpv4List struct {
    ExceptionClientIpv4ListName string `json:"exception-client-ipv4-list-name"`
}


type SlbTemplateClientSslExceptionClientIpv6List struct {
    ExceptionClientIpv6ListName string `json:"exception-client-ipv6-list-name"`
}


type SlbTemplateClientSslExceptionServerIpv4List struct {
    ExceptionServerIpv4ListName string `json:"exception-server-ipv4-list-name"`
}


type SlbTemplateClientSslExceptionServerIpv6List struct {
    ExceptionServerIpv6ListName string `json:"exception-server-ipv6-list-name"`
}


type SlbTemplateClientSslExceptionWebCategory struct {
    ExceptionUncategorized int `json:"exception-uncategorized"`
    ExceptionRealEstate int `json:"exception-real-estate"`
    ExceptionComputerAndInternetSecurity int `json:"exception-computer-and-internet-security"`
    ExceptionFinancialServices int `json:"exception-financial-services"`
    ExceptionBusinessAndEconomy int `json:"exception-business-and-economy"`
    ExceptionComputerAndInternetInfo int `json:"exception-computer-and-internet-info"`
    ExceptionAuctions int `json:"exception-auctions"`
    ExceptionShopping int `json:"exception-shopping"`
    ExceptionCultAndOccult int `json:"exception-cult-and-occult"`
    ExceptionTravel int `json:"exception-travel"`
    ExceptionDrugs int `json:"exception-drugs"`
    ExceptionAdultAndPornography int `json:"exception-adult-and-pornography"`
    ExceptionHomeAndGarden int `json:"exception-home-and-garden"`
    ExceptionMilitary int `json:"exception-military"`
    ExceptionSocialNetwork int `json:"exception-social-network"`
    ExceptionDeadSites int `json:"exception-dead-sites"`
    ExceptionStockAdviceAndTools int `json:"exception-stock-advice-and-tools"`
    ExceptionTrainingAndTools int `json:"exception-training-and-tools"`
    ExceptionDating int `json:"exception-dating"`
    ExceptionSexEducation int `json:"exception-sex-education"`
    ExceptionReligion int `json:"exception-religion"`
    ExceptionEntertainmentAndArts int `json:"exception-entertainment-and-arts"`
    ExceptionPersonalSitesAndBlogs int `json:"exception-personal-sites-and-blogs"`
    ExceptionLegal int `json:"exception-legal"`
    ExceptionLocalInformation int `json:"exception-local-information"`
    ExceptionStreamingMedia int `json:"exception-streaming-media"`
    ExceptionJobSearch int `json:"exception-job-search"`
    ExceptionGambling int `json:"exception-gambling"`
    ExceptionTranslation int `json:"exception-translation"`
    ExceptionReferenceAndResearch int `json:"exception-reference-and-research"`
    ExceptionSharewareAndFreeware int `json:"exception-shareware-and-freeware"`
    ExceptionPeerToPeer int `json:"exception-peer-to-peer"`
    ExceptionMarijuana int `json:"exception-marijuana"`
    ExceptionHacking int `json:"exception-hacking"`
    ExceptionGames int `json:"exception-games"`
    ExceptionPhilosophyAndPolitics int `json:"exception-philosophy-and-politics"`
    ExceptionWeapons int `json:"exception-weapons"`
    ExceptionPayToSurf int `json:"exception-pay-to-surf"`
    ExceptionHuntingAndFishing int `json:"exception-hunting-and-fishing"`
    ExceptionSociety int `json:"exception-society"`
    ExceptionEducationalInstitutions int `json:"exception-educational-institutions"`
    ExceptionOnlineGreetingCards int `json:"exception-online-greeting-cards"`
    ExceptionSports int `json:"exception-sports"`
    ExceptionSwimsuitsAndIntimateApparel int `json:"exception-swimsuits-and-intimate-apparel"`
    ExceptionQuestionable int `json:"exception-questionable"`
    ExceptionKids int `json:"exception-kids"`
    ExceptionHateAndRacism int `json:"exception-hate-and-racism"`
    ExceptionPersonalStorage int `json:"exception-personal-storage"`
    ExceptionViolence int `json:"exception-violence"`
    ExceptionKeyloggersAndMonitoring int `json:"exception-keyloggers-and-monitoring"`
    ExceptionSearchEngines int `json:"exception-search-engines"`
    ExceptionInternetPortals int `json:"exception-internet-portals"`
    ExceptionWebAdvertisements int `json:"exception-web-advertisements"`
    ExceptionCheating int `json:"exception-cheating"`
    ExceptionGross int `json:"exception-gross"`
    ExceptionWebBasedEmail int `json:"exception-web-based-email"`
    ExceptionMalwareSites int `json:"exception-malware-sites"`
    ExceptionPhishingAndOtherFraud int `json:"exception-phishing-and-other-fraud"`
    ExceptionProxyAvoidAndAnonymizers int `json:"exception-proxy-avoid-and-anonymizers"`
    ExceptionSpywareAndAdware int `json:"exception-spyware-and-adware"`
    ExceptionMusic int `json:"exception-music"`
    ExceptionGovernment int `json:"exception-government"`
    ExceptionNudity int `json:"exception-nudity"`
    ExceptionNewsAndMedia int `json:"exception-news-and-media"`
    ExceptionIllegal int `json:"exception-illegal"`
    ExceptionCdns int `json:"exception-cdns"`
    ExceptionInternetCommunications int `json:"exception-internet-communications"`
    ExceptionBotNets int `json:"exception-bot-nets"`
    ExceptionAbortion int `json:"exception-abortion"`
    ExceptionHealthAndMedicine int `json:"exception-health-and-medicine"`
    ExceptionSpamUrls int `json:"exception-spam-urls"`
    ExceptionDynamicallyGeneratedContent int `json:"exception-dynamically-generated-content"`
    ExceptionParkedDomains int `json:"exception-parked-domains"`
    ExceptionAlcoholAndTobacco int `json:"exception-alcohol-and-tobacco"`
    ExceptionImageAndVideoSearch int `json:"exception-image-and-video-search"`
    ExceptionFashionAndBeauty int `json:"exception-fashion-and-beauty"`
    ExceptionRecreationAndHobbies int `json:"exception-recreation-and-hobbies"`
    ExceptionMotorVehicles int `json:"exception-motor-vehicles"`
    ExceptionWebHostingSites int `json:"exception-web-hosting-sites"`
    ExceptionNudityArtistic int `json:"exception-nudity-artistic"`
    ExceptionIllegalPornography int `json:"exception-illegal-pornography"`
}


type SlbTemplateClientSslExceptionWebReputation struct {
    ExceptionTrustworthy int `json:"exception-trustworthy"`
    ExceptionLowRisk int `json:"exception-low-risk"`
    ExceptionModerateRisk int `json:"exception-moderate-risk"`
    ExceptionSuspicious int `json:"exception-suspicious"`
    ExceptionMalicious int `json:"exception-malicious"`
    ExceptionThreshold int `json:"exception-threshold"`
}


type SlbTemplateClientSslForwardProxyTrustedCaLists struct {
    ForwardProxyTrustedCa string `json:"forward-proxy-trusted-ca"`
    FpTrustedCaShared int `json:"fp-trusted-ca-shared"`
}


type SlbTemplateClientSslLocalCertPinList struct {
    LocalCertPinListBypassFailCount int `json:"local-cert-pin-list-bypass-fail-count"`
}


type SlbTemplateClientSslMultiClassList struct {
    MultiClistName string `json:"multi-clist-name"`
}


type SlbTemplateClientSslReqCaLists struct {
    ClientCertificateRequestCa string `json:"client-certificate-Request-CA"`
    ClientCertReqCaShared int `json:"client-cert-req-ca-shared"`
}


type SlbTemplateClientSslServerIpv4List struct {
    ServerIpv4ListName string `json:"server-ipv4-list-name"`
}


type SlbTemplateClientSslServerIpv6List struct {
    ServerIpv6ListName string `json:"server-ipv6-list-name"`
}


type SlbTemplateClientSslServerNameList struct {
    ServerName string `json:"server-name"`
    ServerCert string `json:"server-cert"`
    ServerChain string `json:"server-chain"`
    ServerKey string `json:"server-key"`
    ServerPassphrase string `json:"server-passphrase"`
    ServerEncrypted string `json:"server-encrypted"`
    ServerNameAlternate int `json:"server-name-alternate"`
    ServerShared int `json:"server-shared"`
    SniTemplate int `json:"sni-template"`
    SniTemplateClientSsl string `json:"sni-template-client-ssl"`
    SniSharedPartitionClientSslTemplate int `json:"sni-shared-partition-client-ssl-template"`
    SniTemplateClientSslSharedName string `json:"sni-template-client-ssl-shared-name"`
    ServerNameRegex string `json:"server-name-regex"`
    ServerCertRegex string `json:"server-cert-regex"`
    ServerChainRegex string `json:"server-chain-regex"`
    ServerKeyRegex string `json:"server-key-regex"`
    ServerPassphraseRegex string `json:"server-passphrase-regex"`
    ServerEncryptedRegex string `json:"server-encrypted-regex"`
    ServerNameRegexAlternate int `json:"server-name-regex-alternate"`
    ServerSharedRegex int `json:"server-shared-regex"`
    SniRegexTemplate int `json:"sni-regex-template"`
    SniRegexTemplateClientSsl string `json:"sni-regex-template-client-ssl"`
    SniRegexSharedPartitionClientSslTemplate int `json:"sni-regex-shared-partition-client-ssl-template"`
    SniRegexTemplateClientSslSharedName string `json:"sni-regex-template-client-ssl-shared-name"`
}


type SlbTemplateClientSslStartsWithList struct {
    StartsWith string `json:"starts-with"`
}


type SlbTemplateClientSslWebCategory struct {
    Uncategorized int `json:"uncategorized"`
    RealEstate int `json:"real-estate"`
    ComputerAndInternetSecurity int `json:"computer-and-internet-security"`
    FinancialServices int `json:"financial-services"`
    BusinessAndEconomy int `json:"business-and-economy"`
    ComputerAndInternetInfo int `json:"computer-and-internet-info"`
    Auctions int `json:"auctions"`
    Shopping int `json:"shopping"`
    CultAndOccult int `json:"cult-and-occult"`
    Travel int `json:"travel"`
    Drugs int `json:"drugs"`
    AdultAndPornography int `json:"adult-and-pornography"`
    HomeAndGarden int `json:"home-and-garden"`
    Military int `json:"military"`
    SocialNetwork int `json:"social-network"`
    DeadSites int `json:"dead-sites"`
    StockAdviceAndTools int `json:"stock-advice-and-tools"`
    TrainingAndTools int `json:"training-and-tools"`
    Dating int `json:"dating"`
    SexEducation int `json:"sex-education"`
    Religion int `json:"religion"`
    EntertainmentAndArts int `json:"entertainment-and-arts"`
    PersonalSitesAndBlogs int `json:"personal-sites-and-blogs"`
    Legal int `json:"legal"`
    LocalInformation int `json:"local-information"`
    StreamingMedia int `json:"streaming-media"`
    JobSearch int `json:"job-search"`
    Gambling int `json:"gambling"`
    Translation int `json:"translation"`
    ReferenceAndResearch int `json:"reference-and-research"`
    SharewareAndFreeware int `json:"shareware-and-freeware"`
    PeerToPeer int `json:"peer-to-peer"`
    Marijuana int `json:"marijuana"`
    Hacking int `json:"hacking"`
    Games int `json:"games"`
    PhilosophyAndPolitics int `json:"philosophy-and-politics"`
    Weapons int `json:"weapons"`
    PayToSurf int `json:"pay-to-surf"`
    HuntingAndFishing int `json:"hunting-and-fishing"`
    Society int `json:"society"`
    EducationalInstitutions int `json:"educational-institutions"`
    OnlineGreetingCards int `json:"online-greeting-cards"`
    Sports int `json:"sports"`
    SwimsuitsAndIntimateApparel int `json:"swimsuits-and-intimate-apparel"`
    Questionable int `json:"questionable"`
    Kids int `json:"kids"`
    HateAndRacism int `json:"hate-and-racism"`
    PersonalStorage int `json:"personal-storage"`
    Violence int `json:"violence"`
    KeyloggersAndMonitoring int `json:"keyloggers-and-monitoring"`
    SearchEngines int `json:"search-engines"`
    InternetPortals int `json:"internet-portals"`
    WebAdvertisements int `json:"web-advertisements"`
    Cheating int `json:"cheating"`
    Gross int `json:"gross"`
    WebBasedEmail int `json:"web-based-email"`
    MalwareSites int `json:"malware-sites"`
    PhishingAndOtherFraud int `json:"phishing-and-other-fraud"`
    ProxyAvoidAndAnonymizers int `json:"proxy-avoid-and-anonymizers"`
    SpywareAndAdware int `json:"spyware-and-adware"`
    Music int `json:"music"`
    Government int `json:"government"`
    Nudity int `json:"nudity"`
    NewsAndMedia int `json:"news-and-media"`
    Illegal int `json:"illegal"`
    Cdns int `json:"cdns"`
    InternetCommunications int `json:"internet-communications"`
    BotNets int `json:"bot-nets"`
    Abortion int `json:"abortion"`
    HealthAndMedicine int `json:"health-and-medicine"`
    SpamUrls int `json:"spam-urls"`
    DynamicallyGeneratedContent int `json:"dynamically-generated-content"`
    ParkedDomains int `json:"parked-domains"`
    AlcoholAndTobacco int `json:"alcohol-and-tobacco"`
    ImageAndVideoSearch int `json:"image-and-video-search"`
    FashionAndBeauty int `json:"fashion-and-beauty"`
    RecreationAndHobbies int `json:"recreation-and-hobbies"`
    MotorVehicles int `json:"motor-vehicles"`
    WebHostingSites int `json:"web-hosting-sites"`
    NudityArtistic int `json:"nudity-artistic"`
    IllegalPornography int `json:"illegal-pornography"`
}


type SlbTemplateClientSslWebReputation struct {
    BypassTrustworthy int `json:"bypass-trustworthy"`
    BypassLowRisk int `json:"bypass-low-risk"`
    BypassModerateRisk int `json:"bypass-moderate-risk"`
    BypassSuspicious int `json:"bypass-suspicious"`
    BypassMalicious int `json:"bypass-malicious"`
    BypassThreshold int `json:"bypass-threshold"`
}

func (p *SlbTemplateClientSsl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateClientSsl) getPath() string{
    return "slb/template/client-ssl"
}

func (p *SlbTemplateClientSsl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSsl::Post")
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

func (p *SlbTemplateClientSsl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSsl::Get")
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
func (p *SlbTemplateClientSsl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSsl::Put")
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

func (p *SlbTemplateClientSsl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSsl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
