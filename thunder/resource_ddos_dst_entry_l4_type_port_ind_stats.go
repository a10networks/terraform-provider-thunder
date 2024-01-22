package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryL4TypePortIndStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_l4_type_port_ind_stats`: Statistics for the object port-ind\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryL4TypePortIndStatsRead,

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
						"ddet_ind_pkt_drop_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop Rate Current",
						},
						"ddet_ind_pkt_drop_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop Rate Min",
						},
						"ddet_ind_pkt_drop_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "Pkt Drop Rate Max",
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
						"ddet_ind_fin_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Rate Current",
						},
						"ddet_ind_fin_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Rate Min",
						},
						"ddet_ind_fin_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Rate Max",
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
						"ddet_ind_small_window_ack_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Window ACK Rate Current",
						},
						"ddet_ind_small_window_ack_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Window ACK Rate Min",
						},
						"ddet_ind_small_window_ack_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Window ACK Rate Max",
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
						"ddet_ind_small_payload_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Payload Rate Current",
						},
						"ddet_ind_small_payload_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Payload Rate Min",
						},
						"ddet_ind_small_payload_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Small Payload Rate Max",
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
						"ddet_ind_inb_per_outb_current": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes-to / Bytes-from Current",
						},
						"ddet_ind_inb_per_outb_min": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes-to / Bytes-from Min",
						},
						"ddet_ind_inb_per_outb_max": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes-to / Bytes-from Max",
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
						"ddet_ind_conn_miss_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Miss Rate Current",
						},
						"ddet_ind_conn_miss_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Miss Rate Min",
						},
						"ddet_ind_conn_miss_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Miss Rate Max",
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
						"ddet_ind_data_cpu_util_current": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU Utilization Current",
						},
						"ddet_ind_data_cpu_util_min": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU Utilization Min",
						},
						"ddet_ind_data_cpu_util_max": {
							Type: schema.TypeInt, Optional: true, Description: "Data CPU Utilization Max",
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
						"ddet_ind_frag_rate_current": {
							Type: schema.TypeInt, Optional: true, Description: "Frag Pkt Rate Current",
						},
						"ddet_ind_frag_rate_min": {
							Type: schema.TypeInt, Optional: true, Description: "Frag Pkt Rate Min",
						},
						"ddet_ind_frag_rate_max": {
							Type: schema.TypeInt, Optional: true, Description: "Frag Pkt Rate Max",
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
					},
				},
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryL4TypePortIndStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypePortIndStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4TypePortIndStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryL4TypePortIndStatsStats := setObjectDdosDstEntryL4TypePortIndStatsStats(res)
		d.Set("stats", DdosDstEntryL4TypePortIndStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryL4TypePortIndStatsStats(ret edpt.DataDdosDstEntryL4TypePortIndStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ip_proto_type":                          ret.DtDdosDstEntryL4TypePortIndStats.Stats.IpProtoType,
			"ddet_ind_pkt_rate_current":              ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_rate_current,
			"ddet_ind_pkt_rate_min":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_rate_min,
			"ddet_ind_pkt_rate_max":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_rate_max,
			"ddet_ind_pkt_drop_rate_current":         ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_drop_rate_current,
			"ddet_ind_pkt_drop_rate_min":             ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_drop_rate_min,
			"ddet_ind_pkt_drop_rate_max":             ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_drop_rate_max,
			"ddet_ind_syn_rate_current":              ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_syn_rate_current,
			"ddet_ind_syn_rate_min":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_syn_rate_min,
			"ddet_ind_syn_rate_max":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_syn_rate_max,
			"ddet_ind_fin_rate_current":              ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_fin_rate_current,
			"ddet_ind_fin_rate_min":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_fin_rate_min,
			"ddet_ind_fin_rate_max":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_fin_rate_max,
			"ddet_ind_rst_rate_current":              ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_rst_rate_current,
			"ddet_ind_rst_rate_min":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_rst_rate_min,
			"ddet_ind_rst_rate_max":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_rst_rate_max,
			"ddet_ind_small_window_ack_rate_current": ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_small_window_ack_rate_current,
			"ddet_ind_small_window_ack_rate_min":     ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_small_window_ack_rate_min,
			"ddet_ind_small_window_ack_rate_max":     ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_small_window_ack_rate_max,
			"ddet_ind_empty_ack_rate_current":        ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_empty_ack_rate_current,
			"ddet_ind_empty_ack_rate_min":            ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_empty_ack_rate_min,
			"ddet_ind_empty_ack_rate_max":            ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_empty_ack_rate_max,
			"ddet_ind_small_payload_rate_current":    ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_small_payload_rate_current,
			"ddet_ind_small_payload_rate_min":        ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_small_payload_rate_min,
			"ddet_ind_small_payload_rate_max":        ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_small_payload_rate_max,
			"ddet_ind_pkt_drop_ratio_current":        ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_drop_ratio_current,
			"ddet_ind_pkt_drop_ratio_min":            ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_drop_ratio_min,
			"ddet_ind_pkt_drop_ratio_max":            ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_pkt_drop_ratio_max,
			"ddet_ind_inb_per_outb_current":          ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_inb_per_outb_current,
			"ddet_ind_inb_per_outb_min":              ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_inb_per_outb_min,
			"ddet_ind_inb_per_outb_max":              ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_inb_per_outb_max,
			"ddet_ind_syn_per_fin_rate_current":      ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_syn_per_fin_rate_current,
			"ddet_ind_syn_per_fin_rate_min":          ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_syn_per_fin_rate_min,
			"ddet_ind_syn_per_fin_rate_max":          ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_syn_per_fin_rate_max,
			"ddet_ind_conn_miss_rate_current":        ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_conn_miss_rate_current,
			"ddet_ind_conn_miss_rate_min":            ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_conn_miss_rate_min,
			"ddet_ind_conn_miss_rate_max":            ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_conn_miss_rate_max,
			"ddet_ind_concurrent_conns_current":      ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_concurrent_conns_current,
			"ddet_ind_concurrent_conns_min":          ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_concurrent_conns_min,
			"ddet_ind_concurrent_conns_max":          ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_concurrent_conns_max,
			"ddet_ind_data_cpu_util_current":         ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_data_cpu_util_current,
			"ddet_ind_data_cpu_util_min":             ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_data_cpu_util_min,
			"ddet_ind_data_cpu_util_max":             ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_data_cpu_util_max,
			"ddet_ind_outside_intf_util_current":     ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_outside_intf_util_current,
			"ddet_ind_outside_intf_util_min":         ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_outside_intf_util_min,
			"ddet_ind_outside_intf_util_max":         ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_outside_intf_util_max,
			"ddet_ind_frag_rate_current":             ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_frag_rate_current,
			"ddet_ind_frag_rate_min":                 ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_frag_rate_min,
			"ddet_ind_frag_rate_max":                 ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_frag_rate_max,
			"ddet_ind_bit_rate_current":              ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_bit_rate_current,
			"ddet_ind_bit_rate_min":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_bit_rate_min,
			"ddet_ind_bit_rate_max":                  ret.DtDdosDstEntryL4TypePortIndStats.Stats.Ddet_ind_bit_rate_max,
		},
	}
}

func getObjectDdosDstEntryL4TypePortIndStatsStats(d []interface{}) edpt.DdosDstEntryL4TypePortIndStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypePortIndStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProtoType = in["ip_proto_type"].(int)
		ret.Ddet_ind_pkt_rate_current = in["ddet_ind_pkt_rate_current"].(int)
		ret.Ddet_ind_pkt_rate_min = in["ddet_ind_pkt_rate_min"].(int)
		ret.Ddet_ind_pkt_rate_max = in["ddet_ind_pkt_rate_max"].(int)
		ret.Ddet_ind_pkt_drop_rate_current = in["ddet_ind_pkt_drop_rate_current"].(int)
		ret.Ddet_ind_pkt_drop_rate_min = in["ddet_ind_pkt_drop_rate_min"].(int)
		ret.Ddet_ind_pkt_drop_rate_max = in["ddet_ind_pkt_drop_rate_max"].(int)
		ret.Ddet_ind_syn_rate_current = in["ddet_ind_syn_rate_current"].(int)
		ret.Ddet_ind_syn_rate_min = in["ddet_ind_syn_rate_min"].(int)
		ret.Ddet_ind_syn_rate_max = in["ddet_ind_syn_rate_max"].(int)
		ret.Ddet_ind_fin_rate_current = in["ddet_ind_fin_rate_current"].(int)
		ret.Ddet_ind_fin_rate_min = in["ddet_ind_fin_rate_min"].(int)
		ret.Ddet_ind_fin_rate_max = in["ddet_ind_fin_rate_max"].(int)
		ret.Ddet_ind_rst_rate_current = in["ddet_ind_rst_rate_current"].(int)
		ret.Ddet_ind_rst_rate_min = in["ddet_ind_rst_rate_min"].(int)
		ret.Ddet_ind_rst_rate_max = in["ddet_ind_rst_rate_max"].(int)
		ret.Ddet_ind_small_window_ack_rate_current = in["ddet_ind_small_window_ack_rate_current"].(int)
		ret.Ddet_ind_small_window_ack_rate_min = in["ddet_ind_small_window_ack_rate_min"].(int)
		ret.Ddet_ind_small_window_ack_rate_max = in["ddet_ind_small_window_ack_rate_max"].(int)
		ret.Ddet_ind_empty_ack_rate_current = in["ddet_ind_empty_ack_rate_current"].(int)
		ret.Ddet_ind_empty_ack_rate_min = in["ddet_ind_empty_ack_rate_min"].(int)
		ret.Ddet_ind_empty_ack_rate_max = in["ddet_ind_empty_ack_rate_max"].(int)
		ret.Ddet_ind_small_payload_rate_current = in["ddet_ind_small_payload_rate_current"].(int)
		ret.Ddet_ind_small_payload_rate_min = in["ddet_ind_small_payload_rate_min"].(int)
		ret.Ddet_ind_small_payload_rate_max = in["ddet_ind_small_payload_rate_max"].(int)
		ret.Ddet_ind_pkt_drop_ratio_current = in["ddet_ind_pkt_drop_ratio_current"].(int)
		ret.Ddet_ind_pkt_drop_ratio_min = in["ddet_ind_pkt_drop_ratio_min"].(int)
		ret.Ddet_ind_pkt_drop_ratio_max = in["ddet_ind_pkt_drop_ratio_max"].(int)
		ret.Ddet_ind_inb_per_outb_current = in["ddet_ind_inb_per_outb_current"].(int)
		ret.Ddet_ind_inb_per_outb_min = in["ddet_ind_inb_per_outb_min"].(int)
		ret.Ddet_ind_inb_per_outb_max = in["ddet_ind_inb_per_outb_max"].(int)
		ret.Ddet_ind_syn_per_fin_rate_current = in["ddet_ind_syn_per_fin_rate_current"].(int)
		ret.Ddet_ind_syn_per_fin_rate_min = in["ddet_ind_syn_per_fin_rate_min"].(int)
		ret.Ddet_ind_syn_per_fin_rate_max = in["ddet_ind_syn_per_fin_rate_max"].(int)
		ret.Ddet_ind_conn_miss_rate_current = in["ddet_ind_conn_miss_rate_current"].(int)
		ret.Ddet_ind_conn_miss_rate_min = in["ddet_ind_conn_miss_rate_min"].(int)
		ret.Ddet_ind_conn_miss_rate_max = in["ddet_ind_conn_miss_rate_max"].(int)
		ret.Ddet_ind_concurrent_conns_current = in["ddet_ind_concurrent_conns_current"].(int)
		ret.Ddet_ind_concurrent_conns_min = in["ddet_ind_concurrent_conns_min"].(int)
		ret.Ddet_ind_concurrent_conns_max = in["ddet_ind_concurrent_conns_max"].(int)
		ret.Ddet_ind_data_cpu_util_current = in["ddet_ind_data_cpu_util_current"].(int)
		ret.Ddet_ind_data_cpu_util_min = in["ddet_ind_data_cpu_util_min"].(int)
		ret.Ddet_ind_data_cpu_util_max = in["ddet_ind_data_cpu_util_max"].(int)
		ret.Ddet_ind_outside_intf_util_current = in["ddet_ind_outside_intf_util_current"].(int)
		ret.Ddet_ind_outside_intf_util_min = in["ddet_ind_outside_intf_util_min"].(int)
		ret.Ddet_ind_outside_intf_util_max = in["ddet_ind_outside_intf_util_max"].(int)
		ret.Ddet_ind_frag_rate_current = in["ddet_ind_frag_rate_current"].(int)
		ret.Ddet_ind_frag_rate_min = in["ddet_ind_frag_rate_min"].(int)
		ret.Ddet_ind_frag_rate_max = in["ddet_ind_frag_rate_max"].(int)
		ret.Ddet_ind_bit_rate_current = in["ddet_ind_bit_rate_current"].(int)
		ret.Ddet_ind_bit_rate_min = in["ddet_ind_bit_rate_min"].(int)
		ret.Ddet_ind_bit_rate_max = in["ddet_ind_bit_rate_max"].(int)
	}
	return ret
}

func dataToEndpointDdosDstEntryL4TypePortIndStats(d *schema.ResourceData) edpt.DdosDstEntryL4TypePortIndStats {
	var ret edpt.DdosDstEntryL4TypePortIndStats

	ret.Stats = getObjectDdosDstEntryL4TypePortIndStatsStats(d.Get("stats").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
