package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLsnLogSuppression() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_lsn_log_suppression`: Set LSN system log generation parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingLsnLogSuppressionCreate,
		UpdateContext: resourceLoggingLsnLogSuppressionUpdate,
		ReadContext:   resourceLoggingLsnLogSuppressionRead,
		DeleteContext: resourceLoggingLsnLogSuppressionDelete,

		Schema: map[string]*schema.Schema{
			"count1": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Configure log suppression count (default: 100)",
			},
			"time": {
				Type: schema.TypeInt, Optional: true, Default: 30, Description: "Log generation timeout(default: 30 seconds)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingLsnLogSuppressionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnLogSuppressionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnLogSuppression(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLsnLogSuppressionRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingLsnLogSuppressionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnLogSuppressionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnLogSuppression(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLsnLogSuppressionRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingLsnLogSuppressionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnLogSuppressionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnLogSuppression(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingLsnLogSuppressionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLsnLogSuppressionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLsnLogSuppression(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingLsnLogSuppression(d *schema.ResourceData) edpt.LoggingLsnLogSuppression {
	var ret edpt.LoggingLsnLogSuppression
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.Time = d.Get("time").(int)
	//omit uuid
	return ret
}
