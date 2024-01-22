package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodNtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_ntp`: NTP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodNtpCreate,
		UpdateContext: resourceHealthMonitorMethodNtpUpdate,
		ReadContext:   resourceHealthMonitorMethodNtpRead,
		DeleteContext: resourceHealthMonitorMethodNtpDelete,

		Schema: map[string]*schema.Schema{
			"ntp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "NTP type",
			},
			"ntp_port": {
				Type: schema.TypeInt, Optional: true, Default: 123, Description: "Specify the NTP port, default is 123 (Port Number (default 123))",
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
func resourceHealthMonitorMethodNtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodNtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodNtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodNtpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodNtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodNtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodNtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodNtpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodNtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodNtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodNtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodNtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodNtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodNtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodNtp(d *schema.ResourceData) edpt.HealthMonitorMethodNtp {
	var ret edpt.HealthMonitorMethodNtp
	ret.Inst.Ntp = d.Get("ntp").(int)
	ret.Inst.NtpPort = d.Get("ntp_port").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
