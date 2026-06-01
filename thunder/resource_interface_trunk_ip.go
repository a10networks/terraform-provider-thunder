package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ip`: Global IP configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpCreate,
		UpdateContext: resourceInterfaceTrunkIpUpdate,
		ReadContext:   resourceInterfaceTrunkIpRead,
		DeleteContext: resourceInterfaceTrunkIpDelete,

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
			"dmz": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "DMZ network facing interface for IPv4/v6 traffic",
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
			"max_resp_time": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Maximum Response Time (Max Response Time (Default is 100))",
			},
			"nat": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
						},
						"outside": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
						},
					},
				},
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
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SYN-cookie on the interface",
			},
			"ttl_ignore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore TTL decrement for a received packet",
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
func resourceInterfaceTrunkIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceTrunkIpAddressList(d []interface{}) []edpt.InterfaceTrunkIpAddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpAddressList
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpHelperAddressList(d []interface{}) []edpt.InterfaceTrunkIpHelperAddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpHelperAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpHelperAddressList
		oi.HelperAddress = in["helper_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpNat(d []interface{}) edpt.InterfaceTrunkIpNat {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpNat
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Inside = in["inside"].(int)
		ret.Outside = in["outside"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspf824(d []interface{}) edpt.InterfaceTrunkIpOspf824 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspf824
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceTrunkIpOspfOspfGlobal825(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceTrunkIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobal825(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobal825 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobal825
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceTrunkIpOspfOspfGlobalAuthenticationCfg826(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceTrunkIpOspfOspfGlobalBfdCfg827(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg828(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceTrunkIpOspfOspfGlobalMessageDigestCfg829(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Network = getObjectInterfaceTrunkIpOspfOspfGlobalNetwork831(in["network"].([]interface{}))
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalAuthenticationCfg826(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg826 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalAuthenticationCfg826
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalBfdCfg827(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalBfdCfg827 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalBfdCfg827
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg828(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg828 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalDatabaseFilterCfg828
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIpOspfOspfGlobalMessageDigestCfg829(d []interface{}) []edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg829 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg829, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfg829
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5830(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5830(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5830 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalMessageDigestCfgMd5830
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectInterfaceTrunkIpOspfOspfGlobalNetwork831(d []interface{}) edpt.InterfaceTrunkIpOspfOspfGlobalNetwork831 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpOspfOspfGlobalNetwork831
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

func getSliceInterfaceTrunkIpOspfOspfIpList(d []interface{}) []edpt.InterfaceTrunkIpOspfOspfIpList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpOspfOspfIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceTrunkIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceTrunkIpOspfOspfIpListMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpOspfOspfIpListMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIpRip832(d []interface{}) edpt.InterfaceTrunkIpRip832 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRip832
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceTrunkIpRipAuthentication833(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceTrunkIpRipSendCfg837(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceTrunkIpRipReceiveCfg838(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceTrunkIpRipSplitHorizonCfg839(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthentication833(d []interface{}) edpt.InterfaceTrunkIpRipAuthentication833 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthentication833
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceTrunkIpRipAuthenticationStr834(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceTrunkIpRipAuthenticationMode835(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceTrunkIpRipAuthenticationKeyChain836(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationStr834(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationStr834 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationStr834
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationMode835(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationMode835 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationMode835
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationKeyChain836(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationKeyChain836 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationKeyChain836
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipSendCfg837(d []interface{}) edpt.InterfaceTrunkIpRipSendCfg837 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipSendCfg837
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipReceiveCfg838(d []interface{}) edpt.InterfaceTrunkIpRipReceiveCfg838 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipReceiveCfg838
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipSplitHorizonCfg839(d []interface{}) edpt.InterfaceTrunkIpRipSplitHorizonCfg839 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipSplitHorizonCfg839
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRouter840(d []interface{}) edpt.InterfaceTrunkIpRouter840 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRouter840
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceTrunkIpRouterIsis841(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpRouterIsis841(d []interface{}) edpt.InterfaceTrunkIpRouterIsis841 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRouterIsis841
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTrunkIpStatefulFirewall842(d []interface{}) edpt.InterfaceTrunkIpStatefulFirewall842 {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpStatefulFirewall842
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

func dataToEndpointInterfaceTrunkIp(d *schema.ResourceData) edpt.InterfaceTrunkIp {
	var ret edpt.InterfaceTrunkIp
	ret.Inst.AddressList = getSliceInterfaceTrunkIpAddressList(d.Get("address_list").([]interface{}))
	ret.Inst.AllowPromiscuousVip = d.Get("allow_promiscuous_vip").(int)
	ret.Inst.CacheSpoofingPort = d.Get("cache_spoofing_port").(int)
	ret.Inst.Client = d.Get("client").(int)
	ret.Inst.Dhcp = d.Get("dhcp").(int)
	ret.Inst.Dmz = d.Get("dmz").(int)
	ret.Inst.GenerateMembershipQuery = d.Get("generate_membership_query").(int)
	ret.Inst.HelperAddressList = getSliceInterfaceTrunkIpHelperAddressList(d.Get("helper_address_list").([]interface{}))
	ret.Inst.MaxRespTime = d.Get("max_resp_time").(int)
	ret.Inst.Nat = getObjectInterfaceTrunkIpNat(d.Get("nat").([]interface{}))
	ret.Inst.Ospf = getObjectInterfaceTrunkIpOspf824(d.Get("ospf").([]interface{}))
	ret.Inst.QueryInterval = d.Get("query_interval").(int)
	ret.Inst.Rip = getObjectInterfaceTrunkIpRip832(d.Get("rip").([]interface{}))
	ret.Inst.Router = getObjectInterfaceTrunkIpRouter840(d.Get("router").([]interface{}))
	ret.Inst.Server = d.Get("server").(int)
	ret.Inst.SlbPartitionRedirect = d.Get("slb_partition_redirect").(int)
	ret.Inst.StatefulFirewall = getObjectInterfaceTrunkIpStatefulFirewall842(d.Get("stateful_firewall").([]interface{}))
	ret.Inst.SynCookie = d.Get("syn_cookie").(int)
	ret.Inst.TtlIgnore = d.Get("ttl_ignore").(int)
	ret.Inst.Unnumbered = d.Get("unnumbered").(int)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
