package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwRateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_rate_limit`: View Rate Limit Entries\n\n__PLACEHOLDER__",
		CreateContext: resourceFwRateLimitCreate,
		UpdateContext: resourceFwRateLimitUpdate,
		ReadContext:   resourceFwRateLimitRead,
		DeleteContext: resourceFwRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'ratelimit_used_total_mem': Total Memory Used For Rate-limiting (bytes); 'ratelimit_used_spm_mem': Total SPM Memory Used For Rate-limiting Infra in Bytes; 'ratelimit_used_heap_mem': Total Heap Memory Used For Rate-limiting Infra in Bytes; 'ratelimit_entry_alloc_frm_spm_mem': Total Number of Rate-limit Entries created using SPM Memory; 'ratelimit_high_accurate_entry_alloc_fail': Total Number of Failures to Create Highly Accurate Rate-limit Entries Due to Memory Allocation Failures; 'ratelimit_high_perf_entry_alloc_fail': Total Number of Failures to Create High-Perf Rate-limit Entries Due to Memory Allocation Failures; 'ratelimit_high_perf_entry_secondary_alloc_fail': Total Number of Failures to Allocate Additional Memory to Existing High-Perf Rate-limit Entries; 'ratelimit_entry_alloc_fail_rate_too_high': Total Number of Attempts to Configure Too High Rate Limits; 'ratelimit_entry_alloc_fail_metric_count_gt_supported': Total Number of Failures to Create High-Perf Rate-limit Entries Because of Too Many Metrics; 'ratelimit_entry_count_t2_key': Number of Total Rate-limit Entries; 'ratelimit_entry_count_fw_rule_uid': Number of Rate-limit Entries with Scope Aggregate; 'ratelimit_entry_count_ip_addr': Number of Rate-limit Entries with Scope IPv4 Address; 'ratelimit_entry_count_ip6_addr': Number of Rate-limit Entries with Scope IPv6 Address; 'ratelimit_entry_count_session_id': Number of Rate-limit Entries with Scope Session ID; 'ratelimit_entry_count_rule_ipv4_prefix': Number of Rate-limit Entries with Scope IPv4 Prefix; 'ratelimit_entry_count_rule_ipv6_prefix': Number of Rate-limit Entries with Scope IPv6 Prefix; 'ratelimit_entry_count_parent_uid': Number of Parent Rate-limit Entries with Scope Aggregate; 'ratelimit_entry_count_parent_ipv4_prefix': Number of Parent Rate-limit Entries with Scope IPv4 Prefix; 'ratelimit_entry_count_parent_ipv6_prefix': Number of Parent Rate-limit Entries with Scope IPv6 Prefix; 'ratelimit_infra_generic_errors': Current Number of Generic Errors Encountered in Ratelimit Infra; 'ratelimit_entry_count_allocated': Number of Rate-limit Entries Allocated Totally; 'ratelimit_entry_count_freed': Number of Rate-limit Entries Freed Totally; 'ratelimit_entry_count_rule_ip': Number of Rate-limit Entries with Scope IP; 'ratelimit_entry_count_parent_ip': Number of Parent Rate-limit Entries with Scope IP; 'ratelimit_entry_count_radius_usergroup': The total number of rate-limiting entries with the scope RADIUS user group.; 'ratelimit_entry_count_parent_radius_usergroup': The total number of parent rate-limiting entries with the scope RADIUS user group.; 'ratelimit_entry_count_radius_userid': The total number of rate-limiting entries with the scope RADIUS user ID.; 'ratelimit_entry_count_parent_radius_userid': The total number of parent rate-limiting entries with the scope RADIUS user ID.;",
						},
					},
				},
			},
			"summary": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceFwRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwRateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwRateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwRateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceFwRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwRateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwRateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwRateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceFwRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwRateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwRateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwRateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwRateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwRateLimitSamplingEnable(d []interface{}) []edpt.FwRateLimitSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwRateLimitSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwRateLimitSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwRateLimitSummary436(d []interface{}) edpt.FwRateLimitSummary436 {

	var ret edpt.FwRateLimitSummary436
	return ret
}

func dataToEndpointFwRateLimit(d *schema.ResourceData) edpt.FwRateLimit {
	var ret edpt.FwRateLimit
	ret.Inst.SamplingEnable = getSliceFwRateLimitSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Summary = getObjectFwRateLimitSummary436(d.Get("summary").([]interface{}))
	//omit uuid
	return ret
}
