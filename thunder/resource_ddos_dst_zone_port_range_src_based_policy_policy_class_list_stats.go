package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_stats`: Statistics for the object policy-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsRead,

		Schema: map[string]*schema.Schema{
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_received": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Received",
						},
						"packet_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Dropped",
						},
						"entry_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Learned",
						},
						"entry_count_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Count Overflow",
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
		},
	}
}

func resourceDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats := setObjectDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats(res)
		d.Set("stats", DdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats(ret edpt.DataDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_received":      ret.DtDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats.Stats.Packet_received,
			"packet_dropped":       ret.DtDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats.Stats.Packet_dropped,
			"entry_learned":        ret.DtDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats.Stats.Entry_learned,
			"entry_count_overflow": ret.DtDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats.Stats.Entry_count_overflow,
		},
	}
}

func getObjectDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats(d []interface{}) edpt.DdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_received = in["packet_received"].(int)
		ret.Packet_dropped = in["packet_dropped"].(int)
		ret.Entry_learned = in["entry_learned"].(int)
		ret.Entry_count_overflow = in["entry_count_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats(d *schema.ResourceData) edpt.DdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats {
	var ret edpt.DdosDstZonePortRangeSrcBasedPolicyPolicyClassListStats

	ret.ClassListName = d.Get("class_list_name").(string)

	ret.Stats = getObjectDdosDstZonePortRangeSrcBasedPolicyPolicyClassListStatsStats(d.Get("stats").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)

	ret.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)
	return ret
}
