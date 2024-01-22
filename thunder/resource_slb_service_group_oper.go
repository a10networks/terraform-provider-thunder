package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServiceGroupOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_service_group_oper`: Operational Status for the object service-group\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbServiceGroupOperRead,

		Schema: map[string]*schema.Schema{
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Member name",
						},
						"port": {
							Type: schema.TypeInt, Required: true, Description: "Port number",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hm_key": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hm_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drs_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"drs_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"drs_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"drs_hm_key": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_hm_index": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_priority": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_curr_conn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_pers_conn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_total_conn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_curr_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_total_req": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_total_req_succ": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_rev_pkts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_fwd_pkts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_rev_bts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_fwd_bts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_peak_conn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_rsp_time": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_frsp_time": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"drs_srsp_time": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
									"alt_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alt_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"alt_port": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alt_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"alt_curr_conn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alt_total_conn": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alt_rev_pkts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alt_fwd_pkts": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"alt_peak_conn": {
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "SLB Service Name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"servers_up": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"servers_down": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"servers_disable": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"servers_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stateless_current_rate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stateless_current_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stateless_state": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"stateless_type": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hm_dsr_enable_all_vip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pri_affinity_priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"filter": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sgm_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sgm_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"sgm_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceSlbServiceGroupOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroupOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbServiceGroupOperMemberList := setSliceSlbServiceGroupOperMemberList(res)
		d.Set("member_list", SlbServiceGroupOperMemberList)
		SlbServiceGroupOperOper := setObjectSlbServiceGroupOperOper(res)
		d.Set("oper", SlbServiceGroupOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceSlbServiceGroupOperMemberList(d edpt.DataSlbServiceGroupOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtSlbServiceGroupOper.MemberList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["port"] = item.Port
		in["oper"] = setObjectSlbServiceGroupOperMemberListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectSlbServiceGroupOperMemberListOper(d edpt.SlbServiceGroupOperMemberListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["state"] = d.State

	in["hm_key"] = d.HmKey

	in["hm_index"] = d.HmIndex
	in["drs_list"] = setSliceSlbServiceGroupOperMemberListOperDrsList(d.DrsList)
	in["alt_list"] = setSliceSlbServiceGroupOperMemberListOperAltList(d.AltList)
	result = append(result, in)
	return result
}

func setSliceSlbServiceGroupOperMemberListOperDrsList(d []edpt.SlbServiceGroupOperMemberListOperDrsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["drs_name"] = item.DrsName
		in["drs_state"] = item.DrsState
		in["drs_hm_key"] = item.DrsHmKey
		in["drs_hm_index"] = item.DrsHmIndex
		in["drs_port"] = item.DrsPort
		in["drs_priority"] = item.DrsPriority
		in["drs_curr_conn"] = item.DrsCurrConn
		in["drs_pers_conn"] = item.DrsPersConn
		in["drs_total_conn"] = item.DrsTotalConn
		in["drs_curr_req"] = item.DrsCurrReq
		in["drs_total_req"] = item.DrsTotalReq
		in["drs_total_req_succ"] = item.DrsTotalReqSucc
		in["drs_rev_pkts"] = item.DrsRevPkts
		in["drs_fwd_pkts"] = item.DrsFwdPkts
		in["drs_rev_bts"] = item.DrsRevBts
		in["drs_fwd_bts"] = item.DrsFwdBts
		in["drs_peak_conn"] = item.DrsPeakConn
		in["drs_rsp_time"] = item.DrsRspTime
		in["drs_frsp_time"] = item.DrsFrspTime
		in["drs_srsp_time"] = item.DrsSrspTime
		result = append(result, in)
	}
	return result
}

func setSliceSlbServiceGroupOperMemberListOperAltList(d []edpt.SlbServiceGroupOperMemberListOperAltList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["alt_name"] = item.AltName
		in["alt_port"] = item.AltPort
		in["alt_state"] = item.AltState
		in["alt_curr_conn"] = item.AltCurrConn
		in["alt_total_conn"] = item.AltTotalConn
		in["alt_rev_pkts"] = item.AltRevPkts
		in["alt_fwd_pkts"] = item.AltFwdPkts
		in["alt_peak_conn"] = item.AltPeakConn
		result = append(result, in)
	}
	return result
}

func setObjectSlbServiceGroupOperOper(ret edpt.DataSlbServiceGroupOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":                   ret.DtSlbServiceGroupOper.Oper.State,
			"servers_up":              ret.DtSlbServiceGroupOper.Oper.Servers_up,
			"servers_down":            ret.DtSlbServiceGroupOper.Oper.Servers_down,
			"servers_disable":         ret.DtSlbServiceGroupOper.Oper.Servers_disable,
			"servers_total":           ret.DtSlbServiceGroupOper.Oper.Servers_total,
			"stateless_current_rate":  ret.DtSlbServiceGroupOper.Oper.Stateless_current_rate,
			"stateless_current_usage": ret.DtSlbServiceGroupOper.Oper.Stateless_current_usage,
			"stateless_state":         ret.DtSlbServiceGroupOper.Oper.Stateless_state,
			"stateless_type":          ret.DtSlbServiceGroupOper.Oper.Stateless_type,
			"hm_dsr_enable_all_vip":   ret.DtSlbServiceGroupOper.Oper.Hm_dsr_enable_all_vip,
			"pri_affinity_priority":   ret.DtSlbServiceGroupOper.Oper.Pri_affinity_priority,
			"filter":                  ret.DtSlbServiceGroupOper.Oper.Filter,
			"sgm_list":                setSliceSlbServiceGroupOperOperSgmList(ret.DtSlbServiceGroupOper.Oper.SgmList),
		},
	}
}

func setSliceSlbServiceGroupOperOperSgmList(d []edpt.SlbServiceGroupOperOperSgmList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["sgm_name"] = item.SgmName
		in["sgm_port"] = item.SgmPort
		result = append(result, in)
	}
	return result
}

func getSliceSlbServiceGroupOperMemberList(d []interface{}) []edpt.SlbServiceGroupOperMemberList {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupOperMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupOperMemberList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		oi.Oper = getObjectSlbServiceGroupOperMemberListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServiceGroupOperMemberListOper(d []interface{}) edpt.SlbServiceGroupOperMemberListOper {

	count1 := len(d)
	var ret edpt.SlbServiceGroupOperMemberListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.HmKey = in["hm_key"].(int)
		ret.HmIndex = in["hm_index"].(int)
		ret.DrsList = getSliceSlbServiceGroupOperMemberListOperDrsList(in["drs_list"].([]interface{}))
		ret.AltList = getSliceSlbServiceGroupOperMemberListOperAltList(in["alt_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbServiceGroupOperMemberListOperDrsList(d []interface{}) []edpt.SlbServiceGroupOperMemberListOperDrsList {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupOperMemberListOperDrsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupOperMemberListOperDrsList
		oi.DrsName = in["drs_name"].(string)
		oi.DrsState = in["drs_state"].(string)
		oi.DrsHmKey = in["drs_hm_key"].(int)
		oi.DrsHmIndex = in["drs_hm_index"].(int)
		oi.DrsPort = in["drs_port"].(int)
		oi.DrsPriority = in["drs_priority"].(int)
		oi.DrsCurrConn = in["drs_curr_conn"].(int)
		oi.DrsPersConn = in["drs_pers_conn"].(int)
		oi.DrsTotalConn = in["drs_total_conn"].(int)
		oi.DrsCurrReq = in["drs_curr_req"].(int)
		oi.DrsTotalReq = in["drs_total_req"].(int)
		oi.DrsTotalReqSucc = in["drs_total_req_succ"].(int)
		oi.DrsRevPkts = in["drs_rev_pkts"].(int)
		oi.DrsFwdPkts = in["drs_fwd_pkts"].(int)
		oi.DrsRevBts = in["drs_rev_bts"].(int)
		oi.DrsFwdBts = in["drs_fwd_bts"].(int)
		oi.DrsPeakConn = in["drs_peak_conn"].(int)
		oi.DrsRspTime = in["drs_rsp_time"].(int)
		oi.DrsFrspTime = in["drs_frsp_time"].(int)
		oi.DrsSrspTime = in["drs_srsp_time"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServiceGroupOperMemberListOperAltList(d []interface{}) []edpt.SlbServiceGroupOperMemberListOperAltList {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupOperMemberListOperAltList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupOperMemberListOperAltList
		oi.AltName = in["alt_name"].(string)
		oi.AltPort = in["alt_port"].(int)
		oi.AltState = in["alt_state"].(string)
		oi.AltCurrConn = in["alt_curr_conn"].(int)
		oi.AltTotalConn = in["alt_total_conn"].(int)
		oi.AltRevPkts = in["alt_rev_pkts"].(int)
		oi.AltFwdPkts = in["alt_fwd_pkts"].(int)
		oi.AltPeakConn = in["alt_peak_conn"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServiceGroupOperOper(d []interface{}) edpt.SlbServiceGroupOperOper {

	count1 := len(d)
	var ret edpt.SlbServiceGroupOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.Servers_up = in["servers_up"].(int)
		ret.Servers_down = in["servers_down"].(int)
		ret.Servers_disable = in["servers_disable"].(int)
		ret.Servers_total = in["servers_total"].(int)
		ret.Stateless_current_rate = in["stateless_current_rate"].(int)
		ret.Stateless_current_usage = in["stateless_current_usage"].(int)
		ret.Stateless_state = in["stateless_state"].(int)
		ret.Stateless_type = in["stateless_type"].(int)
		ret.Hm_dsr_enable_all_vip = in["hm_dsr_enable_all_vip"].(int)
		ret.Pri_affinity_priority = in["pri_affinity_priority"].(int)
		ret.Filter = in["filter"].(string)
		ret.SgmList = getSliceSlbServiceGroupOperOperSgmList(in["sgm_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbServiceGroupOperOperSgmList(d []interface{}) []edpt.SlbServiceGroupOperOperSgmList {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupOperOperSgmList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupOperOperSgmList
		oi.SgmName = in["sgm_name"].(string)
		oi.SgmPort = in["sgm_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServiceGroupOper(d *schema.ResourceData) edpt.SlbServiceGroupOper {
	var ret edpt.SlbServiceGroupOper

	ret.MemberList = getSliceSlbServiceGroupOperMemberList(d.Get("member_list").([]interface{}))

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectSlbServiceGroupOperOper(d.Get("oper").([]interface{}))
	return ret
}
