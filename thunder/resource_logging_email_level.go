package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingEmailLevel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_email_level`: Set logging level which sent to email address\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingEmailLevelCreate,
		UpdateContext: resourceLoggingEmailLevelUpdate,
		ReadContext:   resourceLoggingEmailLevelRead,
		DeleteContext: resourceLoggingEmailLevelDelete,

		Schema: map[string]*schema.Schema{
			"email_levelname": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'disable': Do not send log to email address; 'emergency': System unusable log messages      (severity=0); 'alert': Action must be taken immediately   (severity=1); 'critical': Critical conditions      (severity=2); 'notification': Normal but significant conditions (severity=5);",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingEmailLevelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailLevelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailLevel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailLevelRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingEmailLevelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailLevelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailLevel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailLevelRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingEmailLevelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailLevelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailLevel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingEmailLevelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailLevelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailLevel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingEmailLevel(d *schema.ResourceData) edpt.LoggingEmailLevel {
	var ret edpt.LoggingEmailLevel
	ret.Inst.EmailLevelname = d.Get("email_levelname").(string)
	//omit uuid
	return ret
}
