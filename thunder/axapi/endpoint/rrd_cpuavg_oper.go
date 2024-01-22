

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdCpuavgOper struct {
    
    Oper RrdCpuavgOperOper `json:"oper"`

}
type DataRrdCpuavgOper struct {
    DtRrdCpuavgOper RrdCpuavgOper `json:"cpuavg"`
}


type RrdCpuavgOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    CpuUsage []RrdCpuavgOperOperCpuUsage `json:"cpu-usage"`
}


type RrdCpuavgOperOperCpuUsage struct {
    Time int `json:"time"`
    CpuNum int `json:"cpu-num"`
    CpuType int `json:"cpu-type"`
    CpuUsage string `json:"cpu-usage"`
}

func (p *RrdCpuavgOper) GetId() string{
    return "1"
}

func (p *RrdCpuavgOper) getPath() string{
    return "rrd/cpuavg/oper"
}

func (p *RrdCpuavgOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdCpuavgOper,error) {
logger.Println("RrdCpuavgOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdCpuavgOper
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
