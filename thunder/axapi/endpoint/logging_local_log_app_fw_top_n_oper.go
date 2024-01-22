

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAppFwTopNOper struct {
    
    Oper LoggingLocalLogAppFwTopNOperOper `json:"oper"`

}
type DataLoggingLocalLogAppFwTopNOper struct {
    DtLoggingLocalLogAppFwTopNOper LoggingLocalLogAppFwTopNOper `json:"top-n"`
}


type LoggingLocalLogAppFwTopNOperOper struct {
    MaxEntries int `json:"max-entries"`
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    Top string `json:"top"`
    ApplicationName string `json:"application-name"`
    Category string `json:"category"`
    ClientIp string `json:"client-ip"`
    ThreatCategory string `json:"threat-category"`
    ThreatCategoryMatch string `json:"threat-category-match"`
    Action string `json:"action"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAppFwTopNOperOperLogList `json:"log-list"`
}


type LoggingLocalLogAppFwTopNOperOperLogList struct {
    Name string `json:"name"`
    Counter int `json:"counter"`
}

func (p *LoggingLocalLogAppFwTopNOper) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAppFwTopNOper) getPath() string{
    return "logging/local-log/app-fw/top-n/oper"
}

func (p *LoggingLocalLogAppFwTopNOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogAppFwTopNOper,error) {
logger.Println("LoggingLocalLogAppFwTopNOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogAppFwTopNOper
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
