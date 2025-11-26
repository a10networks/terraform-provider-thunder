package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorSecondaryMonitorDebug() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor_secondary_monitor_debug`: Enable store and replay for the entitites\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorSecondaryMonitorDebugCreate,
		UpdateContext: resourceVisibilityMonitorSecondaryMonitorDebugUpdate,
		ReadContext:   resourceVisibilityMonitorSecondaryMonitorDebugRead,
		DeleteContext: resourceVisibilityMonitorSecondaryMonitorDebugDelete,

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
func resourceVisibilityMonitorSecondaryMonitorDebugCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDebugCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDebug(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorDebugRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorDebugUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDebugUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDebug(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorSecondaryMonitorDebugRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorSecondaryMonitorDebugDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDebugDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDebug(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorSecondaryMonitorDebugRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorSecondaryMonitorDebugRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorSecondaryMonitorDebug(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityMonitorSecondaryMonitorDebug(d *schema.ResourceData) edpt.VisibilityMonitorSecondaryMonitorDebug {
	var ret edpt.VisibilityMonitorSecondaryMonitorDebug
	ret.Inst.DebugIpAddr = d.Get("debug_ip_addr").(string)
	ret.Inst.DebugPort = d.Get("debug_port").(int)
	ret.Inst.DebugProtocol = d.Get("debug_protocol").(string)
	//omit uuid
	return ret
}
