package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_stats`: Statistics for the object policy-class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListStatsRead,

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
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats := setObjectDdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats(res)
		d.Set("stats", DdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats(ret edpt.DataDdosDstEntrySrcDstPairPolicyPolicyClassListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"packet_received":      ret.DtDdosDstEntrySrcDstPairPolicyPolicyClassListStats.Stats.Packet_received,
			"packet_dropped":       ret.DtDdosDstEntrySrcDstPairPolicyPolicyClassListStats.Stats.Packet_dropped,
			"entry_learned":        ret.DtDdosDstEntrySrcDstPairPolicyPolicyClassListStats.Stats.Entry_learned,
			"entry_count_overflow": ret.DtDdosDstEntrySrcDstPairPolicyPolicyClassListStats.Stats.Entry_count_overflow,
		},
	}
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Packet_received = in["packet_received"].(int)
		ret.Packet_dropped = in["packet_dropped"].(int)
		ret.Entry_learned = in["entry_learned"].(int)
		ret.Entry_count_overflow = in["entry_count_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListStats(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListStats {
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListStats

	ret.ClassListName = d.Get("class_list_name").(string)

	ret.Stats = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats(d.Get("stats").([]interface{}))

	ret.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
