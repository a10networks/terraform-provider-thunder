package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportToDevice() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_to_device`: specify the target device to import in aVCS\n\n__PLACEHOLDER__",
		CreateContext: resourceImportToDeviceCreate,
		UpdateContext: resourceImportToDeviceUpdate,
		ReadContext:   resourceImportToDeviceRead,
		DeleteContext: resourceImportToDeviceDelete,

		Schema: map[string]*schema.Schema{
			"device": {
				Type: schema.TypeInt, Optional: true, Description: "Device (Device ID)",
			},
			"glm_cert": {
				Type: schema.TypeString, Optional: true, Description: "GLM certificate",
			},
			"glm_license": {
				Type: schema.TypeString, Optional: true, Description: "License File",
			},
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"web_category_license": {
				Type: schema.TypeString, Optional: true, Description: "License file to enable web-category feature",
			},
		},
	}
}
func resourceImportToDeviceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportToDeviceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportToDevice(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportToDeviceRead(ctx, d, meta)
	}
	return diags
}

func resourceImportToDeviceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportToDeviceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportToDevice(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportToDeviceRead(ctx, d, meta)
	}
	return diags
}
func resourceImportToDeviceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportToDeviceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportToDevice(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportToDeviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportToDeviceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportToDevice(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportToDevice(d *schema.ResourceData) edpt.ImportToDevice {
	var ret edpt.ImportToDevice
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.GlmCert = d.Get("glm_cert").(string)
	ret.Inst.GlmLicense = d.Get("glm_license").(string)
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.WebCategoryLicense = d.Get("web_category_license").(string)
	return ret
}
