package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListIpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_rule_list_ip_stats`: Statistics for the object ip\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRuleListIpStatsRead,

		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "Configure a Specific Rule-Set (IP Network Address)",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceCgnv6LsnRuleListIpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListIpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListIpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRuleListIpStatsStats := setObjectCgnv6LsnRuleListIpStatsStats(res)
		d.Set("stats", Cgnv6LsnRuleListIpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRuleListIpStatsStats(ret edpt.DataCgnv6LsnRuleListIpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectCgnv6LsnRuleListIpStatsStats(d []interface{}) edpt.Cgnv6LsnRuleListIpStatsStats {

	var ret edpt.Cgnv6LsnRuleListIpStatsStats
	return ret
}

func dataToEndpointCgnv6LsnRuleListIpStats(d *schema.ResourceData) edpt.Cgnv6LsnRuleListIpStats {
	var ret edpt.Cgnv6LsnRuleListIpStats

	ret.Ipv4Addr = d.Get("ipv4_addr").(string)

	ret.Stats = getObjectCgnv6LsnRuleListIpStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
