package thunder

//Thunder resource InterfaceEthernetIp

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceInterfaceEthernetIp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceEthernetIpCreate,
		UpdateContext: resourceInterfaceEthernetIpUpdate,
		ReadContext:   resourceInterfaceEthernetIpRead,
		DeleteContext: resourceInterfaceEthernetIpDelete,
		Schema: map[string]*schema.Schema{
			"dhcp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv4_netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"allow_promiscuous_vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cache_spoofing_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"helper_address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"helper_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"inside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"outside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ttl_ignore": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"syn_cookie": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_partition_redirect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"generate_membership_query": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"query_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_resp_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stateful_firewall": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"class_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"access_list": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"router": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"isis": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"rip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mode": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"key_chain": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_chain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"send_packet": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"receive_packet": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"send_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"send": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"version": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"receive_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"receive": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"version": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"split_horizon_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ospf": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf_global": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"authentication_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"bfd_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bfd": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"disable": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"database_filter_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"database_filter": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"out": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"dead_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"message_digest_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message_digest_key": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"md5_value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mtu": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"mtu_ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"network": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"broadcast": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"non_broadcast": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"point_to_point": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"point_to_multipoint": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"p2mp_nbma": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf_ip_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"authentication": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"authentication_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"database_filter": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"out": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dead_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"message_digest_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"message_digest_key": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"md5_value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mtu_ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"ifnum": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceEthernetIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernetIp (Inside resourceInterfaceEthernetIpCreate) ")
		name1 := d.Get("ifnum").(string)
		data := dataToInterfaceEthernetIp(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceEthernetIp --")
		d.SetId(name1)
		err := go_thunder.PostInterfaceEthernetIp(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceEthernetIpRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceEthernetIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceEthernetIp (Inside resourceInterfaceEthernetIpRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceEthernetIp(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceInterfaceEthernetIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceEthernetIpRead(ctx, d, meta)
}

func resourceInterfaceEthernetIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceEthernetIpRead(ctx, d, meta)
}
func dataToInterfaceEthernetIp(d *schema.ResourceData) go_thunder.InterfaceEthernetIp {
	var vc go_thunder.InterfaceEthernetIp
	var c go_thunder.InterfaceEthernetIPInstance
	c.InterfaceEthernetIPInstanceDhcp = d.Get("dhcp").(int)

	InterfaceEthernetIPInstanceAddressListCount := d.Get("address_list.#").(int)
	c.InterfaceEthernetIPInstanceAddressListIpv4Address = make([]go_thunder.InterfaceEthernetIPInstanceAddressList, 0, InterfaceEthernetIPInstanceAddressListCount)

	for i := 0; i < InterfaceEthernetIPInstanceAddressListCount; i++ {
		var obj1 go_thunder.InterfaceEthernetIPInstanceAddressList
		prefix1 := fmt.Sprintf("address_list.%d.", i)
		obj1.InterfaceEthernetIPInstanceAddressListIpv4Address = d.Get(prefix1 + "ipv4_address").(string)
		obj1.InterfaceEthernetIPInstanceAddressListIpv4Netmask = d.Get(prefix1 + "ipv4_netmask").(string)
		c.InterfaceEthernetIPInstanceAddressListIpv4Address = append(c.InterfaceEthernetIPInstanceAddressListIpv4Address, obj1)
	}

	c.InterfaceEthernetIPInstanceAllowPromiscuousVip = d.Get("allow_promiscuous_vip").(int)
	c.InterfaceEthernetIPInstanceCacheSpoofingPort = d.Get("cache_spoofing_port").(int)

	InterfaceEthernetIPInstanceHelperAddressListCount := d.Get("helper_address_list.#").(int)
	c.InterfaceEthernetIPInstanceHelperAddressListHelperAddress = make([]go_thunder.InterfaceEthernetIPInstanceHelperAddressList, 0, InterfaceEthernetIPInstanceHelperAddressListCount)

	for i := 0; i < InterfaceEthernetIPInstanceHelperAddressListCount; i++ {
		var obj2 go_thunder.InterfaceEthernetIPInstanceHelperAddressList
		prefix2 := fmt.Sprintf("helper_address_list.%d.", i)
		obj2.InterfaceEthernetIPInstanceHelperAddressListHelperAddress = d.Get(prefix2 + "helper_address").(string)
		c.InterfaceEthernetIPInstanceHelperAddressListHelperAddress = append(c.InterfaceEthernetIPInstanceHelperAddressListHelperAddress, obj2)
	}

	c.InterfaceEthernetIPInstanceInside = d.Get("inside").(int)
	c.InterfaceEthernetIPInstanceOutside = d.Get("outside").(int)
	c.InterfaceEthernetIPInstanceTTLIgnore = d.Get("ttl_ignore").(int)
	c.InterfaceEthernetIPInstanceSynCookie = d.Get("syn_cookie").(int)
	c.InterfaceEthernetIPInstanceSlbPartitionRedirect = d.Get("slb_partition_redirect").(int)
	c.InterfaceEthernetIPInstanceGenerateMembershipQuery = d.Get("generate_membership_query").(int)
	c.InterfaceEthernetIPInstanceQueryInterval = d.Get("query_interval").(int)
	c.InterfaceEthernetIPInstanceMaxRespTime = d.Get("max_resp_time").(int)
	c.InterfaceEthernetIPInstanceClient = d.Get("client").(int)
	c.InterfaceEthernetIPInstanceServer = d.Get("server").(int)

	var obj3 go_thunder.InterfaceEthernetIPInstanceStatefulFirewall
	prefix3 := "stateful_firewall.0."
	obj3.InterfaceEthernetIPInstanceStatefulFirewallInside = d.Get(prefix3 + "inside").(int)
	obj3.InterfaceEthernetIPInstanceStatefulFirewallClassList = d.Get(prefix3 + "class_list").(string)
	obj3.InterfaceEthernetIPInstanceStatefulFirewallOutside = d.Get(prefix3 + "outside").(int)
	obj3.InterfaceEthernetIPInstanceStatefulFirewallAccessList = d.Get(prefix3 + "access_list").(int)
	obj3.InterfaceEthernetIPInstanceStatefulFirewallAclID = d.Get(prefix3 + "acl_id").(int)

	c.InterfaceEthernetIPInstanceStatefulFirewallInside = obj3

	var obj4 go_thunder.InterfaceEthernetIPInstanceRouter
	prefix4 := "router.0."

	var obj4_1 go_thunder.InterfaceEthernetIPInstanceRouterIsis
	prefix4_1 := prefix4 + "isis.0."
	obj4_1.InterfaceEthernetIPInstanceRouterIsisTag = d.Get(prefix4_1 + "tag").(string)

	obj4.InterfaceEthernetIPInstanceRouterIsisTag = obj4_1

	c.InterfaceEthernetIPInstanceRouterIsis = obj4

	var obj5 go_thunder.InterfaceEthernetIPInstanceRip
	prefix5 := "rip.0."

	var obj5_1 go_thunder.InterfaceEthernetIPInstanceRipAuthentication
	prefix5_1 := prefix5 + "authentication.0."

	var obj5_1_1 go_thunder.InterfaceEthernetIPInstanceRipAuthenticationStr
	prefix5_1_1 := prefix5_1 + "str.0."
	obj5_1_1.InterfaceEthernetIPInstanceRipAuthenticationStrString = d.Get(prefix5_1_1 + "string").(string)

	obj5_1.InterfaceEthernetIPInstanceRipAuthenticationStrString = obj5_1_1

	var obj5_1_2 go_thunder.InterfaceEthernetIPInstanceRipAuthenticationMode
	prefix5_1_2 := prefix5_1 + "mode.0."
	obj5_1_2.InterfaceEthernetIPInstanceRipAuthenticationModeMode = d.Get(prefix5_1_2 + "mode").(string)

	obj5_1.InterfaceEthernetIPInstanceRipAuthenticationModeMode = obj5_1_2

	var obj5_1_3 go_thunder.InterfaceEthernetIPInstanceRipAuthenticationKeyChain
	prefix5_1_3 := prefix5_1 + "key_chain.0."
	obj5_1_3.InterfaceEthernetIPInstanceRipAuthenticationKeyChainKeyChain = d.Get(prefix5_1_3 + "key_chain").(string)

	obj5_1.InterfaceEthernetIPInstanceRipAuthenticationKeyChainKeyChain = obj5_1_3

	obj5.InterfaceEthernetIPInstanceRipAuthenticationStr = obj5_1

	obj5.InterfaceEthernetIPInstanceRipSendPacket = d.Get(prefix5 + "send_packet").(int)
	obj5.InterfaceEthernetIPInstanceRipReceivePacket = d.Get(prefix5 + "receive_packet").(int)

	var obj5_2 go_thunder.InterfaceEthernetIPInstanceRipSendCfg
	prefix5_2 := prefix5 + "send_cfg.0."
	obj5_2.InterfaceEthernetIPInstanceRipSendCfgSend = d.Get(prefix5_2 + "send").(int)
	obj5_2.InterfaceEthernetIPInstanceRipSendCfgVersion = d.Get(prefix5_2 + "version").(string)

	obj5.InterfaceEthernetIPInstanceRipSendCfgSend = obj5_2

	var obj5_3 go_thunder.InterfaceEthernetIPInstanceRipReceiveCfg
	prefix5_3 := prefix5 + "receive_cfg.0."
	obj5_3.InterfaceEthernetIPInstanceRipReceiveCfgReceive = d.Get(prefix5_3 + "receive").(int)
	obj5_3.InterfaceEthernetIPInstanceRipReceiveCfgVersion = d.Get(prefix5_3 + "version").(string)

	obj5.InterfaceEthernetIPInstanceRipReceiveCfgReceive = obj5_3

	var obj5_4 go_thunder.InterfaceEthernetIPInstanceRipSplitHorizonCfg
	prefix5_4 := prefix5 + "split_horizon_cfg.0."
	obj5_4.InterfaceEthernetIPInstanceRipSplitHorizonCfgState = d.Get(prefix5_4 + "state").(string)

	obj5.InterfaceEthernetIPInstanceRipSplitHorizonCfgState = obj5_4

	c.InterfaceEthernetIPInstanceRipAuthentication = obj5

	var obj6 go_thunder.InterfaceEthernetIPInstanceOspf
	prefix6 := "ospf.0."

	var obj6_1 go_thunder.InterfaceEthernetIPInstanceOspfOspfGlobal
	prefix6_1 := prefix6 + "ospf_global.0."

	var obj6_1_1 go_thunder.InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfg
	prefix6_1_1 := prefix6_1 + "authentication_cfg.0."
	obj6_1_1.InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication = d.Get(prefix6_1_1 + "authentication").(int)
	obj6_1_1.InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfgValue = d.Get(prefix6_1_1 + "value").(string)

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication = obj6_1_1

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationKey = d.Get(prefix6_1 + "authentication_key").(string)

	var obj6_1_2 go_thunder.InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfg
	prefix6_1_2 := prefix6_1 + "bfd_cfg.0."
	obj6_1_2.InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfgBfd = d.Get(prefix6_1_2 + "bfd").(int)
	obj6_1_2.InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfgDisable = d.Get(prefix6_1_2 + "disable").(int)

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalBfdCfgBfd = obj6_1_2

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalCost = d.Get(prefix6_1 + "cost").(int)

	var obj6_1_3 go_thunder.InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfg
	prefix6_1_3 := prefix6_1 + "database_filter_cfg.0."
	obj6_1_3.InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = d.Get(prefix6_1_3 + "database_filter").(string)
	obj6_1_3.InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfgOut = d.Get(prefix6_1_3 + "out").(int)

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = obj6_1_3

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalDeadInterval = d.Get(prefix6_1 + "dead_interval").(int)
	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalDisable = d.Get(prefix6_1 + "disable").(string)
	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalHelloInterval = d.Get(prefix6_1 + "hello_interval").(int)

	InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgCount := d.Get(prefix6_1 + "message_digest_cfg.#").(int)
	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfg, 0, InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgCount)

	for i := 0; i < InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgCount; i++ {
		var obj6_1_4 go_thunder.InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfg
		prefix6_1_4 := prefix6_1 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj6_1_4.InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey = d.Get(prefix6_1_4 + "message_digest_key").(int)
		obj6_1_4.InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMd5Value = d.Get(prefix6_1_4 + "md5_value").(string)
		obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey = append(obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey, obj6_1_4)
	}

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalMtu = d.Get(prefix6_1 + "mtu").(int)
	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalMtuIgnore = d.Get(prefix6_1 + "mtu_ignore").(int)

	var obj6_1_5 go_thunder.InterfaceEthernetIPInstanceOspfOspfGlobalNetwork
	prefix6_1_5 := prefix6_1 + "network.0."
	obj6_1_5.InterfaceEthernetIPInstanceOspfOspfGlobalNetworkBroadcast = d.Get(prefix6_1_5 + "broadcast").(int)
	obj6_1_5.InterfaceEthernetIPInstanceOspfOspfGlobalNetworkNonBroadcast = d.Get(prefix6_1_5 + "non_broadcast").(int)
	obj6_1_5.InterfaceEthernetIPInstanceOspfOspfGlobalNetworkPointToPoint = d.Get(prefix6_1_5 + "point_to_point").(int)
	obj6_1_5.InterfaceEthernetIPInstanceOspfOspfGlobalNetworkPointToMultipoint = d.Get(prefix6_1_5 + "point_to_multipoint").(int)
	obj6_1_5.InterfaceEthernetIPInstanceOspfOspfGlobalNetworkP2MpNbma = d.Get(prefix6_1_5 + "p2mp_nbma").(int)

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalNetworkBroadcast = obj6_1_5

	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalPriority = d.Get(prefix6_1 + "priority").(int)
	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalRetransmitInterval = d.Get(prefix6_1 + "retransmit_interval").(int)
	obj6_1.InterfaceEthernetIPInstanceOspfOspfGlobalTransmitDelay = d.Get(prefix6_1 + "transmit_delay").(int)

	obj6.InterfaceEthernetIPInstanceOspfOspfGlobalAuthenticationCfg = obj6_1

	InterfaceEthernetIPInstanceOspfOspfIPListCount := d.Get(prefix6 + "ospf_ip_list.#").(int)
	obj6.InterfaceEthernetIPInstanceOspfOspfIPListIPAddr = make([]go_thunder.InterfaceEthernetIPInstanceOspfOspfIPList, 0, InterfaceEthernetIPInstanceOspfOspfIPListCount)

	for i := 0; i < InterfaceEthernetIPInstanceOspfOspfIPListCount; i++ {
		var obj6_2 go_thunder.InterfaceEthernetIPInstanceOspfOspfIPList
		prefix6_2 := prefix6 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListIPAddr = d.Get(prefix6_2 + "ip_addr").(string)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListAuthentication = d.Get(prefix6_2 + "authentication").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListValue = d.Get(prefix6_2 + "value").(string)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListAuthenticationKey = d.Get(prefix6_2 + "authentication_key").(string)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListCost = d.Get(prefix6_2 + "cost").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListDatabaseFilter = d.Get(prefix6_2 + "database_filter").(string)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListOut = d.Get(prefix6_2 + "out").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListDeadInterval = d.Get(prefix6_2 + "dead_interval").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListHelloInterval = d.Get(prefix6_2 + "hello_interval").(int)

		InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgCount := d.Get(prefix6_2 + "message_digest_cfg.#").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfg, 0, InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgCount)

		for i := 0; i < InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgCount; i++ {
			var obj6_2_1 go_thunder.InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfg
			prefix6_2_1 := prefix6_2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj6_2_1.InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey = d.Get(prefix6_2_1 + "message_digest_key").(int)
			obj6_2_1.InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMd5Value = d.Get(prefix6_2_1 + "md5_value").(string)
			obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey = append(obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey, obj6_2_1)
		}

		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListMtuIgnore = d.Get(prefix6_2 + "mtu_ignore").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListPriority = d.Get(prefix6_2 + "priority").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListRetransmitInterval = d.Get(prefix6_2 + "retransmit_interval").(int)
		obj6_2.InterfaceEthernetIPInstanceOspfOspfIPListTransmitDelay = d.Get(prefix6_2 + "transmit_delay").(int)
		obj6.InterfaceEthernetIPInstanceOspfOspfIPListIPAddr = append(obj6.InterfaceEthernetIPInstanceOspfOspfIPListIPAddr, obj6_2)
	}

	c.InterfaceEthernetIPInstanceOspfOspfGlobal = obj6

	vc.InterfaceEthernetIPInstanceDhcp = c
	return vc
}
