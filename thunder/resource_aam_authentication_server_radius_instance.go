package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerRadiusInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server_radius_instance`: RADIUS Authentication Server instance\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerRadiusInstanceCreate,
		UpdateContext: resourceAamAuthenticationServerRadiusInstanceUpdate,
		ReadContext:   resourceAamAuthenticationServerRadiusInstanceRead,
		DeleteContext: resourceAamAuthenticationServerRadiusInstanceDelete,

		Schema: map[string]*schema.Schema{
			"accounting_port": {
				Type: schema.TypeInt, Optional: true, Default: 1813, Description: "Specify the RADIUS server's accounting port, default is 1813",
			},
			"acct_port_hm": {
				Type: schema.TypeString, Optional: true, Description: "Specify accounting port health check method",
			},
			"acct_port_hm_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured accounting port health check configuration",
			},
			"auth_type": {
				Type: schema.TypeString, Optional: true, Description: "'pap': PAP authentication. Default; 'mschapv2': MS-CHAPv2 authentication; 'mschapv2-pap': Use MS-CHAPv2 first. If server doesn't support it, try PAP;",
			},
			"health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check server's health status",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
			},
			"health_check_string": {
				Type: schema.TypeString, Optional: true, Description: "Health monitor name",
			},
			"host": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostip": {
							Type: schema.TypeString, Optional: true, Description: "Server's hostname(Length 1-31) or IP address",
						},
						"hostipv6": {
							Type: schema.TypeString, Optional: true, Description: "Server's IPV6 address",
						},
					},
				},
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the interval time for resend the request (second), default is 3 seconds (The interval time(second), default is 3 seconds)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify RADIUS authentication server name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 1812, Description: "Specify the RADIUS server's authentication port, default is 1812",
			},
			"port_hm": {
				Type: schema.TypeString, Optional: true, Description: "Check port's health status",
			},
			"port_hm_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured port health check configuration",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the retry number for resend the request, default is 5 (The retry number, default is 5)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'authen_success': Authentication Success; 'authen_failure': Authentication Failure; 'authorize_success': Authorization Success; 'authorize_failure': Authorization Failure; 'access_challenge': Access-Challenge Message Receive; 'timeout_error': Timeout; 'other_error': Other Error; 'request': Request; 'accounting-request-sent': Accounting-Request Sent; 'accounting-success': Accounting Success; 'accounting-failure': Accounting Failure;",
						},
					},
				},
			},
			"secret": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the RADIUS server's secret",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "The RADIUS server's secret",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServerRadiusInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadiusInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerRadiusInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerRadiusInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadiusInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerRadiusInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerRadiusInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadiusInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerRadiusInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadiusInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationServerRadiusInstanceHost(d []interface{}) edpt.AamAuthenticationServerRadiusInstanceHost {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerRadiusInstanceHost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerRadiusInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerRadiusInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerRadiusInstance(d *schema.ResourceData) edpt.AamAuthenticationServerRadiusInstance {
	var ret edpt.AamAuthenticationServerRadiusInstance
	ret.Inst.AccountingPort = d.Get("accounting_port").(int)
	ret.Inst.AcctPortHm = d.Get("acct_port_hm").(string)
	ret.Inst.AcctPortHmDisable = d.Get("acct_port_hm_disable").(int)
	ret.Inst.AuthType = d.Get("auth_type").(string)
	//omit encrypted
	ret.Inst.HealthCheck = d.Get("health_check").(int)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckString = d.Get("health_check_string").(string)
	ret.Inst.Host = getObjectAamAuthenticationServerRadiusInstanceHost(d.Get("host").([]interface{}))
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortHm = d.Get("port_hm").(string)
	ret.Inst.PortHmDisable = d.Get("port_hm_disable").(int)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServerRadiusInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Secret = d.Get("secret").(int)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	//omit uuid
	return ret
}
