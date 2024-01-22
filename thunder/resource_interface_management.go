package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceManagement() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_management`: Management interface\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceManagementCreate,
		UpdateContext: resourceInterfaceManagementUpdate,
		ReadContext:   resourceInterfaceManagementRead,
		DeleteContext: resourceInterfaceManagementDelete,

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
						"dhcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
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
			"lldp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rt_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp enable/disable",
									},
									"rx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp rx",
									},
									"tx": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable lldp tx",
									},
								},
							},
						},
						"notification_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp notification configuration",
									},
									"notif_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp notification enable",
									},
								},
							},
						},
						"tx_dot1_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tx_dot1_tlvs": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration",
									},
									"link_aggregation": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface link aggregation information",
									},
									"vlan": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface vlan information",
									},
								},
							},
						},
						"tx_tlvs_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tx_tlvs": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp tx TLVs configuration",
									},
									"exclude": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure which TLVs excluded. All basic TLVs will be included by default",
									},
									"management_address": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp management address",
									},
									"port_description": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp port description",
									},
									"system_capabilities": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system capabilities",
									},
									"system_description": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system description",
									},
									"system_name": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface lldp system name",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"secondary_ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secondary_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Global IP configuration subcommands",
						},
						"ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address",
						},
						"ipv4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
						},
						"dhcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
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
			"speed": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'10': 10 Mbs/sec; '100': 100 Mbs/sec; '1000': 1 Gb/sec; 'auto': Auto Negotiate Speed;  (Interface Speed)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceInterfaceManagementCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceManagementRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceManagementUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceManagementRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceManagementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceManagementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceManagementRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceManagement(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceManagementAccessList(d []interface{}) edpt.InterfaceManagementAccessList {

	count1 := len(d)
	var ret edpt.InterfaceManagementAccessList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AclId = in["acl_id"].(int)
		ret.AclName = in["acl_name"].(string)
	}
	return ret
}

func getObjectInterfaceManagementBroadcastRateLimit(d []interface{}) edpt.InterfaceManagementBroadcastRateLimit {

	count1 := len(d)
	var ret edpt.InterfaceManagementBroadcastRateLimit
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BcastRateLimitEnable = in["bcast_rate_limit_enable"].(int)
		ret.Rate = in["rate"].(int)
	}
	return ret
}

func getObjectInterfaceManagementIp(d []interface{}) edpt.InterfaceManagementIp {

	count1 := len(d)
	var ret edpt.InterfaceManagementIp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.Dhcp = in["dhcp"].(int)
		ret.ControlAppsUseMgmtPort = in["control_apps_use_mgmt_port"].(int)
		ret.DefaultGateway = in["default_gateway"].(string)
	}
	return ret
}

func getSliceInterfaceManagementIpv6(d []interface{}) []edpt.InterfaceManagementIpv6 {

	count1 := len(d)
	ret := make([]edpt.InterfaceManagementIpv6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceManagementIpv6
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.V6AclName = in["v6_acl_name"].(string)
		oi.Inbound = in["inbound"].(int)
		oi.DefaultIpv6Gateway = in["default_ipv6_gateway"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceManagementLldp737(d []interface{}) edpt.InterfaceManagementLldp737 {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldp737
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EnableCfg = getObjectInterfaceManagementLldpEnableCfg738(in["enable_cfg"].([]interface{}))
		ret.NotificationCfg = getObjectInterfaceManagementLldpNotificationCfg739(in["notification_cfg"].([]interface{}))
		ret.TxDot1Cfg = getObjectInterfaceManagementLldpTxDot1Cfg740(in["tx_dot1_cfg"].([]interface{}))
		ret.TxTlvsCfg = getObjectInterfaceManagementLldpTxTlvsCfg741(in["tx_tlvs_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceManagementLldpEnableCfg738(d []interface{}) edpt.InterfaceManagementLldpEnableCfg738 {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpEnableCfg738
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RtEnable = in["rt_enable"].(int)
		ret.Rx = in["rx"].(int)
		ret.Tx = in["tx"].(int)
	}
	return ret
}

func getObjectInterfaceManagementLldpNotificationCfg739(d []interface{}) edpt.InterfaceManagementLldpNotificationCfg739 {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpNotificationCfg739
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notification = in["notification"].(int)
		ret.NotifEnable = in["notif_enable"].(int)
	}
	return ret
}

func getObjectInterfaceManagementLldpTxDot1Cfg740(d []interface{}) edpt.InterfaceManagementLldpTxDot1Cfg740 {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpTxDot1Cfg740
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TxDot1Tlvs = in["tx_dot1_tlvs"].(int)
		ret.LinkAggregation = in["link_aggregation"].(int)
		ret.Vlan = in["vlan"].(int)
	}
	return ret
}

func getObjectInterfaceManagementLldpTxTlvsCfg741(d []interface{}) edpt.InterfaceManagementLldpTxTlvsCfg741 {

	count1 := len(d)
	var ret edpt.InterfaceManagementLldpTxTlvsCfg741
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TxTlvs = in["tx_tlvs"].(int)
		ret.Exclude = in["exclude"].(int)
		ret.ManagementAddress = in["management_address"].(int)
		ret.PortDescription = in["port_description"].(int)
		ret.SystemCapabilities = in["system_capabilities"].(int)
		ret.SystemDescription = in["system_description"].(int)
		ret.SystemName = in["system_name"].(int)
	}
	return ret
}

func getSliceInterfaceManagementSamplingEnable(d []interface{}) []edpt.InterfaceManagementSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.InterfaceManagementSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceManagementSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceManagementSecondaryIp(d []interface{}) edpt.InterfaceManagementSecondaryIp {

	count1 := len(d)
	var ret edpt.InterfaceManagementSecondaryIp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecondaryIp = in["secondary_ip"].(int)
		ret.Ipv4Address = in["ipv4_address"].(string)
		ret.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret.Dhcp = in["dhcp"].(int)
		ret.ControlAppsUseMgmtPort = in["control_apps_use_mgmt_port"].(int)
		ret.DefaultGateway = in["default_gateway"].(string)
	}
	return ret
}

func dataToEndpointInterfaceManagement(d *schema.ResourceData) edpt.InterfaceManagement {
	var ret edpt.InterfaceManagement
	ret.Inst.AccessList = getObjectInterfaceManagementAccessList(d.Get("access_list").([]interface{}))
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.BroadcastRateLimit = getObjectInterfaceManagementBroadcastRateLimit(d.Get("broadcast_rate_limit").([]interface{}))
	ret.Inst.Duplexity = d.Get("duplexity").(string)
	ret.Inst.FlowControl = d.Get("flow_control").(int)
	ret.Inst.Ip = getObjectInterfaceManagementIp(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getSliceInterfaceManagementIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.Lldp = getObjectInterfaceManagementLldp737(d.Get("lldp").([]interface{}))
	ret.Inst.Mtu = d.Get("mtu").(int)
	ret.Inst.SamplingEnable = getSliceInterfaceManagementSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SecondaryIp = getObjectInterfaceManagementSecondaryIp(d.Get("secondary_ip").([]interface{}))
	ret.Inst.Speed = d.Get("speed").(string)
	//omit uuid
	return ret
}
