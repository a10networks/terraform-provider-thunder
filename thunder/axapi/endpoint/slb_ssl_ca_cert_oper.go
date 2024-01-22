

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCaCertOper struct {
    
    Oper SlbSslCaCertOperOper `json:"oper"`

}
type DataSlbSslCaCertOper struct {
    DtSlbSslCaCertOper SlbSslCaCertOper `json:"ssl-ca-cert"`
}


type SlbSslCaCertOperOper struct {
    SortbyName int `json:"sortby-name"`
    SortbyExp int `json:"sortby-exp"`
    SslCerts []SlbSslCaCertOperOperSslCerts `json:"ssl-certs"`
    Partition string `json:"partition"`
    ExactMatch int `json:"exact-match"`
}


type SlbSslCaCertOperOperSslCerts struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Serial string `json:"serial"`
    Notbefore string `json:"notbefore"`
    Notafter string `json:"notafter"`
    CommonName string `json:"common-name"`
    Organization string `json:"organization"`
    Subject string `json:"subject"`
    Issuer string `json:"issuer"`
    NotafterNumber int `json:"notafter-number"`
    Status string `json:"status"`
    Keysize int `json:"keysize"`
}

func (p *SlbSslCaCertOper) GetId() string{
    return "1"
}

func (p *SlbSslCaCertOper) getPath() string{
    return "slb/ssl-ca-cert/oper"
}

func (p *SlbSslCaCertOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCaCertOper,error) {
logger.Println("SlbSslCaCertOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCaCertOper
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
