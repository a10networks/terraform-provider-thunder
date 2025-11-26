package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_config_filter`: Filter for config-replace-{start/end}\n\n__PLACEHOLDER__",
		CreateContext: resourceConfigFilterCreate,
		UpdateContext: resourceConfigFilterUpdate,
		ReadContext:   resourceConfigFilterRead,
		DeleteContext: resourceConfigFilterDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'skip': Skip current lineage during replace; 'accept': Skip other lineages during replace;",
			},
			"lineage": {
				Type: schema.TypeString, Optional: true, Description: "lineage to filter",
			},
		},
	}
}
func resourceConfigFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConfigFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceConfigFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConfigFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceConfigFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceConfigFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointConfigFilter(d *schema.ResourceData) edpt.ConfigFilter {
	var ret edpt.ConfigFilter
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Lineage = d.Get("lineage").(string)
	return ret
}
