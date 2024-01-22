

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAuthenticationTopNOper struct {
    
    Oper LoggingLocalLogAuthenticationTopNOperOper `json:"oper"`

}
type DataLoggingLocalLogAuthenticationTopNOper struct {
    DtLoggingLocalLogAuthenticationTopNOper LoggingLocalLogAuthenticationTopNOper `json:"top-n"`
}


type LoggingLocalLogAuthenticationTopNOperOper struct {
    MaxEntries int `json:"max-entries"`
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    Top string `json:"top"`
    AuthResult string `json:"auth-result"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAuthenticationTopNOperOperLogList `json:"log-list"`
}


type LoggingLocalLogAuthenticationTopNOperOperLogList struct {
    Name string `json:"name"`
    Counter int `json:"counter"`
}

func (p *LoggingLocalLogAuthenticationTopNOper) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAuthenticationTopNOper) getPath() string{
    return "logging/local-log/authentication/top-n/oper"
}

func (p *LoggingLocalLogAuthenticationTopNOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogAuthenticationTopNOper,error) {
logger.Println("LoggingLocalLogAuthenticationTopNOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogAuthenticationTopNOper
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
