package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthExternalEdit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_external_edit`: Edit external health monitor script\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthExternalEditCreate,
		UpdateContext: resourceHealthExternalEditUpdate,
		ReadContext:   resourceHealthExternalEditRead,
		DeleteContext: resourceHealthExternalEditDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Describe health external monitor script briefly",
			},
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "External health monitor script file name",
			},
		},
	}
}
func resourceHealthExternalEditCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalEditCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalEdit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalEditRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthExternalEditUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalEditUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalEdit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthExternalEditRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthExternalEditDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalEditDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalEdit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthExternalEditRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthExternalEditRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthExternalEdit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthExternalEdit(d *schema.ResourceData) edpt.HealthExternalEdit {
	var ret edpt.HealthExternalEdit
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
