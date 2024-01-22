

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdDiskOper struct {
    
    Oper RrdDiskOperOper `json:"oper"`

}
type DataRrdDiskOper struct {
    DtRrdDiskOper RrdDiskOper `json:"disk"`
}


type RrdDiskOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    TotalDisk string `json:"total-disk"`
    DiskUsage []RrdDiskOperOperDiskUsage `json:"disk-usage"`
    Alldynamic int `json:"alldynamic"`
}


type RrdDiskOperOperDiskUsage struct {
    Time int `json:"time"`
    DiskUsage string `json:"disk-usage"`
}

func (p *RrdDiskOper) GetId() string{
    return "1"
}

func (p *RrdDiskOper) getPath() string{
    return "rrd/disk/oper"
}

func (p *RrdDiskOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdDiskOper,error) {
logger.Println("RrdDiskOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdDiskOper
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
