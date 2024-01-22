package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogServer() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_log_server`: Acos logging Server\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsLogServerCreate,
		UpdateContext: resourceAcosEventsLogServerUpdate,
		ReadContext:   resourceAcosEventsLogServerRead,
		DeleteContext: resourceAcosEventsLogServerDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable this Logging Server; 'disable': Disable this Logging Server;",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'msgs_sent': Number of log messages sent;",
									},
								},
							},
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
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'msgs_sent': Number of log messages sent;",
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
func resourceAcosEventsLogServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogServerCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogServer(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogServerRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsLogServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogServerUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogServer(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogServerRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsLogServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogServerDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogServer(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsLogServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogServerRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogServer(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosEventsLogServerPortList(d []interface{}) []edpt.AcosEventsLogServerPortList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsLogServerPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsLogServerPortList
		oi.PortNumber = in["port_number"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.Action = in["action"].(string)
		oi.HealthCheck = in["health_check"].(string)
		oi.HealthCheckDisable = in["health_check_disable"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceAcosEventsLogServerPortListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsLogServerPortListSamplingEnable(d []interface{}) []edpt.AcosEventsLogServerPortListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AcosEventsLogServerPortListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsLogServerPortListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAcosEventsLogServerSamplingEnable(d []interface{}) []edpt.AcosEventsLogServerSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AcosEventsLogServerSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsLogServerSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsLogServer(d *schema.ResourceData) edpt.AcosEventsLogServer {
	var ret edpt.AcosEventsLogServer
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PortList = getSliceAcosEventsLogServerPortList(d.Get("port_list").([]interface{}))
	ret.Inst.ResolveAs = d.Get("resolve_as").(string)
	ret.Inst.SamplingEnable = getSliceAcosEventsLogServerSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServerIpv6Addr = d.Get("server_ipv6_addr").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
