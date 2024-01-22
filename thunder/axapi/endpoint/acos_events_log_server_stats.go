

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsLogServerStats struct {
    
    Name string `json:"name"`
    PortList []AcosEventsLogServerStatsPortList `json:"port-list"`
    Stats AcosEventsLogServerStatsStats `json:"stats"`

}
type DataAcosEventsLogServerStats struct {
    DtAcosEventsLogServerStats AcosEventsLogServerStats `json:"log-server"`
}


type AcosEventsLogServerStatsPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats AcosEventsLogServerStatsPortListStats `json:"stats"`
}


type AcosEventsLogServerStatsPortListStats struct {
}


type AcosEventsLogServerStatsStats struct {
}

func (p *AcosEventsLogServerStats) GetId() string{
    return "1"
}

func (p *AcosEventsLogServerStats) getPath() string{
    
    return "acos-events/log-server/"+p.Name+"/stats"
}

func (p *AcosEventsLogServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAcosEventsLogServerStats,error) {
logger.Println("AcosEventsLogServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAcosEventsLogServerStats
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
