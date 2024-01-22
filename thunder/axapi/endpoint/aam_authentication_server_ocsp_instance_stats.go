

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerOcspInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationServerOcspInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationServerOcspInstanceStats struct {
    DtAamAuthenticationServerOcspInstanceStats AamAuthenticationServerOcspInstanceStats `json:"instance"`
}


type AamAuthenticationServerOcspInstanceStatsStats struct {
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

func (p *AamAuthenticationServerOcspInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerOcspInstanceStats) getPath() string{
    
    return "aam/authentication/server/ocsp/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationServerOcspInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerOcspInstanceStats,error) {
logger.Println("AamAuthenticationServerOcspInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerOcspInstanceStats
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
