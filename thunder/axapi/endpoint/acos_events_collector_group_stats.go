

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsCollectorGroupStats struct {
    
    Name string `json:"name"`
    Stats AcosEventsCollectorGroupStatsStats `json:"stats"`

}
type DataAcosEventsCollectorGroupStats struct {
    DtAcosEventsCollectorGroupStats AcosEventsCollectorGroupStats `json:"collector-group"`
}


type AcosEventsCollectorGroupStatsStats struct {
    Msgs_sent int `json:"msgs_sent"`
}

func (p *AcosEventsCollectorGroupStats) GetId() string{
    return "1"
}

func (p *AcosEventsCollectorGroupStats) getPath() string{
    
    return "acos-events/collector-group/"+p.Name+"/stats"
}

func (p *AcosEventsCollectorGroupStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAcosEventsCollectorGroupStats,error) {
logger.Println("AcosEventsCollectorGroupStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAcosEventsCollectorGroupStats
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
