package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingConfigure() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_configure`: Set logging configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingConfigureCreate,
		UpdateContext: resourceLoggingConfigureUpdate,
		ReadContext:   resourceLoggingConfigureRead,
		DeleteContext: resourceLoggingConfigureDelete,

		Schema: map[string]*schema.Schema{
			"line_feed_in_udp_syslog": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Add newline for syslog over UDP",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingConfigureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingConfigureCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingConfigure(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingConfigureRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingConfigureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingConfigureUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingConfigure(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingConfigureRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingConfigureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingConfigureDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingConfigure(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingConfigureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingConfigureRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingConfigure(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointLoggingConfigure(d *schema.ResourceData) edpt.LoggingConfigure {
	var ret edpt.LoggingConfigure
	ret.Inst.LineFeedInUdpSyslog = d.Get("line_feed_in_udp_syslog").(int)
	//omit uuid
	return ret
}
