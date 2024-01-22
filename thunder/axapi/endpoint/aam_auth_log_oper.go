

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthLogOper struct {
    
    Oper AamAuthLogOperOper `json:"oper"`

}
type DataAamAuthLogOper struct {
    DtAamAuthLogOper AamAuthLogOper `json:"auth-log"`
}


type AamAuthLogOperOper struct {
    Tail int `json:"tail"`
    Last int `json:"last"`
    Top int `json:"top"`
    Target string `json:"target"`
    AuthStatus string `json:"auth-status"`
    Total int `json:"total"`
    LogList []AamAuthLogOperOperLogList `json:"log-list"`
}


type AamAuthLogOperOperLogList struct {
    RecordId int `json:"record-id"`
    Time string `json:"time"`
    UserName string `json:"user-name"`
    UserDomain string `json:"user-domain"`
    ClientIp string `json:"client-ip"`
    DestinationIp string `json:"destination-ip"`
    SourcePort int `json:"source-port"`
    DestinationPort int `json:"destination-port"`
    AuthResult string `json:"auth-result"`
    Counter int `json:"counter"`
}

func (p *AamAuthLogOper) GetId() string{
    return "1"
}

func (p *AamAuthLogOper) getPath() string{
    return "aam/auth-log/oper"
}

func (p *AamAuthLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthLogOper,error) {
logger.Println("AamAuthLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthLogOper
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
