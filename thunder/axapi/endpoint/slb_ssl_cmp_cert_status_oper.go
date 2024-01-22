

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCmpCertStatusOper struct {
    
    Oper SlbSslCmpCertStatusOperOper `json:"oper"`

}
type DataSlbSslCmpCertStatusOper struct {
    DtSlbSslCmpCertStatusOper SlbSslCmpCertStatusOper `json:"ssl-cmp-cert-status"`
}


type SlbSslCmpCertStatusOperOper struct {
    CmpCerts []SlbSslCmpCertStatusOperOperCmpCerts `json:"cmp-certs"`
}


type SlbSslCmpCertStatusOperOperCmpCerts struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Renew string `json:"renew"`
    Rotated int `json:"rotated"`
}

func (p *SlbSslCmpCertStatusOper) GetId() string{
    return "1"
}

func (p *SlbSslCmpCertStatusOper) getPath() string{
    return "slb/ssl-cmp-cert-status/oper"
}

func (p *SlbSslCmpCertStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCmpCertStatusOper,error) {
logger.Println("SlbSslCmpCertStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCmpCertStatusOper
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
