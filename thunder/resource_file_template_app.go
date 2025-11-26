package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileTemplateApp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_file_template_app`: ACT action axAPI\n\n__PLACEHOLDER__",
		CreateContext: resourceFileTemplateAppCreate,
		UpdateContext: resourceFileTemplateAppUpdate,
		ReadContext:   resourceFileTemplateAppRead,
		DeleteContext: resourceFileTemplateAppDelete,

		Schema: map[string]*schema.Schema{
			"act_name": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"action": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"version": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
		},
	}
}
func resourceFileTemplateAppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateAppCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateApp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTemplateAppRead(ctx, d, meta)
	}
	return diags
}

func resourceFileTemplateAppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateAppUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateApp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFileTemplateAppRead(ctx, d, meta)
	}
	return diags
}
func resourceFileTemplateAppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateAppDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateApp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFileTemplateAppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFileTemplateAppRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFileTemplateApp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFileTemplateApp(d *schema.ResourceData) edpt.FileTemplateApp {
	var ret edpt.FileTemplateApp
	ret.Inst.ActName = d.Get("act_name").(string)
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Version = d.Get("version").(string)
	return ret
}
