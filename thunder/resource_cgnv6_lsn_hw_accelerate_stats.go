package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnHwAccelerateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_hw_accelerate_stats`: Statistics for the object hw-accelerate\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnHwAccelerateStatsRead,

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

func resourceCgnv6LsnHwAccelerateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnHwAccelerateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnHwAccelerateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnHwAccelerateStatsStats := setObjectCgnv6LsnHwAccelerateStatsStats(res)
		d.Set("stats", Cgnv6LsnHwAccelerateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnHwAccelerateStatsStats(ret edpt.DataCgnv6LsnHwAccelerateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_create":                  ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryCreate,
			"entry_create_failure":          ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryCreateFailure,
			"entry_create_fail_server_down": ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryCreateFailServerDown,
			"entry_create_fail_max_entry":   ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryCreateFailMaxEntry,
			"entry_free":                    ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryFree,
			"entry_free_opp_entry":          ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryFreeOppEntry,
			"entry_free_no_hw_prog":         ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryFreeNoHwProg,
			"entry_free_no_conn":            ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryFreeNoConn,
			"entry_free_no_sw_entry":        ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryFreeNoSwEntry,
			"entry_counter":                 ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryCounter,
			"entry_age_out":                 ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryAgeOut,
			"entry_age_out_idle":            ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryAgeOutIdle,
			"entry_age_out_tcp_fin":         ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryAgeOutTcpFin,
			"entry_age_out_tcp_rst":         ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryAgeOutTcpRst,
			"entry_age_out_invalid_dst":     ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryAgeOutInvalidDst,
			"entry_force_hw_invalidate":     ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryForceHwInvalidate,
			"entry_invalidate_server_down":  ret.DtCgnv6LsnHwAccelerateStats.Stats.EntryInvalidateServerDown,
			"tcam_create":                   ret.DtCgnv6LsnHwAccelerateStats.Stats.TcamCreate,
			"tcam_free":                     ret.DtCgnv6LsnHwAccelerateStats.Stats.TcamFree,
			"tcam_counter":                  ret.DtCgnv6LsnHwAccelerateStats.Stats.TcamCounter,
		},
	}
}

func getObjectCgnv6LsnHwAccelerateStatsStats(d []interface{}) edpt.Cgnv6LsnHwAccelerateStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnHwAccelerateStatsStats
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

func dataToEndpointCgnv6LsnHwAccelerateStats(d *schema.ResourceData) edpt.Cgnv6LsnHwAccelerateStats {
	var ret edpt.Cgnv6LsnHwAccelerateStats

	ret.Stats = getObjectCgnv6LsnHwAccelerateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
