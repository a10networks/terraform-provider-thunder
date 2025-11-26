package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigReplace() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_config_replace`: Options for config-replace-{start/end}\n\n__PLACEHOLDER__",
		CreateContext: resourceConfigReplaceCreate,
		UpdateContext: resourceConfigReplaceUpdate,
		ReadContext:   resourceConfigReplaceRead,
		DeleteContext: resourceConfigReplaceDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'skip': Skip object and its children; 'accept': Skip siblings of object and its parents;",
			},
			"filter": {
				Type: schema.TypeString, Optional: true, Description: "Set scope of objects by specifying object class",
			},
			"gslb_syncing_off": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Turn off syncing changes to GSLB Group",
			},
			"ignore_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore errors and continue to apply changes",
			},
			"log_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log errors",
			},
		},
	}
}
func resourceConfigReplaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigReplaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigReplace(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConfigReplaceRead(ctx, d, meta)
	}
	return diags
}

func resourceConfigReplaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigReplaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigReplace(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConfigReplaceRead(ctx, d, meta)
	}
	return diags
}
func resourceConfigReplaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigReplaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigReplace(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceConfigReplaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigReplaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigReplace(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointConfigReplace(d *schema.ResourceData) edpt.ConfigReplace {
	var ret edpt.ConfigReplace
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Filter = d.Get("filter").(string)
	ret.Inst.GslbSyncingOff = d.Get("gslb_syncing_off").(int)
	ret.Inst.IgnoreError = d.Get("ignore_error").(int)
	ret.Inst.LogError = d.Get("log_error").(int)
	return ret
}
