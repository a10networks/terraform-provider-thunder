package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHardwareAccelerateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_hardware_accelerate_stats`: Statistics for the object hardware-accelerate\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemHardwareAccelerateStatsRead,

		Schema: map[string]*schema.Schema{
			"slb": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_create": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Created",
									},
									"entry_create_failure": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Created",
									},
									"entry_create_fail_server_down": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Created",
									},
									"entry_create_fail_max_entry": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Created",
									},
									"entry_free": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Freed",
									},
									"entry_free_opp_entry": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Free due to opposite tuple entry ageout event",
									},
									"entry_free_no_hw_prog": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entry Free no hw prog",
									},
									"entry_free_no_conn": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entry Free no matched conn",
									},
									"entry_free_no_sw_entry": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entry Free no software entry",
									},
									"entry_counter": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entry Count",
									},
									"entry_age_out": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Aged Out",
									},
									"entry_age_out_idle": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Aged Out Idle",
									},
									"entry_age_out_tcp_fin": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Aged Out TCP FIN",
									},
									"entry_age_out_tcp_rst": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Aged Out TCP RST",
									},
									"entry_age_out_invalid_dst": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Aged Out invalid dst",
									},
									"entry_force_hw_invalidate": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Force HW Invalidate",
									},
									"entry_invalidate_server_down": {
										Type: schema.TypeInt, Optional: true, Description: "Hardware Entries Invalidate due to server down",
									},
									"tcam_create": {
										Type: schema.TypeInt, Optional: true, Description: "TCAM Entries Created",
									},
									"tcam_free": {
										Type: schema.TypeInt, Optional: true, Description: "TCAM Entries Freed",
									},
									"tcam_counter": {
										Type: schema.TypeInt, Optional: true, Description: "TCAM Entry Count",
									},
								},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hit_counts": {
							Type: schema.TypeInt, Optional: true, Description: "Total packets hit counts",
						},
						"hit_index": {
							Type: schema.TypeInt, Optional: true, Description: "HW Fwd hit index",
						},
						"ipv4_forward_counts": {
							Type: schema.TypeInt, Optional: true, Description: "Total IPv4 hardware forwarded packets",
						},
						"ipv6_forward_counts": {
							Type: schema.TypeInt, Optional: true, Description: "Total IPv6 hardware forwarded packets",
						},
						"hw_fwd_module_status": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forwarder status flags",
						},
						"hw_fwd_prog_reqs": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward programming requests",
						},
						"hw_fwd_prog_errors": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward programming Errors",
						},
						"hw_fwd_flow_singlebit_errors": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward singlebit Errors",
						},
						"hw_fwd_flow_tag_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward tag mismatch errors",
						},
						"hw_fwd_flow_seq_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward sequence mismatch errors",
						},
						"hw_fwd_ageout_drop_count": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward ageout drop count",
						},
						"hw_fwd_invalidation_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward invalid drop count",
						},
						"hw_fwd_flow_hit_index": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow hit index",
						},
						"hw_fwd_flow_reason_flags": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow reason flags",
						},
						"hw_fwd_flow_drop_count": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow drop count",
						},
						"hw_fwd_flow_error_count": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow error count",
						},
						"hw_fwd_flow_unalign_count": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow unalign count",
						},
						"hw_fwd_flow_underflow_count": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow underflow count",
						},
						"hw_fwd_flow_tx_full_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow tx full drop count",
						},
						"hw_fwd_flow_qdr_full_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward flow qdr full drop count",
						},
						"hw_fwd_phyport_mismatch_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward phyport mismatch count",
						},
						"hw_fwd_vlanid_mismatch_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward vlanid mismatch count",
						},
						"hw_fwd_vmid_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward vmid mismatch count",
						},
						"hw_fwd_protocol_mismatch_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward protocol mismatch count",
						},
						"hw_fwd_avail_ipv4_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward available ipv4 entries count",
						},
						"hw_fwd_avail_ipv6_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward available ipv6 entries count",
						},
						"hw_fwd_rate_drop_count": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward rate drop count",
						},
						"hw_fwd_normal_ageout_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward normal ageout received count",
						},
						"hw_fwd_tcp_fin_ageout_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward tcp FIN ageout received count",
						},
						"hw_fwd_tcp_rst_ageout_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward tcp RST ageout received count",
						},
						"hw_fwd_lookup_fail_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward entry lookup fail count",
						},
						"hw_fwd_stats_update_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Hardware forward entry stats update count",
						},
						"hw_fwd_flow_sflow_count": {
							Type: schema.TypeInt, Optional: true, Description: "hardware forward rate drop count",
						},
					},
				},
			},
		},
	}
}

func resourceSystemHardwareAccelerateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemHardwareAccelerateStatsSlb := setObjectSystemHardwareAccelerateStatsSlb(res)
		d.Set("slb", SystemHardwareAccelerateStatsSlb)
		SystemHardwareAccelerateStatsStats := setObjectSystemHardwareAccelerateStatsStats(res)
		d.Set("stats", SystemHardwareAccelerateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemHardwareAccelerateStatsSlb(ret edpt.DataSystemHardwareAccelerateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectSystemHardwareAccelerateStatsSlbStats(ret.DtSystemHardwareAccelerateStats.Slb.Stats),
		},
	}
}

func setObjectSystemHardwareAccelerateStatsSlbStats(d edpt.SystemHardwareAccelerateStatsSlbStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["entry_create"] = d.EntryCreate

	in["entry_create_failure"] = d.EntryCreateFailure

	in["entry_create_fail_server_down"] = d.EntryCreateFailServerDown

	in["entry_create_fail_max_entry"] = d.EntryCreateFailMaxEntry

	in["entry_free"] = d.EntryFree

	in["entry_free_opp_entry"] = d.EntryFreeOppEntry

	in["entry_free_no_hw_prog"] = d.EntryFreeNoHwProg

	in["entry_free_no_conn"] = d.EntryFreeNoConn

	in["entry_free_no_sw_entry"] = d.EntryFreeNoSwEntry

	in["entry_counter"] = d.EntryCounter

	in["entry_age_out"] = d.EntryAgeOut

	in["entry_age_out_idle"] = d.EntryAgeOutIdle

	in["entry_age_out_tcp_fin"] = d.EntryAgeOutTcpFin

	in["entry_age_out_tcp_rst"] = d.EntryAgeOutTcpRst

	in["entry_age_out_invalid_dst"] = d.EntryAgeOutInvalidDst

	in["entry_force_hw_invalidate"] = d.EntryForceHwInvalidate

	in["entry_invalidate_server_down"] = d.EntryInvalidateServerDown

	in["tcam_create"] = d.TcamCreate

	in["tcam_free"] = d.TcamFree

	in["tcam_counter"] = d.TcamCounter
	result = append(result, in)
	return result
}

func setObjectSystemHardwareAccelerateStatsStats(ret edpt.DataSystemHardwareAccelerateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hit_counts":                    ret.DtSystemHardwareAccelerateStats.Stats.HitCounts,
			"hit_index":                     ret.DtSystemHardwareAccelerateStats.Stats.HitIndex,
			"ipv4_forward_counts":           ret.DtSystemHardwareAccelerateStats.Stats.Ipv4ForwardCounts,
			"ipv6_forward_counts":           ret.DtSystemHardwareAccelerateStats.Stats.Ipv6ForwardCounts,
			"hw_fwd_module_status":          ret.DtSystemHardwareAccelerateStats.Stats.HwFwdModuleStatus,
			"hw_fwd_prog_reqs":              ret.DtSystemHardwareAccelerateStats.Stats.HwFwdProgReqs,
			"hw_fwd_prog_errors":            ret.DtSystemHardwareAccelerateStats.Stats.HwFwdProgErrors,
			"hw_fwd_flow_singlebit_errors":  ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowSinglebitErrors,
			"hw_fwd_flow_tag_mismatch":      ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowTagMismatch,
			"hw_fwd_flow_seq_mismatch":      ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowSeqMismatch,
			"hw_fwd_ageout_drop_count":      ret.DtSystemHardwareAccelerateStats.Stats.HwFwdAgeoutDropCount,
			"hw_fwd_invalidation_drop":      ret.DtSystemHardwareAccelerateStats.Stats.HwFwdInvalidationDrop,
			"hw_fwd_flow_hit_index":         ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowHitIndex,
			"hw_fwd_flow_reason_flags":      ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowReasonFlags,
			"hw_fwd_flow_drop_count":        ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowDropCount,
			"hw_fwd_flow_error_count":       ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowErrorCount,
			"hw_fwd_flow_unalign_count":     ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowUnalignCount,
			"hw_fwd_flow_underflow_count":   ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowUnderflowCount,
			"hw_fwd_flow_tx_full_drop":      ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowTxFullDrop,
			"hw_fwd_flow_qdr_full_drop":     ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowQdrFullDrop,
			"hw_fwd_phyport_mismatch_drop":  ret.DtSystemHardwareAccelerateStats.Stats.HwFwdPhyportMismatchDrop,
			"hw_fwd_vlanid_mismatch_drop":   ret.DtSystemHardwareAccelerateStats.Stats.HwFwdVlanidMismatchDrop,
			"hw_fwd_vmid_drop":              ret.DtSystemHardwareAccelerateStats.Stats.HwFwdVmidDrop,
			"hw_fwd_protocol_mismatch_drop": ret.DtSystemHardwareAccelerateStats.Stats.HwFwdProtocolMismatchDrop,
			"hw_fwd_avail_ipv4_entry":       ret.DtSystemHardwareAccelerateStats.Stats.HwFwdAvailIpv4Entry,
			"hw_fwd_avail_ipv6_entry":       ret.DtSystemHardwareAccelerateStats.Stats.HwFwdAvailIpv6Entry,
			"hw_fwd_rate_drop_count":        ret.DtSystemHardwareAccelerateStats.Stats.HwFwdRateDropCount,
			"hw_fwd_normal_ageout_rcvd":     ret.DtSystemHardwareAccelerateStats.Stats.HwFwdNormalAgeoutRcvd,
			"hw_fwd_tcp_fin_ageout_rcvd":    ret.DtSystemHardwareAccelerateStats.Stats.HwFwdTcpFinAgeoutRcvd,
			"hw_fwd_tcp_rst_ageout_rcvd":    ret.DtSystemHardwareAccelerateStats.Stats.HwFwdTcpRstAgeoutRcvd,
			"hw_fwd_lookup_fail_rcvd":       ret.DtSystemHardwareAccelerateStats.Stats.HwFwdLookupFailRcvd,
			"hw_fwd_stats_update_rcvd":      ret.DtSystemHardwareAccelerateStats.Stats.HwFwdStatsUpdateRcvd,
			"hw_fwd_flow_sflow_count":       ret.DtSystemHardwareAccelerateStats.Stats.HwFwdFlowSflowCount,
		},
	}
}

func getObjectSystemHardwareAccelerateStatsSlb(d []interface{}) edpt.SystemHardwareAccelerateStatsSlb {

	count1 := len(d)
	var ret edpt.SystemHardwareAccelerateStatsSlb
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectSystemHardwareAccelerateStatsSlbStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectSystemHardwareAccelerateStatsSlbStats(d []interface{}) edpt.SystemHardwareAccelerateStatsSlbStats {

	count1 := len(d)
	var ret edpt.SystemHardwareAccelerateStatsSlbStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryCreate = in["entry_create"].(int)
		ret.EntryCreateFailure = in["entry_create_failure"].(int)
		ret.EntryCreateFailServerDown = in["entry_create_fail_server_down"].(int)
		ret.EntryCreateFailMaxEntry = in["entry_create_fail_max_entry"].(int)
		ret.EntryFree = in["entry_free"].(int)
		ret.EntryFreeOppEntry = in["entry_free_opp_entry"].(int)
		ret.EntryFreeNoHwProg = in["entry_free_no_hw_prog"].(int)
		ret.EntryFreeNoConn = in["entry_free_no_conn"].(int)
		ret.EntryFreeNoSwEntry = in["entry_free_no_sw_entry"].(int)
		ret.EntryCounter = in["entry_counter"].(int)
		ret.EntryAgeOut = in["entry_age_out"].(int)
		ret.EntryAgeOutIdle = in["entry_age_out_idle"].(int)
		ret.EntryAgeOutTcpFin = in["entry_age_out_tcp_fin"].(int)
		ret.EntryAgeOutTcpRst = in["entry_age_out_tcp_rst"].(int)
		ret.EntryAgeOutInvalidDst = in["entry_age_out_invalid_dst"].(int)
		ret.EntryForceHwInvalidate = in["entry_force_hw_invalidate"].(int)
		ret.EntryInvalidateServerDown = in["entry_invalidate_server_down"].(int)
		ret.TcamCreate = in["tcam_create"].(int)
		ret.TcamFree = in["tcam_free"].(int)
		ret.TcamCounter = in["tcam_counter"].(int)
	}
	return ret
}

func getObjectSystemHardwareAccelerateStatsStats(d []interface{}) edpt.SystemHardwareAccelerateStatsStats {

	count1 := len(d)
	var ret edpt.SystemHardwareAccelerateStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HitCounts = in["hit_counts"].(int)
		ret.HitIndex = in["hit_index"].(int)
		ret.Ipv4ForwardCounts = in["ipv4_forward_counts"].(int)
		ret.Ipv6ForwardCounts = in["ipv6_forward_counts"].(int)
		ret.HwFwdModuleStatus = in["hw_fwd_module_status"].(int)
		ret.HwFwdProgReqs = in["hw_fwd_prog_reqs"].(int)
		ret.HwFwdProgErrors = in["hw_fwd_prog_errors"].(int)
		ret.HwFwdFlowSinglebitErrors = in["hw_fwd_flow_singlebit_errors"].(int)
		ret.HwFwdFlowTagMismatch = in["hw_fwd_flow_tag_mismatch"].(int)
		ret.HwFwdFlowSeqMismatch = in["hw_fwd_flow_seq_mismatch"].(int)
		ret.HwFwdAgeoutDropCount = in["hw_fwd_ageout_drop_count"].(int)
		ret.HwFwdInvalidationDrop = in["hw_fwd_invalidation_drop"].(int)
		ret.HwFwdFlowHitIndex = in["hw_fwd_flow_hit_index"].(int)
		ret.HwFwdFlowReasonFlags = in["hw_fwd_flow_reason_flags"].(int)
		ret.HwFwdFlowDropCount = in["hw_fwd_flow_drop_count"].(int)
		ret.HwFwdFlowErrorCount = in["hw_fwd_flow_error_count"].(int)
		ret.HwFwdFlowUnalignCount = in["hw_fwd_flow_unalign_count"].(int)
		ret.HwFwdFlowUnderflowCount = in["hw_fwd_flow_underflow_count"].(int)
		ret.HwFwdFlowTxFullDrop = in["hw_fwd_flow_tx_full_drop"].(int)
		ret.HwFwdFlowQdrFullDrop = in["hw_fwd_flow_qdr_full_drop"].(int)
		ret.HwFwdPhyportMismatchDrop = in["hw_fwd_phyport_mismatch_drop"].(int)
		ret.HwFwdVlanidMismatchDrop = in["hw_fwd_vlanid_mismatch_drop"].(int)
		ret.HwFwdVmidDrop = in["hw_fwd_vmid_drop"].(int)
		ret.HwFwdProtocolMismatchDrop = in["hw_fwd_protocol_mismatch_drop"].(int)
		ret.HwFwdAvailIpv4Entry = in["hw_fwd_avail_ipv4_entry"].(int)
		ret.HwFwdAvailIpv6Entry = in["hw_fwd_avail_ipv6_entry"].(int)
		ret.HwFwdRateDropCount = in["hw_fwd_rate_drop_count"].(int)
		ret.HwFwdNormalAgeoutRcvd = in["hw_fwd_normal_ageout_rcvd"].(int)
		ret.HwFwdTcpFinAgeoutRcvd = in["hw_fwd_tcp_fin_ageout_rcvd"].(int)
		ret.HwFwdTcpRstAgeoutRcvd = in["hw_fwd_tcp_rst_ageout_rcvd"].(int)
		ret.HwFwdLookupFailRcvd = in["hw_fwd_lookup_fail_rcvd"].(int)
		ret.HwFwdStatsUpdateRcvd = in["hw_fwd_stats_update_rcvd"].(int)
		ret.HwFwdFlowSflowCount = in["hw_fwd_flow_sflow_count"].(int)
	}
	return ret
}

func dataToEndpointSystemHardwareAccelerateStats(d *schema.ResourceData) edpt.SystemHardwareAccelerateStats {
	var ret edpt.SystemHardwareAccelerateStats

	ret.Slb = getObjectSystemHardwareAccelerateStatsSlb(d.Get("slb").([]interface{}))

	ret.Stats = getObjectSystemHardwareAccelerateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
