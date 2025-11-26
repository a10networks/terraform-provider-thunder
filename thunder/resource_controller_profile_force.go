package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProfileForce() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_profile_force`: Controller profile\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerProfileForceCreate,
		UpdateContext: resourceControllerProfileForceUpdate,
		ReadContext:   resourceControllerProfileForceRead,
		DeleteContext: resourceControllerProfileForceDelete,

		Schema: map[string]*schema.Schema{
			"deregister": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "forcefully deregister thunder from harmony controller",
			},
		},
	}
}
func resourceControllerProfileForceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileForceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileForce(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileForceRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerProfileForceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileForceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileForce(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileForceRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerProfileForceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileForceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileForce(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerProfileForceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileForceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileForce(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointControllerProfileForce(d *schema.ResourceData) edpt.ControllerProfileForce {
	var ret edpt.ControllerProfileForce
	ret.Inst.Deregister = d.Get("deregister").(int)
	return ret
}
