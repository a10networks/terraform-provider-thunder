package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback_ip`: Global IP configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackIpCreate,
		UpdateContext: resourceInterfaceLoopbackIpUpdate,
		ReadContext:   resourceInterfaceLoopbackIpRead,
		DeleteContext: resourceInterfaceLoopbackIpDelete,

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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceLoopbackIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLoopbackIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLoopbackIpAddressList(d []interface{}) []edpt.InterfaceLoopbackIpAddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpAddressList
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspf652(d []interface{}) edpt.InterfaceLoopbackIpOspf652 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspf652
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OspfGlobal = getObjectInterfaceLoopbackIpOspfOspfGlobal653(in["ospf_global"].([]interface{}))
		ret.OspfIpList = getSliceInterfaceLoopbackIpOspfOspfIpList(in["ospf_ip_list"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobal653(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobal653 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobal653
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthenticationCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg654(in["authentication_cfg"].([]interface{}))
		ret.AuthenticationKey = in["authentication_key"].(string)
		ret.BfdCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalBfdCfg655(in["bfd_cfg"].([]interface{}))
		ret.Cost = in["cost"].(int)
		ret.DatabaseFilterCfg = getObjectInterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg656(in["database_filter_cfg"].([]interface{}))
		ret.DeadInterval = in["dead_interval"].(int)
		ret.Disable = in["disable"].(string)
		ret.HelloInterval = in["hello_interval"].(int)
		ret.MessageDigestCfg = getSliceInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg657(in["message_digest_cfg"].([]interface{}))
		ret.Mtu = in["mtu"].(int)
		ret.MtuIgnore = in["mtu_ignore"].(int)
		ret.Priority = in["priority"].(int)
		ret.RetransmitInterval = in["retransmit_interval"].(int)
		ret.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg654(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg654 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalAuthenticationCfg654
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.Value = in["value"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalBfdCfg655(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalBfdCfg655 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalBfdCfg655
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg656(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg656 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalDatabaseFilterCfg656
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DatabaseFilter = in["database_filter"].(string)
		ret.Out = in["out"].(int)
	}
	return ret
}

func getSliceInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg657(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg657 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg657, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfg657
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = getObjectInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5658(in["md5"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5658(d []interface{}) edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5658 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpOspfOspfGlobalMessageDigestCfgMd5658
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Md5Value = in["md5_value"].(string)
		//omit encrypted
	}
	return ret
}

func getSliceInterfaceLoopbackIpOspfOspfIpList(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfIpList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfIpList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Authentication = in["authentication"].(int)
		oi.Value = in["value"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.Cost = in["cost"].(int)
		oi.DatabaseFilter = in["database_filter"].(string)
		oi.Out = in["out"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.MessageDigestCfg = getSliceInterfaceLoopbackIpOspfOspfIpListMessageDigestCfg(in["message_digest_cfg"].([]interface{}))
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.Priority = in["priority"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpOspfOspfIpListMessageDigestCfg(d []interface{}) []edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpOspfOspfIpListMessageDigestCfg
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5Value = in["md5_value"].(string)
		//omit encrypted
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRip659(d []interface{}) edpt.InterfaceLoopbackIpRip659 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRip659
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = getObjectInterfaceLoopbackIpRipAuthentication660(in["authentication"].([]interface{}))
		ret.SendPacket = in["send_packet"].(int)
		ret.ReceivePacket = in["receive_packet"].(int)
		ret.SendCfg = getObjectInterfaceLoopbackIpRipSendCfg664(in["send_cfg"].([]interface{}))
		ret.ReceiveCfg = getObjectInterfaceLoopbackIpRipReceiveCfg665(in["receive_cfg"].([]interface{}))
		ret.SplitHorizonCfg = getObjectInterfaceLoopbackIpRipSplitHorizonCfg666(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthentication660(d []interface{}) edpt.InterfaceLoopbackIpRipAuthentication660 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthentication660
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceLoopbackIpRipAuthenticationStr661(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLoopbackIpRipAuthenticationMode662(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLoopbackIpRipAuthenticationKeyChain663(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationStr661(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationStr661 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationStr661
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationMode662(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationMode662 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationMode662
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationKeyChain663(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationKeyChain663 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationKeyChain663
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipSendCfg664(d []interface{}) edpt.InterfaceLoopbackIpRipSendCfg664 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipSendCfg664
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipReceiveCfg665(d []interface{}) edpt.InterfaceLoopbackIpRipReceiveCfg665 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipReceiveCfg665
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipSplitHorizonCfg666(d []interface{}) edpt.InterfaceLoopbackIpRipSplitHorizonCfg666 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipSplitHorizonCfg666
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRouter667(d []interface{}) edpt.InterfaceLoopbackIpRouter667 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRouter667
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Isis = getObjectInterfaceLoopbackIpRouterIsis668(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpRouterIsis668(d []interface{}) edpt.InterfaceLoopbackIpRouterIsis668 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRouterIsis668
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointInterfaceLoopbackIp(d *schema.ResourceData) edpt.InterfaceLoopbackIp {
	var ret edpt.InterfaceLoopbackIp
	ret.Inst.AddressList = getSliceInterfaceLoopbackIpAddressList(d.Get("address_list").([]interface{}))
	ret.Inst.Ospf = getObjectInterfaceLoopbackIpOspf652(d.Get("ospf").([]interface{}))
	ret.Inst.Rip = getObjectInterfaceLoopbackIpRip659(d.Get("rip").([]interface{}))
	ret.Inst.Router = getObjectInterfaceLoopbackIpRouter667(d.Get("router").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
