package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTemplateLoggingDisableLogByDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_template_logging_disable_log_by_destination`: Disable firewall logging by the destination IP address, protocol, and port\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTemplateLoggingDisableLogByDestinationCreate,
		UpdateContext: resourceFwTemplateLoggingDisableLogByDestinationUpdate,
		ReadContext:   resourceFwTemplateLoggingDisableLogByDestinationRead,
		DeleteContext: resourceFwTemplateLoggingDisableLogByDestinationDelete,

		Schema: map[string]*schema.Schema{
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the ICMP traffic",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the ICMP traffic",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the other layer-4 protocols",
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
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the ICMP traffic",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the other layer-4 protocols",
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
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for the other layer-4 protocols",
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
func resourceFwTemplateLoggingDisableLogByDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingDisableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTemplateLoggingDisableLogByDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingDisableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTemplateLoggingDisableLogByDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTemplateLoggingDisableLogByDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingDisableLogByDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingDisableLogByDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwTemplateLoggingDisableLogByDestinationIpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIpList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.TcpList = getSliceFwTemplateLoggingDisableLogByDestinationIpListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceFwTemplateLoggingDisableLogByDestinationIpListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationIpListTcpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIpListTcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIpListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIpListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationIpListUdpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIpListUdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIpListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIpListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationIp6List(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIp6List {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIp6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIp6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.TcpList = getSliceFwTemplateLoggingDisableLogByDestinationIp6ListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceFwTemplateLoggingDisableLogByDestinationIp6ListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationIp6ListTcpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIp6ListTcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIp6ListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIp6ListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationIp6ListUdpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationIp6ListUdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationIp6ListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationIp6ListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationTcpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationTcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingDisableLogByDestinationUdpList(d []interface{}) []edpt.FwTemplateLoggingDisableLogByDestinationUdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingDisableLogByDestinationUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingDisableLogByDestinationUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTemplateLoggingDisableLogByDestination(d *schema.ResourceData) edpt.FwTemplateLoggingDisableLogByDestination {
	var ret edpt.FwTemplateLoggingDisableLogByDestination
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.IpList = getSliceFwTemplateLoggingDisableLogByDestinationIpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Ip6List = getSliceFwTemplateLoggingDisableLogByDestinationIp6List(d.Get("ip6_list").([]interface{}))
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceFwTemplateLoggingDisableLogByDestinationTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceFwTemplateLoggingDisableLogByDestinationUdpList(d.Get("udp_list").([]interface{}))
	//omit uuid
	ret.Inst.Logging_name = d.Get("logging_name").(string)
	return ret
}
