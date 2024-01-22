package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeleteHealthExternal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_delete_health_external`: Address the External Script Program\n\n__PLACEHOLDER__",
		CreateContext: resourceDeleteHealthExternalCreate,
		UpdateContext: resourceDeleteHealthExternalUpdate,
		ReadContext:   resourceDeleteHealthExternalRead,
		DeleteContext: resourceDeleteHealthExternalDelete,

		Schema: map[string]*schema.Schema{
			"file_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify the Program Name",
			},
		},
	}
}
func resourceDeleteHealthExternalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthExternalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthExternal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteHealthExternalRead(ctx, d, meta)
	}
	return diags
}

func resourceDeleteHealthExternalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthExternalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthExternal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDeleteHealthExternalRead(ctx, d, meta)
	}
	return diags
}
func resourceDeleteHealthExternalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthExternalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthExternal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDeleteHealthExternalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDeleteHealthExternalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDeleteHealthExternal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDeleteHealthExternal(d *schema.ResourceData) edpt.DeleteHealthExternal {
	var ret edpt.DeleteHealthExternal
	ret.Inst.FileName = d.Get("file_name").(string)
	return ret
}
