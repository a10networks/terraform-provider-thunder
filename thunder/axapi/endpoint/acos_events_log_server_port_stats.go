

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsLogServerPortStats struct {
    
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats AcosEventsLogServerPortStatsStats `json:"stats"`
    Name string 

}
type DataAcosEventsLogServerPortStats struct {
    DtAcosEventsLogServerPortStats AcosEventsLogServerPortStats `json:"port"`
}


type AcosEventsLogServerPortStatsStats struct {
}

func (p *AcosEventsLogServerPortStats) GetId() string{
    return "1"
}

func (p *AcosEventsLogServerPortStats) getPath() string{
    
    return "acos-events/log-server/" +p.Name + "/port/" +strconv.Itoa(p.PortNumber)+"+"+p.Protocol+"/stats"
}

func (p *AcosEventsLogServerPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAcosEventsLogServerPortStats,error) {
logger.Println("AcosEventsLogServerPortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAcosEventsLogServerPortStats
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
