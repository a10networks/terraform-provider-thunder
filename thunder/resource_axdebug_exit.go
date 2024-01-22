package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebugExit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug_exit`: exit from axdebug and stop capture\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugExitCreate,
		UpdateContext: resourceAxdebugExitUpdate,
		ReadContext:   resourceAxdebugExitRead,
		DeleteContext: resourceAxdebugExitDelete,

		Schema: map[string]*schema.Schema{
			"stop_capture": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "stop capture traffic",
			},
		},
	}
}
func resourceAxdebugExitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugExitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugExit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugExitRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugExitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugExitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugExit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugExitRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugExitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugExitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugExit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugExitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugExitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugExit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAxdebugExit(d *schema.ResourceData) edpt.AxdebugExit {
	var ret edpt.AxdebugExit
	ret.Inst.StopCapture = d.Get("stop_capture").(int)
	return ret
}
