package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServiceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_service_group`: Authentication service group\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServiceGroupCreate,
		UpdateContext: resourceAamAuthenticationServiceGroupUpdate,
		ReadContext:   resourceAamAuthenticationServiceGroupRead,
		DeleteContext: resourceAamAuthenticationServiceGroupDelete,

		Schema: map[string]*schema.Schema{
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable health check",
			},
			"lb_method": {
				Type: schema.TypeString, Optional: true, Description: "'round-robin': Round robin on server level;",
			},
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Member name",
						},
						"port": {
							Type: schema.TypeInt, Required: true, Description: "Port number",
						},
						"member_state": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable member service port; 'disable': Disable member service port;",
						},
						"member_priority": {
							Type: schema.TypeInt, Optional: true, Description: "Priority of Port in the Group",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total established connections; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_status_code_2xx': Total reverse packets inspected status code 2xx; 'total_rev_pkts_inspected_status_code_non_5xx': Total reverse packets inspected status code non 5xx; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total requests successful; 'peak_conn': peak_conn; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'curr_conn_overflow': Current connection counter overflow count;",
									},
								},
							},
						},
						"packet_capture_template": {
							Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify AAM service group name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP AAM service; 'udp': UDP AAM service;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'server_selection_fail_drop': Drops due to Service selection failure; 'server_selection_fail_reset': Resets sent out for Service selection failure; 'service_peak_conn': Peak connection count for the Service Group; 'service_healthy_host': Service Group healthy host count; 'service_unhealthy_host': Service Group unhealthy host count; 'service_req_count': Service Group request count; 'service_resp_count': Service Group response count; 'service_resp_2xx': Service Group response 2xx count; 'service_resp_3xx': Service Group response 3xx count; 'service_resp_4xx': Service Group response 4xx count; 'service_resp_5xx': Service Group response 5xx count; 'service_curr_conn_overflow': Current connection counter overflow count;",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServiceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServiceGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServiceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServiceGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServiceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationServiceGroupMemberList(d []interface{}) []edpt.AamAuthenticationServiceGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServiceGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServiceGroupMemberList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		oi.MemberState = in["member_state"].(string)
		oi.MemberPriority = in["member_priority"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceAamAuthenticationServiceGroupMemberListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServiceGroupMemberListSamplingEnable(d []interface{}) []edpt.AamAuthenticationServiceGroupMemberListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServiceGroupMemberListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServiceGroupMemberListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthenticationServiceGroupSamplingEnable(d []interface{}) []edpt.AamAuthenticationServiceGroupSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServiceGroupSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServiceGroupSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServiceGroup(d *schema.ResourceData) edpt.AamAuthenticationServiceGroup {
	var ret edpt.AamAuthenticationServiceGroup
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.LbMethod = d.Get("lb_method").(string)
	ret.Inst.MemberList = getSliceAamAuthenticationServiceGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServiceGroupSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
