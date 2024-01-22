package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitoredEntityOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitored_entity_oper`: Operational Status for the object monitored-entity\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitoredEntityOperRead,

		Schema: map[string]*schema.Schema{
			"detail": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"primary_keys": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"all_keys": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mon_entity_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"entity_key": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "",
												},
												"flat_oid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_proto": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ha_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"entity_metric_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"current": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"threshold": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"anomaly": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"sec_entity_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"entity_key": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "",
															},
															"flat_oid": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"ipv4_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"ipv6_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_proto": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"mode": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"ha_state": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"entity_metric_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"metric_name": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"current": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"threshold": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"anomaly": {
																			Type: schema.TypeString, Optional: true, Description: "",
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
						"debug": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_keys": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"all_keys": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mon_entity_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"entity_key": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"uuid": {
																Type: schema.TypeString, Optional: true, Computed: true, Description: "",
															},
															"flat_oid": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"ipv4_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"ipv6_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_proto": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"mode": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"ha_state": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"entity_metric_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"metric_name": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"current": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"min": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"max": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"mean": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"threshold": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"std_dev": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"anomaly": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																	},
																},
															},
															"sec_entity_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"entity_key": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"uuid": {
																			Type: schema.TypeString, Optional: true, Computed: true, Description: "",
																		},
																		"flat_oid": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																		"ipv4_addr": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"ipv6_addr": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"l4_proto": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"l4_port": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																		"mode": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"ha_state": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"entity_metric_list": {
																			Type: schema.TypeList, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"metric_name": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"current": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"min": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"max": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"mean": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"threshold": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"std_dev": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"anomaly": {
																						Type: schema.TypeString, Optional: true, Description: "",
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
							},
						},
					},
				},
			},
			"mon_topk": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric_topk_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"metric_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"topk_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"protocol": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"metric_value": {
																Type: schema.TypeString, Optional: true, Description: "",
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
						"sources": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_proto": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"metric_topk_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"topk_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ip_addr": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"metric_value": {
																			Type: schema.TypeString, Optional: true, Description: "",
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
				},
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_keys": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"all_keys": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mon_entity_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entity_key": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "",
									},
									"flat_oid": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"l4_proto": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"l4_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ha_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sec_entity_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"entity_key": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "",
												},
												"flat_oid": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_proto": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"mode": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ha_state": {
													Type: schema.TypeString, Optional: true, Description: "",
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
			"secondary": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"mon_topk": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oper": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_proto": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"metric_topk_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"metric_name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"topk_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ip_addr": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"port": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																		"protocol": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"metric_value": {
																			Type: schema.TypeString, Optional: true, Description: "",
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
									"sources": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"oper": {
													Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv4_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"ipv6_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_proto": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"metric_topk_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"metric_name": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"topk_list": {
																			Type: schema.TypeList, Optional: true, Description: "",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"ip_addr": {
																						Type: schema.TypeString, Optional: true, Description: "",
																					},
																					"metric_value": {
																						Type: schema.TypeString, Optional: true, Description: "",
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
							},
						},
					},
				},
			},
			"sessions": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mon_entity_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"entity_key": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv4_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ipv6_addr": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_proto": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"l4_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"session_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"proto": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"fwd_src_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"fwd_src_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"fwd_dst_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"fwd_dst_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"rev_src_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"rev_src_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"rev_dst_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"rev_dst_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
														},
													},
												},
												"sec_entity_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"entity_key": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"ipv4_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"ipv6_addr": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_proto": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"l4_port": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"session_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"proto": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"fwd_src_ip": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"fwd_src_port": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																		"fwd_dst_ip": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"fwd_dst_port": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																		"rev_src_ip": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"rev_src_port": {
																			Type: schema.TypeInt, Optional: true, Description: "",
																		},
																		"rev_dst_ip": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																		"rev_dst_port": {
																			Type: schema.TypeInt, Optional: true, Description: "",
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
				},
			},
		},
	}
}

func resourceVisibilityMonitoredEntityOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitoredEntityOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitoredEntityOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitoredEntityOperDetail := setObjectVisibilityMonitoredEntityOperDetail(res)
		d.Set("detail", VisibilityMonitoredEntityOperDetail)
		VisibilityMonitoredEntityOperMonTopk := setObjectVisibilityMonitoredEntityOperMonTopk(res)
		d.Set("mon_topk", VisibilityMonitoredEntityOperMonTopk)
		VisibilityMonitoredEntityOperOper := setObjectVisibilityMonitoredEntityOperOper(res)
		d.Set("oper", VisibilityMonitoredEntityOperOper)
		VisibilityMonitoredEntityOperSecondary := setObjectVisibilityMonitoredEntityOperSecondary(res)
		d.Set("secondary", VisibilityMonitoredEntityOperSecondary)
		VisibilityMonitoredEntityOperSessions := setObjectVisibilityMonitoredEntityOperSessions(res)
		d.Set("sessions", VisibilityMonitoredEntityOperSessions)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitoredEntityOperDetail(ret edpt.DataVisibilityMonitoredEntityOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":  setObjectVisibilityMonitoredEntityOperDetailOper(ret.DtVisibilityMonitoredEntityOper.Detail.Oper),
			"debug": setObjectVisibilityMonitoredEntityOperDetailDebug(ret.DtVisibilityMonitoredEntityOper.Detail.Debug),
		},
	}
}

func setObjectVisibilityMonitoredEntityOperDetailOper(d edpt.VisibilityMonitoredEntityOperDetailOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["primary_keys"] = d.PrimaryKeys

	in["all_keys"] = d.AllKeys
	in["mon_entity_list"] = setSliceVisibilityMonitoredEntityOperDetailOperMonEntityList(d.MonEntityList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailOperMonEntityList(d []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList(item.EntityMetricList)
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["threshold"] = item.Threshold
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList(item.EntityMetricList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["threshold"] = item.Threshold
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityOperDetailDebug(d edpt.VisibilityMonitoredEntityOperDetailDebug) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectVisibilityMonitoredEntityOperDetailDebugOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectVisibilityMonitoredEntityOperDetailDebugOper(d edpt.VisibilityMonitoredEntityOperDetailDebugOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["primary_keys"] = d.PrimaryKeys

	in["all_keys"] = d.AllKeys
	in["mon_entity_list"] = setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityList(d.MonEntityList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityList(d []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList(item.EntityMetricList)
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["min"] = item.Min
		in["max"] = item.Max
		in["mean"] = item.Mean
		in["threshold"] = item.Threshold
		in["std_dev"] = item.StdDev
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		in["entity_metric_list"] = setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList(item.EntityMetricList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList(d []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["current"] = item.Current
		in["min"] = item.Min
		in["max"] = item.Max
		in["mean"] = item.Mean
		in["threshold"] = item.Threshold
		in["std_dev"] = item.StdDev
		in["anomaly"] = item.Anomaly
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityOperMonTopk(ret edpt.DataVisibilityMonitoredEntityOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":    setObjectVisibilityMonitoredEntityOperMonTopkOper(ret.DtVisibilityMonitoredEntityOper.MonTopk.Oper),
			"sources": setObjectVisibilityMonitoredEntityOperMonTopkSources(ret.DtVisibilityMonitoredEntityOper.MonTopk.Sources),
		},
	}
}

func setObjectVisibilityMonitoredEntityOperMonTopkOper(d edpt.VisibilityMonitoredEntityOperMonTopkOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["metric_topk_list"] = setSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkList(d.MetricTopkList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkList(d []edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["port"] = item.Port
		in["protocol"] = item.Protocol
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityOperMonTopkSources(d edpt.VisibilityMonitoredEntityOperMonTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectVisibilityMonitoredEntityOperMonTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectVisibilityMonitoredEntityOperMonTopkSourcesOper(d edpt.VisibilityMonitoredEntityOperMonTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ipv4_addr"] = d.Ipv4Addr

	in["ipv6_addr"] = d.Ipv6Addr

	in["l4_proto"] = d.L4Proto

	in["l4_port"] = d.L4Port
	in["metric_topk_list"] = setSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList(d.MetricTopkList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList(d []edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityOperOper(ret edpt.DataVisibilityMonitoredEntityOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"primary_keys":    ret.DtVisibilityMonitoredEntityOper.Oper.PrimaryKeys,
			"all_keys":        ret.DtVisibilityMonitoredEntityOper.Oper.AllKeys,
			"mon_entity_list": setSliceVisibilityMonitoredEntityOperOperMonEntityList(ret.DtVisibilityMonitoredEntityOper.Oper.MonEntityList),
		},
	}
}

func setSliceVisibilityMonitoredEntityOperOperMonEntityList(d []edpt.VisibilityMonitoredEntityOperOperMonEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntityOperOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntityOperOperMonEntityListSecEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		//omit uuid
		in["flat_oid"] = item.FlatOid
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["mode"] = item.Mode
		in["ha_state"] = item.HaState
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityOperSecondary(ret edpt.DataVisibilityMonitoredEntityOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper":     setObjectVisibilityMonitoredEntityOperSecondaryOper(ret.DtVisibilityMonitoredEntityOper.Secondary.Oper),
			"mon_topk": setObjectVisibilityMonitoredEntityOperSecondaryMonTopk(ret.DtVisibilityMonitoredEntityOper.Secondary.MonTopk),
		},
	}
}

func setObjectVisibilityMonitoredEntityOperSecondaryOper(d edpt.VisibilityMonitoredEntityOperSecondaryOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	result = append(result, in)
	return result
}

func setObjectVisibilityMonitoredEntityOperSecondaryMonTopk(d edpt.VisibilityMonitoredEntityOperSecondaryMonTopk) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectVisibilityMonitoredEntityOperSecondaryMonTopkOper(d.Oper)
	in["sources"] = setObjectVisibilityMonitoredEntityOperSecondaryMonTopkSources(d.Sources)
	result = append(result, in)
	return result
}

func setObjectVisibilityMonitoredEntityOperSecondaryMonTopkOper(d edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ipv4_addr"] = d.Ipv4Addr

	in["ipv6_addr"] = d.Ipv6Addr

	in["l4_proto"] = d.L4Proto

	in["l4_port"] = d.L4Port
	in["metric_topk_list"] = setSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList(d.MetricTopkList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList(d []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["port"] = item.Port
		in["protocol"] = item.Protocol
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityOperSecondaryMonTopkSources(d edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSources) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper(d.Oper)
	result = append(result, in)
	return result
}

func setObjectVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper(d edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ipv4_addr"] = d.Ipv4Addr

	in["ipv6_addr"] = d.Ipv6Addr

	in["l4_proto"] = d.L4Proto

	in["l4_port"] = d.L4Port
	in["metric_topk_list"] = setSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList(d.MetricTopkList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList(d []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList(d []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_addr"] = item.IpAddr
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityMonitoredEntityOperSessions(ret edpt.DataVisibilityMonitoredEntityOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVisibilityMonitoredEntityOperSessionsOper(ret.DtVisibilityMonitoredEntityOper.Sessions.Oper),
		},
	}
}

func setObjectVisibilityMonitoredEntityOperSessionsOper(d edpt.VisibilityMonitoredEntityOperSessionsOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["mon_entity_list"] = setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityList(d.MonEntityList)
	result = append(result, in)
	return result
}

func setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityList(d []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["session_list"] = setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList(item.SessionList)
		in["sec_entity_list"] = setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList(item.SecEntityList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList(d []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["proto"] = item.Proto
		in["fwd_src_ip"] = item.FwdSrcIp
		in["fwd_src_port"] = item.FwdSrcPort
		in["fwd_dst_ip"] = item.FwdDstIp
		in["fwd_dst_port"] = item.FwdDstPort
		in["rev_src_ip"] = item.RevSrcIp
		in["rev_src_port"] = item.RevSrcPort
		in["rev_dst_ip"] = item.RevDstIp
		in["rev_dst_port"] = item.RevDstPort
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList(d []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entity_key"] = item.EntityKey
		in["ipv4_addr"] = item.Ipv4Addr
		in["ipv6_addr"] = item.Ipv6Addr
		in["l4_proto"] = item.L4Proto
		in["l4_port"] = item.L4Port
		in["session_list"] = setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList(item.SessionList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList(d []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["proto"] = item.Proto
		in["fwd_src_ip"] = item.FwdSrcIp
		in["fwd_src_port"] = item.FwdSrcPort
		in["fwd_dst_ip"] = item.FwdDstIp
		in["fwd_dst_port"] = item.FwdDstPort
		in["rev_src_ip"] = item.RevSrcIp
		in["rev_src_port"] = item.RevSrcPort
		in["rev_dst_ip"] = item.RevDstIp
		in["rev_dst_port"] = item.RevDstPort
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityMonitoredEntityOperDetail(d []interface{}) edpt.VisibilityMonitoredEntityOperDetail {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperDetail
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperDetailOper(in["oper"].([]interface{}))
		ret.Debug = getObjectVisibilityMonitoredEntityOperDetailDebug(in["debug"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperDetailOper(d []interface{}) edpt.VisibilityMonitoredEntityOperDetailOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperDetailOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrimaryKeys = in["primary_keys"].(int)
		ret.AllKeys = in["all_keys"].(int)
		ret.MonEntityList = getSliceVisibilityMonitoredEntityOperDetailOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		oi.SecEntityList = getSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailOperMonEntityListSecEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperDetailDebug(d []interface{}) edpt.VisibilityMonitoredEntityOperDetailDebug {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperDetailDebug
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperDetailDebugOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperDetailDebugOper(d []interface{}) edpt.VisibilityMonitoredEntityOperDetailDebugOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperDetailDebugOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrimaryKeys = in["primary_keys"].(int)
		ret.AllKeys = in["all_keys"].(int)
		ret.MonEntityList = getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		oi.SecEntityList = getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Min = in["min"].(string)
		oi.Max = in["max"].(string)
		oi.Mean = in["mean"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.StdDev = in["std_dev"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.EntityMetricList = getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList(in["entity_metric_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList(d []interface{}) []edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperDetailDebugOperMonEntityListSecEntityListEntityMetricList
		oi.MetricName = in["metric_name"].(string)
		oi.Current = in["current"].(string)
		oi.Min = in["min"].(string)
		oi.Max = in["max"].(string)
		oi.Mean = in["mean"].(string)
		oi.Threshold = in["threshold"].(string)
		oi.StdDev = in["std_dev"].(string)
		oi.Anomaly = in["anomaly"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperMonTopk(d []interface{}) edpt.VisibilityMonitoredEntityOperMonTopk {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperMonTopk
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperMonTopkOper(in["oper"].([]interface{}))
		ret.Sources = getObjectVisibilityMonitoredEntityOperMonTopkSources(in["sources"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperMonTopkOper(d []interface{}) edpt.VisibilityMonitoredEntityOperMonTopkOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperMonTopkOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MetricTopkList = getSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperMonTopkOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperMonTopkSources(d []interface{}) edpt.VisibilityMonitoredEntityOperMonTopkSources {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperMonTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperMonTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperMonTopkSourcesOper(d []interface{}) edpt.VisibilityMonitoredEntityOperMonTopkSourcesOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperMonTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperMonTopkSourcesOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperOper(d []interface{}) edpt.VisibilityMonitoredEntityOperOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrimaryKeys = in["primary_keys"].(int)
		ret.AllKeys = in["all_keys"].(int)
		ret.MonEntityList = getSliceVisibilityMonitoredEntityOperOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		oi.SecEntityList = getSliceVisibilityMonitoredEntityOperOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		//omit uuid
		oi.FlatOid = in["flat_oid"].(int)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.Mode = in["mode"].(string)
		oi.HaState = in["ha_state"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperSecondary(d []interface{}) edpt.VisibilityMonitoredEntityOperSecondary {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperSecondary
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperSecondaryOper(in["oper"].([]interface{}))
		ret.MonTopk = getObjectVisibilityMonitoredEntityOperSecondaryMonTopk(in["mon_topk"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperSecondaryOper(d []interface{}) edpt.VisibilityMonitoredEntityOperSecondaryOper {

	var ret edpt.VisibilityMonitoredEntityOperSecondaryOper
	return ret
}

func getObjectVisibilityMonitoredEntityOperSecondaryMonTopk(d []interface{}) edpt.VisibilityMonitoredEntityOperSecondaryMonTopk {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperSecondaryMonTopk
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperSecondaryMonTopkOper(in["oper"].([]interface{}))
		ret.Sources = getObjectVisibilityMonitoredEntityOperSecondaryMonTopkSources(in["sources"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperSecondaryMonTopkOper(d []interface{}) edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSecondaryMonTopkOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.Port = in["port"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperSecondaryMonTopkSources(d []interface{}) edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSources {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper(d []interface{}) edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.L4Proto = in["l4_proto"].(string)
		ret.L4Port = in["l4_port"].(int)
		ret.MetricTopkList = getSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSecondaryMonTopkSourcesOperMetricTopkListTopkList
		oi.IpAddr = in["ip_addr"].(string)
		oi.MetricValue = in["metric_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperSessions(d []interface{}) edpt.VisibilityMonitoredEntityOperSessions {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperSessions
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityMonitoredEntityOperSessionsOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityMonitoredEntityOperSessionsOper(d []interface{}) edpt.VisibilityMonitoredEntityOperSessionsOper {

	count1 := len(d)
	var ret edpt.VisibilityMonitoredEntityOperSessionsOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonEntityList = getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityList(in["mon_entity_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityList
		oi.EntityKey = in["entity_key"].(string)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.SessionList = getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList(in["session_list"].([]interface{}))
		oi.SecEntityList = getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList(in["sec_entity_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSessionList
		oi.Proto = in["proto"].(string)
		oi.FwdSrcIp = in["fwd_src_ip"].(string)
		oi.FwdSrcPort = in["fwd_src_port"].(int)
		oi.FwdDstIp = in["fwd_dst_ip"].(string)
		oi.FwdDstPort = in["fwd_dst_port"].(int)
		oi.RevSrcIp = in["rev_src_ip"].(string)
		oi.RevSrcPort = in["rev_src_port"].(int)
		oi.RevDstIp = in["rev_dst_ip"].(string)
		oi.RevDstPort = in["rev_dst_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityList
		oi.EntityKey = in["entity_key"].(string)
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.L4Proto = in["l4_proto"].(string)
		oi.L4Port = in["l4_port"].(int)
		oi.SessionList = getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList(in["session_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList(d []interface{}) []edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitoredEntityOperSessionsOperMonEntityListSecEntityListSessionList
		oi.Proto = in["proto"].(string)
		oi.FwdSrcIp = in["fwd_src_ip"].(string)
		oi.FwdSrcPort = in["fwd_src_port"].(int)
		oi.FwdDstIp = in["fwd_dst_ip"].(string)
		oi.FwdDstPort = in["fwd_dst_port"].(int)
		oi.RevSrcIp = in["rev_src_ip"].(string)
		oi.RevSrcPort = in["rev_src_port"].(int)
		oi.RevDstIp = in["rev_dst_ip"].(string)
		oi.RevDstPort = in["rev_dst_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitoredEntityOper(d *schema.ResourceData) edpt.VisibilityMonitoredEntityOper {
	var ret edpt.VisibilityMonitoredEntityOper

	ret.Detail = getObjectVisibilityMonitoredEntityOperDetail(d.Get("detail").([]interface{}))

	ret.MonTopk = getObjectVisibilityMonitoredEntityOperMonTopk(d.Get("mon_topk").([]interface{}))

	ret.Oper = getObjectVisibilityMonitoredEntityOperOper(d.Get("oper").([]interface{}))

	ret.Secondary = getObjectVisibilityMonitoredEntityOperSecondary(d.Get("secondary").([]interface{}))

	ret.Sessions = getObjectVisibilityMonitoredEntityOperSessions(d.Get("sessions").([]interface{}))
	return ret
}
