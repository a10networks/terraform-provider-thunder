package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceObjectGroupService() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_object_group_service`: Configure Service Object Group\n\n__PLACEHOLDER__",
		CreateContext: resourceObjectGroupServiceCreate,
		UpdateContext: resourceObjectGroupServiceUpdate,
		ReadContext:   resourceObjectGroupServiceRead,
		DeleteContext: resourceObjectGroupServiceDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description of the object-group instance",
			},
			"rules": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"protocol_id": {
							Type: schema.TypeInt, Optional: true, Description: "Protocol ID",
						},
						"tcp_udp": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': Protocol TCP; 'udp': Protocol UDP;",
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Internet Control Message Protocol",
						},
						"icmpv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Internet Control Message Protocol version 6",
						},
						"icmp_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP type number",
						},
						"any_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any ICMP type",
						},
						"special_type": {
							Type: schema.TypeString, Optional: true, Description: "'echo-reply': Type 0, echo reply; 'echo-request': Type 8, echo request; 'info-reply': Type 16, information reply; 'info-request': Type 15, information request; 'mask-reply': Type 18, address mask reply; 'mask-request': Type 17, address mask request; 'parameter-problem': Type 12, parameter problem; 'redirect': Type 5, redirect message; 'source-quench': Type 4, source quench; 'time-exceeded': Type 11, time exceeded; 'timestamp': Type 13, timestamp; 'timestamp-reply': Type 14, timestamp reply; 'dest-unreachable': Type 3, destination unreachable;",
						},
						"any_code": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any ICMP code",
						},
						"icmp_code": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP code number",
						},
						"special_code": {
							Type: schema.TypeString, Optional: true, Description: "'frag-required': Code 4, fragmentation required; 'host-unreachable': Code 1, destination host unreachable; 'network-unreachable': Code 0, destination network unreachable; 'port-unreachable': Code 3, destination port unreachable; 'proto-unreachable': Code 2, destination protocol unreachable; 'route-failed': Code 5, source route failed;",
						},
						"icmpv6_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 type number",
						},
						"v6_any_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any ICMP type",
						},
						"special_v6_type": {
							Type: schema.TypeString, Optional: true, Description: "'dest-unreachable': Type 1, destination unreachable; 'echo-reply': Type 129, echo reply; 'echo-request': Type 128, echo request; 'packet-too-big': Type 2, packet too big; 'param-prob': Type 4, parameter problem; 'time-exceeded': Type 3, time exceeded;",
						},
						"v6_any_code": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any ICMPv6 code",
						},
						"icmpv6_code": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 code number",
						},
						"special_v6_code": {
							Type: schema.TypeString, Optional: true, Description: "'addr-unreachable': Code 3, address unreachable; 'admin-prohibited': Code 1, admin prohibited; 'no-route': Code 0, no route to destination; 'not-neighbour': Code 2, not neighbor; 'port-unreachable': Code 4, destination port unreachable;",
						},
						"source": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Source Port Information",
						},
						"eq_src": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets on a given source port (port number)",
						},
						"gt_src": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a greater source port number",
						},
						"lt_src": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a lower source port number",
						},
						"range_src": {
							Type: schema.TypeInt, Optional: true, Description: "match only packets in the range of source port numbers (Starting Port Number)",
						},
						"port_num_end_src": {
							Type: schema.TypeInt, Optional: true, Description: "Ending Source Port Number",
						},
						"eq_dst": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets on a given destination port (port number)",
						},
						"gt_dst": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a greater destination port number",
						},
						"lt_dst": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a lesser destination port number",
						},
						"range_dst": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets in the range of destination port numbers (Starting Destination Port Number)",
						},
						"port_num_end_dst": {
							Type: schema.TypeInt, Optional: true, Description: "Ending Destination Port Number",
						},
						"alg": {
							Type: schema.TypeString, Optional: true, Description: "'FTP': Specify FTP ALG port range; 'TFTP': Specify TFTP ALG port range; 'SIP': Specify SIP ALG port range; 'DNS': Specify DNS ALG port range; 'PPTP': Specify PPTP ALG port range; 'RTSP': Specify RTSP ALG port range;",
						},
					},
				},
			},
			"svc_name": {
				Type: schema.TypeString, Required: true, Description: "Service Object Group Name",
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
func resourceObjectGroupServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupServiceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupService(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceObjectGroupServiceRead(ctx, d, meta)
	}
	return diags
}

func resourceObjectGroupServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupServiceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupService(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceObjectGroupServiceRead(ctx, d, meta)
	}
	return diags
}
func resourceObjectGroupServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupServiceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupService(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceObjectGroupServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupServiceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupService(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceObjectGroupServiceRules(d []interface{}) []edpt.ObjectGroupServiceRules {

	count1 := len(d)
	ret := make([]edpt.ObjectGroupServiceRules, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ObjectGroupServiceRules
		oi.SeqNum = in["seq_num"].(int)
		oi.ProtocolId = in["protocol_id"].(int)
		oi.TcpUdp = in["tcp_udp"].(string)
		oi.Icmp = in["icmp"].(int)
		oi.Icmpv6 = in["icmpv6"].(int)
		oi.IcmpType = in["icmp_type"].(int)
		oi.AnyType = in["any_type"].(int)
		oi.SpecialType = in["special_type"].(string)
		oi.AnyCode = in["any_code"].(int)
		oi.IcmpCode = in["icmp_code"].(int)
		oi.SpecialCode = in["special_code"].(string)
		oi.Icmpv6Type = in["icmpv6_type"].(int)
		oi.V6AnyType = in["v6_any_type"].(int)
		oi.SpecialV6Type = in["special_v6_type"].(string)
		oi.V6AnyCode = in["v6_any_code"].(int)
		oi.Icmpv6Code = in["icmpv6_code"].(int)
		oi.SpecialV6Code = in["special_v6_code"].(string)
		oi.Source = in["source"].(int)
		oi.EqSrc = in["eq_src"].(int)
		oi.GtSrc = in["gt_src"].(int)
		oi.LtSrc = in["lt_src"].(int)
		oi.RangeSrc = in["range_src"].(int)
		oi.PortNumEndSrc = in["port_num_end_src"].(int)
		oi.EqDst = in["eq_dst"].(int)
		oi.GtDst = in["gt_dst"].(int)
		oi.LtDst = in["lt_dst"].(int)
		oi.RangeDst = in["range_dst"].(int)
		oi.PortNumEndDst = in["port_num_end_dst"].(int)
		oi.Alg = in["alg"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointObjectGroupService(d *schema.ResourceData) edpt.ObjectGroupService {
	var ret edpt.ObjectGroupService
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.Rules = getSliceObjectGroupServiceRules(d.Get("rules").([]interface{}))
	ret.Inst.SvcName = d.Get("svc_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
