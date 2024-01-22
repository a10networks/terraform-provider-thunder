package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationDomainBasicMappingRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_domain_basic_mapping_rule`: Basic mapping rule (BMR)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationDomainBasicMappingRuleCreate,
		UpdateContext: resourceCgnv6MapTranslationDomainBasicMappingRuleUpdate,
		ReadContext:   resourceCgnv6MapTranslationDomainBasicMappingRuleRead,
		DeleteContext: resourceCgnv6MapTranslationDomainBasicMappingRuleDelete,

		Schema: map[string]*schema.Schema{
			"ea_length": {
				Type: schema.TypeInt, Optional: true, Description: "Length of Embedded Address (EA) bits",
			},
			"port_start": {
				Type: schema.TypeInt, Optional: true, Description: "Starting Port, Must be Power of 2 value or zero",
			},
			"prefix_rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "MAP BMR prefix rule name",
						},
						"rule_ipv6_prefix": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 prefix of BMR",
						},
						"rule_ipv4_prefix": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 prefix of BMR",
						},
						"ipv4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "Subnet mask (subnet bigger than /8 is not allowed)",
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
			"rule_ipv4_address_port_settings": {
				Type: schema.TypeString, Optional: true, Description: "'prefix-addr': Each CE is assigned an IPv4 prefix; 'single-addr': Each CE is assigned an IPv4 address; 'shared-addr': Each CE is assigned a shared IPv4 address;",
			},
			"share_ratio": {
				Type: schema.TypeInt, Optional: true, Description: "Port sharing ratio for each NAT IP. Must be Power of 2 value",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6MapTranslationDomainBasicMappingRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainBasicMappingRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationDomainBasicMappingRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainBasicMappingRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationDomainBasicMappingRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationDomainBasicMappingRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleList(d []interface{}) []edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList
		oi.Name = in["name"].(string)
		oi.RuleIpv6Prefix = in["rule_ipv6_prefix"].(string)
		oi.RuleIpv4Prefix = in["rule_ipv4_prefix"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6MapTranslationDomainBasicMappingRule(d *schema.ResourceData) edpt.Cgnv6MapTranslationDomainBasicMappingRule {
	var ret edpt.Cgnv6MapTranslationDomainBasicMappingRule
	ret.Inst.EaLength = d.Get("ea_length").(int)
	ret.Inst.PortStart = d.Get("port_start").(int)
	ret.Inst.PrefixRuleList = getSliceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleList(d.Get("prefix_rule_list").([]interface{}))
	ret.Inst.RuleIpv4AddressPortSettings = d.Get("rule_ipv4_address_port_settings").(string)
	ret.Inst.ShareRatio = d.Get("share_ratio").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
