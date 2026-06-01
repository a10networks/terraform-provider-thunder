package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnelIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_tunnel_ip`: Global IP configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTunnelIpCreate,
		UpdateContext: resourceInterfaceTunnelIpUpdate,
		ReadContext:   resourceInterfaceTunnelIpRead,
		DeleteContext: resourceInterfaceTunnelIpDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dhcp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
						},
						"ip_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"ipv4_netmask": {
										Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
									},
								},
							},
						},
					},
				},
			},
			"generate_membership_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Membership Query",
			},
			"generate_membership_query_val": {
				Type: schema.TypeInt, Optional: true, Default: 125, Description: "1 - 255 (Default is 125)",
			},
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
			},
			"max_resp_time": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Max Response Time (Default is 100)",
			},
			"ospf": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf_global": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
												},
												"value": {
													Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
												},
											},
										},
									},
									"authentication_key": {
										Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
									},
									"bfd_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bfd": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
												},
												"disable": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
												},
											},
										},
									},
									"cost": {
										Type: schema.TypeInt, Optional: true, Description: "Interface cost",
									},
									"database_filter_cfg": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"database_filter": {
													Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
												},
												"out": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
												},
											},
										},
									},
									"dead_interval": {
										Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
									},
									"disable": {
										Type: schema.TypeString, Optional: true, Description: "'all': All functionality;",
									},
									"hello_interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
									},
									"message_digest_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message_digest_key": {
													Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
												},
												"md5": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"md5_value": {
																Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
															},
														},
													},
												},
											},
										},
									},
									"mtu": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF interface MTU (MTU size)",
									},
									"mtu_ignore": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
									},
									"network": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"broadcast": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF broadcast multi-access network",
												},
												"non_broadcast": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF NBMA network",
												},
												"point_to_point": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-point network",
												},
												"point_to_multipoint": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify OSPF point-to-multipoint network",
												},
												"p2mp_nbma": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify non-broadcast point-to-multipoint network",
												},
											},
										},
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
									},
									"retransmit_interval": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
									},
									"transmit_delay": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ospf_ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_addr": {
										Type: schema.TypeString, Required: true, Description: "Address of interface",
									},
									"authentication": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
									},
									"value": {
										Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use no authentication;",
									},
									"authentication_key": {
										Type: schema.TypeString, Optional: true, Description: "Authentication password (key) (The OSPF password (key))",
									},
									"cost": {
										Type: schema.TypeInt, Optional: true, Description: "Interface cost",
									},
									"database_filter": {
										Type: schema.TypeString, Optional: true, Description: "'all': Filter all LSA;",
									},
									"out": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing LSA",
									},
									"dead_interval": {
										Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
									},
									"hello_interval": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
									},
									"message_digest_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message_digest_key": {
													Type: schema.TypeInt, Optional: true, Description: "Message digest authentication password (key) (Key id)",
												},
												"md5_value": {
													Type: schema.TypeString, Optional: true, Description: "The OSPF password (1-16)",
												},
											},
										},
									},
									"mtu_ignore": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
									},
									"priority": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
									},
									"retransmit_interval": {
										Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
									},
									"transmit_delay": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
			},
			"rip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": {
													Type: schema.TypeString, Optional: true, Description: "The RIP authentication string",
												},
											},
										},
									},
									"mode": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type: schema.TypeString, Optional: true, Default: "text", Description: "'md5': Keyed message digest; 'text': Clear text authentication;",
												},
											},
										},
									},
									"key_chain": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_chain": {
													Type: schema.TypeString, Optional: true, Description: "Authentication key-chain (Name of key-chain)",
												},
											},
										},
									},
								},
							},
						},
						"send_packet": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable sending packets through the specified interface",
						},
						"receive_packet": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable receiving packet through the specified interface",
						},
						"send_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"send": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertisement transmission",
									},
									"version": {
										Type: schema.TypeString, Optional: true, Description: "'2': RIP version 2;",
									},
								},
							},
						},
						"receive_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"receive": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertisement reception",
									},
									"version": {
										Type: schema.TypeString, Optional: true, Description: "'2': RIP version 2;",
									},
								},
							},
						},
						"split_horizon_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type: schema.TypeString, Optional: true, Default: "poisoned", Description: "'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTunnelIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTunnelIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTunnelIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTunnelIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTunnelIpAddress(d []interface{}) edpt.InterfaceTunnelIpAddress {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpAddress
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dhcp = in["dhcp"].(int)
		ret.IpCfg = getSliceInterfaceTunnelIpAddressIpCfg(in["ip_cfg"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTunnelIpAddressIpCfg(d []interface{}) []edpt.InterfaceTunnelIpAddressIpCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpAddressIpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpAddressIpCfg
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspf937(d []interface{}) edpt.InterfaceTunnelIpOspf937 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspf937
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceTunnelIpOspfOspfGlobal938(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceTunnelIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobal938(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobal938 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobal938
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceTunnelIpOspfOspfGlobalAuthenticationCfg939(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceTunnelIpOspfOspfGlobalBfdCfg940(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg941(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceTunnelIpOspfOspfGlobalMessageDigestCfg942(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceTunnelIpOspfOspfGlobalNetwork944(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalAuthenticationCfg939(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg939 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg939
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalBfdCfg940(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalBfdCfg940 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalBfdCfg940
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg941(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg941 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg941
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceTunnelIpOspfOspfGlobalMessageDigestCfg942(d []interface{}) []edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg942 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg942, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg942
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5943(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5943(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5943 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5943
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceTunnelIpOspfOspfGlobalNetwork944(d []interface{}) edpt.InterfaceTunnelIpOspfOspfGlobalNetwork944 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpOspfOspfGlobalNetwork944
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Broadcast = in["broadcast"].(int)
		ret.NonBroadcast = in["non_broadcast"].(int)
		ret.PointToPoint = in["point_to_point"].(int)
		ret.PointToMultipoint = in["point_to_multipoint"].(int)
		ret.P2mpNbma = in["p2mp_nbma"].(int)
	}
	return ret
}

func getSliceInterfaceTunnelIpOspfOspfIpList(d []interface{}) []edpt.InterfaceTunnelIpOspfOspfIpList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpOspfOspfIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceTunnelIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceTunnelIpOspfOspfIpListMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpOspfOspfIpListMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpRip945(d []interface{}) edpt.InterfaceTunnelIpRip945 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRip945
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceTunnelIpRipAuthentication946(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceTunnelIpRipSendCfg950(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceTunnelIpRipReceiveCfg951(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceTunnelIpRipSplitHorizonCfg952(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthentication946(d []interface{}) edpt.InterfaceTunnelIpRipAuthentication946 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthentication946
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceTunnelIpRipAuthenticationStr947(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceTunnelIpRipAuthenticationMode948(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceTunnelIpRipAuthenticationKeyChain949(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationStr947(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationStr947 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationStr947
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationMode948(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationMode948 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationMode948
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationKeyChain949(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationKeyChain949 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationKeyChain949
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipSendCfg950(d []interface{}) edpt.InterfaceTunnelIpRipSendCfg950 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipSendCfg950
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipReceiveCfg951(d []interface{}) edpt.InterfaceTunnelIpRipReceiveCfg951 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipReceiveCfg951
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipSplitHorizonCfg952(d []interface{}) edpt.InterfaceTunnelIpRipSplitHorizonCfg952 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipSplitHorizonCfg952
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointInterfaceTunnelIp(d *schema.ResourceData) edpt.InterfaceTunnelIp {
	var ret edpt.InterfaceTunnelIp
	ret.Inst.Address = getObjectInterfaceTunnelIpAddress(d.Get("address").([]interface{}))
	ret.Inst.GenerateMembershipQuery = d.Get("generate_membership_query").(int)
	ret.Inst.GenerateMembershipQueryVal = d.Get("generate_membership_query_val").(int)
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.MaxRespTime = d.Get("max_resp_time").(int)
	ret.Inst.Ospf = getObjectInterfaceTunnelIpOspf937(d.Get("ospf").([]interface{}))
	ret.Inst.Outside = d.Get("outside").(int)
	ret.Inst.Rip = getObjectInterfaceTunnelIpRip945(d.Get("rip").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
