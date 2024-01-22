package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListDomainNameStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_rule_list_domain_name_stats`: Statistics for the object domain-name\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRuleListDomainNameStatsRead,

		Schema: map[string]*schema.Schema{
			"name_domain": {
				Type: schema.TypeString, Required: true, Description: "Configure a Specific Rule-Set (Domain Name)",
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

func resourceCgnv6LsnRuleListDomainNameStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDomainNameStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDomainNameStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRuleListDomainNameStatsStats := setObjectCgnv6LsnRuleListDomainNameStatsStats(res)
		d.Set("stats", Cgnv6LsnRuleListDomainNameStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRuleListDomainNameStatsStats(ret edpt.DataCgnv6LsnRuleListDomainNameStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectCgnv6LsnRuleListDomainNameStatsStats(d []interface{}) edpt.Cgnv6LsnRuleListDomainNameStatsStats {

	var ret edpt.Cgnv6LsnRuleListDomainNameStatsStats
	return ret
}

func dataToEndpointCgnv6LsnRuleListDomainNameStats(d *schema.ResourceData) edpt.Cgnv6LsnRuleListDomainNameStats {
	var ret edpt.Cgnv6LsnRuleListDomainNameStats

	ret.NameDomain = d.Get("name_domain").(string)

	ret.Stats = getObjectCgnv6LsnRuleListDomainNameStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
