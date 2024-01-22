package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_bgp_address_family_ipv6_neighbor_ethernet_neighbor_ipv6`: Specify an ethernet unnumbered neighbor\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Delete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
			},
			"peer_group_name": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"as_number": {
				Type: schema.TypeString, Required: true, Description: "AsNumber",
			},
		},
	}
}
func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(d *schema.ResourceData) edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 {
	var ret edpt.RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.PeerGroupName = d.Get("peer_group_name").(string)
	//omit uuid
	ret.Inst.AsNumber = d.Get("as_number").(string)
	return ret
}
