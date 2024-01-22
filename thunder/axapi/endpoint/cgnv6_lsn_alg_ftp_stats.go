

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgFtpStats struct {
    
    Stats Cgnv6LsnAlgFtpStatsStats `json:"stats"`

}
type DataCgnv6LsnAlgFtpStats struct {
    DtCgnv6LsnAlgFtpStats Cgnv6LsnAlgFtpStats `json:"ftp"`
}


type Cgnv6LsnAlgFtpStatsStats struct {
    PortRequests int `json:"port-requests"`
    EprtRequests int `json:"eprt-requests"`
    LprtRequests int `json:"lprt-requests"`
    PasvReplies int `json:"pasv-replies"`
    EpsvReplies int `json:"epsv-replies"`
    LpsvReplies int `json:"lpsv-replies"`
}

func (p *Cgnv6LsnAlgFtpStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgFtpStats) getPath() string{
    return "cgnv6/lsn/alg/ftp/stats"
}

func (p *Cgnv6LsnAlgFtpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAlgFtpStats,error) {
logger.Println("Cgnv6LsnAlgFtpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAlgFtpStats
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
