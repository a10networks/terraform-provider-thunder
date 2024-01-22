package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_domain_basic_mapping_rule_prefix_rule`: IPv6 and IPv4 prefix rules\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleCreate,
		UpdateContext: resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleUpdate,
		ReadContext:   resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleRead,
		DeleteContext: resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_netmask": {
				Type: schema.TypeString, Optional: true, Description: "Subnet mask (subnet bigger than /8 is not allowed)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "MAP BMR prefix rule name",
			},
			"rule_ipv4_prefix": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 prefix of BMR",
			},
			"rule_ipv6_prefix": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 prefix of BMR",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRulePrefixRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRulePrefixRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRulePrefixRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainBasicMappingRulePrefixRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6MapTranslationDomainBasicMappingRulePrefixRule(d *schema.ResourceData) edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRule {
	var ret edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRule
	ret.Inst.Ipv4Netmask = d.Get("ipv4_netmask").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RuleIpv4Prefix = d.Get("rule_ipv4_prefix").(string)
	ret.Inst.RuleIpv6Prefix = d.Get("rule_ipv6_prefix").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
