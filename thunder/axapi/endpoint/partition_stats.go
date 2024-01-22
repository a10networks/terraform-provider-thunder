

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PartitionStats struct {
    
    PartitionName string `json:"partition-name"`
    Stats PartitionStatsStats `json:"stats"`

}
type DataPartitionStats struct {
    DtPartitionStats PartitionStats `json:"partition"`
}


type PartitionStatsStats struct {
}

func (p *PartitionStats) GetId() string{
    return "1"
}

func (p *PartitionStats) getPath() string{
    
    return "partition/"+p.PartitionName+"/stats"
}

func (p *PartitionStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataPartitionStats,error) {
logger.Println("PartitionStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataPartitionStats
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
