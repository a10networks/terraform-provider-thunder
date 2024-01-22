

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslAcmeCertLogOper struct {
    
    Oper SlbSslAcmeCertLogOperOper `json:"oper"`

}
type DataSlbSslAcmeCertLogOper struct {
    DtSlbSslAcmeCertLogOper SlbSslAcmeCertLogOper `json:"ssl-acme-cert-log"`
}


type SlbSslAcmeCertLogOperOper struct {
    AcmeLogList []SlbSslAcmeCertLogOperOperAcmeLogList `json:"acme-log-list"`
    AcmeLogOffset int `json:"acme-log-offset"`
    AcmeLogOver int `json:"acme-log-over"`
    Name string `json:"name"`
    Follow int `json:"follow"`
    FromStart int `json:"from-start"`
    NumLines int `json:"num-lines"`
}


type SlbSslAcmeCertLogOperOperAcmeLogList struct {
    AcmeLogData string `json:"acme-log-data"`
}

func (p *SlbSslAcmeCertLogOper) GetId() string{
    return "1"
}

func (p *SlbSslAcmeCertLogOper) getPath() string{
    return "slb/ssl-acme-cert-log/oper"
}

func (p *SlbSslAcmeCertLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslAcmeCertLogOper,error) {
logger.Println("SlbSslAcmeCertLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslAcmeCertLogOper
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
