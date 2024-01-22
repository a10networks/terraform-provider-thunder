package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportNgWafCustomPage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_ng_waf_custom_page`: Custom html file for ng-waf\n\n__PLACEHOLDER__",
		CreateContext: resourceImportNgWafCustomPageCreate,
		UpdateContext: resourceImportNgWafCustomPageUpdate,
		ReadContext:   resourceImportNgWafCustomPageRead,
		DeleteContext: resourceImportNgWafCustomPageDelete,

		Schema: map[string]*schema.Schema{
			"custom_page": {
				Type: schema.TypeString, Optional: true, Description: "Custom html file for ng-waf",
			},
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
func resourceImportNgWafCustomPageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafCustomPageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafCustomPage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportNgWafCustomPageRead(ctx, d, meta)
	}
	return diags
}

func resourceImportNgWafCustomPageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafCustomPageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafCustomPage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportNgWafCustomPageRead(ctx, d, meta)
	}
	return diags
}
func resourceImportNgWafCustomPageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafCustomPageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafCustomPage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportNgWafCustomPageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportNgWafCustomPageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportNgWafCustomPage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportNgWafCustomPage(d *schema.ResourceData) edpt.ImportNgWafCustomPage {
	var ret edpt.ImportNgWafCustomPage
	ret.Inst.CustomPage = d.Get("custom_page").(string)
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
