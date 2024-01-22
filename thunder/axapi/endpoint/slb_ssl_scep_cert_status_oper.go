

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslScepCertStatusOper struct {
    
    Oper SlbSslScepCertStatusOperOper `json:"oper"`

}
type DataSlbSslScepCertStatusOper struct {
    DtSlbSslScepCertStatusOper SlbSslScepCertStatusOper `json:"ssl-scep-cert-status"`
}


type SlbSslScepCertStatusOperOper struct {
    ScepCerts []SlbSslScepCertStatusOperOperScepCerts `json:"scep-certs"`
}


type SlbSslScepCertStatusOperOperScepCerts struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Renew string `json:"renew"`
    Rotated int `json:"rotated"`
}

func (p *SlbSslScepCertStatusOper) GetId() string{
    return "1"
}

func (p *SlbSslScepCertStatusOper) getPath() string{
    return "slb/ssl-scep-cert-status/oper"
}

func (p *SlbSslScepCertStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslScepCertStatusOper,error) {
logger.Println("SlbSslScepCertStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslScepCertStatusOper
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
