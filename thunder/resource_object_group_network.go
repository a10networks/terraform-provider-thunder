package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceObjectGroupNetwork() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_object_group_network`: Configure Network Object Group\n\n__PLACEHOLDER__",
		CreateContext: resourceObjectGroupNetworkCreate,
		UpdateContext: resourceObjectGroupNetworkUpdate,
		ReadContext:   resourceObjectGroupNetworkRead,
		DeleteContext: resourceObjectGroupNetworkDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description of the object-group instance",
			},
			"ip_version": {
				Type: schema.TypeString, Optional: true, Description: "'v4': IPv4 rule; 'v6': IPv6 rule;",
			},
			"net_name": {
				Type: schema.TypeString, Required: true, Description: "Network Object Group Name",
			},
			"rules": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"host_v4": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Host Address",
						},
						"host_v6": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Host Address",
						},
						"ip_range_start": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Host Address start",
						},
						"ip_range_end": {
							Type: schema.TypeString, Optional: true, Description: "IPV4 Host address end",
						},
						"ipv6_range_start": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Host Address start",
						},
						"ipv6_range_end": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 Host address end",
						},
						"any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Any host",
						},
						"subnet": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Network Address",
						},
						"rev_subnet_mask": {
							Type: schema.TypeString, Optional: true, Description: "Network Mask. 0=apply, 255=ignore",
						},
						"fw_ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Network Address",
						},
						"ipv6_subnet": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Network Address",
						},
						"fw_ipv6_subnet": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Network Address",
						},
						"obj_network": {
							Type: schema.TypeString, Optional: true, Description: "Network Object",
						},
						"slb_server": {
							Type: schema.TypeString, Optional: true, Description: "Server",
						},
						"slb_vserver": {
							Type: schema.TypeString, Optional: true, Description: "Virtual Server",
						},
					},
				},
			},
			"usage": {
				Type: schema.TypeString, Optional: true, Default: "acl", Description: "'acl': Use for access-lists (default).; 'fw': Use for Firewall rule-set;",
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
func resourceObjectGroupNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupNetworkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupNetwork(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceObjectGroupNetworkRead(ctx, d, meta)
	}
	return diags
}

func resourceObjectGroupNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupNetworkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupNetwork(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceObjectGroupNetworkRead(ctx, d, meta)
	}
	return diags
}
func resourceObjectGroupNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupNetworkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupNetwork(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceObjectGroupNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceObjectGroupNetworkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointObjectGroupNetwork(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceObjectGroupNetworkRules(d []interface{}) []edpt.ObjectGroupNetworkRules {

	count1 := len(d)
	ret := make([]edpt.ObjectGroupNetworkRules, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ObjectGroupNetworkRules
		oi.SeqNum = in["seq_num"].(int)
		oi.HostV4 = in["host_v4"].(string)
		oi.HostV6 = in["host_v6"].(string)
		oi.IpRangeStart = in["ip_range_start"].(string)
		oi.IpRangeEnd = in["ip_range_end"].(string)
		oi.Ipv6RangeStart = in["ipv6_range_start"].(string)
		oi.Ipv6RangeEnd = in["ipv6_range_end"].(string)
		oi.Any = in["any"].(int)
		oi.Subnet = in["subnet"].(string)
		oi.RevSubnetMask = in["rev_subnet_mask"].(string)
		oi.FwIpv4Address = in["fw_ipv4_address"].(string)
		oi.Ipv6Subnet = in["ipv6_subnet"].(string)
		oi.FwIpv6Subnet = in["fw_ipv6_subnet"].(string)
		oi.ObjNetwork = in["obj_network"].(string)
		oi.SlbServer = in["slb_server"].(string)
		oi.SlbVserver = in["slb_vserver"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointObjectGroupNetwork(d *schema.ResourceData) edpt.ObjectGroupNetwork {
	var ret edpt.ObjectGroupNetwork
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.IpVersion = d.Get("ip_version").(string)
	ret.Inst.NetName = d.Get("net_name").(string)
	ret.Inst.Rules = getSliceObjectGroupNetworkRules(d.Get("rules").([]interface{}))
	ret.Inst.Usage = d.Get("usage").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
