package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_sip`: SIP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodSipCreate,
		UpdateContext: resourceHealthMonitorMethodSipUpdate,
		ReadContext:   resourceHealthMonitorMethodSipRead,
		DeleteContext: resourceHealthMonitorMethodSipDelete,

		Schema: map[string]*schema.Schema{
			"expect_response_code": {
				Type: schema.TypeString, Optional: true, Description: "Specify accepted response codes (e.g. 200, 400-430, any) (Format is xxx,xxx-xxx,any (xxx between [100,899]))",
			},
			"register": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send SIP REGISTER message, default is to send OPTION message",
			},
			"sip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP type",
			},
			"sip_hostname": {
				Type: schema.TypeString, Optional: true, Description: "Specify the SIP hostname that used in request",
			},
			"sip_port": {
				Type: schema.TypeInt, Optional: true, Default: 5060, Description: "Specify the SIP port, default is 5060 (Port Number)",
			},
			"sip_tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use TCP for transmission, default is UDP",
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
func resourceHealthMonitorMethodSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodSipRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodSipRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodSip(d *schema.ResourceData) edpt.HealthMonitorMethodSip {
	var ret edpt.HealthMonitorMethodSip
	ret.Inst.ExpectResponseCode = d.Get("expect_response_code").(string)
	ret.Inst.Register = d.Get("register").(int)
	ret.Inst.Sip = d.Get("sip").(int)
	ret.Inst.SipHostname = d.Get("sip_hostname").(string)
	ret.Inst.SipPort = d.Get("sip_port").(int)
	ret.Inst.SipTcp = d.Get("sip_tcp").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
