package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugPacketCapture() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_packet_capture`: Display live captures from packet-capture\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugPacketCaptureCreate,
		UpdateContext: resourceDebugPacketCaptureUpdate,
		ReadContext:   resourceDebugPacketCaptureRead,
		DeleteContext: resourceDebugPacketCaptureDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Specify capture-config. Default all (Capture config)",
			},
			"count1": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum packets to capture. Default is 10 (Specify maximum packet number. For unlimited, specify 0)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugPacketCaptureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketCaptureCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacketCapture(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPacketCaptureRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugPacketCaptureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketCaptureUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacketCapture(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPacketCaptureRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugPacketCaptureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketCaptureDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacketCapture(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugPacketCaptureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketCaptureRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacketCapture(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugPacketCapture(d *schema.ResourceData) edpt.DebugPacketCapture {
	var ret edpt.DebugPacketCapture
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Count1 = d.Get("count1").(int)
	//omit uuid
	return ret
}
