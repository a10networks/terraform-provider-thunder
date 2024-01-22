

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatAlgMgcpStats struct {
    
    Stats Cgnv6FixedNatAlgMgcpStatsStats `json:"stats"`

}
type DataCgnv6FixedNatAlgMgcpStats struct {
    DtCgnv6FixedNatAlgMgcpStats Cgnv6FixedNatAlgMgcpStats `json:"mgcp"`
}


type Cgnv6FixedNatAlgMgcpStatsStats struct {
    Auep int `json:"auep"`
    Aucx int `json:"aucx"`
    Crcx int `json:"crcx"`
    Dlcx int `json:"dlcx"`
    Epcf int `json:"epcf"`
    Mdcx int `json:"mdcx"`
    Ntfy int `json:"ntfy"`
    Rqnt int `json:"rqnt"`
    Rsip int `json:"rsip"`
}

func (p *Cgnv6FixedNatAlgMgcpStats) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatAlgMgcpStats) getPath() string{
    return "cgnv6/fixed-nat/alg/mgcp/stats"
}

func (p *Cgnv6FixedNatAlgMgcpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatAlgMgcpStats,error) {
logger.Println("Cgnv6FixedNatAlgMgcpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatAlgMgcpStats
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
