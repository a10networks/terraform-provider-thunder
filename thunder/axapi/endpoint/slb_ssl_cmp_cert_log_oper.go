

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslCmpCertLogOper struct {
    
    Oper SlbSslCmpCertLogOperOper `json:"oper"`

}
type DataSlbSslCmpCertLogOper struct {
    DtSlbSslCmpCertLogOper SlbSslCmpCertLogOper `json:"ssl-cmp-cert-log"`
}


type SlbSslCmpCertLogOperOper struct {
    CmpLogList []SlbSslCmpCertLogOperOperCmpLogList `json:"cmp-log-list"`
    CmpLogOffset int `json:"cmp-log-offset"`
    CmpLogOver int `json:"cmp-log-over"`
    Name string `json:"name"`
    Follow int `json:"follow"`
    FromStart int `json:"from-start"`
    NumLines int `json:"num-lines"`
}


type SlbSslCmpCertLogOperOperCmpLogList struct {
    CmpLogData string `json:"cmp-log-data"`
}

func (p *SlbSslCmpCertLogOper) GetId() string{
    return "1"
}

func (p *SlbSslCmpCertLogOper) getPath() string{
    return "slb/ssl-cmp-cert-log/oper"
}

func (p *SlbSslCmpCertLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslCmpCertLogOper,error) {
logger.Println("SlbSslCmpCertLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslCmpCertLogOper
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
