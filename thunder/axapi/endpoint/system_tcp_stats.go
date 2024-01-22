

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTcpStats struct {
    
    SamplingEnable []SystemTcpStatsSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

}
type DataSystemTcpStats struct {
    DtSystemTcpStats SystemTcpStats `json:"tcp-stats"`
}


type SystemTcpStatsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemTcpStats) GetId() string{
    return "1"
}

func (p *SystemTcpStats) getPath() string{
    return "system/tcp-stats"
}

func (p *SystemTcpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTcpStats,error) {
logger.Println("SystemTcpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTcpStats
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
