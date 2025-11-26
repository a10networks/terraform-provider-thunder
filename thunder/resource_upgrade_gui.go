package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUpgradeGui() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_upgrade_gui`: Upgrade or Rollback GUI\n\n__PLACEHOLDER__",
		CreateContext: resourceUpgradeGuiCreate,
		UpdateContext: resourceUpgradeGuiUpdate,
		ReadContext:   resourceUpgradeGuiRead,
		DeleteContext: resourceUpgradeGuiDelete,

		Schema: map[string]*schema.Schema{
			"delete": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file": {
							Type: schema.TypeString, Optional: true, Description: "Delete one local GUI image",
						},
					},
				},
			},
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"gui_upload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "image file from AXAPI",
			},
			"image": {
				Type: schema.TypeString, Optional: true, Description: "'pri': Primary image; 'sec': Secondary image;",
			},
			"image_file": {
				Type: schema.TypeString, Optional: true, Description: "image file from AXAPI,",
			},
			"local": {
				Type: schema.TypeString, Optional: true, Description: "Local GUI image name",
			},
			"remote_url": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"rollback": {
				Type: schema.TypeString, Optional: true, Description: "Rollback to a specific local GUI image",
			},
			"source_ip_address": {
				Type: schema.TypeString, Optional: true, Description: "Source IP address",
			},
			"upload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Upload GUI image from remote",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceUpgradeGuiCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeGuiCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeGui(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceUpgradeGuiRead(ctx, d, meta)
	}
	return diags
}

func resourceUpgradeGuiUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeGuiUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeGui(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceUpgradeGuiRead(ctx, d, meta)
	}
	return diags
}
func resourceUpgradeGuiDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeGuiDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeGui(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceUpgradeGuiRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceUpgradeGuiRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointUpgradeGui(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceUpgradeGuiDelete(d []interface{}) []edpt.UpgradeGuiDelete {

	count1 := len(d)
	ret := make([]edpt.UpgradeGuiDelete, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.UpgradeGuiDelete
		oi.File = in["file"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointUpgradeGui(d *schema.ResourceData) edpt.UpgradeGui {
	var ret edpt.UpgradeGui
	ret.Inst.Delete = getSliceUpgradeGuiDelete(d.Get("delete").([]interface{}))
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.GuiUpload = d.Get("gui_upload").(int)
	ret.Inst.Image = d.Get("image").(string)
	ret.Inst.ImageFile = d.Get("image_file").(string)
	ret.Inst.Local = d.Get("local").(string)
	ret.Inst.RemoteUrl = d.Get("remote_url").(string)
	ret.Inst.Rollback = d.Get("rollback").(string)
	ret.Inst.SourceIpAddress = d.Get("source_ip_address").(string)
	ret.Inst.Upload = d.Get("upload").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
