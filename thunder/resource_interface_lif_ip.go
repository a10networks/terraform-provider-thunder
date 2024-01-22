package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ip`: Global IP configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpCreate,
		UpdateContext: resourceInterfaceLifIpUpdate,
		ReadContext:   resourceInterfaceLifIpRead,
		DeleteContext: resourceInterfaceLifIpDelete,

		Schema: map[string]*schema.Schema{
			"address_list": {
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
			"allow_promiscuous_vip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow traffic to be associated with promiscuous VIP",
			},
			"cache_spoofing_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "This interface connects to spoofing cache",
			},
			"dhcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
			},
			"generate_membership_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Membership Query",
			},
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
			},
			"max_resp_time": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Maximum Response Time (Max Response Time (Default is 100))",
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
			"query_interval": {
				Type: schema.TypeInt, Optional: true, Default: 125, Description: "1 - 255 (Default is 125)",
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
										Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-compatible': RIPv1-compatible; '1-2': RIP version 1 & 2;",
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
										Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;",
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
			"router": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"isis": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag": {
										Type: schema.TypeString, Optional: true, Description: "ISO routing area tag",
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
			"unnumbered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface as unnumbered",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLifIpAddressList(d []interface{}) []edpt.InterfaceLifIpAddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpAddressList
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpOspf560(d []interface{}) edpt.InterfaceLifIpOspf560 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspf560
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceLifIpOspfOspfGlobal561(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceLifIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobal561(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobal561 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobal561
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg562(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceLifIpOspfOspfGlobalBfdCfg563(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg564(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg565(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceLifIpOspfOspfGlobalNetwork567(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalAuthenticationCfg562(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg562 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalAuthenticationCfg562
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalBfdCfg563(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalBfdCfg563 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalBfdCfg563
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalDatabaseFilterCfg564(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg564 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg564
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfGlobalMessageDigestCfg565(d []interface{}) []edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg565 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg565, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfg565
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5566(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5566(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5566 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5566
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceLifIpOspfOspfGlobalNetwork567(d []interface{}) edpt.InterfaceLifIpOspfOspfGlobalNetwork567 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpOspfOspfGlobalNetwork567
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

func getSliceInterfaceLifIpOspfOspfIpList(d []interface{}) []edpt.InterfaceLifIpOspfOspfIpList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceLifIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpRip568(d []interface{}) edpt.InterfaceLifIpRip568 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRip568
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceLifIpRipAuthentication569(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceLifIpRipSendCfg573(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceLifIpRipReceiveCfg574(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceLifIpRipSplitHorizonCfg575(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthentication569(d []interface{}) edpt.InterfaceLifIpRipAuthentication569 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthentication569
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceLifIpRipAuthenticationStr570(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLifIpRipAuthenticationMode571(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLifIpRipAuthenticationKeyChain572(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationStr570(d []interface{}) edpt.InterfaceLifIpRipAuthenticationStr570 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationStr570
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationMode571(d []interface{}) edpt.InterfaceLifIpRipAuthenticationMode571 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationMode571
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationKeyChain572(d []interface{}) edpt.InterfaceLifIpRipAuthenticationKeyChain572 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationKeyChain572
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSendCfg573(d []interface{}) edpt.InterfaceLifIpRipSendCfg573 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipSendCfg573
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipReceiveCfg574(d []interface{}) edpt.InterfaceLifIpRipReceiveCfg574 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipReceiveCfg574
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSplitHorizonCfg575(d []interface{}) edpt.InterfaceLifIpRipSplitHorizonCfg575 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipSplitHorizonCfg575
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRouter576(d []interface{}) edpt.InterfaceLifIpRouter576 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRouter576
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceLifIpRouterIsis577(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpRouterIsis577(d []interface{}) edpt.InterfaceLifIpRouterIsis577 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRouterIsis577
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointInterfaceLifIp(d *schema.ResourceData) edpt.InterfaceLifIp {
	var ret edpt.InterfaceLifIp
	ret.Inst.AddressList = getSliceInterfaceLifIpAddressList(d.Get("address_list").([]interface{}))
	ret.Inst.AllowPromiscuousVip = d.Get("allow_promiscuous_vip").(int)
	ret.Inst.CacheSpoofingPort = d.Get("cache_spoofing_port").(int)
	ret.Inst.Dhcp = d.Get("dhcp").(int)
	ret.Inst.GenerateMembershipQuery = d.Get("generate_membership_query").(int)
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.MaxRespTime = d.Get("max_resp_time").(int)
	ret.Inst.Ospf = getObjectInterfaceLifIpOspf560(d.Get("ospf").([]interface{}))
	ret.Inst.Outside = d.Get("outside").(int)
	ret.Inst.QueryInterval = d.Get("query_interval").(int)
	ret.Inst.Rip = getObjectInterfaceLifIpRip568(d.Get("rip").([]interface{}))
	ret.Inst.Router = getObjectInterfaceLifIpRouter576(d.Get("router").([]interface{}))
	ret.Inst.Unnumbered = d.Get("unnumbered").(int)
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
