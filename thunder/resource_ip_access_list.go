package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAccessList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_access_list`: Configure Access List\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAccessListCreate,
		UpdateContext: resourceIpAccessListUpdate,
		ReadContext:   resourceIpAccessListRead,
		DeleteContext: resourceIpAccessListDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "IP Access List Name. Does not support name as digits or start with digit.",
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
						"ip": {
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
						"src_any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any source host",
						},
						"src_host": {
							Type: schema.TypeString, Optional: true, Description: "A single source host (Host address)",
						},
						"src_subnet": {
							Type: schema.TypeString, Optional: true, Description: "Source Address",
						},
						"src_mask": {
							Type: schema.TypeString, Optional: true, Description: "Source Mask 0=apply 255=ignore",
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
						"dst_mask": {
							Type: schema.TypeString, Optional: true, Description: "Destination Mask 0=apply 255=ignore",
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
						"transparent_session_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only log transparent sessions",
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
func resourceIpAccessListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAccessListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAccessList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAccessListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAccessListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAccessListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAccessList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAccessListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAccessListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAccessListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAccessList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAccessListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAccessListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAccessList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpAccessListRules(d []interface{}) []edpt.IpAccessListRules {

	count1 := len(d)
	ret := make([]edpt.IpAccessListRules, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAccessListRules
		oi.SeqNum = in["seq_num"].(int)
		oi.Action = in["action"].(string)
		oi.Remark = in["remark"].(string)
		oi.Icmp = in["icmp"].(int)
		oi.Tcp = in["tcp"].(int)
		oi.Udp = in["udp"].(int)
		oi.Ip = in["ip"].(int)
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
		oi.SrcMask = in["src_mask"].(string)
		oi.SrcObjectGroup = in["src_object_group"].(string)
		oi.SrcEq = in["src_eq"].(int)
		oi.SrcGt = in["src_gt"].(int)
		oi.SrcLt = in["src_lt"].(int)
		oi.SrcRange = in["src_range"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		oi.DstAny = in["dst_any"].(int)
		oi.DstHost = in["dst_host"].(string)
		oi.DstSubnet = in["dst_subnet"].(string)
		oi.DstMask = in["dst_mask"].(string)
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
		oi.TransparentSessionOnly = in["transparent_session_only"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAccessList(d *schema.ResourceData) edpt.IpAccessList {
	var ret edpt.IpAccessList
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Rules = getSliceIpAccessListRules(d.Get("rules").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
