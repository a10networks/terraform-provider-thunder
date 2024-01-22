

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslAcmeCertStatusOper struct {
    
    Oper SlbSslAcmeCertStatusOperOper `json:"oper"`

}
type DataSlbSslAcmeCertStatusOper struct {
    DtSlbSslAcmeCertStatusOper SlbSslAcmeCertStatusOper `json:"ssl-acme-cert-status"`
}


type SlbSslAcmeCertStatusOperOper struct {
    AcmeCerts []SlbSslAcmeCertStatusOperOperAcmeCerts `json:"acme-certs"`
}


type SlbSslAcmeCertStatusOperOperAcmeCerts struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Renew string `json:"renew"`
    Rotated int `json:"rotated"`
    LastEnrollRenewStatus string `json:"last-enroll-renew-status"`
}

func (p *SlbSslAcmeCertStatusOper) GetId() string{
    return "1"
}

func (p *SlbSslAcmeCertStatusOper) getPath() string{
    return "slb/ssl-acme-cert-status/oper"
}

func (p *SlbSslAcmeCertStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslAcmeCertStatusOper,error) {
logger.Println("SlbSslAcmeCertStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslAcmeCertStatusOper
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
