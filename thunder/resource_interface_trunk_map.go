package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkMap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_map`: Configure MAP on this interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkMapCreate,
		UpdateContext: resourceInterfaceTrunkMapUpdate,
		ReadContext:   resourceInterfaceTrunkMapRead,
		DeleteContext: resourceInterfaceTrunkMapDelete,

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
func resourceInterfaceTrunkMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkMapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkMap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkMapRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkMapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkMap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkMapRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkMapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkMap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkMapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkMap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceTrunkMap(d *schema.ResourceData) edpt.InterfaceTrunkMap {
	var ret edpt.InterfaceTrunkMap
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.MapTInside = d.Get("map_t_inside").(int)
	ret.Inst.MapTOutside = d.Get("map_t_outside").(int)
	ret.Inst.Outside = d.Get("outside").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
