package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHardwareAccelerateSlbStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_hardware_accelerate_slb_stats`: Statistics for the object slb\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemHardwareAccelerateSlbStatsRead,

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
	}
}

func resourceSystemHardwareAccelerateSlbStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareAccelerateSlbStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareAccelerateSlbStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemHardwareAccelerateSlbStatsStats := setObjectSystemHardwareAccelerateSlbStatsStats(res)
		d.Set("stats", SystemHardwareAccelerateSlbStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemHardwareAccelerateSlbStatsStats(ret edpt.DataSystemHardwareAccelerateSlbStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_create":                  ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryCreate,
			"entry_create_failure":          ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryCreateFailure,
			"entry_create_fail_server_down": ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryCreateFailServerDown,
			"entry_create_fail_max_entry":   ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryCreateFailMaxEntry,
			"entry_free":                    ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryFree,
			"entry_free_opp_entry":          ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryFreeOppEntry,
			"entry_free_no_hw_prog":         ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryFreeNoHwProg,
			"entry_free_no_conn":            ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryFreeNoConn,
			"entry_free_no_sw_entry":        ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryFreeNoSwEntry,
			"entry_counter":                 ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryCounter,
			"entry_age_out":                 ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryAgeOut,
			"entry_age_out_idle":            ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryAgeOutIdle,
			"entry_age_out_tcp_fin":         ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryAgeOutTcpFin,
			"entry_age_out_tcp_rst":         ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryAgeOutTcpRst,
			"entry_age_out_invalid_dst":     ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryAgeOutInvalidDst,
			"entry_force_hw_invalidate":     ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryForceHwInvalidate,
			"entry_invalidate_server_down":  ret.DtSystemHardwareAccelerateSlbStats.Stats.EntryInvalidateServerDown,
			"tcam_create":                   ret.DtSystemHardwareAccelerateSlbStats.Stats.TcamCreate,
			"tcam_free":                     ret.DtSystemHardwareAccelerateSlbStats.Stats.TcamFree,
			"tcam_counter":                  ret.DtSystemHardwareAccelerateSlbStats.Stats.TcamCounter,
		},
	}
}

func getObjectSystemHardwareAccelerateSlbStatsStats(d []interface{}) edpt.SystemHardwareAccelerateSlbStatsStats {

	count1 := len(d)
	var ret edpt.SystemHardwareAccelerateSlbStatsStats
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

func dataToEndpointSystemHardwareAccelerateSlbStats(d *schema.ResourceData) edpt.SystemHardwareAccelerateSlbStats {
	var ret edpt.SystemHardwareAccelerateSlbStats

	ret.Stats = getObjectSystemHardwareAccelerateSlbStatsStats(d.Get("stats").([]interface{}))
	return ret
}
