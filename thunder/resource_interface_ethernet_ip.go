package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_ip`: Global IP configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetIpCreate,
		UpdateContext: resourceInterfaceEthernetIpUpdate,
		ReadContext:   resourceInterfaceEthernetIpRead,
		DeleteContext: resourceInterfaceEthernetIpDelete,

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
			"client": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client facing interface for IPv4/v6 traffic",
			},
			"dhcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use DHCP to configure IP address",
			},
			"generate_membership_query": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Membership Query",
			},
			"helper_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"helper_address": {
							Type: schema.TypeString, Optional: true, Description: "Helper address for DHCP packets (IP address)",
						},
					},
				},
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
			"server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server facing interface for IPv4/v6 traffic",
			},
			"slb_partition_redirect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Redirect SLB traffic across partition",
			},
			"stateful_firewall": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inside (private) interface for stateful firewall",
						},
						"class_list": {
							Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outside (public) interface for stateful firewall",
						},
						"access_list": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Access-list for traffic from the outside",
						},
						"acl_id": {
							Type: schema.TypeInt, Optional: true, Description: "ACL id",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"syn_cookie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure Enable SYN-cookie on the interface",
			},
			"ttl_ignore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet before sending out",
			},
			"unnumbered": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface as unnumbered",
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
func resourceInterfaceEthernetIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceEthernetIpAddressList(d []interface{}) []edpt.InterfaceEthernetIpAddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpAddressList
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpHelperAddressList(d []interface{}) []edpt.InterfaceEthernetIpHelperAddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpHelperAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpHelperAddressList
		oi.HelperAddress = in["helper_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspf447(d []interface{}) edpt.InterfaceEthernetIpOspf447 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspf447
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceEthernetIpOspfOspfGlobal448(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceEthernetIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobal448(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobal448 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobal448
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceEthernetIpOspfOspfGlobalAuthenticationCfg449(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceEthernetIpOspfOspfGlobalBfdCfg450(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg451(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceEthernetIpOspfOspfGlobalMessageDigestCfg452(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceEthernetIpOspfOspfGlobalNetwork454(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalAuthenticationCfg449(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg449 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg449
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalBfdCfg450(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalBfdCfg450 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalBfdCfg450
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg451(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg451 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg451
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIpOspfOspfGlobalMessageDigestCfg452(d []interface{}) []edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg452 {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg452, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg452
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5453(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5453(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5453 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5453
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceEthernetIpOspfOspfGlobalNetwork454(d []interface{}) edpt.InterfaceEthernetIpOspfOspfGlobalNetwork454 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpOspfOspfGlobalNetwork454
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

func getSliceInterfaceEthernetIpOspfOspfIpList(d []interface{}) []edpt.InterfaceEthernetIpOspfOspfIpList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpOspfOspfIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceEthernetIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceEthernetIpOspfOspfIpListMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIpOspfOspfIpListMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIpRip455(d []interface{}) edpt.InterfaceEthernetIpRip455 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRip455
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceEthernetIpRipAuthentication456(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceEthernetIpRipSendCfg460(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceEthernetIpRipReceiveCfg461(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceEthernetIpRipSplitHorizonCfg462(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthentication456(d []interface{}) edpt.InterfaceEthernetIpRipAuthentication456 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthentication456
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceEthernetIpRipAuthenticationStr457(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceEthernetIpRipAuthenticationMode458(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceEthernetIpRipAuthenticationKeyChain459(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthenticationStr457(d []interface{}) edpt.InterfaceEthernetIpRipAuthenticationStr457 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthenticationStr457
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthenticationMode458(d []interface{}) edpt.InterfaceEthernetIpRipAuthenticationMode458 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthenticationMode458
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipAuthenticationKeyChain459(d []interface{}) edpt.InterfaceEthernetIpRipAuthenticationKeyChain459 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipAuthenticationKeyChain459
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipSendCfg460(d []interface{}) edpt.InterfaceEthernetIpRipSendCfg460 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipSendCfg460
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipReceiveCfg461(d []interface{}) edpt.InterfaceEthernetIpRipReceiveCfg461 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipReceiveCfg461
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRipSplitHorizonCfg462(d []interface{}) edpt.InterfaceEthernetIpRipSplitHorizonCfg462 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRipSplitHorizonCfg462
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceEthernetIpRouter463(d []interface{}) edpt.InterfaceEthernetIpRouter463 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRouter463
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceEthernetIpRouterIsis464(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceEthernetIpRouterIsis464(d []interface{}) edpt.InterfaceEthernetIpRouterIsis464 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpRouterIsis464
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceEthernetIpStatefulFirewall465(d []interface{}) edpt.InterfaceEthernetIpStatefulFirewall465 {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIpStatefulFirewall465
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.ClassList = in["class_list"].(string)
		ret.Outside = in["outside"].(int)
		ret.AccessList = in["access_list"].(int)
		ret.AclId = in["acl_id"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointInterfaceEthernetIp(d *schema.ResourceData) edpt.InterfaceEthernetIp {
	var ret edpt.InterfaceEthernetIp
	ret.Inst.AddressList = getSliceInterfaceEthernetIpAddressList(d.Get("address_list").([]interface{}))
	ret.Inst.AllowPromiscuousVip = d.Get("allow_promiscuous_vip").(int)
	ret.Inst.CacheSpoofingPort = d.Get("cache_spoofing_port").(int)
	ret.Inst.Client = d.Get("client").(int)
	ret.Inst.Dhcp = d.Get("dhcp").(int)
	ret.Inst.GenerateMembershipQuery = d.Get("generate_membership_query").(int)
	ret.Inst.HelperAddressList = getSliceInterfaceEthernetIpHelperAddressList(d.Get("helper_address_list").([]interface{}))
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.MaxRespTime = d.Get("max_resp_time").(int)
	ret.Inst.Ospf = getObjectInterfaceEthernetIpOspf447(d.Get("ospf").([]interface{}))
	ret.Inst.Outside = d.Get("outside").(int)
	ret.Inst.QueryInterval = d.Get("query_interval").(int)
	ret.Inst.Rip = getObjectInterfaceEthernetIpRip455(d.Get("rip").([]interface{}))
	ret.Inst.Router = getObjectInterfaceEthernetIpRouter463(d.Get("router").([]interface{}))
	ret.Inst.Server = d.Get("server").(int)
	ret.Inst.SlbPartitionRedirect = d.Get("slb_partition_redirect").(int)
	ret.Inst.StatefulFirewall = getObjectInterfaceEthernetIpStatefulFirewall465(d.Get("stateful_firewall").([]interface{}))
	ret.Inst.SynCookie = d.Get("syn_cookie").(int)
	ret.Inst.TtlIgnore = d.Get("ttl_ignore").(int)
	ret.Inst.Unnumbered = d.Get("unnumbered").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
