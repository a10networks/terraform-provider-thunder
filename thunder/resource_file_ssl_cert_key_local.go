package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSslCertKeyLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_ssl_cert_key_local`: ssl certificate and key file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSslCertKeyLocalCreate,
		UpdateContext: resourceFileSslCertKeyLocalUpdate,
		ReadContext:   resourceFileSslCertKeyLocalRead,
		DeleteContext: resourceFileSslCertKeyLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'replace': replace;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "ssl certificate local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark keys as non-exportable",
			},
		},
	}
}
func resourceFileSslCertKeyLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertKeyLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCertKeyLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslCertKeyLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSslCertKeyLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertKeyLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCertKeyLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslCertKeyLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSslCertKeyLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertKeyLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCertKeyLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSslCertKeyLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslCertKeyLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslCertKeyLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileSslCertKeyLocal(d *schema.ResourceData) edpt.FileSslCertKeyLocal {
	var ret edpt.FileSslCertKeyLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.Secured = d.Get("secured").(int)
	return ret
}
