

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAuthenticationLogOper struct {
    
    Oper LoggingLocalLogAuthenticationLogOperOper `json:"oper"`

}
type DataLoggingLocalLogAuthenticationLogOper struct {
    DtLoggingLocalLogAuthenticationLogOper LoggingLocalLogAuthenticationLogOper `json:"log"`
}


type LoggingLocalLogAuthenticationLogOperOper struct {
    MaxEntries int `json:"max-entries"`
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAuthenticationLogOperOperLogList `json:"log-list"`
}


type LoggingLocalLogAuthenticationLogOperOperLogList struct {
    Time string `json:"time"`
    UserName string `json:"user-name"`
    UserDomain string `json:"user-domain"`
    ClientIp string `json:"client-ip"`
    DestinationIp string `json:"destination-ip"`
    SourcePort int `json:"source-port"`
    DestinationPort int `json:"destination-port"`
    AuthResult string `json:"auth-result"`
}

func (p *LoggingLocalLogAuthenticationLogOper) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAuthenticationLogOper) getPath() string{
    return "logging/local-log/authentication/log/oper"
}

func (p *LoggingLocalLogAuthenticationLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogAuthenticationLogOper,error) {
logger.Println("LoggingLocalLogAuthenticationLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogAuthenticationLogOper
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
