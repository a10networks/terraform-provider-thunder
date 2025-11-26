package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorReplayDebugFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_replay_debug_file`: Replay the debug entity file\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorReplayDebugFileCreate,
		UpdateContext: resourceVisibilityMonitorReplayDebugFileUpdate,
		ReadContext:   resourceVisibilityMonitorReplayDebugFileRead,
		DeleteContext: resourceVisibilityMonitorReplayDebugFileDelete,

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
func resourceVisibilityMonitorReplayDebugFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorReplayDebugFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorReplayDebugFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorReplayDebugFileRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorReplayDebugFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorReplayDebugFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorReplayDebugFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorReplayDebugFileRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorReplayDebugFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorReplayDebugFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorReplayDebugFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorReplayDebugFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorReplayDebugFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorReplayDebugFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorReplayDebugFile(d *schema.ResourceData) edpt.VisibilityMonitorReplayDebugFile {
	var ret edpt.VisibilityMonitorReplayDebugFile
	ret.Inst.DebugIpAddr = d.Get("debug_ip_addr").(string)
	ret.Inst.DebugPort = d.Get("debug_port").(int)
	ret.Inst.DebugProtocol = d.Get("debug_protocol").(string)
	return ret
}
