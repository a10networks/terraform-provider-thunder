package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_udp`: UDP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodUdpCreate,
		UpdateContext: resourceHealthMonitorMethodUdpUpdate,
		ReadContext:   resourceHealthMonitorMethodUdpRead,
		DeleteContext: resourceHealthMonitorMethodUdpDelete,

		Schema: map[string]*schema.Schema{
			"force_up_with_single_healthcheck": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force Up with no response at the first time",
			},
			"udp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "UDP type",
			},
			"udp_port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify UDP port (Specify port number)",
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
func resourceHealthMonitorMethodUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodUdp(d *schema.ResourceData) edpt.HealthMonitorMethodUdp {
	var ret edpt.HealthMonitorMethodUdp
	ret.Inst.ForceUpWithSingleHealthcheck = d.Get("force_up_with_single_healthcheck").(int)
	ret.Inst.Udp = d.Get("udp").(int)
	ret.Inst.UdpPort = d.Get("udp_port").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
