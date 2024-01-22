package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_server`: Server\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbServerCreate,
		UpdateContext: resourceSlbServerUpdate,
		ReadContext:   resourceSlbServerRead,
		DeleteContext: resourceSlbServerDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable this Real Server; 'disable': Disable this Real Server; 'disable-with-health-check': disable real server, but health check work;",
			},
			"alternate_server": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alternate": {
							Type: schema.TypeInt, Optional: true, Description: "Alternate Server (Alternate Server Number)",
						},
						"alternate_name": {
							Type: schema.TypeString, Optional: true, Description: "Alternate Name",
						},
					},
				},
			},
			"conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 64000000, Description: "Connection Limit",
			},
			"conn_resume": {
				Type: schema.TypeInt, Optional: true, Description: "Connection Resume (Connection Resume (min active conn before resume taking new conn))",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "ethernet interface",
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on real server",
			},
			"external_ip": {
				Type: schema.TypeString, Optional: true, Description: "External IP address for NAT of GSLB",
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
			"health_check_shared": {
				Type: schema.TypeString, Optional: true, Description: "Health Check Monitor (Health monitor name)",
			},
			"host": {
				Type: schema.TypeString, Optional: true, Description: "IP Address",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 address Mapping of GSLB",
			},
			"l2_health_check_path": {
				Type: schema.TypeString, Optional: true, Description: "L2 health check path",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server Name",
			},
			"no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
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
						"range": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Port range (Port range value - used for vip-to-rport-mapping and vport-rport range mapping)",
						},
						"template_port": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "Port template (Port template name)",
						},
						"shared_partition_port_template": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a port template from shared partition",
						},
						"template_port_shared": {
							Type: schema.TypeString, Optional: true, Description: "Port Template Name",
						},
						"template_server_ssl": {
							Type: schema.TypeString, Optional: true, Description: "Server side SSL template (Server side SSL Name)",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable; 'disable-with-health-check': disable port, but health check work;",
						},
						"no_ssl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No SSL",
						},
						"health_check": {
							Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
						},
						"shared_rport_health_check": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a health-check from shared partition",
						},
						"rport_health_check_shared": {
							Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
						},
						"health_check_follow_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify which port to follow for health status (Port Number)",
						},
						"follow_port_protocol": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"health_check_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable health check",
						},
						"support_http2": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Starting HTTP/2 with Prior Knowledge",
						},
						"only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force using HTTP/2 with Prior Knowledge all the time",
						},
						"weight": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Port Weight (Connection Weight)",
						},
						"conn_limit": {
							Type: schema.TypeInt, Optional: true, Default: 64000000, Description: "Connection Limit",
						},
						"no_logging": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
						},
						"conn_resume": {
							Type: schema.TypeInt, Optional: true, Description: "Connection Resume",
						},
						"stats_data_action": {
							Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for real server port; 'stats-data-disable': Disable statistical data collection for real server port;",
						},
						"extended_stats": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on real server port",
						},
						"alternate_port": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alternate": {
										Type: schema.TypeInt, Optional: true, Description: "Alternate Server Number",
									},
									"alternate_name": {
										Type: schema.TypeString, Optional: true, Description: "Alternate Name",
									},
									"alternate_server_port": {
										Type: schema.TypeInt, Optional: true, Description: "Port (Alternate Server Port Value)",
									},
								},
							},
						},
						"auth_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_principal_name": {
										Type: schema.TypeString, Optional: true, Description: "Service Principal Name (Kerberos principal name)",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_req': Current requests; 'total_req': Total Requests; 'total_req_succ': Total requests succ; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total connections; 'last_total_conn': Last total connections; 'peak_conn': Peak connections; 'es_resp_200': Response status 200; 'es_resp_300': Response status 300; 'es_resp_400': Response status 400; 'es_resp_500': Response status 500; 'es_resp_other': Response status other; 'es_req_count': Total proxy requests; 'es_resp_count': Total proxy response; 'es_resp_invalid_http': Total non-http response; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_good_status_code': Total reverse packets with good status code inspected; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'resp-count': Total Response Count; 'resp-1xx': Response status 1xx; 'resp-2xx': Response status 2xx; 'resp-3xx': Response status 3xx; 'resp-4xx': Response status 4xx; 'resp-5xx': Response status 5xx; 'resp-other': Response status Other; 'resp-latency': Time to First Response Byte; 'curr_pconn': Current persistent connections;",
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
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'total-conn': Total established connections; 'fwd-pkt': Forward Packets Processed; 'rev-pkt': Reverse Packets Processed; 'peak-conn': Peak number of established connections; 'total_req': Total Requests processed; 'total_req_succ': Total Requests succeeded; 'curr_ssl_conn': Current SSL connections established; 'total_ssl_conn': Total SSL connections established; 'total_fwd_bytes': Bytes processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_pkts': Packets processed in reverse direction; 'ip_only_lb_fwd_bytes': IP-Only-LB Bytes processed in forward direction; 'ip_only_lb_rev_bytes': IP-Only-LB Bytes processed in reverse direction; 'ip_only_lb_fwd_pkts': IP-Only-LB Packets processed in forward direction; 'ip_only_lb_rev_pkts': IP-Only-LB Packets processed in reverse direction;",
						},
					},
				},
			},
			"server_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address",
			},
			"shared_partition_health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a health-check from shared partition",
			},
			"shared_partition_server_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a server template from shared partition",
			},
			"slow_start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Slowly ramp up the connection number after server is up (start from 128, then double every 10 sec till 4096)",
			},
			"spoofing_cache": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "This server is a spoofing cache",
			},
			"stats_data_action": {
				Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for real server; 'stats-data-disable': Disable statistical data collection for real server;",
			},
			"template_link_cost": {
				Type: schema.TypeString, Optional: true, Description: "Link-Cost template (Link-Cost template name)",
			},
			"template_server": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "Server template (Server template name)",
			},
			"template_server_shared": {
				Type: schema.TypeString, Optional: true, Description: "Server Template Name",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "trunk interface",
			},
			"use_aam_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Using aam server. For health check, please configure it in aam server",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Weight for this Real Server (Connection Weight)",
			},
		},
	}
}
func resourceSlbServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServerRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServerRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbServerAlternateServer(d []interface{}) []edpt.SlbServerAlternateServer {

	count1 := len(d)
	ret := make([]edpt.SlbServerAlternateServer, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerAlternateServer
		oi.Alternate = in["alternate"].(int)
		oi.AlternateName = in["alternate_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerPortList(d []interface{}) []edpt.SlbServerPortList {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Range = in["range"].(int)
		oi.TemplatePort = in["template_port"].(string)
		oi.SharedPartitionPortTemplate = in["shared_partition_port_template"].(int)
		oi.TemplatePortShared = in["template_port_shared"].(string)
		oi.TemplateServerSsl = in["template_server_ssl"].(string)
		oi.Action = in["action"].(string)
		oi.NoSsl = in["no_ssl"].(int)
		oi.HealthCheck = in["health_check"].(string)
		oi.SharedRportHealthCheck = in["shared_rport_health_check"].(int)
		oi.RportHealthCheckShared = in["rport_health_check_shared"].(string)
		oi.HealthCheckFollowPort = in["health_check_follow_port"].(int)
		oi.FollowPortProtocol = in["follow_port_protocol"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		oi.SupportHttp2 = in["support_http2"].(int)
		oi.Only = in["only"].(int)
		oi.Weight = in["weight"].(int)
		oi.ConnLimit = in["conn_limit"].(int)
		oi.NoLogging = in["no_logging"].(int)
		oi.ConnResume = in["conn_resume"].(int)
		oi.StatsDataAction = in["stats_data_action"].(string)
		oi.ExtendedStats = in["extended_stats"].(int)
		oi.AlternatePort = getSliceSlbServerPortListAlternatePort(in["alternate_port"].([]interface{}))
		oi.AuthCfg = getObjectSlbServerPortListAuthCfg(in["auth_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbServerPortListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.PacketCaptureTemplate = in["packet_capture_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerPortListAlternatePort(d []interface{}) []edpt.SlbServerPortListAlternatePort {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortListAlternatePort, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortListAlternatePort
		oi.Alternate = in["alternate"].(int)
		oi.AlternateName = in["alternate_name"].(string)
		oi.AlternateServerPort = in["alternate_server_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServerPortListAuthCfg(d []interface{}) edpt.SlbServerPortListAuthCfg {

	count1 := len(d)
	var ret edpt.SlbServerPortListAuthCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServicePrincipalName = in["service_principal_name"].(string)
	}
	return ret
}

func getSliceSlbServerPortListSamplingEnable(d []interface{}) []edpt.SlbServerPortListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbServerPortListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServerSamplingEnable(d []interface{}) []edpt.SlbServerSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbServerSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServer(d *schema.ResourceData) edpt.SlbServer {
	var ret edpt.SlbServer
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AlternateServer = getSliceSlbServerAlternateServer(d.Get("alternate_server").([]interface{}))
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnResume = d.Get("conn_resume").(int)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.ExternalIp = d.Get("external_ip").(string)
	ret.Inst.FqdnName = d.Get("fqdn_name").(string)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckShared = d.Get("health_check_shared").(string)
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	ret.Inst.L2HealthCheckPath = d.Get("l2_health_check_path").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NoLogging = d.Get("no_logging").(int)
	ret.Inst.PortList = getSliceSlbServerPortList(d.Get("port_list").([]interface{}))
	ret.Inst.ResolveAs = d.Get("resolve_as").(string)
	ret.Inst.SamplingEnable = getSliceSlbServerSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	ret.Inst.SharedPartitionHealthCheck = d.Get("shared_partition_health_check").(int)
	ret.Inst.SharedPartitionServerTemplate = d.Get("shared_partition_server_template").(int)
	ret.Inst.SlowStart = d.Get("slow_start").(int)
	ret.Inst.SpoofingCache = d.Get("spoofing_cache").(int)
	ret.Inst.StatsDataAction = d.Get("stats_data_action").(string)
	ret.Inst.TemplateLinkCost = d.Get("template_link_cost").(string)
	ret.Inst.TemplateServer = d.Get("template_server").(string)
	ret.Inst.TemplateServerShared = d.Get("template_server_shared").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	ret.Inst.UseAamServer = d.Get("use_aam_server").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	return ret
}
