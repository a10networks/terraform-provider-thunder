package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_oper`: Operational Status for the object rule-set\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetOperRead,

		Schema: map[string]*schema.Schema{
			"application": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category_stat": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_stat": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rule": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rule_set_only": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rule_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"stat_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"category": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"type": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"conns": {
																Type: schema.TypeInt, Optional: true, Description: "",
															},
															"bytes": {
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule set name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"policy_unmatched_drop": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_permit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_deny": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_reset": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_rule_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rule_stats": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rule_hitcount": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rule_action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rule_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"total_hit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_permit_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_deny_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_reset_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_permit_packets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_deny_packets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_reset_packets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_packets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_active_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_active_udp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_active_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_active_others": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"show_total_stats": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"topn_rules": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Rule name",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hitcount": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"last_hitcount_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"permitbytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"denybytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resetbytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"totalbytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"permitpackets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"denypackets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"resetpackets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"totalpackets": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"activesessiontcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"activesessionudp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"activesessionicmp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"activesessionsctp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"activesessionother": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"activesessiontotal": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sessiontcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sessionudp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sessionicmp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sessionsctp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sessionother": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sessiontotal": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ratelimitdrops": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"rules_by_zone": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"from": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"to": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rule_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"source_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"source": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																	},
																},
															},
															"dest_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dest": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																	},
																},
															},
															"service_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"service": {
																			Type: schema.TypeString, Optional: true, Description: "",
																		},
																	},
																},
															},
															"dscp_list": {
																Type: schema.TypeList, Optional: true, Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dscp": {
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
			"track_app_rule_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
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
	}
}

func resourceRuleSetOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetOperApplication := setObjectRuleSetOperApplication(res)
		d.Set("application", RuleSetOperApplication)
		RuleSetOperOper := setObjectRuleSetOperOper(res)
		d.Set("oper", RuleSetOperOper)
		RuleSetOperRuleList := setSliceRuleSetOperRuleList(res)
		d.Set("rule_list", RuleSetOperRuleList)
		RuleSetOperRulesByZone := setObjectRuleSetOperRulesByZone(res)
		d.Set("rules_by_zone", RuleSetOperRulesByZone)
		RuleSetOperTrackAppRuleList := setObjectRuleSetOperTrackAppRuleList(res)
		d.Set("track_app_rule_list", RuleSetOperTrackAppRuleList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetOperApplication(ret edpt.DataRuleSetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectRuleSetOperApplicationOper(ret.DtRuleSetOper.Application.Oper),
		},
	}
}

func setObjectRuleSetOperApplicationOper(d edpt.RuleSetOperApplicationOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["category_stat"] = d.CategoryStat

	in["app_stat"] = d.AppStat

	in["protocol"] = d.Protocol

	in["rule"] = d.Rule

	in["rule_set_only"] = d.RuleSetOnly
	in["rule_list"] = setSliceRuleSetOperApplicationOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceRuleSetOperApplicationOperRuleList(d []edpt.RuleSetOperApplicationOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stat_list"] = setSliceRuleSetOperApplicationOperRuleListStatList(item.StatList)
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetOperApplicationOperRuleListStatList(d []edpt.RuleSetOperApplicationOperRuleListStatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["category"] = item.Category
		in["type"] = item.Type
		in["conns"] = item.Conns
		in["bytes"] = item.Bytes
		result = append(result, in)
	}
	return result
}

func setObjectRuleSetOperOper(ret edpt.DataRuleSetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"policy_status":         ret.DtRuleSetOper.Oper.PolicyStatus,
			"policy_unmatched_drop": ret.DtRuleSetOper.Oper.PolicyUnmatchedDrop,
			"policy_permit":         ret.DtRuleSetOper.Oper.PolicyPermit,
			"policy_deny":           ret.DtRuleSetOper.Oper.PolicyDeny,
			"policy_reset":          ret.DtRuleSetOper.Oper.PolicyReset,
			"policy_rule_count":     ret.DtRuleSetOper.Oper.PolicyRuleCount,
			"rule_stats":            setSliceRuleSetOperOperRuleStats(ret.DtRuleSetOper.Oper.RuleStats),
			"total_hit":             ret.DtRuleSetOper.Oper.TotalHit,
			"total_permit_bytes":    ret.DtRuleSetOper.Oper.TotalPermitBytes,
			"total_deny_bytes":      ret.DtRuleSetOper.Oper.TotalDenyBytes,
			"total_reset_bytes":     ret.DtRuleSetOper.Oper.TotalResetBytes,
			"total_bytes":           ret.DtRuleSetOper.Oper.TotalBytes,
			"total_permit_packets":  ret.DtRuleSetOper.Oper.TotalPermitPackets,
			"total_deny_packets":    ret.DtRuleSetOper.Oper.TotalDenyPackets,
			"total_reset_packets":   ret.DtRuleSetOper.Oper.TotalResetPackets,
			"total_packets":         ret.DtRuleSetOper.Oper.TotalPackets,
			"total_active_tcp":      ret.DtRuleSetOper.Oper.TotalActiveTcp,
			"total_active_udp":      ret.DtRuleSetOper.Oper.TotalActiveUdp,
			"total_active_icmp":     ret.DtRuleSetOper.Oper.TotalActiveIcmp,
			"total_active_others":   ret.DtRuleSetOper.Oper.TotalActiveOthers,
			"show_total_stats":      ret.DtRuleSetOper.Oper.ShowTotalStats,
			"topn_rules":            ret.DtRuleSetOper.Oper.TopnRules,
		},
	}
}

func setSliceRuleSetOperOperRuleStats(d []edpt.RuleSetOperOperRuleStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rule_name"] = item.RuleName
		in["rule_hitcount"] = item.RuleHitcount
		in["rule_action"] = item.RuleAction
		in["rule_status"] = item.RuleStatus
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetOperRuleList(d edpt.DataRuleSetOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtRuleSetOper.RuleList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["oper"] = setObjectRuleSetOperRuleListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectRuleSetOperRuleListOper(d edpt.RuleSetOperRuleListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hitcount"] = d.Hitcount

	in["last_hitcount_time"] = d.LastHitcountTime

	in["action"] = d.Action

	in["status"] = d.Status

	in["permitbytes"] = d.Permitbytes

	in["denybytes"] = d.Denybytes

	in["resetbytes"] = d.Resetbytes

	in["totalbytes"] = d.Totalbytes

	in["permitpackets"] = d.Permitpackets

	in["denypackets"] = d.Denypackets

	in["resetpackets"] = d.Resetpackets

	in["totalpackets"] = d.Totalpackets

	in["activesessiontcp"] = d.Activesessiontcp

	in["activesessionudp"] = d.Activesessionudp

	in["activesessionicmp"] = d.Activesessionicmp

	in["activesessionsctp"] = d.Activesessionsctp

	in["activesessionother"] = d.Activesessionother

	in["activesessiontotal"] = d.Activesessiontotal

	in["sessiontcp"] = d.Sessiontcp

	in["sessionudp"] = d.Sessionudp

	in["sessionicmp"] = d.Sessionicmp

	in["sessionsctp"] = d.Sessionsctp

	in["sessionother"] = d.Sessionother

	in["sessiontotal"] = d.Sessiontotal

	in["ratelimitdrops"] = d.Ratelimitdrops
	result = append(result, in)
	return result
}

func setObjectRuleSetOperRulesByZone(ret edpt.DataRuleSetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectRuleSetOperRulesByZoneOper(ret.DtRuleSetOper.RulesByZone.Oper),
		},
	}
}

func setObjectRuleSetOperRulesByZoneOper(d edpt.RuleSetOperRulesByZoneOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["group_list"] = setSliceRuleSetOperRulesByZoneOperGroupList(d.GroupList)
	result = append(result, in)
	return result
}

func setSliceRuleSetOperRulesByZoneOperGroupList(d []edpt.RuleSetOperRulesByZoneOperGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["from"] = item.From
		in["to"] = item.To
		in["rule_list"] = setSliceRuleSetOperRulesByZoneOperGroupListRuleList(item.RuleList)
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetOperRulesByZoneOperGroupListRuleList(d []edpt.RuleSetOperRulesByZoneOperGroupListRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["action"] = item.Action
		in["source_list"] = setSliceRuleSetOperRulesByZoneOperGroupListRuleListSourceList(item.SourceList)
		in["dest_list"] = setSliceRuleSetOperRulesByZoneOperGroupListRuleListDestList(item.DestList)
		in["service_list"] = setSliceRuleSetOperRulesByZoneOperGroupListRuleListServiceList(item.ServiceList)
		in["dscp_list"] = setSliceRuleSetOperRulesByZoneOperGroupListRuleListDscpList(item.DscpList)
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetOperRulesByZoneOperGroupListRuleListSourceList(d []edpt.RuleSetOperRulesByZoneOperGroupListRuleListSourceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["source"] = item.Source
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetOperRulesByZoneOperGroupListRuleListDestList(d []edpt.RuleSetOperRulesByZoneOperGroupListRuleListDestList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dest"] = item.Dest
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetOperRulesByZoneOperGroupListRuleListServiceList(d []edpt.RuleSetOperRulesByZoneOperGroupListRuleListServiceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["service"] = item.Service
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetOperRulesByZoneOperGroupListRuleListDscpList(d []edpt.RuleSetOperRulesByZoneOperGroupListRuleListDscpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dscp"] = item.Dscp
		result = append(result, in)
	}
	return result
}

func setObjectRuleSetOperTrackAppRuleList(ret edpt.DataRuleSetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectRuleSetOperTrackAppRuleListOper(ret.DtRuleSetOper.TrackAppRuleList.Oper),
		},
	}
}

func setObjectRuleSetOperTrackAppRuleListOper(d edpt.RuleSetOperTrackAppRuleListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["rule_list"] = setSliceRuleSetOperTrackAppRuleListOperRuleList(d.RuleList)
	result = append(result, in)
	return result
}

func setSliceRuleSetOperTrackAppRuleListOperRuleList(d []edpt.RuleSetOperTrackAppRuleListOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func getObjectRuleSetOperApplication(d []interface{}) edpt.RuleSetOperApplication {

	count1 := len(d)
	var ret edpt.RuleSetOperApplication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectRuleSetOperApplicationOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectRuleSetOperApplicationOper(d []interface{}) edpt.RuleSetOperApplicationOper {

	count1 := len(d)
	var ret edpt.RuleSetOperApplicationOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CategoryStat = in["category_stat"].(string)
		ret.AppStat = in["app_stat"].(string)
		ret.Protocol = in["protocol"].(int)
		ret.Rule = in["rule"].(string)
		ret.RuleSetOnly = in["rule_set_only"].(int)
		ret.RuleList = getSliceRuleSetOperApplicationOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceRuleSetOperApplicationOperRuleList(d []interface{}) []edpt.RuleSetOperApplicationOperRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperApplicationOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperApplicationOperRuleList
		oi.Name = in["name"].(string)
		oi.StatList = getSliceRuleSetOperApplicationOperRuleListStatList(in["stat_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetOperApplicationOperRuleListStatList(d []interface{}) []edpt.RuleSetOperApplicationOperRuleListStatList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperApplicationOperRuleListStatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperApplicationOperRuleListStatList
		oi.Name = in["name"].(string)
		oi.Category = in["category"].(string)
		oi.Type = in["type"].(string)
		oi.Conns = in["conns"].(int)
		oi.Bytes = in["bytes"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRuleSetOperOper(d []interface{}) edpt.RuleSetOperOper {

	count1 := len(d)
	var ret edpt.RuleSetOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PolicyStatus = in["policy_status"].(string)
		ret.PolicyUnmatchedDrop = in["policy_unmatched_drop"].(int)
		ret.PolicyPermit = in["policy_permit"].(int)
		ret.PolicyDeny = in["policy_deny"].(int)
		ret.PolicyReset = in["policy_reset"].(int)
		ret.PolicyRuleCount = in["policy_rule_count"].(int)
		ret.RuleStats = getSliceRuleSetOperOperRuleStats(in["rule_stats"].([]interface{}))
		ret.TotalHit = in["total_hit"].(int)
		ret.TotalPermitBytes = in["total_permit_bytes"].(int)
		ret.TotalDenyBytes = in["total_deny_bytes"].(int)
		ret.TotalResetBytes = in["total_reset_bytes"].(int)
		ret.TotalBytes = in["total_bytes"].(int)
		ret.TotalPermitPackets = in["total_permit_packets"].(int)
		ret.TotalDenyPackets = in["total_deny_packets"].(int)
		ret.TotalResetPackets = in["total_reset_packets"].(int)
		ret.TotalPackets = in["total_packets"].(int)
		ret.TotalActiveTcp = in["total_active_tcp"].(int)
		ret.TotalActiveUdp = in["total_active_udp"].(int)
		ret.TotalActiveIcmp = in["total_active_icmp"].(int)
		ret.TotalActiveOthers = in["total_active_others"].(int)
		ret.ShowTotalStats = in["show_total_stats"].(string)
		ret.TopnRules = in["topn_rules"].(string)
	}
	return ret
}

func getSliceRuleSetOperOperRuleStats(d []interface{}) []edpt.RuleSetOperOperRuleStats {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperOperRuleStats, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperOperRuleStats
		oi.RuleName = in["rule_name"].(string)
		oi.RuleHitcount = in["rule_hitcount"].(int)
		oi.RuleAction = in["rule_action"].(string)
		oi.RuleStatus = in["rule_status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetOperRuleList(d []interface{}) []edpt.RuleSetOperRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperRuleList
		oi.Name = in["name"].(string)
		oi.Oper = getObjectRuleSetOperRuleListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRuleSetOperRuleListOper(d []interface{}) edpt.RuleSetOperRuleListOper {

	count1 := len(d)
	var ret edpt.RuleSetOperRuleListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hitcount = in["hitcount"].(int)
		ret.LastHitcountTime = in["last_hitcount_time"].(string)
		ret.Action = in["action"].(string)
		ret.Status = in["status"].(string)
		ret.Permitbytes = in["permitbytes"].(int)
		ret.Denybytes = in["denybytes"].(int)
		ret.Resetbytes = in["resetbytes"].(int)
		ret.Totalbytes = in["totalbytes"].(int)
		ret.Permitpackets = in["permitpackets"].(int)
		ret.Denypackets = in["denypackets"].(int)
		ret.Resetpackets = in["resetpackets"].(int)
		ret.Totalpackets = in["totalpackets"].(int)
		ret.Activesessiontcp = in["activesessiontcp"].(int)
		ret.Activesessionudp = in["activesessionudp"].(int)
		ret.Activesessionicmp = in["activesessionicmp"].(int)
		ret.Activesessionsctp = in["activesessionsctp"].(int)
		ret.Activesessionother = in["activesessionother"].(int)
		ret.Activesessiontotal = in["activesessiontotal"].(int)
		ret.Sessiontcp = in["sessiontcp"].(int)
		ret.Sessionudp = in["sessionudp"].(int)
		ret.Sessionicmp = in["sessionicmp"].(int)
		ret.Sessionsctp = in["sessionsctp"].(int)
		ret.Sessionother = in["sessionother"].(int)
		ret.Sessiontotal = in["sessiontotal"].(int)
		ret.Ratelimitdrops = in["ratelimitdrops"].(int)
	}
	return ret
}

func getObjectRuleSetOperRulesByZone(d []interface{}) edpt.RuleSetOperRulesByZone {

	count1 := len(d)
	var ret edpt.RuleSetOperRulesByZone
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectRuleSetOperRulesByZoneOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectRuleSetOperRulesByZoneOper(d []interface{}) edpt.RuleSetOperRulesByZoneOper {

	count1 := len(d)
	var ret edpt.RuleSetOperRulesByZoneOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GroupList = getSliceRuleSetOperRulesByZoneOperGroupList(in["group_list"].([]interface{}))
	}
	return ret
}

func getSliceRuleSetOperRulesByZoneOperGroupList(d []interface{}) []edpt.RuleSetOperRulesByZoneOperGroupList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperRulesByZoneOperGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperRulesByZoneOperGroupList
		oi.From = in["from"].(string)
		oi.To = in["to"].(string)
		oi.RuleList = getSliceRuleSetOperRulesByZoneOperGroupListRuleList(in["rule_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetOperRulesByZoneOperGroupListRuleList(d []interface{}) []edpt.RuleSetOperRulesByZoneOperGroupListRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperRulesByZoneOperGroupListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperRulesByZoneOperGroupListRuleList
		oi.Name = in["name"].(string)
		oi.Action = in["action"].(string)
		oi.SourceList = getSliceRuleSetOperRulesByZoneOperGroupListRuleListSourceList(in["source_list"].([]interface{}))
		oi.DestList = getSliceRuleSetOperRulesByZoneOperGroupListRuleListDestList(in["dest_list"].([]interface{}))
		oi.ServiceList = getSliceRuleSetOperRulesByZoneOperGroupListRuleListServiceList(in["service_list"].([]interface{}))
		oi.DscpList = getSliceRuleSetOperRulesByZoneOperGroupListRuleListDscpList(in["dscp_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetOperRulesByZoneOperGroupListRuleListSourceList(d []interface{}) []edpt.RuleSetOperRulesByZoneOperGroupListRuleListSourceList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperRulesByZoneOperGroupListRuleListSourceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperRulesByZoneOperGroupListRuleListSourceList
		oi.Source = in["source"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetOperRulesByZoneOperGroupListRuleListDestList(d []interface{}) []edpt.RuleSetOperRulesByZoneOperGroupListRuleListDestList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperRulesByZoneOperGroupListRuleListDestList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperRulesByZoneOperGroupListRuleListDestList
		oi.Dest = in["dest"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetOperRulesByZoneOperGroupListRuleListServiceList(d []interface{}) []edpt.RuleSetOperRulesByZoneOperGroupListRuleListServiceList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperRulesByZoneOperGroupListRuleListServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperRulesByZoneOperGroupListRuleListServiceList
		oi.Service = in["service"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetOperRulesByZoneOperGroupListRuleListDscpList(d []interface{}) []edpt.RuleSetOperRulesByZoneOperGroupListRuleListDscpList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperRulesByZoneOperGroupListRuleListDscpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperRulesByZoneOperGroupListRuleListDscpList
		oi.Dscp = in["dscp"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRuleSetOperTrackAppRuleList(d []interface{}) edpt.RuleSetOperTrackAppRuleList {

	count1 := len(d)
	var ret edpt.RuleSetOperTrackAppRuleList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectRuleSetOperTrackAppRuleListOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectRuleSetOperTrackAppRuleListOper(d []interface{}) edpt.RuleSetOperTrackAppRuleListOper {

	count1 := len(d)
	var ret edpt.RuleSetOperTrackAppRuleListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceRuleSetOperTrackAppRuleListOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceRuleSetOperTrackAppRuleListOperRuleList(d []interface{}) []edpt.RuleSetOperTrackAppRuleListOperRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetOperTrackAppRuleListOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetOperTrackAppRuleListOperRuleList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRuleSetOper(d *schema.ResourceData) edpt.RuleSetOper {
	var ret edpt.RuleSetOper

	ret.Application = getObjectRuleSetOperApplication(d.Get("application").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectRuleSetOperOper(d.Get("oper").([]interface{}))

	ret.RuleList = getSliceRuleSetOperRuleList(d.Get("rule_list").([]interface{}))

	ret.RulesByZone = getObjectRuleSetOperRulesByZone(d.Get("rules_by_zone").([]interface{}))

	ret.TrackAppRuleList = getObjectRuleSetOperTrackAppRuleList(d.Get("track_app_rule_list").([]interface{}))
	return ret
}
