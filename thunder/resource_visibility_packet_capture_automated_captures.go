package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureAutomatedCaptures() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_automated_captures`: Predefined set of automated captures\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureAutomatedCapturesCreate,
		UpdateContext: resourceVisibilityPacketCaptureAutomatedCapturesUpdate,
		ReadContext:   resourceVisibilityPacketCaptureAutomatedCapturesRead,
		DeleteContext: resourceVisibilityPacketCaptureAutomatedCapturesDelete,

		Schema: map[string]*schema.Schema{
			"slb_port_tmpl_error_code_return_inc": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Trigger capture when there is 3xx or 4xx or 5xx responses from server",
			},
			"slb_port_tmpl_high_error_code_return": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Trigger capture when there is high number of 3xx or 4xx or 5xx responses from server",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureAutomatedCapturesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureAutomatedCapturesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureAutomatedCaptures(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureAutomatedCapturesRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureAutomatedCapturesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureAutomatedCapturesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureAutomatedCaptures(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureAutomatedCapturesRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureAutomatedCapturesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureAutomatedCapturesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureAutomatedCaptures(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureAutomatedCapturesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureAutomatedCapturesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureAutomatedCaptures(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureAutomatedCaptures(d *schema.ResourceData) edpt.VisibilityPacketCaptureAutomatedCaptures {
	var ret edpt.VisibilityPacketCaptureAutomatedCaptures
	ret.Inst.Slb_port_tmpl_error_code_return_inc = d.Get("slb_port_tmpl_error_code_return_inc").(int)
	ret.Inst.Slb_port_tmpl_high_error_code_return = d.Get("slb_port_tmpl_high_error_code_return").(int)
	//omit uuid
	return ret
}
