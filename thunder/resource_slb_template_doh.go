package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDoh() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_doh`: DNS over HTTP(s) template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDohCreate,
		UpdateContext: resourceSlbTemplateDohUpdate,
		ReadContext:   resourceSlbTemplateDohRead,
		DeleteContext: resourceSlbTemplateDohDelete,

		Schema: map[string]*schema.Schema{
			"conn_reuse": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable Connection Reuse; 'disable': Disable Connection-Reuse (Default);",
			},
			"dns": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "DNS Template Name",
			},
			"dns_retry": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retry_interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "DNS Retry Interval value 1 - 400 in units of 100ms, default is 10 (default is 1000ms) (1 - 400 in units of 100ms, default is 10 (1000ms/1sec))",
						},
						"after_timeout": {
							Type: schema.TypeString, Optional: true, Default: "close", Description: "'close': Close client side connection; 'retry-with-tcp': Retry DNS query to server using TCP (If UDP was tried initially. Close after.);",
						},
						"max_trials": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Total number of times to try DNS query to server before closing client connection, default 3",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"forwarder": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forwarding_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "SLB VIP IPv4 address to forward DOH query (IP address)",
						},
						"v4_internal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Try to find this IP as a VIP in this L3v Partition and forward it internally to the VIP",
						},
						"v4_port": {
							Type: schema.TypeInt, Optional: true, Default: 53, Description: "Forwarding port number, Default is 53",
						},
						"v4_l4_proto": {
							Type: schema.TypeString, Optional: true, Default: "both", Description: "'tcp': Use TCP only when forwarding DNS traffic; 'udp': Use UDP only when forwarding DNS traffic; 'both': Use UDP 1st and if unreachable, retry with TCP when forwarding DNS traffic;",
						},
						"forwarding_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "SLB VIP IPv6 address to forward DOH query (IP address)",
						},
						"v6_internal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Try to find this IP as a VIP in this L3v Partition and forward it internally to the VIP",
						},
						"v6_port": {
							Type: schema.TypeInt, Optional: true, Default: 53, Description: "Forwarding port number, Default is 53",
						},
						"v6_l4_proto": {
							Type: schema.TypeString, Optional: true, Default: "both", Description: "'tcp': Use TCP only when forwarding DNS traffic; 'udp': Use UDP only when forwarding DNS traffic; 'both': Use UDP 1st and if unreachable, retry with TCP when forwarding DNS traffic;",
						},
						"tcp_service_group": {
							Type: schema.TypeString, Optional: true, Description: "Bind a TCP Service Group to the template (Service Group Name)",
						},
						"udp_service_group": {
							Type: schema.TypeString, Optional: true, Description: "Bind a UDP Service Group to the template (Service Group Name)",
						},
						"bypass_doh": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward valid DoH HTTP request as is, no DNS packet extraction (Bypass DoH)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS over HTTP(s) Template Name",
			},
			"non_dns_request": {
				Type: schema.TypeString, Optional: true, Default: "reject", Description: "'allow': Forward Non-DoH request to http server bound to vport; 'reject': Reject Non-DoH requests with HTTP 400 Bad Request (Default);",
			},
			"reject_status_code": {
				Type: schema.TypeString, Optional: true, Default: "400", Description: "'400': Status Code 400 BAD Request (Default); '500': Status Code 500 Internal Server Error; '501': Status Code 501 Not Implemented;",
			},
			"shared_partition_dns_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a DNS template from shared partition",
			},
			"shared_partition_tcp_proxy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a TCP Proxy template from shared partition",
			},
			"snat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Source NAT pool or pool group",
			},
			"source_nat": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'auto': Perform Source NAT Auto for service-group(Default) (Not supported with forwarding-ip); 'disable': Don't perform source-nat for server side DNS queries; 'pool': Perform Source NAT with specific pool;",
			},
			"tcp_proxy": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "TCP Proxy Template Name",
			},
			"template_dns_shared": {
				Type: schema.TypeString, Optional: true, Description: "DNS Template name",
			},
			"template_tcp_proxy_shared": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template name",
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
func resourceSlbTemplateDohCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDoh(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDohRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDohUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDoh(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDohRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDohDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDoh(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDohRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDoh(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateDohDnsRetry1445(d []interface{}) edpt.SlbTemplateDohDnsRetry1445 {

	count1 := len(d)
	var ret edpt.SlbTemplateDohDnsRetry1445
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RetryInterval = in["retry_interval"].(int)
		ret.AfterTimeout = in["after_timeout"].(string)
		ret.MaxTrials = in["max_trials"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSlbTemplateDohForwarder1446(d []interface{}) edpt.SlbTemplateDohForwarder1446 {

	count1 := len(d)
	var ret edpt.SlbTemplateDohForwarder1446
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ForwardingIpv4 = in["forwarding_ipv4"].(string)
		ret.V4Internal = in["v4_internal"].(int)
		ret.V4Port = in["v4_port"].(int)
		ret.V4L4Proto = in["v4_l4_proto"].(string)
		ret.ForwardingIpv6 = in["forwarding_ipv6"].(string)
		ret.V6Internal = in["v6_internal"].(int)
		ret.V6Port = in["v6_port"].(int)
		ret.V6L4Proto = in["v6_l4_proto"].(string)
		ret.TcpServiceGroup = in["tcp_service_group"].(string)
		ret.UdpServiceGroup = in["udp_service_group"].(string)
		ret.BypassDoh = in["bypass_doh"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSlbTemplateDoh(d *schema.ResourceData) edpt.SlbTemplateDoh {
	var ret edpt.SlbTemplateDoh
	ret.Inst.ConnReuse = d.Get("conn_reuse").(string)
	ret.Inst.Dns = d.Get("dns").(string)
	ret.Inst.DnsRetry = getObjectSlbTemplateDohDnsRetry1445(d.Get("dns_retry").([]interface{}))
	ret.Inst.Forwarder = getObjectSlbTemplateDohForwarder1446(d.Get("forwarder").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NonDnsRequest = d.Get("non_dns_request").(string)
	ret.Inst.RejectStatusCode = d.Get("reject_status_code").(string)
	ret.Inst.SharedPartitionDnsTemplate = d.Get("shared_partition_dns_template").(int)
	ret.Inst.SharedPartitionTcpProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	ret.Inst.SnatPool = d.Get("snat_pool").(string)
	ret.Inst.SourceNat = d.Get("source_nat").(string)
	ret.Inst.TcpProxy = d.Get("tcp_proxy").(string)
	ret.Inst.TemplateDnsShared = d.Get("template_dns_shared").(string)
	ret.Inst.TemplateTcpProxyShared = d.Get("template_tcp_proxy_shared").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
