package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangePortIndStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_port_ind_stats`: Statistics for the object port-ind\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangePortIndStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_proto_type": {
							Type: schema.TypeInt, Optional: true, Description: "IP Protocol Type",
						},
						"ddet_ind_pkt_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Rate Current",
						},
						"ddet_ind_pkt_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Rate Min",
						},
						"ddet_ind_pkt_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Rate Max",
						},
						"ddet_ind_pkt_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Rate Adaptive Threshold",
						},
						"ddet_ind_pkt_drop_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop Rate Current",
						},
						"ddet_ind_pkt_drop_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop Rate Min",
						},
						"ddet_ind_pkt_drop_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop Rate Max",
						},
						"ddet_ind_pkt_drop_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop Rate Adaptive Threshold",
						},
						"ddet_ind_syn_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Current",
						},
						"ddet_ind_syn_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Min",
						},
						"ddet_ind_syn_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Max",
						},
						"ddet_ind_syn_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Adaptive Threshold",
						},
						"ddet_ind_fin_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Rate Current",
						},
						"ddet_ind_fin_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Rate Min",
						},
						"ddet_ind_fin_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Rate Max",
						},
						"ddet_ind_fin_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Rate Adaptive Threshold",
						},
						"ddet_ind_rst_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP RST Rate Current",
						},
						"ddet_ind_rst_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP RST Rate Min",
						},
						"ddet_ind_rst_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP RST Rate Max",
						},
						"ddet_ind_rst_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP RST Rate Adaptive Threshold",
						},
						"ddet_ind_small_window_ack_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Window ACK Rate Current",
						},
						"ddet_ind_small_window_ack_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Window ACK Rate Min",
						},
						"ddet_ind_small_window_ack_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Window ACK Rate Max",
						},
						"ddet_ind_small_window_ack_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Window ACK Rate Adaptive Threshold",
						},
						"ddet_ind_empty_ack_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Empty ACK Rate Current",
						},
						"ddet_ind_empty_ack_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Empty ACK Rate Min",
						},
						"ddet_ind_empty_ack_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Empty ACK Rate Max",
						},
						"ddet_ind_empty_ack_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Empty ACK Rate Adaptive Threshold",
						},
						"ddet_ind_small_payload_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Payload Rate Current",
						},
						"ddet_ind_small_payload_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Payload Rate Min",
						},
						"ddet_ind_small_payload_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Payload Rate Max",
						},
						"ddet_ind_small_payload_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Payload Rate Adaptive Threshold",
						},
						"ddet_ind_pkt_drop_ratio_current": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop / Pkt Rcvd Current",
						},
						"ddet_ind_pkt_drop_ratio_min": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop / Pkt Rcvd Min",
						},
						"ddet_ind_pkt_drop_ratio_max": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop / Pkt Rcvd Max",
						},
						"ddet_ind_pkt_drop_ratio_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop / Pkt Rcvd Adaptive Threshold",
						},
						"ddet_ind_inb_per_outb_current": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes-to / Bytes-from Current",
						},
						"ddet_ind_inb_per_outb_min": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes-to / Bytes-from Min",
						},
						"ddet_ind_inb_per_outb_max": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes-to / Bytes-from Max",
						},
						"ddet_ind_inb_per_outb_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes-to / Bytes-from Adaptive Threshold",
						},
						"ddet_ind_syn_per_fin_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate / FIN Rate Current",
						},
						"ddet_ind_syn_per_fin_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate / FIN Rate Min",
						},
						"ddet_ind_syn_per_fin_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate / FIN Rate Max",
						},
						"ddet_ind_syn_per_fin_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate / FIN Rate Adaptive Threshold",
						},
						"ddet_ind_conn_miss_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Miss Rate Current",
						},
						"ddet_ind_conn_miss_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Miss Rate Min",
						},
						"ddet_ind_conn_miss_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Miss Rate Max",
						},
						"ddet_ind_conn_miss_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Miss Rate Adaptive Threshold",
						},
						"ddet_ind_concurrent_conns_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP/UDP Concurrent Sessions Current",
						},
						"ddet_ind_concurrent_conns_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP/UDP Concurrent Sessions Min",
						},
						"ddet_ind_concurrent_conns_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP/UDP Concurrent Sessions Max",
						},
						"ddet_ind_concurrent_conns_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP/UDP Concurrent Sessions Adaptive Threshold",
						},
						"ddet_ind_data_cpu_util_current": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU Utilization Current",
						},
						"ddet_ind_data_cpu_util_min": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU Utilization Min",
						},
						"ddet_ind_data_cpu_util_max": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU Utilization Max",
						},
						"ddet_ind_data_cpu_util_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU Utilization Adaptive Threshold",
						},
						"ddet_ind_outside_intf_util_current": {
							Type: schema.TypeInt, Optional: true, Description: "Outside Interface Utilization Current",
						},
						"ddet_ind_outside_intf_util_min": {
							Type: schema.TypeInt, Optional: true, Description: "Outside Interface Utilization Min",
						},
						"ddet_ind_outside_intf_util_max": {
							Type: schema.TypeInt, Optional: true, Description: "Outside Interface Utilization Max",
						},
						"ddet_ind_outside_intf_util_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Outside Interface Utilization Adaptive Threshold",
						},
						"ddet_ind_frag_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "Frag Pkt Rate Current",
						},
						"ddet_ind_frag_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "Frag Pkt Rate Min",
						},
						"ddet_ind_frag_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "Frag Pkt Rate Max",
						},
						"ddet_ind_frag_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Frag Pkt Rate Adaptive Threshold",
						},
						"ddet_ind_bit_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "Bit Rate Current",
						},
						"ddet_ind_bit_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "Bit Rate Min",
						},
						"ddet_ind_bit_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "Bit Rate Max",
						},
						"ddet_ind_bit_rate_adaptive_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Bit Rate Adaptive Threshold",
						},
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}

func resourceDdosDstZonePortRangePortIndStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangePortIndStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangePortIndStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangePortIndStatsStats := setObjectDdosDstZonePortRangePortIndStatsStats(res)
		d.Set("stats", DdosDstZonePortRangePortIndStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangePortIndStatsStats(ret edpt.DataDdosDstZonePortRangePortIndStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_proto_type":                                     ret.DtDdosDstZonePortRangePortIndStats.Stats.IpProtoType,
			"ddet_ind_pkt_rate_current":                         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_rate_current,
			"ddet_ind_pkt_rate_min":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_rate_min,
			"ddet_ind_pkt_rate_max":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_rate_max,
			"ddet_ind_pkt_rate_adaptive_threshold":              ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_rate_adaptive_threshold,
			"ddet_ind_pkt_drop_rate_current":                    ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_rate_current,
			"ddet_ind_pkt_drop_rate_min":                        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_rate_min,
			"ddet_ind_pkt_drop_rate_max":                        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_rate_max,
			"ddet_ind_pkt_drop_rate_adaptive_threshold":         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_rate_adaptive_threshold,
			"ddet_ind_syn_rate_current":                         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_rate_current,
			"ddet_ind_syn_rate_min":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_rate_min,
			"ddet_ind_syn_rate_max":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_rate_max,
			"ddet_ind_syn_rate_adaptive_threshold":              ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_rate_adaptive_threshold,
			"ddet_ind_fin_rate_current":                         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_fin_rate_current,
			"ddet_ind_fin_rate_min":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_fin_rate_min,
			"ddet_ind_fin_rate_max":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_fin_rate_max,
			"ddet_ind_fin_rate_adaptive_threshold":              ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_fin_rate_adaptive_threshold,
			"ddet_ind_rst_rate_current":                         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_rst_rate_current,
			"ddet_ind_rst_rate_min":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_rst_rate_min,
			"ddet_ind_rst_rate_max":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_rst_rate_max,
			"ddet_ind_rst_rate_adaptive_threshold":              ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_rst_rate_adaptive_threshold,
			"ddet_ind_small_window_ack_rate_current":            ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_window_ack_rate_current,
			"ddet_ind_small_window_ack_rate_min":                ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_window_ack_rate_min,
			"ddet_ind_small_window_ack_rate_max":                ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_window_ack_rate_max,
			"ddet_ind_small_window_ack_rate_adaptive_threshold": ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_window_ack_rate_adaptive_threshold,
			"ddet_ind_empty_ack_rate_current":                   ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_empty_ack_rate_current,
			"ddet_ind_empty_ack_rate_min":                       ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_empty_ack_rate_min,
			"ddet_ind_empty_ack_rate_max":                       ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_empty_ack_rate_max,
			"ddet_ind_empty_ack_rate_adaptive_threshold":        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_empty_ack_rate_adaptive_threshold,
			"ddet_ind_small_payload_rate_current":               ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_payload_rate_current,
			"ddet_ind_small_payload_rate_min":                   ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_payload_rate_min,
			"ddet_ind_small_payload_rate_max":                   ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_payload_rate_max,
			"ddet_ind_small_payload_rate_adaptive_threshold":    ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_small_payload_rate_adaptive_threshold,
			"ddet_ind_pkt_drop_ratio_current":                   ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_ratio_current,
			"ddet_ind_pkt_drop_ratio_min":                       ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_ratio_min,
			"ddet_ind_pkt_drop_ratio_max":                       ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_ratio_max,
			"ddet_ind_pkt_drop_ratio_adaptive_threshold":        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_pkt_drop_ratio_adaptive_threshold,
			"ddet_ind_inb_per_outb_current":                     ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_inb_per_outb_current,
			"ddet_ind_inb_per_outb_min":                         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_inb_per_outb_min,
			"ddet_ind_inb_per_outb_max":                         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_inb_per_outb_max,
			"ddet_ind_inb_per_outb_adaptive_threshold":          ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_inb_per_outb_adaptive_threshold,
			"ddet_ind_syn_per_fin_rate_current":                 ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_per_fin_rate_current,
			"ddet_ind_syn_per_fin_rate_min":                     ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_per_fin_rate_min,
			"ddet_ind_syn_per_fin_rate_max":                     ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_per_fin_rate_max,
			"ddet_ind_syn_per_fin_rate_adaptive_threshold":      ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_syn_per_fin_rate_adaptive_threshold,
			"ddet_ind_conn_miss_rate_current":                   ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_conn_miss_rate_current,
			"ddet_ind_conn_miss_rate_min":                       ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_conn_miss_rate_min,
			"ddet_ind_conn_miss_rate_max":                       ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_conn_miss_rate_max,
			"ddet_ind_conn_miss_rate_adaptive_threshold":        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_conn_miss_rate_adaptive_threshold,
			"ddet_ind_concurrent_conns_current":                 ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_concurrent_conns_current,
			"ddet_ind_concurrent_conns_min":                     ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_concurrent_conns_min,
			"ddet_ind_concurrent_conns_max":                     ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_concurrent_conns_max,
			"ddet_ind_concurrent_conns_adaptive_threshold":      ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_concurrent_conns_adaptive_threshold,
			"ddet_ind_data_cpu_util_current":                    ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_data_cpu_util_current,
			"ddet_ind_data_cpu_util_min":                        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_data_cpu_util_min,
			"ddet_ind_data_cpu_util_max":                        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_data_cpu_util_max,
			"ddet_ind_data_cpu_util_adaptive_threshold":         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_data_cpu_util_adaptive_threshold,
			"ddet_ind_outside_intf_util_current":                ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_outside_intf_util_current,
			"ddet_ind_outside_intf_util_min":                    ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_outside_intf_util_min,
			"ddet_ind_outside_intf_util_max":                    ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_outside_intf_util_max,
			"ddet_ind_outside_intf_util_adaptive_threshold":     ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_outside_intf_util_adaptive_threshold,
			"ddet_ind_frag_rate_current":                        ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_frag_rate_current,
			"ddet_ind_frag_rate_min":                            ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_frag_rate_min,
			"ddet_ind_frag_rate_max":                            ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_frag_rate_max,
			"ddet_ind_frag_rate_adaptive_threshold":             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_frag_rate_adaptive_threshold,
			"ddet_ind_bit_rate_current":                         ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_bit_rate_current,
			"ddet_ind_bit_rate_min":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_bit_rate_min,
			"ddet_ind_bit_rate_max":                             ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_bit_rate_max,
			"ddet_ind_bit_rate_adaptive_threshold":              ret.DtDdosDstZonePortRangePortIndStats.Stats.Ddet_ind_bit_rate_adaptive_threshold,
		},
	}
}

func getObjectDdosDstZonePortRangePortIndStatsStats(d []interface{}) edpt.DdosDstZonePortRangePortIndStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangePortIndStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProtoType = in["ip_proto_type"].(int)
		ret.Ddet_ind_pkt_rate_current = in["ddet_ind_pkt_rate_current"].(int)
		ret.Ddet_ind_pkt_rate_min = in["ddet_ind_pkt_rate_min"].(int)
		ret.Ddet_ind_pkt_rate_max = in["ddet_ind_pkt_rate_max"].(int)
		ret.Ddet_ind_pkt_rate_adaptive_threshold = in["ddet_ind_pkt_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_pkt_drop_rate_current = in["ddet_ind_pkt_drop_rate_current"].(int)
		ret.Ddet_ind_pkt_drop_rate_min = in["ddet_ind_pkt_drop_rate_min"].(int)
		ret.Ddet_ind_pkt_drop_rate_max = in["ddet_ind_pkt_drop_rate_max"].(int)
		ret.Ddet_ind_pkt_drop_rate_adaptive_threshold = in["ddet_ind_pkt_drop_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_syn_rate_current = in["ddet_ind_syn_rate_current"].(int)
		ret.Ddet_ind_syn_rate_min = in["ddet_ind_syn_rate_min"].(int)
		ret.Ddet_ind_syn_rate_max = in["ddet_ind_syn_rate_max"].(int)
		ret.Ddet_ind_syn_rate_adaptive_threshold = in["ddet_ind_syn_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_fin_rate_current = in["ddet_ind_fin_rate_current"].(int)
		ret.Ddet_ind_fin_rate_min = in["ddet_ind_fin_rate_min"].(int)
		ret.Ddet_ind_fin_rate_max = in["ddet_ind_fin_rate_max"].(int)
		ret.Ddet_ind_fin_rate_adaptive_threshold = in["ddet_ind_fin_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_rst_rate_current = in["ddet_ind_rst_rate_current"].(int)
		ret.Ddet_ind_rst_rate_min = in["ddet_ind_rst_rate_min"].(int)
		ret.Ddet_ind_rst_rate_max = in["ddet_ind_rst_rate_max"].(int)
		ret.Ddet_ind_rst_rate_adaptive_threshold = in["ddet_ind_rst_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_small_window_ack_rate_current = in["ddet_ind_small_window_ack_rate_current"].(int)
		ret.Ddet_ind_small_window_ack_rate_min = in["ddet_ind_small_window_ack_rate_min"].(int)
		ret.Ddet_ind_small_window_ack_rate_max = in["ddet_ind_small_window_ack_rate_max"].(int)
		ret.Ddet_ind_small_window_ack_rate_adaptive_threshold = in["ddet_ind_small_window_ack_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_empty_ack_rate_current = in["ddet_ind_empty_ack_rate_current"].(int)
		ret.Ddet_ind_empty_ack_rate_min = in["ddet_ind_empty_ack_rate_min"].(int)
		ret.Ddet_ind_empty_ack_rate_max = in["ddet_ind_empty_ack_rate_max"].(int)
		ret.Ddet_ind_empty_ack_rate_adaptive_threshold = in["ddet_ind_empty_ack_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_small_payload_rate_current = in["ddet_ind_small_payload_rate_current"].(int)
		ret.Ddet_ind_small_payload_rate_min = in["ddet_ind_small_payload_rate_min"].(int)
		ret.Ddet_ind_small_payload_rate_max = in["ddet_ind_small_payload_rate_max"].(int)
		ret.Ddet_ind_small_payload_rate_adaptive_threshold = in["ddet_ind_small_payload_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_pkt_drop_ratio_current = in["ddet_ind_pkt_drop_ratio_current"].(int)
		ret.Ddet_ind_pkt_drop_ratio_min = in["ddet_ind_pkt_drop_ratio_min"].(int)
		ret.Ddet_ind_pkt_drop_ratio_max = in["ddet_ind_pkt_drop_ratio_max"].(int)
		ret.Ddet_ind_pkt_drop_ratio_adaptive_threshold = in["ddet_ind_pkt_drop_ratio_adaptive_threshold"].(int)
		ret.Ddet_ind_inb_per_outb_current = in["ddet_ind_inb_per_outb_current"].(int)
		ret.Ddet_ind_inb_per_outb_min = in["ddet_ind_inb_per_outb_min"].(int)
		ret.Ddet_ind_inb_per_outb_max = in["ddet_ind_inb_per_outb_max"].(int)
		ret.Ddet_ind_inb_per_outb_adaptive_threshold = in["ddet_ind_inb_per_outb_adaptive_threshold"].(int)
		ret.Ddet_ind_syn_per_fin_rate_current = in["ddet_ind_syn_per_fin_rate_current"].(int)
		ret.Ddet_ind_syn_per_fin_rate_min = in["ddet_ind_syn_per_fin_rate_min"].(int)
		ret.Ddet_ind_syn_per_fin_rate_max = in["ddet_ind_syn_per_fin_rate_max"].(int)
		ret.Ddet_ind_syn_per_fin_rate_adaptive_threshold = in["ddet_ind_syn_per_fin_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_conn_miss_rate_current = in["ddet_ind_conn_miss_rate_current"].(int)
		ret.Ddet_ind_conn_miss_rate_min = in["ddet_ind_conn_miss_rate_min"].(int)
		ret.Ddet_ind_conn_miss_rate_max = in["ddet_ind_conn_miss_rate_max"].(int)
		ret.Ddet_ind_conn_miss_rate_adaptive_threshold = in["ddet_ind_conn_miss_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_concurrent_conns_current = in["ddet_ind_concurrent_conns_current"].(int)
		ret.Ddet_ind_concurrent_conns_min = in["ddet_ind_concurrent_conns_min"].(int)
		ret.Ddet_ind_concurrent_conns_max = in["ddet_ind_concurrent_conns_max"].(int)
		ret.Ddet_ind_concurrent_conns_adaptive_threshold = in["ddet_ind_concurrent_conns_adaptive_threshold"].(int)
		ret.Ddet_ind_data_cpu_util_current = in["ddet_ind_data_cpu_util_current"].(int)
		ret.Ddet_ind_data_cpu_util_min = in["ddet_ind_data_cpu_util_min"].(int)
		ret.Ddet_ind_data_cpu_util_max = in["ddet_ind_data_cpu_util_max"].(int)
		ret.Ddet_ind_data_cpu_util_adaptive_threshold = in["ddet_ind_data_cpu_util_adaptive_threshold"].(int)
		ret.Ddet_ind_outside_intf_util_current = in["ddet_ind_outside_intf_util_current"].(int)
		ret.Ddet_ind_outside_intf_util_min = in["ddet_ind_outside_intf_util_min"].(int)
		ret.Ddet_ind_outside_intf_util_max = in["ddet_ind_outside_intf_util_max"].(int)
		ret.Ddet_ind_outside_intf_util_adaptive_threshold = in["ddet_ind_outside_intf_util_adaptive_threshold"].(int)
		ret.Ddet_ind_frag_rate_current = in["ddet_ind_frag_rate_current"].(int)
		ret.Ddet_ind_frag_rate_min = in["ddet_ind_frag_rate_min"].(int)
		ret.Ddet_ind_frag_rate_max = in["ddet_ind_frag_rate_max"].(int)
		ret.Ddet_ind_frag_rate_adaptive_threshold = in["ddet_ind_frag_rate_adaptive_threshold"].(int)
		ret.Ddet_ind_bit_rate_current = in["ddet_ind_bit_rate_current"].(int)
		ret.Ddet_ind_bit_rate_min = in["ddet_ind_bit_rate_min"].(int)
		ret.Ddet_ind_bit_rate_max = in["ddet_ind_bit_rate_max"].(int)
		ret.Ddet_ind_bit_rate_adaptive_threshold = in["ddet_ind_bit_rate_adaptive_threshold"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangePortIndStats(d *schema.ResourceData) edpt.DdosDstZonePortRangePortIndStats {
	var ret edpt.DdosDstZonePortRangePortIndStats

	ret.Stats = getObjectDdosDstZonePortRangePortIndStatsStats(d.Get("stats").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
