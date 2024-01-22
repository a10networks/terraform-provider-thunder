

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCertRevokeStats struct {
    
    Stats SlbSslCertRevokeStatsStats `json:"stats"`

}
type DataSlbSslCertRevokeStats struct {
    DtSlbSslCertRevokeStats SlbSslCertRevokeStats `json:"ssl-cert-revoke"`
}


type SlbSslCertRevokeStatsStats struct {
    Ocsp_stapling_response_good int `json:"ocsp_stapling_response_good"`
    Ocsp_chain_status_good int `json:"ocsp_chain_status_good"`
    Ocsp_chain_status_revoked int `json:"ocsp_chain_status_revoked"`
    Ocsp_chain_status_unknown int `json:"ocsp_chain_status_unknown"`
    Ocsp_request int `json:"ocsp_request"`
    Ocsp_response int `json:"ocsp_response"`
    Ocsp_connection_error int `json:"ocsp_connection_error"`
    Ocsp_uri_not_found int `json:"ocsp_uri_not_found"`
    Ocsp_uri_https int `json:"ocsp_uri_https"`
    Ocsp_uri_unsupported int `json:"ocsp_uri_unsupported"`
    Ocsp_response_status_good int `json:"ocsp_response_status_good"`
    Ocsp_response_status_revoked int `json:"ocsp_response_status_revoked"`
    Ocsp_response_status_unknown int `json:"ocsp_response_status_unknown"`
    Ocsp_cache_status_good int `json:"ocsp_cache_status_good"`
    Ocsp_cache_status_revoked int `json:"ocsp_cache_status_revoked"`
    Ocsp_cache_miss int `json:"ocsp_cache_miss"`
    Ocsp_cache_expired int `json:"ocsp_cache_expired"`
    Ocsp_other_error int `json:"ocsp_other_error"`
    Ocsp_response_no_nonce int `json:"ocsp_response_no_nonce"`
    Ocsp_response_nonce_error int `json:"ocsp_response_nonce_error"`
    Crl_request int `json:"crl_request"`
    Crl_response int `json:"crl_response"`
    Crl_connection_error int `json:"crl_connection_error"`
    Crl_uri_not_found int `json:"crl_uri_not_found"`
    Crl_uri_https int `json:"crl_uri_https"`
    Crl_uri_unsupported int `json:"crl_uri_unsupported"`
    Crl_response_status_good int `json:"crl_response_status_good"`
    Crl_response_status_revoked int `json:"crl_response_status_revoked"`
    Crl_response_status_unknown int `json:"crl_response_status_unknown"`
    Crl_cache_status_good int `json:"crl_cache_status_good"`
    Crl_cache_status_revoked int `json:"crl_cache_status_revoked"`
    Crl_other_error int `json:"crl_other_error"`
}

func (p *SlbSslCertRevokeStats) GetId() string{
    return "1"
}

func (p *SlbSslCertRevokeStats) getPath() string{
    return "slb/ssl-cert-revoke/stats"
}

func (p *SlbSslCertRevokeStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCertRevokeStats,error) {
logger.Println("SlbSslCertRevokeStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCertRevokeStats
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
