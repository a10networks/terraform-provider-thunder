package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCmcov() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cmcov`: CM code coverage\n\n__PLACEHOLDER__",
		CreateContext: resourceCmcovCreate,
		UpdateContext: resourceCmcovUpdate,
		ReadContext:   resourceCmcovRead,
		DeleteContext: resourceCmcovDelete,

		Schema: map[string]*schema.Schema{
			"dump": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dump code coverage data and and generate report}",
			},
			"export": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export code coverage report",
			},
		},
	}
}
func resourceCmcovCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCmcovCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCmcov(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCmcovRead(ctx, d, meta)
	}
	return diags
}

func resourceCmcovUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCmcovUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCmcov(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCmcovRead(ctx, d, meta)
	}
	return diags
}
func resourceCmcovDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCmcovDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCmcov(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCmcovRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCmcovRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCmcov(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCmcov(d *schema.ResourceData) edpt.Cmcov {
	var ret edpt.Cmcov
	ret.Inst.Dump = d.Get("dump").(int)
	ret.Inst.Export = d.Get("export").(int)
	return ret
}
