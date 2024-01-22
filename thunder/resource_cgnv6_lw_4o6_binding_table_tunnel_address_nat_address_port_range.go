package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lw_4o6_binding_table_tunnel_address_nat_address_port_range`: Single Port or Port Range Start\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeCreate,
		UpdateContext: resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeUpdate,
		ReadContext:   resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeRead,
		DeleteContext: resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeDelete,

		Schema: map[string]*schema.Schema{
			"port_end": {
				Type: schema.TypeInt, Required: true, Description: "Port Range End",
			},
			"port_start": {
				Type: schema.TypeInt, Required: true, Description: "Single Port or Port Range Start",
			},
			"tunnel_endpoint_address": {
				Type: schema.TypeString, Required: true, Description: "Configure LW-4over6 IPIP Tunnel Endpoint Address (LW-4over6 Tunnel Endpoint Address)",
			},
			"ipv4_nat_addr": {
				Type: schema.TypeString, Required: true, Description: "Ipv4NatAddr",
			},
			"ipv6_tunnel_addr": {
				Type: schema.TypeString, Required: true, Description: "Ipv6TunnelAddr",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange(d *schema.ResourceData) edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange {
	var ret edpt.Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange
	ret.Inst.PortEnd = d.Get("port_end").(int)
	ret.Inst.PortStart = d.Get("port_start").(int)
	ret.Inst.TunnelEndpointAddress = d.Get("tunnel_endpoint_address").(string)
	ret.Inst.Ipv4NatAddr = d.Get("ipv4_nat_addr").(string)
	ret.Inst.Ipv6TunnelAddr = d.Get("ipv6_tunnel_addr").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
