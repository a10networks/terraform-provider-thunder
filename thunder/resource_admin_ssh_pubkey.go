package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdminSshPubkey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_admin_ssh_pubkey`: Config openssh authorized public keys management\n\n__PLACEHOLDER__",
		CreateContext: resourceAdminSshPubkeyCreate,
		UpdateContext: resourceAdminSshPubkeyUpdate,
		ReadContext:   resourceAdminSshPubkeyRead,
		DeleteContext: resourceAdminSshPubkeyDelete,

		Schema: map[string]*schema.Schema{
			"delete": {
				Type: schema.TypeInt, Optional: true, Description: "Delete an authorized public key (SSH key index)",
			},
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"import": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Import an authorized public key",
			},
			"list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "List all authorized public keys",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"user": {
				Type: schema.TypeString, Required: true, Description: "User",
			},
		},
	}
}
func resourceAdminSshPubkeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSshPubkeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSshPubkey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminSshPubkeyRead(ctx, d, meta)
	}
	return diags
}

func resourceAdminSshPubkeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSshPubkeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSshPubkey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAdminSshPubkeyRead(ctx, d, meta)
	}
	return diags
}
func resourceAdminSshPubkeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSshPubkeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSshPubkey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAdminSshPubkeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAdminSshPubkeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAdminSshPubkey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAdminSshPubkey(d *schema.ResourceData) edpt.AdminSshPubkey {
	var ret edpt.AdminSshPubkey
	ret.Inst.Delete = d.Get("delete").(int)
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Import = d.Get("import").(int)
	ret.Inst.List = d.Get("list").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.User = d.Get("user").(string)
	return ret
}
