package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationAccount() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_account`: Authentication account configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationAccountCreate,
		UpdateContext: resourceAamAuthenticationAccountUpdate,
		ReadContext:   resourceAamAuthenticationAccountRead,
		DeleteContext: resourceAamAuthenticationAccountDelete,

		Schema: map[string]*schema.Schema{
			"kerberos_spn_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Specify AD account name",
						},
						"realm": {
							Type: schema.TypeString, Optional: true, Description: "Specify Kerberos realm",
						},
						"account": {
							Type: schema.TypeString, Optional: true, Description: "Specify domain account for SPN",
						},
						"service_principal_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify service principal name",
						},
						"password": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify password of domain account",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "Password of AD account",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'request-normal': Total Normal Request; 'request-dropped': Total Dropped Request; 'response-success': Total Success Response; 'response-failure': Total Failure Response; 'response-error': Total Error Response; 'response-timeout': Total Timeout Response; 'response-other': Total Other Response;",
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
func resourceAamAuthenticationAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccount(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationAccountRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccount(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationAccountRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccount(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccount(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationAccountKerberosSpnList(d []interface{}) []edpt.AamAuthenticationAccountKerberosSpnList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationAccountKerberosSpnList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationAccountKerberosSpnList
		oi.Name = in["name"].(string)
		oi.Realm = in["realm"].(string)
		oi.Account = in["account"].(string)
		oi.ServicePrincipalName = in["service_principal_name"].(string)
		oi.Password = in["password"].(int)
		oi.SecretString = in["secret_string"].(string)
		//omit encrypted
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationAccountSamplingEnable(d []interface{}) []edpt.AamAuthenticationAccountSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationAccountSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationAccountSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationAccount(d *schema.ResourceData) edpt.AamAuthenticationAccount {
	var ret edpt.AamAuthenticationAccount
	ret.Inst.KerberosSpnList = getSliceAamAuthenticationAccountKerberosSpnList(d.Get("kerberos_spn_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceAamAuthenticationAccountSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
