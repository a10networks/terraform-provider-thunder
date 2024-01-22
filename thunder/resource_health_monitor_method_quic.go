package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodQuic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_quic`: QUIC type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodQuicCreate,
		UpdateContext: resourceHealthMonitorMethodQuicUpdate,
		ReadContext:   resourceHealthMonitorMethodQuicRead,
		DeleteContext: resourceHealthMonitorMethodQuicDelete,

		Schema: map[string]*schema.Schema{
			"quic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "QUIC type",
			},
			"quic_port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "Specify QUIC port (Port Number (default 443))",
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
func resourceHealthMonitorMethodQuicCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodQuicCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodQuic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodQuicRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodQuicUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodQuicUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodQuic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodQuicRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodQuicDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodQuicDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodQuic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodQuicRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodQuicRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodQuic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodQuic(d *schema.ResourceData) edpt.HealthMonitorMethodQuic {
	var ret edpt.HealthMonitorMethodQuic
	ret.Inst.Quic = d.Get("quic").(int)
	ret.Inst.QuicPort = d.Get("quic_port").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
