package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebugCapture() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug_capture`: Dump packets\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugCaptureCreate,
		UpdateContext: resourceAxdebugCaptureUpdate,
		ReadContext:   resourceAxdebugCaptureRead,
		DeleteContext: resourceAxdebugCaptureDelete,

		Schema: map[string]*schema.Schema{
			"brief": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print basic packet information",
			},
			"current_slot": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only for current-slot of chassis",
			},
			"detail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include packet payload",
			},
			"no_stop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Non-stop execution",
			},
			"save": {
				Type: schema.TypeString, Optional: true, Description: "Save packets into file (Specify filename to save packets)",
			},
		},
	}
}
func resourceAxdebugCaptureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugCaptureCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugCapture(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugCaptureRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugCaptureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugCaptureUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugCapture(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugCaptureRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugCaptureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugCaptureDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugCapture(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugCaptureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugCaptureRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugCapture(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAxdebugCapture(d *schema.ResourceData) edpt.AxdebugCapture {
	var ret edpt.AxdebugCapture
	ret.Inst.Brief = d.Get("brief").(int)
	ret.Inst.CurrentSlot = d.Get("current_slot").(int)
	ret.Inst.Detail = d.Get("detail").(int)
	ret.Inst.NoStop = d.Get("no_stop").(int)
	ret.Inst.Save = d.Get("save").(string)
	return ret
}
