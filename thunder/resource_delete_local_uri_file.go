package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteLocalUriFile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_local_uri_file`: Local URI files for http response\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteLocalUriFileCreate,
		UpdateContext: resourceDeleteLocalUriFileUpdate,
		ReadContext:   resourceDeleteLocalUriFileRead,
		DeleteContext: resourceDeleteLocalUriFileDelete,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "Local File Name",
			},
		},
	}
}
func resourceDeleteLocalUriFileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteLocalUriFileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteLocalUriFile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteLocalUriFileRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteLocalUriFileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteLocalUriFileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteLocalUriFile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteLocalUriFileRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteLocalUriFileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteLocalUriFileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteLocalUriFile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteLocalUriFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteLocalUriFileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteLocalUriFile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteLocalUriFile(d *schema.ResourceData) edpt.DeleteLocalUriFile {
	var ret edpt.DeleteLocalUriFile
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
