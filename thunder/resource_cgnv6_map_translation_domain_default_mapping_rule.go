package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationDomainDefaultMappingRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_domain_default_mapping_rule`: Default mapping rule (DMR)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationDomainDefaultMappingRuleCreate,
		UpdateContext: resourceCgnv6MapTranslationDomainDefaultMappingRuleUpdate,
		ReadContext:   resourceCgnv6MapTranslationDomainDefaultMappingRuleRead,
		DeleteContext: resourceCgnv6MapTranslationDomainDefaultMappingRuleDelete,

		Schema: map[string]*schema.Schema{
			"rule_ipv6_prefix": {
				Type: schema.TypeString, Optional: true, Description: "Rule IPv6 prefix",
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
func resourceCgnv6MapTranslationDomainDefaultMappingRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainDefaultMappingRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainDefaultMappingRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainDefaultMappingRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationDomainDefaultMappingRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainDefaultMappingRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainDefaultMappingRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainDefaultMappingRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationDomainDefaultMappingRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainDefaultMappingRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainDefaultMappingRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationDomainDefaultMappingRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainDefaultMappingRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomainDefaultMappingRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6MapTranslationDomainDefaultMappingRule(d *schema.ResourceData) edpt.Cgnv6MapTranslationDomainDefaultMappingRule {
	var ret edpt.Cgnv6MapTranslationDomainDefaultMappingRule
	ret.Inst.RuleIpv6Prefix = d.Get("rule_ipv6_prefix").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
