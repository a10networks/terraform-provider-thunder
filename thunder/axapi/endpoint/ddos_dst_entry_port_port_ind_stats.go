

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortPortIndStats struct {
    
    Stats DdosDstEntryPortPortIndStatsStats `json:"stats"`
    Protocol string 
    PortNum string 
    DstEntryName string 

}
type DataDdosDstEntryPortPortIndStats struct {
    DtDdosDstEntryPortPortIndStats DdosDstEntryPortPortIndStats `json:"port-ind"`
}


type DdosDstEntryPortPortIndStatsStats struct {
    IpProtoType int `json:"ip-proto-type"`
    Ddet_ind_pkt_rate_current int `json:"ddet_ind_pkt_rate_current"`
    Ddet_ind_pkt_rate_min int `json:"ddet_ind_pkt_rate_min"`
    Ddet_ind_pkt_rate_max int `json:"ddet_ind_pkt_rate_max"`
    Ddet_ind_pkt_drop_rate_current int `json:"ddet_ind_pkt_drop_rate_current"`
    Ddet_ind_pkt_drop_rate_min int `json:"ddet_ind_pkt_drop_rate_min"`
    Ddet_ind_pkt_drop_rate_max int `json:"ddet_ind_pkt_drop_rate_max"`
    Ddet_ind_syn_rate_current int `json:"ddet_ind_syn_rate_current"`
    Ddet_ind_syn_rate_min int `json:"ddet_ind_syn_rate_min"`
    Ddet_ind_syn_rate_max int `json:"ddet_ind_syn_rate_max"`
    Ddet_ind_fin_rate_current int `json:"ddet_ind_fin_rate_current"`
    Ddet_ind_fin_rate_min int `json:"ddet_ind_fin_rate_min"`
    Ddet_ind_fin_rate_max int `json:"ddet_ind_fin_rate_max"`
    Ddet_ind_rst_rate_current int `json:"ddet_ind_rst_rate_current"`
    Ddet_ind_rst_rate_min int `json:"ddet_ind_rst_rate_min"`
    Ddet_ind_rst_rate_max int `json:"ddet_ind_rst_rate_max"`
    Ddet_ind_small_window_ack_rate_current int `json:"ddet_ind_small_window_ack_rate_current"`
    Ddet_ind_small_window_ack_rate_min int `json:"ddet_ind_small_window_ack_rate_min"`
    Ddet_ind_small_window_ack_rate_max int `json:"ddet_ind_small_window_ack_rate_max"`
    Ddet_ind_empty_ack_rate_current int `json:"ddet_ind_empty_ack_rate_current"`
    Ddet_ind_empty_ack_rate_min int `json:"ddet_ind_empty_ack_rate_min"`
    Ddet_ind_empty_ack_rate_max int `json:"ddet_ind_empty_ack_rate_max"`
    Ddet_ind_small_payload_rate_current int `json:"ddet_ind_small_payload_rate_current"`
    Ddet_ind_small_payload_rate_min int `json:"ddet_ind_small_payload_rate_min"`
    Ddet_ind_small_payload_rate_max int `json:"ddet_ind_small_payload_rate_max"`
    Ddet_ind_pkt_drop_ratio_current int `json:"ddet_ind_pkt_drop_ratio_current"`
    Ddet_ind_pkt_drop_ratio_min int `json:"ddet_ind_pkt_drop_ratio_min"`
    Ddet_ind_pkt_drop_ratio_max int `json:"ddet_ind_pkt_drop_ratio_max"`
    Ddet_ind_inb_per_outb_current int `json:"ddet_ind_inb_per_outb_current"`
    Ddet_ind_inb_per_outb_min int `json:"ddet_ind_inb_per_outb_min"`
    Ddet_ind_inb_per_outb_max int `json:"ddet_ind_inb_per_outb_max"`
    Ddet_ind_syn_per_fin_rate_current int `json:"ddet_ind_syn_per_fin_rate_current"`
    Ddet_ind_syn_per_fin_rate_min int `json:"ddet_ind_syn_per_fin_rate_min"`
    Ddet_ind_syn_per_fin_rate_max int `json:"ddet_ind_syn_per_fin_rate_max"`
    Ddet_ind_conn_miss_rate_current int `json:"ddet_ind_conn_miss_rate_current"`
    Ddet_ind_conn_miss_rate_min int `json:"ddet_ind_conn_miss_rate_min"`
    Ddet_ind_conn_miss_rate_max int `json:"ddet_ind_conn_miss_rate_max"`
    Ddet_ind_concurrent_conns_current int `json:"ddet_ind_concurrent_conns_current"`
    Ddet_ind_concurrent_conns_min int `json:"ddet_ind_concurrent_conns_min"`
    Ddet_ind_concurrent_conns_max int `json:"ddet_ind_concurrent_conns_max"`
    Ddet_ind_data_cpu_util_current int `json:"ddet_ind_data_cpu_util_current"`
    Ddet_ind_data_cpu_util_min int `json:"ddet_ind_data_cpu_util_min"`
    Ddet_ind_data_cpu_util_max int `json:"ddet_ind_data_cpu_util_max"`
    Ddet_ind_outside_intf_util_current int `json:"ddet_ind_outside_intf_util_current"`
    Ddet_ind_outside_intf_util_min int `json:"ddet_ind_outside_intf_util_min"`
    Ddet_ind_outside_intf_util_max int `json:"ddet_ind_outside_intf_util_max"`
    Ddet_ind_frag_rate_current int `json:"ddet_ind_frag_rate_current"`
    Ddet_ind_frag_rate_min int `json:"ddet_ind_frag_rate_min"`
    Ddet_ind_frag_rate_max int `json:"ddet_ind_frag_rate_max"`
    Ddet_ind_bit_rate_current int `json:"ddet_ind_bit_rate_current"`
    Ddet_ind_bit_rate_min int `json:"ddet_ind_bit_rate_min"`
    Ddet_ind_bit_rate_max int `json:"ddet_ind_bit_rate_max"`
}

func (p *DdosDstEntryPortPortIndStats) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortPortIndStats) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/port/" +p.PortNum + "+" +p.Protocol + "/port-ind/stats"
}

func (p *DdosDstEntryPortPortIndStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryPortPortIndStats,error) {
logger.Println("DdosDstEntryPortPortIndStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryPortPortIndStats
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
