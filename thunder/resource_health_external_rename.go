package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthExternalRename() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_external_rename`: Rename external health monitor script\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthExternalRenameCreate,
		UpdateContext: resourceHealthExternalRenameUpdate,
		ReadContext:   resourceHealthExternalRenameRead,
		DeleteContext: resourceHealthExternalRenameDelete,

		Schema: map[string]*schema.Schema{
			"dst_file": {
				Type: schema.TypeString, Optional: true, Description: "Destination external health monitor script file name",
			},
			"src_file": {
				Type: schema.TypeString, Optional: true, Description: "Source external health monitor script file name",
			},
		},
	}
}
func resourceHealthExternalRenameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalRenameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalRename(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalRenameRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthExternalRenameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalRenameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalRename(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalRenameRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthExternalRenameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalRenameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalRename(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthExternalRenameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalRenameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalRename(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthExternalRename(d *schema.ResourceData) edpt.HealthExternalRename {
	var ret edpt.HealthExternalRename
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.SrcFile = d.Get("src_file").(string)
	return ret
}
