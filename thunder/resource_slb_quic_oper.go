package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbQuicOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_quic_oper`: Operational Status for the object quic\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbQuicOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fwd_source": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fwd_source_cid": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fwd_dest": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fwd_dest_cid": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fwd_state": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fwd_flags": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"fwd_active_scids": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fwd_active_scid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"fwd_available_scids": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fwd_available_scid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"fwd_retired_scids": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fwd_retired_scid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"fwd_active_dcids": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fwd_active_dcid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"fwd_available_dcids": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fwd_available_dcid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"fwd_retired_dcids": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"fwd_retired_dcid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
									"reverse_tuples": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rev_source": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_source_cid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_dest": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_dest_cid": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_flags": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"rev_active_scids": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rev_active_scid": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"rev_available_scids": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rev_available_scid": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"rev_retired_scids": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rev_retired_scid": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"rev_active_dcids": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rev_active_dcid": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"rev_available_dcids": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rev_available_dcid": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"rev_retired_dcids": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"rev_retired_dcid": {
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
						"total_sessions": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbQuicOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbQuicOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbQuicOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbQuicOperOper := setObjectSlbQuicOperOper(res)
		d.Set("oper", SlbQuicOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbQuicOperOper(ret edpt.DataSlbQuicOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_list":   setSliceSlbQuicOperOperSessionList(ret.DtSlbQuicOper.Oper.SessionList),
			"total_sessions": ret.DtSlbQuicOper.Oper.TotalSessions,
		},
	}
}

func setSliceSlbQuicOperOperSessionList(d []edpt.SlbQuicOperOperSessionList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fwd_source"] = item.FwdSource
		in["fwd_source_cid"] = item.FwdSourceCid
		in["fwd_dest"] = item.FwdDest
		in["fwd_dest_cid"] = item.FwdDestCid
		in["fwd_state"] = item.FwdState
		in["fwd_flags"] = item.FwdFlags
		in["fwd_active_scids"] = setSliceSlbQuicOperOperSessionListFwdActiveScids(item.FwdActiveScids)
		in["fwd_available_scids"] = setSliceSlbQuicOperOperSessionListFwdAvailableScids(item.FwdAvailableScids)
		in["fwd_retired_scids"] = setSliceSlbQuicOperOperSessionListFwdRetiredScids(item.FwdRetiredScids)
		in["fwd_active_dcids"] = setSliceSlbQuicOperOperSessionListFwdActiveDcids(item.FwdActiveDcids)
		in["fwd_available_dcids"] = setSliceSlbQuicOperOperSessionListFwdAvailableDcids(item.FwdAvailableDcids)
		in["fwd_retired_dcids"] = setSliceSlbQuicOperOperSessionListFwdRetiredDcids(item.FwdRetiredDcids)
		in["reverse_tuples"] = setSliceSlbQuicOperOperSessionListReverseTuples(item.ReverseTuples)
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListFwdActiveScids(d []edpt.SlbQuicOperOperSessionListFwdActiveScids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fwd_active_scid"] = item.FwdActiveScid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListFwdAvailableScids(d []edpt.SlbQuicOperOperSessionListFwdAvailableScids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fwd_available_scid"] = item.FwdAvailableScid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListFwdRetiredScids(d []edpt.SlbQuicOperOperSessionListFwdRetiredScids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fwd_retired_scid"] = item.FwdRetiredScid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListFwdActiveDcids(d []edpt.SlbQuicOperOperSessionListFwdActiveDcids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fwd_active_dcid"] = item.FwdActiveDcid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListFwdAvailableDcids(d []edpt.SlbQuicOperOperSessionListFwdAvailableDcids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fwd_available_dcid"] = item.FwdAvailableDcid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListFwdRetiredDcids(d []edpt.SlbQuicOperOperSessionListFwdRetiredDcids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["fwd_retired_dcid"] = item.FwdRetiredDcid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListReverseTuples(d []edpt.SlbQuicOperOperSessionListReverseTuples) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rev_source"] = item.RevSource
		in["rev_source_cid"] = item.RevSourceCid
		in["rev_dest"] = item.RevDest
		in["rev_dest_cid"] = item.RevDestCid
		in["rev_state"] = item.RevState
		in["rev_flags"] = item.RevFlags
		in["rev_active_scids"] = setSliceSlbQuicOperOperSessionListReverseTuplesRevActiveScids(item.RevActiveScids)
		in["rev_available_scids"] = setSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableScids(item.RevAvailableScids)
		in["rev_retired_scids"] = setSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredScids(item.RevRetiredScids)
		in["rev_active_dcids"] = setSliceSlbQuicOperOperSessionListReverseTuplesRevActiveDcids(item.RevActiveDcids)
		in["rev_available_dcids"] = setSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableDcids(item.RevAvailableDcids)
		in["rev_retired_dcids"] = setSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredDcids(item.RevRetiredDcids)
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListReverseTuplesRevActiveScids(d []edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveScids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rev_active_scid"] = item.RevActiveScid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableScids(d []edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableScids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rev_available_scid"] = item.RevAvailableScid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredScids(d []edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredScids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rev_retired_scid"] = item.RevRetiredScid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListReverseTuplesRevActiveDcids(d []edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveDcids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rev_active_dcid"] = item.RevActiveDcid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableDcids(d []edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableDcids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rev_available_dcid"] = item.RevAvailableDcid
		result = append(result, in)
	}
	return result
}

func setSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredDcids(d []edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredDcids) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rev_retired_dcid"] = item.RevRetiredDcid
		result = append(result, in)
	}
	return result
}

func getObjectSlbQuicOperOper(d []interface{}) edpt.SlbQuicOperOper {

	count1 := len(d)
	var ret edpt.SlbQuicOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionList = getSliceSlbQuicOperOperSessionList(in["session_list"].([]interface{}))
		ret.TotalSessions = in["total_sessions"].(int)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionList(d []interface{}) []edpt.SlbQuicOperOperSessionList {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionList
		oi.FwdSource = in["fwd_source"].(string)
		oi.FwdSourceCid = in["fwd_source_cid"].(string)
		oi.FwdDest = in["fwd_dest"].(string)
		oi.FwdDestCid = in["fwd_dest_cid"].(string)
		oi.FwdState = in["fwd_state"].(string)
		oi.FwdFlags = in["fwd_flags"].(string)
		oi.FwdActiveScids = getSliceSlbQuicOperOperSessionListFwdActiveScids(in["fwd_active_scids"].([]interface{}))
		oi.FwdAvailableScids = getSliceSlbQuicOperOperSessionListFwdAvailableScids(in["fwd_available_scids"].([]interface{}))
		oi.FwdRetiredScids = getSliceSlbQuicOperOperSessionListFwdRetiredScids(in["fwd_retired_scids"].([]interface{}))
		oi.FwdActiveDcids = getSliceSlbQuicOperOperSessionListFwdActiveDcids(in["fwd_active_dcids"].([]interface{}))
		oi.FwdAvailableDcids = getSliceSlbQuicOperOperSessionListFwdAvailableDcids(in["fwd_available_dcids"].([]interface{}))
		oi.FwdRetiredDcids = getSliceSlbQuicOperOperSessionListFwdRetiredDcids(in["fwd_retired_dcids"].([]interface{}))
		oi.ReverseTuples = getSliceSlbQuicOperOperSessionListReverseTuples(in["reverse_tuples"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListFwdActiveScids(d []interface{}) []edpt.SlbQuicOperOperSessionListFwdActiveScids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListFwdActiveScids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListFwdActiveScids
		oi.FwdActiveScid = in["fwd_active_scid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListFwdAvailableScids(d []interface{}) []edpt.SlbQuicOperOperSessionListFwdAvailableScids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListFwdAvailableScids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListFwdAvailableScids
		oi.FwdAvailableScid = in["fwd_available_scid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListFwdRetiredScids(d []interface{}) []edpt.SlbQuicOperOperSessionListFwdRetiredScids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListFwdRetiredScids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListFwdRetiredScids
		oi.FwdRetiredScid = in["fwd_retired_scid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListFwdActiveDcids(d []interface{}) []edpt.SlbQuicOperOperSessionListFwdActiveDcids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListFwdActiveDcids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListFwdActiveDcids
		oi.FwdActiveDcid = in["fwd_active_dcid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListFwdAvailableDcids(d []interface{}) []edpt.SlbQuicOperOperSessionListFwdAvailableDcids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListFwdAvailableDcids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListFwdAvailableDcids
		oi.FwdAvailableDcid = in["fwd_available_dcid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListFwdRetiredDcids(d []interface{}) []edpt.SlbQuicOperOperSessionListFwdRetiredDcids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListFwdRetiredDcids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListFwdRetiredDcids
		oi.FwdRetiredDcid = in["fwd_retired_dcid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListReverseTuples(d []interface{}) []edpt.SlbQuicOperOperSessionListReverseTuples {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListReverseTuples, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListReverseTuples
		oi.RevSource = in["rev_source"].(string)
		oi.RevSourceCid = in["rev_source_cid"].(string)
		oi.RevDest = in["rev_dest"].(string)
		oi.RevDestCid = in["rev_dest_cid"].(string)
		oi.RevState = in["rev_state"].(string)
		oi.RevFlags = in["rev_flags"].(string)
		oi.RevActiveScids = getSliceSlbQuicOperOperSessionListReverseTuplesRevActiveScids(in["rev_active_scids"].([]interface{}))
		oi.RevAvailableScids = getSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableScids(in["rev_available_scids"].([]interface{}))
		oi.RevRetiredScids = getSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredScids(in["rev_retired_scids"].([]interface{}))
		oi.RevActiveDcids = getSliceSlbQuicOperOperSessionListReverseTuplesRevActiveDcids(in["rev_active_dcids"].([]interface{}))
		oi.RevAvailableDcids = getSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableDcids(in["rev_available_dcids"].([]interface{}))
		oi.RevRetiredDcids = getSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredDcids(in["rev_retired_dcids"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListReverseTuplesRevActiveScids(d []interface{}) []edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveScids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveScids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveScids
		oi.RevActiveScid = in["rev_active_scid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableScids(d []interface{}) []edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableScids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableScids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableScids
		oi.RevAvailableScid = in["rev_available_scid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredScids(d []interface{}) []edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredScids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredScids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredScids
		oi.RevRetiredScid = in["rev_retired_scid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListReverseTuplesRevActiveDcids(d []interface{}) []edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveDcids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveDcids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListReverseTuplesRevActiveDcids
		oi.RevActiveDcid = in["rev_active_dcid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListReverseTuplesRevAvailableDcids(d []interface{}) []edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableDcids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableDcids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListReverseTuplesRevAvailableDcids
		oi.RevAvailableDcid = in["rev_available_dcid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbQuicOperOperSessionListReverseTuplesRevRetiredDcids(d []interface{}) []edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredDcids {

	count1 := len(d)
	ret := make([]edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredDcids, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbQuicOperOperSessionListReverseTuplesRevRetiredDcids
		oi.RevRetiredDcid = in["rev_retired_dcid"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbQuicOper(d *schema.ResourceData) edpt.SlbQuicOper {
	var ret edpt.SlbQuicOper

	ret.Oper = getObjectSlbQuicOperOper(d.Get("oper").([]interface{}))
	return ret
}
