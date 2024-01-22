package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListDomainListNameStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_lsn_rule_list_domain_list_name_stats`: Statistics for the object domain-list-name\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6LsnRuleListDomainListNameStatsRead,

		Schema: map[string]*schema.Schema{
			"name_domain_list": {
				Type: schema.TypeString, Required: true, Description: "Configure a Specific Rule-Set (Domain List Name)",
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

func resourceCgnv6LsnRuleListDomainListNameStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDomainListNameStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDomainListNameStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6LsnRuleListDomainListNameStatsStats := setObjectCgnv6LsnRuleListDomainListNameStatsStats(res)
		d.Set("stats", Cgnv6LsnRuleListDomainListNameStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6LsnRuleListDomainListNameStatsStats(ret edpt.DataCgnv6LsnRuleListDomainListNameStats) []interface{} {
	return []interface{}{
		map[string]interface{}{},
	}
}

func getObjectCgnv6LsnRuleListDomainListNameStatsStats(d []interface{}) edpt.Cgnv6LsnRuleListDomainListNameStatsStats {

	var ret edpt.Cgnv6LsnRuleListDomainListNameStatsStats
	return ret
}

func dataToEndpointCgnv6LsnRuleListDomainListNameStats(d *schema.ResourceData) edpt.Cgnv6LsnRuleListDomainListNameStats {
	var ret edpt.Cgnv6LsnRuleListDomainListNameStats

	ret.NameDomainList = d.Get("name_domain_list").(string)

	ret.Stats = getObjectCgnv6LsnRuleListDomainListNameStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
