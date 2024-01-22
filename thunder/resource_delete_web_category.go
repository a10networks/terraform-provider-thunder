package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteWebCategory() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_web_category`: Delete web-category database\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteWebCategoryCreate,
		UpdateContext: resourceDeleteWebCategoryUpdate,
		ReadContext:   resourceDeleteWebCategoryRead,
		DeleteContext: resourceDeleteWebCategoryDelete,

		Schema: map[string]*schema.Schema{
			"database": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete web-category database",
			},
		},
	}
}
func resourceDeleteWebCategoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteWebCategoryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteWebCategory(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteWebCategoryRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteWebCategoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteWebCategoryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteWebCategory(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteWebCategoryRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteWebCategoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteWebCategoryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteWebCategory(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteWebCategoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteWebCategoryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteWebCategory(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteWebCategory(d *schema.ResourceData) edpt.DeleteWebCategory {
	var ret edpt.DeleteWebCategory
	ret.Inst.Database = d.Get("database").(int)
	return ret
}
