package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwPerInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_per_instance_stats`: Statistics for the object per-instance\n\n__PLACEHOLDER__",
		ReadContext: resourceFwPerInstanceStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Created",
						},
						"data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Freed",
						},
						"tcp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Created",
						},
						"tcp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Freed",
						},
						"udp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Created",
						},
						"udp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Freed",
						},
						"dyn_blist_sess_created": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist Session Created",
						},
						"dyn_blist_sess_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist Session Freed",
						},
						"dyn_blist_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist Packet Drop",
						},
						"dyn_blist_version_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist - Version Mismatch",
						},
						"dyn_blist_pkt_rate_low": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist - Pkt Rate Low",
						},
						"dyn_blist_pkt_rate_high": {
							Type: schema.TypeInt, Optional: true, Description: "Dynamic Blacklist - Pkt Rate High",
						},
					},
				},
			},
		},
	}
}

func resourceFwPerInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwPerInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwPerInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwPerInstanceStatsStats := setObjectFwPerInstanceStatsStats(res)
		d.Set("stats", FwPerInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwPerInstanceStatsStats(ret edpt.DataFwPerInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"key_name":                   ret.DtFwPerInstanceStats.Stats.KeyName,
			"data_session_created":       ret.DtFwPerInstanceStats.Stats.Data_session_created,
			"data_session_freed":         ret.DtFwPerInstanceStats.Stats.Data_session_freed,
			"tcp_fullcone_created":       ret.DtFwPerInstanceStats.Stats.Tcp_fullcone_created,
			"tcp_fullcone_freed":         ret.DtFwPerInstanceStats.Stats.Tcp_fullcone_freed,
			"udp_fullcone_created":       ret.DtFwPerInstanceStats.Stats.Udp_fullcone_created,
			"udp_fullcone_freed":         ret.DtFwPerInstanceStats.Stats.Udp_fullcone_freed,
			"dyn_blist_sess_created":     ret.DtFwPerInstanceStats.Stats.Dyn_blist_sess_created,
			"dyn_blist_sess_freed":       ret.DtFwPerInstanceStats.Stats.Dyn_blist_sess_freed,
			"dyn_blist_pkt_drop":         ret.DtFwPerInstanceStats.Stats.Dyn_blist_pkt_drop,
			"dyn_blist_version_mismatch": ret.DtFwPerInstanceStats.Stats.Dyn_blist_version_mismatch,
			"dyn_blist_pkt_rate_low":     ret.DtFwPerInstanceStats.Stats.Dyn_blist_pkt_rate_low,
			"dyn_blist_pkt_rate_high":    ret.DtFwPerInstanceStats.Stats.Dyn_blist_pkt_rate_high,
		},
	}
}

func getObjectFwPerInstanceStatsStats(d []interface{}) edpt.FwPerInstanceStatsStats {

	count1 := len(d)
	var ret edpt.FwPerInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyName = in["key_name"].(string)
		ret.Data_session_created = in["data_session_created"].(int)
		ret.Data_session_freed = in["data_session_freed"].(int)
		ret.Tcp_fullcone_created = in["tcp_fullcone_created"].(int)
		ret.Tcp_fullcone_freed = in["tcp_fullcone_freed"].(int)
		ret.Udp_fullcone_created = in["udp_fullcone_created"].(int)
		ret.Udp_fullcone_freed = in["udp_fullcone_freed"].(int)
		ret.Dyn_blist_sess_created = in["dyn_blist_sess_created"].(int)
		ret.Dyn_blist_sess_freed = in["dyn_blist_sess_freed"].(int)
		ret.Dyn_blist_pkt_drop = in["dyn_blist_pkt_drop"].(int)
		ret.Dyn_blist_version_mismatch = in["dyn_blist_version_mismatch"].(int)
		ret.Dyn_blist_pkt_rate_low = in["dyn_blist_pkt_rate_low"].(int)
		ret.Dyn_blist_pkt_rate_high = in["dyn_blist_pkt_rate_high"].(int)
	}
	return ret
}

func dataToEndpointFwPerInstanceStats(d *schema.ResourceData) edpt.FwPerInstanceStats {
	var ret edpt.FwPerInstanceStats

	ret.Stats = getObjectFwPerInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
