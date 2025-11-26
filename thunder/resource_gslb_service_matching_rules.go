package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceMatchingRules() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_service_matching_rules`: Configure service-matching rules for a GSLB zone\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbServiceMatchingRulesCreate,
		UpdateContext: resourceGslbServiceMatchingRulesUpdate,
		ReadContext:   resourceGslbServiceMatchingRulesRead,
		DeleteContext: resourceGslbServiceMatchingRulesDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable rule matching on associated zone (default is enabled)",
			},
			"hitcount_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable hit count of rule matching",
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number of rule",
						},
						"domain_match_type": {
							Type: schema.TypeString, Optional: true, Description: "'equals': Domain name equals to string; 'contains': Domain name contains string; 'starts-with': Domain name starts with string; 'ends-with': Domain name ends with string;",
						},
						"domain_match_string": {
							Type: schema.TypeString, Optional: true, Description: "Domain name string",
						},
						"src_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "Client source IPv4 subnet",
						},
						"src_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Client source IPv6 subnet",
						},
						"health_state": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gslb_site": {
										Type: schema.TypeString, Optional: true, Description: "Health state of GSLB site (GSLB site name)",
									},
									"site_state": {
										Type: schema.TypeString, Optional: true, Description: "'AllUp': AllUp; 'Down': Down; 'AllUp-or-PartUp': AllUp-or-PartUp; 'Down-or-PartUp': Down-or-PartUp;",
									},
									"gslb_service_ip": {
										Type: schema.TypeString, Optional: true, Description: "Health state of GSLB service-ip (GSLB service-ip name)",
									},
									"service_ip_state": {
										Type: schema.TypeString, Optional: true, Description: "'Up': Up; 'Down': Down;",
									},
									"slb_server": {
										Type: schema.TypeString, Optional: true, Description: "Health state of SLB server (SLB server name)",
									},
									"slb_svr_state": {
										Type: schema.TypeString, Optional: true, Description: "'Up': Up; 'Down': Down;",
									},
								},
							},
						},
						"service": {
							Type: schema.TypeString, Optional: true, Description: "Name of zone service to be redirected",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone": {
				Type: schema.TypeString, Required: true, Description: "GSLB zone name",
			},
		},
	}
}
func resourceGslbServiceMatchingRulesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRules(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceMatchingRulesRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbServiceMatchingRulesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRules(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceMatchingRulesRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbServiceMatchingRulesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRules(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbServiceMatchingRulesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRules(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbServiceMatchingRulesRuleList(d []interface{}) []edpt.GslbServiceMatchingRulesRuleList {

	count1 := len(d)
	ret := make([]edpt.GslbServiceMatchingRulesRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceMatchingRulesRuleList
		oi.SeqNum = in["seq_num"].(int)
		oi.DomainMatchType = in["domain_match_type"].(string)
		oi.DomainMatchString = in["domain_match_string"].(string)
		oi.SrcIpv4 = in["src_ipv4"].(string)
		oi.SrcIpv6 = in["src_ipv6"].(string)
		oi.HealthState = getSliceGslbServiceMatchingRulesRuleListHealthState(in["health_state"].([]interface{}))
		oi.Service = in["service"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbServiceMatchingRulesRuleListHealthState(d []interface{}) []edpt.GslbServiceMatchingRulesRuleListHealthState {

	count1 := len(d)
	ret := make([]edpt.GslbServiceMatchingRulesRuleListHealthState, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceMatchingRulesRuleListHealthState
		oi.GslbSite = in["gslb_site"].(string)
		oi.SiteState = in["site_state"].(string)
		oi.GslbServiceIp = in["gslb_service_ip"].(string)
		oi.ServiceIpState = in["service_ip_state"].(string)
		oi.SlbServer = in["slb_server"].(string)
		oi.SlbSvrState = in["slb_svr_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbServiceMatchingRules(d *schema.ResourceData) edpt.GslbServiceMatchingRules {
	var ret edpt.GslbServiceMatchingRules
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.HitcountEnable = d.Get("hitcount_enable").(int)
	ret.Inst.RuleList = getSliceGslbServiceMatchingRulesRuleList(d.Get("rule_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Zone = d.Get("zone").(string)
	return ret
}
