package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatPerInstanceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_per_instance_stats`: Statistics for the object per-instance\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatPerInstanceStatsRead,

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
						"total_nat_in_use": {
							Type: schema.TypeInt, Optional: true, Description: "Total NAT addresses in use",
						},
						"active_subscriber_added": {
							Type: schema.TypeInt, Optional: true, Description: "Active Subscriber Added",
						},
						"active_subscriber_removed": {
							Type: schema.TypeInt, Optional: true, Description: "Active Subscriber Removed",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatPerInstanceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatPerInstanceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatPerInstanceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatPerInstanceStatsStats := setObjectCgnv6FixedNatPerInstanceStatsStats(res)
		d.Set("stats", Cgnv6FixedNatPerInstanceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatPerInstanceStatsStats(ret edpt.DataCgnv6FixedNatPerInstanceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"key_name":                  ret.DtCgnv6FixedNatPerInstanceStats.Stats.KeyName,
			"data_session_created":      ret.DtCgnv6FixedNatPerInstanceStats.Stats.Data_session_created,
			"data_session_freed":        ret.DtCgnv6FixedNatPerInstanceStats.Stats.Data_session_freed,
			"tcp_fullcone_created":      ret.DtCgnv6FixedNatPerInstanceStats.Stats.Tcp_fullcone_created,
			"tcp_fullcone_freed":        ret.DtCgnv6FixedNatPerInstanceStats.Stats.Tcp_fullcone_freed,
			"udp_fullcone_created":      ret.DtCgnv6FixedNatPerInstanceStats.Stats.Udp_fullcone_created,
			"udp_fullcone_freed":        ret.DtCgnv6FixedNatPerInstanceStats.Stats.Udp_fullcone_freed,
			"tcp_allocated":             ret.DtCgnv6FixedNatPerInstanceStats.Stats.Tcp_allocated,
			"tcp_freed":                 ret.DtCgnv6FixedNatPerInstanceStats.Stats.Tcp_freed,
			"udp_allocated":             ret.DtCgnv6FixedNatPerInstanceStats.Stats.Udp_allocated,
			"udp_freed":                 ret.DtCgnv6FixedNatPerInstanceStats.Stats.Udp_freed,
			"icmp_allocated":            ret.DtCgnv6FixedNatPerInstanceStats.Stats.Icmp_allocated,
			"icmp_freed":                ret.DtCgnv6FixedNatPerInstanceStats.Stats.Icmp_freed,
			"total_nat_in_use":          ret.DtCgnv6FixedNatPerInstanceStats.Stats.Total_nat_in_use,
			"active_subscriber_added":   ret.DtCgnv6FixedNatPerInstanceStats.Stats.ActiveSubscriberAdded,
			"active_subscriber_removed": ret.DtCgnv6FixedNatPerInstanceStats.Stats.ActiveSubscriberRemoved,
		},
	}
}

func getObjectCgnv6FixedNatPerInstanceStatsStats(d []interface{}) edpt.Cgnv6FixedNatPerInstanceStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatPerInstanceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyName = in["key_name"].(string)
		ret.Data_session_created = in["data_session_created"].(int)
		ret.Data_session_freed = in["data_session_freed"].(int)
		ret.Tcp_fullcone_created = in["tcp_fullcone_created"].(int)
		ret.Tcp_fullcone_freed = in["tcp_fullcone_freed"].(int)
		ret.Udp_fullcone_created = in["udp_fullcone_created"].(int)
		ret.Udp_fullcone_freed = in["udp_fullcone_freed"].(int)
		ret.Tcp_allocated = in["tcp_allocated"].(int)
		ret.Tcp_freed = in["tcp_freed"].(int)
		ret.Udp_allocated = in["udp_allocated"].(int)
		ret.Udp_freed = in["udp_freed"].(int)
		ret.Icmp_allocated = in["icmp_allocated"].(int)
		ret.Icmp_freed = in["icmp_freed"].(int)
		ret.Total_nat_in_use = in["total_nat_in_use"].(int)
		ret.ActiveSubscriberAdded = in["active_subscriber_added"].(int)
		ret.ActiveSubscriberRemoved = in["active_subscriber_removed"].(int)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatPerInstanceStats(d *schema.ResourceData) edpt.Cgnv6FixedNatPerInstanceStats {
	var ret edpt.Cgnv6FixedNatPerInstanceStats

	ret.Stats = getObjectCgnv6FixedNatPerInstanceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
