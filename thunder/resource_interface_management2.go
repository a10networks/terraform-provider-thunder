package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceManagement2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_management2`: Management 2 interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceManagement2Create,
		UpdateContext: resourceInterfaceManagement2Update,
		ReadContext:   resourceInterfaceManagement2Read,
		DeleteContext: resourceInterfaceManagement2Delete,

		Schema: map[string]*schema.Schema{
			"access_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type: schema.TypeInt, Optional: true, Description: "ACL id",
						},
						"acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply an access list (Named Access List)",
						},
					},
				},
			},
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Management Port; 'disable': Disable Management Port;",
			},
			"broadcast_rate_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bcast_rate_limit_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rate limit the l2 broadcast packet on mgmt port",
						},
						"rate": {
							Type: schema.TypeInt, Optional: true, Default: 500, Description: "packets per second. Default is 500. (packets per second. Please specify an even number. Default is 500)",
						},
					},
				},
			},
			"duplexity": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'Full': Full; 'Half': Half; 'auto': Auto;",
			},
			"flow_control": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable 802.3x flow control on full duplex port",
			},
			"ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
						},
						"control_apps_use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Control applications use management port",
						},
						"default_gateway": {
							Type: schema.TypeString, Optional: true, Description: "Set default gateway (Default gateway address)",
						},
					},
				},
			},
			"ipv6": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Set the IPv6 address of an interface",
						},
						"address_type": {
							Type: schema.TypeString, Optional: true, Description: "'link-local': Configure an IPv6 link local address;",
						},
						"v6_acl_name": {
							Type: schema.TypeString, Optional: true, Description: "Apply ACL rules to incoming packets on this interface (Named Access List)",
						},
						"inbound": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ACL applied on incoming packets to this interface",
						},
						"default_ipv6_gateway": {
							Type: schema.TypeString, Optional: true, Description: "Set default gateway (Default gateway address)",
						},
					},
				},
			},
			"mtu": {
				Type: schema.TypeInt, Optional: true, Description: "Interface mtu (Interface MTU, default 1 (min MTU is 1280 for IPv6))",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets_input': Input packets; 'bytes_input': Input bytes; 'received_broadcasts': Received broadcasts; 'received_multicasts': Received multicasts; 'received_unicasts': Received unicasts; 'input_errors': Input errors; 'crc': CRC; 'frame': Frames; 'input_err_short': Runts; 'input_err_long': Giants; 'packets_output': Output packets; 'bytes_output': Output bytes; 'transmitted_broadcasts': Transmitted broadcasts; 'transmitted_multicasts': Transmitted multicasts; 'transmitted_unicasts': Transmitted unicasts; 'output_errors': Output errors; 'collisions': Collisions;",
						},
					},
				},
			},
			"speed": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'10': 10 Mbs/sec; '100': 100 Mbs/sec; '1000': 1 Gb/sec; 'auto': Auto Negotiate Speed;  (Interface Speed)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceInterfaceManagement2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagement2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceManagement2Read(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceManagement2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagement2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceManagement2Read(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceManagement2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagement2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceManagement2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagement2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceManagement2AccessList(d []interface{}) edpt.InterfaceManagement2AccessList {

	count1 := len(d)
	var ret edpt.InterfaceManagement2AccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
	}
	return ret
}

func getObjectInterfaceManagement2BroadcastRateLimit(d []interface{}) edpt.InterfaceManagement2BroadcastRateLimit {

	count1 := len(d)
	var ret edpt.InterfaceManagement2BroadcastRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BcastRateLimitEnable = in["bcast_rate_limit_enable"].(int)
		ret.Rate = in["rate"].(int)
	}
	return ret
}

func getObjectInterfaceManagement2Ip(d []interface{}) edpt.InterfaceManagement2Ip {

	count1 := len(d)
	var ret edpt.InterfaceManagement2Ip
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.ControlAppsUseMgmtPort = in["control_apps_use_mgmt_port"].(int)
		ret.DefaultGateway = in["default_gateway"].(string)
	}
	return ret
}

func getSliceInterfaceManagement2Ipv6(d []interface{}) []edpt.InterfaceManagement2Ipv6 {

	count1 := len(d)
	ret := make([]edpt.InterfaceManagement2Ipv6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceManagement2Ipv6
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.V6AclName = in["v6_acl_name"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.DefaultIpv6Gateway = in["default_ipv6_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceManagement2SamplingEnable(d []interface{}) []edpt.InterfaceManagement2SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.InterfaceManagement2SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceManagement2SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceManagement2(d *schema.ResourceData) edpt.InterfaceManagement2 {
	var ret edpt.InterfaceManagement2
	ret.Inst.AccessList = getObjectInterfaceManagement2AccessList(d.Get("access_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.BroadcastRateLimit = getObjectInterfaceManagement2BroadcastRateLimit(d.Get("broadcast_rate_limit").([]interface{}))
	ret.Inst.Duplexity = d.Get("duplexity").(string)
	ret.Inst.FlowControl = d.Get("flow_control").(int)
	ret.Inst.Ip = getObjectInterfaceManagement2Ip(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getSliceInterfaceManagement2Ipv6(d.Get("ipv6").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.SamplingEnable = getSliceInterfaceManagement2SamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Speed = d.Get("speed").(string)
	//omit uuid
	return ret
}
