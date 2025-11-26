package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorDebug() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_debug`: Enable store and replay for the entitites\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorDebugCreate,
		UpdateContext: resourceVisibilityMonitorDebugUpdate,
		ReadContext:   resourceVisibilityMonitorDebugRead,
		DeleteContext: resourceVisibilityMonitorDebugDelete,

		Schema: map[string]*schema.Schema{
			"debug_ip_addr": {
				Type: schema.TypeString, Required: true, Description: "Specify source/dest ip addr",
			},
			"debug_port": {
				Type: schema.TypeInt, Required: true, Description: "Specify port",
			},
			"debug_protocol": {
				Type: schema.TypeString, Required: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityMonitorDebugCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDebugCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDebug(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorDebugRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorDebugUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDebugUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDebug(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorDebugRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorDebugDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDebugDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDebug(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorDebugRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDebugRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorDebug(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorDebug(d *schema.ResourceData) edpt.VisibilityMonitorDebug {
	var ret edpt.VisibilityMonitorDebug
	ret.Inst.DebugIpAddr = d.Get("debug_ip_addr").(string)
	ret.Inst.DebugPort = d.Get("debug_port").(int)
	ret.Inst.DebugProtocol = d.Get("debug_protocol").(string)
	//omit uuid
	return ret
}
