package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspfRoute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf_route`: OSPFv3 route information\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfRouteCreate,
		UpdateContext: resourceDebugOspfRouteUpdate,
		ReadContext:   resourceDebugOspfRouteRead,
		DeleteContext: resourceDebugOspfRouteDelete,

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
func resourceDebugOspfRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfRouteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfRoute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfRouteRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfRouteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfRoute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfRouteRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfRouteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfRoute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfRouteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspfRoute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugOspfRoute(d *schema.ResourceData) edpt.DebugOspfRoute {
	var ret edpt.DebugOspfRoute
	ret.Inst.Ase = d.Get("ase").(int)
	ret.Inst.Ia = d.Get("ia").(int)
	ret.Inst.Install = d.Get("install").(int)
	ret.Inst.Spf = d.Get("spf").(int)
	//omit uuid
	return ret
}
