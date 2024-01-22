

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbServerGroupStats struct {
    
    Name string `json:"name"`
    Stats SlbServerGroupStatsStats `json:"stats"`

}
type DataSlbServerGroupStats struct {
    DtSlbServerGroupStats SlbServerGroupStats `json:"server-group"`
}


type SlbServerGroupStatsStats struct {
    DummyConn int `json:"dummy-conn"`
}

func (p *SlbServerGroupStats) GetId() string{
    return "1"
}

func (p *SlbServerGroupStats) getPath() string{
    
    return "slb/server-group/"+p.Name+"/stats"
}

func (p *SlbServerGroupStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbServerGroupStats,error) {
logger.Println("SlbServerGroupStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbServerGroupStats
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
