package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteNgWafCustomPage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_ng_waf_custom_page`: Custom html file for ng-waf\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteNgWafCustomPageCreate,
		UpdateContext: resourceDeleteNgWafCustomPageUpdate,
		ReadContext:   resourceDeleteNgWafCustomPageRead,
		DeleteContext: resourceDeleteNgWafCustomPageDelete,

		Schema: map[string]*schema.Schema{
			"page_name": {
				Type: schema.TypeString, Optional: true, Description: "Local File Name",
			},
		},
	}
}
func resourceDeleteNgWafCustomPageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteNgWafCustomPageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteNgWafCustomPage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteNgWafCustomPageRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteNgWafCustomPageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteNgWafCustomPageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteNgWafCustomPage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteNgWafCustomPageRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteNgWafCustomPageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteNgWafCustomPageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteNgWafCustomPage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteNgWafCustomPageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteNgWafCustomPageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteNgWafCustomPage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteNgWafCustomPage(d *schema.ResourceData) edpt.DeleteNgWafCustomPage {
	var ret edpt.DeleteNgWafCustomPage
	ret.Inst.PageName = d.Get("page_name").(string)
	return ret
}
