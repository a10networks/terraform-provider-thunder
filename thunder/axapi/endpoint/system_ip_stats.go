

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpStats struct {
    
    SamplingEnable []SystemIpStatsSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

}
type DataSystemIpStats struct {
    DtSystemIpStats SystemIpStats `json:"ip-stats"`
}


type SystemIpStatsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemIpStats) GetId() string{
    return "1"
}

func (p *SystemIpStats) getPath() string{
    return "system/ip-stats"
}

func (p *SystemIpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIpStats,error) {
logger.Println("SystemIpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIpStats
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
