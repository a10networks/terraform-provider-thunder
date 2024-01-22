package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnelMap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_tunnel_map`: Configure MAP on this interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTunnelMapCreate,
		UpdateContext: resourceInterfaceTunnelMapUpdate,
		ReadContext:   resourceInterfaceTunnelMapRead,
		DeleteContext: resourceInterfaceTunnelMapDelete,

		Schema: map[string]*schema.Schema{
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP inside interface (connected to MAP domains)",
			},
			"map_t_inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP inside interface (connected to MAP domains)",
			},
			"map_t_outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP outside interface",
			},
			"outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure MAP outside interface",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTunnelMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelMapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelMap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelMapRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTunnelMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelMapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelMap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelMapRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTunnelMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelMapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelMap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTunnelMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelMapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelMap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceTunnelMap(d *schema.ResourceData) edpt.InterfaceTunnelMap {
	var ret edpt.InterfaceTunnelMap
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.MapTInside = d.Get("map_t_inside").(int)
	ret.Inst.MapTOutside = d.Get("map_t_outside").(int)
	ret.Inst.Outside = d.Get("outside").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
