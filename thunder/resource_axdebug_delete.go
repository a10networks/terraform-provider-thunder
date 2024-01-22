package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebugDelete() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug_delete`: delete axdebug file\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugDeleteCreate,
		UpdateContext: resourceAxdebugDeleteUpdate,
		ReadContext:   resourceAxdebugDeleteRead,
		DeleteContext: resourceAxdebugDeleteDelete,

		Schema: map[string]*schema.Schema{
			"capture_file": {
				Type: schema.TypeString, Optional: true, Description: "Delete a capture file (Specify target filename to change)",
			},
			"config_file": {
				Type: schema.TypeString, Optional: true, Description: "Delete AXDebug config file (Specify target filename to change)",
			},
		},
	}
}
func resourceAxdebugDeleteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugDeleteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugDelete(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugDeleteRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugDeleteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugDeleteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugDelete(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugDeleteRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugDeleteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugDeleteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugDelete(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugDeleteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugDeleteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugDelete(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAxdebugDelete(d *schema.ResourceData) edpt.AxdebugDelete {
	var ret edpt.AxdebugDelete
	ret.Inst.CaptureFile = d.Get("capture_file").(string)
	ret.Inst.ConfigFile = d.Get("config_file").(string)
	return ret
}
