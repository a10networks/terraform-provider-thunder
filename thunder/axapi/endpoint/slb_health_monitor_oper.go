

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthMonitorOper struct {
    
    Oper SlbHealthMonitorOperOper `json:"oper"`

}
type DataSlbHealthMonitorOper struct {
    DtSlbHealthMonitorOper SlbHealthMonitorOper `json:"health-monitor"`
}


type SlbHealthMonitorOperOper struct {
    HealthMonitorList []SlbHealthMonitorOperOperHealthMonitorList `json:"health-monitor-list"`
}


type SlbHealthMonitorOperOperHealthMonitorList struct {
    Name string `json:"name"`
    Interval int `json:"interval"`
    Retries int `json:"retries"`
    Timeout int `json:"timeout"`
    UpRetries int `json:"up-retries"`
    Method string `json:"method"`
    Status string `json:"status"`
    Partition string `json:"partition"`
    SslRefresh int `json:"ssl-refresh"`
    PinId int `json:"pin-id"`
    PinProcessIndex int `json:"pin-process-index"`
    AllPartitions int `json:"all-partitions"`
}

func (p *SlbHealthMonitorOper) GetId() string{
    return "1"
}

func (p *SlbHealthMonitorOper) getPath() string{
    return "slb/health-monitor/oper"
}

func (p *SlbHealthMonitorOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthMonitorOper,error) {
logger.Println("SlbHealthMonitorOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthMonitorOper
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
