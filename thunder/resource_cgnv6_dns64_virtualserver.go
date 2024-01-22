package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Dns64Virtualserver() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_dns64_virtualserver`: Create a DNS64 Virtual Server\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Dns64VirtualserverCreate,
		UpdateContext: resourceCgnv6Dns64VirtualserverUpdate,
		ReadContext:   resourceCgnv6Dns64VirtualserverRead,
		DeleteContext: resourceCgnv6Dns64VirtualserverDelete,

		Schema: map[string]*schema.Schema{
			"enable_disable_action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Virtual Server (default); 'disable': Disable Virtual Server;",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
			},
			"ip_address": {
				Type: schema.TypeString, Optional: true, Description: "IP Address",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "CGNV6 Virtual Server Name",
			},
			"netmask": {
				Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
			},
			"policy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Policy template",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type: schema.TypeInt, Required: true, Description: "Port",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns-udp': DNS service over UDP;",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
						},
						"pool": {
							Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
						},
						"auto": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure auto NAT for the vport",
						},
						"precedence": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set auto NAT pool as higher precedence for source NAT",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_conn': Current connection; 'total_l4_conn': Total L4 connections; 'total_l7_conn': Total L7 connections; 'toatal_tcp_conn': Total TCP connections; 'total_conn': Total connections; 'total_fwd_bytes': Total forward bytes; 'total_fwd_pkts': Total forward packets; 'total_rev_bytes': Total reverse bytes; 'total_rev_pkts': Total reverse packets; 'total_dns_pkts': Total DNS packets; 'total_mf_dns_pkts': Total MF DNS packets; 'es_total_failure_actions': Total failure actions; 'compression_bytes_before': Data into compression engine; 'compression_bytes_after': Data out of compression engine; 'compression_hit': Number of requests compressed; 'compression_miss': Number of requests NOT compressed; 'compression_miss_no_client': Compression miss no client; 'compression_miss_template_exclusion': Compression miss template exclusion; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total successful requests; 'peak_conn': Peak connections; 'curr_conn_rate': Current connection rate; 'last_rsp_time': Last response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time;",
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
			"template_policy": {
				Type: schema.TypeString, Optional: true, Description: "Policy template name",
			},
			"use_if_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Interface IP",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Join a vrrp group (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceCgnv6Dns64VirtualserverCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64Virtualserver(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Dns64VirtualserverRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Dns64VirtualserverUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64Virtualserver(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Dns64VirtualserverRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Dns64VirtualserverDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64Virtualserver(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Dns64VirtualserverRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Dns64VirtualserverRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Dns64Virtualserver(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6Dns64VirtualserverPortList(d []interface{}) []edpt.Cgnv6Dns64VirtualserverPortList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Dns64VirtualserverPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Dns64VirtualserverPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Action = in["action"].(string)
		oi.Pool = in["pool"].(string)
		oi.Auto = in["auto"].(int)
		oi.Precedence = in["precedence"].(int)
		oi.ServiceGroup = in["service_group"].(string)
		oi.TemplateDns = in["template_dns"].(string)
		oi.TemplatePolicy = in["template_policy"].(string)
		oi.AclIdList = getSliceCgnv6Dns64VirtualserverPortListAclIdList(in["acl_id_list"].([]interface{}))
		oi.AclNameList = getSliceCgnv6Dns64VirtualserverPortListAclNameList(in["acl_name_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceCgnv6Dns64VirtualserverPortListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Dns64VirtualserverPortListAclIdList(d []interface{}) []edpt.Cgnv6Dns64VirtualserverPortListAclIdList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Dns64VirtualserverPortListAclIdList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Dns64VirtualserverPortListAclIdList
		oi.AclId = in["acl_id"].(int)
		oi.AclIdSrcNatPool = in["acl_id_src_nat_pool"].(string)
		oi.AclIdSeqNum = in["acl_id_seq_num"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Dns64VirtualserverPortListAclNameList(d []interface{}) []edpt.Cgnv6Dns64VirtualserverPortListAclNameList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Dns64VirtualserverPortListAclNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Dns64VirtualserverPortListAclNameList
		oi.AclName = in["acl_name"].(string)
		oi.AclNameSrcNatPool = in["acl_name_src_nat_pool"].(string)
		oi.AclNameSeqNum = in["acl_name_seq_num"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6Dns64VirtualserverPortListSamplingEnable(d []interface{}) []edpt.Cgnv6Dns64VirtualserverPortListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6Dns64VirtualserverPortListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6Dns64VirtualserverPortListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6Dns64Virtualserver(d *schema.ResourceData) edpt.Cgnv6Dns64Virtualserver {
	var ret edpt.Cgnv6Dns64Virtualserver
	ret.Inst.EnableDisableAction = d.Get("enable_disable_action").(string)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Netmask = d.Get("netmask").(string)
	ret.Inst.Policy = d.Get("policy").(int)
	ret.Inst.PortList = getSliceCgnv6Dns64VirtualserverPortList(d.Get("port_list").([]interface{}))
	ret.Inst.TemplatePolicy = d.Get("template_policy").(string)
	ret.Inst.UseIfIp = d.Get("use_if_ip").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
