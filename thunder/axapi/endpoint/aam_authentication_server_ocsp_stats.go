

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerOcspStats struct {
    
    InstanceList []AamAuthenticationServerOcspStatsInstanceList `json:"instance-list"`
    Stats AamAuthenticationServerOcspStatsStats `json:"stats"`

}
type DataAamAuthenticationServerOcspStats struct {
    DtAamAuthenticationServerOcspStats AamAuthenticationServerOcspStats `json:"ocsp"`
}


type AamAuthenticationServerOcspStatsInstanceList struct {
    Name string `json:"name"`
    Stats AamAuthenticationServerOcspStatsInstanceListStats `json:"stats"`
}


type AamAuthenticationServerOcspStatsInstanceListStats struct {
    Request int `json:"request"`
    CertificateGood int `json:"certificate-good"`
    CertificateRevoked int `json:"certificate-revoked"`
    CertificateUnknown int `json:"certificate-unknown"`
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingRequest int `json:"stapling-request"`
    StaplingCertificateGood int `json:"stapling-certificate-good"`
    StaplingCertificateRevoked int `json:"stapling-certificate-revoked"`
    StaplingCertificateUnknown int `json:"stapling-certificate-unknown"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
}


type AamAuthenticationServerOcspStatsStats struct {
    StaplingCertificateGood int `json:"stapling-certificate-good"`
    StaplingCertificateRevoked int `json:"stapling-certificate-revoked"`
    StaplingCertificateUnknown int `json:"stapling-certificate-unknown"`
    StaplingRequestNormal int `json:"stapling-request-normal"`
    StaplingRequestDropped int `json:"stapling-request-dropped"`
    StaplingResponseSuccess int `json:"stapling-response-success"`
    StaplingResponseFailure int `json:"stapling-response-failure"`
    StaplingResponseError int `json:"stapling-response-error"`
    StaplingResponseTimeout int `json:"stapling-response-timeout"`
    StaplingResponseOther int `json:"stapling-response-other"`
    RequestNormal int `json:"request-normal"`
    RequestDropped int `json:"request-dropped"`
    ResponseSuccess int `json:"response-success"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
}

func (p *AamAuthenticationServerOcspStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerOcspStats) getPath() string{
    return "aam/authentication/server/ocsp/stats"
}

func (p *AamAuthenticationServerOcspStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerOcspStats,error) {
logger.Println("AamAuthenticationServerOcspStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerOcspStats
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
