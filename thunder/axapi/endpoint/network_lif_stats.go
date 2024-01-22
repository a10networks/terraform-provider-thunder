

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkLifStats struct {
    
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`

}
type DataNetworkLifStats struct {
    DtNetworkLifStats NetworkLifStats `json:"lif-stats"`
}

func (p *NetworkLifStats) GetId() string{
    return "1"
}

func (p *NetworkLifStats) getPath() string{
    return "network/lif-stats"
}

func (p *NetworkLifStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkLifStats,error) {
logger.Println("NetworkLifStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkLifStats
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
