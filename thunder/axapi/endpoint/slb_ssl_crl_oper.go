

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCrlOper struct {
    
    Oper SlbSslCrlOperOper `json:"oper"`

}
type DataSlbSslCrlOper struct {
    DtSlbSslCrlOper SlbSslCrlOper `json:"ssl-crl"`
}


type SlbSslCrlOperOper struct {
    Vserver string `json:"vserver"`
    Port int `json:"port"`
    CrlInfo []SlbSslCrlOperOperCrlInfo `json:"crl-info"`
}


type SlbSslCrlOperOperCrlInfo struct {
    Issuer string `json:"issuer"`
    Status string `json:"status"`
}

func (p *SlbSslCrlOper) GetId() string{
    return "1"
}

func (p *SlbSslCrlOper) getPath() string{
    return "slb/ssl-crl/oper"
}

func (p *SlbSslCrlOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCrlOper,error) {
logger.Println("SlbSslCrlOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCrlOper
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
