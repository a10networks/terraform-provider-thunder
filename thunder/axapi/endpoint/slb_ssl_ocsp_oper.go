

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslOcspOper struct {
    
    Oper SlbSslOcspOperOper `json:"oper"`

}
type DataSlbSslOcspOper struct {
    DtSlbSslOcspOper SlbSslOcspOper `json:"ssl-ocsp"`
}


type SlbSslOcspOperOper struct {
    TotalEntries int `json:"total-entries"`
    CachedEntries []SlbSslOcspOperOperCachedEntries `json:"cached-entries"`
}


type SlbSslOcspOperOperCachedEntries struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Subject string `json:"subject"`
    Length int `json:"length"`
    Uri string `json:"uri"`
    Expire int `json:"expire"`
    Hits int `json:"hits"`
}

func (p *SlbSslOcspOper) GetId() string{
    return "1"
}

func (p *SlbSslOcspOper) getPath() string{
    return "slb/ssl-ocsp/oper"
}

func (p *SlbSslOcspOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslOcspOper,error) {
logger.Println("SlbSslOcspOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslOcspOper
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
