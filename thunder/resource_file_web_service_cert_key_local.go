package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileWebServiceCertKeyLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_web_service_cert_key_local`: Import web service's private key and certificate\n\n__PLACEHOLDER__",
		CreateContext: resourceFileWebServiceCertKeyLocalCreate,
		UpdateContext: resourceFileWebServiceCertKeyLocalUpdate,
		ReadContext:   resourceFileWebServiceCertKeyLocalRead,
		DeleteContext: resourceFileWebServiceCertKeyLocalDelete,

		Schema: map[string]*schema.Schema{
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
		},
	}
}
func resourceFileWebServiceCertKeyLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebServiceCertKeyLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebServiceCertKeyLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileWebServiceCertKeyLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileWebServiceCertKeyLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebServiceCertKeyLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebServiceCertKeyLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileWebServiceCertKeyLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileWebServiceCertKeyLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebServiceCertKeyLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebServiceCertKeyLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileWebServiceCertKeyLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebServiceCertKeyLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebServiceCertKeyLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileWebServiceCertKeyLocal(d *schema.ResourceData) edpt.FileWebServiceCertKeyLocal {
	var ret edpt.FileWebServiceCertKeyLocal
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	return ret
}
