package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServiceGroupMember() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_service_group_member`: Service Group Member\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbServiceGroupMemberCreate,
		UpdateContext: resourceSlbServiceGroupMemberUpdate,
		ReadContext:   resourceSlbServiceGroupMemberRead,
		DeleteContext: resourceSlbServiceGroupMemberDelete,

		Schema: map[string]*schema.Schema{
			"fqdn_name": {
				Type: schema.TypeString, Optional: true, Description: "Server hostname - Not applicable if real server is already defined",
			},
			"host": {
				Type: schema.TypeString, Optional: true, Description: "IP Address - Not applicable if real server is already defined",
			},
			"member_priority": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Priority of Port in the Group (Priority of Port in the Group, default is 1)",
			},
			"member_state": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable member service port; 'disable': Disable member service port; 'disable-with-health-check': disable member service port, but health check work;",
			},
			"member_stats_data_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable statistical data collection",
			},
			"member_template": {
				Type: schema.TypeString, Optional: true, Description: "Real server port template (Real server port template name)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Member name",
			},
			"service_group_name": {
				Type: schema.TypeString, Required: true, Description: "service_group_name name",
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port number",
			},
			"resolve_as": {
				Type: schema.TypeString, Optional: true, Default: "resolve-to-ipv4", Description: "'resolve-to-ipv4': Use A Query only to resolve FQDN; 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total established connections; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_status_code_2xx': Total reverse packets inspected status code 2xx; 'total_rev_pkts_inspected_status_code_non_5xx': Total reverse packets inspected status code non 5xx; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total requests successful; 'peak_conn': Peak connections; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'curr_conn_overflow': Current connection counter overflow count; 'state_flaps': State flaps count;",
						},
					},
				},
			},
			"server_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 Address - Not applicable if real server is already defined",
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
func resourceSlbServiceGroupMemberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupMemberCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroupMember(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServiceGroupMemberRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbServiceGroupMemberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupMemberUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroupMember(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServiceGroupMemberRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbServiceGroupMemberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupMemberDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroupMember(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbServiceGroupMemberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupMemberRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroupMember(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbServiceGroupMemberSamplingEnable(d []interface{}) []edpt.SlbServiceGroupMemberSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupMemberSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupMemberSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServiceGroupMember(d *schema.ResourceData) edpt.SlbServiceGroupMember {
	var ret edpt.SlbServiceGroupMember
	ret.Inst.FqdnName = d.Get("fqdn_name").(string)
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.MemberPriority = d.Get("member_priority").(int)
	ret.Inst.MemberState = d.Get("member_state").(string)
	ret.Inst.MemberStatsDataDisable = d.Get("member_stats_data_disable").(int)
	ret.Inst.MemberTemplate = d.Get("member_template").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ServiceGroupName = d.Get("service_group_name").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.ResolveAs = d.Get("resolve_as").(string)
	ret.Inst.SamplingEnable = getSliceSlbServiceGroupMemberSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
