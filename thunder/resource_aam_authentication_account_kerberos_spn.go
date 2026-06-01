package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationAccountKerberosSpn() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_account_kerberos_spn`: AD domain account associated with a SPN\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationAccountKerberosSpnCreate,
		UpdateContext: resourceAamAuthenticationAccountKerberosSpnUpdate,
		ReadContext:   resourceAamAuthenticationAccountKerberosSpnRead,
		DeleteContext: resourceAamAuthenticationAccountKerberosSpnDelete,

		Schema: map[string]*schema.Schema{
			"account": {
				Type: schema.TypeString, Optional: true, Description: "Specify domain account for SPN",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify AD account name",
			},
			"password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify password of domain account",
			},
			"realm": {
				Type: schema.TypeString, Optional: true, Description: "Specify Kerberos realm",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "Password of AD account",
			},
			"service_principal_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify service principal name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationAccountKerberosSpnCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountKerberosSpnCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccountKerberosSpn(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationAccountKerberosSpnRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationAccountKerberosSpnUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountKerberosSpnUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccountKerberosSpn(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationAccountKerberosSpnRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationAccountKerberosSpnDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountKerberosSpnDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccountKerberosSpn(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationAccountKerberosSpnRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationAccountKerberosSpnRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationAccountKerberosSpn(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthenticationAccountKerberosSpn(d *schema.ResourceData) edpt.AamAuthenticationAccountKerberosSpn {
	var ret edpt.AamAuthenticationAccountKerberosSpn
	ret.Inst.Account = d.Get("account").(string)
	//omit encrypted
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Password = d.Get("password").(int)
	ret.Inst.Realm = d.Get("realm").(string)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.ServicePrincipalName = d.Get("service_principal_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
