package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileWebCategoryLicenseLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_web_category_license_local`: License file to enable web-category feature activation\n\n__PLACEHOLDER__",
		CreateContext: resourceFileWebCategoryLicenseLocalCreate,
		UpdateContext: resourceFileWebCategoryLicenseLocalUpdate,
		ReadContext:   resourceFileWebCategoryLicenseLocalRead,
		DeleteContext: resourceFileWebCategoryLicenseLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'import': import;",
			},
			"device": {
				Type: schema.TypeInt, Optional: true, Description: "Device (Device ID)",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "Web-Category license local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "Full path of the uploaded file",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable management port for backend",
			},
		},
	}
}
func resourceFileWebCategoryLicenseLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebCategoryLicenseLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebCategoryLicenseLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileWebCategoryLicenseLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileWebCategoryLicenseLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebCategoryLicenseLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebCategoryLicenseLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileWebCategoryLicenseLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileWebCategoryLicenseLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebCategoryLicenseLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebCategoryLicenseLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileWebCategoryLicenseLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileWebCategoryLicenseLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileWebCategoryLicenseLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileWebCategoryLicenseLocal(d *schema.ResourceData) edpt.FileWebCategoryLicenseLocal {
	var ret edpt.FileWebCategoryLicenseLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
