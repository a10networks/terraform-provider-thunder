

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgMgcpStats struct {
    
    Stats Cgnv6LsnAlgMgcpStatsStats `json:"stats"`

}
type DataCgnv6LsnAlgMgcpStats struct {
    DtCgnv6LsnAlgMgcpStats Cgnv6LsnAlgMgcpStats `json:"mgcp"`
}


type Cgnv6LsnAlgMgcpStatsStats struct {
    Auep int `json:"auep"`
    Aucx int `json:"aucx"`
    Crcx int `json:"crcx"`
    Dlcx int `json:"dlcx"`
    Epcf int `json:"epcf"`
    Mdcx int `json:"mdcx"`
    Ntfy int `json:"ntfy"`
    Rqnt int `json:"rqnt"`
    Rsip int `json:"rsip"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
}

func (p *Cgnv6LsnAlgMgcpStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgMgcpStats) getPath() string{
    return "cgnv6/lsn/alg/mgcp/stats"
}

func (p *Cgnv6LsnAlgMgcpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAlgMgcpStats,error) {
logger.Println("Cgnv6LsnAlgMgcpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAlgMgcpStats
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
