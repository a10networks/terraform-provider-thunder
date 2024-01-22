package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCliMonitorInterval() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_cli_monitor_interval`: Config cli-server health monitor interval\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemCliMonitorIntervalCreate,
		UpdateContext: resourceSystemCliMonitorIntervalUpdate,
		ReadContext:   resourceSystemCliMonitorIntervalRead,
		DeleteContext: resourceSystemCliMonitorIntervalDelete,

		Schema: map[string]*schema.Schema{
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "one interval is 300ms (0 = disable)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemCliMonitorIntervalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCliMonitorIntervalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCliMonitorInterval(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCliMonitorIntervalRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemCliMonitorIntervalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCliMonitorIntervalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCliMonitorInterval(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemCliMonitorIntervalRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemCliMonitorIntervalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCliMonitorIntervalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCliMonitorInterval(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemCliMonitorIntervalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCliMonitorIntervalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCliMonitorInterval(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemCliMonitorInterval(d *schema.ResourceData) edpt.SystemCliMonitorInterval {
	var ret edpt.SystemCliMonitorInterval
	ret.Inst.Interval = d.Get("interval").(int)
	//omit uuid
	return ret
}
