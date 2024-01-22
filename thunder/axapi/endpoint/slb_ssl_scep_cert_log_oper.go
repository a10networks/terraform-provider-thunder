

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslScepCertLogOper struct {
    
    Oper SlbSslScepCertLogOperOper `json:"oper"`

}
type DataSlbSslScepCertLogOper struct {
    DtSlbSslScepCertLogOper SlbSslScepCertLogOper `json:"ssl-scep-cert-log"`
}


type SlbSslScepCertLogOperOper struct {
    ScepLogList []SlbSslScepCertLogOperOperScepLogList `json:"scep-log-list"`
    ScepLogOffset int `json:"scep-log-offset"`
    ScepLogOver int `json:"scep-log-over"`
    Name string `json:"name"`
    Follow int `json:"follow"`
    FromStart int `json:"from-start"`
    NumLines int `json:"num-lines"`
}


type SlbSslScepCertLogOperOperScepLogList struct {
    ScepLogData string `json:"scep-log-data"`
}

func (p *SlbSslScepCertLogOper) GetId() string{
    return "1"
}

func (p *SlbSslScepCertLogOper) getPath() string{
    return "slb/ssl-scep-cert-log/oper"
}

func (p *SlbSslScepCertLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslScepCertLogOper,error) {
logger.Println("SlbSslScepCertLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslScepCertLogOper
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
