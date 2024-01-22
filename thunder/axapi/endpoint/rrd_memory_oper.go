

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdMemoryOper struct {
    
    Oper RrdMemoryOperOper `json:"oper"`

}
type DataRrdMemoryOper struct {
    DtRrdMemoryOper RrdMemoryOper `json:"memory"`
}


type RrdMemoryOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    MemUsage []RrdMemoryOperOperMemUsage `json:"mem-usage"`
}


type RrdMemoryOperOperMemUsage struct {
    Time int `json:"time"`
    MemUsage string `json:"mem-usage"`
}

func (p *RrdMemoryOper) GetId() string{
    return "1"
}

func (p *RrdMemoryOper) getPath() string{
    return "rrd/memory/oper"
}

func (p *RrdMemoryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdMemoryOper,error) {
logger.Println("RrdMemoryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdMemoryOper
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
