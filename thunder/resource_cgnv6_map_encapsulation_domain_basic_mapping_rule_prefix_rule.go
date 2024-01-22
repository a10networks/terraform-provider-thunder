package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_encapsulation_domain_basic_mapping_rule_prefix_rule`: IPv6 and IPv4 prefix rules\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleCreate,
		UpdateContext: resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleUpdate,
		ReadContext:   resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleRead,
		DeleteContext: resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleDelete,

		Schema: map[string]*schema.Schema{
			"ea_length": {
				Type: schema.TypeInt, Optional: true, Description: "Length of Embedded Address (EA) bits",
			},
			"ipv4_address_port_settings": {
				Type: schema.TypeString, Optional: true, Description: "'prefix-addr': Each CE is assigned an IPv4 prefix; 'single-addr': Each CE is assigned an IPv4 address; 'shared-addr': Each CE is assigned a shared IPv4 address;",
			},
			"ipv4_netmask": {
				Type: schema.TypeString, Optional: true, Description: "Subnet mask (subnet bigger than /8 is not allowed)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "MAP BMR prefix rule name",
			},
			"port_start": {
				Type: schema.TypeInt, Optional: true, Description: "Starting Port, Must be Power of 2 value or zero",
			},
			"rule_ipv4_prefix": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 prefix of BMR",
			},
			"rule_ipv6_prefix": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 prefix of BMR",
			},
			"share_ratio": {
				Type: schema.TypeInt, Optional: true, Description: "Port sharing ratio for each NAT IP. Must be Power of 2 value",
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
func resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRulePrefixRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRulePrefixRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRulePrefixRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRulePrefixRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRulePrefixRule(d *schema.ResourceData) edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRule {
	var ret edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRule
	ret.Inst.EaLength = d.Get("ea_length").(int)
	ret.Inst.Ipv4AddressPortSettings = d.Get("ipv4_address_port_settings").(string)
	ret.Inst.Ipv4Netmask = d.Get("ipv4_netmask").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PortStart = d.Get("port_start").(int)
	ret.Inst.RuleIpv4Prefix = d.Get("rule_ipv4_prefix").(string)
	ret.Inst.RuleIpv6Prefix = d.Get("rule_ipv6_prefix").(string)
	ret.Inst.ShareRatio = d.Get("share_ratio").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
