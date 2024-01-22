package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodRtsp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_rtsp`: RTSP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodRtspCreate,
		UpdateContext: resourceHealthMonitorMethodRtspUpdate,
		ReadContext:   resourceHealthMonitorMethodRtspRead,
		DeleteContext: resourceHealthMonitorMethodRtspDelete,

		Schema: map[string]*schema.Schema{
			"rtsp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RTSP type",
			},
			"rtsp_port": {
				Type: schema.TypeInt, Optional: true, Default: 554, Description: "Specify RTSP port, default is 554 (Port Number (default 554))",
			},
			"rtspurl": {
				Type: schema.TypeString, Optional: true, Description: "Specify URL string (Specify the path on the server)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodRtspCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRtspCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRtsp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodRtspRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodRtspUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRtspUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRtsp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodRtspRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodRtspDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRtspDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRtsp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodRtspRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodRtspRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodRtsp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodRtsp(d *schema.ResourceData) edpt.HealthMonitorMethodRtsp {
	var ret edpt.HealthMonitorMethodRtsp
	ret.Inst.Rtsp = d.Get("rtsp").(int)
	ret.Inst.RtspPort = d.Get("rtsp_port").(int)
	ret.Inst.Rtspurl = d.Get("rtspurl").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
