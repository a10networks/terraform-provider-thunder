

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAppFwOper struct {
    
    DotPlot LoggingLocalLogAppFwOperDotPlot `json:"dot-plot"`
    Oper LoggingLocalLogAppFwOperOper `json:"oper"`
    TopN LoggingLocalLogAppFwOperTopN `json:"top-n"`

}
type DataLoggingLocalLogAppFwOper struct {
    DtLoggingLocalLogAppFwOper LoggingLocalLogAppFwOper `json:"app-fw"`
}


type LoggingLocalLogAppFwOperDotPlot struct {
    Oper LoggingLocalLogAppFwOperDotPlotOper `json:"oper"`
}


type LoggingLocalLogAppFwOperDotPlotOper struct {
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    ApplicationName string `json:"application-name"`
    ClientIp string `json:"client-ip"`
    Data string `json:"data"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAppFwOperDotPlotOperLogList `json:"log-list"`
}


type LoggingLocalLogAppFwOperDotPlotOperLogList struct {
    Counter int `json:"counter"`
}


type LoggingLocalLogAppFwOperOper struct {
    MaxEntries int `json:"max-entries"`
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    SourceThreatCategoryMatch string `json:"source-threat-category-match"`
    DestinationThreatCategoryMatch string `json:"destination-threat-category-match"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAppFwOperOperLogList `json:"log-list"`
}


type LoggingLocalLogAppFwOperOperLogList struct {
    Time string `json:"time"`
    ApplicationName string `json:"application-name"`
    Category string `json:"category"`
    ClientIp string `json:"client-ip"`
    DestinationIp string `json:"destination-ip"`
    SourcePort int `json:"source-port"`
    DestinationPort int `json:"destination-port"`
    SourceThreatListName string `json:"source-threat-list-name"`
    DestinationThreatListName string `json:"destination-threat-list-name"`
    SourceThreatCategory string `json:"source-threat-category"`
    DestinationThreatCategory string `json:"destination-threat-category"`
    Action string `json:"action"`
    PolicyName string `json:"policy-name"`
    RuleName string `json:"rule-name"`
    Bytes int `json:"bytes"`
}


type LoggingLocalLogAppFwOperTopN struct {
    Oper LoggingLocalLogAppFwOperTopNOper `json:"oper"`
}


type LoggingLocalLogAppFwOperTopNOper struct {
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
    LogList []LoggingLocalLogAppFwOperTopNOperLogList `json:"log-list"`
}


type LoggingLocalLogAppFwOperTopNOperLogList struct {
    Name string `json:"name"`
    Counter int `json:"counter"`
}

func (p *LoggingLocalLogAppFwOper) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAppFwOper) getPath() string{
    return "logging/local-log/app-fw/oper"
}

func (p *LoggingLocalLogAppFwOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogAppFwOper,error) {
logger.Println("LoggingLocalLogAppFwOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogAppFwOper
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
