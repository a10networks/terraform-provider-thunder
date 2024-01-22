package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_dns`: DNS Global Options\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbDnsCreate,
		UpdateContext: resourceGslbDnsUpdate,
		ReadContext:   resourceGslbDnsRead,
		DeleteContext: resourceGslbDnsDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "none", Description: "'none': No action (default); 'drop': Drop query; 'reject': Send refuse response; 'ignore': Send empty response;",
			},
			"logging": {
				Type: schema.TypeString, Optional: true, Default: "none", Description: "'none': No logging (default); 'query': DNS Query; 'response': DNS Response; 'both': Both DNS Query and Response;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total-query': Total number of DNS queries received; 'total-response': Total number of DNS replies sent to clients; 'bad-packet-query': Number of queries with incorrect data length; 'bad-packet-response': Number of replies with incorrect data length; 'bad-header-query': Number of queries with incorrect header; 'bad-header-response': Number of replies with incorrect header; 'bad-format-query': Number of queries with incorrect format; 'bad-format-response': Number of replies with incorrect format; 'bad-service-query': Number of queries with unknown service; 'bad-service-response': Number of replies with unknown service; 'bad-class-query': Number of queries with incorrect class; 'bad-class-response': Number of replies with incorrect class; 'bad-type-query': Number of queries with incorrect type; 'bad-type-response': Number of replies with incorrect type; 'no_answer': Number of replies with unknown server IP; 'metric_health_check': Metric Health Check Hit; 'metric_weighted_ip': Metric Weighted IP Hit; 'metric_weighted_site': Metric Weighted Site Hit; 'metric_capacity': Metric Capacity Hit; 'metric_active_server': Metric Active Server Hit; 'metric_easy_rdt': Metric Easy RDT Hit; 'metric_active_rdt': Metric Active RDT Hit; 'metric_geographic': Metric Geographic Hit; 'metric_connection_load': Metric Connection Load Hit; 'metric_number_of_sessions': Metric Number of Sessions Hit; 'metric_active_weight': Metric Active Weight Hit; 'metric_admin_preference': Metric Admin Preference Hit; 'metric_bandwidth_quality': Metric Bandwidth Quality Hit; 'metric_bandwidth_cost': Metric Bandwidth Cost Hit; 'metric_user': Metric User Hit; 'metric_least_reponse': Metric Least Reponse Hit; 'metric_admin_ip': Metric Admin IP Hit; 'metric_round_robin': Metric Round Robin Hit;",
						},
					},
				},
			},
			"template": {
				Type: schema.TypeString, Optional: true, Description: "Logging template (Logging Template Name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbDnsSamplingEnable(d []interface{}) []edpt.GslbDnsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbDnsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbDnsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbDns(d *schema.ResourceData) edpt.GslbDns {
	var ret edpt.GslbDns
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.SamplingEnable = getSliceGslbDnsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Template = d.Get("template").(string)
	//omit uuid
	return ret
}
