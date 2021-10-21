package thunder

//Thunder resource InterfaceVeIp

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceInterfaceVeIp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceVeIpCreate,
		UpdateContext: resourceInterfaceVeIpUpdate,
		ReadContext:   resourceInterfaceVeIpRead,
		DeleteContext: resourceInterfaceVeIpDelete,
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

func resourceInterfaceVeIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceVeIp (Inside resourceInterfaceVeIpCreate) ")
		name1 := d.Get("ifnum").(string)
		data := dataToInterfaceVeIp(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceVeIp --")
		d.SetId(name1)
		err := go_thunder.PostInterfaceVeIp(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceVeIpRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceVeIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceVeIp (Inside resourceInterfaceVeIpRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceVeIp(client.Token, name1, client.Host)
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

func resourceInterfaceVeIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceVeIpRead(ctx, d, meta)
}

func resourceInterfaceVeIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceVeIpRead(ctx, d, meta)
}
func dataToInterfaceVeIp(d *schema.ResourceData) go_thunder.InterfaceVeIp {
	var vc go_thunder.InterfaceVeIp
	var c go_thunder.InterfaceVeIPInstance
	c.InterfaceVeIPInstanceDhcp = d.Get("dhcp").(int)

	InterfaceVeIPInstanceAddressListCount := d.Get("address_list.#").(int)
	c.InterfaceVeIPInstanceAddressListIpv4Address = make([]go_thunder.InterfaceVeIPInstanceAddressList, 0, InterfaceVeIPInstanceAddressListCount)

	for i := 0; i < InterfaceVeIPInstanceAddressListCount; i++ {
		var obj1 go_thunder.InterfaceVeIPInstanceAddressList
		prefix1 := fmt.Sprintf("address_list.%d.", i)
		obj1.InterfaceVeIPInstanceAddressListIpv4Address = d.Get(prefix1 + "ipv4_address").(string)
		obj1.InterfaceVeIPInstanceAddressListIpv4Netmask = d.Get(prefix1 + "ipv4_netmask").(string)
		c.InterfaceVeIPInstanceAddressListIpv4Address = append(c.InterfaceVeIPInstanceAddressListIpv4Address, obj1)
	}

	c.InterfaceVeIPInstanceAllowPromiscuousVip = d.Get("allow_promiscuous_vip").(int)
	c.InterfaceVeIPInstanceClient = d.Get("client").(int)
	c.InterfaceVeIPInstanceServer = d.Get("server").(int)

	InterfaceVeIPInstanceHelperAddressListCount := d.Get("helper_address_list.#").(int)
	c.InterfaceVeIPInstanceHelperAddressListHelperAddress = make([]go_thunder.InterfaceVeIPInstanceHelperAddressList, 0, InterfaceVeIPInstanceHelperAddressListCount)

	for i := 0; i < InterfaceVeIPInstanceHelperAddressListCount; i++ {
		var obj2 go_thunder.InterfaceVeIPInstanceHelperAddressList
		prefix2 := fmt.Sprintf("helper_address_list.%d.", i)
		obj2.InterfaceVeIPInstanceHelperAddressListHelperAddress = d.Get(prefix2 + "helper_address").(string)
		c.InterfaceVeIPInstanceHelperAddressListHelperAddress = append(c.InterfaceVeIPInstanceHelperAddressListHelperAddress, obj2)
	}

	c.InterfaceVeIPInstanceInside = d.Get("inside").(int)
	c.InterfaceVeIPInstanceOutside = d.Get("outside").(int)
	c.InterfaceVeIPInstanceTTLIgnore = d.Get("ttl_ignore").(int)
	c.InterfaceVeIPInstanceSynCookie = d.Get("syn_cookie").(int)
	c.InterfaceVeIPInstanceSlbPartitionRedirect = d.Get("slb_partition_redirect").(int)
	c.InterfaceVeIPInstanceGenerateMembershipQuery = d.Get("generate_membership_query").(int)
	c.InterfaceVeIPInstanceQueryInterval = d.Get("query_interval").(int)
	c.InterfaceVeIPInstanceMaxRespTime = d.Get("max_resp_time").(int)

	var obj3 go_thunder.InterfaceVeIPInstanceStatefulFirewall
	prefix3 := "stateful_firewall.0."
	obj3.InterfaceVeIPInstanceStatefulFirewallInside = d.Get(prefix3 + "inside").(int)
	obj3.InterfaceVeIPInstanceStatefulFirewallClassList = d.Get(prefix3 + "class_list").(string)
	obj3.InterfaceVeIPInstanceStatefulFirewallOutside = d.Get(prefix3 + "outside").(int)
	obj3.InterfaceVeIPInstanceStatefulFirewallAccessList = d.Get(prefix3 + "access_list").(int)
	obj3.InterfaceVeIPInstanceStatefulFirewallAclID = d.Get(prefix3 + "acl_id").(int)

	c.InterfaceVeIPInstanceStatefulFirewallInside = obj3

	var obj4 go_thunder.InterfaceVeIPInstanceRouter
	prefix4 := "router.0."

	var obj4_1 go_thunder.InterfaceVeIPInstanceRouterIsis
	prefix4_1 := prefix4 + "isis.0."
	obj4_1.InterfaceVeIPInstanceRouterIsisTag = d.Get(prefix4_1 + "tag").(string)

	obj4.InterfaceVeIPInstanceRouterIsisTag = obj4_1

	c.InterfaceVeIPInstanceRouterIsis = obj4

	var obj5 go_thunder.InterfaceVeIPInstanceRip
	prefix5 := "rip.0."

	var obj5_1 go_thunder.InterfaceVeIPInstanceRipAuthentication
	prefix5_1 := prefix5 + "authentication.0."

	var obj5_1_1 go_thunder.InterfaceVeIPInstanceRipAuthenticationStr
	prefix5_1_1 := prefix5_1 + "str.0."
	obj5_1_1.InterfaceVeIPInstanceRipAuthenticationStrString = d.Get(prefix5_1_1 + "string").(string)

	obj5_1.InterfaceVeIPInstanceRipAuthenticationStrString = obj5_1_1

	var obj5_1_2 go_thunder.InterfaceVeIPInstanceRipAuthenticationMode
	prefix5_1_2 := prefix5_1 + "mode.0."
	obj5_1_2.InterfaceVeIPInstanceRipAuthenticationModeMode = d.Get(prefix5_1_2 + "mode").(string)

	obj5_1.InterfaceVeIPInstanceRipAuthenticationModeMode = obj5_1_2

	var obj5_1_3 go_thunder.InterfaceVeIPInstanceRipAuthenticationKeyChain
	prefix5_1_3 := prefix5_1 + "key_chain.0."
	obj5_1_3.InterfaceVeIPInstanceRipAuthenticationKeyChainKeyChain = d.Get(prefix5_1_3 + "key_chain").(string)

	obj5_1.InterfaceVeIPInstanceRipAuthenticationKeyChainKeyChain = obj5_1_3

	obj5.InterfaceVeIPInstanceRipAuthenticationStr = obj5_1

	obj5.InterfaceVeIPInstanceRipSendPacket = d.Get(prefix5 + "send_packet").(int)
	obj5.InterfaceVeIPInstanceRipReceivePacket = d.Get(prefix5 + "receive_packet").(int)

	var obj5_2 go_thunder.InterfaceVeIPInstanceRipSendCfg
	prefix5_2 := prefix5 + "send_cfg.0."
	obj5_2.InterfaceVeIPInstanceRipSendCfgSend = d.Get(prefix5_2 + "send").(int)
	obj5_2.InterfaceVeIPInstanceRipSendCfgVersion = d.Get(prefix5_2 + "version").(string)

	obj5.InterfaceVeIPInstanceRipSendCfgSend = obj5_2

	var obj5_3 go_thunder.InterfaceVeIPInstanceRipReceiveCfg
	prefix5_3 := prefix5 + "receive_cfg.0."
	obj5_3.InterfaceVeIPInstanceRipReceiveCfgReceive = d.Get(prefix5_3 + "receive").(int)
	obj5_3.InterfaceVeIPInstanceRipReceiveCfgVersion = d.Get(prefix5_3 + "version").(string)

	obj5.InterfaceVeIPInstanceRipReceiveCfgReceive = obj5_3

	var obj5_4 go_thunder.InterfaceVeIPInstanceRipSplitHorizonCfg
	prefix5_4 := prefix5 + "split_horizon_cfg.0."
	obj5_4.InterfaceVeIPInstanceRipSplitHorizonCfgState = d.Get(prefix5_4 + "state").(string)

	obj5.InterfaceVeIPInstanceRipSplitHorizonCfgState = obj5_4

	c.InterfaceVeIPInstanceRipAuthentication = obj5

	var obj6 go_thunder.InterfaceVeIPInstanceOspf
	prefix6 := "ospf.0."

	var obj6_1 go_thunder.InterfaceVeIPInstanceOspfOspfGlobal
	prefix6_1 := prefix6 + "ospf_global.0."

	var obj6_1_1 go_thunder.InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfg
	prefix6_1_1 := prefix6_1 + "authentication_cfg.0."
	obj6_1_1.InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication = d.Get(prefix6_1_1 + "authentication").(int)
	obj6_1_1.InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfgValue = d.Get(prefix6_1_1 + "value").(string)

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfgAuthentication = obj6_1_1

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalAuthenticationKey = d.Get(prefix6_1 + "authentication_key").(string)

	var obj6_1_2 go_thunder.InterfaceVeIPInstanceOspfOspfGlobalBfdCfg
	prefix6_1_2 := prefix6_1 + "bfd_cfg.0."
	obj6_1_2.InterfaceVeIPInstanceOspfOspfGlobalBfdCfgBfd = d.Get(prefix6_1_2 + "bfd").(int)
	obj6_1_2.InterfaceVeIPInstanceOspfOspfGlobalBfdCfgDisable = d.Get(prefix6_1_2 + "disable").(int)

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalBfdCfgBfd = obj6_1_2

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalCost = d.Get(prefix6_1 + "cost").(int)

	var obj6_1_3 go_thunder.InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfg
	prefix6_1_3 := prefix6_1 + "database_filter_cfg.0."
	obj6_1_3.InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = d.Get(prefix6_1_3 + "database_filter").(string)
	obj6_1_3.InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfgOut = d.Get(prefix6_1_3 + "out").(int)

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = obj6_1_3

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalDeadInterval = d.Get(prefix6_1 + "dead_interval").(int)
	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalDisable = d.Get(prefix6_1 + "disable").(string)
	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalHelloInterval = d.Get(prefix6_1 + "hello_interval").(int)

	InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgCount := d.Get(prefix6_1 + "message_digest_cfg.#").(int)
	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfg, 0, InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgCount)

	for i := 0; i < InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgCount; i++ {
		var obj6_1_4 go_thunder.InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfg
		prefix6_1_4 := prefix6_1 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj6_1_4.InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey = d.Get(prefix6_1_4 + "message_digest_key").(int)
		obj6_1_4.InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMd5Value = d.Get(prefix6_1_4 + "md5_value").(string)
		obj6_1.InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey = append(obj6_1.InterfaceVeIPInstanceOspfOspfGlobalMessageDigestCfgMessageDigestKey, obj6_1_4)
	}

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalMtu = d.Get(prefix6_1 + "mtu").(int)
	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalMtuIgnore = d.Get(prefix6_1 + "mtu_ignore").(int)

	var obj6_1_5 go_thunder.InterfaceVeIPInstanceOspfOspfGlobalNetwork
	prefix6_1_5 := prefix6_1 + "network.0."
	obj6_1_5.InterfaceVeIPInstanceOspfOspfGlobalNetworkBroadcast = d.Get(prefix6_1_5 + "broadcast").(int)
	obj6_1_5.InterfaceVeIPInstanceOspfOspfGlobalNetworkNonBroadcast = d.Get(prefix6_1_5 + "non_broadcast").(int)
	obj6_1_5.InterfaceVeIPInstanceOspfOspfGlobalNetworkPointToPoint = d.Get(prefix6_1_5 + "point_to_point").(int)
	obj6_1_5.InterfaceVeIPInstanceOspfOspfGlobalNetworkPointToMultipoint = d.Get(prefix6_1_5 + "point_to_multipoint").(int)
	obj6_1_5.InterfaceVeIPInstanceOspfOspfGlobalNetworkP2MpNbma = d.Get(prefix6_1_5 + "p2mp_nbma").(int)

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalNetworkBroadcast = obj6_1_5

	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalPriority = d.Get(prefix6_1 + "priority").(int)
	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalRetransmitInterval = d.Get(prefix6_1 + "retransmit_interval").(int)
	obj6_1.InterfaceVeIPInstanceOspfOspfGlobalTransmitDelay = d.Get(prefix6_1 + "transmit_delay").(int)

	obj6.InterfaceVeIPInstanceOspfOspfGlobalAuthenticationCfg = obj6_1

	InterfaceVeIPInstanceOspfOspfIPListCount := d.Get(prefix6 + "ospf_ip_list.#").(int)
	obj6.InterfaceVeIPInstanceOspfOspfIPListIPAddr = make([]go_thunder.InterfaceVeIPInstanceOspfOspfIPList, 0, InterfaceVeIPInstanceOspfOspfIPListCount)

	for i := 0; i < InterfaceVeIPInstanceOspfOspfIPListCount; i++ {
		var obj6_2 go_thunder.InterfaceVeIPInstanceOspfOspfIPList
		prefix6_2 := prefix6 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListIPAddr = d.Get(prefix6_2 + "ip_addr").(string)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListAuthentication = d.Get(prefix6_2 + "authentication").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListValue = d.Get(prefix6_2 + "value").(string)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListAuthenticationKey = d.Get(prefix6_2 + "authentication_key").(string)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListCost = d.Get(prefix6_2 + "cost").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListDatabaseFilter = d.Get(prefix6_2 + "database_filter").(string)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListOut = d.Get(prefix6_2 + "out").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListDeadInterval = d.Get(prefix6_2 + "dead_interval").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListHelloInterval = d.Get(prefix6_2 + "hello_interval").(int)

		InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgCount := d.Get(prefix6_2 + "message_digest_cfg.#").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfg, 0, InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgCount)

		for i := 0; i < InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgCount; i++ {
			var obj6_2_1 go_thunder.InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfg
			prefix6_2_1 := prefix6_2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj6_2_1.InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey = d.Get(prefix6_2_1 + "message_digest_key").(int)
			obj6_2_1.InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMd5Value = d.Get(prefix6_2_1 + "md5_value").(string)
			obj6_2.InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey = append(obj6_2.InterfaceVeIPInstanceOspfOspfIPListMessageDigestCfgMessageDigestKey, obj6_2_1)
		}

		obj6_2.InterfaceVeIPInstanceOspfOspfIPListMtuIgnore = d.Get(prefix6_2 + "mtu_ignore").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListPriority = d.Get(prefix6_2 + "priority").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListRetransmitInterval = d.Get(prefix6_2 + "retransmit_interval").(int)
		obj6_2.InterfaceVeIPInstanceOspfOspfIPListTransmitDelay = d.Get(prefix6_2 + "transmit_delay").(int)
		obj6.InterfaceVeIPInstanceOspfOspfIPListIPAddr = append(obj6.InterfaceVeIPInstanceOspfOspfIPListIPAddr, obj6_2)
	}

	c.InterfaceVeIPInstanceOspfOspfGlobal = obj6

	vc.InterfaceVeIPInstanceDhcp = c
	return vc
}
