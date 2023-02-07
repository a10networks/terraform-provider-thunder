package thunder

//Thunder resource InterfaceManagement

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceManagement() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceManagementCreate,
		UpdateContext: resourceInterfaceManagementUpdate,
		ReadContext:   resourceInterfaceManagementRead,
		DeleteContext: resourceInterfaceManagementDelete,
		Schema: map[string]*schema.Schema{
			"broadcast_rate_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"bcast_rate_limit_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"access_list": {
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
						"acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ipv6": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"inbound": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipv6_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"v6_acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"default_ipv6_gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"default_gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"control_apps_use_mgmt_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipv4_netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dhcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"flow_control": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
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
			"lldp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tx_tlvs_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_description": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"system_name": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"system_capabilities": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tx_tlvs": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_description": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"exclude": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"management_address": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"tx_dot1_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"link_aggregation": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"vlan": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tx_dot1_tlvs": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"enable_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rt_enable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tx": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"rx": {
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
						"notification_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"notif_enable": {
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
			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"duplexity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"secondary_ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"default_gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"control_apps_use_mgmt_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipv4_netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dhcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"secondary_ip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceManagementCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceManagement (Inside resourceInterfaceManagementCreate) ")

		data := dataToInterfaceManagement(d)
		logger.Println("[INFO] received V from method data to InterfaceManagement --")
		d.SetId("1")
		err := go_thunder.PostInterfaceManagement(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceManagementRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceManagementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading InterfaceManagement (Inside resourceInterfaceManagementRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceManagement(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceInterfaceManagementUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceManagementRead(ctx, d, meta)
}

func resourceInterfaceManagementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceInterfaceManagementRead(ctx, d, meta)
}

func dataToInterfaceManagement(d *schema.ResourceData) go_thunder.Management {
	var vc go_thunder.Management
	var c go_thunder.ManagementInstance
	var obj1 go_thunder.Lldp
	prefix := "lldp.0."

	var obj1_1 go_thunder.TxDot1Cfg
	prefix1 := prefix + "tx_dot1_cfg.0."
	obj1_1.LinkAggregation = d.Get(prefix1 + "link_aggregation").(int)
	obj1_1.Vlan = d.Get(prefix1 + "vlan").(int)
	obj1_1.TxDot1Tlvs = d.Get(prefix1 + "tx_dot1_tlvs").(int)
	obj1.LinkAggregation = obj1_1

	var obj1_2 go_thunder.NotificationCfg
	prefix2 := prefix + "notification_cfg.0."
	obj1_2.Notification = d.Get(prefix2 + "notification").(int)
	obj1_2.NotifEnable = d.Get(prefix2 + "notif_enable").(int)
	obj1.Notification = obj1_2

	var obj1_3 go_thunder.EnableCfg
	prefix3 := prefix + "enable_cfg.0."
	obj1_3.Rx = d.Get(prefix3 + "rx").(int)
	obj1_3.Tx = d.Get(prefix3 + "tx").(int)
	obj1_3.RtEnable = d.Get(prefix3 + "rt_enable").(int)
	obj1.Rx = obj1_3

	var obj1_4 go_thunder.TxTlvsCfg
	prefix4 := prefix + "tx_tlvs_cfg.0."
	obj1_4.SystemCapabilities = d.Get(prefix4 + "system_capabilities").(int)
	obj1_4.SystemDescription = d.Get(prefix4 + "system_description").(int)
	obj1_4.ManagementAddress = d.Get(prefix4 + "management_address").(int)
	obj1_4.TxTlvs = d.Get(prefix4 + "tx_tlvs").(int)
	obj1_4.Exclude = d.Get(prefix4 + "exclude").(int)
	obj1_4.PortDescription = d.Get(prefix4 + "port_description").(int)
	obj1_4.SystemName = d.Get(prefix4 + "system_name").(int)
	obj1.SystemCapabilities = obj1_4

	c.LinkAggregation = obj1

	c.FlowControl = d.Get("flow_control").(int)

	var obj2 go_thunder.BroadcastRateLimit
	prefix = "broadcast_rate_limit.0."
	obj2.Rate = d.Get(prefix + "rate").(int)
	obj2.BcastRateLimitEnable = d.Get(prefix + "bcast_rate_limit_enable").(int)
	c.Rate = obj2

	c.Duplexity = d.Get("duplexity").(string)

	var obj3 go_thunder.ManagementIP
	prefix = "ip.0."
	obj3.Dhcp = d.Get(prefix + "dhcp").(int)
	obj3.Ipv4Address = d.Get(prefix + "ipv4_address").(string)
	obj3.ControlAppsUseMgmtPort = d.Get(prefix + "control_apps_use_mgmt_port").(int)
	obj3.DefaultGateway = d.Get(prefix + "default_gateway").(string)
	obj3.Ipv4Netmask = d.Get(prefix + "ipv4_netmask").(string)
	c.Dhcp = obj3

	var obj4 go_thunder.SecondaryIP
	prefix = "secondary_ip.0."
	obj4.Ipv4Netmask = d.Get(prefix + "ipv4_netmask").(string)
	obj4.ControlAppsUseMgmtPort = d.Get(prefix + "control_apps_use_mgmt_port").(int)
	obj4.SecondaryIP = d.Get(prefix + "secondary_ip").(int)
	obj4.DefaultGateway = d.Get(prefix + "default_gateway").(string)
	obj4.Dhcp = d.Get(prefix + "dhcp").(int)
	obj4.Ipv4Address = d.Get(prefix + "ipv4_address").(string)
	c.Ipv4Netmask = obj4

	var obj5 go_thunder.AccessList
	prefix = "access_list.0."
	obj5.ACLName = d.Get(prefix + "acl_name").(string)
	obj5.ACLID = d.Get(prefix + "acl_id").(int)
	c.ACLName = obj5

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.ManagementSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj6 go_thunder.ManagementSamplingEnable
		prefix = fmt.Sprintf("sampling_enable.%d.", i)
		obj6.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj6)
	}

	Ipv6Count := d.Get("ipv6.#").(int)
	c.DefaultIpv6Gateway = make([]go_thunder.Ipv6, 0, Ipv6Count)

	for i := 0; i < Ipv6Count; i++ {
		var obj7 go_thunder.Ipv6
		prefix = fmt.Sprintf("ipv6.%d.", i)
		obj7.DefaultIpv6Gateway = d.Get(prefix + "default_ipv6_gateway").(string)
		obj7.Inbound = d.Get(prefix + "inbound").(int)
		obj7.AddressType = d.Get(prefix + "address_type").(string)
		obj7.Ipv6Addr = d.Get(prefix + "ipv6_addr").(string)
		obj7.V6ACLName = d.Get(prefix + "v6_acl_name").(string)
		c.DefaultIpv6Gateway = append(c.DefaultIpv6Gateway, obj7)
	}

	c.Action = d.Get("action").(string)
	c.Speed = d.Get("speed").(string)

	vc.UUID = c
	return vc
}
