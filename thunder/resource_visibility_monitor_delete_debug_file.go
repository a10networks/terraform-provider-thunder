package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorDeleteDebugFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_delete_debug_file`: Delete the debug entity file\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorDeleteDebugFileCreate,
		UpdateContext: resourceVisibilityMonitorDeleteDebugFileUpdate,
		ReadContext:   resourceVisibilityMonitorDeleteDebugFileRead,
		DeleteContext: resourceVisibilityMonitorDeleteDebugFileDelete,

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
func resourceVisibilityMonitorDeleteDebugFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDeleteDebugFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDeleteDebugFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorDeleteDebugFileRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorDeleteDebugFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDeleteDebugFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDeleteDebugFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorDeleteDebugFileRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorDeleteDebugFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDeleteDebugFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDeleteDebugFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorDeleteDebugFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDeleteDebugFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDeleteDebugFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorDeleteDebugFile(d *schema.ResourceData) edpt.VisibilityMonitorDeleteDebugFile {
	var ret edpt.VisibilityMonitorDeleteDebugFile
	ret.Inst.DebugIpAddr = d.Get("debug_ip_addr").(string)
	ret.Inst.DebugPort = d.Get("debug_port").(int)
	ret.Inst.DebugProtocol = d.Get("debug_protocol").(string)
	return ret
}
