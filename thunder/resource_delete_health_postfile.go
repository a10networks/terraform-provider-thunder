package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteHealthPostfile() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_health_postfile`: Address the HTTP Post data file\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteHealthPostfileCreate,
		UpdateContext: resourceDeleteHealthPostfileUpdate,
		ReadContext:   resourceDeleteHealthPostfileRead,
		DeleteContext: resourceDeleteHealthPostfileDelete,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the File Name",
			},
		},
	}
}
func resourceDeleteHealthPostfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthPostfileCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthPostfile(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteHealthPostfileRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteHealthPostfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthPostfileUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthPostfile(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteHealthPostfileRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteHealthPostfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthPostfileDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthPostfile(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteHealthPostfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthPostfileRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthPostfile(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteHealthPostfile(d *schema.ResourceData) edpt.DeleteHealthPostfile {
	var ret edpt.DeleteHealthPostfile
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
