

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIp6Stats struct {
    
    SamplingEnable []SystemIp6StatsSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

}
type DataSystemIp6Stats struct {
    DtSystemIp6Stats SystemIp6Stats `json:"ip6-stats"`
}


type SystemIp6StatsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemIp6Stats) GetId() string{
    return "1"
}

func (p *SystemIp6Stats) getPath() string{
    return "system/ip6-stats"
}

func (p *SystemIp6Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemIp6Stats,error) {
logger.Println("SystemIp6Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemIp6Stats
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
