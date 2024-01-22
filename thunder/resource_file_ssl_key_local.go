package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileSslKeyLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_ssl_key_local`: ssl key file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileSslKeyLocalCreate,
		UpdateContext: resourceFileSslKeyLocalUpdate,
		ReadContext:   resourceFileSslKeyLocalRead,
		DeleteContext: resourceFileSslKeyLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'create': create; 'import': import; 'export': export; 'replace': replace;",
			},
			"dst_file": {
				Type: schema.TypeString, Optional: true, Description: "destination file name for copy and rename action",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "ssl key local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark as non-exportable",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFileSslKeyLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslKeyLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslKeyLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileSslKeyLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslKeyLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileSslKeyLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileSslKeyLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslKeyLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileSslKeyLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileSslKeyLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileSslKeyLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileSslKeyLocal(d *schema.ResourceData) edpt.FileSslKeyLocal {
	var ret edpt.FileSslKeyLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.Secured = d.Get("secured").(int)
	//omit uuid
	return ret
}
