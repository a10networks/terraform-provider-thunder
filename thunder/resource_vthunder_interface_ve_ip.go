package thunder

//Thunder resource InterfaceVeIP

import (
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceVeIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceVeIPCreate,
		Update: resourceInterfaceVeIPUpdate,
		Read:   resourceInterfaceVeIPRead,
		Delete: resourceInterfaceVeIPDelete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateful_firewall": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"access_list": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"class_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_partition_redirect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"allow_promiscuous_vip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_resp_time": {
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
			"inside": {
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ospf": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf_ip_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ip_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"mtu_ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"database_filter": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dead_interval": {
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
												"md5": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"encrypted": {
																Type:        schema.TypeString,
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
											},
										},
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"out": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"authentication_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"authentication": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf_global": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"mtu_ignore": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dead_interval": {
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
												"md5": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"encrypted": {
																Type:        schema.TypeString,
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
											},
										},
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
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
												"point_to_multipoint": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"point_to_point": {
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
									"mtu": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"transmit_delay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"authentication_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"authentication": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"authentication_key": {
										Type:        schema.TypeString,
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
								},
							},
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
			"outside": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"receive_packet": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
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
						"send_packet": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
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
						"send_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"send": {
										Type:        schema.TypeInt,
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
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
					},
				},
			},
			"query_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"client": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"generate_membership_query": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dhcp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ttl_ignore": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceVeIPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceVeIP (Inside resourceInterfaceVeIPCreate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceVeIP(d)
		logger.Println("[INFO] received V from method data to InterfaceVeIP --")
		d.SetId(strconv.Itoa(name))
		go_thunder.PostInterfaceVeIP(client.Token, name, data, client.Host)

		return resourceInterfaceVeIPRead(d, meta)

	}
	return nil
}

func resourceInterfaceVeIPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceVeIP (Inside resourceInterfaceVeIPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceVeIP(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceInterfaceVeIPUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceVeIPRead(d, meta)
}

func resourceInterfaceVeIPDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceInterfaceVeIPRead(d, meta)
}

func dataToInterfaceVeIP(d *schema.ResourceData) go_thunder.VeIP {
	var vc go_thunder.VeIP
	var c go_thunder.VeIPInstance
	c.GenerateMembershipQuery = d.Get("generate_membership_query").(int)

	AddressListCount := d.Get("address_list.#").(int)
	c.Ipv4Address = make([]go_thunder.VeIPAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj1 go_thunder.VeIPAddressList
		prefix := fmt.Sprintf("address_list.%d.", i)
		obj1.Ipv4Address = d.Get(prefix + "ipv4_address").(string)
		obj1.Ipv4Netmask = d.Get(prefix + "ipv4_netmask").(string)
		c.Ipv4Address = append(c.Ipv4Address, obj1)
	}

	c.Inside = d.Get("inside").(int)
	c.AllowPromiscuousVip = d.Get("allow_promiscuous_vip").(int)

	HelperAddressListCount := d.Get("helper_address_list.#").(int)
	c.HelperAddress = make([]go_thunder.VeIPHelperAddressList, 0, HelperAddressListCount)

	for i := 0; i < HelperAddressListCount; i++ {
		var obj2 go_thunder.VeIPHelperAddressList
		prefix := fmt.Sprintf("helper_address_list.%d.", i)
		obj2.HelperAddress = d.Get(prefix + "helper_address").(string)
		c.HelperAddress = append(c.HelperAddress, obj2)
	}

	c.MaxRespTime = d.Get("max_resp_time").(int)
	c.QueryInterval = d.Get("query_interval").(int)
	c.Outside = d.Get("outside").(int)
	c.Client = d.Get("client").(int)

	var obj3 go_thunder.VeIPStatefulFirewall
	prefix := "stateful_firewall.0."
	obj3.ClassList = d.Get(prefix + "class_list").(string)
	obj3.Inside = d.Get(prefix + "inside").(int)
	obj3.Outside = d.Get(prefix + "outside").(int)
	obj3.ACLID = d.Get(prefix + "acl_id").(int)
	obj3.AccessList = d.Get(prefix + "access_list").(int)
	c.ClassList = obj3

	var obj4 go_thunder.VeIPRip
	prefix = "rip.0."

	var obj4_1 go_thunder.VeIPReceiveCfg
	prefix1 := prefix + "receive_cfg.0."
	obj4_1.Receive = d.Get(prefix1 + "receive").(int)
	obj4_1.Version = d.Get(prefix1 + "version").(string)
	obj4.Receive = obj4_1

	obj4.ReceivePacket = d.Get(prefix + "receive_packet").(int)

	var obj4_2 go_thunder.VeIPSplitHorizonCfg
	prefix1 = prefix + "split_horizon_cfg.0."
	obj4_2.State = d.Get(prefix1 + "state").(string)
	obj4.State = obj4_2

	var obj4_3 go_thunder.VeIPAuthentication
	prefix1 = prefix + "authentication.0."

	var obj4_3_1 go_thunder.VeIPKeyChain
	prefix2 := prefix1 + "key_chain.0."
	obj4_3_1.KeyChain = d.Get(prefix2 + "key_chain").(string)
	obj4_3.KeyChain = obj4_3_1

	var obj4_3_2 go_thunder.VeIPMode
	prefix2 = prefix1 + "mode.0."
	obj4_3_2.Mode = d.Get(prefix2 + "mode").(string)
	obj4_3.Mode = obj4_3_2

	var obj4_3_3 go_thunder.VeIPStr
	prefix2 = prefix1 + "str.0."
	obj4_3_3.String = d.Get(prefix2 + "string").(string)
	obj4_3.String = obj4_3_3

	obj4.KeyChain = obj4_3

	var obj4_4 go_thunder.VeIPSendCfg
	prefix1 = prefix + "send_cfg.0."
	obj4_4.Version = d.Get(prefix1 + "version").(string)
	obj4_4.Send = d.Get(prefix1 + "send").(int)
	obj4.Version = obj4_4

	obj4.SendPacket = d.Get(prefix + "send_packet").(int)
	c.Receive = obj4

	c.TTLIgnore = d.Get("ttl_ignore").(int)

	var obj5 go_thunder.VeIPRouter
	prefix = "router.0."

	var obj5_1 go_thunder.VeIPIsis
	prefix1 = prefix + "isis.0."
	obj5_1.Tag = d.Get(prefix1 + "tag").(string)
	obj5.Tag = obj5_1

	c.Tag = obj5

	c.Dhcp = d.Get("dhcp").(int)
	c.Server = d.Get("server").(int)

	var obj6 go_thunder.VeIPOspf
	prefix = "ospf.0."

	OspfIpListCount := d.Get(prefix + "ospf_ip_list.#").(int)
	obj6.DeadInterval = make([]go_thunder.VeIPOspfIPList, 0, OspfIpListCount)

	for i := 0; i < OspfIpListCount; i++ {
		var obj6_1 go_thunder.VeIPOspfIPList
		prefix1 = prefix + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj6_1.DeadInterval = d.Get(prefix1 + "dead_interval").(int)
		obj6_1.AuthenticationKey = d.Get(prefix1 + "authentication_key").(string)
		obj6_1.MtuIgnore = d.Get(prefix1 + "mtu_ignore").(int)
		obj6_1.TransmitDelay = d.Get(prefix1 + "transmit_delay").(int)
		obj6_1.Value = d.Get(prefix1 + "value").(string)
		obj6_1.Priority = d.Get(prefix1 + "priority").(int)
		obj6_1.Authentication = d.Get(prefix1 + "authentication").(int)
		obj6_1.Cost = d.Get(prefix1 + "cost").(int)
		obj6_1.DatabaseFilter = d.Get(prefix1 + "database_filter").(string)
		obj6_1.HelloInterval = d.Get(prefix1 + "hello_interval").(int)
		obj6_1.IPAddr = d.Get(prefix1 + "ip_addr").(string)
		obj6_1.RetransmitInterval = d.Get(prefix1 + "retransmit_interval").(int)

		MessageDigestCfgCount := d.Get(prefix1 + "message_digest_cfg.#").(int)
		obj6_1.MessageDigestKey = make([]go_thunder.VeIPMessageDigestCfg, 0, MessageDigestCfgCount)

		for i := 0; i < MessageDigestCfgCount; i++ {
			var obj6_1_1 go_thunder.VeIPMessageDigestCfg
			prefix2 = prefix1 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj6_1_1.MessageDigestKey = d.Get(prefix2 + "message_digest_key").(int)

			var obj6_1_1_1 go_thunder.VeIPMd5
			prefix3 := prefix2 + "md5.0."
			obj6_1_1_1.Md5Value = d.Get(prefix3 + "md5_value").(string)
			obj6_1_1_1.Encrypted = d.Get(prefix3 + "encrypted").(string)
			obj6_1_1.Md5Value = obj6_1_1_1

			obj6_1.MessageDigestKey = append(obj6_1.MessageDigestKey, obj6_1_1)
		}

		obj6_1.Out = d.Get(prefix1 + "out").(int)
		obj6.DeadInterval = append(obj6.DeadInterval, obj6_1)
	}

	var obj6_2 go_thunder.VeIPOspfGlobal
	prefix1 = prefix + "ospf_global.0."
	obj6_2.Cost = d.Get(prefix1 + "cost").(int)
	obj6_2.DeadInterval = d.Get(prefix1 + "dead_interval").(int)
	obj6_2.AuthenticationKey = d.Get(prefix1 + "authentication_key").(string)

	var obj6_2_1 go_thunder.VeIPNetwork
	prefix2 = prefix1 + "network.0."
	obj6_2_1.Broadcast = d.Get(prefix2 + "broadcast").(int)
	obj6_2_1.PointToMultipoint = d.Get(prefix2 + "point_to_multipoint").(int)
	obj6_2_1.NonBroadcast = d.Get(prefix2 + "non_broadcast").(int)
	obj6_2_1.PointToPoint = d.Get(prefix2 + "point_to_point").(int)
	obj6_2_1.P2MpNbma = d.Get(prefix2 + "p2mp_nbma").(int)
	obj6_2.Broadcast = obj6_2_1

	obj6_2.MtuIgnore = d.Get(prefix1 + "mtu_ignore").(int)
	obj6_2.TransmitDelay = d.Get(prefix1 + "transmit_delay").(int)

	var obj6_2_2 go_thunder.VeIPAuthenticationCfg
	prefix2 = prefix1 + "authentication_cfg.0."
	obj6_2_2.Authentication = d.Get(prefix2 + "authentication").(int)
	obj6_2_2.Value = d.Get(prefix2 + "value").(string)
	obj6_2.Authentication = obj6_2_2

	obj6_2.RetransmitInterval = d.Get(prefix1 + "retransmit_interval").(int)

	var obj6_2_3 go_thunder.VeIPBfdCfg
	prefix2 = prefix1 + "bfd_cfg.0."
	obj6_2_3.Disable = d.Get(prefix2 + "disable").(int)
	obj6_2_3.Bfd = d.Get(prefix2 + "bfd").(int)
	obj6_2.Bfd = obj6_2_3

	obj6_2.Disable = d.Get(prefix1 + "disable").(string)
	obj6_2.HelloInterval = d.Get(prefix1 + "hello_interval").(int)

	var obj6_2_4 go_thunder.VeIPDatabaseFilterCfg
	prefix2 = prefix1 + "database_filter_cfg.0."
	obj6_2_4.DatabaseFilter = d.Get(prefix2 + "database_filter").(string)
	obj6_2_4.Out = d.Get(prefix2 + "out").(int)
	obj6_2.DatabaseFilter = obj6_2_4

	obj6_2.Priority = d.Get(prefix1 + "priority").(int)
	obj6_2.Mtu = d.Get(prefix1 + "mtu").(int)

	MessageDigestCfgCount := d.Get(prefix1 + "message_digest_cfg.#").(int)
	obj6_2.MessageDigestKey = make([]go_thunder.VeIPMessageDigestCfg, 0, MessageDigestCfgCount)

	for i := 0; i < MessageDigestCfgCount; i++ {
		var obj5 go_thunder.VeIPMessageDigestCfg
		prefix2 = prefix1 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj5.MessageDigestKey = d.Get(prefix2 + "message_digest_key").(int)

		var obj1 go_thunder.VeIPMd5
		prefix3 := prefix2 + "md5.0."
		obj1.Md5Value = d.Get(prefix3 + "md5_value").(string)
		obj1.Encrypted = d.Get(prefix3 + "encrypted").(string)
		obj5.Md5Value = obj1

		obj6_2.MessageDigestKey = append(obj6_2.MessageDigestKey, obj5)
	}

	obj6.Cost = obj6_2

	c.DeadInterval = obj6

	c.SlbPartitionRedirect = d.Get("slb_partition_redirect").(int)

	vc.UUID = c
	return vc
}
