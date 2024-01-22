package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_dns`: DNS Packet Statistics\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemDnsCreate,
		UpdateContext: resourceSystemDnsUpdate,
		ReadContext:   resourceSystemDnsRead,
		DeleteContext: resourceSystemDnsDelete,

		Schema: map[string]*schema.Schema{
			"recursive_nameserver": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"follow_shared": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the configured name servers of shared partition",
						},
						"server_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv4 server address",
									},
									"v4_desc": {
										Type: schema.TypeString, Optional: true, Description: "Description for this ipv4 address",
									},
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "Specify IPv6 server address",
									},
									"v6_desc": {
										Type: schema.TypeString, Optional: true, Description: "Description for this ipv6 address",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'slb_req': No. of requests; 'slb_resp': No. of responses; 'slb_no_resp': No. of requests with no response; 'slb_req_rexmit': No. of requests retransmit; 'slb_resp_no_match': No. of requests and responses with no match; 'slb_no_resource': No. of resource failures; 'nat_req': (NAT) No. of requests; 'nat_resp': (NAT) No. of responses; 'nat_no_resp': (NAT) No. of resource failures; 'nat_req_rexmit': (NAT) No. of request retransmits; 'nat_resp_no_match': (NAT) No. of requests with no response; 'nat_no_resource': (NAT) No. of resource failures; 'nat_xid_reused': (NAT) No. of requests reusing a transaction id; 'filter_type_drop': Total Query Type Drop; 'filter_class_drop': Total Query Class Drop; 'filter_type_any_drop': Total Query ANY Type Drop; 'slb_dns_client_ssl_succ': No. of client ssl success; 'slb_dns_server_ssl_succ': No. of server ssl success; 'slb_dns_udp_conn': No. of backend udp connections; 'slb_dns_udp_conn_succ': No. of backend udp conn established; 'slb_dns_padding_to_server_removed': slb_dns_padding_to_server_removed; 'slb_dns_padding_to_client_added': slb_dns_padding_to_client_added; 'slb_dns_edns_subnet_to_server_removed': slb_dns_edns_subnet_to_server_removed; 'slb_dns_udp_retransmit': slb_dns_udp_retransmit; 'slb_dns_udp_retransmit_fail': slb_dns_udp_retransmit_fail; 'rpz_action_drop': RPZ Action Drop; 'rpz_action_pass_thru': RPZ Action Pass Through; 'rpz_action_tcp_only': RPZ Action TCP Only; 'rpz_action_nxdomain': RPZ Action NXDOMAIN; 'rpz_action_nodata': RPZ Action NODATA; 'rpz_action_local_data': RPZ Action Local Data; 'slb_drop': DNS requests drop; 'nat_slb_drop': (NAT)DNS requests drop; 'invalid_q_len_to_udp': invalid query length to conver to UDP; 'slb_dns_edns_ecs_received': Number of ecs from client received; 'slb_dns_edns_ecs_inserted': Number of ecs inserted;",
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
func resourceSystemDnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDns(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDnsRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemDnsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDns(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDnsRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemDnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDns(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemDnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDns(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemDnsRecursiveNameserver1571(d []interface{}) edpt.SystemDnsRecursiveNameserver1571 {

	count1 := len(d)
	var ret edpt.SystemDnsRecursiveNameserver1571
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		ret.ServerList = getSliceSystemDnsRecursiveNameserverServerList1572(in["server_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSystemDnsRecursiveNameserverServerList1572(d []interface{}) []edpt.SystemDnsRecursiveNameserverServerList1572 {

	count1 := len(d)
	ret := make([]edpt.SystemDnsRecursiveNameserverServerList1572, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsRecursiveNameserverServerList1572
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.V4Desc = in["v4_desc"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.V6Desc = in["v6_desc"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemDnsSamplingEnable(d []interface{}) []edpt.SystemDnsSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SystemDnsSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemDns(d *schema.ResourceData) edpt.SystemDns {
	var ret edpt.SystemDns
	ret.Inst.RecursiveNameserver = getObjectSystemDnsRecursiveNameserver1571(d.Get("recursive_nameserver").([]interface{}))
	ret.Inst.SamplingEnable = getSliceSystemDnsSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
