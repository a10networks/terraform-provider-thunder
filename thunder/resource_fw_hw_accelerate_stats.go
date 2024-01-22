package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwHwAccelerateStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_hw_accelerate_stats`: Statistics for the object hw-accelerate\n\n__PLACEHOLDER__",
		ReadContext: resourceFwHwAccelerateStatsRead,

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

func resourceFwHwAccelerateStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwHwAccelerateStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwHwAccelerateStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwHwAccelerateStatsStats := setObjectFwHwAccelerateStatsStats(res)
		d.Set("stats", FwHwAccelerateStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwHwAccelerateStatsStats(ret edpt.DataFwHwAccelerateStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_create":                  ret.DtFwHwAccelerateStats.Stats.EntryCreate,
			"entry_create_failure":          ret.DtFwHwAccelerateStats.Stats.EntryCreateFailure,
			"entry_create_fail_server_down": ret.DtFwHwAccelerateStats.Stats.EntryCreateFailServerDown,
			"entry_create_fail_max_entry":   ret.DtFwHwAccelerateStats.Stats.EntryCreateFailMaxEntry,
			"entry_free":                    ret.DtFwHwAccelerateStats.Stats.EntryFree,
			"entry_free_opp_entry":          ret.DtFwHwAccelerateStats.Stats.EntryFreeOppEntry,
			"entry_free_no_hw_prog":         ret.DtFwHwAccelerateStats.Stats.EntryFreeNoHwProg,
			"entry_free_no_conn":            ret.DtFwHwAccelerateStats.Stats.EntryFreeNoConn,
			"entry_free_no_sw_entry":        ret.DtFwHwAccelerateStats.Stats.EntryFreeNoSwEntry,
			"entry_counter":                 ret.DtFwHwAccelerateStats.Stats.EntryCounter,
			"entry_age_out":                 ret.DtFwHwAccelerateStats.Stats.EntryAgeOut,
			"entry_age_out_idle":            ret.DtFwHwAccelerateStats.Stats.EntryAgeOutIdle,
			"entry_age_out_tcp_fin":         ret.DtFwHwAccelerateStats.Stats.EntryAgeOutTcpFin,
			"entry_age_out_tcp_rst":         ret.DtFwHwAccelerateStats.Stats.EntryAgeOutTcpRst,
			"entry_age_out_invalid_dst":     ret.DtFwHwAccelerateStats.Stats.EntryAgeOutInvalidDst,
			"entry_force_hw_invalidate":     ret.DtFwHwAccelerateStats.Stats.EntryForceHwInvalidate,
			"entry_invalidate_server_down":  ret.DtFwHwAccelerateStats.Stats.EntryInvalidateServerDown,
			"tcam_create":                   ret.DtFwHwAccelerateStats.Stats.TcamCreate,
			"tcam_free":                     ret.DtFwHwAccelerateStats.Stats.TcamFree,
			"tcam_counter":                  ret.DtFwHwAccelerateStats.Stats.TcamCounter,
		},
	}
}

func getObjectFwHwAccelerateStatsStats(d []interface{}) edpt.FwHwAccelerateStatsStats {

	count1 := len(d)
	var ret edpt.FwHwAccelerateStatsStats
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

func dataToEndpointFwHwAccelerateStats(d *schema.ResourceData) edpt.FwHwAccelerateStats {
	var ret edpt.FwHwAccelerateStats

	ret.Stats = getObjectFwHwAccelerateStatsStats(d.Get("stats").([]interface{}))
	return ret
}
