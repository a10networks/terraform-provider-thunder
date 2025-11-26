package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LoggingPoolBasedLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_logging_pool_based_log`: Change logging period for pool-based logs, e.g. NAT Quota Exceeded\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LoggingPoolBasedLogCreate,
		UpdateContext: resourceCgnv6LoggingPoolBasedLogUpdate,
		ReadContext:   resourceCgnv6LoggingPoolBasedLogRead,
		DeleteContext: resourceCgnv6LoggingPoolBasedLogDelete,

		Schema: map[string]*schema.Schema{
			"cycle": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Logging cycle",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LoggingPoolBasedLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingPoolBasedLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingPoolBasedLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingPoolBasedLogRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LoggingPoolBasedLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingPoolBasedLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingPoolBasedLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingPoolBasedLogRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LoggingPoolBasedLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingPoolBasedLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingPoolBasedLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LoggingPoolBasedLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingPoolBasedLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingPoolBasedLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LoggingPoolBasedLog(d *schema.ResourceData) edpt.Cgnv6LoggingPoolBasedLog {
	var ret edpt.Cgnv6LoggingPoolBasedLog
	ret.Inst.Cycle = d.Get("cycle").(int)
	//omit uuid
	return ret
}
