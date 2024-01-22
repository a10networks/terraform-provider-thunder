

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemDpdkStats struct {
    
    SamplingEnable []SystemDpdkStatsSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

}
type DataSystemDpdkStats struct {
    DtSystemDpdkStats SystemDpdkStats `json:"dpdk-stats"`
}


type SystemDpdkStatsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemDpdkStats) GetId() string{
    return "1"
}

func (p *SystemDpdkStats) getPath() string{
    return "system/dpdk-stats"
}

func (p *SystemDpdkStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemDpdkStats,error) {
logger.Println("SystemDpdkStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemDpdkStats
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
