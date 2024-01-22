package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingEmailFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_email_filter`: Logging via email filter settings\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingEmailFilterCreate,
		UpdateContext: resourceLoggingEmailFilterUpdate,
		ReadContext:   resourceLoggingEmailFilterRead,
		DeleteContext: resourceLoggingEmailFilterDelete,

		Schema: map[string]*schema.Schema{
			"expression": {
				Type: schema.TypeString, Optional: true, Description: "Reverse Polish Notation, consists of level 0-7, module AFLEX/HMON/..., pattern log-content-pattern, and or/and/not",
			},
			"filter_id": {
				Type: schema.TypeInt, Required: true, Description: "Logging via email filter settings",
			},
			"trigger": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Trigger email, override buffer settings",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingEmailFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingEmailFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingEmailFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingEmailFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingEmailFilter(d *schema.ResourceData) edpt.LoggingEmailFilter {
	var ret edpt.LoggingEmailFilter
	ret.Inst.Expression = d.Get("expression").(string)
	ret.Inst.FilterId = d.Get("filter_id").(int)
	ret.Inst.Trigger = d.Get("trigger").(int)
	//omit uuid
	return ret
}
