package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingBuffered() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_buffered`: Set buffered logging parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingBufferedCreate,
		UpdateContext: resourceLoggingBufferedUpdate,
		ReadContext:   resourceLoggingBufferedRead,
		DeleteContext: resourceLoggingBufferedDelete,

		Schema: map[string]*schema.Schema{
			"buffersize": {
				Type: schema.TypeInt, Optional: true, Default: 30000, Description: "Logging buffer size (in items), default 30000",
			},
			"levelname": {
				Type: schema.TypeString, Optional: true, Default: "debugging", Description: "'disable': Do not send log to buffer; 'emergency': System unusable log messages      (severity=0); 'alert': Action must be taken immediately  (severity=1); 'critical': Critical conditions               (severity=2); 'error': Error conditions                  (severity=3); 'warning': Warning conditions                (severity=4); 'notification': Normal but significant conditions (severity=5); 'information': Informational messages            (severity=6); 'debugging': Debug level messages              (severity=7);",
			},
			"partition_buffersize": {
				Type: schema.TypeInt, Optional: true, Default: 3000, Description: "Logging buffer size (in items), default 3000",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingBufferedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingBufferedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingBuffered(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingBufferedRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingBufferedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingBufferedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingBuffered(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingBufferedRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingBufferedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingBufferedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingBuffered(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingBufferedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingBufferedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingBuffered(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingBuffered(d *schema.ResourceData) edpt.LoggingBuffered {
	var ret edpt.LoggingBuffered
	ret.Inst.Buffersize = d.Get("buffersize").(int)
	ret.Inst.Levelname = d.Get("levelname").(string)
	ret.Inst.PartitionBuffersize = d.Get("partition_buffersize").(int)
	//omit uuid
	return ret
}
