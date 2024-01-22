package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_src_based_policy_policy_class_list_stats`: Statistics for the object policy-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsRead,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats := setObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats(res)
		d.Set("stats", DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats(ret edpt.DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_received":      ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats.Stats.Packet_received,
			"packet_dropped":       ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats.Stats.Packet_dropped,
			"entry_learned":        ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats.Stats.Entry_learned,
			"entry_count_overflow": ret.DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats.Stats.Entry_count_overflow,
		},
	}
}

func getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats(d []interface{}) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_received = in["packet_received"].(int)
		ret.Packet_dropped = in["packet_dropped"].(int)
		ret.Entry_learned = in["entry_learned"].(int)
		ret.Entry_count_overflow = in["entry_count_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats {
	var ret edpt.DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats

	ret.ClassListName = d.Get("class_list_name").(string)

	ret.Stats = getObjectDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats(d.Get("stats").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)

	ret.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)

	ret.PortNum = d.Get("port_num").(string)
	return ret
}
