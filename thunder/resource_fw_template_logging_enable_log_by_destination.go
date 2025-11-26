package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTemplateLoggingEnableLogByDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_template_logging_enable_log_by_destination`: Enable firewall logging by the destination IP address, protocol, and port\n\n__PLACEHOLDER__",
		CreateContext: resourceFwTemplateLoggingEnableLogByDestinationCreate,
		UpdateContext: resourceFwTemplateLoggingEnableLogByDestinationUpdate,
		ReadContext:   resourceFwTemplateLoggingEnableLogByDestinationRead,
		DeleteContext: resourceFwTemplateLoggingEnableLogByDestinationDelete,

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
func resourceFwTemplateLoggingEnableLogByDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingEnableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceFwTemplateLoggingEnableLogByDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwTemplateLoggingEnableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceFwTemplateLoggingEnableLogByDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwTemplateLoggingEnableLogByDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTemplateLoggingEnableLogByDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTemplateLoggingEnableLogByDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwTemplateLoggingEnableLogByDestinationIpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIpList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.TcpList = getSliceFwTemplateLoggingEnableLogByDestinationIpListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceFwTemplateLoggingEnableLogByDestinationIpListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationIpListTcpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIpListTcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIpListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIpListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationIpListUdpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIpListUdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIpListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIpListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationIp6List(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIp6List {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIp6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIp6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.TcpList = getSliceFwTemplateLoggingEnableLogByDestinationIp6ListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceFwTemplateLoggingEnableLogByDestinationIp6ListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationIp6ListTcpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIp6ListTcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIp6ListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIp6ListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationIp6ListUdpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationIp6ListUdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationIp6ListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationIp6ListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationTcpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationTcpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTemplateLoggingEnableLogByDestinationUdpList(d []interface{}) []edpt.FwTemplateLoggingEnableLogByDestinationUdpList {

	count1 := len(d)
	ret := make([]edpt.FwTemplateLoggingEnableLogByDestinationUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTemplateLoggingEnableLogByDestinationUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTemplateLoggingEnableLogByDestination(d *schema.ResourceData) edpt.FwTemplateLoggingEnableLogByDestination {
	var ret edpt.FwTemplateLoggingEnableLogByDestination
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.IpList = getSliceFwTemplateLoggingEnableLogByDestinationIpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Ip6List = getSliceFwTemplateLoggingEnableLogByDestinationIp6List(d.Get("ip6_list").([]interface{}))
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceFwTemplateLoggingEnableLogByDestinationTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceFwTemplateLoggingEnableLogByDestinationUdpList(d.Get("udp_list").([]interface{}))
	//omit uuid
	ret.Inst.Logging_name = d.Get("logging_name").(string)
	return ret
}
