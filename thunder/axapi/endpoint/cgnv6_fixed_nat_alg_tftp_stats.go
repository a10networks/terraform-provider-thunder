

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatAlgTftpStats struct {
    
    Stats Cgnv6FixedNatAlgTftpStatsStats `json:"stats"`

}
type DataCgnv6FixedNatAlgTftpStats struct {
    DtCgnv6FixedNatAlgTftpStats Cgnv6FixedNatAlgTftpStats `json:"tftp"`
}


type Cgnv6FixedNatAlgTftpStatsStats struct {
    SessionCreated int `json:"session-created"`
}

func (p *Cgnv6FixedNatAlgTftpStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatAlgTftpStats) getPath() string{
    return "cgnv6/fixed-nat/alg/tftp/stats"
}

func (p *Cgnv6FixedNatAlgTftpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatAlgTftpStats,error) {
logger.Println("Cgnv6FixedNatAlgTftpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatAlgTftpStats
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
