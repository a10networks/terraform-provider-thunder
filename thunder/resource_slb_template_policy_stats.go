package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_stats`: Statistics for the object policy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Policy template name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fwd_policy_dns_unresolved": {
							Type: schema.TypeInt, Optional: true, Description: "Forward-policy unresolved DNS queries",
						},
						"fwd_policy_dns_outstanding": {
							Type: schema.TypeInt, Optional: true, Description: "Forward-policy current DNS outstanding requests",
						},
						"fwd_policy_snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Forward-policy source-nat translation failure",
						},
						"fwd_policy_hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of forward-policy requests for this policy template",
						},
						"fwd_policy_forward_to_internet": {
							Type: schema.TypeInt, Optional: true, Description: "Number of forward-policy requests forwarded to internet",
						},
						"fwd_policy_forward_to_service_group": {
							Type: schema.TypeInt, Optional: true, Description: "Number of forward-policy requests forwarded to service group",
						},
						"fwd_policy_forward_to_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Number of forward-policy requests forwarded to proxy",
						},
						"fwd_policy_policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Number of forward-policy requests dropped",
						},
						"fwd_policy_source_match_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Forward-policy requests without matching source rule",
						},
						"exp_client_hello_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Expected Client HELLO requests not found",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplatePolicyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyStatsStats := setObjectSlbTemplatePolicyStatsStats(res)
		d.Set("stats", SlbTemplatePolicyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyStatsStats(ret edpt.DataSlbTemplatePolicyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"fwd_policy_dns_unresolved":           ret.DtSlbTemplatePolicyStats.Stats.FwdPolicyDnsUnresolved,
			"fwd_policy_dns_outstanding":          ret.DtSlbTemplatePolicyStats.Stats.FwdPolicyDnsOutstanding,
			"fwd_policy_snat_fail":                ret.DtSlbTemplatePolicyStats.Stats.FwdPolicySnatFail,
			"fwd_policy_hits":                     ret.DtSlbTemplatePolicyStats.Stats.FwdPolicyHits,
			"fwd_policy_forward_to_internet":      ret.DtSlbTemplatePolicyStats.Stats.FwdPolicyForwardToInternet,
			"fwd_policy_forward_to_service_group": ret.DtSlbTemplatePolicyStats.Stats.FwdPolicyForwardToServiceGroup,
			"fwd_policy_forward_to_proxy":         ret.DtSlbTemplatePolicyStats.Stats.FwdPolicyForwardToProxy,
			"fwd_policy_policy_drop":              ret.DtSlbTemplatePolicyStats.Stats.FwdPolicyPolicyDrop,
			"fwd_policy_source_match_not_found":   ret.DtSlbTemplatePolicyStats.Stats.FwdPolicySourceMatchNotFound,
			"exp_client_hello_not_found":          ret.DtSlbTemplatePolicyStats.Stats.ExpClientHelloNotFound,
		},
	}
}

func getObjectSlbTemplatePolicyStatsStats(d []interface{}) edpt.SlbTemplatePolicyStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FwdPolicyDnsUnresolved = in["fwd_policy_dns_unresolved"].(int)
		ret.FwdPolicyDnsOutstanding = in["fwd_policy_dns_outstanding"].(int)
		ret.FwdPolicySnatFail = in["fwd_policy_snat_fail"].(int)
		ret.FwdPolicyHits = in["fwd_policy_hits"].(int)
		ret.FwdPolicyForwardToInternet = in["fwd_policy_forward_to_internet"].(int)
		ret.FwdPolicyForwardToServiceGroup = in["fwd_policy_forward_to_service_group"].(int)
		ret.FwdPolicyForwardToProxy = in["fwd_policy_forward_to_proxy"].(int)
		ret.FwdPolicyPolicyDrop = in["fwd_policy_policy_drop"].(int)
		ret.FwdPolicySourceMatchNotFound = in["fwd_policy_source_match_not_found"].(int)
		ret.ExpClientHelloNotFound = in["exp_client_hello_not_found"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyStats(d *schema.ResourceData) edpt.SlbTemplatePolicyStats {
	var ret edpt.SlbTemplatePolicyStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplatePolicyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
