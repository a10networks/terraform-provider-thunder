

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgTftpStats struct {
    
    Stats Cgnv6LsnAlgTftpStatsStats `json:"stats"`

}
type DataCgnv6LsnAlgTftpStats struct {
    DtCgnv6LsnAlgTftpStats Cgnv6LsnAlgTftpStats `json:"tftp"`
}


type Cgnv6LsnAlgTftpStatsStats struct {
    SessionCreated int `json:"session-created"`
}

func (p *Cgnv6LsnAlgTftpStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgTftpStats) getPath() string{
    return "cgnv6/lsn/alg/tftp/stats"
}

func (p *Cgnv6LsnAlgTftpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAlgTftpStats,error) {
logger.Println("Cgnv6LsnAlgTftpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAlgTftpStats
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
