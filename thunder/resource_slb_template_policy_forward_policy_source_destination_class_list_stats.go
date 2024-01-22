package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_source_destination_class_list_stats`: Statistics for the object class-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListStatsRead,

		Schema: map[string]*schema.Schema{
			"dest_class_list": {
				Type: schema.TypeString, Required: true, Description: "Destination Class List Name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests matching this destination rule",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationClassListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationClassListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats := setObjectSlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicySourceDestinationClassListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtSlbTemplatePolicyForwardPolicySourceDestinationClassListStats.Stats.Hits,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationClassListStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassListStats {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationClassListStats

	ret.DestClassList = d.Get("dest_class_list").(string)

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicySourceDestinationClassListStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
