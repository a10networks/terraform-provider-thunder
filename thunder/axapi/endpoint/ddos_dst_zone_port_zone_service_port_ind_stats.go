

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServicePortIndStats struct {
    
    Stats DdosDstZonePortZoneServicePortIndStatsStats `json:"stats"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServicePortIndStats struct {
    DtDdosDstZonePortZoneServicePortIndStats DdosDstZonePortZoneServicePortIndStats `json:"port-ind"`
}


type DdosDstZonePortZoneServicePortIndStatsStats struct {
    IpProtoType int `json:"ip-proto-type"`
    Ddet_ind_pkt_rate_current int `json:"ddet_ind_pkt_rate_current"`
    Ddet_ind_pkt_rate_min int `json:"ddet_ind_pkt_rate_min"`
    Ddet_ind_pkt_rate_max int `json:"ddet_ind_pkt_rate_max"`
    Ddet_ind_pkt_rate_adaptive_threshold int `json:"ddet_ind_pkt_rate_adaptive_threshold"`
    Ddet_ind_pkt_drop_rate_current int `json:"ddet_ind_pkt_drop_rate_current"`
    Ddet_ind_pkt_drop_rate_min int `json:"ddet_ind_pkt_drop_rate_min"`
    Ddet_ind_pkt_drop_rate_max int `json:"ddet_ind_pkt_drop_rate_max"`
    Ddet_ind_pkt_drop_rate_adaptive_threshold int `json:"ddet_ind_pkt_drop_rate_adaptive_threshold"`
    Ddet_ind_syn_rate_current int `json:"ddet_ind_syn_rate_current"`
    Ddet_ind_syn_rate_min int `json:"ddet_ind_syn_rate_min"`
    Ddet_ind_syn_rate_max int `json:"ddet_ind_syn_rate_max"`
    Ddet_ind_syn_rate_adaptive_threshold int `json:"ddet_ind_syn_rate_adaptive_threshold"`
    Ddet_ind_fin_rate_current int `json:"ddet_ind_fin_rate_current"`
    Ddet_ind_fin_rate_min int `json:"ddet_ind_fin_rate_min"`
    Ddet_ind_fin_rate_max int `json:"ddet_ind_fin_rate_max"`
    Ddet_ind_fin_rate_adaptive_threshold int `json:"ddet_ind_fin_rate_adaptive_threshold"`
    Ddet_ind_rst_rate_current int `json:"ddet_ind_rst_rate_current"`
    Ddet_ind_rst_rate_min int `json:"ddet_ind_rst_rate_min"`
    Ddet_ind_rst_rate_max int `json:"ddet_ind_rst_rate_max"`
    Ddet_ind_rst_rate_adaptive_threshold int `json:"ddet_ind_rst_rate_adaptive_threshold"`
    Ddet_ind_small_window_ack_rate_current int `json:"ddet_ind_small_window_ack_rate_current"`
    Ddet_ind_small_window_ack_rate_min int `json:"ddet_ind_small_window_ack_rate_min"`
    Ddet_ind_small_window_ack_rate_max int `json:"ddet_ind_small_window_ack_rate_max"`
    Ddet_ind_small_window_ack_rate_adaptive_threshold int `json:"ddet_ind_small_window_ack_rate_adaptive_threshold"`
    Ddet_ind_empty_ack_rate_current int `json:"ddet_ind_empty_ack_rate_current"`
    Ddet_ind_empty_ack_rate_min int `json:"ddet_ind_empty_ack_rate_min"`
    Ddet_ind_empty_ack_rate_max int `json:"ddet_ind_empty_ack_rate_max"`
    Ddet_ind_empty_ack_rate_adaptive_threshold int `json:"ddet_ind_empty_ack_rate_adaptive_threshold"`
    Ddet_ind_small_payload_rate_current int `json:"ddet_ind_small_payload_rate_current"`
    Ddet_ind_small_payload_rate_min int `json:"ddet_ind_small_payload_rate_min"`
    Ddet_ind_small_payload_rate_max int `json:"ddet_ind_small_payload_rate_max"`
    Ddet_ind_small_payload_rate_adaptive_threshold int `json:"ddet_ind_small_payload_rate_adaptive_threshold"`
    Ddet_ind_pkt_drop_ratio_current int `json:"ddet_ind_pkt_drop_ratio_current"`
    Ddet_ind_pkt_drop_ratio_min int `json:"ddet_ind_pkt_drop_ratio_min"`
    Ddet_ind_pkt_drop_ratio_max int `json:"ddet_ind_pkt_drop_ratio_max"`
    Ddet_ind_pkt_drop_ratio_adaptive_threshold int `json:"ddet_ind_pkt_drop_ratio_adaptive_threshold"`
    Ddet_ind_inb_per_outb_current int `json:"ddet_ind_inb_per_outb_current"`
    Ddet_ind_inb_per_outb_min int `json:"ddet_ind_inb_per_outb_min"`
    Ddet_ind_inb_per_outb_max int `json:"ddet_ind_inb_per_outb_max"`
    Ddet_ind_inb_per_outb_adaptive_threshold int `json:"ddet_ind_inb_per_outb_adaptive_threshold"`
    Ddet_ind_syn_per_fin_rate_current int `json:"ddet_ind_syn_per_fin_rate_current"`
    Ddet_ind_syn_per_fin_rate_min int `json:"ddet_ind_syn_per_fin_rate_min"`
    Ddet_ind_syn_per_fin_rate_max int `json:"ddet_ind_syn_per_fin_rate_max"`
    Ddet_ind_syn_per_fin_rate_adaptive_threshold int `json:"ddet_ind_syn_per_fin_rate_adaptive_threshold"`
    Ddet_ind_conn_miss_rate_current int `json:"ddet_ind_conn_miss_rate_current"`
    Ddet_ind_conn_miss_rate_min int `json:"ddet_ind_conn_miss_rate_min"`
    Ddet_ind_conn_miss_rate_max int `json:"ddet_ind_conn_miss_rate_max"`
    Ddet_ind_conn_miss_rate_adaptive_threshold int `json:"ddet_ind_conn_miss_rate_adaptive_threshold"`
    Ddet_ind_concurrent_conns_current int `json:"ddet_ind_concurrent_conns_current"`
    Ddet_ind_concurrent_conns_min int `json:"ddet_ind_concurrent_conns_min"`
    Ddet_ind_concurrent_conns_max int `json:"ddet_ind_concurrent_conns_max"`
    Ddet_ind_concurrent_conns_adaptive_threshold int `json:"ddet_ind_concurrent_conns_adaptive_threshold"`
    Ddet_ind_data_cpu_util_current int `json:"ddet_ind_data_cpu_util_current"`
    Ddet_ind_data_cpu_util_min int `json:"ddet_ind_data_cpu_util_min"`
    Ddet_ind_data_cpu_util_max int `json:"ddet_ind_data_cpu_util_max"`
    Ddet_ind_data_cpu_util_adaptive_threshold int `json:"ddet_ind_data_cpu_util_adaptive_threshold"`
    Ddet_ind_outside_intf_util_current int `json:"ddet_ind_outside_intf_util_current"`
    Ddet_ind_outside_intf_util_min int `json:"ddet_ind_outside_intf_util_min"`
    Ddet_ind_outside_intf_util_max int `json:"ddet_ind_outside_intf_util_max"`
    Ddet_ind_outside_intf_util_adaptive_threshold int `json:"ddet_ind_outside_intf_util_adaptive_threshold"`
    Ddet_ind_frag_rate_current int `json:"ddet_ind_frag_rate_current"`
    Ddet_ind_frag_rate_min int `json:"ddet_ind_frag_rate_min"`
    Ddet_ind_frag_rate_max int `json:"ddet_ind_frag_rate_max"`
    Ddet_ind_frag_rate_adaptive_threshold int `json:"ddet_ind_frag_rate_adaptive_threshold"`
    Ddet_ind_bit_rate_current int `json:"ddet_ind_bit_rate_current"`
    Ddet_ind_bit_rate_min int `json:"ddet_ind_bit_rate_min"`
    Ddet_ind_bit_rate_max int `json:"ddet_ind_bit_rate_max"`
    Ddet_ind_bit_rate_adaptive_threshold int `json:"ddet_ind_bit_rate_adaptive_threshold"`
}

func (p *DdosDstZonePortZoneServicePortIndStats) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServicePortIndStats) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/port-ind/stats"
}

func (p *DdosDstZonePortZoneServicePortIndStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServicePortIndStats,error) {
logger.Println("DdosDstZonePortZoneServicePortIndStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServicePortIndStats
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
