package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6OspfRoute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf_route`: OSPFv3 route information\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfRouteCreate,
		UpdateContext: resourceDebugIpv6OspfRouteUpdate,
		ReadContext:   resourceDebugIpv6OspfRouteRead,
		DeleteContext: resourceDebugIpv6OspfRouteDelete,

		Schema: map[string]*schema.Schema{
			"ase": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "External route calculation information",
			},
			"ia": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-Area route calculation information",
			},
			"install": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Route installation information",
			},
			"spf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SPF calculation information",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugIpv6OspfRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfRouteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfRoute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfRouteRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfRouteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfRoute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfRouteRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfRouteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfRoute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfRouteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6OspfRoute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugIpv6OspfRoute(d *schema.ResourceData) edpt.DebugIpv6OspfRoute {
	var ret edpt.DebugIpv6OspfRoute
	ret.Inst.Ase = d.Get("ase").(int)
	ret.Inst.Ia = d.Get("ia").(int)
	ret.Inst.Install = d.Get("install").(int)
	ret.Inst.Spf = d.Get("spf").(int)
	//omit uuid
	return ret
}
