package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwMatchOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_match_oper`: Operational Status for the object match\n\n__PLACEHOLDER__",
		ReadContext: resourceFwMatchOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_msg": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"active_access_control_rule_set": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"active_traffic_control_rule_set": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"rule_ids": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"matching_rule_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"matching_rule": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"matching_rules_fetched": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"matching_rules_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"show_all": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vlan": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tcp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"udp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"icmpv6": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"icmp_type": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwMatchOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwMatchOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwMatchOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwMatchOperOper := setObjectFwMatchOperOper(res)
		d.Set("oper", FwMatchOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwMatchOperOper(ret edpt.DataFwMatchOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"error_msg":                       ret.DtFwMatchOper.Oper.ErrorMsg,
			"active_access_control_rule_set":  ret.DtFwMatchOper.Oper.ActiveAccessControlRuleSet,
			"active_traffic_control_rule_set": ret.DtFwMatchOper.Oper.ActiveTrafficControlRuleSet,
			"rule_ids":                        setSliceFwMatchOperOperRuleIds(ret.DtFwMatchOper.Oper.RuleIds),
			"matching_rules_fetched":          ret.DtFwMatchOper.Oper.MatchingRulesFetched,
			"matching_rules_total":            ret.DtFwMatchOper.Oper.MatchingRulesTotal,
			"show_all":                        ret.DtFwMatchOper.Oper.ShowAll,
			"vlan":                            ret.DtFwMatchOper.Oper.Vlan,
			"src_ipv4_addr":                   ret.DtFwMatchOper.Oper.SrcIpv4Addr,
			"dst_ipv4_addr":                   ret.DtFwMatchOper.Oper.DstIpv4Addr,
			"src_ipv6_addr":                   ret.DtFwMatchOper.Oper.SrcIpv6Addr,
			"dst_ipv6_addr":                   ret.DtFwMatchOper.Oper.DstIpv6Addr,
			"tcp":                             ret.DtFwMatchOper.Oper.Tcp,
			"udp":                             ret.DtFwMatchOper.Oper.Udp,
			"icmp":                            ret.DtFwMatchOper.Oper.Icmp,
			"icmpv6":                          ret.DtFwMatchOper.Oper.Icmpv6,
			"src_port":                        ret.DtFwMatchOper.Oper.SrcPort,
			"dst_port":                        ret.DtFwMatchOper.Oper.DstPort,
			"icmp_type":                       ret.DtFwMatchOper.Oper.IcmpType,
		},
	}
}

func setSliceFwMatchOperOperRuleIds(d []edpt.FwMatchOperOperRuleIds) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["matching_rule_type"] = item.MatchingRuleType
		in["matching_rule"] = item.MatchingRule
		result = append(result, in)
	}
	return result
}

func getObjectFwMatchOperOper(d []interface{}) edpt.FwMatchOperOper {

	count1 := len(d)
	var ret edpt.FwMatchOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ErrorMsg = in["error_msg"].(string)
		ret.ActiveAccessControlRuleSet = in["active_access_control_rule_set"].(string)
		ret.ActiveTrafficControlRuleSet = in["active_traffic_control_rule_set"].(string)
		ret.RuleIds = getSliceFwMatchOperOperRuleIds(in["rule_ids"].([]interface{}))
		ret.MatchingRulesFetched = in["matching_rules_fetched"].(int)
		ret.MatchingRulesTotal = in["matching_rules_total"].(int)
		ret.ShowAll = in["show_all"].(int)
		ret.Vlan = in["vlan"].(int)
		ret.SrcIpv4Addr = in["src_ipv4_addr"].(string)
		ret.DstIpv4Addr = in["dst_ipv4_addr"].(string)
		ret.SrcIpv6Addr = in["src_ipv6_addr"].(string)
		ret.DstIpv6Addr = in["dst_ipv6_addr"].(string)
		ret.Tcp = in["tcp"].(int)
		ret.Udp = in["udp"].(int)
		ret.Icmp = in["icmp"].(int)
		ret.Icmpv6 = in["icmpv6"].(int)
		ret.SrcPort = in["src_port"].(int)
		ret.DstPort = in["dst_port"].(int)
		ret.IcmpType = in["icmp_type"].(int)
	}
	return ret
}

func getSliceFwMatchOperOperRuleIds(d []interface{}) []edpt.FwMatchOperOperRuleIds {

	count1 := len(d)
	ret := make([]edpt.FwMatchOperOperRuleIds, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwMatchOperOperRuleIds
		oi.MatchingRuleType = in["matching_rule_type"].(int)
		oi.MatchingRule = in["matching_rule"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwMatchOper(d *schema.ResourceData) edpt.FwMatchOper {
	var ret edpt.FwMatchOper

	ret.Oper = getObjectFwMatchOperOper(d.Get("oper").([]interface{}))
	return ret
}
