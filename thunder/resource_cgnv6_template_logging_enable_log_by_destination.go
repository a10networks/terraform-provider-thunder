package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateLoggingEnableLogByDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_logging_enable_log_by_destination`: Enable CGNAT logging by the destination IP address, protocol, and port\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateLoggingEnableLogByDestinationCreate,
		UpdateContext: resourceCgnv6TemplateLoggingEnableLogByDestinationUpdate,
		ReadContext:   resourceCgnv6TemplateLoggingEnableLogByDestinationRead,
		DeleteContext: resourceCgnv6TemplateLoggingEnableLogByDestinationDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the ICMP traffic",
			},
			"ip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_addr": {
							Type: schema.TypeString, Required: true, Description: "Configure an IP subnet",
						},
						"tcp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
									},
									"tcp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Port Range End",
									},
								},
							},
						},
						"udp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
									},
									"udp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Port Range End",
									},
								},
							},
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the ICMP traffic",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the other layer-4 protocols",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"ip6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type: schema.TypeString, Required: true, Description: "Configure an IPv6 subnet",
						},
						"tcp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
									},
									"tcp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Port Range End",
									},
								},
							},
						},
						"udp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
									},
									"udp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Port Range End",
									},
								},
							},
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the ICMP traffic",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the other layer-4 protocols",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"others": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for the other layer-4 protocols",
			},
			"tcp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
						},
						"tcp_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Port Range End",
						},
					},
				},
			},
			"udp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
						},
						"udp_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Port Range End",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"logging_name": {
				Type: schema.TypeString, Required: true, Description: "Logging_name",
			},
		},
	}
}
func resourceCgnv6TemplateLoggingEnableLogByDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingEnableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateLoggingEnableLogByDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingEnableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateLoggingEnableLogByDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateLoggingEnableLogByDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingEnableLogByDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingEnableLogByDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.TcpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationIpListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationIpListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIpListTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpListTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIpListUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpListUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIpListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIp6List(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6List {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.TcpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationIp6ListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationIp6ListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIp6ListTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6ListTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6ListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6ListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationIp6ListUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6ListUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6ListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationIp6ListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingEnableLogByDestinationUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingEnableLogByDestinationUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingEnableLogByDestinationUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingEnableLogByDestinationUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6TemplateLoggingEnableLogByDestination(d *schema.ResourceData) edpt.Cgnv6TemplateLoggingEnableLogByDestination {
	var ret edpt.Cgnv6TemplateLoggingEnableLogByDestination
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.IpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationIpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Ip6List = getSliceCgnv6TemplateLoggingEnableLogByDestinationIp6List(d.Get("ip6_list").([]interface{}))
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceCgnv6TemplateLoggingEnableLogByDestinationUdpList(d.Get("udp_list").([]interface{}))
	//omit uuid
	ret.Inst.Logging_name = d.Get("logging_name").(string)
	return ret
}
