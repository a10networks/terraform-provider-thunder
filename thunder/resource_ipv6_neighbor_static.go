package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NeighborStatic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_neighbor_static`: static IPv6 Neighbor commands\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6NeighborStaticCreate,
		UpdateContext: resourceIpv6NeighborStaticUpdate,
		ReadContext:   resourceIpv6NeighborStaticRead,
		DeleteContext: resourceIpv6NeighborStaticDelete,

		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Port Value)",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "IPV6 address",
			},
			"mac": {
				Type: schema.TypeString, Optional: true, Description: "MAC Address",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk group",
			},
			"tunnel": {
				Type: schema.TypeInt, Optional: true, Description: "Tunnel interface",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeInt, Required: true, Description: "VLAN ID",
			},
		},
	}
}
func resourceIpv6NeighborStaticCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NeighborStaticCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NeighborStatic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NeighborStaticRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6NeighborStaticUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NeighborStaticUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NeighborStatic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NeighborStaticRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6NeighborStaticDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NeighborStaticDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NeighborStatic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6NeighborStaticRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NeighborStaticRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NeighborStatic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6NeighborStatic(d *schema.ResourceData) edpt.Ipv6NeighborStatic {
	var ret edpt.Ipv6NeighborStatic
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.Mac = d.Get("mac").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	ret.Inst.Tunnel = d.Get("tunnel").(int)
	//omit uuid
	ret.Inst.Vlan = d.Get("vlan").(int)
	return ret
}
