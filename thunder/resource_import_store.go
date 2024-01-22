package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportStore() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_store`: Create store name for remote url\n\n__PLACEHOLDER__",
		CreateContext: resourceImportStoreCreate,
		UpdateContext: resourceImportStoreUpdate,
		ReadContext:   resourceImportStoreRead,
		DeleteContext: resourceImportStoreDelete,

		Schema: map[string]*schema.Schema{
			"create": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create an import store profile",
			},
			"delete": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an import store profile",
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
func resourceImportStoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportStoreCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportStore(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportStoreRead(ctx, d, meta)
	}
	return diags
}

func resourceImportStoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportStoreUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportStore(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportStoreRead(ctx, d, meta)
	}
	return diags
}
func resourceImportStoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportStoreDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportStore(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportStoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportStoreRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportStore(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportStore(d *schema.ResourceData) edpt.ImportStore {
	var ret edpt.ImportStore
	ret.Inst.Create = d.Get("create").(int)
	ret.Inst.Delete = d.Get("delete").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	return ret
}
