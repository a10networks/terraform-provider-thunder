package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingEmailBuffer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_email_buffer`: Logging via email buffering settings\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingEmailBufferCreate,
		UpdateContext: resourceLoggingEmailBufferUpdate,
		ReadContext:   resourceLoggingEmailBufferRead,
		DeleteContext: resourceLoggingEmailBufferDelete,

		Schema: map[string]*schema.Schema{
			"number": {
				Type: schema.TypeInt, Optional: true, Default: 50, Description: "Number of log messages that can be buffered (Number of log messages that can be buffered, default 50)",
			},
			"time": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Number of minutes a log message can stay in buffer (Number of minutes a log message can stay in buffer, default 10)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingEmailBufferCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailBufferCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailBuffer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailBufferRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingEmailBufferUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailBufferUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailBuffer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingEmailBufferRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingEmailBufferDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailBufferDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailBuffer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingEmailBufferRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingEmailBufferRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingEmailBuffer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingEmailBuffer(d *schema.ResourceData) edpt.LoggingEmailBuffer {
	var ret edpt.LoggingEmailBuffer
	ret.Inst.Number = d.Get("number").(int)
	ret.Inst.Time = d.Get("time").(int)
	//omit uuid
	return ret
}
