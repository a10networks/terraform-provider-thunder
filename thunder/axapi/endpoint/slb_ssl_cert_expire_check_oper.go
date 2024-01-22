

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCertExpireCheckOper struct {
    
    Oper SlbSslCertExpireCheckOperOper `json:"oper"`

}
type DataSlbSslCertExpireCheckOper struct {
    DtSlbSslCertExpireCheckOper SlbSslCertExpireCheckOper `json:"ssl-cert-expire-check"`
}


type SlbSslCertExpireCheckOperOper struct {
    ExpireCheckStatus string `json:"expire-check-status"`
    EmailAddress string `json:"email-address"`
    EmailAddress2 string `json:"email-address2"`
    Before int `json:"before"`
    Interval int `json:"interval"`
    SslException []SlbSslCertExpireCheckOperOperSslException `json:"ssl-exception"`
}


type SlbSslCertExpireCheckOperOperSslException struct {
    ExceptionCert string `json:"exception-cert"`
}

func (p *SlbSslCertExpireCheckOper) GetId() string{
    return "1"
}

func (p *SlbSslCertExpireCheckOper) getPath() string{
    return "slb/ssl-cert-expire-check/oper"
}

func (p *SlbSslCertExpireCheckOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCertExpireCheckOper,error) {
logger.Println("SlbSslCertExpireCheckOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCertExpireCheckOper
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
