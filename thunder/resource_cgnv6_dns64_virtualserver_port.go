package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Dns64VirtualserverPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_dns64_virtualserver_port`: Virtual Port\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Dns64VirtualserverPortCreate,
		UpdateContext: resourceCgnv6Dns64VirtualserverPortUpdate,
		ReadContext:   resourceCgnv6Dns64VirtualserverPortRead,
		DeleteContext: resourceCgnv6Dns64VirtualserverPortDelete,

		Schema: map[string]*schema.Schema{
			"acl_id_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type: schema.TypeInt, Optional: true, Description: "ACL id VPORT",
						},
						"acl_id_src_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"acl_id_seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
					},
				},
			},
			"acl_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply an access list name (Named Access List)",
						},
						"acl_name_src_nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Policy based Source NAT (NAT Pool or Pool Group)",
						},
						"acl_name_seq_num": {
							Type: schema.TypeInt, Optional: true, Description: "Specify ACL precedence (sequence-number)",
						},
					},
				},
			},
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
			},
			"auto": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure auto NAT for the vport",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"pool": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"port_number": {
				Type: schema.TypeInt, Required: true, Description: "Port",
			},
			"precedence": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set auto NAT pool as higher precedence for source NAT",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-udp': DNS service over UDP;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_conn': Current connection; 'total_l4_conn': Total L4 connections; 'total_l7_conn': Total L7 connections; 'toatal_tcp_conn': Total TCP connections; 'total_conn': Total connections; 'total_fwd_bytes': Total forward bytes; 'total_fwd_pkts': Total forward packets; 'total_rev_bytes': Total reverse bytes; 'total_rev_pkts': Total reverse packets; 'total_dns_pkts': Total DNS packets; 'total_mf_dns_pkts': Total MF DNS packets; 'es_total_failure_actions': Total failure actions; 'compression_bytes_before': Data into compression engine; 'compression_bytes_after': Data out of compression engine; 'compression_hit': Number of requests compressed; 'compression_miss': Number of requests NOT compressed; 'compression_miss_no_client': Compression miss no client; 'compression_miss_template_exclusion': Compression miss template exclusion; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total successful requests; 'peak_conn': Peak connections; 'curr_conn_rate': Current connection rate; 'last_rsp_time': Last response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time;",
						},
					},
				},
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a Service Group to this Virtual Server (Service Group Name)",
			},
			"template_dns": {
				Type: schema.TypeString, Optional: true, Description: "DNS template (DNS template name)",
			},
			"template_policy": {
				Type: schema.TypeString, Optional: true, Description: "Policy Template (Policy template name)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6Dns64VirtualserverPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64VirtualserverPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Dns64VirtualserverPortRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Dns64VirtualserverPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64VirtualserverPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Dns64VirtualserverPortRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Dns64VirtualserverPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64VirtualserverPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Dns64VirtualserverPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64VirtualserverPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Dns64VirtualserverPortAclIdList(d []interface{}) []edpt.Cgnv6Dns64VirtualserverPortAclIdList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Dns64VirtualserverPortAclIdList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Dns64VirtualserverPortAclIdList
		oi.AclId = in["acl_id"].(int)
		oi.AclIdSrcNatPool = in["acl_id_src_nat_pool"].(string)
		oi.AclIdSeqNum = in["acl_id_seq_num"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Dns64VirtualserverPortAclNameList(d []interface{}) []edpt.Cgnv6Dns64VirtualserverPortAclNameList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Dns64VirtualserverPortAclNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Dns64VirtualserverPortAclNameList
		oi.AclName = in["acl_name"].(string)
		oi.AclNameSrcNatPool = in["acl_name_src_nat_pool"].(string)
		oi.AclNameSeqNum = in["acl_name_seq_num"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Dns64VirtualserverPortSamplingEnable(d []interface{}) []edpt.Cgnv6Dns64VirtualserverPortSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Dns64VirtualserverPortSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Dns64VirtualserverPortSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Dns64VirtualserverPort(d *schema.ResourceData) edpt.Cgnv6Dns64VirtualserverPort {
	var ret edpt.Cgnv6Dns64VirtualserverPort
	ret.Inst.AclIdList = getSliceCgnv6Dns64VirtualserverPortAclIdList(d.Get("acl_id_list").([]interface{}))
	ret.Inst.AclNameList = getSliceCgnv6Dns64VirtualserverPortAclNameList(d.Get("acl_name_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Auto = d.Get("auto").(int)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Pool = d.Get("pool").(string)
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Precedence = d.Get("precedence").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SamplingEnable = getSliceCgnv6Dns64VirtualserverPortSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.TemplateDns = d.Get("template_dns").(string)
	ret.Inst.TemplatePolicy = d.Get("template_policy").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
