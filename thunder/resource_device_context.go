package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceContext() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_device_context`: Switch context to a particular device for VCS config\n\n__PLACEHOLDER__",
		CreateContext: resourceDeviceContextCreate,
		UpdateContext: resourceDeviceContextUpdate,
		ReadContext:   resourceDeviceContextRead,
		DeleteContext: resourceDeviceContextDelete,

		Schema: map[string]*schema.Schema{
			"device_id": {
				Type: schema.TypeInt, Optional: true, Description: "Device ID",
			},
		},
	}
}
func resourceDeviceContextCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeviceContextCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeviceContext(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeviceContextRead(ctx, d, meta)
	}
	return diags
}

func resourceDeviceContextUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeviceContextUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeviceContext(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeviceContextRead(ctx, d, meta)
	}
	return diags
}
func resourceDeviceContextDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeviceContextDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeviceContext(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeviceContextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeviceContextRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeviceContext(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeviceContext(d *schema.ResourceData) edpt.DeviceContext {
	var ret edpt.DeviceContext
	ret.Inst.DeviceId = d.Get("device_id").(int)
	return ret
}
