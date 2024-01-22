package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosExecScript() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_exec_script`: Execute scripts\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosExecScriptCreate,
		UpdateContext: resourceDdosExecScriptUpdate,
		ReadContext:   resourceDdosExecScriptRead,
		DeleteContext: resourceDdosExecScriptDelete,

		Schema: map[string]*schema.Schema{
			"alert_type": {
				Type: schema.TypeInt, Optional: true, Description: "1: UDP Pkt Rate 2: TCP Pkt Rate 3: ICMP Pkt Rate",
			},
			"exec_script_ip_portocol": {
				Type: schema.TypeString, Optional: true, Description: "'icmp-v4': ip-proto icmp-v4; 'icmp-v6': ip-proto icmp-v6; 'other': ip-proto other; 'gre': ip-proto gre; 'ipv4-encap': ip-proto IPv4 Encapsulation; 'ipv6-encap': ip-proto IPv6 Encapsulation;",
			},
			"exec_script_port_other_protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP Port; 'udp': UDP Port;",
			},
			"level": {
				Type: schema.TypeInt, Optional: true, Description: "Current Level",
			},
			"mock": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use mock data",
			},
			"port_num": {
				Type: schema.TypeInt, Optional: true, Description: "Port Number",
			},
			"port_other": {
				Type: schema.TypeString, Optional: true, Description: "'other': other;",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-tcp': SIP-TCP Port; 'sip-udp': SIP-UDP Port; 'quic': QUIC Port;",
			},
			"protocol_num": {
				Type: schema.TypeInt, Optional: true, Description: "Protocol Number",
			},
			"script": {
				Type: schema.TypeString, Optional: true, Description: "Specify script to execute",
			},
			"src_ip": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IP address",
						},
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "IP Subnet",
						},
					},
				},
			},
			"src_ipv6": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 address",
						},
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 Subnet",
						},
					},
				},
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Timeout (Default: 10 seconds, Mock Default: 2 seconds)",
			},
			"zone": {
				Type: schema.TypeString, Optional: true, Description: "DST Zone name",
			},
		},
	}
}
func resourceDdosExecScriptCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosExecScriptCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosExecScript(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosExecScriptRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosExecScriptUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosExecScriptUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosExecScript(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosExecScriptRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosExecScriptDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosExecScriptDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosExecScript(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosExecScriptRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosExecScriptRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosExecScript(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosExecScriptSrcIp(d []interface{}) []edpt.DdosExecScriptSrcIp {

	count1 := len(d)
	ret := make([]edpt.DdosExecScriptSrcIp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosExecScriptSrcIp
		oi.IpAddr = in["ip_addr"].(string)
		oi.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosExecScriptSrcIpv6(d []interface{}) []edpt.DdosExecScriptSrcIpv6 {

	count1 := len(d)
	ret := make([]edpt.DdosExecScriptSrcIpv6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosExecScriptSrcIpv6
		oi.Ip6Addr = in["ip6_addr"].(string)
		oi.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosExecScript(d *schema.ResourceData) edpt.DdosExecScript {
	var ret edpt.DdosExecScript
	ret.Inst.AlertType = d.Get("alert_type").(int)
	ret.Inst.ExecScriptIpPortocol = d.Get("exec_script_ip_portocol").(string)
	ret.Inst.ExecScriptPortOtherProtocol = d.Get("exec_script_port_other_protocol").(string)
	ret.Inst.Level = d.Get("level").(int)
	ret.Inst.Mock = d.Get("mock").(int)
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.PortOther = d.Get("port_other").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.ProtocolNum = d.Get("protocol_num").(int)
	ret.Inst.Script = d.Get("script").(string)
	ret.Inst.SrcIp = getSliceDdosExecScriptSrcIp(d.Get("src_ip").([]interface{}))
	ret.Inst.SrcIpv6 = getSliceDdosExecScriptSrcIpv6(d.Get("src_ipv6").([]interface{}))
	ret.Inst.Threshold = d.Get("threshold").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.Zone = d.Get("zone").(string)
	return ret
}
