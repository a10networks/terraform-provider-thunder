package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_monitor`: Monitors application ports\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortMonitorCreate,
		UpdateContext: resourceIpAppProtocolPortMonitorUpdate,
		ReadContext:   resourceIpAppProtocolPortMonitorRead,
		DeleteContext: resourceIpAppProtocolPortMonitorDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enables global port processing",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enables global port processing",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpAppProtocolPortMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpAppProtocolPortMonitor(d *schema.ResourceData) edpt.IpAppProtocolPortMonitor {
	var ret edpt.IpAppProtocolPortMonitor
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
