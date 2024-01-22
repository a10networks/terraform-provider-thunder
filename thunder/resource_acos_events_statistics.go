package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsStatistics() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_statistics`: acos events global statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsStatisticsCreate,
		UpdateContext: resourceAcosEventsStatisticsUpdate,
		ReadContext:   resourceAcosEventsStatisticsRead,
		DeleteContext: resourceAcosEventsStatisticsDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'msg_sent': Messages sent, to Remote; 'msg_sent_logdb': Messages sent, to LogDB; 'msg_dropped_format_not_defined': Messages Dropped, format not defined; 'msg_dropped_malloc_failure': Messages Dropped, malloc failure; 'msg_dropped_no_template': Messages Dropped, no active template; 'msg_dropped_selector': Messages Dropped, selector does not enable msg; 'msg_dropped_too_long': Messages Dropped, invalid length; 'msg_dropped_craft_fail': Messages Dropped, msg crafting failed; 'msg_dropped_local_log_ratelimit': Messages Dropped, local log ratelimited; 'msg_dropped_remote_log_ratelimit': Messages Dropped, remote log ratelimited; 'msg_dropped_send_failed': Messages Dropped, send failed; 'msg_dropped_no_active_member': Messages Dropped, no active member in collector grp; 'msg_dropped_route_fail': Messages Dropped, Route lookup failed; 'msg_dropped_other': Messages Dropped, unexpected error; 'no_template': Message API called, with no active template; 'msg_dropped_lost_during_config_change': Messages Dropped, lost during config change; 'local_enqueue_pass': Messages enqueue to Logd passed; 'msg_sent_to_logd': Messages sent to Logd via IPC; 'msg_retry_after_socket_fail': Messages retried to be sent to Logd via IPC; 'msg_sent_direct_syslog': Messages sent to syslog directly from axlog; 'msg_dropped_send_to_logd_fail': Messages Dropped, send to Logd via IPC failed; 'msg_dropped_trylock_fail': Messages Dropped, Trylock failed in axlog; 'msg_dropped_remote_cplane_log_ratelimit': Messages Dropped, Remote cplane log ratelimited; 'msg_dropped_remote_dplane_log_ratelimit': Messages Dropped, Remote dplane log ratelimited; 'msg_dropped_local_enqueue_failed': Messages Dropped, Enqueue to Logd failed; 'msg_dropped_grp_not_used': Messages Dropped, Collector group not used; 'msg_sent_remote_cplane': Messages Sent, to remote in logd; 'msg_dropped_no_template_logd': Messages Dropped, no active template in Logd; 'msg_dropped_lost_during_config_change_logd': Messages Dropped, lost during config change in Logd; 'msg_dropped_craft_fail_logd': Messages Dropped, msg crafting failed in Logd; 'msg_dropped_send_failed_logd': Messages Dropped, send failed in Logd; 'msg_dropped_no_active_member_logd': Messages Dropped, no active member in collector grp in Logd; 'msg_dropped_other_logd': Messages Dropped, unexpected error in Logd; 'msg_dropped_invalid_part': Messages Dropped, Invalid partition Id; 'acos_evt_test_logs_ticks': Number of ticks when running ACOS Event Test Logs; 'param_msg_sent_to_hc': Parameterized log sent to HC; 'param_msg_sent_fail': Parameterized log send to HC failed; 'param_msg_encode_fail': Parameterized log AVRO encoding failed;",
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
func resourceAcosEventsStatisticsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsStatisticsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsStatistics(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsStatisticsRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsStatisticsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsStatisticsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsStatistics(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsStatisticsRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsStatisticsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsStatisticsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsStatistics(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsStatisticsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsStatisticsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsStatistics(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosEventsStatisticsSamplingEnable(d []interface{}) []edpt.AcosEventsStatisticsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AcosEventsStatisticsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsStatisticsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsStatistics(d *schema.ResourceData) edpt.AcosEventsStatistics {
	var ret edpt.AcosEventsStatistics
	ret.Inst.SamplingEnable = getSliceAcosEventsStatisticsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
