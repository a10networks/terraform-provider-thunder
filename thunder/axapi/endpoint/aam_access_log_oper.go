

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAccessLogOper struct {
    
    Oper AamAccessLogOperOper `json:"oper"`

}
type DataAamAccessLogOper struct {
    DtAamAccessLogOper AamAccessLogOper `json:"access-log"`
}


type AamAccessLogOperOper struct {
    Tail int `json:"tail"`
    Last int `json:"last"`
    Top int `json:"top"`
    Target string `json:"target"`
    AccessStatus string `json:"access-status"`
    HostName string `json:"host-name"`
    Total int `json:"total"`
    LogList []AamAccessLogOperOperLogList `json:"log-list"`
}


type AamAccessLogOperOperLogList struct {
    RecordId int `json:"record-id"`
    Time string `json:"time"`
    UserName string `json:"user-name"`
    UserDomain string `json:"user-domain"`
    ClientIp string `json:"client-ip"`
    DestinationIp string `json:"destination-ip"`
    SourcePort int `json:"source-port"`
    DestinationPort int `json:"destination-port"`
    Policy string `json:"policy"`
    Action string `json:"action"`
    VipName string `json:"vip-name"`
    VipPort int `json:"vip-port"`
    Host string `json:"host"`
    Uri string `json:"uri"`
    WebCategory string `json:"web-category"`
    SslStatus string `json:"ssl-status"`
    Counter int `json:"counter"`
}

func (p *AamAccessLogOper) GetId() string{
    return "1"
}

func (p *AamAccessLogOper) getPath() string{
    return "aam/access-log/oper"
}

func (p *AamAccessLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAccessLogOper,error) {
logger.Println("AamAccessLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAccessLogOper
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
