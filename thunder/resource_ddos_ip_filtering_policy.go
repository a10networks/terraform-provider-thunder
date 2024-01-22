package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosIpFilteringPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_ip_filtering_policy`: IP Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosIpFilteringPolicyCreate,
		UpdateContext: resourceDdosIpFilteringPolicyUpdate,
		ReadContext:   resourceDdosIpFilteringPolicyRead,
		DeleteContext: resourceDdosIpFilteringPolicyDelete,

		Schema: map[string]*schema.Schema{
			"default_action": {
				Type: schema.TypeString, Optional: true, Default: "permit", Description: "'drop': Drop all the packets not meet any rule; 'permit': Forward all the packets not meet any rule (Default);",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS ip-filtering-policy name",
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop the packet (default); 'permit': Let the packet skip all afterword address filters; 'blacklist': Blacklist with glid; 'bypass': Bypass all the ddos process;",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"src_ip": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Subnet address",
						},
						"src_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Subnet address",
						},
						"dst_ip": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Subnet address",
						},
						"dst_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Subnet address",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': TCP; 'udp': UDP; 'icmp-v4': ICMP; 'icmp-v6': ICMPv6; 'number': Specify IP protocol number;",
						},
						"proto_num": {
							Type: schema.TypeInt, Optional: true, Description: "IP proto number",
						},
						"src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with the port number",
						},
						"src_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets in the range of port numbers (Starting Port Number)",
						},
						"src_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ending Port Number",
						},
						"dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with the port number",
						},
						"dst_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets in the range of port numbers (Starting Port Number)",
						},
						"dst_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ending Port Number",
						},
						"tcp_flag": {
							Type: schema.TypeString, Optional: true, Description: "'match-all': not = 0 match = 1; 'none-of': not = 1 match = 0; 'match-any': not = 0 match = 0;",
						},
						"tcp_flags_bitmask": {
							Type: schema.TypeString, Optional: true, Description: "Bitmask in Hex",
						},
						"icmp_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP message type",
						},
						"icmp_code": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP code",
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
		},
	}
}
func resourceDdosIpFilteringPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosIpFilteringPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosIpFilteringPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosIpFilteringPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosIpFilteringPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosIpFilteringPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosIpFilteringPolicyRuleList(d []interface{}) []edpt.DdosIpFilteringPolicyRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosIpFilteringPolicyRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosIpFilteringPolicyRuleList
		oi.Seq = in["seq"].(int)
		oi.Action = in["action"].(string)
		oi.Glid = in["glid"].(string)
		oi.SrcIp = in["src_ip"].(string)
		oi.SrcIpv6 = in["src_ipv6"].(string)
		oi.DstIp = in["dst_ip"].(string)
		oi.DstIpv6 = in["dst_ipv6"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.ProtoNum = in["proto_num"].(int)
		oi.SrcPort = in["src_port"].(int)
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		oi.DstPort = in["dst_port"].(int)
		oi.DstPortStart = in["dst_port_start"].(int)
		oi.DstPortEnd = in["dst_port_end"].(int)
		oi.TcpFlag = in["tcp_flag"].(string)
		oi.TcpFlagsBitmask = in["tcp_flags_bitmask"].(string)
		oi.IcmpType = in["icmp_type"].(int)
		oi.IcmpCode = in["icmp_code"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosIpFilteringPolicy(d *schema.ResourceData) edpt.DdosIpFilteringPolicy {
	var ret edpt.DdosIpFilteringPolicy
	ret.Inst.DefaultAction = d.Get("default_action").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RuleList = getSliceDdosIpFilteringPolicyRuleList(d.Get("rule_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
