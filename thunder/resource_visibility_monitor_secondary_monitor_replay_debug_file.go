package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorSecondaryMonitorReplayDebugFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_secondary_monitor_replay_debug_file`: Replay the debug entity file\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorSecondaryMonitorReplayDebugFileCreate,
		UpdateContext: resourceVisibilityMonitorSecondaryMonitorReplayDebugFileUpdate,
		ReadContext:   resourceVisibilityMonitorSecondaryMonitorReplayDebugFileRead,
		DeleteContext: resourceVisibilityMonitorSecondaryMonitorReplayDebugFileDelete,

		Schema: map[string]*schema.Schema{
			"debug_ip_addr": {
				Type: schema.TypeString, Required: true, Description: "Specify source/dest ip addr",
			},
			"debug_port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify port",
			},
			"debug_protocol": {
				Type: schema.TypeString, Optional: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
			},
		},
	}
}
func resourceVisibilityMonitorSecondaryMonitorReplayDebugFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorReplayDebugFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorReplayDebugFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorReplayDebugFileRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorReplayDebugFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorReplayDebugFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorReplayDebugFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorReplayDebugFileRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorSecondaryMonitorReplayDebugFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorReplayDebugFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorReplayDebugFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorReplayDebugFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorReplayDebugFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorReplayDebugFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorSecondaryMonitorReplayDebugFile(d *schema.ResourceData) edpt.VisibilityMonitorSecondaryMonitorReplayDebugFile {
	var ret edpt.VisibilityMonitorSecondaryMonitorReplayDebugFile
	ret.Inst.DebugIpAddr = d.Get("debug_ip_addr").(string)
	ret.Inst.DebugPort = d.Get("debug_port").(int)
	ret.Inst.DebugProtocol = d.Get("debug_protocol").(string)
	return ret
}
