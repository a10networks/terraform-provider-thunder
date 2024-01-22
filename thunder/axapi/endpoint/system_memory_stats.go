

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemMemoryStats struct {
    
    Stats SystemMemoryStatsStats `json:"stats"`

}
type DataSystemMemoryStats struct {
    DtSystemMemoryStats SystemMemoryStats `json:"memory"`
}


type SystemMemoryStatsStats struct {
    UsagePercentage int `json:"usage-percentage"`
}

func (p *SystemMemoryStats) GetId() string{
    return "1"
}

func (p *SystemMemoryStats) getPath() string{
    return "system/memory/stats"
}

func (p *SystemMemoryStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemMemoryStats,error) {
logger.Println("SystemMemoryStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemMemoryStats
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
