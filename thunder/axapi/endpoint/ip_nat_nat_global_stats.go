

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatNatGlobalStats struct {
    
    Stats IpNatNatGlobalStatsStats `json:"stats"`

}
type DataIpNatNatGlobalStats struct {
    DtIpNatNatGlobalStats IpNatNatGlobalStats `json:"nat-global"`
}


type IpNatNatGlobalStatsStats struct {
}

func (p *IpNatNatGlobalStats) GetId() string{
    return "1"
}

func (p *IpNatNatGlobalStats) getPath() string{
    return "ip/nat/nat-global/stats"
}

func (p *IpNatNatGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpNatNatGlobalStats,error) {
logger.Println("IpNatNatGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpNatNatGlobalStats
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
