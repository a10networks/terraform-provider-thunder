package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityZbarDestOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_zbar_dest_oper`: Operational Status for the object dest\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityZbarDestOperRead,

		Schema: map[string]*schema.Schema{
			"bad_sources": {
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
									"port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"multi_bad_src_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"src_ip": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"ind_value": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"state": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"drop_cnt": {
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
						"ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"phase": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"tuple_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"zbar_multi_ind_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ind_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ind_total_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"zbar_ind_min": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"zbar_ind_max": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"zbar_ind_slot_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"slot_id": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"start": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"end": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"agg": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"avg": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"src_count": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"zbar_slot_src_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"src_ip": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
															"src_ind_value": {
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
	}
}

func resourceVisibilityZbarDestOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityZbarDestOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityZbarDestOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityZbarDestOperBadSources := setObjectVisibilityZbarDestOperBadSources(res)
		d.Set("bad_sources", VisibilityZbarDestOperBadSources)
		VisibilityZbarDestOperOper := setObjectVisibilityZbarDestOperOper(res)
		d.Set("oper", VisibilityZbarDestOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityZbarDestOperBadSources(ret edpt.DataVisibilityZbarDestOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectVisibilityZbarDestOperBadSourcesOper(ret.DtVisibilityZbarDestOper.BadSources.Oper),
		},
	}
}

func setObjectVisibilityZbarDestOperBadSourcesOper(d edpt.VisibilityZbarDestOperBadSourcesOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ipv4_addr"] = d.Ipv4Addr

	in["ipv6_addr"] = d.Ipv6Addr

	in["port"] = d.Port

	in["protocol"] = d.Protocol
	in["multi_bad_src_list"] = setSliceVisibilityZbarDestOperBadSourcesOperMultiBadSrcList(d.MultiBadSrcList)
	result = append(result, in)
	return result
}

func setSliceVisibilityZbarDestOperBadSourcesOperMultiBadSrcList(d []edpt.VisibilityZbarDestOperBadSourcesOperMultiBadSrcList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_ip"] = item.SrcIp
		in["ind_value"] = item.IndValue
		in["state"] = item.State
		in["drop_cnt"] = item.DropCnt
		result = append(result, in)
	}
	return result
}

func setObjectVisibilityZbarDestOperOper(ret edpt.DataVisibilityZbarDestOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ipv4_addr":           ret.DtVisibilityZbarDestOper.Oper.Ipv4Addr,
			"ipv6_addr":           ret.DtVisibilityZbarDestOper.Oper.Ipv6Addr,
			"port":                ret.DtVisibilityZbarDestOper.Oper.Port,
			"protocol":            ret.DtVisibilityZbarDestOper.Oper.Protocol,
			"phase":               ret.DtVisibilityZbarDestOper.Oper.Phase,
			"tuple_count":         ret.DtVisibilityZbarDestOper.Oper.TupleCount,
			"zbar_multi_ind_list": setSliceVisibilityZbarDestOperOperZbarMultiIndList(ret.DtVisibilityZbarDestOper.Oper.ZbarMultiIndList),
		},
	}
}

func setSliceVisibilityZbarDestOperOperZbarMultiIndList(d []edpt.VisibilityZbarDestOperOperZbarMultiIndList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ind_name"] = item.IndName
		in["ind_total_count"] = item.IndTotalCount
		in["zbar_ind_min"] = item.ZbarIndMin
		in["zbar_ind_max"] = item.ZbarIndMax
		in["zbar_ind_slot_list"] = setSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList(item.ZbarIndSlotList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList(d []edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["slot_id"] = item.SlotId
		in["start"] = item.Start
		in["end"] = item.End
		in["agg"] = item.Agg
		in["avg"] = item.Avg
		in["src_count"] = item.SrcCount
		in["zbar_slot_src_list"] = setSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList(item.ZbarSlotSrcList)
		result = append(result, in)
	}
	return result
}

func setSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList(d []edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["src_ip"] = item.SrcIp
		in["src_ind_value"] = item.SrcIndValue
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityZbarDestOperBadSources(d []interface{}) edpt.VisibilityZbarDestOperBadSources {

	count1 := len(d)
	var ret edpt.VisibilityZbarDestOperBadSources
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectVisibilityZbarDestOperBadSourcesOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectVisibilityZbarDestOperBadSourcesOper(d []interface{}) edpt.VisibilityZbarDestOperBadSourcesOper {

	count1 := len(d)
	var ret edpt.VisibilityZbarDestOperBadSourcesOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.Port = in["port"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.MultiBadSrcList = getSliceVisibilityZbarDestOperBadSourcesOperMultiBadSrcList(in["multi_bad_src_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityZbarDestOperBadSourcesOperMultiBadSrcList(d []interface{}) []edpt.VisibilityZbarDestOperBadSourcesOperMultiBadSrcList {

	count1 := len(d)
	ret := make([]edpt.VisibilityZbarDestOperBadSourcesOperMultiBadSrcList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityZbarDestOperBadSourcesOperMultiBadSrcList
		oi.SrcIp = in["src_ip"].(string)
		oi.IndValue = in["ind_value"].(int)
		oi.State = in["state"].(string)
		oi.DropCnt = in["drop_cnt"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityZbarDestOperOper(d []interface{}) edpt.VisibilityZbarDestOperOper {

	count1 := len(d)
	var ret edpt.VisibilityZbarDestOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.Port = in["port"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.Phase = in["phase"].(string)
		ret.TupleCount = in["tuple_count"].(int)
		ret.ZbarMultiIndList = getSliceVisibilityZbarDestOperOperZbarMultiIndList(in["zbar_multi_ind_list"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityZbarDestOperOperZbarMultiIndList(d []interface{}) []edpt.VisibilityZbarDestOperOperZbarMultiIndList {

	count1 := len(d)
	ret := make([]edpt.VisibilityZbarDestOperOperZbarMultiIndList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityZbarDestOperOperZbarMultiIndList
		oi.IndName = in["ind_name"].(string)
		oi.IndTotalCount = in["ind_total_count"].(int)
		oi.ZbarIndMin = in["zbar_ind_min"].(int)
		oi.ZbarIndMax = in["zbar_ind_max"].(int)
		oi.ZbarIndSlotList = getSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList(in["zbar_ind_slot_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList(d []interface{}) []edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList {

	count1 := len(d)
	ret := make([]edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotList
		oi.SlotId = in["slot_id"].(int)
		oi.Start = in["start"].(int)
		oi.End = in["end"].(int)
		oi.Agg = in["agg"].(int)
		oi.Avg = in["avg"].(int)
		oi.SrcCount = in["src_count"].(int)
		oi.ZbarSlotSrcList = getSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList(in["zbar_slot_src_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList(d []interface{}) []edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList {

	count1 := len(d)
	ret := make([]edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityZbarDestOperOperZbarMultiIndListZbarIndSlotListZbarSlotSrcList
		oi.SrcIp = in["src_ip"].(string)
		oi.SrcIndValue = in["src_ind_value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityZbarDestOper(d *schema.ResourceData) edpt.VisibilityZbarDestOper {
	var ret edpt.VisibilityZbarDestOper

	ret.BadSources = getObjectVisibilityZbarDestOperBadSources(d.Get("bad_sources").([]interface{}))

	ret.Oper = getObjectVisibilityZbarDestOperOper(d.Get("oper").([]interface{}))
	return ret
}
