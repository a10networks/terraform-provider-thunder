package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorSecondaryMonitorDeleteDebugFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_secondary_monitor_delete_debug_file`: Delete the debug entity file\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileCreate,
		UpdateContext: resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileUpdate,
		ReadContext:   resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileRead,
		DeleteContext: resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileDelete,

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
func resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDeleteDebugFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDeleteDebugFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDeleteDebugFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDeleteDebugFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDeleteDebugFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorSecondaryMonitorDeleteDebugFile(d *schema.ResourceData) edpt.VisibilityMonitorSecondaryMonitorDeleteDebugFile {
	var ret edpt.VisibilityMonitorSecondaryMonitorDeleteDebugFile
	ret.Inst.DebugIpAddr = d.Get("debug_ip_addr").(string)
	ret.Inst.DebugPort = d.Get("debug_port").(int)
	ret.Inst.DebugProtocol = d.Get("debug_protocol").(string)
	return ret
}
