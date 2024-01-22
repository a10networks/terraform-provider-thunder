package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthExternalCopy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_external_copy`: Copy external health monitor script\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthExternalCopyCreate,
		UpdateContext: resourceHealthExternalCopyUpdate,
		ReadContext:   resourceHealthExternalCopyRead,
		DeleteContext: resourceHealthExternalCopyDelete,

		Schema: map[string]*schema.Schema{
			"dst_file": {
				Type: schema.TypeString, Optional: true, Description: "Destination external health monitor script file name",
			},
			"src_file": {
				Type: schema.TypeString, Optional: true, Description: "Source external health monitor script file name",
			},
		},
	}
}
func resourceHealthExternalCopyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCopyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCopy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalCopyRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthExternalCopyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCopyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCopy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalCopyRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthExternalCopyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCopyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCopy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthExternalCopyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalCopyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalCopy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthExternalCopy(d *schema.ResourceData) edpt.HealthExternalCopy {
	var ret edpt.HealthExternalCopy
	ret.Inst.DstFile = d.Get("dst_file").(string)
	ret.Inst.SrcFile = d.Get("src_file").(string)
	return ret
}
