package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RouteSourceLif() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_route_source_lif`: Lif interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RouteSourceLifCreate,
		UpdateContext: resourceIpv6RouteSourceLifUpdate,
		ReadContext:   resourceIpv6RouteSourceLifRead,
		DeleteContext: resourceIpv6RouteSourceLifDelete,

		Schema: map[string]*schema.Schema{
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Lif interface name",
			},
			"nexthop_ip": {
				Type: schema.TypeString, Required: true, Description: "Nexthop IPv6 address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6RouteSourceLifCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceLifCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceLif(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteSourceLifRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RouteSourceLifUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceLifUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceLif(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RouteSourceLifRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RouteSourceLifDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceLifDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceLif(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RouteSourceLifRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RouteSourceLifRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RouteSourceLif(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RouteSourceLif(d *schema.ResourceData) edpt.Ipv6RouteSourceLif {
	var ret edpt.Ipv6RouteSourceLif
	ret.Inst.Ifname = d.Get("ifname").(string)
	ret.Inst.NexthopIp = d.Get("nexthop_ip").(string)
	//omit uuid
	return ret
}
