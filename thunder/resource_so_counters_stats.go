package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSoCountersStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_so_counters_stats`: Statistics for the object so-counters\n\n__PLACEHOLDER__",
		ReadContext: resourceSoCountersStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"so_pkts_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Total data packets received",
						},
						"so_redirect_pkts_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total packets redirected out of node",
						},
						"so_pkts_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Total packets dropped",
						},
						"so_redirected_pkts_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Total redirected packets received on node",
						},
						"so_fw_shadow_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "FW Shadow Session created",
						},
						"so_pkts_traffic_map_not_found_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Traffic MAP Not Found Drop",
						},
						"so_slb_pkts_redirect_conn_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "Total SLB redirect conns aged out",
						},
						"so_pkts_scaleout_not_active_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Scaleout Not Active Drop",
						},
						"so_pkts_slb_nat_reserve_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total SLB NAT reserve failures",
						},
						"so_pkts_slb_nat_release_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total SLB NAT release failures",
						},
						"so_pkts_dest_mac_mismatch_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Destination MAC Mistmatch Drop",
						},
						"so_pkts_l2redirect_dest_mac_zero_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Destination MAC Address zero Drop",
						},
						"so_pkts_l2redirect_interface_not_up": {
							Type: schema.TypeInt, Optional: true, Description: "L2redirect Intf is not UP",
						},
						"so_pkts_l2redirect_invalid_redirect_info_error": {
							Type: schema.TypeInt, Optional: true, Description: "Redirect Table Error due to invalid redirect info",
						},
						"so_pkts_l3_redirect_encap_error_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect encap error drop during transmission",
						},
						"so_pkts_l3_redirect_inner_mac_zero_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect inner mac zero drop during transmission",
						},
						"so_pkts_l3_redirect_decap_vlan_sanity_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect Decap VLAN Sanity Drop during receipt",
						},
						"so_pkts_l3_redirect_decap_non_ipv4_vxlan_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect received non ipv4 VXLAN packet",
						},
						"so_pkts_l3_redirect_decap_rx_encap_params_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect decap Rx encap params error Drop",
						},
						"so_pkts_l3_redirect_table_error": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect Table error Drop",
						},
						"so_pkts_l3_redirect_rcvd_in_l2_mode_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Recevied l3 redirect packets in L2 mode Drop",
						},
						"so_pkts_l3_redirect_fragmentation_error": {
							Type: schema.TypeInt, Optional: true, Description: "L3 redirect encap Fragmentation error",
						},
						"so_pkts_l3_redirect_table_no_entry_found": {
							Type: schema.TypeInt, Optional: true, Description: "L3 redirect Table no redirect entry found error",
						},
						"so_pkts_l3_redirect_invalid_dev_dir": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect Invalid Device direction during transmission",
						},
						"so_pkts_l3_redirect_chassis_dest_mac_error": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect RX multi-slot Destination MAC Error",
						},
						"so_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect ipv4 packet after encap more than max jumbo size",
						},
						"so_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L3 Redirect tx ipv6 packet after encap more than max jumbo size",
						},
						"so_pkts_l3_redirect_too_large_pkts_in_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Received L3 Redirected fragmented packets too large",
						},
						"so_pkts_l2redirect_vlan_retrieval_error": {
							Type: schema.TypeInt, Optional: true, Description: "L2 redirect pkt vlan not retrieved",
						},
						"so_pkts_l2redirect_port_retrieval_error": {
							Type: schema.TypeInt, Optional: true, Description: "L2 redirect pkt port not retrieved",
						},
						"so_slb_shadow_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "SLB Shadow Session created",
						},
					},
				},
			},
		},
	}
}

func resourceSoCountersStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSoCountersStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSoCountersStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SoCountersStatsStats := setObjectSoCountersStatsStats(res)
		d.Set("stats", SoCountersStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSoCountersStatsStats(ret edpt.DataSoCountersStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"so_pkts_rcvd":                                   ret.DtSoCountersStats.Stats.So_pkts_rcvd,
			"so_redirect_pkts_sent":                          ret.DtSoCountersStats.Stats.So_redirect_pkts_sent,
			"so_pkts_dropped":                                ret.DtSoCountersStats.Stats.So_pkts_dropped,
			"so_redirected_pkts_rcvd":                        ret.DtSoCountersStats.Stats.So_redirected_pkts_rcvd,
			"so_fw_shadow_session_created":                   ret.DtSoCountersStats.Stats.So_fw_shadow_session_created,
			"so_pkts_traffic_map_not_found_drop":             ret.DtSoCountersStats.Stats.So_pkts_traffic_map_not_found_drop,
			"so_slb_pkts_redirect_conn_aged_out":             ret.DtSoCountersStats.Stats.So_slb_pkts_redirect_conn_aged_out,
			"so_pkts_scaleout_not_active_drop":               ret.DtSoCountersStats.Stats.So_pkts_scaleout_not_active_drop,
			"so_pkts_slb_nat_reserve_fail":                   ret.DtSoCountersStats.Stats.So_pkts_slb_nat_reserve_fail,
			"so_pkts_slb_nat_release_fail":                   ret.DtSoCountersStats.Stats.So_pkts_slb_nat_release_fail,
			"so_pkts_dest_mac_mismatch_drop":                 ret.DtSoCountersStats.Stats.So_pkts_dest_mac_mismatch_drop,
			"so_pkts_l2redirect_dest_mac_zero_drop":          ret.DtSoCountersStats.Stats.So_pkts_l2redirect_dest_mac_zero_drop,
			"so_pkts_l2redirect_interface_not_up":            ret.DtSoCountersStats.Stats.So_pkts_l2redirect_interface_not_up,
			"so_pkts_l2redirect_invalid_redirect_info_error": ret.DtSoCountersStats.Stats.So_pkts_l2redirect_invalid_redirect_info_error,
			"so_pkts_l3_redirect_encap_error_drop":           ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_encap_error_drop,
			"so_pkts_l3_redirect_inner_mac_zero_drop":        ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_inner_mac_zero_drop,
			"so_pkts_l3_redirect_decap_vlan_sanity_drop":     ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_decap_vlan_sanity_drop,
			"so_pkts_l3_redirect_decap_non_ipv4_vxlan_drop":  ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_decap_non_ipv4_vxlan_drop,
			"so_pkts_l3_redirect_decap_rx_encap_params_drop": ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_decap_rx_encap_params_drop,
			"so_pkts_l3_redirect_table_error":                ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_table_error,
			"so_pkts_l3_redirect_rcvd_in_l2_mode_drop":       ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_rcvd_in_l2_mode_drop,
			"so_pkts_l3_redirect_fragmentation_error":        ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_fragmentation_error,
			"so_pkts_l3_redirect_table_no_entry_found":       ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_table_no_entry_found,
			"so_pkts_l3_redirect_invalid_dev_dir":            ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_invalid_dev_dir,
			"so_pkts_l3_redirect_chassis_dest_mac_error":     ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_chassis_dest_mac_error,
			"so_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop": ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop,
			"so_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop": ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop,
			"so_pkts_l3_redirect_too_large_pkts_in_drop":     ret.DtSoCountersStats.Stats.So_pkts_l3_redirect_too_large_pkts_in_drop,
			"so_pkts_l2redirect_vlan_retrieval_error":        ret.DtSoCountersStats.Stats.So_pkts_l2redirect_vlan_retrieval_error,
			"so_pkts_l2redirect_port_retrieval_error":        ret.DtSoCountersStats.Stats.So_pkts_l2redirect_port_retrieval_error,
			"so_slb_shadow_session_created":                  ret.DtSoCountersStats.Stats.So_slb_shadow_session_created,
		},
	}
}

func getObjectSoCountersStatsStats(d []interface{}) edpt.SoCountersStatsStats {

	count1 := len(d)
	var ret edpt.SoCountersStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.So_pkts_rcvd = in["so_pkts_rcvd"].(int)
		ret.So_redirect_pkts_sent = in["so_redirect_pkts_sent"].(int)
		ret.So_pkts_dropped = in["so_pkts_dropped"].(int)
		ret.So_redirected_pkts_rcvd = in["so_redirected_pkts_rcvd"].(int)
		ret.So_fw_shadow_session_created = in["so_fw_shadow_session_created"].(int)
		ret.So_pkts_traffic_map_not_found_drop = in["so_pkts_traffic_map_not_found_drop"].(int)
		ret.So_slb_pkts_redirect_conn_aged_out = in["so_slb_pkts_redirect_conn_aged_out"].(int)
		ret.So_pkts_scaleout_not_active_drop = in["so_pkts_scaleout_not_active_drop"].(int)
		ret.So_pkts_slb_nat_reserve_fail = in["so_pkts_slb_nat_reserve_fail"].(int)
		ret.So_pkts_slb_nat_release_fail = in["so_pkts_slb_nat_release_fail"].(int)
		ret.So_pkts_dest_mac_mismatch_drop = in["so_pkts_dest_mac_mismatch_drop"].(int)
		ret.So_pkts_l2redirect_dest_mac_zero_drop = in["so_pkts_l2redirect_dest_mac_zero_drop"].(int)
		ret.So_pkts_l2redirect_interface_not_up = in["so_pkts_l2redirect_interface_not_up"].(int)
		ret.So_pkts_l2redirect_invalid_redirect_info_error = in["so_pkts_l2redirect_invalid_redirect_info_error"].(int)
		ret.So_pkts_l3_redirect_encap_error_drop = in["so_pkts_l3_redirect_encap_error_drop"].(int)
		ret.So_pkts_l3_redirect_inner_mac_zero_drop = in["so_pkts_l3_redirect_inner_mac_zero_drop"].(int)
		ret.So_pkts_l3_redirect_decap_vlan_sanity_drop = in["so_pkts_l3_redirect_decap_vlan_sanity_drop"].(int)
		ret.So_pkts_l3_redirect_decap_non_ipv4_vxlan_drop = in["so_pkts_l3_redirect_decap_non_ipv4_vxlan_drop"].(int)
		ret.So_pkts_l3_redirect_decap_rx_encap_params_drop = in["so_pkts_l3_redirect_decap_rx_encap_params_drop"].(int)
		ret.So_pkts_l3_redirect_table_error = in["so_pkts_l3_redirect_table_error"].(int)
		ret.So_pkts_l3_redirect_rcvd_in_l2_mode_drop = in["so_pkts_l3_redirect_rcvd_in_l2_mode_drop"].(int)
		ret.So_pkts_l3_redirect_fragmentation_error = in["so_pkts_l3_redirect_fragmentation_error"].(int)
		ret.So_pkts_l3_redirect_table_no_entry_found = in["so_pkts_l3_redirect_table_no_entry_found"].(int)
		ret.So_pkts_l3_redirect_invalid_dev_dir = in["so_pkts_l3_redirect_invalid_dev_dir"].(int)
		ret.So_pkts_l3_redirect_chassis_dest_mac_error = in["so_pkts_l3_redirect_chassis_dest_mac_error"].(int)
		ret.So_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop = in["so_pkts_l3_redirect_encap_ipv4_jumbo_frag_drop"].(int)
		ret.So_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop = in["so_pkts_l3_redirect_encap_ipv6_jumbo_frag_drop"].(int)
		ret.So_pkts_l3_redirect_too_large_pkts_in_drop = in["so_pkts_l3_redirect_too_large_pkts_in_drop"].(int)
		ret.So_pkts_l2redirect_vlan_retrieval_error = in["so_pkts_l2redirect_vlan_retrieval_error"].(int)
		ret.So_pkts_l2redirect_port_retrieval_error = in["so_pkts_l2redirect_port_retrieval_error"].(int)
		ret.So_slb_shadow_session_created = in["so_slb_shadow_session_created"].(int)
	}
	return ret
}

func dataToEndpointSoCountersStats(d *schema.ResourceData) edpt.SoCountersStats {
	var ret edpt.SoCountersStats

	ret.Stats = getObjectSoCountersStatsStats(d.Get("stats").([]interface{}))
	return ret
}
