package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_server`: Firewall logging Server\n\n__PLACEHOLDER__",
		CreateContext: resourceFwServerCreate,
		UpdateContext: resourceFwServerUpdate,
		ReadContext:   resourceFwServerRead,
		DeleteContext: resourceFwServerDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable this Real Server; 'disable': Disable this Real Server;",
			},
			"fqdn_name": {
				Type: schema.TypeString, Optional: true, Description: "Server hostname",
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check Monitor (Health monitor name)",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable configured health check configuration",
			},
			"host": {
				Type: schema.TypeString, Optional: true, Description: "IP Address",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Name",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_number": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
						},
						"health_check": {
							Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
						},
						"health_check_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable health check",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_conn': Current connections; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total request success; 'total_fwd_bytes': Forward bytes; 'total_fwd_pkts': Forward packets; 'total_rev_bytes': Reverse bytes; 'total_rev_pkts': Reverse packets; 'total_conn': Total connections; 'last_total_conn': Last total connections; 'peak_conn': Peak connections; 'es_resp_200': Response status 200; 'es_resp_300': Response status 300; 'es_resp_400': Response status 400; 'es_resp_500': Response status 500; 'es_resp_other': Response status other; 'es_req_count': Total proxy request; 'es_resp_count': Total proxy Response; 'es_resp_invalid_http': Total non-http response; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_good_status_code': Total reverse packets with good status code inspected; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time;",
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
			"resolve_as": {
				Type: schema.TypeString, Optional: true, Default: "resolve-to-ipv4", Description: "'resolve-to-ipv4': Use A Query only to resolve FQDN; 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr-conn': Current connections; 'total-conn': Total connections; 'fwd-pkt': Forward packets; 'rev-pkt': Reverse Packets; 'peak-conn': Peak connections;",
						},
					},
				},
			},
			"server_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address",
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
func resourceFwServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwServerRead(ctx, d, meta)
	}
	return diags
}

func resourceFwServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwServerRead(ctx, d, meta)
	}
	return diags
}
func resourceFwServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwServerPortList(d []interface{}) []edpt.FwServerPortList {

	count1 := len(d)
	ret := make([]edpt.FwServerPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwServerPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Action = in["action"].(string)
		oi.HealthCheck = in["health_check"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceFwServerPortListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwServerPortListSamplingEnable(d []interface{}) []edpt.FwServerPortListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwServerPortListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwServerPortListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwServerSamplingEnable(d []interface{}) []edpt.FwServerSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwServerSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwServerSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwServer(d *schema.ResourceData) edpt.FwServer {
	var ret edpt.FwServer
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.FqdnName = d.Get("fqdn_name").(string)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PortList = getSliceFwServerPortList(d.Get("port_list").([]interface{}))
	ret.Inst.ResolveAs = d.Get("resolve_as").(string)
	ret.Inst.SamplingEnable = getSliceFwServerSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
