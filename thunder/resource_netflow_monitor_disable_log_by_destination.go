package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitorDisableLogByDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor_disable_log_by_destination`: Disable logging by destination ip address protocol and port\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorDisableLogByDestinationCreate,
		UpdateContext: resourceNetflowMonitorDisableLogByDestinationUpdate,
		ReadContext:   resourceNetflowMonitorDisableLogByDestinationRead,
		DeleteContext: resourceNetflowMonitorDisableLogByDestinationDelete,

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
func resourceNetflowMonitorDisableLogByDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDisableLogByDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDisableLogByDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorDisableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorDisableLogByDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDisableLogByDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDisableLogByDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorDisableLogByDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorDisableLogByDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDisableLogByDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDisableLogByDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorDisableLogByDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDisableLogByDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitorDisableLogByDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetflowMonitorDisableLogByDestinationIpList(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIpList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIpList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.TcpList = getSliceNetflowMonitorDisableLogByDestinationIpListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceNetflowMonitorDisableLogByDestinationIpListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIpListTcpList(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIpListTcpList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIpListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIpListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIpListUdpList(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIpListUdpList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIpListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIpListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIp6List(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIp6List {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIp6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIp6List
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.TcpList = getSliceNetflowMonitorDisableLogByDestinationIp6ListTcpList(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceNetflowMonitorDisableLogByDestinationIp6ListUdpList(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIp6ListTcpList(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIp6ListTcpList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIp6ListTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIp6ListTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIp6ListUdpList(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIp6ListUdpList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIp6ListUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIp6ListUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationTcpList(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationTcpList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationTcpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationTcpList
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationUdpList(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationUdpList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationUdpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationUdpList
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetflowMonitorDisableLogByDestination(d *schema.ResourceData) edpt.NetflowMonitorDisableLogByDestination {
	var ret edpt.NetflowMonitorDisableLogByDestination
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.IpList = getSliceNetflowMonitorDisableLogByDestinationIpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Ip6List = getSliceNetflowMonitorDisableLogByDestinationIp6List(d.Get("ip6_list").([]interface{}))
	ret.Inst.Others = d.Get("others").(int)
	ret.Inst.TcpList = getSliceNetflowMonitorDisableLogByDestinationTcpList(d.Get("tcp_list").([]interface{}))
	ret.Inst.UdpList = getSliceNetflowMonitorDisableLogByDestinationUdpList(d.Get("udp_list").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
