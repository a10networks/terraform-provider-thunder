package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslExpireCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ssl_expire_check`: SSL certificate expiration check\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSslExpireCheckCreate,
		UpdateContext: resourceSlbSslExpireCheckUpdate,
		ReadContext:   resourceSlbSslExpireCheckRead,
		DeleteContext: resourceSlbSslExpireCheckDelete,

		Schema: map[string]*schema.Schema{
			"before": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "The number of days in advance notice before expiration, default is 5",
			},
			"exception": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'add': Add an exception; 'delete': Delete an exception; 'clean': Delete all exception;",
						},
						"certificate_name": {
							Type: schema.TypeString, Optional: true, Description: "The certificate name",
						},
					},
				},
			},
			"expire_address1": {
				Type: schema.TypeString, Optional: true, Description: "Email address for certificate expiration check",
			},
			"interval_days": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "The interval of days notice after expiration, default is 2",
			},
			"ssl_expire_email_address": {
				Type: schema.TypeString, Optional: true, Description: "Config Email address for certificate expiration check",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbSslExpireCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslExpireCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSslExpireCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslExpireCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSslExpireCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSslExpireCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbSslExpireCheckException1414(d []interface{}) edpt.SlbSslExpireCheckException1414 {

	count1 := len(d)
	var ret edpt.SlbSslExpireCheckException1414
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.CertificateName = in["certificate_name"].(string)
	}
	return ret
}

func dataToEndpointSlbSslExpireCheck(d *schema.ResourceData) edpt.SlbSslExpireCheck {
	var ret edpt.SlbSslExpireCheck
	ret.Inst.Before = d.Get("before").(int)
	ret.Inst.Exception = getObjectSlbSslExpireCheckException1414(d.Get("exception").([]interface{}))
	ret.Inst.ExpireAddress1 = d.Get("expire_address1").(string)
	ret.Inst.IntervalDays = d.Get("interval_days").(int)
	ret.Inst.SslExpireEmailAddress = d.Get("ssl_expire_email_address").(string)
	//omit uuid
	return ret
}
