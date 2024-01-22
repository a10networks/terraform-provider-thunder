

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeStats struct {
    
    Ifnum int `json:"ifnum"`
    Stats InterfaceVeStatsStats `json:"stats"`

}
type DataInterfaceVeStats struct {
    DtInterfaceVeStats InterfaceVeStats `json:"ve"`
}


type InterfaceVeStatsStats struct {
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
    Rate_pkt_sent int `json:"rate_pkt_sent"`
    Rate_byte_sent int `json:"rate_byte_sent"`
    Rate_pkt_rcvd int `json:"rate_pkt_rcvd"`
    Rate_byte_rcvd int `json:"rate_byte_rcvd"`
    Load_interval int `json:"load_interval"`
}

func (p *InterfaceVeStats) GetId() string{
    return "1"
}

func (p *InterfaceVeStats) getPath() string{
    
    return "interface/ve/" +strconv.Itoa(p.Ifnum)+"/stats"
}

func (p *InterfaceVeStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceVeStats,error) {
logger.Println("InterfaceVeStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceVeStats
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
