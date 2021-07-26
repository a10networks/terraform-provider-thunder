package thunder

//Thunder resource RouteMap

import (
	"context"
	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"strings"
	"util"
)

func resourceRouteMap() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouteMapCreate,
		UpdateContext: resourceRouteMapUpdate,
		ReadContext:   resourceRouteMapRead,
		DeleteContext: resourceRouteMapDelete,
		Schema: map[string]*schema.Schema{
			"tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
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
			"match": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_path": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"community": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"exact_match": {
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
						"extcommunity": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"extcommunity_l_name": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"exact_match": {
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
						"group": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"ha_state": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"scaleout": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"operational_state": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"interface": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"loopback": {
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
									"tunnel": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"local_preference": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"val": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"origin": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"egp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"igp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"incomplete": {
										Type:        schema.TypeInt,
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
									"address": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"acl2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"prefix_list": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
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
									"next_hop": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"acl2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"prefix_list_1": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
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
									"peer": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"acl2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"rib": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"exact": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"reachable": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"unreachable": {
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
						"ipv6": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_1": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"prefix_list_2": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
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
									"next_hop_1": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"next_hop_acl_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"v6_addr": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"prefix_list_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"peer_1": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"acl1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"acl2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"rib": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"exact": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"reachable": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"unreachable": {
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
						"metric": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"route_type": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"external": {
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
											},
										},
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
									"value": {
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
					},
				},
			},
			"set": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"next_hop": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
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
						"ddos": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"class_list_cid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"zone": {
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
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"next_hop_1": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"local": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
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
						"level": {
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
								},
							},
						},
						"metric": {
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
								},
							},
						},
						"metric_type": {
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
								},
							},
						},
						"tag": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"aggregator": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aggregator_as": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"asn": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"ip": {
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
						"as_path": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prepend": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"num": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"num2": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"atomic_aggregate": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"comm_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v_std": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"delete": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_exp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v_exp_delete": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"name_delete": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"community": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dampening_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dampening": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_half_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_reuse": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_supress": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_max_supress": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"dampening_penalty": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"extcommunity": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rt": {
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
											},
										},
									},
									"soo": {
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
											},
										},
									},
								},
							},
						},
						"local_preference": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"val": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"originator_id": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"originator_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"weight": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"weight_val": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"origin": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"egp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"igp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"incomplete": {
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
					},
				},
			},
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouteMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouteMap (Inside resourceRouteMapCreate) ")
		name1 := d.Get("tag").(string)
		name3 := d.Get("sequence").(int)
		name2 := d.Get("action").(string)
		data := dataToRouteMap(d)
		logger.Println("[INFO] received formatted data from method data to RouteMap --")
		d.SetId(name1 + "," + name2 + "," + strconv.Itoa(name3))
		err := go_thunder.PostRouteMap(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouteMapRead(ctx, d, meta)

	}
	return diags
}

func resourceRouteMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouteMap (Inside resourceRouteMapRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouteMap(client.Token, name1, name2, name3, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceRouteMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Modifying RouteMap   (Inside resourceRouteMapUpdate) ")
		data := dataToRouteMap(d)
		logger.Println("[INFO] received formatted data from method data to RouteMap ")
		err := go_thunder.PutRouteMap(client.Token, name1, name2, name3, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouteMapRead(ctx, d, meta)

	}
	return diags
}

func resourceRouteMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		name3 := id[2]
		logger.Println("[INFO] Deleting instance (Inside resourceRouteMapDelete) " + name1)
		err := go_thunder.DeleteRouteMap(client.Token, name1, name2, name3, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToRouteMap(d *schema.ResourceData) go_thunder.RouteMap {
	var vc go_thunder.RouteMap
	var c go_thunder.RouteMapInstance
	c.Tag = d.Get("tag").(string)
	c.Action = d.Get("action").(string)
	c.Sequence = d.Get("sequence").(int)
	c.UserTag = d.Get("user_tag").(string)

	var obj1 go_thunder.RouteMapMatch
	prefix1 := "match.0."

	var obj1_1 go_thunder.RouteMapAsPath1
	prefix1_1 := prefix1 + "as_path.0."
	obj1_1.Name = d.Get(prefix1_1 + "name").(string)

	obj1.Name = obj1_1

	var obj1_2 go_thunder.RouteMapCommunity
	prefix1_2 := prefix1 + "community.0."

	var obj1_2_1 go_thunder.RouteMapNameCfg
	prefix1_2_1 := prefix1_2 + "name_cfg.0."
	obj1_2_1.Name = d.Get(prefix1_2_1 + "name").(string)
	obj1_2_1.ExactMatch = d.Get(prefix1_2_1 + "exact_match").(int)

	obj1_2.Name = obj1_2_1

	obj1.NameCfg = obj1_2

	var obj1_3 go_thunder.RouteMapExtcommunity
	prefix1_3 := prefix1 + "extcommunity.0."

	var obj1_3_1 go_thunder.RouteMapExtcommunityLName
	prefix1_3_1 := prefix1_3 + "extcommunity_l_name.0."
	obj1_3_1.ExtcommunityLName = d.Get(prefix1_3_1 + "name").(string)
	obj1_3_1.ExactMatch = d.Get(prefix1_3_1 + "exact_match").(int)

	obj1_3.ExtcommunityLName = obj1_3_1

	obj1.ExtcommunityLName = obj1_3

	var obj1_4 go_thunder.RouteMapGroup
	prefix1_4 := prefix1 + "group.0."
	obj1_4.GroupID = d.Get(prefix1_4 + "group_id").(int)
	obj1_4.HaState = d.Get(prefix1_4 + "ha_state").(string)

	obj1.GroupID = obj1_4

	var obj1_5 go_thunder.RouteMapScaleout
	prefix1_5 := prefix1 + "scaleout.0."
	obj1_5.ClusterID = d.Get(prefix1_5 + "cluster_id").(int)
	obj1_5.OperationalState = d.Get(prefix1_5 + "operational_state").(string)

	obj1.ClusterID = obj1_5

	var obj1_6 go_thunder.RouteMapInterface
	prefix1_6 := prefix1 + "interface.0."
	obj1_6.Ethernet = d.Get(prefix1_6 + "ethernet").(int)
	obj1_6.Loopback = d.Get(prefix1_6 + "loopback").(int)
	obj1_6.Trunk = d.Get(prefix1_6 + "trunk").(int)
	obj1_6.Ve = d.Get(prefix1_6 + "ve").(int)
	obj1_6.Tunnel = d.Get(prefix1_6 + "tunnel").(int)

	obj1.Ethernet = obj1_6

	var obj1_7 go_thunder.RouteMapLocalPreference
	prefix1_7 := prefix1 + "local_preference.0."
	obj1_7.Val = d.Get(prefix1_7 + "val").(int)

	obj1.Val = obj1_7

	var obj1_8 go_thunder.RouteMapOrigin
	prefix1_8 := prefix1 + "origin.0."
	obj1_8.Egp = d.Get(prefix1_8 + "egp").(int)
	obj1_8.Igp = d.Get(prefix1_8 + "igp").(int)
	obj1_8.Incomplete = d.Get(prefix1_8 + "incomplete").(int)

	obj1.Egp = obj1_8

	var obj1_9 go_thunder.RouteMapIp1
	prefix1_9 := prefix1 + "ip.0."

	var obj1_9_1 go_thunder.RouteMapAddress
	prefix1_9_1 := prefix1_9 + "address.0."
	obj1_9_1.Acl1 = d.Get(prefix1_9_1 + "acl1").(int)
	obj1_9_1.Acl2 = d.Get(prefix1_9_1 + "acl2").(int)
	obj1_9_1.RouteMapAddressName = d.Get(prefix1_9_1 + "name").(string)

	var obj1_9_1_1 go_thunder.RouteMapPrefixList
	prefix1_9_1_1 := prefix1_9_1 + "prefix_list.0."
	obj1_9_1_1.Name = d.Get(prefix1_9_1_1 + "name").(string)

	obj1_9_1.Name = obj1_9_1_1

	obj1_9.Acl1 = obj1_9_1

	var obj1_9_2 go_thunder.RouteMapNextHopIp
	prefix1_9_2 := prefix1_9 + "next_hop.0."
	obj1_9_2.Acl1 = d.Get(prefix1_9_2 + "acl1").(int)
	obj1_9_2.Acl2 = d.Get(prefix1_9_2 + "acl2").(int)
	obj1_9_2.RouteMapNextHopIpName = d.Get(prefix1_9_2 + "name").(string)

	var obj1_9_2_1 go_thunder.RouteMapPrefixList1
	prefix1_9_2_1 := prefix1_9_2 + "prefix_list_1.0."
	obj1_9_2_1.Name = d.Get(prefix1_9_2_1 + "name").(string)

	obj1_9_2.Name = obj1_9_2_1

	obj1_9.Acl2 = obj1_9_2

	var obj1_9_3 go_thunder.RouteMapPeer
	prefix1_9_3 := prefix1_9 + "peer.0."
	obj1_9_3.Acl1 = d.Get(prefix1_9_3 + "acl1").(int)
	obj1_9_3.Acl2 = d.Get(prefix1_9_3 + "acl2").(int)
	obj1_9_3.Name = d.Get(prefix1_9_3 + "name").(string)

	obj1_9.Name = obj1_9_3

	var obj1_9_4 go_thunder.RouteMapRib
	prefix1_9_4 := prefix1_9 + "rib.0."
	obj1_9_4.Exact = d.Get(prefix1_9_4 + "exact").(string)
	obj1_9_4.Reachable = d.Get(prefix1_9_4 + "reachable").(string)
	obj1_9_4.Unreachable = d.Get(prefix1_9_4 + "unreachable").(string)

	obj1_9.Exact = obj1_9_4

	obj1.Address = obj1_9

	var obj1_10 go_thunder.RouteMapIpv6
	prefix1_10 := prefix1 + "ipv6.0."

	var obj1_10_1 go_thunder.RouteMapAddress1
	prefix1_10_1 := prefix1_10 + "address_1.0."
	obj1_10_1.RouteMapAddress1Name = d.Get(prefix1_10_1 + "name").(string)

	var obj1_10_1_1 go_thunder.RouteMapPrefixList2
	prefix1_10_1_1 := prefix1_10_1 + "prefix_list_2.0."
	obj1_10_1_1.Name = d.Get(prefix1_10_1_1 + "name").(string)

	obj1_10_1.Name = obj1_10_1_1

	obj1_10.Name = obj1_10_1

	var obj1_10_2 go_thunder.RouteMapNextHopIpv6
	prefix1_10_2 := prefix1_10 + "next_hop_1.0."
	obj1_10_2.NextHopAclName = d.Get(prefix1_10_2 + "next_hop_acl_name").(string)
	obj1_10_2.V6Addr = d.Get(prefix1_10_2 + "v6_addr").(string)
	obj1_10_2.PrefixListName = d.Get(prefix1_10_2 + "prefix_list_name").(string)

	obj1_10.NextHopAclName = obj1_10_2

	var obj1_10_3 go_thunder.RouteMapPeer
	prefix1_10_3 := prefix1_10 + "peer_1.0."
	obj1_10_3.Acl1 = d.Get(prefix1_10_3 + "acl1").(int)
	obj1_10_3.Acl2 = d.Get(prefix1_10_3 + "acl2").(int)
	obj1_10_3.Name = d.Get(prefix1_10_3 + "name").(string)

	obj1_10.Acl1 = obj1_10_3

	var obj1_10_4 go_thunder.RouteMapRib
	prefix1_10_4 := prefix1_10 + "rib.0."
	obj1_10_4.Exact = d.Get(prefix1_10_4 + "exact").(string)
	obj1_10_4.Reachable = d.Get(prefix1_10_4 + "reachable").(string)
	obj1_10_4.Unreachable = d.Get(prefix1_10_4 + "unreachable").(string)

	obj1_10.Exact = obj1_10_4

	obj1.Address1 = obj1_10

	var obj1_11 go_thunder.RouteMapMetricMap
	prefix1_11 := prefix1 + "metric.0."
	obj1_11.Value = d.Get(prefix1_11 + "value").(int)

	obj1.Value = obj1_11

	var obj1_12 go_thunder.RouteMapRouteType
	prefix1_12 := prefix1 + "route_type.0."

	var obj1_12_1 go_thunder.RouteMapExternal
	prefix1_12_1 := prefix1_12 + "external.0."
	obj1_12_1.ValueExternal = d.Get(prefix1_12_1 + "value").(string)

	obj1_12.ValueExternal = obj1_12_1

	obj1.ValueExternal = obj1_12

	var obj1_13 go_thunder.RouteMapTag
	prefix1_13 := prefix1 + "tag.0."
	obj1_13.ValueTag = d.Get(prefix1_13 + "value").(int)

	obj1.ValueTag = obj1_13

	c.Name = obj1

	var obj2 go_thunder.RouteMapSet
	prefix2 := "set.0."

	var obj2_1 go_thunder.RouteMapIp2
	prefix2_1 := prefix2 + "ip.0."

	var obj2_1_1 go_thunder.RouteMapNextHopSetIp
	prefix2_1_1 := prefix2_1 + "next_hop.0."
	obj2_1_1.Address = d.Get(prefix2_1_1 + "address").(string)

	obj2_1.Address = obj2_1_1

	obj2.Address = obj2_1

	var obj2_2 go_thunder.RouteMapDdos
	prefix2_2 := prefix2 + "ddos.0."
	obj2_2.ClassListName = d.Get(prefix2_2 + "class_list_name").(string)
	obj2_2.ClassListCid = d.Get(prefix2_2 + "class_list_cid").(int)
	obj2_2.Zone = d.Get(prefix2_2 + "zone").(string)

	obj2.ClassListName = obj2_2

	var obj2_3 go_thunder.RouteMapIpv62
	prefix2_3 := prefix2 + "ipv6.0."

	var obj2_3_1 go_thunder.RouteMapNextHop1
	prefix2_3_1 := prefix2_3 + "next_hop_1.0."
	obj2_3_1.RouteMapNextHop1Address = d.Get(prefix2_3_1 + "address").(string)

	var obj2_3_1_1 go_thunder.RouteMapLocal
	prefix2_3_1_1 := prefix2_3_1 + "local.0."
	obj2_3_1_1.Address = d.Get(prefix2_3_1_1 + "address").(string)

	obj2_3_1.Address = obj2_3_1_1

	obj2_3.RouteMapNextHop1Address = obj2_3_1

	obj2.RouteMapNextHop1Address = obj2_3

	var obj2_4 go_thunder.RouteMapLevel
	prefix2_4 := prefix2 + "level.0."
	obj2_4.ValueLevel1 = d.Get(prefix2_4 + "value").(string)

	obj2.ValueLevel1 = obj2_4

	var obj2_5 go_thunder.RouteMapMetric
	prefix2_5 := prefix2 + "metric.0."
	obj2_5.Value = d.Get(prefix2_5 + "value").(string)
	obj2_5.ValueMetric = d.Get(prefix2_5 + "value").(string)
	obj2_5.ValueLevel2 = d.Get(prefix2_5 + "value").(string)

	obj2.ValueLevel2 = obj2_5

	var obj2_6 go_thunder.RouteMapMetricType
	prefix2_6 := prefix2 + "metric_type.0."
	obj2_6.ValueLevel3 = d.Get(prefix2_6 + "value").(string)

	obj2.ValueLevel3 = obj2_6

	var obj2_7 go_thunder.RouteMapTag
	prefix2_7 := prefix2 + "tag.0."
	obj2_7.ValueTag = d.Get(prefix2_7 + "value").(int)

	obj2.ValueTag = obj2_7

	var obj2_8 go_thunder.RouteMapAggregator
	prefix2_8 := prefix2 + "aggregator.0."

	var obj2_8_1 go_thunder.RouteMapAggregatorAs
	prefix2_8_1 := prefix2_8 + "aggregator_as.0."
	obj2_8_1.Asn = d.Get(prefix2_8_1 + "asn").(int)
	obj2_8_1.IP = d.Get(prefix2_8_1 + "ip").(string)

	obj2_8.Asn = obj2_8_1

	obj2.AggregatorAs = obj2_8

	var obj2_9 go_thunder.RouteMapAsPath
	prefix2_9 := prefix2 + "as_path.0."
	obj2_9.Prepend = d.Get(prefix2_9 + "prepend").(string)
	obj2_9.Num = d.Get(prefix2_9 + "num").(int)
	obj2_9.Num2 = d.Get(prefix2_9 + "num2").(int)

	obj2.Prepend = obj2_9

	obj2.AtomicAggregate = d.Get(prefix2 + "atomic_aggregate").(int)

	var obj2_10 go_thunder.RouteMapCommList
	prefix2_10 := prefix2 + "comm_list.0."
	obj2_10.VStd = d.Get(prefix2_10 + "v_std").(int)
	obj2_10.Delete = d.Get(prefix2_10 + "delete").(int)
	obj2_10.VExp = d.Get(prefix2_10 + "v_exp").(int)
	obj2_10.VExpDelete = d.Get(prefix2_10 + "v_exp_delete").(int)
	obj2_10.Name = d.Get(prefix2_10 + "name").(string)
	obj2_10.NameDelete = d.Get(prefix2_10 + "name_delete").(int)

	obj2.VStd = obj2_10

	obj2.Community = d.Get(prefix2 + "community").(string)

	var obj2_11 go_thunder.RouteMapDampeningCfg
	prefix2_11 := prefix2 + "dampening_cfg.0."
	obj2_11.Dampening = d.Get(prefix2_11 + "dampening").(int)
	obj2_11.DampeningHalfTime = d.Get(prefix2_11 + "dampening_half_time").(int)
	obj2_11.DampeningReuse = d.Get(prefix2_11 + "dampening_reuse").(int)
	obj2_11.DampeningSupress = d.Get(prefix2_11 + "dampening_supress").(int)
	obj2_11.DampeningMaxSupress = d.Get(prefix2_11 + "dampening_max_supress").(int)
	obj2_11.DampeningPenalty = d.Get(prefix2_11 + "dampening_penalty").(int)

	obj2.Dampening = obj2_11

	var obj2_12 go_thunder.RouteMapExtcommunity1
	prefix2_12 := prefix2 + "extcommunity.0."

	var obj2_12_1 go_thunder.RouteMapRt
	prefix2_12_1 := prefix2_12 + "rt.0."
	obj2_12_1.ValueRt = d.Get(prefix2_12_1 + "value").(string)

	obj2_12.ValueRt = obj2_12_1

	var obj2_12_2 go_thunder.RouteMapSoo
	prefix2_12_2 := prefix2_12 + "soo.0."
	obj2_12_2.ValueSoo = d.Get(prefix2_12_2 + "value").(string)

	obj2_12.ValueSoo = obj2_12_2

	obj2.ValueRt = obj2_12

	var obj2_13 go_thunder.RouteMapLocalPreference
	prefix2_13 := prefix2 + "local_preference.0."
	obj2_13.Val = d.Get(prefix2_13 + "val").(int)

	obj2.Val = obj2_13

	var obj2_14 go_thunder.RouteMapOriginatorID
	prefix2_14 := prefix2 + "originator_id.0."
	obj2_14.OriginatorIP = d.Get(prefix2_14 + "originator_ip").(string)

	obj2.OriginatorIP = obj2_14

	var obj2_15 go_thunder.RouteMapWeight
	prefix2_15 := prefix2 + "weight.0."
	obj2_15.WeightVal = d.Get(prefix2_15 + "weight_val").(int)

	obj2.WeightVal = obj2_15

	var obj2_16 go_thunder.RouteMapOrigin
	prefix2_16 := prefix2 + "origin.0."
	obj2_16.Egp = d.Get(prefix2_16 + "egp").(int)
	obj2_16.Igp = d.Get(prefix2_16 + "igp").(int)
	obj2_16.Incomplete = d.Get(prefix2_16 + "incomplete").(int)

	obj2.Egp = obj2_16

	c.AggregatorAs = obj2

	vc.Tag = c
	return vc
}
