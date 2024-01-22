

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatAlgFtpStats struct {
    
    Stats Cgnv6FixedNatAlgFtpStatsStats `json:"stats"`

}
type DataCgnv6FixedNatAlgFtpStats struct {
    DtCgnv6FixedNatAlgFtpStats Cgnv6FixedNatAlgFtpStats `json:"ftp"`
}


type Cgnv6FixedNatAlgFtpStatsStats struct {
    PortRequests int `json:"port-requests"`
    EprtRequests int `json:"eprt-requests"`
    LprtRequests int `json:"lprt-requests"`
    PasvReplies int `json:"pasv-replies"`
    EpsvReplies int `json:"epsv-replies"`
    LpsvReplies int `json:"lpsv-replies"`
}

func (p *Cgnv6FixedNatAlgFtpStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatAlgFtpStats) getPath() string{
    return "cgnv6/fixed-nat/alg/ftp/stats"
}

func (p *Cgnv6FixedNatAlgFtpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatAlgFtpStats,error) {
logger.Println("Cgnv6FixedNatAlgFtpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatAlgFtpStats
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
