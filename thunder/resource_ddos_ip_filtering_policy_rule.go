package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosIpFilteringPolicyRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_ip_filtering_policy_rule`: IP filter rule configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosIpFilteringPolicyRuleCreate,
		UpdateContext: resourceDdosIpFilteringPolicyRuleUpdate,
		ReadContext:   resourceDdosIpFilteringPolicyRuleRead,
		DeleteContext: resourceDdosIpFilteringPolicyRuleDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop the packet (default); 'permit': Let the packet skip all afterword address filters; 'blacklist': Blacklist with glid; 'bypass': Bypass all the ddos process;",
			},
			"dst_ip": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Subnet address",
			},
			"dst_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Subnet address",
			},
			"dst_port": {
				Type: schema.TypeInt, Optional: true, Description: "Match only packets with the port number",
			},
			"dst_port_end": {
				Type: schema.TypeInt, Optional: true, Description: "Ending Port Number",
			},
			"dst_port_start": {
				Type: schema.TypeInt, Optional: true, Description: "Match only packets in the range of port numbers (Starting Port Number)",
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"icmp_code": {
				Type: schema.TypeInt, Optional: true, Description: "ICMP code",
			},
			"icmp_type": {
				Type: schema.TypeInt, Optional: true, Description: "ICMP message type",
			},
			"proto_num": {
				Type: schema.TypeInt, Optional: true, Description: "IP proto number",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP; 'udp': UDP; 'icmp-v4': ICMP; 'icmp-v6': ICMPv6; 'number': Specify IP protocol number;",
			},
			"seq": {
				Type: schema.TypeInt, Required: true, Description: "Sequence number",
			},
			"src_ip": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Subnet address",
			},
			"src_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Subnet address",
			},
			"src_port": {
				Type: schema.TypeInt, Optional: true, Description: "Match only packets with the port number",
			},
			"src_port_end": {
				Type: schema.TypeInt, Optional: true, Description: "Ending Port Number",
			},
			"src_port_start": {
				Type: schema.TypeInt, Optional: true, Description: "Match only packets in the range of port numbers (Starting Port Number)",
			},
			"tcp_flag": {
				Type: schema.TypeString, Optional: true, Description: "'match-all': not = 0 match = 1; 'none-of': not = 1 match = 0; 'match-any': not = 0 match = 0;",
			},
			"tcp_flags_bitmask": {
				Type: schema.TypeString, Optional: true, Description: "Bitmask in Hex",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDdosIpFilteringPolicyRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicyRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosIpFilteringPolicyRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosIpFilteringPolicyRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicyRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosIpFilteringPolicyRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosIpFilteringPolicyRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicyRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosIpFilteringPolicyRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpFilteringPolicyRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpFilteringPolicyRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosIpFilteringPolicyRule(d *schema.ResourceData) edpt.DdosIpFilteringPolicyRule {
	var ret edpt.DdosIpFilteringPolicyRule
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstIp = d.Get("dst_ip").(string)
	ret.Inst.DstIpv6 = d.Get("dst_ipv6").(string)
	ret.Inst.DstPort = d.Get("dst_port").(int)
	ret.Inst.DstPortEnd = d.Get("dst_port_end").(int)
	ret.Inst.DstPortStart = d.Get("dst_port_start").(int)
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.IcmpCode = d.Get("icmp_code").(int)
	ret.Inst.IcmpType = d.Get("icmp_type").(int)
	ret.Inst.ProtoNum = d.Get("proto_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Seq = d.Get("seq").(int)
	ret.Inst.SrcIp = d.Get("src_ip").(string)
	ret.Inst.SrcIpv6 = d.Get("src_ipv6").(string)
	ret.Inst.SrcPort = d.Get("src_port").(int)
	ret.Inst.SrcPortEnd = d.Get("src_port_end").(int)
	ret.Inst.SrcPortStart = d.Get("src_port_start").(int)
	ret.Inst.TcpFlag = d.Get("tcp_flag").(string)
	ret.Inst.TcpFlagsBitmask = d.Get("tcp_flags_bitmask").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
