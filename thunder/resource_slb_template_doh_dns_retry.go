package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDohDnsRetry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_doh_dns_retry`: DNS over HTTP(s) template retry policy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDohDnsRetryCreate,
		UpdateContext: resourceSlbTemplateDohDnsRetryUpdate,
		ReadContext:   resourceSlbTemplateDohDnsRetryRead,
		DeleteContext: resourceSlbTemplateDohDnsRetryDelete,

		Schema: map[string]*schema.Schema{
			"after_timeout": {
				Type: schema.TypeString, Optional: true, Default: "close", Description: "'close': Close client side connection; 'retry-with-tcp': Retry DNS query to server using TCP (If UDP was tried initially. Close after.);",
			},
			"max_trials": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Total number of times to try DNS query to server before closing client connection, default 3",
			},
			"retry_interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "DNS Retry Interval value 1 - 400 in units of 100ms, default is 10 (default is 1000ms) (1 - 400 in units of 100ms, default is 10 (1000ms/1sec))",
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
func resourceSlbTemplateDohDnsRetryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohDnsRetryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohDnsRetry(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDohDnsRetryRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDohDnsRetryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohDnsRetryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohDnsRetry(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDohDnsRetryRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDohDnsRetryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohDnsRetryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohDnsRetry(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDohDnsRetryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohDnsRetryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohDnsRetry(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDohDnsRetry(d *schema.ResourceData) edpt.SlbTemplateDohDnsRetry {
	var ret edpt.SlbTemplateDohDnsRetry
	ret.Inst.AfterTimeout = d.Get("after_timeout").(string)
	ret.Inst.MaxTrials = d.Get("max_trials").(int)
	ret.Inst.RetryInterval = d.Get("retry_interval").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
