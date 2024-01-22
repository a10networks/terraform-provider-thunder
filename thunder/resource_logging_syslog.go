package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingSyslog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_syslog`: Set logging level which sent to syslog host\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingSyslogCreate,
		UpdateContext: resourceLoggingSyslogUpdate,
		ReadContext:   resourceLoggingSyslogRead,
		DeleteContext: resourceLoggingSyslogDelete,

		Schema: map[string]*schema.Schema{
			"syslog_levelname": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'disable': Do not send log to syslog; 'emergency': System unusable log messages      (severity=0); 'alert': Action must be taken immediately  (severity=1); 'critical': Critical conditions               (severity=2); 'error': Error conditions                  (severity=3); 'warning': Warning conditions                (severity=4); 'notification': Normal but significant conditions (severity=5); 'information': Informational messages            (severity=6); 'debugging': Debug level messages              (severity=7);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingSyslogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingSyslogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingSyslog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingSyslogRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingSyslogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingSyslogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingSyslog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingSyslogRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingSyslogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingSyslogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingSyslog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingSyslogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingSyslogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingSyslog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingSyslog(d *schema.ResourceData) edpt.LoggingSyslog {
	var ret edpt.LoggingSyslog
	ret.Inst.SyslogLevelname = d.Get("syslog_levelname").(string)
	//omit uuid
	return ret
}
