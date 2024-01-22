package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerConfigReplace() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_harmony_controller_config_replace`: Enable config-replace-mode\n\n__PLACEHOLDER__",
		CreateContext: resourceHarmonyControllerConfigReplaceCreate,
		UpdateContext: resourceHarmonyControllerConfigReplaceUpdate,
		ReadContext:   resourceHarmonyControllerConfigReplaceRead,
		DeleteContext: resourceHarmonyControllerConfigReplaceDelete,

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
func resourceHarmonyControllerConfigReplaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerConfigReplaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerConfigReplace(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerConfigReplaceRead(ctx, d, meta)
	}
	return diags
}

func resourceHarmonyControllerConfigReplaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerConfigReplaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerConfigReplace(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerConfigReplaceRead(ctx, d, meta)
	}
	return diags
}
func resourceHarmonyControllerConfigReplaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerConfigReplaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerConfigReplace(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHarmonyControllerConfigReplaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerConfigReplaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerConfigReplace(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHarmonyControllerConfigReplace(d *schema.ResourceData) edpt.HarmonyControllerConfigReplace {
	var ret edpt.HarmonyControllerConfigReplace
	ret.Inst.Status = d.Get("status").(string)
	//omit uuid
	return ret
}
