package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListDomainIpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_rule_list_domain_ip_stats`: Statistics for the object domain-ip\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRuleListDomainIpStatsRead,

		Schema: map[string]*schema.Schema{
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

func resourceCgnv6LsnRuleListDomainIpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDomainIpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDomainIpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRuleListDomainIpStatsStats := setObjectCgnv6LsnRuleListDomainIpStatsStats(res)
		d.Set("stats", Cgnv6LsnRuleListDomainIpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRuleListDomainIpStatsStats(ret edpt.DataCgnv6LsnRuleListDomainIpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectCgnv6LsnRuleListDomainIpStatsStats(d []interface{}) edpt.Cgnv6LsnRuleListDomainIpStatsStats {

	var ret edpt.Cgnv6LsnRuleListDomainIpStatsStats
	return ret
}

func dataToEndpointCgnv6LsnRuleListDomainIpStats(d *schema.ResourceData) edpt.Cgnv6LsnRuleListDomainIpStats {
	var ret edpt.Cgnv6LsnRuleListDomainIpStats

	ret.Stats = getObjectCgnv6LsnRuleListDomainIpStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
