package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceMatchingRulesRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_service_matching_rules_rule`: Service-matching rule\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbServiceMatchingRulesRuleCreate,
		UpdateContext: resourceGslbServiceMatchingRulesRuleUpdate,
		ReadContext:   resourceGslbServiceMatchingRulesRuleRead,
		DeleteContext: resourceGslbServiceMatchingRulesRuleDelete,

		Schema: map[string]*schema.Schema{
			"domain_match_string": {
				Type: schema.TypeString, Optional: true, Description: "Domain name string",
			},
			"domain_match_type": {
				Type: schema.TypeString, Optional: true, Description: "'equals': Domain name equals to string; 'contains': Domain name contains string; 'starts-with': Domain name starts with string; 'ends-with': Domain name ends with string;",
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
			"seq_num": {
				Type: schema.TypeInt, Required: true, Description: "Sequence number of rule",
			},
			"service": {
				Type: schema.TypeString, Optional: true, Description: "Name of zone service to be redirected",
			},
			"src_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "Client source IPv4 subnet",
			},
			"src_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "Client source IPv6 subnet",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone": {
				Type: schema.TypeString, Required: true, Description: "Zone",
			},
		},
	}
}
func resourceGslbServiceMatchingRulesRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRulesRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceMatchingRulesRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbServiceMatchingRulesRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRulesRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceMatchingRulesRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbServiceMatchingRulesRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRulesRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbServiceMatchingRulesRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceMatchingRulesRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceMatchingRulesRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbServiceMatchingRulesRuleHealthState(d []interface{}) []edpt.GslbServiceMatchingRulesRuleHealthState {

	count1 := len(d)
	ret := make([]edpt.GslbServiceMatchingRulesRuleHealthState, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceMatchingRulesRuleHealthState
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

func dataToEndpointGslbServiceMatchingRulesRule(d *schema.ResourceData) edpt.GslbServiceMatchingRulesRule {
	var ret edpt.GslbServiceMatchingRulesRule
	ret.Inst.DomainMatchString = d.Get("domain_match_string").(string)
	ret.Inst.DomainMatchType = d.Get("domain_match_type").(string)
	ret.Inst.HealthState = getSliceGslbServiceMatchingRulesRuleHealthState(d.Get("health_state").([]interface{}))
	ret.Inst.SeqNum = d.Get("seq_num").(int)
	ret.Inst.Service = d.Get("service").(string)
	ret.Inst.SrcIpv4 = d.Get("src_ipv4").(string)
	ret.Inst.SrcIpv6 = d.Get("src_ipv6").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Zone = d.Get("zone").(string)
	return ret
}
