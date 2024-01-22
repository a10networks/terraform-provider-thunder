package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHotfix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_hotfix`: hotfix System\n\n__PLACEHOLDER__",
		CreateContext: resourceHotfixCreate,
		UpdateContext: resourceHotfixUpdate,
		ReadContext:   resourceHotfixRead,
		DeleteContext: resourceHotfixDelete,

		Schema: map[string]*schema.Schema{
			"apply": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"source_ip_address": {
							Type: schema.TypeString, Optional: true, Description: "Source ip address",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
						"image_file": {
							Type: schema.TypeString, Optional: true, Description: "image file from AXAPI",
						},
					},
				},
			},
			"revoke": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "revoke a hotfix",
			},
		},
	}
}
func resourceHotfixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHotfixRead(ctx, d, meta)
	}
	return diags
}

func resourceHotfixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHotfixRead(ctx, d, meta)
	}
	return diags
}
func resourceHotfixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHotfixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHotfixApply438(d []interface{}) edpt.HotfixApply438 {

	count1 := len(d)
	var ret edpt.HotfixApply438
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.SourceIpAddress = in["source_ip_address"].(string)
		ret.FileUrl = in["file_url"].(string)
		ret.ImageFile = in["image_file"].(string)
	}
	return ret
}

func dataToEndpointHotfix(d *schema.ResourceData) edpt.Hotfix {
	var ret edpt.Hotfix
	ret.Inst.Apply = getObjectHotfixApply438(d.Get("apply").([]interface{}))
	ret.Inst.Revoke = d.Get("revoke").(int)
	return ret
}
