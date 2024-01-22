package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_binding_table_tunnel_address_nat_address`: NAT IPv4 address\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressCreate,
		UpdateContext: resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressUpdate,
		ReadContext:   resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressRead,
		DeleteContext: resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_nat_addr": {
				Type: schema.TypeString, Required: true, Description: "NAT IPv4 Address",
			},
			"port_range_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_start": {
							Type: schema.TypeInt, Required: true, Description: "Single Port or Port Range Start",
						},
						"port_end": {
							Type: schema.TypeInt, Required: true, Description: "Port Range End",
						},
						"tunnel_endpoint_address": {
							Type: schema.TypeString, Required: true, Description: "Configure LW-4over6 IPIP Tunnel Endpoint Address (LW-4over6 Tunnel Endpoint Address)",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
			"ipv6_tunnel_addr": {
				Type: schema.TypeString, Required: true, Description: "Ipv6TunnelAddr",
			},
		},
	}
}
func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeList(d []interface{}) []edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeList
		oi.PortStart = in["port_start"].(int)
		oi.PortEnd = in["port_end"].(int)
		oi.TunnelEndpointAddress = in["tunnel_endpoint_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddress(d *schema.ResourceData) edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddress {
	var ret edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddress
	ret.Inst.Ipv4NatAddr = d.Get("ipv4_nat_addr").(string)
	ret.Inst.PortRangeList = getSliceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeList(d.Get("port_range_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Ipv6TunnelAddr = d.Get("ipv6_tunnel_addr").(string)
	return ret
}
