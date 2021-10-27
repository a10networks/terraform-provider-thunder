package thunder

//Thunder resource RuleSet

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceRuleSet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRuleSetCreate,
		UpdateContext: resourceRuleSetUpdate,
		ReadContext:   resourceRuleSetRead,
		DeleteContext: resourceRuleSetDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"session_statistic": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"packet_capture_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rule_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"remark": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"status": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reset_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"listen_on_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"vpn_ipsec_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"forward_listen_on_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"listen_on_port_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fw_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fwlog": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cgnv6_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"forward_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lidlog": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reset_lidlog": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"listen_on_port_lidlog": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cgnv6_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cgnv6_fixed_nat_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cgnv6_lsn_lid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cgnv6_lsn_log": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip_version": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"gtp_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_class_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_geoloc_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_geoloc_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_geoloc_list_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"src_ipv4_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_ipv6_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"source_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_ip_subnet": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"src_ipv6_subnet": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"src_obj_network": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"src_obj_grp_network": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"src_slb_server": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"src_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_zone_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"src_threat_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_class_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_geoloc_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_geoloc_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_geoloc_list_shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dst_ipv4_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_ipv6_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dest_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dst_ip_subnet": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dst_ipv6_subnet": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dst_obj_network": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dst_obj_grp_network": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dst_slb_server": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dst_slb_vserver": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"dst_domain_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_zone_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dst_threat_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocols": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"proto_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"obj_grp_service": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"icmp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"icmpv6": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"icmp_type": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"special_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"icmp_code": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"special_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"icmpv6_type": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"special_v6_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"icmpv6_code": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"special_v6_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"eq_src_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"gt_src_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lt_src_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"range_src_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_num_end_src": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"eq_dst_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"gt_dst_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"lt_dst_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"range_dst_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_num_end_dst": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"sctp_template": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"alg": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"idle_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dscp_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dscp_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dscp_range_start": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dscp_range_end": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"application_any": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"app_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"obj_grp_application": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"protocol": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"protocol_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"track_application": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
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
						"packet_capture_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action_group": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"permit_log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"reset_log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"deny_log": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"listen_on_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"forward": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ipsec": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"vpn_ipsec_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"cgnv6": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"cgnv6_policy": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"cgnv6_lsn_lid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"permit_limit_policy": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"permit_respond_to_user_mac": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"reset_respond_to_user_mac": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"set_dscp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dscp_value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"dscp_number": {
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
						"move_rule": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"location": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"target_rule": {
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
			"rules_by_zone": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"application": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"track_app_rule_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"app": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"tag": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceRuleSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating RuleSet (Inside resourceRuleSetCreate) ")
		name1 := d.Get("name").(string)
		data := dataToRuleSet(d)
		logger.Println("[INFO] received formatted data from method data to RuleSet --")
		d.SetId(name1)
		err := go_thunder.PostRuleSet(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRuleSetRead(ctx, d, meta)

	}
	return diags
}

func resourceRuleSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RuleSet (Inside resourceRuleSetRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRuleSet(client.Token, name1, client.Host)
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

func resourceRuleSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying RuleSet   (Inside resourceRuleSetUpdate) ")
		data := dataToRuleSet(d)
		logger.Println("[INFO] received formatted data from method data to RuleSet ")
		err := go_thunder.PutRuleSet(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRuleSetRead(ctx, d, meta)

	}
	return diags
}

func resourceRuleSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceRuleSetDelete) " + name1)
		err := go_thunder.DeleteRuleSet(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToRuleSet(d *schema.ResourceData) go_thunder.RuleSet {
	var vc go_thunder.RuleSet
	var c go_thunder.RuleSetInstance
	c.RuleSetInstanceName = d.Get("name").(string)
	c.RuleSetInstanceSessionStatistic = d.Get("session_statistic").(string)
	c.RuleSetInstanceRemark = d.Get("remark").(string)
	c.RuleSetInstanceUserTag = d.Get("user_tag").(string)

	RuleSetInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.RuleSetInstanceSamplingEnableCounters1 = make([]go_thunder.RuleSetInstanceSamplingEnable, 0, RuleSetInstanceSamplingEnableCount)

	for i := 0; i < RuleSetInstanceSamplingEnableCount; i++ {
		var obj1 go_thunder.RuleSetInstanceSamplingEnable
		prefix1 := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.RuleSetInstanceSamplingEnableCounters1 = d.Get(prefix1 + "counters1").(string)
		c.RuleSetInstanceSamplingEnableCounters1 = append(c.RuleSetInstanceSamplingEnableCounters1, obj1)
	}

	c.RuleSetInstancePacketCaptureTemplate = d.Get("packet_capture_template").(string)

	RuleSetInstanceRuleListCount := d.Get("rule_list.#").(int)
	c.RuleSetInstanceRuleListStatus = make([]go_thunder.RuleSetInstanceRuleList, 0, RuleSetInstanceRuleListCount)

	for i := 0; i < RuleSetInstanceRuleListCount; i++ {
		var obj2 go_thunder.RuleSetInstanceRuleList
		prefix2 := fmt.Sprintf("rule_list.%d.", i)
		obj2.RuleSetInstanceRuleListName = d.Get(prefix2 + "name").(string)
		obj2.RuleSetInstanceRuleListRemark = d.Get(prefix2 + "remark").(string)
		obj2.RuleSetInstanceRuleListStatus = d.Get(prefix2 + "status").(string)
		obj2.RuleSetInstanceRuleListAction = d.Get(prefix2 + "action").(string)
		obj2.RuleSetInstanceRuleListLog = d.Get(prefix2 + "log").(int)
		obj2.RuleSetInstanceRuleListResetLid = d.Get(prefix2 + "reset_lid").(int)
		obj2.RuleSetInstanceRuleListListenOnPort = d.Get(prefix2 + "listen_on_port").(int)
		obj2.RuleSetInstanceRuleListPolicy = d.Get(prefix2 + "policy").(string)
		obj2.RuleSetInstanceRuleListVpnIpsecName = d.Get(prefix2 + "vpn_ipsec_name").(string)
		obj2.RuleSetInstanceRuleListForwardListenOnPort = d.Get(prefix2 + "forward_listen_on_port").(int)
		obj2.RuleSetInstanceRuleListLid = d.Get(prefix2 + "lid").(int)
		obj2.RuleSetInstanceRuleListListenOnPortLid = d.Get(prefix2 + "listen_on_port_lid").(int)
		obj2.RuleSetInstanceRuleListFwLog = d.Get(prefix2 + "fw_log").(int)
		obj2.RuleSetInstanceRuleListFwlog = d.Get(prefix2 + "fwlog").(int)
		obj2.RuleSetInstanceRuleListCgnv6Log = d.Get(prefix2 + "cgnv6_log").(int)
		obj2.RuleSetInstanceRuleListForwardLog = d.Get(prefix2 + "forward_log").(int)
		obj2.RuleSetInstanceRuleListLidlog = d.Get(prefix2 + "lidlog").(int)
		obj2.RuleSetInstanceRuleListResetLidlog = d.Get(prefix2 + "reset_lidlog").(int)
		obj2.RuleSetInstanceRuleListListenOnPortLidlog = d.Get(prefix2 + "listen_on_port_lidlog").(int)
		obj2.RuleSetInstanceRuleListCgnv6Policy = d.Get(prefix2 + "cgnv6_policy").(string)
		obj2.RuleSetInstanceRuleListCgnv6FixedNatLog = d.Get(prefix2 + "cgnv6_fixed_nat_log").(int)
		obj2.RuleSetInstanceRuleListCgnv6LsnLid = d.Get(prefix2 + "cgnv6_lsn_lid").(int)
		obj2.RuleSetInstanceRuleListCgnv6LsnLog = d.Get(prefix2 + "cgnv6_lsn_log").(int)
		obj2.RuleSetInstanceRuleListIPVersion = d.Get(prefix2 + "ip_version").(string)
		obj2.RuleSetInstanceRuleListGtpTemplate = d.Get(prefix2 + "gtp_template").(string)
		obj2.RuleSetInstanceRuleListSrcClassList = d.Get(prefix2 + "src_class_list").(string)
		obj2.RuleSetInstanceRuleListSrcGeolocName = d.Get(prefix2 + "src_geoloc_name").(string)
		obj2.RuleSetInstanceRuleListSrcGeolocList = d.Get(prefix2 + "src_geoloc_list").(string)
		obj2.RuleSetInstanceRuleListSrcGeolocListShared = d.Get(prefix2 + "src_geoloc_list_shared").(int)
		obj2.RuleSetInstanceRuleListSrcIpv4Any = d.Get(prefix2 + "src_ipv4_any").(string)
		obj2.RuleSetInstanceRuleListSrcIpv6Any = d.Get(prefix2 + "src_ipv6_any").(string)

		RuleSetInstanceRuleListSourceListCount := d.Get(prefix2 + "source_list.#").(int)
		obj2.RuleSetInstanceRuleListSourceListSrcIPSubnet = make([]go_thunder.RuleSetInstanceRuleListSourceList, 0, RuleSetInstanceRuleListSourceListCount)

		for i := 0; i < RuleSetInstanceRuleListSourceListCount; i++ {
			var obj2_1 go_thunder.RuleSetInstanceRuleListSourceList
			prefix2_1 := prefix2 + fmt.Sprintf("source_list.%d.", i)
			obj2_1.RuleSetInstanceRuleListSourceListSrcIPSubnet = d.Get(prefix2_1 + "src_ip_subnet").(string)
			obj2_1.RuleSetInstanceRuleListSourceListSrcIpv6Subnet = d.Get(prefix2_1 + "src_ipv6_subnet").(string)
			obj2_1.RuleSetInstanceRuleListSourceListSrcObjNetwork = d.Get(prefix2_1 + "src_obj_network").(string)
			obj2_1.RuleSetInstanceRuleListSourceListSrcObjGrpNetwork = d.Get(prefix2_1 + "src_obj_grp_network").(string)
			obj2_1.RuleSetInstanceRuleListSourceListSrcSlbServer = d.Get(prefix2_1 + "src_slb_server").(string)
			obj2.RuleSetInstanceRuleListSourceListSrcIPSubnet = append(obj2.RuleSetInstanceRuleListSourceListSrcIPSubnet, obj2_1)
		}

		obj2.RuleSetInstanceRuleListSrcZone = d.Get(prefix2 + "src_zone").(string)
		obj2.RuleSetInstanceRuleListSrcZoneAny = d.Get(prefix2 + "src_zone_any").(string)
		obj2.RuleSetInstanceRuleListSrcThreatList = d.Get(prefix2 + "src_threat_list").(string)
		obj2.RuleSetInstanceRuleListDstClassList = d.Get(prefix2 + "dst_class_list").(string)
		obj2.RuleSetInstanceRuleListDstGeolocName = d.Get(prefix2 + "dst_geoloc_name").(string)
		obj2.RuleSetInstanceRuleListDstGeolocList = d.Get(prefix2 + "dst_geoloc_list").(string)
		obj2.RuleSetInstanceRuleListDstGeolocListShared = d.Get(prefix2 + "dst_geoloc_list_shared").(int)
		obj2.RuleSetInstanceRuleListDstIpv4Any = d.Get(prefix2 + "dst_ipv4_any").(string)
		obj2.RuleSetInstanceRuleListDstIpv6Any = d.Get(prefix2 + "dst_ipv6_any").(string)

		RuleSetInstanceRuleListDestListCount := d.Get(prefix2 + "dest_list.#").(int)
		obj2.RuleSetInstanceRuleListDestListDstIPSubnet = make([]go_thunder.RuleSetInstanceRuleListDestList, 0, RuleSetInstanceRuleListDestListCount)

		for i := 0; i < RuleSetInstanceRuleListDestListCount; i++ {
			var obj2_2 go_thunder.RuleSetInstanceRuleListDestList
			prefix2_2 := prefix2 + fmt.Sprintf("dest_list.%d.", i)
			obj2_2.RuleSetInstanceRuleListDestListDstIPSubnet = d.Get(prefix2_2 + "dst_ip_subnet").(string)
			obj2_2.RuleSetInstanceRuleListDestListDstIpv6Subnet = d.Get(prefix2_2 + "dst_ipv6_subnet").(string)
			obj2_2.RuleSetInstanceRuleListDestListDstObjNetwork = d.Get(prefix2_2 + "dst_obj_network").(string)
			obj2_2.RuleSetInstanceRuleListDestListDstObjGrpNetwork = d.Get(prefix2_2 + "dst_obj_grp_network").(string)
			obj2_2.RuleSetInstanceRuleListDestListDstSlbServer = d.Get(prefix2_2 + "dst_slb_server").(string)
			obj2_2.RuleSetInstanceRuleListDestListDstSlbVserver = d.Get(prefix2_2 + "dst_slb_vserver").(string)
			obj2.RuleSetInstanceRuleListDestListDstIPSubnet = append(obj2.RuleSetInstanceRuleListDestListDstIPSubnet, obj2_2)
		}

		obj2.RuleSetInstanceRuleListDstDomainList = d.Get(prefix2 + "dst_domain_list").(string)
		obj2.RuleSetInstanceRuleListDstZone = d.Get(prefix2 + "dst_zone").(string)
		obj2.RuleSetInstanceRuleListDstZoneAny = d.Get(prefix2 + "dst_zone_any").(string)
		obj2.RuleSetInstanceRuleListDstThreatList = d.Get(prefix2 + "dst_threat_list").(string)
		obj2.RuleSetInstanceRuleListServiceAny = d.Get(prefix2 + "service_any").(string)

		RuleSetInstanceRuleListServiceListCount := d.Get(prefix2 + "service_list.#").(int)
		obj2.RuleSetInstanceRuleListServiceListProtocols = make([]go_thunder.RuleSetInstanceRuleListServiceList, 0, RuleSetInstanceRuleListServiceListCount)

		for i := 0; i < RuleSetInstanceRuleListServiceListCount; i++ {
			var obj2_3 go_thunder.RuleSetInstanceRuleListServiceList
			prefix2_3 := prefix2 + fmt.Sprintf("service_list.%d.", i)
			obj2_3.RuleSetInstanceRuleListServiceListProtocols = d.Get(prefix2_3 + "protocols").(string)
			obj2_3.RuleSetInstanceRuleListServiceListProtoID = d.Get(prefix2_3 + "proto_id").(int)
			obj2_3.RuleSetInstanceRuleListServiceListObjGrpService = d.Get(prefix2_3 + "obj_grp_service").(string)
			obj2_3.RuleSetInstanceRuleListServiceListIcmp = d.Get(prefix2_3 + "icmp").(int)
			obj2_3.RuleSetInstanceRuleListServiceListIcmpv6 = d.Get(prefix2_3 + "icmpv6").(int)
			obj2_3.RuleSetInstanceRuleListServiceListIcmpType = d.Get(prefix2_3 + "icmp_type").(int)
			obj2_3.RuleSetInstanceRuleListServiceListSpecialType = d.Get(prefix2_3 + "special_type").(string)
			obj2_3.RuleSetInstanceRuleListServiceListIcmpCode = d.Get(prefix2_3 + "icmp_code").(int)
			obj2_3.RuleSetInstanceRuleListServiceListSpecialCode = d.Get(prefix2_3 + "special_code").(string)
			obj2_3.RuleSetInstanceRuleListServiceListIcmpv6Type = d.Get(prefix2_3 + "icmpv6_type").(int)
			obj2_3.RuleSetInstanceRuleListServiceListSpecialV6Type = d.Get(prefix2_3 + "special_v6_type").(string)
			obj2_3.RuleSetInstanceRuleListServiceListIcmpv6Code = d.Get(prefix2_3 + "icmpv6_code").(int)
			obj2_3.RuleSetInstanceRuleListServiceListSpecialV6Code = d.Get(prefix2_3 + "special_v6_code").(string)
			obj2_3.RuleSetInstanceRuleListServiceListEqSrcPort = d.Get(prefix2_3 + "eq_src_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListGtSrcPort = d.Get(prefix2_3 + "gt_src_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListLtSrcPort = d.Get(prefix2_3 + "lt_src_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListRangeSrcPort = d.Get(prefix2_3 + "range_src_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListPortNumEndSrc = d.Get(prefix2_3 + "port_num_end_src").(int)
			obj2_3.RuleSetInstanceRuleListServiceListEqDstPort = d.Get(prefix2_3 + "eq_dst_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListGtDstPort = d.Get(prefix2_3 + "gt_dst_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListLtDstPort = d.Get(prefix2_3 + "lt_dst_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListRangeDstPort = d.Get(prefix2_3 + "range_dst_port").(int)
			obj2_3.RuleSetInstanceRuleListServiceListPortNumEndDst = d.Get(prefix2_3 + "port_num_end_dst").(int)
			obj2_3.RuleSetInstanceRuleListServiceListSctpTemplate = d.Get(prefix2_3 + "sctp_template").(string)
			obj2_3.RuleSetInstanceRuleListServiceListAlg = d.Get(prefix2_3 + "alg").(string)
			obj2.RuleSetInstanceRuleListServiceListProtocols = append(obj2.RuleSetInstanceRuleListServiceListProtocols, obj2_3)
		}

		obj2.RuleSetInstanceRuleListIdleTimeout = d.Get(prefix2 + "idle_timeout").(int)

		RuleSetInstanceRuleListDscpListCount := d.Get(prefix2 + "dscp_list.#").(int)
		obj2.RuleSetInstanceRuleListDscpListDscpValue = make([]go_thunder.RuleSetInstanceRuleListDscpList, 0, RuleSetInstanceRuleListDscpListCount)

		for i := 0; i < RuleSetInstanceRuleListDscpListCount; i++ {
			var obj2_4 go_thunder.RuleSetInstanceRuleListDscpList
			prefix2_4 := prefix2 + fmt.Sprintf("dscp_list.%d.", i)
			obj2_4.RuleSetInstanceRuleListDscpListDscpValue = d.Get(prefix2_4 + "dscp_value").(string)
			obj2_4.RuleSetInstanceRuleListDscpListDscpRangeStart = d.Get(prefix2_4 + "dscp_range_start").(int)
			obj2_4.RuleSetInstanceRuleListDscpListDscpRangeEnd = d.Get(prefix2_4 + "dscp_range_end").(int)
			obj2.RuleSetInstanceRuleListDscpListDscpValue = append(obj2.RuleSetInstanceRuleListDscpListDscpValue, obj2_4)
		}

		obj2.RuleSetInstanceRuleListApplicationAny = d.Get(prefix2 + "application_any").(string)

		RuleSetInstanceRuleListAppListCount := d.Get(prefix2 + "app_list.#").(int)
		obj2.RuleSetInstanceRuleListAppListObjGrpApplication = make([]go_thunder.RuleSetInstanceRuleListAppList, 0, RuleSetInstanceRuleListAppListCount)

		for i := 0; i < RuleSetInstanceRuleListAppListCount; i++ {
			var obj2_5 go_thunder.RuleSetInstanceRuleListAppList
			prefix2_5 := prefix2 + fmt.Sprintf("app_list.%d.", i)
			obj2_5.RuleSetInstanceRuleListAppListObjGrpApplication = d.Get(prefix2_5 + "obj_grp_application").(string)
			obj2_5.RuleSetInstanceRuleListAppListProtocol = d.Get(prefix2_5 + "protocol").(string)
			obj2_5.RuleSetInstanceRuleListAppListProtocolTag = d.Get(prefix2_5 + "protocol_tag").(string)
			obj2.RuleSetInstanceRuleListAppListObjGrpApplication = append(obj2.RuleSetInstanceRuleListAppListObjGrpApplication, obj2_5)
		}

		obj2.RuleSetInstanceRuleListTrackApplication = d.Get(prefix2 + "track_application").(int)
		obj2.RuleSetInstanceRuleListUserTag = d.Get(prefix2 + "user_tag").(string)

		RuleSetInstanceRuleListSamplingEnableCount := d.Get(prefix2 + "sampling_enable.#").(int)
		obj2.RuleSetInstanceRuleListSamplingEnableCounters1 = make([]go_thunder.RuleSetInstanceRuleListSamplingEnable, 0, RuleSetInstanceRuleListSamplingEnableCount)

		for i := 0; i < RuleSetInstanceRuleListSamplingEnableCount; i++ {
			var obj2_6 go_thunder.RuleSetInstanceRuleListSamplingEnable
			prefix2_6 := prefix2 + fmt.Sprintf("sampling_enable.%d.", i)
			obj2_6.RuleSetInstanceRuleListSamplingEnableCounters1 = d.Get(prefix2_6 + "counters1").(string)
			obj2.RuleSetInstanceRuleListSamplingEnableCounters1 = append(obj2.RuleSetInstanceRuleListSamplingEnableCounters1, obj2_6)
		}

		obj2.RuleSetInstanceRuleListPacketCaptureTemplate = d.Get(prefix2 + "packet_capture_template").(string)

		var obj2_7 go_thunder.RuleSetInstanceRuleListActionGroup
		prefix2_7 := prefix2 + "action_group.0."
		obj2_7.RuleSetInstanceRuleListActionGroupType = d.Get(prefix2_7 + "type").(string)
		obj2_7.RuleSetInstanceRuleListActionGroupPermitLog = d.Get(prefix2_7 + "permit_log").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupResetLog = d.Get(prefix2_7 + "reset_log").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupDenyLog = d.Get(prefix2_7 + "deny_log").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupListenOnPort = d.Get(prefix2_7 + "listen_on_port").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupForward = d.Get(prefix2_7 + "forward").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupIpsec = d.Get(prefix2_7 + "ipsec").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupVpnIpsecName = d.Get(prefix2_7 + "vpn_ipsec_name").(string)
		obj2_7.RuleSetInstanceRuleListActionGroupCgnv6 = d.Get(prefix2_7 + "cgnv6").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupCgnv6Policy = d.Get(prefix2_7 + "cgnv6_policy").(string)
		obj2_7.RuleSetInstanceRuleListActionGroupCgnv6LsnLid = d.Get(prefix2_7 + "cgnv6_lsn_lid").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupPermitLimitPolicy = d.Get(prefix2_7 + "permit_limit_policy").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupPermitRespondToUserMac = d.Get(prefix2_7 + "permit_respond_to_user_mac").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupResetRespondToUserMac = d.Get(prefix2_7 + "reset_respond_to_user_mac").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupSetDscp = d.Get(prefix2_7 + "set_dscp").(int)
		obj2_7.RuleSetInstanceRuleListActionGroupDscpValue = d.Get(prefix2_7 + "dscp_value").(string)
		obj2_7.RuleSetInstanceRuleListActionGroupDscpNumber = d.Get(prefix2_7 + "dscp_number").(int)

		obj2.RuleSetInstanceRuleListActionGroupType = obj2_7

		var obj2_8 go_thunder.RuleSetInstanceRuleListMoveRule
		prefix2_8 := prefix2 + "move_rule.0."
		obj2_8.RuleSetInstanceRuleListMoveRuleLocation = d.Get(prefix2_8 + "location").(string)
		obj2_8.RuleSetInstanceRuleListMoveRuleTargetRule = d.Get(prefix2_8 + "target_rule").(string)

		obj2.RuleSetInstanceRuleListMoveRuleLocation = obj2_8

		c.RuleSetInstanceRuleListStatus = append(c.RuleSetInstanceRuleListStatus, obj2)
	}

	var obj3 go_thunder.RuleSetInstanceRulesByZone
	prefix3 := "rules_by_zone.0."

	RuleSetInstanceRulesByZoneSamplingEnableCount := d.Get(prefix3 + "sampling_enable.#").(int)
	obj3.RuleSetInstanceRulesByZoneSamplingEnableCounters1 = make([]go_thunder.RuleSetInstanceRulesByZoneSamplingEnable, 0, RuleSetInstanceRulesByZoneSamplingEnableCount)

	for i := 0; i < RuleSetInstanceRulesByZoneSamplingEnableCount; i++ {
		var obj3_1 go_thunder.RuleSetInstanceRulesByZoneSamplingEnable
		prefix3_1 := prefix3 + fmt.Sprintf("sampling_enable.%d.", i)
		obj3_1.RuleSetInstanceRulesByZoneSamplingEnableCounters1 = d.Get(prefix3_1 + "counters1").(string)
		obj3.RuleSetInstanceRulesByZoneSamplingEnableCounters1 = append(obj3.RuleSetInstanceRulesByZoneSamplingEnableCounters1, obj3_1)
	}

	c.RuleSetInstanceRulesByZoneUUID = obj3

	vc.RuleSetInstanceName = c
	return vc
}
