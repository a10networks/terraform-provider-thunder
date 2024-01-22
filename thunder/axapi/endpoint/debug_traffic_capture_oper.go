

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugTrafficCaptureOper struct {
    
    Oper DebugTrafficCaptureOperOper `json:"oper"`

}
type DataDebugTrafficCaptureOper struct {
    DtDebugTrafficCaptureOper DebugTrafficCaptureOper `json:"debug-traffic-capture"`
}


type DebugTrafficCaptureOperOper struct {
    StatusInfo []DebugTrafficCaptureOperOperStatusInfo `json:"status-info"`
}


type DebugTrafficCaptureOperOperStatusInfo struct {
    Name string `json:"name"`
    Status string `json:"status"`
    End_time string `json:"end_time"`
    End_reason string `json:"end_reason"`
    Pkt_count int `json:"pkt_count"`
    Pkt_dropped int `json:"pkt_dropped"`
    Disk_size int `json:"disk_size"`
}

func (p *DebugTrafficCaptureOper) GetId() string{
    return "1"
}

func (p *DebugTrafficCaptureOper) getPath() string{
    return "debug-traffic-capture/oper"
}

func (p *DebugTrafficCaptureOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDebugTrafficCaptureOper,error) {
logger.Println("DebugTrafficCaptureOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDebugTrafficCaptureOper
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
