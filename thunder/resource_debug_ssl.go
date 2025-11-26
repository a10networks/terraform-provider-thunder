package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugSsl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ssl`: Debug SSL operation\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugSslCreate,
		UpdateContext: resourceDebugSslUpdate,
		ReadContext:   resourceDebugSslRead,
		DeleteContext: resourceDebugSslDelete,

		Schema: map[string]*schema.Schema{
			"client_server": {
				Type: schema.TypeString, Optional: true, Description: "'clientside': clientside SSL connection; 'serverside': serverside SSL connection;",
			},
			"payload_dump_max": {
				Type: schema.TypeInt, Optional: true, Description: "Application payloads exceeding this limit will be truncated to dump. Set to 0 to disable payload dumping",
			},
			"payload_dump_string": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dump the application payload as strings",
			},
			"read_write": {
				Type: schema.TypeString, Optional: true, Description: "'read': read record; 'write': write record;",
			},
			"record_end": {
				Type: schema.TypeInt, Optional: true, Description: "The last record number to print debug messages. Default: 0 (no end)",
			},
			"record_start": {
				Type: schema.TypeInt, Optional: true, Description: "The first record number to debug prints messages. Default is 0",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugSslCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSslCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSsl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSslRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugSslUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSslUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSsl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSslRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugSslDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSslDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSsl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugSslRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSslRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSsl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugSsl(d *schema.ResourceData) edpt.DebugSsl {
	var ret edpt.DebugSsl
	ret.Inst.ClientServer = d.Get("client_server").(string)
	ret.Inst.PayloadDumpMax = d.Get("payload_dump_max").(int)
	ret.Inst.PayloadDumpString = d.Get("payload_dump_string").(int)
	ret.Inst.ReadWrite = d.Get("read_write").(string)
	ret.Inst.RecordEnd = d.Get("record_end").(int)
	ret.Inst.RecordStart = d.Get("record_start").(int)
	//omit uuid
	return ret
}
