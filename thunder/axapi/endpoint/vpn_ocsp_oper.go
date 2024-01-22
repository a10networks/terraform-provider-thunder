

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnOcspOper struct {
    
    Oper VpnOcspOperOper `json:"oper"`

}
type DataVpnOcspOper struct {
    DtVpnOcspOper VpnOcspOper `json:"ocsp"`
}


type VpnOcspOperOper struct {
    OcspList []VpnOcspOperOperOcspList `json:"ocsp-list"`
    TotalOcsps int `json:"total-ocsps"`
}


type VpnOcspOperOperOcspList struct {
    Subject string `json:"subject"`
    Issuer string `json:"issuer"`
    Validity string `json:"validity"`
    CertificateStatus string `json:"certificate-status"`
}

func (p *VpnOcspOper) GetId() string{
    return "1"
}

func (p *VpnOcspOper) getPath() string{
    return "vpn/ocsp/oper"
}

func (p *VpnOcspOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnOcspOper,error) {
logger.Println("VpnOcspOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnOcspOper
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
