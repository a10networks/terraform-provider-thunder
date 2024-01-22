package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateLoggingDisableLogByDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_logging_disable_log_by_destination`: Disable logging by destination ip address protocol and port\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateLoggingDisableLogByDestinationCreate,
		UpdateContext: resourceCgnv6TemplateLoggingDisableLogByDestinationUpdate,
		ReadContext:   resourceCgnv6TemplateLoggingDisableLogByDestinationRead,
		DeleteContext: resourceCgnv6TemplateLoggingDisableLogByDestinationDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for other L4 protocols",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for other L4 protocols",
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
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for other L4 protocols",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6TemplateLoggingDisableLogByDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingDisableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateLoggingDisableLogByDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingDisableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateLoggingDisableLogByDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateLoggingDisableLogByDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDisableLogByDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLoggingDisableLogByDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6List(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6List {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationTcpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationTcpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationUdpList(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationUdpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6TemplateLoggingDisableLogByDestination(d *schema.ResourceData) edpt.Cgnv6TemplateLoggingDisableLogByDestination {
	var ret edpt.Cgnv6TemplateLoggingDisableLogByDestination
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.IpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Ip6List = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6List(d.Get("ip6_list").([]interface{}))
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationUdpList(d.Get("udp_list").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
