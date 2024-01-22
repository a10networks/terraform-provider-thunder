package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTemplateLoggingSessionPeriodicLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_template_logging_session_periodic_log`: Periodic log for long lived sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTemplateLoggingSessionPeriodicLogCreate,
		UpdateContext: resourceFwTemplateLoggingSessionPeriodicLogUpdate,
		ReadContext:   resourceFwTemplateLoggingSessionPeriodicLogRead,
		DeleteContext: resourceFwTemplateLoggingSessionPeriodicLogDelete,

		Schema: map[string]*schema.Schema{
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "Logging time interval (minutes) for long lived sessions",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceFwTemplateLoggingSessionPeriodicLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSessionPeriodicLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSessionPeriodicLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingSessionPeriodicLogRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTemplateLoggingSessionPeriodicLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSessionPeriodicLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSessionPeriodicLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingSessionPeriodicLogRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTemplateLoggingSessionPeriodicLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSessionPeriodicLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSessionPeriodicLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTemplateLoggingSessionPeriodicLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingSessionPeriodicLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingSessionPeriodicLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwTemplateLoggingSessionPeriodicLog(d *schema.ResourceData) edpt.FwTemplateLoggingSessionPeriodicLog {
	var ret edpt.FwTemplateLoggingSessionPeriodicLog
	ret.Inst.Interval = d.Get("interval").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
