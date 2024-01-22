package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayKerberosInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_kerberos_instance`: Kerberos Authentication Relay\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayKerberosInstanceCreate,
		UpdateContext: resourceAamAuthenticationRelayKerberosInstanceUpdate,
		ReadContext:   resourceAamAuthenticationRelayKerberosInstanceRead,
		DeleteContext: resourceAamAuthenticationRelayKerberosInstanceDelete,

		Schema: map[string]*schema.Schema{
			"kerberos_account": {
				Type: schema.TypeString, Optional: true, Description: "Specify the kerberos account name",
			},
			"kerberos_kdc": {
				Type: schema.TypeString, Optional: true, Description: "Specify the kerberos kdc ip or host name",
			},
			"kerberos_kdc_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Specify an authentication service group as multiple KDCs",
			},
			"kerberos_realm": {
				Type: schema.TypeString, Optional: true, Description: "Specify the kerberos realm",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Kerberos authentication relay name",
			},
			"password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify password of Kerberos password",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 88, Description: "Specify The KDC port, default is 88",
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
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "The kerberos client password",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify timeout for kerberos transport, default is 10 seconds (The timeout, default is 10 seconds)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationRelayKerberosInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberosInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayKerberosInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayKerberosInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberosInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayKerberosInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayKerberosInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberosInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayKerberosInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayKerberosInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayKerberosInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelayKerberosInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayKerberosInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayKerberosInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayKerberosInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayKerberosInstance(d *schema.ResourceData) edpt.AamAuthenticationRelayKerberosInstance {
	var ret edpt.AamAuthenticationRelayKerberosInstance
	//omit encrypted
	ret.Inst.KerberosAccount = d.Get("kerberos_account").(string)
	ret.Inst.KerberosKdc = d.Get("kerberos_kdc").(string)
	ret.Inst.KerberosKdcServiceGroup = d.Get("kerberos_kdc_service_group").(string)
	ret.Inst.KerberosRealm = d.Get("kerberos_realm").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Password = d.Get("password").(int)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelayKerberosInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
