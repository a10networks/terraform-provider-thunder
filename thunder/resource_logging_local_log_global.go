package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_logging_local_log_global`: \n\n__PLACEHOLDER__",
		CreateContext: resourceLoggingLocalLogGlobalCreate,
		UpdateContext: resourceLoggingLocalLogGlobalUpdate,
		ReadContext:   resourceLoggingLocalLogGlobalRead,
		DeleteContext: resourceLoggingLocalLogGlobalDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'enqueue': Total local-log enqueue; 'enqueue-full': Total local-log queue full; 'enqueue-error': Total local-log enqueue error; 'dequeue': Total local-log dequeue; 'dequeue-error': Total local-log dequeue processing error; 'raw-log': Total local-log raw logs; 'raw-log-error': Total raw log logging error; 'log-summarized': Total raw log summarized; 'l1-log-summarized': Total layer 1 log summarized; 'l2-log-summarized': Total layer 2 log summarized; 'log-summarized-error': Total local-log summarization error; 'aam-db': Total local-log AAM raw database; 'ep-db': Total local-log EP raw database; 'fw-db': Total local-log Firewall raw database; 'aam-top-user-db': Total local-log AAM top user summary database; 'ep-top-user-db': Total local-log EP top user summary database; 'ep-top-src-db': Total local-log EP top client summary database; 'ep-top-dst-db': Total local-log EP top destination summary database; 'ep-top-domain-db': Total local-log EP top domain summary database; 'ep-top-web-category-db': Total local-log EP top web-category summary database; 'ep-top-host-db': Total local-log EP top host summary database; 'fw-top-app-db': Total local-log Firewall top application summary database; 'fw-top-src-db': Total local-log Firewall top source summary database; 'fw-top-app-src-db': Total local-log Firewall top application and source summary database; 'fw-top-category-db': Total local-log Firewall top category summary database; 'db-erro': Total local-log database create error; 'query': Total local-log axapi query; 'response': Total local-log axapi response; 'query-error': Total local-log axapi query error; 'fw-top-thr-db': Total local-log Firewall top threat summary database; 'fw-top-thr-src-db': Total local-log Firewall top threat and source summary database;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLoggingLocalLogGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLocalLogGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceLoggingLocalLogGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLoggingLocalLogGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceLoggingLocalLogGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLoggingLocalLogGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceLoggingLocalLogGlobalSamplingEnable(d []interface{}) []edpt.LoggingLocalLogGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.LoggingLocalLogGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.LoggingLocalLogGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointLoggingLocalLogGlobal(d *schema.ResourceData) edpt.LoggingLocalLogGlobal {
	var ret edpt.LoggingLocalLogGlobal
	ret.Inst.SamplingEnable = getSliceLoggingLocalLogGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
