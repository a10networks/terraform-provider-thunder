package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSlbServerPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_server_port`: Real Server Port\n\n",
		CreateContext: resourceSlbServerPortCreate,
		UpdateContext: resourceSlbServerPortUpdate,
		ReadContext:   resourceSlbServerPortRead,
		DeleteContext: resourceSlbServerPortDelete,
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable; 'disable-with-health-check': disable port, but health check work;",
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable", "disable-with-health-check"}, false),
			},
			"alternate_port": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alternate": {
							Type: schema.TypeInt, Optional: true, Description: "Alternate Server Number",
							ValidateFunc: validation.IntBetween(1, 16),
						},
						"alternate_name": {
							Type: schema.TypeString, Optional: true, Description: "Alternate Name",
							ValidateFunc: validation.StringLenBetween(1, 127),
						},
						"alternate_server_port": {
							Type: schema.TypeInt, Optional: true, Description: "Port (Alternate Server Port Value)",
							ValidateFunc: validation.IntBetween(1, 65535),
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
							ValidateFunc: validation.StringLenBetween(1, 128),
						},
					},
				},
			},
			"conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 64000000, Description: "Connection Limit",
				ValidateFunc: validation.IntBetween(1, 64000000),
			},
			"conn_resume": {
				Type: schema.TypeInt, Optional: true, Description: "Connection Resume",
				ValidateFunc: validation.IntBetween(1, 1000000),
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on real server port",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"follow_port_protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
				ValidateFunc:  validation.StringLenBetween(1, 63),
				ConflictsWith: []string{"shared_rport_health_check", "health_check_follow_port", "health_check_disable"},
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable health check",
				ValidateFunc:  validation.IntBetween(0, 1),
				ConflictsWith: []string{"health_check", "shared_rport_health_check", "rport_health_check_shared", "health_check_follow_port"},
			},
			"health_check_follow_port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify which port to follow for health status (Port Number)",
				ValidateFunc:  validation.IntBetween(1, 65534),
				ConflictsWith: []string{"health_check", "health_check_disable"},
			},
			"no_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not log connection over limit event",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"no_ssl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "No SSL",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
				ValidateFunc: validation.StringLenBetween(1, 128),
			},
			"port_number": {
				Type: schema.TypeInt, Required: true, ForceNew: true, Description: "Port Number",
				ValidateFunc: validation.IntBetween(0, 65534),
			},
			"protocol": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
				ValidateFunc: validation.StringInSlice([]string{"tcp", "udp"}, false),
			},
			"range": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Port range (Port range value - used for vip-to-rport-mapping and vport-rport range mapping)",
				ValidateFunc: validation.IntBetween(0, 254),
			},
			"rport_health_check_shared": {
				Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
				ValidateFunc:  validation.StringLenBetween(1, 63),
				ConflictsWith: []string{"health_check_disable"},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_req': Current requests; 'total_req': Total Requests; 'total_req_succ': Total requests succ; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total connections; 'last_total_conn': Last total connections; 'peak_conn': Peak connections; 'es_resp_200': Response status 200; 'es_resp_300': Response status 300; 'es_resp_400': Response status 400; 'es_resp_500': Response status 500; 'es_resp_other': Response status other; 'es_req_count': Total proxy requests; 'es_resp_count': Total proxy response; 'es_resp_invalid_http': Total non-http response; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_good_status_code': Total reverse packets with good status code inspected; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'resp-count': Total Response Count; 'resp-1xx': Response status 1xx; 'resp-2xx': Response status 2xx; 'resp-3xx': Response status 3xx; 'resp-4xx': Response status 4xx; 'resp-5xx': Response status 5xx; 'resp-other': Response status Other; 'resp-latency': Time to First Response Byte; 'curr_pconn': Current persistent connections;",
							ValidateFunc: validation.StringInSlice([]string{"all", "curr_req", "total_req", "total_req_succ", "total_fwd_bytes", "total_fwd_pkts", "total_rev_bytes", "total_rev_pkts", "total_conn", "last_total_conn", "peak_conn", "es_resp_200", "es_resp_300", "es_resp_400", "es_resp_500", "es_resp_other", "es_req_count", "es_resp_count", "es_resp_invalid_http", "total_rev_pkts_inspected", "total_rev_pkts_inspected_good_status_code", "response_time", "fastest_rsp_time", "slowest_rsp_time", "curr_ssl_conn", "total_ssl_conn", "resp-count", "resp-1xx", "resp-2xx", "resp-3xx", "resp-4xx", "resp-5xx", "resp-other", "resp-latency", "curr_pconn"}, false),
						},
					},
				},
			},
			"server_name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Server Name",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"shared_partition_port_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a port template from shared partition",
				ValidateFunc:  validation.IntBetween(0, 1),
				ConflictsWith: []string{"template_port"},
			},
			"shared_rport_health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a health-check from shared partition",
				ValidateFunc:  validation.IntBetween(0, 1),
				ConflictsWith: []string{"health_check", "health_check_disable"},
			},
			"stats_data_action": {
				Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for real server port; 'stats-data-disable': Disable statistical data collection for real server port;",
				ValidateFunc: validation.StringInSlice([]string{"stats-data-enable", "stats-data-disable"}, false),
			},
			"support_http2": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Starting HTTP/2 with Prior Knowledge",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"template_port": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "Port template (Port template name)",
				ValidateFunc:  validation.StringLenBetween(1, 127),
				ConflictsWith: []string{"shared_partition_port_template"},
			},
			"template_port_shared": {
				Type: schema.TypeString, Optional: true, Description: "Port Template Name",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"template_server_ssl": {
				Type: schema.TypeString, Optional: true, Description: "Server side SSL template (Server side SSL Name)",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"weight": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Port Weight (Connection Weight)",
				ValidateFunc: validation.IntBetween(1, 1000),
			},
		},
	}
}

func resourceSlbServerPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServerPortRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbServerPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbServerPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServerPortRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbServerPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServerPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServerPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbServerPortAlternatePort(d []interface{}) []edpt.SlbServerPortAlternatePort {
	count := len(d)
	ret := make([]edpt.SlbServerPortAlternatePort, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortAlternatePort
		oi.Alternate = in["alternate"].(int)
		oi.AlternateName = in["alternate_name"].(string)
		oi.AlternateServerPort = in["alternate_server_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServerPortAuthCfg(d []interface{}) edpt.SlbServerPortAuthCfg {
	count := len(d)
	var ret edpt.SlbServerPortAuthCfg
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.ServicePrincipalName = in["service_principal_name"].(string)
	}
	return ret
}

func getSliceSlbServerPortSamplingEnable(d []interface{}) []edpt.SlbServerPortSamplingEnable {
	count := len(d)
	ret := make([]edpt.SlbServerPortSamplingEnable, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServerPortSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServerPort(d *schema.ResourceData) edpt.SlbServerPort {
	var ret edpt.SlbServerPort
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AlternatePort = getSliceSlbServerPortAlternatePort(d.Get("alternate_port").([]interface{}))
	ret.Inst.AuthCfg = getObjectSlbServerPortAuthCfg(d.Get("auth_cfg").([]interface{}))
	ret.Inst.ConnLimit = d.Get("conn_limit").(int)
	ret.Inst.ConnResume = d.Get("conn_resume").(int)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.FollowPortProtocol = d.Get("follow_port_protocol").(string)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckFollowPort = d.Get("health_check_follow_port").(int)
	ret.Inst.NoLogging = d.Get("no_logging").(int)
	ret.Inst.NoSsl = d.Get("no_ssl").(int)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.PortNumber = d.Get("port_number").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Range = d.Get("range").(int)
	ret.Inst.RportHealthCheckShared = d.Get("rport_health_check_shared").(string)
	ret.Inst.SamplingEnable = getSliceSlbServerPortSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerName = d.Get("server_name").(string)
	ret.Inst.SharedPartitionPortTemplate = d.Get("shared_partition_port_template").(int)
	ret.Inst.SharedRportHealthCheck = d.Get("shared_rport_health_check").(int)
	ret.Inst.StatsDataAction = d.Get("stats_data_action").(string)
	ret.Inst.SupportHttp2 = d.Get("support_http2").(int)
	ret.Inst.TemplatePort = d.Get("template_port").(string)
	ret.Inst.TemplatePortShared = d.Get("template_port_shared").(string)
	ret.Inst.TemplateServerSsl = d.Get("template_server_ssl").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Weight = d.Get("weight").(int)
	return ret
}
