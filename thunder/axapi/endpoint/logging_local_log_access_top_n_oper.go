

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAccessTopNOper struct {
    
    Oper LoggingLocalLogAccessTopNOperOper `json:"oper"`

}
type DataLoggingLocalLogAccessTopNOper struct {
    DtLoggingLocalLogAccessTopNOper LoggingLocalLogAccessTopNOper `json:"top-n"`
}


type LoggingLocalLogAccessTopNOperOper struct {
    MaxEntries int `json:"max-entries"`
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    Top string `json:"top"`
    Action string `json:"action"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAccessTopNOperOperLogList `json:"log-list"`
}


type LoggingLocalLogAccessTopNOperOperLogList struct {
    Name string `json:"name"`
    Counter int `json:"counter"`
}

func (p *LoggingLocalLogAccessTopNOper) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAccessTopNOper) getPath() string{
    return "logging/local-log/access/top-n/oper"
}

func (p *LoggingLocalLogAccessTopNOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogAccessTopNOper,error) {
logger.Println("LoggingLocalLogAccessTopNOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogAccessTopNOper
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
