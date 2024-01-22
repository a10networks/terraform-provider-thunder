package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSslExpireCheckException() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ssl_expire_check_exception`: Config the exception which will not be checked\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSslExpireCheckExceptionCreate,
		UpdateContext: resourceSlbSslExpireCheckExceptionUpdate,
		ReadContext:   resourceSlbSslExpireCheckExceptionRead,
		DeleteContext: resourceSlbSslExpireCheckExceptionDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'add': Add an exception; 'delete': Delete an exception; 'clean': Delete all exception;",
			},
			"certificate_name": {
				Type: schema.TypeString, Optional: true, Description: "The certificate name",
			},
		},
	}
}
func resourceSlbSslExpireCheckExceptionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckExceptionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheckException(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslExpireCheckExceptionRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSslExpireCheckExceptionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckExceptionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheckException(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSslExpireCheckExceptionRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSslExpireCheckExceptionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckExceptionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheckException(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSslExpireCheckExceptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSslExpireCheckExceptionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSslExpireCheckException(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbSslExpireCheckException(d *schema.ResourceData) edpt.SlbSslExpireCheckException {
	var ret edpt.SlbSslExpireCheckException
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.CertificateName = d.Get("certificate_name").(string)
	return ret
}
