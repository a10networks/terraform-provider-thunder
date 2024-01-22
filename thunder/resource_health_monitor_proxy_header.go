package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorProxyHeader() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_proxy_header`: Proxy header command\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorProxyHeaderCreate,
		UpdateContext: resourceHealthMonitorProxyHeaderUpdate,
		ReadContext:   resourceHealthMonitorProxyHeaderRead,
		DeleteContext: resourceHealthMonitorProxyHeaderDelete,

		Schema: map[string]*schema.Schema{
			"proxy_header_ver": {
				Type: schema.TypeString, Optional: true, Description: "'v1': version 1; 'v2': version 2;  (version number)",
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
func resourceHealthMonitorProxyHeaderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorProxyHeaderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorProxyHeader(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorProxyHeaderRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorProxyHeaderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorProxyHeaderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorProxyHeader(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorProxyHeaderRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorProxyHeaderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorProxyHeaderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorProxyHeader(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorProxyHeaderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorProxyHeaderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorProxyHeader(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorProxyHeader(d *schema.ResourceData) edpt.HealthMonitorProxyHeader {
	var ret edpt.HealthMonitorProxyHeader
	ret.Inst.ProxyHeaderVer = d.Get("proxy_header_ver").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
