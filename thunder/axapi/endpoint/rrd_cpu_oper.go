

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdCpuOper struct {
    
    Oper RrdCpuOperOper `json:"oper"`

}
type DataRrdCpuOper struct {
    DtRrdCpuOper RrdCpuOper `json:"cpu"`
}


type RrdCpuOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    CpuUsage []RrdCpuOperOperCpuUsage `json:"cpu-usage"`
}


type RrdCpuOperOperCpuUsage struct {
    Time int `json:"time"`
    CpuId int `json:"cpu-id"`
    CpuType int `json:"cpu-type"`
    CpuUsage string `json:"cpu-usage"`
}

func (p *RrdCpuOper) GetId() string{
    return "1"
}

func (p *RrdCpuOper) getPath() string{
    return "rrd/cpu/oper"
}

func (p *RrdCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdCpuOper,error) {
logger.Println("RrdCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdCpuOper
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
