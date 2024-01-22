package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUpgradeCf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_upgrade_cf`: Compact flash\n\n__PLACEHOLDER__",
		CreateContext: resourceUpgradeCfCreate,
		UpdateContext: resourceUpgradeCfUpdate,
		ReadContext:   resourceUpgradeCfRead,
		DeleteContext: resourceUpgradeCfDelete,

		Schema: map[string]*schema.Schema{
			"device": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"image": {
				Type: schema.TypeString, Optional: true, Description: "'pri': Primary image;",
			},
			"image_file": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"local": {
				Type: schema.TypeString, Optional: true, Description: "Use image from local VCS image repository (Specify an image name, format: aximage_XX_XX_XX_XX.tar.gz)",
			},
			"reboot_after_upgrade": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "reboot system after upgrade is done",
			},
			"staggered_upgrade_mode": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "in staggered upgrade mode",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceUpgradeCfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeCfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeCf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceUpgradeCfRead(ctx, d, meta)
	}
	return diags
}

func resourceUpgradeCfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeCfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeCf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceUpgradeCfRead(ctx, d, meta)
	}
	return diags
}
func resourceUpgradeCfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeCfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeCf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceUpgradeCfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeCfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeCf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointUpgradeCf(d *schema.ResourceData) edpt.UpgradeCf {
	var ret edpt.UpgradeCf
	ret.Inst.Device = d.Get("device").(int)
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Image = d.Get("image").(string)
	ret.Inst.ImageFile = d.Get("image_file").(string)
	ret.Inst.Local = d.Get("local").(string)
	ret.Inst.RebootAfterUpgrade = d.Get("reboot_after_upgrade").(int)
	ret.Inst.StaggeredUpgradeMode = d.Get("staggered_upgrade_mode").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
