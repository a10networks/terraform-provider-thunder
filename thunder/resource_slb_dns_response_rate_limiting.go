package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsResponseRateLimiting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_dns_response_rate_limiting`: Configure DNS Response-Rate-Limiting\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbDnsResponseRateLimitingCreate,
		UpdateContext: resourceSlbDnsResponseRateLimitingUpdate,
		ReadContext:   resourceSlbDnsResponseRateLimitingRead,
		DeleteContext: resourceSlbDnsResponseRateLimitingDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_entries': Current Entry Count; 'total_created': Total Entry Created; 'total_inserted': Total Entry Inserted; 'total_withdrew': Total Entry Withdrew; 'total_ready_to_free': Total Entry Ready To Free; 'total_freed': Total Entry Freed; 'total_logs': Total Logs; 'total_overflow_entry_hits': Total Overflow Entry Hits; 'total_refill': Total Refills; 'total_credit_exceeded': Total Credit Exceeded; 'other_thread_refill': Other Thread Refilling; 'err_entry_create_failed': Entry Creation Failure; 'err_entry_create_oom': Entry Creation Out of Memory; 'err_entry_ext_create_oom': Entry Extension Creation Out of Memory; 'err_entry_insert_failed': Entry Insert Failed; 'err_vport_fail_match': Failed to Match Vport;",
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
func resourceSlbDnsResponseRateLimitingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsResponseRateLimitingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsResponseRateLimiting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDnsResponseRateLimitingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbDnsResponseRateLimitingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsResponseRateLimitingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsResponseRateLimiting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbDnsResponseRateLimitingRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbDnsResponseRateLimitingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsResponseRateLimitingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsResponseRateLimiting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbDnsResponseRateLimitingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsResponseRateLimitingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsResponseRateLimiting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbDnsResponseRateLimitingSamplingEnable(d []interface{}) []edpt.SlbDnsResponseRateLimitingSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbDnsResponseRateLimitingSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsResponseRateLimitingSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDnsResponseRateLimiting(d *schema.ResourceData) edpt.SlbDnsResponseRateLimiting {
	var ret edpt.SlbDnsResponseRateLimiting
	ret.Inst.SamplingEnable = getSliceSlbDnsResponseRateLimitingSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
