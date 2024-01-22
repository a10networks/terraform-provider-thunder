

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosResourceTrackingCpuOper struct {
    
    Oper DdosResourceTrackingCpuOperOper `json:"oper"`

}
type DataDdosResourceTrackingCpuOper struct {
    DtDdosResourceTrackingCpuOper DdosResourceTrackingCpuOper `json:"cpu"`
}


type DdosResourceTrackingCpuOperOper struct {
    IfShowErrorStr int `json:"if-show-error-str"`
    ErrorStr string `json:"error-str"`
    Timestamps []DdosResourceTrackingCpuOperOperTimestamps `json:"timestamps"`
    MaxCount int `json:"max-count"`
}


type DdosResourceTrackingCpuOperOperTimestamps struct {
    Timestamp string `json:"timestamp"`
    AverageCpuPercent string `json:"average-cpu-percent"`
    Entries []DdosResourceTrackingCpuOperOperTimestampsEntries `json:"entries"`
}


type DdosResourceTrackingCpuOperOperTimestampsEntries struct {
    Entry string `json:"entry"`
    Address string `json:"address"`
    RelativeCpuPercent string `json:"relative-cpu-percent"`
    AbsoluteCpuPercent string `json:"absolute-cpu-percent"`
}

func (p *DdosResourceTrackingCpuOper) GetId() string{
    return "1"
}

func (p *DdosResourceTrackingCpuOper) getPath() string{
    return "ddos/resource-tracking/cpu/oper"
}

func (p *DdosResourceTrackingCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosResourceTrackingCpuOper,error) {
logger.Println("DdosResourceTrackingCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosResourceTrackingCpuOper
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
