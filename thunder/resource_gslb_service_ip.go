package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbServiceIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_service_ip`: Service IP\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbServiceIpCreate,
		UpdateContext: resourceGslbServiceIpUpdate,
		ReadContext:   resourceGslbServiceIpRead,
		DeleteContext: resourceGslbServiceIpDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable this GSLB server; 'disable': Disable this GSLB server;",
			},
			"external_ip": {
				Type: schema.TypeString, Optional: true, Description: "External IP address for NAT",
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check Monitor (Monitor Name)",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Health Check Monitor",
			},
			"health_check_protocol_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable GSLB Protocol Health Monitor",
			},
			"ip_address": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"ipv6": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 address Mapping (Applicable only when service-ip has an IPv4 Address)",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address",
			},
			"node_name": {
				Type: schema.TypeString, Required: true, Description: "Service-IP Name",
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"port_proto": {
							Type: schema.TypeString, Required: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable this GSLB server port; 'disable': Disable this GSLB server port;",
						},
						"health_check": {
							Type: schema.TypeString, Optional: true, Description: "Health Check Monitor (Monitor Name)",
						},
						"health_check_follow_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify which port to follow for health status (Port Number)",
						},
						"follow_port_protocol": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
						},
						"health_check_protocol_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable GSLB Protocol Health Monitor",
						},
						"health_check_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable Health Check Monitor",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'active': Active Servers; 'current': Current Connections;",
									},
								},
							},
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hits': Number of times the service IP has been selected; 'recent': Recent hits;",
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
func resourceGslbServiceIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceIpRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbServiceIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbServiceIpRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbServiceIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbServiceIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbServiceIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbServiceIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbServiceIpPortList(d []interface{}) []edpt.GslbServiceIpPortList {

	count1 := len(d)
	ret := make([]edpt.GslbServiceIpPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceIpPortList
		oi.PortNum = in["port_num"].(int)
		oi.PortProto = in["port_proto"].(string)
		oi.Action = in["action"].(string)
		oi.HealthCheck = in["health_check"].(string)
		oi.HealthCheckFollowPort = in["health_check_follow_port"].(int)
		oi.FollowPortProtocol = in["follow_port_protocol"].(string)
		oi.HealthCheckProtocolDisable = in["health_check_protocol_disable"].(int)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceGslbServiceIpPortListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbServiceIpPortListSamplingEnable(d []interface{}) []edpt.GslbServiceIpPortListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbServiceIpPortListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceIpPortListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbServiceIpSamplingEnable(d []interface{}) []edpt.GslbServiceIpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.GslbServiceIpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbServiceIpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbServiceIp(d *schema.ResourceData) edpt.GslbServiceIp {
	var ret edpt.GslbServiceIp
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ExternalIp = d.Get("external_ip").(string)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.HealthCheckProtocolDisable = d.Get("health_check_protocol_disable").(int)
	ret.Inst.IpAddress = d.Get("ip_address").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(string)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.NodeName = d.Get("node_name").(string)
	ret.Inst.PortList = getSliceGslbServiceIpPortList(d.Get("port_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceGslbServiceIpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
