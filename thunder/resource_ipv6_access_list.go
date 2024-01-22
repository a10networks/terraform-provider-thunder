package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6AccessList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_access_list`: Configure a IPv6 Access List\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6AccessListCreate,
		UpdateContext: resourceIpv6AccessListUpdate,
		ReadContext:   resourceIpv6AccessListRead,
		DeleteContext: resourceIpv6AccessListDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Named Access List",
			},
			"rules": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Number",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'deny': Deny; 'permit': Permit; 'l3-vlan-fwd-disable': Disable L3 forwarding between VLANs;",
						},
						"remark": {
							Type: schema.TypeString, Optional: true, Description: "Access list entry comment (Notes for this ACL)",
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Internet Control Message Protocol",
						},
						"tcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "protocol TCP",
						},
						"udp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "protocol UDP",
						},
						"ipv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any Internet Protocol",
						},
						"service_obj_group": {
							Type: schema.TypeString, Optional: true, Description: "Service object group (Source object group name)",
						},
						"geo_location": {
							Type: schema.TypeString, Optional: true, Description: "Specify geo-location name",
						},
						"icmp_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP type number",
						},
						"any_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any ICMP type",
						},
						"special_type": {
							Type: schema.TypeString, Optional: true, Description: "'echo-reply': Type 129, echo reply; 'echo-request': help Type 128, echo request; 'packet-too-big': Type 2, packet too big; 'param-prob': Type 4, parameter problem; 'time-exceeded': Type 3, time exceeded; 'dest-unreachable': Type 1, destination unreachable;",
						},
						"any_code": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any ICMP code",
						},
						"icmp_code": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP code number",
						},
						"special_code": {
							Type: schema.TypeString, Optional: true, Description: "'addr-unreachable': Code 3, address unreachable; 'admin-prohibited': Code 1, admin prohibited; 'no-route': Code 0, no route to destination; 'not-neighbour': Code 2, not neighbor; 'port-unreachable': Code 4, destination port unreachable;",
						},
						"src_any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any source host",
						},
						"src_host": {
							Type: schema.TypeString, Optional: true, Description: "A single source host (Host address)",
						},
						"src_subnet": {
							Type: schema.TypeString, Optional: true, Description: "Source Address",
						},
						"src_object_group": {
							Type: schema.TypeString, Optional: true, Description: "Network object group (Source network object group name)",
						},
						"src_eq": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets on a given source port (port number)",
						},
						"src_gt": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a greater port number",
						},
						"src_lt": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a lower port number",
						},
						"src_range": {
							Type: schema.TypeInt, Optional: true, Description: "match only packets in the range of port numbers (Starting Port Number)",
						},
						"src_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ending Port Number",
						},
						"dst_any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any destination host",
						},
						"dst_host": {
							Type: schema.TypeString, Optional: true, Description: "A single destination host (Host address)",
						},
						"dst_subnet": {
							Type: schema.TypeString, Optional: true, Description: "Destination Address",
						},
						"dst_object_group": {
							Type: schema.TypeString, Optional: true, Description: "Destination network object group name",
						},
						"dst_eq": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets on a given destination port (port number)",
						},
						"dst_gt": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a greater port number",
						},
						"dst_lt": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets with a lesser port number",
						},
						"dst_range": {
							Type: schema.TypeInt, Optional: true, Description: "Match only packets in the range of port numbers (Starting Destination Port Number)",
						},
						"dst_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Edning Destination Port Number",
						},
						"fragments": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP fragments",
						},
						"vlan": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
						},
						"ethernet": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
						},
						"trunk": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet trunk (trunk number)",
						},
						"dscp": {
							Type: schema.TypeInt, Optional: true, Description: "DSCP",
						},
						"established": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP established",
						},
						"acl_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log matches against this entry",
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
func resourceIpv6AccessListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AccessListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6AccessList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6AccessListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6AccessListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AccessListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6AccessList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6AccessListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6AccessListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AccessListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6AccessList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6AccessListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6AccessListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6AccessList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpv6AccessListRules(d []interface{}) []edpt.Ipv6AccessListRules {

	count1 := len(d)
	ret := make([]edpt.Ipv6AccessListRules, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Ipv6AccessListRules
		oi.SeqNum = in["seq_num"].(int)
		oi.Action = in["action"].(string)
		oi.Remark = in["remark"].(string)
		oi.Icmp = in["icmp"].(int)
		oi.Tcp = in["tcp"].(int)
		oi.Udp = in["udp"].(int)
		oi.Ipv6 = in["ipv6"].(int)
		oi.ServiceObjGroup = in["service_obj_group"].(string)
		oi.GeoLocation = in["geo_location"].(string)
		oi.IcmpType = in["icmp_type"].(int)
		oi.AnyType = in["any_type"].(int)
		oi.SpecialType = in["special_type"].(string)
		oi.AnyCode = in["any_code"].(int)
		oi.IcmpCode = in["icmp_code"].(int)
		oi.SpecialCode = in["special_code"].(string)
		oi.SrcAny = in["src_any"].(int)
		oi.SrcHost = in["src_host"].(string)
		oi.SrcSubnet = in["src_subnet"].(string)
		oi.SrcObjectGroup = in["src_object_group"].(string)
		oi.SrcEq = in["src_eq"].(int)
		oi.SrcGt = in["src_gt"].(int)
		oi.SrcLt = in["src_lt"].(int)
		oi.SrcRange = in["src_range"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		oi.DstAny = in["dst_any"].(int)
		oi.DstHost = in["dst_host"].(string)
		oi.DstSubnet = in["dst_subnet"].(string)
		oi.DstObjectGroup = in["dst_object_group"].(string)
		oi.DstEq = in["dst_eq"].(int)
		oi.DstGt = in["dst_gt"].(int)
		oi.DstLt = in["dst_lt"].(int)
		oi.DstRange = in["dst_range"].(int)
		oi.DstPortEnd = in["dst_port_end"].(int)
		oi.Fragments = in["fragments"].(int)
		oi.Vlan = in["vlan"].(int)
		oi.Ethernet = in["ethernet"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Dscp = in["dscp"].(int)
		oi.Established = in["established"].(int)
		oi.AclLog = in["acl_log"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpv6AccessList(d *schema.ResourceData) edpt.Ipv6AccessList {
	var ret edpt.Ipv6AccessList
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Rules = getSliceIpv6AccessListRules(d.Get("rules").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
