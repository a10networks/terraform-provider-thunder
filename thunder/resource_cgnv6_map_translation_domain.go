package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapTranslationDomain() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_translation_domain`: MAP Translation domain\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapTranslationDomainCreate,
		UpdateContext: resourceCgnv6MapTranslationDomainUpdate,
		ReadContext:   resourceCgnv6MapTranslationDomainRead,
		DeleteContext: resourceCgnv6MapTranslationDomainDelete,

		Schema: map[string]*schema.Schema{
			"basic_mapping_rule": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_ipv4_address_port_settings": {
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
					},
				},
			},
			"default_mapping_rule": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_ipv6_prefix": {
							Type: schema.TypeString, Optional: true, Description: "Rule IPv6 prefix",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "MAP-T domain description",
			},
			"health_check_gateway": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_gateway": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 Gateway",
									},
								},
							},
						},
						"ipv6_address_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_gateway": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 Gateway",
									},
								},
							},
						},
						"withdraw_route": {
							Type: schema.TypeString, Optional: true, Default: "any-link-failure", Description: "'all-link-failure': Withdraw routes on health-check failure of all IPv4 gateways or all IPv6 gateways; 'any-link-failure': Withdraw routes on health-check failure of any gateway (default);",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mtu": {
				Type: schema.TypeInt, Optional: true, Description: "Domain MTU",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "MAP-T domain name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'inbound_packet_received': Inbound IPv4 Packets Received; 'inbound_frag_packet_received': Inbound IPv4 Fragment Packets Received; 'inbound_addr_port_validation_failed': Inbound IPv4 Destination Address Port Validation Failed; 'inbound_rev_lookup_failed': Inbound IPv4 Reverse Route Lookup Failed; 'inbound_dest_unreachable': Inbound IPv6 Destination Address Unreachable; 'outbound_packet_received': Outbound IPv6 Packets Received; 'outbound_frag_packet_received': Outbound IPv6 Fragment Packets Received; 'outbound_addr_validation_failed': Outbound IPv6 Source Address Validation Failed; 'outbound_rev_lookup_failed': Outbound IPv6 Reverse Route Lookup Failed; 'outbound_dest_unreachable': Outbound IPv4 Destination Address Unreachable; 'packet_mtu_exceeded': Packet Exceeded MTU; 'frag_icmp_sent': ICMP Packet Too Big Sent; 'interface_not_configured': Interfaces not Configured Dropped; 'bmr_prefixrules_configured': BMR prefix rules configured; 'helper_count': Helper Count; 'active_dhcpv6_leases': Active DHCPv6 leases;",
						},
					},
				},
			},
			"tcp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mss_clamp": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mss_clamp_type": {
										Type: schema.TypeString, Optional: true, Default: "none", Description: "'fixed': Specify a fixed max value for the TCP MSS; 'none': No TCP MSS clamping(default); 'subtract': Specify the value to subtract from the TCP MSS;",
									},
									"mss_value": {
										Type: schema.TypeInt, Optional: true, Description: "The max value allowed for the TCP MSS",
									},
									"mss_subtract": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the value to subtract from the TCP MSS",
									},
									"min": {
										Type: schema.TypeInt, Optional: true, Default: 516, Description: "Specify the min value allowed for the TCP MSS (Specify the min value allowed for the TCP MSS (default: 516))",
									},
								},
							},
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
		},
	}
}
func resourceCgnv6MapTranslationDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomain(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapTranslationDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomain(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapTranslationDomainRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapTranslationDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomain(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapTranslationDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapTranslationDomainRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapTranslationDomain(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6MapTranslationDomainBasicMappingRule102(d []interface{}) edpt.Cgnv6MapTranslationDomainBasicMappingRule102 {

	count1 := len(d)
	var ret edpt.Cgnv6MapTranslationDomainBasicMappingRule102
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleIpv4AddressPortSettings = in["rule_ipv4_address_port_settings"].(string)
		ret.EaLength = in["ea_length"].(int)
		ret.ShareRatio = in["share_ratio"].(int)
		ret.PortStart = in["port_start"].(int)
		//omit uuid
		ret.PrefixRuleList = getSliceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleList103(in["prefix_rule_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6MapTranslationDomainBasicMappingRulePrefixRuleList103(d []interface{}) []edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList103 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList103, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList103
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

func getObjectCgnv6MapTranslationDomainDefaultMappingRule104(d []interface{}) edpt.Cgnv6MapTranslationDomainDefaultMappingRule104 {

	count1 := len(d)
	var ret edpt.Cgnv6MapTranslationDomainDefaultMappingRule104
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleIpv6Prefix = in["rule_ipv6_prefix"].(string)
		//omit uuid
	}
	return ret
}

func getObjectCgnv6MapTranslationDomainHealthCheckGateway105(d []interface{}) edpt.Cgnv6MapTranslationDomainHealthCheckGateway105 {

	count1 := len(d)
	var ret edpt.Cgnv6MapTranslationDomainHealthCheckGateway105
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceCgnv6MapTranslationDomainHealthCheckGatewayAddressList106(in["address_list"].([]interface{}))
		ret.Ipv6AddressList = getSliceCgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList107(in["ipv6_address_list"].([]interface{}))
		ret.WithdrawRoute = in["withdraw_route"].(string)
		//omit uuid
	}
	return ret
}

func getSliceCgnv6MapTranslationDomainHealthCheckGatewayAddressList106(d []interface{}) []edpt.Cgnv6MapTranslationDomainHealthCheckGatewayAddressList106 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapTranslationDomainHealthCheckGatewayAddressList106, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapTranslationDomainHealthCheckGatewayAddressList106
		oi.Ipv4Gateway = in["ipv4_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList107(d []interface{}) []edpt.Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList107 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList107, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList107
		oi.Ipv6Gateway = in["ipv6_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6MapTranslationDomainSamplingEnable(d []interface{}) []edpt.Cgnv6MapTranslationDomainSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapTranslationDomainSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapTranslationDomainSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6MapTranslationDomainTcp(d []interface{}) edpt.Cgnv6MapTranslationDomainTcp {

	count1 := len(d)
	var ret edpt.Cgnv6MapTranslationDomainTcp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MssClamp = getObjectCgnv6MapTranslationDomainTcpMssClamp(in["mss_clamp"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6MapTranslationDomainTcpMssClamp(d []interface{}) edpt.Cgnv6MapTranslationDomainTcpMssClamp {

	count1 := len(d)
	var ret edpt.Cgnv6MapTranslationDomainTcpMssClamp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MssClampType = in["mss_clamp_type"].(string)
		ret.MssValue = in["mss_value"].(int)
		ret.MssSubtract = in["mss_subtract"].(int)
		ret.Min = in["min"].(int)
	}
	return ret
}

func dataToEndpointCgnv6MapTranslationDomain(d *schema.ResourceData) edpt.Cgnv6MapTranslationDomain {
	var ret edpt.Cgnv6MapTranslationDomain
	ret.Inst.BasicMappingRule = getObjectCgnv6MapTranslationDomainBasicMappingRule102(d.Get("basic_mapping_rule").([]interface{}))
	ret.Inst.DefaultMappingRule = getObjectCgnv6MapTranslationDomainDefaultMappingRule104(d.Get("default_mapping_rule").([]interface{}))
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.HealthCheckGateway = getObjectCgnv6MapTranslationDomainHealthCheckGateway105(d.Get("health_check_gateway").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6MapTranslationDomainSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Tcp = getObjectCgnv6MapTranslationDomainTcp(d.Get("tcp").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
