package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6BindingTableTunnelAddress() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_binding_table_tunnel_address`: Tunnel IPv6 Endpoint Address\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6BindingTableTunnelAddressCreate,
		UpdateContext: resourceCgnv6Lw4o6BindingTableTunnelAddressUpdate,
		ReadContext:   resourceCgnv6Lw4o6BindingTableTunnelAddressRead,
		DeleteContext: resourceCgnv6Lw4o6BindingTableTunnelAddressDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_tunnel_addr": {
				Type: schema.TypeString, Required: true, Description: "Tunnel IPv6 Endpoint Address",
			},
			"nat_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_nat_addr": {
							Type: schema.TypeString, Required: true, Description: "NAT IPv4 Address",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6Lw4o6BindingTableTunnelAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddress(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableTunnelAddressRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableTunnelAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddress(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableTunnelAddressRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6BindingTableTunnelAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddress(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableTunnelAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddress(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Lw4o6BindingTableTunnelAddressNatAddressList(d []interface{}) []edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressList
		oi.Ipv4NatAddr = in["ipv4_nat_addr"].(string)
		oi.UserTag = in["user_tag"].(string)
		oi.PortRangeList = getSliceCgnv6Lw4o6BindingTableTunnelAddressNatAddressListPortRangeList(in["port_range_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Lw4o6BindingTableTunnelAddressNatAddressListPortRangeList(d []interface{}) []edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressListPortRangeList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressListPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressListPortRangeList
		oi.PortStart = in["port_start"].(int)
		oi.PortEnd = in["port_end"].(int)
		oi.TunnelEndpointAddress = in["tunnel_endpoint_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6BindingTableTunnelAddress(d *schema.ResourceData) edpt.Cgnv6Lw4o6BindingTableTunnelAddress {
	var ret edpt.Cgnv6Lw4o6BindingTableTunnelAddress
	ret.Inst.Ipv6TunnelAddr = d.Get("ipv6_tunnel_addr").(string)
	ret.Inst.NatAddressList = getSliceCgnv6Lw4o6BindingTableTunnelAddressNatAddressList(d.Get("nat_address_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
