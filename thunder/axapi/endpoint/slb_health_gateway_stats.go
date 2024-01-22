

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthGatewayStats struct {
    
    Stats SlbHealthGatewayStatsStats `json:"stats"`

}
type DataSlbHealthGatewayStats struct {
    DtSlbHealthGatewayStats SlbHealthGatewayStats `json:"health-gateway"`
}


type SlbHealthGatewayStatsStats struct {
    Total_sent int `json:"total_sent"`
    Total_retry_sent int `json:"total_retry_sent"`
    Total_timeout int `json:"total_timeout"`
}

func (p *SlbHealthGatewayStats) GetId() string{
    return "1"
}

func (p *SlbHealthGatewayStats) getPath() string{
    return "slb/health-gateway/stats"
}

func (p *SlbHealthGatewayStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthGatewayStats,error) {
logger.Println("SlbHealthGatewayStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthGatewayStats
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
