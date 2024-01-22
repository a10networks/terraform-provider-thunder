package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBackupStore() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_backup_store`: Save backup store information\n\n__PLACEHOLDER__",
		CreateContext: resourceBackupStoreCreate,
		UpdateContext: resourceBackupStoreUpdate,
		ReadContext:   resourceBackupStoreRead,
		DeleteContext: resourceBackupStoreDelete,

		Schema: map[string]*schema.Schema{
			"creat_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create store",
						},
						"store_name": {
							Type: schema.TypeString, Optional: true, Description: "profile name to store remote url",
						},
						"remote_file": {
							Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
						},
					},
				},
			},
			"delete_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete store",
						},
						"store_name_del": {
							Type: schema.TypeString, Optional: true, Description: "profile name for deleting",
						},
					},
				},
			},
		},
	}
}
func resourceBackupStoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupStoreCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupStore(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupStoreRead(ctx, d, meta)
	}
	return diags
}

func resourceBackupStoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupStoreUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupStore(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceBackupStoreRead(ctx, d, meta)
	}
	return diags
}
func resourceBackupStoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupStoreDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupStore(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBackupStoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceBackupStoreRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointBackupStore(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectBackupStoreCreatCfg(d []interface{}) edpt.BackupStoreCreatCfg {

	count1 := len(d)
	var ret edpt.BackupStoreCreatCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Create = in["create"].(int)
		ret.StoreName = in["store_name"].(string)
		ret.RemoteFile = in["remote_file"].(string)
	}
	return ret
}

func getObjectBackupStoreDeleteCfg(d []interface{}) edpt.BackupStoreDeleteCfg {

	count1 := len(d)
	var ret edpt.BackupStoreDeleteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Delete = in["delete"].(int)
		ret.StoreNameDel = in["store_name_del"].(string)
	}
	return ret
}

func dataToEndpointBackupStore(d *schema.ResourceData) edpt.BackupStore {
	var ret edpt.BackupStore
	ret.Inst.CreatCfg = getObjectBackupStoreCreatCfg(d.Get("creat_cfg").([]interface{}))
	ret.Inst.DeleteCfg = getObjectBackupStoreDeleteCfg(d.Get("delete_cfg").([]interface{}))
	return ret
}
