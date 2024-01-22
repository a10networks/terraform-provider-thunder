package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ServiceGroupMemberOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_service_group_member_oper`: Operational Status for the object member\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6ServiceGroupMemberOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Member name",
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
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port number",
			},
		},
	}
}

func resourceCgnv6ServiceGroupMemberOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ServiceGroupMemberOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ServiceGroupMemberOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6ServiceGroupMemberOperOper := setObjectCgnv6ServiceGroupMemberOperOper(res)
		d.Set("oper", Cgnv6ServiceGroupMemberOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6ServiceGroupMemberOperOper(ret edpt.DataCgnv6ServiceGroupMemberOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"state":    ret.DtCgnv6ServiceGroupMemberOper.Oper.State,
			"hm_key":   ret.DtCgnv6ServiceGroupMemberOper.Oper.HmKey,
			"hm_index": ret.DtCgnv6ServiceGroupMemberOper.Oper.HmIndex,
			"drs_list": setSliceCgnv6ServiceGroupMemberOperOperDrsList(ret.DtCgnv6ServiceGroupMemberOper.Oper.DrsList),
			"alt_list": setSliceCgnv6ServiceGroupMemberOperOperAltList(ret.DtCgnv6ServiceGroupMemberOper.Oper.AltList),
		},
	}
}

func setSliceCgnv6ServiceGroupMemberOperOperDrsList(d []edpt.Cgnv6ServiceGroupMemberOperOperDrsList) []map[string]interface{} {
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

func setSliceCgnv6ServiceGroupMemberOperOperAltList(d []edpt.Cgnv6ServiceGroupMemberOperOperAltList) []map[string]interface{} {
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

func getObjectCgnv6ServiceGroupMemberOperOper(d []interface{}) edpt.Cgnv6ServiceGroupMemberOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6ServiceGroupMemberOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
		ret.HmKey = in["hm_key"].(int)
		ret.HmIndex = in["hm_index"].(int)
		ret.DrsList = getSliceCgnv6ServiceGroupMemberOperOperDrsList(in["drs_list"].([]interface{}))
		ret.AltList = getSliceCgnv6ServiceGroupMemberOperOperAltList(in["alt_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6ServiceGroupMemberOperOperDrsList(d []interface{}) []edpt.Cgnv6ServiceGroupMemberOperOperDrsList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServiceGroupMemberOperOperDrsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServiceGroupMemberOperOperDrsList
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

func getSliceCgnv6ServiceGroupMemberOperOperAltList(d []interface{}) []edpt.Cgnv6ServiceGroupMemberOperOperAltList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6ServiceGroupMemberOperOperAltList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6ServiceGroupMemberOperOperAltList
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

func dataToEndpointCgnv6ServiceGroupMemberOper(d *schema.ResourceData) edpt.Cgnv6ServiceGroupMemberOper {
	var ret edpt.Cgnv6ServiceGroupMemberOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectCgnv6ServiceGroupMemberOperOper(d.Get("oper").([]interface{}))

	ret.Port = d.Get("port").(int)
	return ret
}
