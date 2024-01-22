

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HarmonyControllerTunnelStatsOper struct {
    
    Oper HarmonyControllerTunnelStatsOperOper `json:"oper"`

}
type DataHarmonyControllerTunnelStatsOper struct {
    DtHarmonyControllerTunnelStatsOper HarmonyControllerTunnelStatsOper `json:"tunnel-stats"`
}


type HarmonyControllerTunnelStatsOperOper struct {
    Status string `json:"status"`
    BytesSent int `json:"bytes-sent"`
    BytesRecieved int `json:"bytes-recieved"`
    NumberOfErrors int `json:"number-of-errors"`
    Uptime string `json:"uptime"`
    ErrorMessage string `json:"error-message"`
}

func (p *HarmonyControllerTunnelStatsOper) GetId() string{
    return "1"
}

func (p *HarmonyControllerTunnelStatsOper) getPath() string{
    return "harmony-controller/tunnel-stats/oper"
}

func (p *HarmonyControllerTunnelStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataHarmonyControllerTunnelStatsOper,error) {
logger.Println("HarmonyControllerTunnelStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataHarmonyControllerTunnelStatsOper
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
