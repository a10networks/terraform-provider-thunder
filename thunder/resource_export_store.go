package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExportStore() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_export_store`: Create store name for remote url\n\n__PLACEHOLDER__",
		CreateContext: resourceExportStoreCreate,
		UpdateContext: resourceExportStoreUpdate,
		ReadContext:   resourceExportStoreRead,
		DeleteContext: resourceExportStoreDelete,

		Schema: map[string]*schema.Schema{
			"create": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create an export store profile",
			},
			"delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an export store profile",
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "profile name to store remote url",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
		},
	}
}
func resourceExportStoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportStoreCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportStore(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceExportStoreRead(ctx, d, meta)
	}
	return diags
}

func resourceExportStoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportStoreUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportStore(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceExportStoreRead(ctx, d, meta)
	}
	return diags
}
func resourceExportStoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportStoreDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportStore(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceExportStoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportStoreRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExportStore(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointExportStore(d *schema.ResourceData) edpt.ExportStore {
	var ret edpt.ExportStore
	ret.Inst.Create = d.Get("create").(int)
	ret.Inst.Delete = d.Get("delete").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	return ret
}
