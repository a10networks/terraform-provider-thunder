

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkStats struct {
    
    Ifnum int `json:"ifnum"`
    Stats InterfaceTrunkStatsStats `json:"stats"`

}
type DataInterfaceTrunkStats struct {
    DtInterfaceTrunkStats InterfaceTrunkStats `json:"trunk"`
}


type InterfaceTrunkStatsStats struct {
    Num_pkts int `json:"num_pkts"`
    Num_total_bytes int `json:"num_total_bytes"`
    Num_unicast_pkts int `json:"num_unicast_pkts"`
    Num_broadcast_pkts int `json:"num_broadcast_pkts"`
    Num_multicast_pkts int `json:"num_multicast_pkts"`
    Num_tx_pkts int `json:"num_tx_pkts"`
    Num_total_tx_bytes int `json:"num_total_tx_bytes"`
    Num_unicast_tx_pkts int `json:"num_unicast_tx_pkts"`
    Num_broadcast_tx_pkts int `json:"num_broadcast_tx_pkts"`
    Num_multicast_tx_pkts int `json:"num_multicast_tx_pkts"`
    Dropped_dis_rx_pkts int `json:"dropped_dis_rx_pkts"`
    Dropped_rx_pkts int `json:"dropped_rx_pkts"`
    Dropped_dis_tx_pkts int `json:"dropped_dis_tx_pkts"`
    Dropped_tx_pkts int `json:"dropped_tx_pkts"`
}

func (p *InterfaceTrunkStats) GetId() string{
    return "1"
}

func (p *InterfaceTrunkStats) getPath() string{
    
    return "interface/trunk/" +strconv.Itoa(p.Ifnum)+"/stats"
}

func (p *InterfaceTrunkStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceTrunkStats,error) {
logger.Println("InterfaceTrunkStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceTrunkStats
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
