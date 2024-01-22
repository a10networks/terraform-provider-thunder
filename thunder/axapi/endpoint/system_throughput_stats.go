

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemThroughputStats struct {
    
    Stats SystemThroughputStatsStats `json:"stats"`

}
type DataSystemThroughputStats struct {
    DtSystemThroughputStats SystemThroughputStats `json:"throughput"`
}


type SystemThroughputStatsStats struct {
    GlobalSystemThroughputBitsPerSec int `json:"global-system-throughput-bits-per-sec"`
    PerPartThroughputBitsPerSec int `json:"per-part-throughput-bits-per-sec"`
}

func (p *SystemThroughputStats) GetId() string{
    return "1"
}

func (p *SystemThroughputStats) getPath() string{
    return "system/throughput/stats"
}

func (p *SystemThroughputStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemThroughputStats,error) {
logger.Println("SystemThroughputStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemThroughputStats
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
