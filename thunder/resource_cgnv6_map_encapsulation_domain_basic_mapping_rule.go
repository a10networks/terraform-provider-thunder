package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapEncapsulationDomainBasicMappingRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_encapsulation_domain_basic_mapping_rule`: Basic mapping rule (BMR)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapEncapsulationDomainBasicMappingRuleCreate,
		UpdateContext: resourceCgnv6MapEncapsulationDomainBasicMappingRuleUpdate,
		ReadContext:   resourceCgnv6MapEncapsulationDomainBasicMappingRuleRead,
		DeleteContext: resourceCgnv6MapEncapsulationDomainBasicMappingRuleDelete,

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
						"ipv4_address_port_settings": {
							Type: schema.TypeString, Optional: true, Description: "'prefix-addr': Each CE is assigned an IPv4 prefix; 'single-addr': Each CE is assigned an IPv4 address; 'shared-addr': Each CE is assigned a shared IPv4 address;",
						},
						"ea_length": {
							Type: schema.TypeInt, Optional: true, Description: "Length of Embedded Address (EA) bits",
						},
						"share_ratio": {
							Type: schema.TypeInt, Optional: true, Description: "Port sharing ratio for each NAT IP. Must be Power of 2 value",
						},
						"port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Starting Port, Must be Power of 2 value or zero",
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
func resourceCgnv6MapEncapsulationDomainBasicMappingRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationDomainBasicMappingRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapEncapsulationDomainBasicMappingRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationDomainBasicMappingRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapEncapsulationDomainBasicMappingRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapEncapsulationDomainBasicMappingRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainBasicMappingRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList(d []interface{}) []edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList
		oi.Name = in["name"].(string)
		oi.RuleIpv6Prefix = in["rule_ipv6_prefix"].(string)
		oi.RuleIpv4Prefix = in["rule_ipv4_prefix"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		oi.Ipv4AddressPortSettings = in["ipv4_address_port_settings"].(string)
		oi.EaLength = in["ea_length"].(int)
		oi.ShareRatio = in["share_ratio"].(int)
		oi.PortStart = in["port_start"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6MapEncapsulationDomainBasicMappingRule(d *schema.ResourceData) edpt.Cgnv6MapEncapsulationDomainBasicMappingRule {
	var ret edpt.Cgnv6MapEncapsulationDomainBasicMappingRule
	ret.Inst.EaLength = d.Get("ea_length").(int)
	ret.Inst.PortStart = d.Get("port_start").(int)
	ret.Inst.PrefixRuleList = getSliceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList(d.Get("prefix_rule_list").([]interface{}))
	ret.Inst.RuleIpv4AddressPortSettings = d.Get("rule_ipv4_address_port_settings").(string)
	ret.Inst.ShareRatio = d.Get("share_ratio").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
