package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServerRadius() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_server_radius`: RADIUS Authentication Server\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServerRadiusCreate,
		UpdateContext: resourceAamAuthenticationServerRadiusUpdate,
		ReadContext:   resourceAamAuthenticationServerRadiusRead,
		DeleteContext: resourceAamAuthenticationServerRadiusDelete,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify RADIUS authentication server name",
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
						"secret": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the RADIUS server's secret",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "The RADIUS server's secret",
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
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Specify the interval time for resend the request (second), default is 3 seconds (The interval time(second), default is 3 seconds)",
						},
						"retry": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the retry number for resend the request, default is 5 (The retry number, default is 5)",
						},
						"health_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Check server's health status",
						},
						"health_check_string": {
							Type: schema.TypeString, Optional: true, Description: "Health monitor name",
						},
						"health_check_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
						},
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
						"packet_capture_template": {
							Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'authen_success': Total Authentication Success; 'authen_failure': Total Authentication Failure; 'authorize_success': Total Authorization Success; 'authorize_failure': Total Authorization Failure; 'access_challenge': Total Access-Challenge Message Receive; 'timeout_error': Total Timeout; 'other_error': Total Other Error; 'request': Total Request; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response; 'job-start-error': Total Job Start Error; 'polling-control-error': Total Polling Control Error; 'accounting-request-sent': Accounting-Request Sent; 'accounting-success': Accounting Success; 'accounting-failure': Accounting Failure;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServerRadiusCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadius(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerRadiusRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServerRadiusUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadius(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServerRadiusRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServerRadiusDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadius(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServerRadiusRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServerRadiusRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServerRadius(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationServerRadiusInstanceList(d []interface{}) []edpt.AamAuthenticationServerRadiusInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusInstanceList
		oi.Name = in["name"].(string)
		oi.Host = getObjectAamAuthenticationServerRadiusInstanceListHost(in["host"].([]interface{}))
		oi.Secret = in["secret"].(int)
		oi.SecretString = in["secret_string"].(string)
		//omit encrypted
		oi.Port = in["port"].(int)
		oi.PortHm = in["port_hm"].(string)
		oi.PortHmDisable = in["port_hm_disable"].(int)
		oi.Interval = in["interval"].(int)
		oi.Retry = in["retry"].(int)
		oi.HealthCheck = in["health_check"].(int)
		oi.HealthCheckString = in["health_check_string"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		oi.AccountingPort = in["accounting_port"].(int)
		oi.AcctPortHm = in["acct_port_hm"].(string)
		oi.AcctPortHmDisable = in["acct_port_hm_disable"].(int)
		oi.AuthType = in["auth_type"].(string)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationServerRadiusInstanceListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationServerRadiusInstanceListHost(d []interface{}) edpt.AamAuthenticationServerRadiusInstanceListHost {

	count1 := len(d)
	var ret edpt.AamAuthenticationServerRadiusInstanceListHost
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hostip = in["hostip"].(string)
		ret.Hostipv6 = in["hostipv6"].(string)
	}
	return ret
}

func getSliceAamAuthenticationServerRadiusInstanceListSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerRadiusInstanceListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusInstanceListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusInstanceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServerRadiusSamplingEnable(d []interface{}) []edpt.AamAuthenticationServerRadiusSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServerRadiusSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServerRadiusSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServerRadius(d *schema.ResourceData) edpt.AamAuthenticationServerRadius {
	var ret edpt.AamAuthenticationServerRadius
	ret.Inst.InstanceList = getSliceAamAuthenticationServerRadiusInstanceList(d.Get("instance_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServerRadiusSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
