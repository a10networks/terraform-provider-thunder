package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSessionsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sessions_oper`: Operational Status for the object sessions\n\n__PLACEHOLDER__",
		ReadContext: resourceSessionsOperRead,
		Schema: map[string]*schema.Schema{
			"ext": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"session_ext_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"alloc": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"free": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"fail": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"cpu_round_robin_fail": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alloc_exceed": {
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
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"forward_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"forward_dest": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"reverse_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"reverse_dest": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peak_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hash": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"flags": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"100ms": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sip_call_id": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"app_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"service_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rserver_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"category_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"duration": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_idx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hash_idx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ddos_total_fwd_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ddos_total_rev_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ddos_total_out_of_order": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ddos_total_zero_window": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ddos_total_retrans": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ddos_current_pkt_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ddos_exceeded_pkt_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"extension_fields_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ext_field_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ext_field_val": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"dns_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"app_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nat_ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"src_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"name_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dest_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nat_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"thread": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"bucket": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"app_category": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"app": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"l4_protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_helper_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fw_ip_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_rule": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_dest_zone": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_src_zone": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_dest_obj": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_src_obj": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_dest_obj_grp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_src_obj_grp": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_dest_rserver": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_src_rserver": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fw_dest_vserver": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"application": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"session_id": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"zone_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sport_rate_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sport_rate_limit_curr": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"src_ipv6_prefix": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"dst_ipv6_prefix": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"check_inside_user": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rev_dest_teid": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"msisdn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"msisdn_val": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"imsi": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"imsi_val": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"gtp_msg_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"gtp_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"full_width": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ext_filter_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"smp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"session_smp_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"alloc": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"free": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alloc_fail": {
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
			"smp_table": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src4": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"src6": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dst4": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"dst6": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"srcport": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"dstport": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"ttl": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"payload": {
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

func resourceSessionsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)

		SetExt := setObjectSessionsOperExt(res)
		SetOpr := setObjectSessionsOperOper(res)
		SetSmp := setObjectSessionsOperSmp(res)
		SetSmpTable := setObjectSessionsOperSmpTable(res)

		d.SetId(obj.GetId())
		d.Set("ext", SetExt)
		d.Set("oper", SetOpr)
		d.Set("smp", SetSmp)
		d.Set("smp_table", SetSmpTable)

		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSessionsOperExt(res edpt.Sessionss) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectSessionsOperExtOper(res.Datasession.Ext.Oper)
	result = append(result, in)
	return result
}

func setObjectSessionsOperExtOper(d edpt.SessionsOperExtOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["session_ext_list"] = setSliceSessionsOperExtOperSessionExtList(d.SessionExtList)
	result = append(result, in)
	return result
}

func setSliceSessionsOperExtOperSessionExtList(d []edpt.SessionsOperExtOperSessionExtList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["alloc"] = item.Alloc
		in["free"] = item.Free
		in["fail"] = item.Fail
		in["cpu_round_robin_fail"] = item.CpuRoundRobinFail
		in["alloc_exceed"] = item.AllocExceed
		result = append(result, in)
	}
	return result
}

func setObjectSessionsOperOper(res edpt.Sessionss) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":            setSliceSessionsOperOperSessionList(res.Datasession.Oper.SessionList),
			"total_sessions":          res.Datasession.Oper.TotalSessions,
			"app_sessions":            res.Datasession.Oper.TotalSessions,
			"filter_type":             res.Datasession.Oper.Filter_type,
			"src_ipv4_addr":           res.Datasession.Oper.SrcIpv4Addr,
			"dst_ipv4_addr":           res.Datasession.Oper.DstIpv4Addr,
			"nat_ipv4_addr":           res.Datasession.Oper.NatIpv4Addr,
			"src_ipv6_addr":           res.Datasession.Oper.SrcIpv6Addr,
			"dst_ipv6_addr":           res.Datasession.Oper.DstIpv6Addr,
			"name_str":                res.Datasession.Oper.NameStr,
			"dest_port":               res.Datasession.Oper.DestPort,
			"src_port":                res.Datasession.Oper.SrcPort,
			"nat_port":                res.Datasession.Oper.NatPort,
			"thread":                  res.Datasession.Oper.Thread,
			"bucket":                  res.Datasession.Oper.Bucket,
			"app_category":            res.Datasession.Oper.AppCategory,
			"app":                     res.Datasession.Oper.App,
			"l4_protocol":             res.Datasession.Oper.L4Protocol,
			"fw_helper_sessions":      res.Datasession.Oper.FwHelperSessions,
			"fw_ip_type":              res.Datasession.Oper.FwIpType,
			"fw_rule":                 res.Datasession.Oper.FwRule,
			"fw_dest_zone":            res.Datasession.Oper.FwDestZone,
			"fw_src_zone":             res.Datasession.Oper.FwSrcZone,
			"fw_dest_obj":             res.Datasession.Oper.FwDestObj,
			"fw_src_obj":              res.Datasession.Oper.FwSrcObj,
			"fw_dest_obj_grp":         res.Datasession.Oper.FwDestObjGrp,
			"fw_src_obj_grp":          res.Datasession.Oper.FwSrcObjGrp,
			"fw_dest_rserver":         res.Datasession.Oper.FwDestRserver,
			"fw_src_rserver":          res.Datasession.Oper.FwSrcRserver,
			"fw_dest_vserver":         res.Datasession.Oper.FwDestVserver,
			"application":             res.Datasession.Oper.Application,
			"session_id":              res.Datasession.Oper.SessionId,
			"zone_name":               res.Datasession.Oper.ZoneName,
			"sport_rate_limit_exceed": res.Datasession.Oper.SportRateLimitExceed,
			"sport_rate_limit_curr":   res.Datasession.Oper.SportRateLimitCurr,
			"src_ipv6_prefix":         res.Datasession.Oper.SrcIpv6Prefix,
			"dst_ipv6_prefix":         res.Datasession.Oper.DstIpv6Prefix,
			"check_inside_user":       res.Datasession.Oper.CheckInsideUser,
			"rev_dest_teid":           res.Datasession.Oper.RevDestTeid,
			"msisdn":                  res.Datasession.Oper.Msisdn,
			"msisdn_val":              res.Datasession.Oper.MsisdnVal,
			"imsi":                    res.Datasession.Oper.Imsi,
			"imsi_val":                res.Datasession.Oper.ImsiVal,
			"gtp_msg_type":            res.Datasession.Oper.GtpMsgType,
			"gtp_version":             res.Datasession.Oper.GtpVersion,
			"full_width":              res.Datasession.Oper.FullWidth,
			"ext_filter_name":         res.Datasession.Oper.ExtFilterName,
		},
	}
}

func setSliceSessionsOperOperSessionList(d []edpt.SessionsOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["protocol"] = item.Protocol
		in["forward_source"] = item.ForwardSource
		in["forward_dest"] = item.ForwardDest
		in["reverse_source"] = item.ReverseSource
		in["reverse_dest"] = item.ReverseDest
		in["rate"] = item.Rate
		in["limit"] = item.Limit
		in["drop"] = item.Drop
		in["peak_rate"] = item.PeakRate
		in["age"] = item.Age
		in["hash"] = item.Hash
		in["flags"] = item.Flags
		in["app_type"] = item.AppType
		in["100ms"] = item.Ms100
		in["sip_call_id"] = item.Sip_call_id
		in["app_name"] = item.AppName
		in["service_name"] = item.ServiceName
		in["rserver_name"] = item.RserverName
		in["category_name"] = item.CategoryName
		in["bytes"] = item.Bytes
		in["duration"] = item.Duration
		in["conn_idx"] = item.Conn_idx
		in["hash_idx"] = item.Hash_idx
		in["ddos_total_fwd_bytes"] = item.Ddos_total_fwd_bytes
		in["ddos_total_rev_bytes"] = item.Ddos_total_rev_bytes
		in["ddos_total_out_of_order"] = item.Ddos_total_out_of_order
		in["ddos_total_zero_window"] = item.Ddos_total_zero_window
		in["ddos_total_retrans"] = item.Ddos_total_retrans
		in["ddos_current_pkt_rate"] = item.Ddos_current_pkt_rate
		in["ddos_exceeded_pkt_rate"] = item.Ddos_exceeded_pkt_rate
		in["extension_fields_list"] = setSliceSessionsOperOperSessionListExtensionFieldsList(item.ExtensionFieldsList)
		in["dns_id"] = item.Dns_id
		result = append(result, in)
	}
	return result
}

func setSliceSessionsOperOperSessionListExtensionFieldsList(d []edpt.SessionsOperOperSessionListExtensionFieldsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ext_field_name"] = item.ExtFieldName
		in["ext_field_val"] = item.ExtFieldVal
		result = append(result, in)
	}
	return result
}

func setObjectSessionsOperSmp(res edpt.Sessionss) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectSessionsOperSmpOper(res.Datasession.Smp.Oper)
	result = append(result, in)
	return result
}

func setObjectSessionsOperSmpOper(res edpt.SessionsOperSmpOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["session_smp_list"] = setSliceSessionsOperSmpOperSessionSmpList(res.SessionSmpList)
	result = append(result, in)
	return result
}

func setSliceSessionsOperSmpOperSessionSmpList(d []edpt.SessionsOperSmpOperSessionSmpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["type"] = item.Type
		in["alloc"] = item.Alloc
		in["free"] = item.Free
		in["alloc_fail"] = item.AllocFail
		result = append(result, in)
	}
	return result
}

func setObjectSessionsOperSmpTable(res edpt.Sessionss) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["oper"] = setObjectSessionsOperSmpTableOper(res.Datasession.SmpTable.Oper)
	result = append(result, in)
	return result
}

func setObjectSessionsOperSmpTableOper(d edpt.SessionsOperSmpTableOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["entry_list"] = setSliceSessionsOperSmpTableOperEntryList(d.EntryList)
	result = append(result, in)
	return result
}

func setSliceSessionsOperSmpTableOperEntryList(d []edpt.SessionsOperSmpTableOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src4"] = item.Src4
		in["src6"] = item.Src6
		in["dst4"] = item.Dst4
		in["dst6"] = item.Dst6
		in["srcport"] = item.Srcport
		in["dstport"] = item.Dstport
		in["ttl"] = item.Ttl
		in["type"] = item.Type
		in["payload"] = item.Payload
		result = append(result, in)
	}
	return result
}

func getObjectSessionsOperExt(d []interface{}) edpt.SessionsOperExt {
	count := len(d)
	var ret edpt.SessionsOperExt
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectSessionsOperExtOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectSessionsOperExtOper(d []interface{}) edpt.SessionsOperExtOper {
	count := len(d)
	var ret edpt.SessionsOperExtOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionExtList = getSliceSessionsOperExtOperSessionExtList(in["session_ext_list"].([]interface{}))
	}
	return ret
}

func getSliceSessionsOperExtOperSessionExtList(d []interface{}) []edpt.SessionsOperExtOperSessionExtList {
	count := len(d)
	ret := make([]edpt.SessionsOperExtOperSessionExtList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsOperExtOperSessionExtList
		oi.Type = in["type"].(string)
		oi.Alloc = in["alloc"].(int)
		oi.Free = in["free"].(int)
		oi.Fail = in["fail"].(int)
		oi.CpuRoundRobinFail = in["cpu_round_robin_fail"].(int)
		oi.AllocExceed = in["alloc_exceed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSessionsOperOper(d []interface{}) edpt.SessionsOperOper {
	count := len(d)
	var ret edpt.SessionsOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceSessionsOperOperSessionList(in["session_list"].([]interface{}))
		ret.TotalSessions = in["total_sessions"].(int)
		ret.AppSessions = in["app_sessions"].(int)
		ret.Filter_type = in["filter_type"].(string)
		ret.SrcIpv4Addr = in["src_ipv4_addr"].(string)
		ret.DstIpv4Addr = in["dst_ipv4_addr"].(string)
		ret.NatIpv4Addr = in["nat_ipv4_addr"].(string)
		ret.SrcIpv6Addr = in["src_ipv6_addr"].(string)
		ret.DstIpv6Addr = in["dst_ipv6_addr"].(string)
		ret.NameStr = in["name_str"].(string)
		ret.DestPort = in["dest_port"].(int)
		ret.SrcPort = in["src_port"].(int)
		ret.NatPort = in["nat_port"].(int)
		ret.Thread = in["thread"].(int)
		ret.Bucket = in["bucket"].(int)
		ret.AppCategory = in["app_category"].(string)
		ret.App = in["app"].(string)
		ret.L4Protocol = in["l4_protocol"].(string)
		ret.FwHelperSessions = in["fw_helper_sessions"].(int)
		ret.FwIpType = in["fw_ip_type"].(string)
		ret.FwRule = in["fw_rule"].(string)
		ret.FwDestZone = in["fw_dest_zone"].(string)
		ret.FwSrcZone = in["fw_src_zone"].(string)
		ret.FwDestObj = in["fw_dest_obj"].(string)
		ret.FwSrcObj = in["fw_src_obj"].(string)
		ret.FwDestObjGrp = in["fw_dest_obj_grp"].(string)
		ret.FwSrcObjGrp = in["fw_src_obj_grp"].(string)
		ret.FwDestRserver = in["fw_dest_rserver"].(string)
		ret.FwSrcRserver = in["fw_src_rserver"].(string)
		ret.FwDestVserver = in["fw_dest_vserver"].(string)
		ret.Application = in["application"].(string)
		ret.SessionId = in["session_id"].(string)
		ret.ZoneName = in["zone_name"].(string)
		ret.SportRateLimitExceed = in["sport_rate_limit_exceed"].(int)
		ret.SportRateLimitCurr = in["sport_rate_limit_curr"].(int)
		ret.SrcIpv6Prefix = in["src_ipv6_prefix"].(string)
		ret.DstIpv6Prefix = in["dst_ipv6_prefix"].(string)
		ret.CheckInsideUser = in["check_inside_user"].(int)
		ret.RevDestTeid = in["rev_dest_teid"].(int)
		ret.Msisdn = in["msisdn"].(int)
		ret.MsisdnVal = in["msisdn_val"].(string)
		ret.Imsi = in["imsi"].(int)
		ret.ImsiVal = in["imsi_val"].(string)
		ret.GtpMsgType = in["gtp_msg_type"].(string)
		ret.GtpVersion = in["gtp_version"].(string)
		ret.FullWidth = in["full_width"].(int)
		ret.ExtFilterName = in["ext_filter_name"].(string)
	}
	return ret
}

func getSliceSessionsOperOperSessionList(d []interface{}) []edpt.SessionsOperOperSessionList {
	count := len(d)
	ret := make([]edpt.SessionsOperOperSessionList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsOperOperSessionList
		oi.Protocol = in["protocol"].(string)
		oi.ForwardSource = in["forward_source"].(string)
		oi.ForwardDest = in["forward_dest"].(string)
		oi.ReverseSource = in["reverse_source"].(string)
		oi.ReverseDest = in["reverse_dest"].(string)
		oi.Rate = in["rate"].(int)
		oi.Limit = in["limit"].(int)
		oi.Drop = in["drop"].(int)
		oi.PeakRate = in["peak_rate"].(int)
		oi.Age = in["age"].(int)
		oi.Hash = in["hash"].(int)
		oi.Flags = in["flags"].(string)
		oi.AppType = in["app_type"].(string)
		oi.Ms100 = in["100ms"].(string)
		oi.Sip_call_id = in["sip_call_id"].(string)
		oi.AppName = in["app_name"].(string)
		oi.ServiceName = in["service_name"].(string)
		oi.RserverName = in["rserver_name"].(string)
		oi.CategoryName = in["category_name"].(string)
		oi.Bytes = in["bytes"].(int)
		oi.Duration = in["duration"].(int)
		oi.Conn_idx = in["conn_idx"].(int)
		oi.Hash_idx = in["hash_idx"].(int)
		oi.Ddos_total_fwd_bytes = in["ddos_total_fwd_bytes"].(int)
		oi.Ddos_total_rev_bytes = in["ddos_total_rev_bytes"].(int)
		oi.Ddos_total_out_of_order = in["ddos_total_out_of_order"].(int)
		oi.Ddos_total_zero_window = in["ddos_total_zero_window"].(int)
		oi.Ddos_total_retrans = in["ddos_total_retrans"].(int)
		oi.Ddos_current_pkt_rate = in["ddos_current_pkt_rate"].(int)
		oi.Ddos_exceeded_pkt_rate = in["ddos_exceeded_pkt_rate"].(int)
		oi.ExtensionFieldsList = getSliceSessionsOperOperSessionListExtensionFieldsList(in["extension_fields_list"].([]interface{}))
		oi.Dns_id = in["dns_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSessionsOperOperSessionListExtensionFieldsList(d []interface{}) []edpt.SessionsOperOperSessionListExtensionFieldsList {
	count := len(d)
	ret := make([]edpt.SessionsOperOperSessionListExtensionFieldsList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsOperOperSessionListExtensionFieldsList
		oi.ExtFieldName = in["ext_field_name"].(string)
		oi.ExtFieldVal = in["ext_field_val"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSessionsOperSmp(d []interface{}) edpt.SessionsOperSmp {
	count := len(d)
	var ret edpt.SessionsOperSmp
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectSessionsOperSmpOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectSessionsOperSmpOper(d []interface{}) edpt.SessionsOperSmpOper {
	count := len(d)
	var ret edpt.SessionsOperSmpOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionSmpList = getSliceSessionsOperSmpOperSessionSmpList(in["session_smp_list"].([]interface{}))
	}
	return ret
}

func getSliceSessionsOperSmpOperSessionSmpList(d []interface{}) []edpt.SessionsOperSmpOperSessionSmpList {
	count := len(d)
	ret := make([]edpt.SessionsOperSmpOperSessionSmpList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsOperSmpOperSessionSmpList
		oi.Type = in["type"].(string)
		oi.Alloc = in["alloc"].(int)
		oi.Free = in["free"].(int)
		oi.AllocFail = in["alloc_fail"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSessionsOperSmpTable(d []interface{}) edpt.SessionsOperSmpTable {
	count := len(d)
	var ret edpt.SessionsOperSmpTable
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectSessionsOperSmpTableOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectSessionsOperSmpTableOper(d []interface{}) edpt.SessionsOperSmpTableOper {
	count := len(d)
	var ret edpt.SessionsOperSmpTableOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceSessionsOperSmpTableOperEntryList(in["entry_list"].([]interface{}))
	}
	return ret
}

func getSliceSessionsOperSmpTableOperEntryList(d []interface{}) []edpt.SessionsOperSmpTableOperEntryList {
	count := len(d)
	ret := make([]edpt.SessionsOperSmpTableOperEntryList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SessionsOperSmpTableOperEntryList
		oi.Src4 = in["src4"].(string)
		oi.Src6 = in["src6"].(string)
		oi.Dst4 = in["dst4"].(string)
		oi.Dst6 = in["dst6"].(string)
		oi.Srcport = in["srcport"].(int)
		oi.Dstport = in["dstport"].(int)
		oi.Ttl = in["ttl"].(int)
		oi.Type = in["type"].(string)
		oi.Payload = in["payload"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSessionsOper(d *schema.ResourceData) edpt.SessionsOper {
	var ret edpt.SessionsOper
	ret.Ext = getObjectSessionsOperExt(d.Get("ext").([]interface{}))
	ret.Oper = getObjectSessionsOperOper(d.Get("oper").([]interface{}))
	ret.Smp = getObjectSessionsOperSmp(d.Get("smp").([]interface{}))
	ret.SmpTable = getObjectSessionsOperSmpTable(d.Get("smp_table").([]interface{}))
	return ret
}
