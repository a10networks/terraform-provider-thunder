package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ServiceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_service_group`: Service Group\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6ServiceGroupCreate,
		UpdateContext: resourceCgnv6ServiceGroupUpdate,
		ReadContext:   resourceCgnv6ServiceGroupRead,
		DeleteContext: resourceCgnv6ServiceGroupDelete,

		Schema: map[string]*schema.Schema{
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_conn': Current connections; 'total_fwd_bytes': Total forward bytes; 'total_fwd_pkts': Total forward packets; 'total_rev_bytes': Total reverse bytes; 'total_rev_pkts': Total reverse packets; 'total_conn': Total connections; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_status_code_2xx': Total reverse packets inspected status code 2xx; 'total_rev_pkts_inspected_status_code_non_5xx': Total reverse packets inspected status code non 5xx; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total requests success; 'peak_conn': Peak connections; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'curr_conn_overflow': Current connection counter overflow count; 'state_flaps': State flaps count;",
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "CGNV6 Service Name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP LB service; 'udp': UDP LB service;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'server_selection_fail_drop': Service selection fail drop; 'server_selection_fail_reset': Service selection fail reset; 'service_peak_conn': Service peak connection;",
						},
					},
				},
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share with partition",
			},
			"shared_group": {
				Type: schema.TypeString, Optional: true, Description: "Share with a partition group (Partition Group Name)",
			},
			"shared_partition": {
				Type: schema.TypeString, Optional: true, Description: "Share with a single partition (Partition Name)",
			},
			"traffic_replication_mirror_ip_repl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replaces IP with server-IP",
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
func resourceCgnv6ServiceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServiceGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServiceGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ServiceGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6ServiceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServiceGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServiceGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ServiceGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6ServiceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServiceGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServiceGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6ServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServiceGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServiceGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6ServiceGroupMemberList(d []interface{}) []edpt.Cgnv6ServiceGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServiceGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServiceGroupMemberList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceCgnv6ServiceGroupMemberListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6ServiceGroupMemberListSamplingEnable(d []interface{}) []edpt.Cgnv6ServiceGroupMemberListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServiceGroupMemberListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServiceGroupMemberListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6ServiceGroupSamplingEnable(d []interface{}) []edpt.Cgnv6ServiceGroupSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServiceGroupSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServiceGroupSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6ServiceGroup(d *schema.ResourceData) edpt.Cgnv6ServiceGroup {
	var ret edpt.Cgnv6ServiceGroup
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.MemberList = getSliceCgnv6ServiceGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6ServiceGroupSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.SharedGroup = d.Get("shared_group").(string)
	ret.Inst.SharedPartition = d.Get("shared_partition").(string)
	ret.Inst.TrafficReplicationMirrorIpRepl = d.Get("traffic_replication_mirror_ip_repl").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
