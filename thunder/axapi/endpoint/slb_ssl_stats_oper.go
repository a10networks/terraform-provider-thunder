

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslStatsOper struct {
    
    Oper SlbSslStatsOperOper `json:"oper"`

}
type DataSlbSslStatsOper struct {
    DtSlbSslStatsOper SlbSslStatsOper `json:"ssl-stats"`
}


type SlbSslStatsOperOper struct {
    SslModuleType string `json:"ssl-module-type"`
    ConfigModuleType string `json:"config-module-type"`
    ThalesHsmStatus string `json:"thales-hsm-status"`
    SslModulesCount int `json:"ssl-modules-count"`
    CurrentClientsideConnections int `json:"current-clientside-connections"`
    TotalClientsideConnections int `json:"total-clientside-connections"`
    CurrentServersideConnections int `json:"current-serverside-connections"`
    TotalServersideConnections int `json:"total-serverside-connections"`
    NonSslBypassConnections int `json:"non-ssl-bypass-connections"`
    TotalReuseClientSsl int `json:"total-reuse-client-ssl"`
    TotalReuseServerSsl int `json:"total-reuse-server-ssl"`
    TotalReuseSessionTicketClientSsl int `json:"total-reuse-session-ticket-client-ssl"`
    TotalReuseSessionTicketServerSsl int `json:"total-reuse-session-ticket-server-ssl"`
    TotalClientsideEarlyDataConnections int `json:"total-clientside-early-data-connections"`
    TotalServersideEarlyDataConnections int `json:"total-serverside-early-data-connections"`
    TotalClientsideFailedEarlyDataConnections int `json:"total-clientside-failed-early-data-connections"`
    TotalServersideFailedEarlyDataConnections int `json:"total-serverside-failed-early-data-connections"`
    FailedHandshakes int `json:"failed-handshakes"`
    FailedCrypto int `json:"failed-crypto"`
    SslMemoryUsage int `json:"ssl-memory-usage"`
    ServerCertErrors int `json:"server-cert-errors"`
    ClientCertAuthFail int `json:"client-cert-auth-fail"`
    CaVerificationFailures int `json:"ca-verification-failures"`
    HwContextTotal int `json:"hw-context-total"`
    HwContextUsage int `json:"hw-context-usage"`
    HwContextAllocFail int `json:"hw-context-alloc-fail"`
    HwRingFull int `json:"hw-ring-full"`
    RecordTooBig int `json:"record-too-big"`
    ClientsslContextMallocFail int `json:"clientssl-context-malloc-fail"`
    MaxSslContexts int `json:"max-ssl-contexts"`
    CurrSslContexts int `json:"curr-ssl-contexts"`
    StaticSslContexts int `json:"static-ssl-contexts"`
    DynamicSslContexts int `json:"dynamic-ssl-contexts"`
    BypassFailsafeSessions int `json:"bypass-failsafe-sessions"`
    BypassUsenameSessions int `json:"bypass-usename-sessions"`
    BypassAdGroupSessions int `json:"bypass-ad-group-sessions"`
    BypassSniSessions int `json:"bypass-sni-sessions"`
    BypassCertSubjectSessions int `json:"bypass-cert-subject-sessions"`
    BypassCertIssuerSessions int `json:"bypass-cert-issuer-sessions"`
    BypassCertSanSessions int `json:"bypass-cert-san-sessions"`
    BypassNoSniSessions int `json:"bypass-no-sni-sessions"`
    ResetNoSniSessions int `json:"reset-no-sni-sessions"`
    BypassEsniSessions int `json:"bypass-esni-sessions"`
    DropEsniSessions int `json:"drop-esni-sessions"`
    BypassClientAuthSessions int `json:"bypass-client-auth-sessions"`
    SslHandshakeFailure int `json:"ssl-handshake-failure"`
    CryptioOperations int `json:"cryptio-operations"`
    TcpFailures int `json:"tcp-failures"`
    CertificateVerificationFailures int `json:"certificate-verification-failures"`
    CertificateSigningFailures int `json:"certificate-signing-failures"`
    InvalidOcspStaplingResponse int `json:"invalid-ocsp-stapling-response"`
    RevokedOcspResponse int `json:"revoked-ocsp-response"`
    UnsupportedSslVersion int `json:"unsupported-ssl-version"`
    ClientSslFatalAlert int `json:"client-ssl-fatal-alert"`
    ClientSslFinReset int `json:"client-ssl-fin-reset"`
    SslSessionFinReset int `json:"ssl-session-fin-reset"`
    ServerSslFatalAlert int `json:"server-ssl-fatal-alert"`
    ServerSslFinReset int `json:"server-ssl-fin-reset"`
    ClientSslInternalError int `json:"client-ssl-internal-error"`
    ClientSslUnknownError int `json:"client-ssl-unknown-error"`
    ServerSslInternalError int `json:"server-ssl-internal-error"`
    ServerSslUnknownError int `json:"server-ssl-unknown-error"`
    FpFailHandshake int `json:"fp-fail-handshake"`
    FpFailTcpError int `json:"fp-fail-tcp-error"`
    FpFailVerifyCert int `json:"fp-fail-verify-cert"`
    FpFailAflex int `json:"fp-fail-aflex"`
    UtilEnableStatus string `json:"util-enable-status"`
    SslModuleStats []SlbSslStatsOperOperSslModuleStats `json:"ssl-module-stats"`
}


type SlbSslStatsOperOperSslModuleStats struct {
    SslModulesIndex int `json:"ssl-modules-index"`
    TotalEnabledCryptoEngines int `json:"total-enabled-crypto-engines"`
    TotalAvailableCryptoEngines int `json:"total-available-crypto-engines"`
    RequestsHandled int `json:"requests-handled"`
    UtilPercentage int `json:"util-percentage"`
    SecRate1 int `json:"sec-rate1"`
    SecRate5 int `json:"sec-rate5"`
    SecRate10 int `json:"sec-rate10"`
    SecRate30 int `json:"sec-rate30"`
    SecRate60 int `json:"sec-rate60"`
}

func (p *SlbSslStatsOper) GetId() string{
    return "1"
}

func (p *SlbSslStatsOper) getPath() string{
    return "slb/ssl-stats/oper"
}

func (p *SlbSslStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslStatsOper,error) {
logger.Println("SlbSslStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslStatsOper
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
