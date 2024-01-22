package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_number_src_based_policy_policy_class_list_stats`: Statistics for the object policy-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsRead,

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
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats := setObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats(res)
		d.Set("stats", DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats(ret edpt.DataDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_received":      ret.DtDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats.Stats.Packet_received,
			"packet_dropped":       ret.DtDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats.Stats.Packet_dropped,
			"entry_learned":        ret.DtDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats.Stats.Entry_learned,
			"entry_count_overflow": ret.DtDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats.Stats.Entry_count_overflow,
		},
	}
}

func getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_received = in["packet_received"].(int)
		ret.Packet_dropped = in["packet_dropped"].(int)
		ret.Entry_learned = in["entry_learned"].(int)
		ret.Entry_count_overflow = in["entry_count_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats {
	var ret edpt.DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats

	ret.ClassListName = d.Get("class_list_name").(string)

	ret.Stats = getObjectDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats(d.Get("stats").([]interface{}))

	ret.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
