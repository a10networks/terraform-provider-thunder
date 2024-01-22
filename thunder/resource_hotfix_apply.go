package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHotfixApply() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_hotfix_apply`: apply a hotfix\n\n__PLACEHOLDER__",
		CreateContext: resourceHotfixApplyCreate,
		UpdateContext: resourceHotfixApplyUpdate,
		ReadContext:   resourceHotfixApplyRead,
		DeleteContext: resourceHotfixApplyDelete,

		Schema: map[string]*schema.Schema{
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"image_file": {
				Type: schema.TypeString, Optional: true, Description: "image file from AXAPI",
			},
			"source_ip_address": {
				Type: schema.TypeString, Optional: true, Description: "Source ip address",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceHotfixApplyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixApplyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfixApply(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHotfixApplyRead(ctx, d, meta)
	}
	return diags
}

func resourceHotfixApplyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixApplyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfixApply(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHotfixApplyRead(ctx, d, meta)
	}
	return diags
}
func resourceHotfixApplyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixApplyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfixApply(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHotfixApplyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHotfixApplyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHotfixApply(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHotfixApply(d *schema.ResourceData) edpt.HotfixApply {
	var ret edpt.HotfixApply
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.ImageFile = d.Get("image_file").(string)
	ret.Inst.SourceIpAddress = d.Get("source_ip_address").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
