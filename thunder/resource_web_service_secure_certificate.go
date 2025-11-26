package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebServiceSecureCertificate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_service_secure_certificate`: Web-service secure certificate operation\n\n__PLACEHOLDER__",
		CreateContext: resourceWebServiceSecureCertificateCreate,
		UpdateContext: resourceWebServiceSecureCertificateUpdate,
		ReadContext:   resourceWebServiceSecureCertificateRead,
		DeleteContext: resourceWebServiceSecureCertificateDelete,

		Schema: map[string]*schema.Schema{
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"load": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load WEB certificate",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceWebServiceSecureCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureCertificateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureCertificate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureCertificateRead(ctx, d, meta)
	}
	return diags
}

func resourceWebServiceSecureCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureCertificateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureCertificate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureCertificateRead(ctx, d, meta)
	}
	return diags
}
func resourceWebServiceSecureCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureCertificateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureCertificate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebServiceSecureCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureCertificateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecureCertificate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointWebServiceSecureCertificate(d *schema.ResourceData) edpt.WebServiceSecureCertificate {
	var ret edpt.WebServiceSecureCertificate
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Load = d.Get("load").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
