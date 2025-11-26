package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerConfigReplace() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_config_replace`: Enable config-replace-mode\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerConfigReplaceCreate,
		UpdateContext: resourceControllerConfigReplaceUpdate,
		ReadContext:   resourceControllerConfigReplaceRead,
		DeleteContext: resourceControllerConfigReplaceDelete,

		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable config replace mode; 'disable': disable config replace mode;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceControllerConfigReplaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerConfigReplaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerConfigReplace(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerConfigReplaceRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerConfigReplaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerConfigReplaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerConfigReplace(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerConfigReplaceRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerConfigReplaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerConfigReplaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerConfigReplace(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerConfigReplaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerConfigReplaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerConfigReplace(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointControllerConfigReplace(d *schema.ResourceData) edpt.ControllerConfigReplace {
	var ret edpt.ControllerConfigReplace
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	return ret
}
