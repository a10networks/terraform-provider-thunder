package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateTcpProgressionTrackingProfiling() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_tcp_progression_tracking_profiling`: Configure and enable TCP Progression Tracking Profiling\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateTcpProgressionTrackingProfilingCreate,
		UpdateContext: resourceDdosZoneTemplateTcpProgressionTrackingProfilingUpdate,
		ReadContext:   resourceDdosZoneTemplateTcpProgressionTrackingProfilingRead,
		DeleteContext: resourceDdosZoneTemplateTcpProgressionTrackingProfilingDelete,

		Schema: map[string]*schema.Schema{
			"profiling_connection_life_model": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for connection model",
			},
			"profiling_request_response_model": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for request response model",
			},
			"profiling_time_window_model": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto-config progression tracking learning for time window model",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"tcp_name": {
				Type: schema.TypeString, Required: true, Description: "Tcp_name",
			},
		},
	}
}
func resourceDdosZoneTemplateTcpProgressionTrackingProfilingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingProfilingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingProfiling(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpProgressionTrackingProfilingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateTcpProgressionTrackingProfilingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingProfilingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingProfiling(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateTcpProgressionTrackingProfilingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateTcpProgressionTrackingProfilingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingProfilingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingProfiling(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateTcpProgressionTrackingProfilingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateTcpProgressionTrackingProfilingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateTcpProgressionTrackingProfiling(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateTcpProgressionTrackingProfiling(d *schema.ResourceData) edpt.DdosZoneTemplateTcpProgressionTrackingProfiling {
	var ret edpt.DdosZoneTemplateTcpProgressionTrackingProfiling
	ret.Inst.ProfilingConnectionLifeModel = d.Get("profiling_connection_life_model").(int)
	ret.Inst.ProfilingRequestResponseModel = d.Get("profiling_request_response_model").(int)
	ret.Inst.ProfilingTimeWindowModel = d.Get("profiling_time_window_model").(int)
	//omit uuid
	ret.Inst.Tcp_name = d.Get("tcp_name").(string)
	return ret
}
