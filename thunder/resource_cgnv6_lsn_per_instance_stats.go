package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnPerInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_per_instance_stats`: Statistics for the object per-instance\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnPerInstanceStatsRead,

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
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Session Created",
						},
						"tcp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Session Freed",
						},
						"udp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Session Created",
						},
						"udp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Session Freed",
						},
						"user_quota_created": {
							Type: schema.TypeInt, Optional: true, Description: "User-Quota Created",
						},
						"user_quota_put_in_del_q": {
							Type: schema.TypeInt, Optional: true, Description: "User-Quota Freed",
						},
						"tcp_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Ports Allocated",
						},
						"tcp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Ports Freed",
						},
						"udp_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Ports Allocated",
						},
						"udp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Ports Freed",
						},
						"icmp_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Ports Allocated",
						},
						"icmp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Ports Freed",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6LsnPerInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnPerInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnPerInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnPerInstanceStatsStats := setObjectCgnv6LsnPerInstanceStatsStats(res)
		d.Set("stats", Cgnv6LsnPerInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnPerInstanceStatsStats(ret edpt.DataCgnv6LsnPerInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"key_name":                ret.DtCgnv6LsnPerInstanceStats.Stats.KeyName,
			"data_session_created":    ret.DtCgnv6LsnPerInstanceStats.Stats.Data_session_created,
			"data_session_freed":      ret.DtCgnv6LsnPerInstanceStats.Stats.Data_session_freed,
			"tcp_fullcone_created":    ret.DtCgnv6LsnPerInstanceStats.Stats.Tcp_fullcone_created,
			"tcp_fullcone_freed":      ret.DtCgnv6LsnPerInstanceStats.Stats.Tcp_fullcone_freed,
			"udp_fullcone_created":    ret.DtCgnv6LsnPerInstanceStats.Stats.Udp_fullcone_created,
			"udp_fullcone_freed":      ret.DtCgnv6LsnPerInstanceStats.Stats.Udp_fullcone_freed,
			"user_quota_created":      ret.DtCgnv6LsnPerInstanceStats.Stats.User_quota_created,
			"user_quota_put_in_del_q": ret.DtCgnv6LsnPerInstanceStats.Stats.User_quota_put_in_del_q,
			"tcp_allocated":           ret.DtCgnv6LsnPerInstanceStats.Stats.Tcp_allocated,
			"tcp_freed":               ret.DtCgnv6LsnPerInstanceStats.Stats.Tcp_freed,
			"udp_allocated":           ret.DtCgnv6LsnPerInstanceStats.Stats.Udp_allocated,
			"udp_freed":               ret.DtCgnv6LsnPerInstanceStats.Stats.Udp_freed,
			"icmp_allocated":          ret.DtCgnv6LsnPerInstanceStats.Stats.Icmp_allocated,
			"icmp_freed":              ret.DtCgnv6LsnPerInstanceStats.Stats.Icmp_freed,
		},
	}
}

func getObjectCgnv6LsnPerInstanceStatsStats(d []interface{}) edpt.Cgnv6LsnPerInstanceStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6LsnPerInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyName = in["key_name"].(string)
		ret.Data_session_created = in["data_session_created"].(int)
		ret.Data_session_freed = in["data_session_freed"].(int)
		ret.Tcp_fullcone_created = in["tcp_fullcone_created"].(int)
		ret.Tcp_fullcone_freed = in["tcp_fullcone_freed"].(int)
		ret.Udp_fullcone_created = in["udp_fullcone_created"].(int)
		ret.Udp_fullcone_freed = in["udp_fullcone_freed"].(int)
		ret.User_quota_created = in["user_quota_created"].(int)
		ret.User_quota_put_in_del_q = in["user_quota_put_in_del_q"].(int)
		ret.Tcp_allocated = in["tcp_allocated"].(int)
		ret.Tcp_freed = in["tcp_freed"].(int)
		ret.Udp_allocated = in["udp_allocated"].(int)
		ret.Udp_freed = in["udp_freed"].(int)
		ret.Icmp_allocated = in["icmp_allocated"].(int)
		ret.Icmp_freed = in["icmp_freed"].(int)
	}
	return ret
}

func dataToEndpointCgnv6LsnPerInstanceStats(d *schema.ResourceData) edpt.Cgnv6LsnPerInstanceStats {
	var ret edpt.Cgnv6LsnPerInstanceStats

	ret.Stats = getObjectCgnv6LsnPerInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
