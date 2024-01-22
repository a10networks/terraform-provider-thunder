

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCertRevokeOper struct {
    
    Oper SlbSslCertRevokeOperOper `json:"oper"`

}
type DataSlbSslCertRevokeOper struct {
    DtSlbSslCertRevokeOper SlbSslCertRevokeOper `json:"ssl-cert-revoke"`
}


type SlbSslCertRevokeOperOper struct {
    Vserver string `json:"vserver"`
    Port int `json:"port"`
}

func (p *SlbSslCertRevokeOper) GetId() string{
    return "1"
}

func (p *SlbSslCertRevokeOper) getPath() string{
    return "slb/ssl-cert-revoke/oper"
}

func (p *SlbSslCertRevokeOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCertRevokeOper,error) {
logger.Println("SlbSslCertRevokeOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCertRevokeOper
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
