package thunder

//Thunder resource VrrpAVrid

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"util"
)

func resourceVrrpAVrid() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVrrpAVridCreate,
		UpdateContext: resourceVrrpAVridUpdate,
		ReadContext:   resourceVrrpAVridRead,
		DeleteContext: resourceVrrpAVridDelete,
		Schema: map[string]*schema.Schema{
			"vrid_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"floating_ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ip_address_part_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address_partition": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv6_address_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ethernet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"trunk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ve": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipv6_address_part_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_address_partition": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ethernet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"trunk": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ve": {
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
			"preempt_mode": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold": {
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
			"follow": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid_lead": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"pair_follow": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pair_follow": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"vrid_lead": {
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
			"user_tag": {
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
			"blade_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fail_over_policy_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"tracking_options": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethernet": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"priority_cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"route": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_destination_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_destination": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"mask": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gateway": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"distance": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"protocol": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ipv6_destination_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_destination": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gatewayv6": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"distance": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"protocol": {
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
									"trunk_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trunk": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"priority_cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"per_port_pri": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"bgp": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bgp_ipv4_address_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bgp_ipv4_address": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"bgp_ipv6_address_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bgp_ipv6_address": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
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
									"vlan_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vlan": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"timeout": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"priority_cost": {
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
									"gateway": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_gateway_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_address": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
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
												"ipv6_gateway_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_address": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"priority_cost": {
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVrrpAVridCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating VrrpAVrid (Inside resourceVrrpAVridCreate) ")
		name1 := d.Get("vrid_val").(int)
		data := dataToVrrpAVrid(d)
		logger.Println("[INFO] received formatted data from method data to VrrpAVrid --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostVrrpAVrid(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceVrrpAVridRead(ctx, d, meta)

	}
	return diags
}

func resourceVrrpAVridRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading VrrpAVrid (Inside resourceVrrpAVridRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetVrrpAVrid(client.Token, name1, client.Host)
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

func resourceVrrpAVridUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying VrrpAVrid   (Inside resourceVrrpAVridUpdate) ")
		data := dataToVrrpAVrid(d)
		logger.Println("[INFO] received formatted data from method data to VrrpAVrid ")
		err := go_thunder.PutVrrpAVrid(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceVrrpAVridRead(ctx, d, meta)

	}
	return diags
}

func resourceVrrpAVridDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceVrrpAVridDelete) " + name1)
		err := go_thunder.DeleteVrrpAVrid(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToVrrpAVrid(d *schema.ResourceData) go_thunder.VrrpAVrid {
	var vc go_thunder.VrrpAVrid
	var c go_thunder.VrrpAVridInstance
	c.VrrpAVridInstanceVridVal = d.Get("vrid_val").(int)

	var obj1 go_thunder.VrrpAVridInstanceFloatingIP
	prefix1 := "floating_ip.0."

	VrrpAVridInstanceFloatingIPIPAddressCfgCount := d.Get(prefix1 + "ip_address_cfg.#").(int)
	obj1.VrrpAVridInstanceFloatingIPIPAddressCfgIPAddress = make([]go_thunder.VrrpAVridInstanceFloatingIPIPAddressCfg, 0, VrrpAVridInstanceFloatingIPIPAddressCfgCount)

	for i := 0; i < VrrpAVridInstanceFloatingIPIPAddressCfgCount; i++ {
		var obj1_1 go_thunder.VrrpAVridInstanceFloatingIPIPAddressCfg
		prefix1_1 := prefix1 + fmt.Sprintf("ip_address_cfg.%d.", i)
		obj1_1.VrrpAVridInstanceFloatingIPIPAddressCfgIPAddress = d.Get(prefix1_1 + "ip_address").(string)
		obj1.VrrpAVridInstanceFloatingIPIPAddressCfgIPAddress = append(obj1.VrrpAVridInstanceFloatingIPIPAddressCfgIPAddress, obj1_1)
	}

	VrrpAVridInstanceFloatingIPIPAddressPartCfgCount := d.Get(prefix1 + "ip_address_part_cfg.#").(int)
	obj1.VrrpAVridInstanceFloatingIPIPAddressPartCfgIPAddressPartition = make([]go_thunder.VrrpAVridInstanceFloatingIPIPAddressPartCfg, 0, VrrpAVridInstanceFloatingIPIPAddressPartCfgCount)

	for i := 0; i < VrrpAVridInstanceFloatingIPIPAddressPartCfgCount; i++ {
		var obj1_2 go_thunder.VrrpAVridInstanceFloatingIPIPAddressPartCfg
		prefix1_2 := prefix1 + fmt.Sprintf("ip_address_part_cfg.%d.", i)
		obj1_2.VrrpAVridInstanceFloatingIPIPAddressPartCfgIPAddressPartition = d.Get(prefix1_2 + "ip_address_partition").(string)
		obj1.VrrpAVridInstanceFloatingIPIPAddressPartCfgIPAddressPartition = append(obj1.VrrpAVridInstanceFloatingIPIPAddressPartCfgIPAddressPartition, obj1_2)
	}

	VrrpAVridInstanceFloatingIPIpv6AddressCfgCount := d.Get(prefix1 + "ipv6_address_cfg.#").(int)
	obj1.VrrpAVridInstanceFloatingIPIpv6AddressCfgIpv6Address = make([]go_thunder.VrrpAVridInstanceFloatingIPIpv6AddressCfg, 0, VrrpAVridInstanceFloatingIPIpv6AddressCfgCount)

	for i := 0; i < VrrpAVridInstanceFloatingIPIpv6AddressCfgCount; i++ {
		var obj1_3 go_thunder.VrrpAVridInstanceFloatingIPIpv6AddressCfg
		prefix1_3 := prefix1 + fmt.Sprintf("ipv6_address_cfg.%d.", i)
		obj1_3.VrrpAVridInstanceFloatingIPIpv6AddressCfgIpv6Address = d.Get(prefix1_3 + "ipv6_address").(string)
		obj1_3.VrrpAVridInstanceFloatingIPIpv6AddressCfgEthernet = d.Get(prefix1_3 + "ethernet").(int)
		obj1_3.VrrpAVridInstanceFloatingIPIpv6AddressCfgTrunk = d.Get(prefix1_3 + "trunk").(int)
		obj1_3.VrrpAVridInstanceFloatingIPIpv6AddressCfgVe = d.Get(prefix1_3 + "ve").(int)
		obj1.VrrpAVridInstanceFloatingIPIpv6AddressCfgIpv6Address = append(obj1.VrrpAVridInstanceFloatingIPIpv6AddressCfgIpv6Address, obj1_3)
	}

	VrrpAVridInstanceFloatingIPIpv6AddressPartCfgCount := d.Get(prefix1 + "ipv6_address_part_cfg.#").(int)
	obj1.VrrpAVridInstanceFloatingIPIpv6AddressPartCfgIpv6AddressPartition = make([]go_thunder.VrrpAVridInstanceFloatingIPIpv6AddressPartCfg, 0, VrrpAVridInstanceFloatingIPIpv6AddressPartCfgCount)

	for i := 0; i < VrrpAVridInstanceFloatingIPIpv6AddressPartCfgCount; i++ {
		var obj1_4 go_thunder.VrrpAVridInstanceFloatingIPIpv6AddressPartCfg
		prefix1_4 := prefix1 + fmt.Sprintf("ipv6_address_part_cfg.%d.", i)
		obj1_4.VrrpAVridInstanceFloatingIPIpv6AddressPartCfgIpv6AddressPartition = d.Get(prefix1_4 + "ipv6_address_partition").(string)
		obj1_4.VrrpAVridInstanceFloatingIPIpv6AddressPartCfgEthernet = d.Get(prefix1_4 + "ethernet").(int)
		obj1_4.VrrpAVridInstanceFloatingIPIpv6AddressPartCfgTrunk = d.Get(prefix1_4 + "trunk").(int)
		obj1_4.VrrpAVridInstanceFloatingIPIpv6AddressPartCfgVe = d.Get(prefix1_4 + "ve").(int)
		obj1.VrrpAVridInstanceFloatingIPIpv6AddressPartCfgIpv6AddressPartition = append(obj1.VrrpAVridInstanceFloatingIPIpv6AddressPartCfgIpv6AddressPartition, obj1_4)
	}

	c.VrrpAVridInstanceFloatingIPIPAddressCfg = obj1

	var obj2 go_thunder.VrrpAVridInstancePreemptMode
	prefix2 := "preempt_mode.0."
	obj2.VrrpAVridInstancePreemptModeThreshold = d.Get(prefix2 + "threshold").(int)
	obj2.VrrpAVridInstancePreemptModeDisable = d.Get(prefix2 + "disable").(int)

	c.VrrpAVridInstancePreemptModeThreshold = obj2

	var obj3 go_thunder.VrrpAVridInstanceFollow
	prefix3 := "follow.0."
	obj3.VrrpAVridInstanceFollowVridLead = d.Get(prefix3 + "vrid_lead").(string)

	c.VrrpAVridInstanceFollowVridLead = obj3

	var obj4 go_thunder.VrrpAVridInstancePairFollow
	prefix4 := "pair_follow.0."
	obj4.VrrpAVridInstancePairFollowPairFollow = d.Get(prefix4 + "pair_follow").(int)
	obj4.VrrpAVridInstancePairFollowVridLead = d.Get(prefix4 + "vrid_lead").(string)

	c.VrrpAVridInstancePairFollowPairFollow = obj4

	c.VrrpAVridInstanceUserTag = d.Get("user_tag").(string)

	VrrpAVridInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.VrrpAVridInstanceSamplingEnableCounters1 = make([]go_thunder.VrrpAVridInstanceSamplingEnable, 0, VrrpAVridInstanceSamplingEnableCount)

	for i := 0; i < VrrpAVridInstanceSamplingEnableCount; i++ {
		var obj5 go_thunder.VrrpAVridInstanceSamplingEnable
		prefix5 := fmt.Sprintf("sampling_enable.%d.", i)
		obj5.VrrpAVridInstanceSamplingEnableCounters1 = d.Get(prefix5 + "counters1").(string)
		c.VrrpAVridInstanceSamplingEnableCounters1 = append(c.VrrpAVridInstanceSamplingEnableCounters1, obj5)
	}

	var obj6 go_thunder.VrrpAVridInstanceBladeParameters
	prefix6 := "blade_parameters.0."
	obj6.VrrpAVridInstanceBladeParametersPriority = d.Get(prefix6 + "priority").(int)
	obj6.VrrpAVridInstanceBladeParametersFailOverPolicyTemplate = d.Get(prefix6 + "fail_over_policy_template").(string)

	var obj6_1 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptions
	prefix6_1 := prefix6 + "tracking_options.0."

	VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceCount := d.Get(prefix6_1 + "interface.#").(int)
	obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceEthernet = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsInterface, 0, VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceCount; i++ {
		var obj6_1_1 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsInterface
		prefix6_1_1 := prefix6_1 + fmt.Sprintf("interface.%d.", i)
		obj6_1_1.VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceEthernet = d.Get(prefix6_1_1 + "ethernet").(int)
		obj6_1_1.VrrpAVridInstanceBladeParametersTrackingOptionsInterfacePriorityCost = d.Get(prefix6_1_1 + "priority_cost").(int)
		obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceEthernet = append(obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceEthernet, obj6_1_1)
	}

	var obj6_1_2 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsRoute
	prefix6_1_2 := prefix6_1 + "route.0."

	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgCount := d.Get(prefix6_1_2 + "ip_destination_cfg.#").(int)
	obj6_1_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgIPDestination = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfg, 0, VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgCount; i++ {
		var obj6_1_2_1 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfg
		prefix6_1_2_1 := prefix6_1_2 + fmt.Sprintf("ip_destination_cfg.%d.", i)
		obj6_1_2_1.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgIPDestination = d.Get(prefix6_1_2_1 + "ip_destination").(string)
		obj6_1_2_1.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgMask = d.Get(prefix6_1_2_1 + "mask").(string)
		obj6_1_2_1.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgPriorityCost = d.Get(prefix6_1_2_1 + "priority_cost").(int)
		obj6_1_2_1.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgGateway = d.Get(prefix6_1_2_1 + "gateway").(string)
		obj6_1_2_1.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgDistance = d.Get(prefix6_1_2_1 + "distance").(int)
		obj6_1_2_1.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgProtocol = d.Get(prefix6_1_2_1 + "protocol").(string)
		obj6_1_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgIPDestination = append(obj6_1_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgIPDestination, obj6_1_2_1)
	}

	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgCount := d.Get(prefix6_1_2 + "ipv6_destination_cfg.#").(int)
	obj6_1_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgIpv6Destination = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfg, 0, VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgCount; i++ {
		var obj6_1_2_2 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfg
		prefix6_1_2_2 := prefix6_1_2 + fmt.Sprintf("ipv6_destination_cfg.%d.", i)
		obj6_1_2_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgIpv6Destination = d.Get(prefix6_1_2_2 + "ipv6_destination").(string)
		obj6_1_2_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgPriorityCost = d.Get(prefix6_1_2_2 + "priority_cost").(int)
		obj6_1_2_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgGatewayv6 = d.Get(prefix6_1_2_2 + "gatewayv6").(string)
		obj6_1_2_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgDistance = d.Get(prefix6_1_2_2 + "distance").(int)
		obj6_1_2_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgProtocol = d.Get(prefix6_1_2_2 + "protocol").(string)
		obj6_1_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgIpv6Destination = append(obj6_1_2.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgIpv6Destination, obj6_1_2_2)
	}

	obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfg = obj6_1_2

	VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgCount := d.Get(prefix6_1 + "trunk_cfg.#").(int)
	obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgTrunk = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfg, 0, VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgCount; i++ {
		var obj6_1_3 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfg
		prefix6_1_3 := prefix6_1 + fmt.Sprintf("trunk_cfg.%d.", i)
		obj6_1_3.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgTrunk = d.Get(prefix6_1_3 + "trunk").(int)
		obj6_1_3.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgPriorityCost = d.Get(prefix6_1_3 + "priority_cost").(int)
		obj6_1_3.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgPerPortPri = d.Get(prefix6_1_3 + "per_port_pri").(int)
		obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgTrunk = append(obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgTrunk, obj6_1_3)
	}

	var obj6_1_4 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsBgp
	prefix6_1_4 := prefix6_1 + "bgp.0."

	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgCount := d.Get(prefix6_1_4 + "bgp_ipv4_address_cfg.#").(int)
	obj6_1_4.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgBgpIpv4Address = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg, 0, VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgCount; i++ {
		var obj6_1_4_1 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg
		prefix6_1_4_1 := prefix6_1_4 + fmt.Sprintf("bgp_ipv4_address_cfg.%d.", i)
		obj6_1_4_1.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgBgpIpv4Address = d.Get(prefix6_1_4_1 + "bgp_ipv4_address").(string)
		obj6_1_4_1.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgPriorityCost = d.Get(prefix6_1_4_1 + "priority_cost").(int)
		obj6_1_4.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgBgpIpv4Address = append(obj6_1_4.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgBgpIpv4Address, obj6_1_4_1)
	}

	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgCount := d.Get(prefix6_1_4 + "bgp_ipv6_address_cfg.#").(int)
	obj6_1_4.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgBgpIpv6Address = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg, 0, VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgCount; i++ {
		var obj6_1_4_2 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg
		prefix6_1_4_2 := prefix6_1_4 + fmt.Sprintf("bgp_ipv6_address_cfg.%d.", i)
		obj6_1_4_2.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgBgpIpv6Address = d.Get(prefix6_1_4_2 + "bgp_ipv6_address").(string)
		obj6_1_4_2.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgPriorityCost = d.Get(prefix6_1_4_2 + "priority_cost").(int)
		obj6_1_4.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgBgpIpv6Address = append(obj6_1_4.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgBgpIpv6Address, obj6_1_4_2)
	}

	obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg = obj6_1_4

	VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgCount := d.Get(prefix6_1 + "vlan_cfg.#").(int)
	obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgVlan = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfg, 0, VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgCount; i++ {
		var obj6_1_5 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfg
		prefix6_1_5 := prefix6_1 + fmt.Sprintf("vlan_cfg.%d.", i)
		obj6_1_5.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgVlan = d.Get(prefix6_1_5 + "vlan").(int)
		obj6_1_5.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgTimeout = d.Get(prefix6_1_5 + "timeout").(int)
		obj6_1_5.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgPriorityCost = d.Get(prefix6_1_5 + "priority_cost").(int)
		obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgVlan = append(obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgVlan, obj6_1_5)
	}

	var obj6_1_6 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsGateway
	prefix6_1_6 := prefix6_1 + "gateway.0."

	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListCount := d.Get(prefix6_1_6 + "ipv4_gateway_list.#").(int)
	obj6_1_6.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListIPAddress = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayList, 0, VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListCount; i++ {
		var obj6_1_6_1 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayList
		prefix6_1_6_1 := prefix6_1_6 + fmt.Sprintf("ipv4_gateway_list.%d.", i)
		obj6_1_6_1.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListIPAddress = d.Get(prefix6_1_6_1 + "ip_address").(string)
		obj6_1_6_1.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListPriorityCost = d.Get(prefix6_1_6_1 + "priority_cost").(int)
		obj6_1_6.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListIPAddress = append(obj6_1_6.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListIPAddress, obj6_1_6_1)
	}

	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListCount := d.Get(prefix6_1_6 + "ipv6_gateway_list.#").(int)
	obj6_1_6.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListIpv6Address = make([]go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayList, 0, VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListCount)

	for i := 0; i < VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListCount; i++ {
		var obj6_1_6_2 go_thunder.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayList
		prefix6_1_6_2 := prefix6_1_6 + fmt.Sprintf("ipv6_gateway_list.%d.", i)
		obj6_1_6_2.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListIpv6Address = d.Get(prefix6_1_6_2 + "ipv6_address").(string)
		obj6_1_6_2.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListPriorityCost = d.Get(prefix6_1_6_2 + "priority_cost").(int)
		obj6_1_6.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListIpv6Address = append(obj6_1_6.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListIpv6Address, obj6_1_6_2)
	}

	obj6_1.VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayList = obj6_1_6

	obj6.VrrpAVridInstanceBladeParametersTrackingOptionsInterface = obj6_1

	c.VrrpAVridInstanceBladeParametersPriority = obj6

	vc.VrrpAVridInstanceVridVal = c
	return vc
}
