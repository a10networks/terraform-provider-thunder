

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTunnelStats struct {
    
    Ifnum int `json:"ifnum"`
    Stats InterfaceTunnelStatsStats `json:"stats"`

}
type DataInterfaceTunnelStats struct {
    DtInterfaceTunnelStats InterfaceTunnelStats `json:"tunnel"`
}


type InterfaceTunnelStatsStats struct {
    NumRxPkts int `json:"num-rx-pkts"`
    NumTotalRxBytes int `json:"num-total-rx-bytes"`
    NumTxPkts int `json:"num-tx-pkts"`
    NumTotalTxBytes int `json:"num-total-tx-bytes"`
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Rate_pkt_sent int `json:"rate_pkt_sent"`
    Rate_byte_sent int `json:"rate_byte_sent"`
    Rate_pkt_rcvd int `json:"rate_pkt_rcvd"`
    Rate_byte_rcvd int `json:"rate_byte_rcvd"`
}

func (p *InterfaceTunnelStats) GetId() string{
    return "1"
}

func (p *InterfaceTunnelStats) getPath() string{
    
    return "interface/tunnel/" +strconv.Itoa(p.Ifnum)+"/stats"
}

func (p *InterfaceTunnelStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceTunnelStats,error) {
logger.Println("InterfaceTunnelStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceTunnelStats
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
