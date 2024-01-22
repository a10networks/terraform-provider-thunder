package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6BindingTable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_binding_table`: Configure LW-4over6 Binding Table\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6BindingTableCreate,
		UpdateContext: resourceCgnv6Lw4o6BindingTableUpdate,
		ReadContext:   resourceCgnv6Lw4o6BindingTableRead,
		DeleteContext: resourceCgnv6Lw4o6BindingTableDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "LW-4over6 Binding Table Name",
			},
			"tunnel_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_tunnel_addr": {
							Type: schema.TypeString, Required: true, Description: "Tunnel IPv6 Endpoint Address",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
		},
	}
}
func resourceCgnv6Lw4o6BindingTableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6BindingTableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Lw4o6BindingTableTunnelAddressList(d []interface{}) []edpt.Cgnv6Lw4o6BindingTableTunnelAddressList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6BindingTableTunnelAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6BindingTableTunnelAddressList
		oi.Ipv6TunnelAddr = in["ipv6_tunnel_addr"].(string)
		oi.UserTag = in["user_tag"].(string)
		oi.NatAddressList = getSliceCgnv6Lw4o6BindingTableTunnelAddressListNatAddressList(in["nat_address_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Lw4o6BindingTableTunnelAddressListNatAddressList(d []interface{}) []edpt.Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressList
		oi.Ipv4NatAddr = in["ipv4_nat_addr"].(string)
		oi.UserTag = in["user_tag"].(string)
		oi.PortRangeList = getSliceCgnv6Lw4o6BindingTableTunnelAddressListNatAddressListPortRangeList(in["port_range_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Lw4o6BindingTableTunnelAddressListNatAddressListPortRangeList(d []interface{}) []edpt.Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressListPortRangeList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressListPortRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressListPortRangeList
		oi.PortStart = in["port_start"].(int)
		oi.PortEnd = in["port_end"].(int)
		oi.TunnelEndpointAddress = in["tunnel_endpoint_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Lw4o6BindingTable(d *schema.ResourceData) edpt.Cgnv6Lw4o6BindingTable {
	var ret edpt.Cgnv6Lw4o6BindingTable
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TunnelAddressList = getSliceCgnv6Lw4o6BindingTableTunnelAddressList(d.Get("tunnel_address_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	return ret
}
