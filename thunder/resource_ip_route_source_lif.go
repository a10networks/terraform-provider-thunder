package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpRouteSourceLif() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_route_source_lif`: Lif interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpRouteSourceLifCreate,
		UpdateContext: resourceIpRouteSourceLifUpdate,
		ReadContext:   resourceIpRouteSourceLifRead,
		DeleteContext: resourceIpRouteSourceLifDelete,

		Schema: map[string]*schema.Schema{
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Lif interface name",
			},
			"nexthop_ip": {
				Type: schema.TypeString, Required: true, Description: "Nexthop IP address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpRouteSourceLifCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceLifCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceLif(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteSourceLifRead(ctx, d, meta)
	}
	return diags
}

func resourceIpRouteSourceLifUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceLifUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceLif(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRouteSourceLifRead(ctx, d, meta)
	}
	return diags
}
func resourceIpRouteSourceLifDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceLifDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceLif(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpRouteSourceLifRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRouteSourceLifRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRouteSourceLif(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpRouteSourceLif(d *schema.ResourceData) edpt.IpRouteSourceLif {
	var ret edpt.IpRouteSourceLif
	ret.Inst.Ifname = d.Get("ifname").(string)
	ret.Inst.NexthopIp = d.Get("nexthop_ip").(string)
	//omit uuid
	return ret
}
