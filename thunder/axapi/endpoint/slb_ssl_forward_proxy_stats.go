

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslForwardProxyStats struct {
    
    Stats SlbSslForwardProxyStatsStats `json:"stats"`

}
type DataSlbSslForwardProxyStats struct {
    DtSlbSslForwardProxyStats SlbSslForwardProxyStats `json:"ssl-forward-proxy"`
}


type SlbSslForwardProxyStatsStats struct {
    Cert_create int `json:"cert_create"`
    Cert_expr int `json:"cert_expr"`
    Cert_hit int `json:"cert_hit"`
    Cert_miss int `json:"cert_miss"`
    Conn_bypass int `json:"conn_bypass"`
    Conn_inspect int `json:"conn_inspect"`
    BypassFailsafeSslSessions int `json:"bypass-failsafe-ssl-sessions"`
    BypassSniSessions int `json:"bypass-sni-sessions"`
    BypassClientAuthSessions int `json:"bypass-client-auth-sessions"`
    FailedInSslHandshakes int `json:"failed-in-ssl-handshakes"`
    FailedInCryptoOperations int `json:"failed-in-crypto-operations"`
    FailedInTcp int `json:"failed-in-tcp"`
    FailedInCertificateVerification int `json:"failed-in-certificate-verification"`
    FailedInCertificateSigning int `json:"failed-in-certificate-signing"`
    InvalidOcspStaplingResponse int `json:"invalid-ocsp-stapling-response"`
    RevokedOcspResponse int `json:"revoked-ocsp-response"`
    UnsupportedSslVersion int `json:"unsupported-ssl-version"`
    CertificatesInCache int `json:"certificates-in-cache"`
    ConnectionsFailed int `json:"connections-failed"`
    AflexBypass int `json:"aflex-bypass"`
    BypassCertSubjectSessions int `json:"bypass-cert-subject-sessions"`
    BypassCertIssuerSessions int `json:"bypass-cert-issuer-sessions"`
    BypassCertSanSessions int `json:"bypass-cert-san-sessions"`
    BypassNoSniSessions int `json:"bypass-no-sni-sessions"`
    ResetNoSniSessions int `json:"reset-no-sni-sessions"`
    BypassEsniSessions int `json:"bypass-esni-sessions"`
    DropEsniSessions int `json:"drop-esni-sessions"`
    BypassUsernameSessions int `json:"bypass-username-sessions"`
    BypassAdGroupSessions int `json:"bypass-ad-group-sessions"`
    Tot_conn_in_buff int `json:"tot_conn_in_buff"`
    Curr_conn_in_buff int `json:"curr_conn_in_buff"`
    Async_conn_timeout int `json:"async_conn_timeout"`
    Async_conn_limit_drop int `json:"async_conn_limit_drop"`
    Cert_in_cache int `json:"cert_in_cache"`
    BypassClientIpSessions int `json:"bypass-client-ip-sessions"`
    BypassServerIpSessions int `json:"bypass-server-ip-sessions"`
}

func (p *SlbSslForwardProxyStats) GetId() string{
    return "1"
}

func (p *SlbSslForwardProxyStats) getPath() string{
    return "slb/ssl-forward-proxy/stats"
}

func (p *SlbSslForwardProxyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslForwardProxyStats,error) {
logger.Println("SlbSslForwardProxyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslForwardProxyStats
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
