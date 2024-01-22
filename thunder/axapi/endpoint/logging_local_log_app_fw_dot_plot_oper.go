

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAppFwDotPlotOper struct {
    
    Oper LoggingLocalLogAppFwDotPlotOperOper `json:"oper"`

}
type DataLoggingLocalLogAppFwDotPlotOper struct {
    DtLoggingLocalLogAppFwDotPlotOper LoggingLocalLogAppFwDotPlotOper `json:"dot-plot"`
}


type LoggingLocalLogAppFwDotPlotOperOper struct {
    StartTime string `json:"start-time"`
    Interval string `json:"interval"`
    IntervalPosition string `json:"interval-position"`
    ApplicationName string `json:"application-name"`
    ClientIp string `json:"client-ip"`
    Data string `json:"data"`
    Total int `json:"total"`
    LogList []LoggingLocalLogAppFwDotPlotOperOperLogList `json:"log-list"`
}


type LoggingLocalLogAppFwDotPlotOperOperLogList struct {
    Counter int `json:"counter"`
}

func (p *LoggingLocalLogAppFwDotPlotOper) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAppFwDotPlotOper) getPath() string{
    return "logging/local-log/app-fw/dot-plot/oper"
}

func (p *LoggingLocalLogAppFwDotPlotOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataLoggingLocalLogAppFwDotPlotOper,error) {
logger.Println("LoggingLocalLogAppFwDotPlotOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataLoggingLocalLogAppFwDotPlotOper
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
