package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayKerberos() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_kerberos`: Kerberos Authentication Relay\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayKerberosCreate,
		UpdateContext: resourceAamAuthenticationRelayKerberosUpdate,
		ReadContext:   resourceAamAuthenticationRelayKerberosRead,
		DeleteContext: resourceAamAuthenticationRelayKerberosDelete,

		Schema: map[string]*schema.Schema{
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify Kerberos authentication relay name",
						},
						"kerberos_realm": {
							Type: schema.TypeString, Optional: true, Description: "Specify the kerberos realm",
						},
						"kerberos_kdc": {
							Type: schema.TypeString, Optional: true, Description: "Specify the kerberos kdc ip or host name",
						},
						"kerberos_kdc_service_group": {
							Type: schema.TypeString, Optional: true, Description: "Specify an authentication service group as multiple KDCs",
						},
						"kerberos_account": {
							Type: schema.TypeString, Optional: true, Description: "Specify the kerberos account name",
						},
						"password": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify password of Kerberos password",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "The kerberos client password",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Default: 88, Description: "Specify The KDC port, default is 88",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify timeout for kerberos transport, default is 10 seconds (The timeout, default is 10 seconds)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'request-send': Request Send; 'response-receive': Response Receive; 'current-requests-of-user': Current Pending Requests of User; 'tickets': Tickets;",
									},
								},
							},
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'request-send': Total Request Send; 'response-get': Total Response Get; 'timeout-error': Total Timeout; 'other-error': Total Other Error; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response; 'job-start-error': Total Job Start Error; 'polling-control-error': Total Polling Control Error;",
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
func resourceAamAuthenticationRelayKerberosCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberos(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayKerberosRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayKerberosUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberos(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayKerberosRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayKerberosDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberos(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayKerberosRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberos(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelayKerberosInstanceList(d []interface{}) []edpt.AamAuthenticationRelayKerberosInstanceList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayKerberosInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayKerberosInstanceList
		oi.Name = in["name"].(string)
		oi.KerberosRealm = in["kerberos_realm"].(string)
		oi.KerberosKdc = in["kerberos_kdc"].(string)
		oi.KerberosKdcServiceGroup = in["kerberos_kdc_service_group"].(string)
		oi.KerberosAccount = in["kerberos_account"].(string)
		oi.Password = in["password"].(int)
		oi.SecretString = in["secret_string"].(string)
		//omit encrypted
		oi.Port = in["port"].(int)
		oi.Timeout = in["timeout"].(int)
		//omit uuid
		oi.SamplingEnable = getSliceAamAuthenticationRelayKerberosInstanceListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationRelayKerberosInstanceListSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayKerberosInstanceListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayKerberosInstanceListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayKerberosInstanceListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationRelayKerberosSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayKerberosSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayKerberosSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayKerberosSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayKerberos(d *schema.ResourceData) edpt.AamAuthenticationRelayKerberos {
	var ret edpt.AamAuthenticationRelayKerberos
	ret.Inst.InstanceList = getSliceAamAuthenticationRelayKerberosInstanceList(d.Get("instance_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelayKerberosSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
