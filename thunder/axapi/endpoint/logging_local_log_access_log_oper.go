

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAccessLogOper struct {
    
    Oper LoggingLocalLogAccessLogOperOper `json:"oper"`

}
type DataLoggingLocalLogAccessLogOper struct {
    DtLoggingLocalLogAccessLogOper LoggingLocalLogAccessLogOper `json:"log"`
}


type LoggingLocalLogAccessLogOperOper struct {
    MaxEntries int `json:"max-entries"`
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAccessLogOperOperLogList `json:"log-list"`
}


type LoggingLocalLogAccessLogOperOperLogList struct {
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
    WebDomain string `json:"web-domain"`
    Uri string `json:"uri"`
    WebCategory string `json:"web-category"`
    WebReputation string `json:"web-reputation"`
    SslStatus string `json:"ssl-status"`
    Reason string `json:"reason"`
}

func (p *LoggingLocalLogAccessLogOper) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAccessLogOper) getPath() string{
    return "logging/local-log/access/log/oper"
}

func (p *LoggingLocalLogAccessLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogAccessLogOper,error) {
logger.Println("LoggingLocalLogAccessLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogAccessLogOper
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
