package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6MapEncapsulationDomain() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_map_encapsulation_domain`: MAP Encapsulation domain\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6MapEncapsulationDomainCreate,
		UpdateContext: resourceCgnv6MapEncapsulationDomainUpdate,
		ReadContext:   resourceCgnv6MapEncapsulationDomainRead,
		DeleteContext: resourceCgnv6MapEncapsulationDomainDelete,

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
					},
				},
			},
			"description": {
				Type: schema.TypeString, Optional: true, Description: "MAP-E domain description",
			},
			"format": {
				Type: schema.TypeString, Optional: true, Description: "'draft-03': Construct IPv6 Interface Identifier according to draft-03;",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "MAP-E domain name",
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
			"tunnel_endpoint_address": {
				Type: schema.TypeString, Optional: true, Description: "Tunnel Endpoint Address for MAP-E domain",
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
func resourceCgnv6MapEncapsulationDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomain(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationDomainRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6MapEncapsulationDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomain(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6MapEncapsulationDomainRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6MapEncapsulationDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomain(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6MapEncapsulationDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6MapEncapsulationDomainRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6MapEncapsulationDomain(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6MapEncapsulationDomainBasicMappingRule97(d []interface{}) edpt.Cgnv6MapEncapsulationDomainBasicMappingRule97 {

	count1 := len(d)
	var ret edpt.Cgnv6MapEncapsulationDomainBasicMappingRule97
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleIpv4AddressPortSettings = in["rule_ipv4_address_port_settings"].(string)
		ret.EaLength = in["ea_length"].(int)
		ret.ShareRatio = in["share_ratio"].(int)
		ret.PortStart = in["port_start"].(int)
		//omit uuid
		ret.PrefixRuleList = getSliceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList98(in["prefix_rule_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList98(d []interface{}) []edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList98 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList98, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList98
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

func getObjectCgnv6MapEncapsulationDomainHealthCheckGateway99(d []interface{}) edpt.Cgnv6MapEncapsulationDomainHealthCheckGateway99 {

	count1 := len(d)
	var ret edpt.Cgnv6MapEncapsulationDomainHealthCheckGateway99
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AddressList = getSliceCgnv6MapEncapsulationDomainHealthCheckGatewayAddressList100(in["address_list"].([]interface{}))
		ret.Ipv6AddressList = getSliceCgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList101(in["ipv6_address_list"].([]interface{}))
		ret.WithdrawRoute = in["withdraw_route"].(string)
		//omit uuid
	}
	return ret
}

func getSliceCgnv6MapEncapsulationDomainHealthCheckGatewayAddressList100(d []interface{}) []edpt.Cgnv6MapEncapsulationDomainHealthCheckGatewayAddressList100 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapEncapsulationDomainHealthCheckGatewayAddressList100, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapEncapsulationDomainHealthCheckGatewayAddressList100
		oi.Ipv4Gateway = in["ipv4_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList101(d []interface{}) []edpt.Cgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList101 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList101, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList101
		oi.Ipv6Gateway = in["ipv6_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6MapEncapsulationDomainSamplingEnable(d []interface{}) []edpt.Cgnv6MapEncapsulationDomainSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6MapEncapsulationDomainSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6MapEncapsulationDomainSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6MapEncapsulationDomain(d *schema.ResourceData) edpt.Cgnv6MapEncapsulationDomain {
	var ret edpt.Cgnv6MapEncapsulationDomain
	ret.Inst.BasicMappingRule = getObjectCgnv6MapEncapsulationDomainBasicMappingRule97(d.Get("basic_mapping_rule").([]interface{}))
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.Format = d.Get("format").(string)
	ret.Inst.HealthCheckGateway = getObjectCgnv6MapEncapsulationDomainHealthCheckGateway99(d.Get("health_check_gateway").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6MapEncapsulationDomainSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.TunnelEndpointAddress = d.Get("tunnel_endpoint_address").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
