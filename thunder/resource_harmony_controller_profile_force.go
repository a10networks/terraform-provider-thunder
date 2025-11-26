package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerProfileForce() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_harmony_controller_profile_force`: Harmony controller profile\n\n__PLACEHOLDER__",
		CreateContext: resourceHarmonyControllerProfileForceCreate,
		UpdateContext: resourceHarmonyControllerProfileForceUpdate,
		ReadContext:   resourceHarmonyControllerProfileForceRead,
		DeleteContext: resourceHarmonyControllerProfileForceDelete,

		Schema: map[string]*schema.Schema{
			"deregister": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "forcefully deregister thunder from harmony controller",
			},
		},
	}
}
func resourceHarmonyControllerProfileForceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileForceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileForce(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileForceRead(ctx, d, meta)
	}
	return diags
}

func resourceHarmonyControllerProfileForceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileForceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileForce(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileForceRead(ctx, d, meta)
	}
	return diags
}
func resourceHarmonyControllerProfileForceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileForceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileForce(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHarmonyControllerProfileForceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileForceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileForce(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHarmonyControllerProfileForce(d *schema.ResourceData) edpt.HarmonyControllerProfileForce {
	var ret edpt.HarmonyControllerProfileForce
	ret.Inst.Deregister = d.Get("deregister").(int)
	return ret
}
