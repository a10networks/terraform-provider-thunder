package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileGlmLicenseLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_glm_license_local`: glm license file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceFileGlmLicenseLocalCreate,
		UpdateContext: resourceFileGlmLicenseLocalUpdate,
		ReadContext:   resourceFileGlmLicenseLocalRead,
		DeleteContext: resourceFileGlmLicenseLocalDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'import': import;",
			},
			"device": {
				Type: schema.TypeInt, Optional: true, Description: "Device (Device ID)",
			},
			"file": {
				Type: schema.TypeString, Optional: true, Description: "glm license local file name",
			},
			"file_handle": {
				Type: schema.TypeString, Optional: true, Description: "full path of the uploaded file",
			},
		},
	}
}
func resourceFileGlmLicenseLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGlmLicenseLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGlmLicenseLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGlmLicenseLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceFileGlmLicenseLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGlmLicenseLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGlmLicenseLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileGlmLicenseLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceFileGlmLicenseLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGlmLicenseLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGlmLicenseLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileGlmLicenseLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileGlmLicenseLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileGlmLicenseLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileGlmLicenseLocal(d *schema.ResourceData) edpt.FileGlmLicenseLocal {
	var ret edpt.FileGlmLicenseLocal
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.File = d.Get("file").(string)
	ret.Inst.FileHandle = d.Get("file_handle").(string)
	return ret
}
