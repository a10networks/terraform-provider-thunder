package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64HwAccelerateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat64_hw_accelerate_stats`: Statistics for the object hw-accelerate\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Nat64HwAccelerateStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_create": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Created",
						},
						"entry_create_failure": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entry Creation Failed",
						},
						"entry_create_fail_server_down": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entry Creation Failed - server down",
						},
						"entry_create_fail_max_entry": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entry Creation Failed - max entries exceeded",
						},
						"entry_free": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Freed",
						},
						"entry_free_opp_entry": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Freed - opposite tuple entry aged-out",
						},
						"entry_free_no_hw_prog": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entry Freed - no HW prog",
						},
						"entry_free_no_conn": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entry Freed - no matched conn",
						},
						"entry_free_no_sw_entry": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entry Freed - no software entry",
						},
						"entry_counter": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Count",
						},
						"entry_age_out": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Aged Out",
						},
						"entry_age_out_idle": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Aged Out - idle timeout",
						},
						"entry_age_out_tcp_fin": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Aged Out - TCP FIN",
						},
						"entry_age_out_tcp_rst": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Aged Out - TCP RST",
						},
						"entry_age_out_invalid_dst": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Aged Out - invalid dst",
						},
						"entry_force_hw_invalidate": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Force HW Invalidate",
						},
						"entry_invalidate_server_down": {
							Type: schema.TypeInt, Optional: true, Description: "HW Entries Invalidate due to server down",
						},
						"tcam_create": {
							Type: schema.TypeInt, Optional: true, Description: "TCAM Flows Created",
						},
						"tcam_free": {
							Type: schema.TypeInt, Optional: true, Description: "TCAM Flows Freed",
						},
						"tcam_counter": {
							Type: schema.TypeInt, Optional: true, Description: "TCAM Flow Count",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Nat64HwAccelerateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64HwAccelerateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64HwAccelerateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Nat64HwAccelerateStatsStats := setObjectCgnv6Nat64HwAccelerateStatsStats(res)
		d.Set("stats", Cgnv6Nat64HwAccelerateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Nat64HwAccelerateStatsStats(ret edpt.DataCgnv6Nat64HwAccelerateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_create":                  ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryCreate,
			"entry_create_failure":          ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryCreateFailure,
			"entry_create_fail_server_down": ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryCreateFailServerDown,
			"entry_create_fail_max_entry":   ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryCreateFailMaxEntry,
			"entry_free":                    ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryFree,
			"entry_free_opp_entry":          ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryFreeOppEntry,
			"entry_free_no_hw_prog":         ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryFreeNoHwProg,
			"entry_free_no_conn":            ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryFreeNoConn,
			"entry_free_no_sw_entry":        ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryFreeNoSwEntry,
			"entry_counter":                 ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryCounter,
			"entry_age_out":                 ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryAgeOut,
			"entry_age_out_idle":            ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryAgeOutIdle,
			"entry_age_out_tcp_fin":         ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryAgeOutTcpFin,
			"entry_age_out_tcp_rst":         ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryAgeOutTcpRst,
			"entry_age_out_invalid_dst":     ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryAgeOutInvalidDst,
			"entry_force_hw_invalidate":     ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryForceHwInvalidate,
			"entry_invalidate_server_down":  ret.DtCgnv6Nat64HwAccelerateStats.Stats.EntryInvalidateServerDown,
			"tcam_create":                   ret.DtCgnv6Nat64HwAccelerateStats.Stats.TcamCreate,
			"tcam_free":                     ret.DtCgnv6Nat64HwAccelerateStats.Stats.TcamFree,
			"tcam_counter":                  ret.DtCgnv6Nat64HwAccelerateStats.Stats.TcamCounter,
		},
	}
}

func getObjectCgnv6Nat64HwAccelerateStatsStats(d []interface{}) edpt.Cgnv6Nat64HwAccelerateStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Nat64HwAccelerateStatsStats
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

func dataToEndpointCgnv6Nat64HwAccelerateStats(d *schema.ResourceData) edpt.Cgnv6Nat64HwAccelerateStats {
	var ret edpt.Cgnv6Nat64HwAccelerateStats

	ret.Stats = getObjectCgnv6Nat64HwAccelerateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
