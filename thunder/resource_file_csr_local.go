package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileCsrLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_csr_local`: ssl certificate signing request file\n\n__PLACEHOLDER__",
		CreateContext: resourceFileCsrLocalCreate,
		UpdateContext: resourceFileCsrLocalUpdate,
		ReadContext:   resourceFileCsrLocalRead,
		DeleteContext: resourceFileCsrLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'export': export;",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "Specify the File Name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileCsrLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCsrLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileCsrLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileCsrLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileCsrLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileCsrLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileCsrLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileCsrLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileCsrLocal(d *schema.ResourceData) edpt.FileCsrLocal {
	var ret edpt.FileCsrLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	//omit uuid
	return ret
}
