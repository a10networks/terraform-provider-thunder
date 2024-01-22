package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_src_based_policy_policy_class_list_stats`: Statistics for the object policy-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsRead,

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
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "PortOther",
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
		},
	}
}

func resourceDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats := setObjectDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats(res)
		d.Set("stats", DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats(ret edpt.DataDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_received":      ret.DtDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats.Stats.Packet_received,
			"packet_dropped":       ret.DtDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats.Stats.Packet_dropped,
			"entry_learned":        ret.DtDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats.Stats.Entry_learned,
			"entry_count_overflow": ret.DtDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats.Stats.Entry_count_overflow,
		},
	}
}

func getObjectDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_received = in["packet_received"].(int)
		ret.Packet_dropped = in["packet_dropped"].(int)
		ret.Entry_learned = in["entry_learned"].(int)
		ret.Entry_count_overflow = in["entry_count_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats {
	var ret edpt.DdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStats

	ret.ClassListName = d.Get("class_list_name").(string)

	ret.Stats = getObjectDdosDstZonePortZoneServiceOtherSrcBasedPolicyPolicyClassListStatsStats(d.Get("stats").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)

	ret.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	return ret
}
