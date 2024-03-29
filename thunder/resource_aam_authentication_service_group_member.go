package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationServiceGroupMember() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_service_group_member`: Authentication service group member\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationServiceGroupMemberCreate,
		UpdateContext: resourceAamAuthenticationServiceGroupMemberUpdate,
		ReadContext:   resourceAamAuthenticationServiceGroupMemberRead,
		DeleteContext: resourceAamAuthenticationServiceGroupMemberDelete,

		Schema: map[string]*schema.Schema{
			"member_priority": {
				Type: schema.TypeInt, Optional: true, Description: "Priority of Port in the Group",
			},
			"member_state": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable member service port; 'disable': Disable member service port;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Member name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port number",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationServiceGroupMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupMemberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroupMember(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServiceGroupMemberRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationServiceGroupMemberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupMemberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroupMember(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationServiceGroupMemberRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationServiceGroupMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupMemberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroupMember(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationServiceGroupMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationServiceGroupMemberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationServiceGroupMember(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationServiceGroupMemberSamplingEnable(d []interface{}) []edpt.AamAuthenticationServiceGroupMemberSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationServiceGroupMemberSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationServiceGroupMemberSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationServiceGroupMember(d *schema.ResourceData) edpt.AamAuthenticationServiceGroupMember {
	var ret edpt.AamAuthenticationServiceGroupMember
	ret.Inst.MemberPriority = d.Get("member_priority").(int)
	ret.Inst.MemberState = d.Get("member_state").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.SamplingEnable = getSliceAamAuthenticationServiceGroupMemberSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
