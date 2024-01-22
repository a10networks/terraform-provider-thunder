package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportNgWafModule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_ng_waf_module`: WAF module\n\n__PLACEHOLDER__",
		CreateContext: resourceImportNgWafModuleCreate,
		UpdateContext: resourceImportNgWafModuleUpdate,
		ReadContext:   resourceImportNgWafModuleRead,
		DeleteContext: resourceImportNgWafModuleDelete,

		Schema: map[string]*schema.Schema{
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "Profile name for remote url",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management interface for reachability",
			},
		},
	}
}
func resourceImportNgWafModuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafModuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafModule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportNgWafModuleRead(ctx, d, meta)
	}
	return diags
}

func resourceImportNgWafModuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafModuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafModule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportNgWafModuleRead(ctx, d, meta)
	}
	return diags
}
func resourceImportNgWafModuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafModuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafModule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportNgWafModuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafModuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafModule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportNgWafModule(d *schema.ResourceData) edpt.ImportNgWafModule {
	var ret edpt.ImportNgWafModule
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
