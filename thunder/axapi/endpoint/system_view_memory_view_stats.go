

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewMemoryViewStats struct {
    
    Stats SystemViewMemoryViewStatsStats `json:"stats"`

}
type DataSystemViewMemoryViewStats struct {
    DtSystemViewMemoryViewStats SystemViewMemoryViewStats `json:"memory-view"`
}


type SystemViewMemoryViewStatsStats struct {
    UsagePercentage int `json:"usage-percentage"`
}

func (p *SystemViewMemoryViewStats) GetId() string{
    return "1"
}

func (p *SystemViewMemoryViewStats) getPath() string{
    return "system-view/memory-view/stats"
}

func (p *SystemViewMemoryViewStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemViewMemoryViewStats,error) {
logger.Println("SystemViewMemoryViewStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemViewMemoryViewStats
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
