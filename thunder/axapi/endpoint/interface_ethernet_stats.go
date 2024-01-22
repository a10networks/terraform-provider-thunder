

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetStats struct {
    
    Ifnum int `json:"ifnum"`
    Stats InterfaceEthernetStatsStats `json:"stats"`

}
type DataInterfaceEthernetStats struct {
    DtInterfaceEthernetStats InterfaceEthernetStats `json:"ethernet"`
}


type InterfaceEthernetStatsStats struct {
    Packets_input int `json:"packets_input"`
    Bytes_input int `json:"bytes_input"`
    Received_broadcasts int `json:"received_broadcasts"`
    Received_multicasts int `json:"received_multicasts"`
    Received_unicasts int `json:"received_unicasts"`
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Frame int `json:"frame"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Packets_output int `json:"packets_output"`
    Bytes_output int `json:"bytes_output"`
    Transmitted_broadcasts int `json:"transmitted_broadcasts"`
    Transmitted_multicasts int `json:"transmitted_multicasts"`
    Transmitted_unicasts int `json:"transmitted_unicasts"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Rate_pkt_sent int `json:"rate_pkt_sent"`
    Rate_byte_sent int `json:"rate_byte_sent"`
    Rate_pkt_rcvd int `json:"rate_pkt_rcvd"`
    Rate_byte_rcvd int `json:"rate_byte_rcvd"`
    Load_interval int `json:"load_interval"`
    Drops int `json:"drops"`
    Input_utilization int `json:"input_utilization"`
    Output_utilization int `json:"output_utilization"`
}

func (p *InterfaceEthernetStats) GetId() string{
    return "1"
}

func (p *InterfaceEthernetStats) getPath() string{
    
    return "interface/ethernet/" +strconv.Itoa(p.Ifnum)+"/stats"
}

func (p *InterfaceEthernetStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceEthernetStats,error) {
logger.Println("InterfaceEthernetStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceEthernetStats
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
