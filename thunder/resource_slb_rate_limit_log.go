package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRateLimitLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_rate_limit_log`: Configure rate limit logging\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbRateLimitLogCreate,
		UpdateContext: resourceSlbRateLimitLogUpdate,
		ReadContext:   resourceSlbRateLimitLogRead,
		DeleteContext: resourceSlbRateLimitLogDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_log_times': Total log times; 'total_log_msg': Total log messages; 'local_log_msg': Local log messages; 'remote_log_msg': Remote log messages; 'local_log_rate': Local rate (per sec); 'remote_log_rate': Remote rate (per sec); 'msg_too_big': Log message too big; 'buff_alloc_fail': Buffer alloc fail; 'no_route': No route; 'buff_send_fail': Buffer send fail; 'alloc_conn': Log-session alloc; 'free_conn': Log-session free; 'conn_alloc_fail': Log-session alloc fail; 'no_repeat_msg': No repeat message; 'local_log_dropped': Local log dropped due to rate-limit;",
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
func resourceSlbRateLimitLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRateLimitLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRateLimitLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbRateLimitLogRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbRateLimitLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRateLimitLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRateLimitLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbRateLimitLogRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbRateLimitLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRateLimitLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRateLimitLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbRateLimitLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRateLimitLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRateLimitLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbRateLimitLogSamplingEnable(d []interface{}) []edpt.SlbRateLimitLogSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbRateLimitLogSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRateLimitLogSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbRateLimitLog(d *schema.ResourceData) edpt.SlbRateLimitLog {
	var ret edpt.SlbRateLimitLog
	ret.Inst.SamplingEnable = getSliceSlbRateLimitLogSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
