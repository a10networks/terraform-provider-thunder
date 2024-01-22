package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlmSend() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_glm_send`: Trigger GLM Connect\n\n__PLACEHOLDER__",
		CreateContext: resourceGlmSendCreate,
		UpdateContext: resourceGlmSendUpdate,
		ReadContext:   resourceGlmSendRead,
		DeleteContext: resourceGlmSendDelete,

		Schema: map[string]*schema.Schema{
			"ha_status": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send a ELM HA status request",
			},
			"license_request": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Immediately send a single GLM license request",
			},
		},
	}
}
func resourceGlmSendCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmSendCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmSend(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmSendRead(ctx, d, meta)
	}
	return diags
}

func resourceGlmSendUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmSendUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmSend(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmSendRead(ctx, d, meta)
	}
	return diags
}
func resourceGlmSendDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmSendDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmSend(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGlmSendRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmSendRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlmSend(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointGlmSend(d *schema.ResourceData) edpt.GlmSend {
	var ret edpt.GlmSend
	ret.Inst.HaStatus = d.Get("ha_status").(int)
	ret.Inst.LicenseRequest = d.Get("license_request").(int)
	return ret
}
