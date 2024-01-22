package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugTrafficCapture() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_traffic_capture`: Debug Traffic Capture Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugTrafficCaptureCreate,
		UpdateContext: resourceDebugTrafficCaptureUpdate,
		ReadContext:   resourceDebugTrafficCaptureRead,
		DeleteContext: resourceDebugTrafficCaptureDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Debug Traffic Capture",
			},
			"maximum_file_size": {
				Type: schema.TypeInt, Optional: true, Description: "Specify pcapng filesize in MB (default 100)",
			},
			"maximum_history_recordings": {
				Type: schema.TypeInt, Optional: true, Description: "Number of crashed samples (default 5)",
			},
			"record_direction_type": {
				Type: schema.TypeString, Optional: true, Description: "'ingress-only': Record all ingress traffic; 'all-direction': Record all ingress and egress traffic;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugTrafficCaptureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTrafficCaptureCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTrafficCapture(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugTrafficCaptureRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugTrafficCaptureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTrafficCaptureUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTrafficCapture(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugTrafficCaptureRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugTrafficCaptureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTrafficCaptureDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTrafficCapture(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugTrafficCaptureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTrafficCaptureRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTrafficCapture(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugTrafficCapture(d *schema.ResourceData) edpt.DebugTrafficCapture {
	var ret edpt.DebugTrafficCapture
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.MaximumFileSize = d.Get("maximum_file_size").(int)
	ret.Inst.MaximumHistoryRecordings = d.Get("maximum_history_recordings").(int)
	ret.Inst.RecordDirectionType = d.Get("record_direction_type").(string)
	//omit uuid
	return ret
}
